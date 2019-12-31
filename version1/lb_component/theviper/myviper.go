package theviper

import (
	"github.com/spf13/viper"
	"sync"
)

var config_viper *viper.Viper
var rwlock sync.RWMutex

func init()  {
	vp := viper.New()
	vp.SetConfigType("yaml")
	vp.AddConfigPath("D:/GoProject/src/lb_component/config/")
	//vp.SetConfigFile("lb_config.yaml")
	vp.SetConfigName("lb_config")
	ReadConfig(vp)

	config_viper = vp
}

func ReadConfig(vp *viper.Viper) error {
	rwlock.Lock()
	defer rwlock.Unlock()
	err := vp.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func GetConfig() *viper.Viper {
	return config_viper
}


