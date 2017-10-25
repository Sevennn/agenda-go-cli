package service

import (
	"entity/storage"
	"entity/date"
	"entity/meeting"
	"entity/user"
)

func MeetingQuery(userName string, title string) []meeting.Meeting {
	var rlist := storage.QueryMeeting(func(t meeting.Meeting){
		
	})

}