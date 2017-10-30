package service

import (
	"agenda-go-cli/entity"
	"agenda-go-cli/loghelper"
	"log"
)


var curuserinfoPath = "/src/agenda-go-cli/data/curuser.txt"
var errLog *log.Logger
type User entity.User
type Meeting entity.Meeting
func init() {
	errLog = loghelper.Error
}
func UserLogout() bool {
	if err := entity.Logout(); err != nil {
		return false
	} else {
		return true
	}
}
func GetCurUser() (entity.User,bool) {
	if cu,err := entity.GetCurUser(); err != nil {
		return cu, false
	} else {
		return cu, true
	}
}
func UserLogin(username string, password string) bool {
	user := entity.QueryUser(func (u *entity.User) bool {
		if u.Name == username && u.Password == password {
			return true
		}
		return false
	})
	if len(user) == 0 {
		errLog.Println("Login: User not Exist")
		return false
	}
	entity.SetCurUser(&user[0])
	if err := entity.Sync(); err != nil {
		errLog.Println("Login: error occurred when set curuser")
		return false
	}
	return true
}

func UserRegister(username string, password string, email string, phone string) (bool, error) {
	user := entity.QueryUser(func (u *entity.User) bool {
		return u.Name == username
	})
	if len(user) == 1 {
		errLog.Println("User Register: Already exist username")
		return false, nil
	}
	entity.CreateUser(&entity.User{username, password, email, phone})
	if err := entity.Sync(); err != nil {
		return true, err
	}
	return true, nil
}


func DeleteUser(username string) bool {
	entity.DeleteUser(func (u *entity.User) bool {
		return u.Name == username
	})
	entity.UpdateMeeting(
		func(m *entity.Meeting) bool {
			return m.IsParticipator(username)
		},
		func(m *entity.Meeting) {
			m.DeleteParticipator(username)
		})
	entity.DeleteMeeting(func(m *entity.Meeting) bool {
		return m.Sponsor == username || len(m.GetParticipator()) == 0
	})
	if err := entity.Sync(); err != nil {
		return false
	}
	return UserLogout()
}

func ListAllUser() []entity.User {
	return entity.QueryUser(func (u *entity.User) bool {
		return true
	})
}

func CreateMeeting(username string, title string, startDate string, endDate string, participator []string) bool {
	for _, i := range participator {
		if username == i {
			errLog.Println("Create Meeting: sponsor can't be participator")
			return false
		}
		l := entity.QueryUser(func (u *entity.User) bool{
			return u.Name == i
		})
		if (len(l) == 0) {
			errLog.Println("Create Meeting: no such a user : ", i)
			return false
		}
		dc := 0
		for _, j := range participator {
			if j == i {
				dc++
				if dc == 2 {
					errLog.Println("Create Meeting: duplicate participator")
					return false
				}
			}
		}
	}
	sTime,err := entity.StringToDate(startDate)
	if err != nil {
		errLog.Println("Create Meeting: Wrong Date")
		return false
	}
	eTime,err := entity.StringToDate(endDate)
	if err != nil {
		errLog.Println("Create Meeting: Wrong Date")
		return false
	}
	if eTime.LessThan(sTime) == true {
		errLog.Println("Create Meeting: Start Time greater than end time")
		return false
	}
	for _, p := range participator {
		l := entity.QueryMeeting(func (m *entity.Meeting) bool {
			if m.Sponsor == p || m.IsParticipator(p) {
				if m.StartDate.LessOrEqual(sTime) && m.EndDate.MoreThan(sTime) {
					return true
				}
				if m.StartDate.LessThan(eTime) && m.EndDate.GreateOrEqual(eTime) {
					return true
				}
				if m.StartDate.GreateOrEqual(sTime) && m.EndDate.LessOrEqual(eTime) {
					return true
				}
			}
			return false
		})
		if len(l) > 0 {
			errLog.Println("Create Meeting: ",p," time conflict")
			return false
		}
	}
	tu := entity.QueryUser(func (u *entity.User) bool {
		return u.Name == username
	})
	if len(tu) == 0 {
		errLog.Println("Create Meeting: Sponsor ", username, " not exist")
		return false
	}
	l := entity.QueryMeeting(func (m *entity.Meeting) bool {
		if m.Sponsor == username || m.IsParticipator(username) {
			if m.StartDate.LessOrEqual(sTime) && m.EndDate.MoreThan(sTime) {
				return true
			}
			if m.StartDate.LessThan(eTime) && m.EndDate.GreateOrEqual(eTime) {
				return true
			}
			if m.StartDate.GreateOrEqual(sTime) && m.EndDate.LessOrEqual(eTime) {
				return true
			}
		}
		return false
	})

	if len(l) > 0 {
		errLog.Println("Create Meeting: ", username, " time conflict")
		return false;
	}
	entity.CreateMeeting(&entity.Meeting{username, participator,sTime,eTime, title})
	if err := entity.Sync(); err != nil {
		return false
	}
	return true
}

