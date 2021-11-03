package config

import (
	"fmt"
	"testing"
	"time"

	"github.com/allegro/bigcache/v3"
)

func TestGetString(t *testing.T) {
	time.Sleep(time.Second * 3)
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "获取配置",
			args: args{
				key: "bj_base_url",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetString(tt.args.key)
			fmt.Println(got)
		})
	}
}

func TestBigcache(t *testing.T) {
	c, err := bigcache.NewBigCache(bigcache.Config{
		LifeWindow:  time.Second * 5,
		CleanWindow: time.Second*5,
		Shards:      1024,
		// OnRemove: func(key string, entry []byte) {
		// 	t.Log("OnRemove", key)
		// },
		OnRemoveWithReason: func(key string, entry []byte, reason bigcache.RemoveReason) {
			t.Log("OnRemoveWithReason", key)
		},
	})
	t.Log(err)
	// t.Log(c==nil)
	// return

	c.Set("a", []byte("a"))
	c.Set("b", []byte("b"))
	
	go func() {
		for {
			a, _ := c.Get("a")
			b, _ := c.Get("b")
			t.Log(a, b)
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		time.Sleep(time.Second * 3)
		c.Delete("b")
	}()

	time.Sleep(time.Second * 20)
}
