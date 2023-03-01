package models

import (
	"testing"
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

func TestLogUserModel(t *testing.T) {
	a := &types.AuditTrail{
		Revision:  1,
		CreatedAt: time.Date(2022, 4, 1, 9, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2022, 4, 1, 9, 0, 0, 0, time.UTC),
	}
	type args struct {
		historyUserId int
		activeUserId  int
		nickname      string
		email         string
		password      string
		icon          string
		roll          string
		auditTrail    *types.AuditTrail
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				historyUserId: 1,
				activeUserId:  1,
				nickname:      "ニックネーム",
				email:         "abc@a.com",
				password:      "password",
				icon:          "http://iamge.iamge",
				roll:          "member",
				auditTrail:    a,
			},
			wantErr: true,
		},
		{
			name: "ng",
			args: args{
				historyUserId: -1,
				activeUserId:  -1,
				nickname:      "",
				email:         "",
				password:      "",
				icon:          "http://iamge.iamge",
				roll:          "member",
				auditTrail:    a,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewLogUserModel(
				tt.args.historyUserId,
				tt.args.activeUserId,
				tt.args.nickname,
				tt.args.email,
				tt.args.password,
				tt.args.icon,
				tt.args.roll,
				int(tt.args.auditTrail.Revision),
				tt.args.auditTrail.CreatedAt,
				tt.args.auditTrail.UpdatedAt,
			); (err != nil) == tt.wantErr {
				t.Errorf("NewActiveUserModel() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
