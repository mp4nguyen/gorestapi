package websocketCtrl

import (
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.org/restapi/utils"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

type socketIn struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func Socket(w http.ResponseWriter, r *http.Request) {
	in := socketIn{}
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
		log.Printf("recv: %s\n", message)
		err = json.Unmarshal(message, &in)
		log.Println(" in obj = ", in)
		utils.LogError("Unmarshal socket ", err)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
