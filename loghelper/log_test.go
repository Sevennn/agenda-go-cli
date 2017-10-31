package loghelper

import (
	"reflect"
	"testing"
)

func TestGetGOPATH(t *testing.T) {
	sP := GetGOPATH()
	tests := []struct {
		name string
		want *string
	}{
		{"GOPATH", sP},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGOPATH(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGOPATH() = %v, want %v", got, tt.want)
			}
		})
	}
}
