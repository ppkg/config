package config

import (
	"context"
	"os"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/maybgit/glog"
	"github.com/tricobbler/config/proto"
	"google.golang.org/grpc"
)

var cache *bigcache.BigCache

func init() {
	cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(time.Second * 60))
}

//根据key获取配置信息
func Get(key string) (string, error) {
	if v, err := cache.Get(key); err == nil {
		glog.Info("get cache ", App.AppConfigService.Appid, key)
		return string(v), nil
	}

	client := getClient()
	defer client.close()
	out, err := client.RPC.GetConfigValue(client.Ctx, &proto.ConfigRequest{
		AppName:     App.AppConfigService.Appid,
		EnvCode:     os.Getenv("ASPNETCORE_ENVIRONMENT"),
		VersionName: os.Getenv("APPVERSION"),
		Key:         key,
	})
	var value string
	if out != nil && out.Code == 200 {
		value = out.Message
	}
	cache.Set(key, []byte(value))
	glog.Info("set cache ", App.AppConfigService.Appid, key)

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
	Conn *grpc.ClientConn
	Ctx  context.Context
	Cf   context.CancelFunc
	RPC  proto.AppConfigClient
}

func getClient() *configClient {
	var client configClient
	var err error
	if client.Conn, err = grpc.Dial(App.AppConfigService.Address, grpc.WithInsecure(), grpc.WithBlock()); err != nil {
		glog.Error(err)
		return nil
	} else {
		client.RPC = proto.NewAppConfigClient(client.Conn)
		client.Ctx = context.Background()
		client.Ctx, client.Cf = context.WithTimeout(client.Ctx, time.Second*5)
		return &client
	}
}

//关闭链接
func (c *configClient) close() {
	c.Conn.Close()
	c.Cf()
}
