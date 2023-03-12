package models

import (
	"testing"
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

func TestActiveCategoryModel(t *testing.T) {
	a := &types.AuditTrail{
		Revision:  1,
		CreatedAt: time.Date(2022, 4, 1, 9, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2022, 4, 1, 9, 0, 0, 0, time.UTC),
	}
	type args struct {
		categoryId   int
		categoryName string
		auditTrail   *types.AuditTrail
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				categoryId:   1,
				categoryName: "category name",
				auditTrail:   a,
			},
			wantErr: true,
		},
		{
			name: "ng (categoryId negative number)",
			args: args{
				categoryId:   -1,
				categoryName: "category name",
				auditTrail:   a,
			},
			wantErr: false,
		},
		{
			name: "ng (set nil category name)",
			args: args{
				categoryId:   1,
				categoryName: "",
				auditTrail:   a,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewActiveCategoryModel(
				tt.args.categoryId,
				tt.args.categoryName,
				int(tt.args.auditTrail.Revision),
				tt.args.auditTrail.CreatedAt,
				tt.args.auditTrail.UpdatedAt,
			); (err != nil) == tt.wantErr {
				t.Errorf("NewActiveCategoryModel() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
