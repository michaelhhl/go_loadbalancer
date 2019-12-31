package loadbalancer

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"lb_component/theviper"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var insts []*Instance
var lb_algorithm string
var rwlock sync.RWMutex

func init()  {
	vp := theviper.GetConfig()
	lb_port := vp.GetInt("lb_port")
	lb_algorithm = vp.GetString("lb_algorithm")
	lb_hosts := vp.GetStringSlice("lb_hosts")
	rwlock.Lock()
	for _, host := range lb_hosts{
		one := NewInstance(host, lb_port)
		insts = append(insts, one)
	}
	rwlock.Unlock()

	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		lb_hosts := vp.GetStringSlice("lb_hosts")
		for _, host := range lb_hosts{
			fmt.Println("the host is:", host)
		}
		refresh(vp)
	})
}

func GetBalancerHostPort() (inst *Instance) {
	var balancer Balancer
	switch lb_algorithm {
	case "random":
		balancer = &RandomBalancer{}
	case "round":
		balancer = &RoundBalancer{}
	}
	inst, err := balancer.GetBalancer(insts)
	if err != nil {
		log.Fatal("Get balance instance failed: %v", err)
	}
	return
}

func refresh(vp *viper.Viper)  {
	lb_port := vp.GetInt("lb_port")
	lb_hosts := vp.GetStringSlice("lb_hosts")
	rwlock.Lock()
	insts = insts[0:0]
	for _, host := range lb_hosts{
		one := NewInstance(host, lb_port)
		insts = append(insts, one)
	}
	rwlock.Unlock()
}
