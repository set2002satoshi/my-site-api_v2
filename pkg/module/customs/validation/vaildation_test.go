package validation

import "testing"

func TestValidEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				email: "ok@example.com",
			},
			wantErr: true,
		},
		{
			name: "ng format",
			args: args{
				email: "example.com",
			},
			wantErr: false,
		},
		{
			name: "ng nil",
			args: args{
				email: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidEmail(tt.args.email); got != tt.wantErr {
				t.Errorf("ValidEmail() = %v, want %v", got, tt)
			}
		})
	}
}



