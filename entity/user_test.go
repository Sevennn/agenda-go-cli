package entity

import (
	"testing"
)

func TestUser_GetPassword(t *testing.T) {
	tests := []struct {
		name   string
		user User
		want   string
	}{
		{"Password: Long", User{"", "aaaaaaaaaaaaaaaaaaaaaaaaa", "", ""}, "aaaaaaaaaaaaaaaaaaaaaaaaa"},
		{"Password: 汉字", User{"", "汉字", "", ""}, "汉字"},
		{"Password: nil", User{"", "", "", ""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.GetPassword(); got != tt.want {
				t.Errorf("User.GetPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
