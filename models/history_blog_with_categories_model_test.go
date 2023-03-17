package models

import (
	"testing"
)

func TestHistoryBlogWithCategoryModel(t *testing.T) {
	type args struct {
		id         int
		activeId   int
		categoryId int
		blogId     int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				id:         1,
				activeId:   1,
				categoryId: 1,
				blogId:     1,
			},
			wantErr: true,
		},
		{
			name: "ng (id negative number)",
			args: args{
				id:         -1,
				activeId:   1,
				categoryId: 1,
				blogId:     1,
			},
			wantErr: false,
		},
		{
			name: "ng (activeId negative number)",
			args: args{
				id:         1,
				activeId:   -1,
				categoryId: 1,
				blogId:     1,
			},
			wantErr: false,
		},
		{
			name: "ng (categoryId negative number)",
			args: args{
				id:         1,
				activeId:   1,
				categoryId: -1,
				blogId:     1,
			},
			wantErr: false,
		},
		{
			name: "ng (blogId negative number)",
			args: args{
				id:         1,
				activeId:   1,
				categoryId: 1,
				blogId:     -1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewHistoryBlogWithCategoryModel(
				tt.args.id,
				tt.args.activeId,
				tt.args.categoryId,
				tt.args.blogId,
			); (err != nil) == tt.wantErr {
				t.Errorf("NewHistoryBlogWithCategoryModel() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
