package entity

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	type args struct {
		date Date
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"leap year", args{Date{2000, 2, 29, 1, 1}}, true},
		{"leap year", args{Date{2004, 2, 29, 1, 1}}, true},
		{"not leap year", args{Date{1500, 2, 29, 1, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.date); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_MoreThan(t *testing.T) {
	type args struct {
		date Date
	}
	tests := []struct {
		name   string
		date Date
		args   args
		want   bool
	}{
		{"month: 2 > 3", Date{2000, 2, 29, 1, 1}, args{Date{2000, 3, 1, 20, 1}}, false},
		{"minute: 1 > 1", Date{2000, 2, 29, 1, 1}, args{Date{2000, 2, 29, 1, 1}}, false},
		{"vaild: 1 > 0", Date{2000, 2, 28, 1, 1}, args{Date{2000, 2, 28, 1, 0}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.MoreThan(tt.args.date); got != tt.want {
				t.Errorf("Date.MoreThan() = %v, want %v", got, tt.want)
			}
		})
	}
}
