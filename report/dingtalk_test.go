package report

import (
	"game_go/system"
	"testing"
)

func TestTextMessage_Send(t *testing.T) {

	t.Run("test", func(t *testing.T) {
		receiver := &TextMessage{
			At: AtContent{
				AtMobiles: []string{
					"",
				},
				AtUserIds: nil,
				IsAtAll:   true,
			},
			Text: TextContent{
				Content: "error:我来了",
			},
			MsgType: "text",
		}
		cfg, err := system.NewConfig("../", "keeplive", "ini")
		if err != nil {
			t.Errorf("read config Error: %v", err)
			return
		}
		res := receiver.Send(cfg)
		if res != nil {
			t.Errorf("res: %v", res)
			return
		}
	})
}

func TestSecretInfo_MakeSign(t *testing.T) {
	type fields struct {
		AccessToken string
		Secret      string
	}
	type args struct {
		timestamp int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "sign test",
			fields: fields{
				AccessToken: "",
				Secret:      "abc",
			},
			args: args{
				timestamp: 123,
			},
			want: "RO0k2PnUe%2F79SlwRdbYxibEjLl8mGorWVEKzdJXDR5A%3D",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := &SecretInfo{
				AccessToken: tt.fields.AccessToken,
				Secret:      tt.fields.Secret,
			}
			if got := receiver.MakeSign(tt.args.timestamp); got != tt.want {
				t.Errorf("MakeSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
