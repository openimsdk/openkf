package test

import (
	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"reflect"
	"testing"
)

func TestGetBool(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-GetBool-1",
			args: args{
				key: "app.debug",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := config.GetBool(tt.args.key); got != tt.want {
				t.Errorf("GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBoolOrDefault(t *testing.T) {
	type args struct {
		key          string
		defaultValue bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-GetBoolOrDefault-1",
			args: args{
				key: "app.debug",
			},
			want: true,
		},
		{
			name: "test-GetBoolOrDefault-2",
			args: args{
				key:          "app.test",
				defaultValue: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := config.GetBoolOrDefault(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetBoolOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-GetInt-1",
			args: args{
				key: "server.max_file_size",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := config.GetInt(tt.args.key); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIntOrDefault(t *testing.T) {
	type args struct {
		key          string
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-GetIntOrDefault-1",
			args: args{
				key: "server.max_file_size",
			},
			want: 10,
		},
		{
			name: "test-GetIntOrDefault-2",
			args: args{
				key:          "server.test",
				defaultValue: 20,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := config.GetIntOrDefault(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetIntOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInterface(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "test-GetInterface-1",
			args: args{
				key: "app.log_file",
			},
			want: "logs/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := config.GetInterface(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-GetString-1",
			args: args{
				key: "app.log_file",
			},
			want: "logs/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := config.GetString(tt.args.key); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStringOrDefault(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-GetStringOrDefault-1",
			args: args{
				key: "app.log_file",
			},
			want: "logs/",
		},
		{
			name: "test-GetStringOrDefault-2",
			args: args{
				key:          "app.test",
				defaultValue: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := config.GetStringOrDefault(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetStringOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
