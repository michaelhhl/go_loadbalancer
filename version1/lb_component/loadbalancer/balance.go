package loadbalancer

type Balancer interface {
	GetBalancer([] *Instance) (*Instance, error)
}


