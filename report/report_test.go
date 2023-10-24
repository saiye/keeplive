package report

import (
	"game_go/system"
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
			if err := ReportHttpServiceInfo(cfg); (err != nil) != tt.wantErr {
				t.Errorf("ReportHttpServiceInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
