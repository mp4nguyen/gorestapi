package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"bitbucket.org/restapi/models/serviceRegistryMdl"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:9090", "http service address")

var upgrader = websocket.Upgrader{} // use default options

var registry serviceRegistryMdl.Registry = serviceRegistryMdl.Registry{}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func extractNameVersion(target *url.URL) (name, version string, err error) {
	path := target.Path
	// Trim the leading `/`
	if len(path) > 1 && path[0] == '/' {
		path = path[1:]
	}
	// Explode on `/` and make sure we have at least
	// 2 elements (service name and version)
	tmp := strings.Split(path, "/")
	if len(tmp) < 2 {
		return "", "", fmt.Errorf("Invalid path")
	}
	name, version = tmp[0], tmp[1]
	// Rewrite the request's path without the prefix.
	target.Path = "/" + strings.Join(tmp[2:], "/")
	return name, version, nil
}

// NewMultipleHostReverseProxy creates a reverse proxy that will randomly
// select a host from the passed `targets`
func NewMultipleHostReverseProxy() *httputil.ReverseProxy {
	serviceName := ""
	director := func(req *http.Request) {
		name, version, err := extractNameVersion(req.URL)
		if err != nil {
			log.Print(err)
			return
		}
		serviceName = name + "/" + version
		endpoints, ok := registry[serviceName]
		if !ok {
			log.Printf("Service/Version not found")
			return
		}
		req.URL.Scheme = "http"
		req.URL.Host = endpoints.Servers[rand.Int()%len(endpoints.Servers)].Host
	}

	// return &httputil.ReverseProxy{
	// 	Director: director,
	// }

	return &httputil.ReverseProxy{
		Director: director,
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				println("CALLING PROXY")
				return http.ProxyFromEnvironment(req)
			},
			Dial: func(network, addr string) (net.Conn, error) {
				println("CALLING DIAL ADDR = ", addr)
				conn, err := (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial(network, addr)

				if err != nil {
					println("Error during DIAL:", err.Error())

					registry.RemoveServer(&registry, serviceName, addr)
					// conn, err = (&net.Dialer{
					// 	Timeout:   30 * time.Second,
					// 	KeepAlive: 30 * time.Second,
					// }).Dial(network, "localhost:9090/error")

				}
				return conn, err
			},
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}

}

// NewMultipleHostReverseProxy creates a reverse proxy that will randomly
// select a host from the passed `targets`
var serverIndex int = 0

func nextServerIndex(noOfServer int) {
	if serverIndex == noOfServer {
		serverIndex = 0
	} else {
		serverIndex++
	}
}

//
// func NewMultipleHostReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
//
// 	director := func(req *http.Request) {
// 		println("CALLING DIRECTOR")
// 		target := targets[serverIndex]
// 		req.URL.Scheme = target.Scheme
// 		req.URL.Host = target.Host
// 		req.URL.Path = target.Path
// 		nextServerIndex(len(targets) - 1)
// 	}
// 	return &httputil.ReverseProxy{
// 		Director: director,
// 		Transport: &http.Transport{
// 			Proxy: func(req *http.Request) (*url.URL, error) {
// 				println("CALLING PROXY")
// 				return http.ProxyFromEnvironment(req)
// 			},
// 			Dial: func(network, addr string) (net.Conn, error) {
// 				println("CALLING DIAL ADDR = ", addr)
// 				conn, err := (&net.Dialer{
// 					Timeout:   30 * time.Second,
// 					KeepAlive: 30 * time.Second,
// 				}).Dial(network, addr)
//
// 				if err != nil {
// 					println("Error during DIAL:", err.Error())
//
// 					conn, err = (&net.Dialer{
// 						Timeout:   30 * time.Second,
// 						KeepAlive: 30 * time.Second,
// 					}).Dial(network, targets[serverIndex].Host)
//
// 				}
// 				return conn, err
// 			},
// 			TLSHandshakeTimeout: 10 * time.Second,
// 		},
// 	}
// }

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func addServerHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery, " : ", r.URL.Query().Get("serviceName"), " ", r.URL.Query().Get("serverPath"))
	registry.AddServer(&registry, r.URL.Query().Get("serviceName"), r.URL.Query().Get("serverPath"))
	fmt.Println("registry =", registry)
	fmt.Fprintln(w, "OK")
}

func removeServerHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery, " : ", r.URL.Query().Get("serviceName"), " ", r.URL.Query().Get("serverPath"))
	registry.RemoveServer(&registry, r.URL.Query().Get("serviceName"), r.URL.Query().Get("serverPath"))
	fmt.Println("registry =", registry)
	fmt.Fprintln(w, "OK")
}

func main() {
	proxy := NewMultipleHostReverseProxy()
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/home", home)
	http.HandleFunc("/addServer", addServerHandleFunc)
	http.HandleFunc("/removeServer", addServerHandleFunc)
	http.HandleFunc("/", Logger(proxy))
	fmt.Println("Proxy Server is running on port : 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))

}

func Logger(inner *httputil.ReverseProxy) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<head>
<meta charset="utf-8">
<script>
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server,
"Send" to send a message to the server and "Close" to close the connection.
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
