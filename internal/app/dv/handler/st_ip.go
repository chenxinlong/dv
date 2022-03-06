package handler

import (
	"io/ioutil"
	"net"
	"net/http"

	aw "github.com/deanishe/awgo"
)

const (
	HandlerIP = "ip"
)

type IPHandler struct {
	typ HandlerType
}

func NewHandlerIP() *IPHandler {
	return &IPHandler{
		typ: HandlerTypShortTag,
	}
}

func (g IPHandler) GetType() HandlerType {
	return g.typ
}

func (g IPHandler) Handle(e Event) (item *aw.Item) {
	intranetIP := g.GetIntranetIP()
	remoteIP := g.GetRemoteIP()
	e.wf.NewItem("[local] : " + intranetIP).Copytext(intranetIP)
	e.wf.NewItem("[remote] : " + remoteIP).Copytext(remoteIP)

	return
}

func (g IPHandler) GetIntranetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "get intranet ip failed, err = " + err.Error()
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func (g IPHandler) GetRemoteIP() string {
	resp, err := http.Get("https://myip.ipip.net")
	if err != nil {
		return "get remote ip failed, err = " + err.Error()
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "get remote ip failed, err = " + err.Error()
	}

	return string(b)
}
