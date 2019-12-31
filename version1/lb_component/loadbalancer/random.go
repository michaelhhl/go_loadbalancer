package loadbalancer

import (
	"errors"
	"math/rand"
)

type RandomBalancer struct {}

func (p *RandomBalancer) GetBalancer(insts [] *Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No Instance")
		return
	}

	rwlock.RLock()
	defer rwlock.RUnlock()
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]
	return
}
