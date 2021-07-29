package config

import (
	"fmt"
	"testing"
	"time"
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
