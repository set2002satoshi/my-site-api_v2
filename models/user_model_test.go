package models

import (
	"testing"
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

func TestActiveUserModel(t *testing.T) {
	a := &types.AuditTrail{
		Revision:  1,
		CreatedAt: time.Date(2022, 4, 1, 9, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2022, 4, 1, 9, 0, 0, 0, time.UTC),
	}
	i := &types.ImageTypeFileOrURL{
		ImgFile:      nil,
		ImgURL:       "http://iamge.iamge",
		DataTypeFlag: true,
	}
	type args struct {
		userId     int
		nickname   string
		email      string
		password   string
		icon       *types.ImageTypeFileOrURL
		roll       string
		auditTrail *types.AuditTrail
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				userId:     1,
				nickname:   "ニックネーム",
				email:      "abc@a.com",
				password:   "password",
				icon:       i,
				roll:       "member",
				auditTrail: a,
			},
			wantErr: true,
		},
		{
			name: "ng",
			args: args{
				userId:     -1,
				nickname:   "",
				email:      "",
				password:   "",
				icon:       i,
				roll:       "",
				auditTrail: a,
			},
			wantErr: false,
		},
		{
			name: "ng valid roll",
			args: args{
				userId:     1,
				nickname:   "ニックネーム",
				email:      "abc@a.com",
				password:   "password",
				icon:       i,
				roll:       "aaaa",
				auditTrail: a,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewActiveUserModel(
				tt.args.userId,
				tt.args.nickname,
				tt.args.email,
				tt.args.password,
				tt.args.icon.ImgFile,
				tt.args.icon.ImgURL,
				tt.args.icon.DataTypeFlag,
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
