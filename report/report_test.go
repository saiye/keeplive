package report

import (
	"game_go/system"
	"github.com/spf13/viper"
	"testing"
)

func TestReportHttpServiceInfo(t *testing.T) {
	cfg, err1 := system.GetCfg("../")
	if err1 != nil {
		t.Errorf("system.NewConfig error = %v, ", err1)
	}
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := HttpServiceInfo(cfg); (err != nil) != tt.wantErr {
				t.Errorf("HttpServiceInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetHttpServiceErrorInfo(t *testing.T) {
	cfg, err1 := system.GetCfg("../")
	if err1 != nil {
		t.Errorf("GetHttpServiceErrorInfo( error = %v, ", err1)
	}
	type args struct {
		cfg *viper.Viper
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test http check",
			args: args{
				cfg: cfg,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetHttpServiceErrorInfo(tt.args.cfg)
			if !(len(got) == 2) {
				t.Errorf("GetHttpServiceErrorInfo() = %v", got)
			}

		})
	}
}
