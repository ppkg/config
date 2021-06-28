package config

import (
	"testing"
	"time"
)

func TestGetString(t *testing.T) {
	time.Sleep(time.Second*3)
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
				key: "redis.Addr",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.args.key); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}
