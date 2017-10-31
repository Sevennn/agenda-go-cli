package entity

import (
	"testing"
)

func TestMeeting_IsParticipator(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name      string
		meeting Meeting
		args      args
		want      bool
	}{
		{"Query Sponsor", Meeting{"u1", []string{"u2", "u3"}, Date{}, Date{}, ""}, args{"u1"}, false},
		{"Query Participator", Meeting{"u1", []string{"u2", "u3"}, Date{}, Date{}, ""}, args{"u2"}, true},
		{"Query not Participator", Meeting{"u1", []string{"u2", "u3"}, Date{}, Date{}, ""}, args{"u4"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.meeting.IsParticipator(tt.args.username); got != tt.want {
				t.Errorf("Meeting.IsParticipator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMeeting_AddParticipator(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name      string
		meeting *Meeting
		args      args
		want      bool
	}{
		{"Add Sponsor", &Meeting{"汉字", []string{"u2", "u3"}, Date{}, Date{}, ""}, args{"汉字"}, false},
		{"Add Participator", &Meeting{"u1", []string{"u2", "u3"}, Date{}, Date{}, ""}, args{"u2"}, false},
		{"Add not Participator", &Meeting{"汉字", []string{"u2", "u3"}, Date{}, Date{}, ""}, args{"汉"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.meeting.AddParticipator(tt.args.username); got != tt.want {
				t.Errorf("Meeting.AddParticipator() = %v, want %v", got, tt.want)
			}
		})
	}
}
