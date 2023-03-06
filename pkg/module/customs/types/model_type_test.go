package types

import (
	"testing"
	"time"
)

func TestAuditTrail(t *testing.T) {
	type args struct {
		revision  int
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok",
			args: args{
				revision:  1,
				createdAt: time.Date(2022, 4, 1, 9, 0, 0, 0, time.UTC),
				updatedAt: time.Date(2022, 4, 1, 9, 0, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "ng negative number",
			args: args{
				revision:  -1,
				createdAt: time.Date(2022, 4, 1, 9, 0, 0, 0, time.UTC),
				updatedAt: time.Date(2022, 4, 1, 9, 0, 0, 0, time.UTC),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewAuditTrail(
				tt.args.revision,
				tt.args.createdAt,
				tt.args.updatedAt,
			); (err != nil) == tt.want {
				t.Errorf("NewAuditTrail() = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestImageTypeFileOrURL(t *testing.T) {
	type args struct {
		// ImgFile *multipart.FileHeader,
		ImgURL       string
		ImgKey       string
		DataTypeFlag bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok",
			args: args{
				// ImgFile:      nil,
				ImgURL:       "http://iamge.iamge",
				ImgKey:       "iam_key",
				DataTypeFlag: true,
			},
			want: true,
		},
		{
			name: "ng valid flag",
			args: args{
				// ImgFile:      nil,
				ImgURL:       "http://iamge.iamge",
				ImgKey:       "iam_key",
				DataTypeFlag: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewImageTypeFileOrURL(
				nil,
				tt.args.ImgURL,
				tt.args.ImgKey,
				tt.args.DataTypeFlag,
			); (err != nil) == tt.want {
				t.Errorf("NewAuditTrail() = %v, want %v", err, tt.want)
			}
		})
	}
}
