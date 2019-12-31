package loadbalancer

import "strconv"

type Instance struct {
	Host string    `json:"host"`
	Port int       `json:"port"`
}

func NewInstance(host string, port int) *Instance {
	return &Instance{
		Host: host,
		Port: port,
	}
}

func (p *Instance) GetHost() string {
	return p.Host
}

func (p *Instance) GetPort() int {
	return p.Port
}

func (p *Instance) String() string {
	return p.Host + ":" + strconv.Itoa(p.Port)
}
