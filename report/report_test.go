package report

import "testing"

func TestReportMessage(t *testing.T) {
	type args struct {
		configDir string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "system info",
			args: args{
				configDir: "../",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReportMessage(tt.args.configDir)
		})
	}
}
