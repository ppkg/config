package config

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/maybgit/glog"
)

var App AppSetting

type AppSetting struct {
	ConfigService struct {
		Appid   string
		Address string
	}
}

func init() {
	glog.Info("AppConfigService init...")
	glog.Info(App.ConfigService.Appid, App.ConfigService.Address)
	//初始化配置
	_, err := toml.DecodeFile("appsetting.toml", &App)
	if err != nil {
		glog.Error(err)
	}
}
