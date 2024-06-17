package config

import "testing"

func TestInitConfig1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{"default",
			args{"../../config/default.conf.json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitConfig(tt.args.path)
		})
	}
}
