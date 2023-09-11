package utils

import (
	"github.com/gofrs/uuid"
	"testing"
)

func TestGenUUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test-GenUUID",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenUUID(); got == tt.want {
				t.Errorf("GenUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenUUIDWithoutHyphen(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test-GenUUIDWithoutHyphen",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenUUIDWithoutHyphen(); got == tt.want {
				t.Errorf("GenUUIDWithoutHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeCanonical(t *testing.T) {
	type args struct {
		dst []byte
		u   uuid.UUID
	}
	var buf [36]byte
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-encodeCanonical",
			args: args{
				dst: buf[:],
				u:   uuid.Must(uuid.NewV4()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodeCanonical(tt.args.dst, tt.args.u)
		})
	}
}

func Test_toString(t *testing.T) {
	type args struct {
		u uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-toString",
			args: args{
				u: uuid.Must(uuid.NewV4()),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toString(tt.args.u); got == tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComparePassword(t *testing.T) {
	type args struct {
		password          string
		encryptedPassword string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-ComparePassword",
			args: args{
				password:          "123456",
				encryptedPassword: "fa9ffa9a6017dee3be4e4063d108b69b54c4c8b3e9465a65b8d3191481655141",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComparePassword(tt.args.password, tt.args.encryptedPassword); got != tt.want {
				t.Errorf("ComparePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-EncryptPassword",
			args: args{
				password: "123456",
			},
			want: "fa9ffa9a6017dee3be4e4063d108b69b54c4c8b3e9465a65b8d3191481655141",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptPassword(tt.args.password); got != tt.want {
				t.Errorf("EncryptPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToString(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-IntToString",
			args: args{
				i: 1,
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToString(tt.args.i); got != tt.want {
				t.Errorf("IntToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-IsValidEmail",
			args: args{
				email: "123@qq.com",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmail(tt.args.email); got != tt.want {
				t.Errorf("IsValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-StringToInt",
			args: args{
				i: "1",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInt(tt.args.i); got != tt.want {
				t.Errorf("StringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
