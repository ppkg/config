package config

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/ppkg/glog"
)

var App AppSetting

type AppSetting struct {
	ConfigService struct {
		Appid   string
		Address string
	}
}

func init() {
	//初始化配置
	_, err := toml.DecodeFile("appsetting.toml", &App)
	if err != nil {
		glog.Error(err)
		log.Fatal(err)
	}
	glog.Info("AppConfigService init...")
	glog.Info(App.ConfigService.Appid, App.ConfigService.Address)
}
