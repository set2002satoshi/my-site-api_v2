package models

import (
	"testing"
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

func TestHistoryBlogModel(t *testing.T) {
	a := &types.AuditTrail{
		Revision:  1,
		CreatedAt: time.Date(2022, 4, 1, 9, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2022, 4, 1, 9, 0, 0, 0, time.UTC),
	}
	type args struct {
		blogId      int
		activeId    int
		userId      int
		title       string
		context     string
		categoryIds []*HistoryCategoryModel
		auditTrail  *types.AuditTrail
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				blogId:      1,
				activeId:    1,
				userId:      1,
				title:       "test title",
				context:     "test context",
				categoryIds: []*HistoryCategoryModel{},
				auditTrail:  a,
			},
			wantErr: true,
		},
		{
			name: "ng blog id (negative number)",
			args: args{
				blogId:      -1,
				activeId:    1,
				userId:      1,
				title:       "test title",
				context:     "test context",
				categoryIds: []*HistoryCategoryModel{},
				auditTrail:  a,
			},
			wantErr: false,
		},
		{
			name: "ng active id (negative number)",
			args: args{
				blogId:      1,
				activeId:    -1,
				userId:      1,
				title:       "test title",
				context:     "test context",
				categoryIds: []*HistoryCategoryModel{},
				auditTrail:  a,
			},
			wantErr: false,
		},
		{
			name: "ng (userId zero)",
			args: args{
				blogId:      1,
				activeId:    1,
				userId:      0,
				title:       "test title",
				context:     "test context",
				categoryIds: []*HistoryCategoryModel{},
				auditTrail:  a,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewHistoryBlogModel(
				tt.args.blogId,
				tt.args.activeId,
				tt.args.userId,
				tt.args.title,
				tt.args.context,
				tt.args.categoryIds,
				int(tt.args.auditTrail.Revision),
				tt.args.auditTrail.CreatedAt,
				tt.args.auditTrail.UpdatedAt,
			); (err != nil) == tt.wantErr {
				t.Errorf("NewHistoryBlogModel() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
