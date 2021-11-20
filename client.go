package config

import (
	"context"
	"os"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/maybgit/glog"
	"github.com/ppkg/config/proto"
	"github.com/ppkg/pool"
)

var (
	cache *bigcache.BigCache
	pol   pool.Pool
)

func init() {
	cache, _ = bigcache.NewBigCache(bigcache.Config{
		Shards:      1024,
		LifeWindow:  time.Second * 60,
		CleanWindow: time.Microsecond * 1,
	})
	pol, _ = pool.New([]string{App.ConfigService.Address}, pool.DefaultOptions)
	glog.Info("初始化缓存，连接池")
}

//根据key获取配置信息
func Get(key string) (string, error) {
	if v, err := cache.Get(key); err == nil {
		glog.Info("get cache ", App.ConfigService.Appid, key)
		return string(v), nil
	}

	client := getClient()
	defer client.Conn.Close()

	out, err := client.RPC.GetConfigValue(client.Ctx, &proto.ConfigRequest{
		AppName:     App.ConfigService.Appid,
		EnvCode:     os.Getenv("ASPNETCORE_ENVIRONMENT"),
		VersionName: os.Getenv("APPVERSION"),
		Key:         key,
	})
	var value string
	if out != nil && out.Code == 200 {
		value = out.Message
	} else {
		glog.Error(out, err)
	}
	cache.Set(key, []byte(value))
	glog.Info("set cache ", App.ConfigService.Appid, key)

	return value, err
}

func GetString(key string) string {
	str, err := Get(key)
	if err != nil {
		glog.Error("app-config GetString error ", key, err)
	}
	return str
}

type configClient struct {
	Conn pool.Conn
	Ctx  context.Context
	Cf   context.CancelFunc
	RPC  proto.AppconfigClient
}

func getClient() *configClient {
	var err error
	var client configClient
	client.Conn, err = pol.Get()
	if err != nil {
		glog.Error(err)
		return nil
	}
	client.RPC = proto.NewAppconfigClient(client.Conn.Value())
	client.Ctx = context.Background()
	client.Ctx, client.Cf = context.WithTimeout(client.Ctx, time.Second*5)
	return &client
}
