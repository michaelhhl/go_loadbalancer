package loadbalancer

import "errors"

var curIndex int
type RoundBalancer struct {
}

func (p *RoundBalancer) GetBalancer(insts [] *Instance) (inst *Instance, err error) {

	if len(insts) == 0 {
		err = errors.New("No Instance")
		return
	}

	rwlock.RLock()
	defer rwlock.RUnlock()
	lens := len(insts)
	if curIndex >= lens {
		curIndex = 0
	}
	inst = insts[curIndex]
	curIndex = (curIndex + 1) % lens
	return
}
