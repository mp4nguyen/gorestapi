package serviceRegistryMdl

import "log"

type Registry map[string]*ServerGroup

type ServerGroup struct {
	Servers     []*Server
	DiedServers []*Server
}

type Server struct {
	Host   string
	IsLive bool
}

func (r Registry) AddServer(registry *Registry, serviceName string, serverPath string) {
	log.Println("AddServer: serviceName = ", serviceName, " serverPath = ", serverPath)
	service, ok := (*registry)[serviceName] //*registry[serviceName]
	if !ok {
		serverGroup := ServerGroup{Servers: []*Server{&Server{Host: serverPath, IsLive: true}}, DiedServers: []*Server{}}
		(*registry)[serviceName] = &serverGroup
	} else {
		service.Servers = append(service.Servers, &Server{Host: serverPath, IsLive: true})
	}
}

func removeServerArray(s []*Server, i int) []*Server {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (r Registry) RemoveServer(registry *Registry, serviceName string, serverPath string) {
	log.Println("RemoveServer: serviceName = ", serviceName, " serverPath = ", serverPath)
	service, ok := (*registry)[serviceName]
	if ok {
		for index, server := range service.Servers {
			if server.Host == serverPath {
				service.Servers = removeServerArray(service.Servers, index)
				break
			}
		}
	}
	if len(service.Servers) == 0 {
		delete(*registry, serviceName)
	}
}
