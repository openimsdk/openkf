package test

import (
	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestInitLogger(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试-InitLogger"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.Config.App.Debug = false
			log.InitLogger()
			config.Config.App.Debug = true
			log.InitLogger()
			t.Run("测试-GetLogger", GetLogger)
			t.Run("测试-Debug", Debug)
			t.Run("测试-DebugFormat", DebugFormat)
			t.Run("测试-Info", Info)
			t.Run("测试-InfoFormat", InfoFormat)
			t.Run("测试-Error", Error)
			t.Run("测试-ErrorFormat", ErrorFormat)
			t.Run("测试-Panic", Panic)
			t.Run("测试-PanicFormat", PanicFormat)
		})
	}
}
func GetLogger(t *testing.T) {
	tests := []struct {
		name string
		want *logrus.Logger
	}{
		{
			name: "测试-GetLogger",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := log.GetLogger(); got == tt.want {
				t.Errorf("GetLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Debug(t *testing.T) {
	type args struct {
		Operation string
		args      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "测试-Debug-1",
			args: args{
				Operation: "test-debug-1",
				args:      []interface{}{},
			},
		},
		{
			name: "测试-Debug-2",
			args: args{
				Operation: "test-debug-2",
				args:      []interface{}{"arg2-1", "arg2-2"},
			},
		},
		{
			name: "测试-Debug-3",
			args: args{
				Operation: "test-debug-3",
				args:      []interface{}{"arg3", 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Debug(tt.args.Operation, tt.args.args...)
		})
	}
}

func DebugFormat(t *testing.T) {
	type args struct {
		Operation string
		format    string
		args      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{

		{
			name: "测试-DebugFormat-1",
			args: args{
				Operation: "test-DebugFormat-1",
				format:    "",
				args:      []interface{}{},
			},
		},
		{
			name: "测试-DebugFormat-2",
			args: args{
				Operation: "test-DebugFormat-2",
				format:    "%+v,%+v",
				args:      []interface{}{"arg2-1", "arg2-2"},
			},
		},
		{
			name: "测试-DebugFormat-3",
			args: args{
				Operation: "test-DebugFormat-3",
				format:    "%+v,%+v",
				args:      []interface{}{"arg3", 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Debugf(tt.args.Operation, tt.args.format, tt.args.args...)
		})
	}
}

func Error(t *testing.T) {
	type args struct {
		Operation string
		args      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{

		{
			name: "测试-Error-1",
			args: args{
				Operation: "test-Error-1",
				args:      []interface{}{},
			},
		},
		{
			name: "测试-Error-2",
			args: args{
				Operation: "test-Error-2",
				args:      []interface{}{"arg2-1", "arg2-2"},
			},
		},
		{
			name: "测试-Error-3",
			args: args{
				Operation: "test-Error-3",
				args:      []interface{}{"arg3", 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Error(tt.args.Operation, tt.args.args...)
		})
	}
}

func ErrorFormat(t *testing.T) {
	type args struct {
		Operation string
		format    string
		args      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{

		{
			name: "测试-ErrorFormat-1",
			args: args{
				Operation: "test-ErrorFormat-1",
				format:    "",
				args:      []interface{}{},
			},
		},
		{
			name: "测试-ErrorFormat-2",
			args: args{
				Operation: "test-ErrorFormat-2",
				format:    "%+v,%+v",
				args:      []interface{}{"arg2-1", "arg2-2"},
			},
		},
		{
			name: "测试-ErrorFormat-3",
			args: args{
				Operation: "test-ErrorFormat-3",
				format:    "%+v,%+v",
				args:      []interface{}{"arg3", 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Errorf(tt.args.Operation, tt.args.format, tt.args.args...)
		})
	}
}

func Info(t *testing.T) {
	type args struct {
		Operation string
		args      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{

		{
			name: "测试-Info-1",
			args: args{
				Operation: "test-Info-1",
				args:      []interface{}{},
			},
		},
		{
			name: "测试-Info-2",
			args: args{
				Operation: "test-Info-2",
				args:      []interface{}{"arg2-1", "arg2-2"},
			},
		},
		{
			name: "测试-Info-3",
			args: args{
				Operation: "test-Info-3",
				args:      []interface{}{"arg3", 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Info(tt.args.Operation, tt.args.args...)
		})
	}
}

func InfoFormat(t *testing.T) {
	type args struct {
		Operation string
		format    string
		args      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{

		{
			name: "测试-InfoFormat-1",
			args: args{
				Operation: "test-InfoFormat-1",
				format:    "",
				args:      []interface{}{},
			},
		},
		{
			name: "测试-InfoFormat-2",
			args: args{
				Operation: "test-InfoFormat-2",
				format:    "%+v,%+v",
				args:      []interface{}{"arg2-1", "arg2-2"},
			},
		},
		{
			name: "测试-InfoFormat-3",
			args: args{
				Operation: "test-InfoFormat-3",
				format:    "%+v,%+v",
				args:      []interface{}{"arg3", 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Infof(tt.args.Operation, tt.args.format, tt.args.args...)
		})
	}
}

func Panic(t *testing.T) {
	type args struct {
		Operation string
		args      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{

		{
			name: "测试-Panic-1",
			args: args{
				Operation: "test-Panic-1",
				args:      []interface{}{},
			},
		},
		{
			name: "测试-Panic-2",
			args: args{
				Operation: "test-Panic-2",
				args:      []interface{}{"arg2-1", "arg2-2"},
			},
		},
		{
			name: "测试-Panic-3",
			args: args{
				Operation: "test-Panic-3",
				args:      []interface{}{"arg3", 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			func() {
				defer func() {
					if recover() != nil {
						return
					}
				}()
				log.Panic(tt.args.Operation, tt.args.args...)

			}()
		})
	}
}

func PanicFormat(t *testing.T) {
	type args struct {
		Operation string
		format    string
		args      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{

		{
			name: "测试-PanicFormat-1",
			args: args{
				Operation: "test-PanicFormat-1",
				format:    "",
				args:      []interface{}{},
			},
		},
		{
			name: "测试-PanicFormat-2",
			args: args{
				Operation: "test-PanicFormat-2",
				format:    "%+v,%+v",
				args:      []interface{}{"arg2-1", "arg2-2"},
			},
		},
		{
			name: "测试-PanicFormat-3",
			args: args{
				Operation: "test-PanicFormat-3",
				format:    "%+v,%+v",
				args:      []interface{}{"arg3", 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			func() {
				defer func() {
					if recover() != nil {
						return
					}
				}()
				log.Panicf(tt.args.Operation, tt.args.format, tt.args.args...)
			}()
		})
	}
}