func QueryMeeting(username, startDate, endDate string) ([]entity.Meeting,bool) {
	sTime,err := entity.StringToDate(startDate)
	var m []entity.Meeting
	if err != nil {
		errLog.Println("Query Meeting: Wrong StartDate")
		return m,false
	}
	eTime,err := entity.StringToDate(endDate)
	if err != nil {
		errLog.Println("Query Meeting: Wrong EndDate")
		return m,false
	}
	if eTime.LessThan(sTime) == true {
		errLog.Println("Query Meeting: Start Time greater than end time")
		return m, false
	}

	tm := entity.QueryMeeting(func (m *entity.Meeting) bool {
		if m.Sponsor == username || m.IsParticipator(username) {
			if m.StartDate.LessOrEqual(sTime) && m.EndDate.MoreThan(sTime) {
				return true
			}
			if m.StartDate.LessOrEqual(eTime) && m.EndDate.GreateOrEqual(eTime) {
				return true
			}
			if m.StartDate.GreateOrEqual(sTime) && m.EndDate.LessOrEqual(eTime) {
				return true
			}
		}
		return false
	})
	return tm,true
}

func DeleteMeeting(username, title string) int {
	return entity.DeleteMeeting(func (m *entity.Meeting) bool {
		return m.Sponsor == username && m.Title == title
	})
}

func QuitMeeting(username string, title string) bool {
	flag :=entity.QueryMeeting(func (m *entity.Meeting) bool {
		return m.Title == title && m.IsParticipator(username) == true
	})
	if len(flag) == 0 {
		return false
	}
	entity.UpdateMeeting(func (m *entity.Meeting) bool {
		return m.IsParticipator(username) == true && m.Title == title
	}, func (m *entity.Meeting) {
		m.DeleteParticipator(username)
	})
	entity.DeleteMeeting(func (m *entity.Meeting) bool {
		return len(m.GetParticipator()) == 0
	})
	return true
}

func ClearMeeting(username string) (int,bool) {
	cm := entity.DeleteMeeting(func (m *entity.Meeting) bool {
		return m.Sponsor == username
	})
	if err := entity.Sync(); err != nil {
		errLog.Println("Clear Meeting: Delete failed")
		return cm, false
	} else {
		return cm, true
	}
}

func AddMeetingParticipator(username string, title string, participators []string) bool {
	for _, p := range participators {
		uc := entity.QueryUser(func (u *entity.User) bool {
			return u.Name == p
		})
		if len(uc) == 0 {
			errLog.Println("Add Meeting Participator: No such a user: ", p)
			return false
		}
		qm := entity.QueryMeeting(func (m *entity.Meeting) bool {
			return m.Sponsor == username && m.IsParticipator(p)
		})
		if len(qm) != 0 {
			errLog.Println("Add Meeting Participator: ",p, "Already in meeting")
			return false
		}
	}
	mt := entity.UpdateMeeting(func (m *entity.Meeting) bool {
		return m.Sponsor == username && m.Title == title
	}, func (m *entity.Meeting) {
		for _,p := range participators {
			m.AddParticipator(p)
		}
	})
	if mt == 0 {
		errLog.Println("Add Meeting Participator: no such meeting")
		return false
	}
	if err := entity.Sync(); err != nil {
		return false
	}
	return true
}

func RemoveMeetingParticipator(username string, title string, participators []string) bool {
	for _, p := range participators {
		uc := entity.QueryUser(func (u *entity.User) bool {
			return u.Name == p
		})
		if len(uc) == 0 {
			errLog.Println("Remove Meeting Participator: No such a user: ", p)
			return false
		}
		qm := entity.QueryMeeting(func (m *entity.Meeting) bool {
			return m.Sponsor == username && m.IsParticipator(p)
		})
		if len(qm) == 0 {
			errLog.Println("Remove Meeting Participator: Not in Meeting :", p)
			return false
		}
	}
	mt := entity.UpdateMeeting(func (m *entity.Meeting) bool {
		return m.Sponsor == username && m.Title == title
	}, func (m *entity.Meeting) {
		for _, p := range participators {
			m.DeleteParticipator(p)
		}
	})
	if mt == 0 {
		errLog.Println("Remove Meeting Participator: no such a meeting: ", title)
		return false
	}
	entity.DeleteMeeting(func(m *entity.Meeting) bool {
		return m.Sponsor == username && len(m.GetParticipator()) == 0
	})
	if err := entity.Sync(); err != nil {
		return false
	}
	return true
}
