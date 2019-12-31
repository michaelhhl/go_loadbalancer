package main

import (
	"encoding/json"
	"fmt"
	"lb_component/loadbalancer"
	"lb_component/theviper"
	"net/http"
)

var local_host_port string

func init()  {
	vp := theviper.GetConfig()
	local_host_port = vp.GetString("local_host_port")
}

func getTimeServiceServer(w http.ResponseWriter, r *http.Request)  {
	instance := loadbalancer.GetBalancerHostPort()
	if jsonStr, err := json.Marshal(instance); err == nil {
		fmt.Fprintln(w, string(jsonStr))
	}
}
func main()  {

	http.HandleFunc("/timeservice", getTimeServiceServer)
	http.ListenAndServe(local_host_port, nil)
	//i := 0
	//for {
	//	instance := loadbalancer.GetBalancerHostPort()
	//	fmt.Printf(" the balancer %d is: %s \n", i, instance)
	//	i = i + 1
	//	time.Sleep(time.Second * 5)
	//}
}
