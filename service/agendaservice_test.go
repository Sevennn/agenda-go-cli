package service

import (
	"agenda-go-cli/entity"
	"agenda-go-cli/loghelper"
	"reflect"
	"testing"
)

var users = []entity.User{
	{Name: "TEST_AGENDASERVICE1", Password: "up", Email: "TEST_AGENDASERVICE1@q", Phone: "123"},
	{Name: "TEST_AGENDASERVICE2", Password: "ua", Email: "TEST_AGENDASERVICE2@q", Phone: "456"},
	{Name: "TEST_AGENDASERVICE3", Password: "ua", Email: "TEST_AGENDASERVICE3@q", Phone: "789"},
}

var dates = []entity.Date{
	{Year: 2000, Month: 2, Day: 28, Hour: 0, Minute: 0},
	{Year: 2000, Month: 2, Day: 29, Hour: 1, Minute: 1},
	{Year: 2001, Month: 2, Day: 28, Hour: 1, Minute: 1},
}

var meetings = []entity.Meeting{
	{Sponsor:users[0].Name, Participators:[]string{users[1].Name, users[2].Name}, StartDate:dates[0], EndDate:dates[1], Title:"TEST_AGENDASERVICE1 TEST_AGENDASERVICE2 TEST_AGENDASERVICE3"},
	{Sponsor:users[0].Name, Participators:[]string{}, StartDate:dates[1], EndDate:dates[2],Title:"only TEST_AGENDASERVICE1"},
	{Sponsor:users[1].Name, Participators:[]string{users[2].Name}, StartDate:dates[1], EndDate:dates[2], Title:"TEST_AGENDASERVICE2 TEST_AGENDASERVICE3"},
}

func _testLogin(u *entity.User) {
	UserLogout()
	if re, err := UserRegister(users[0].Name, users[0].Password, users[0].Email, users[0].Phone); err != nil {
		loghelper.Error.Println("UserRegister() Fail: ", re, err)
	}
	if re, err := UserRegister(users[1].Name, users[1].Password, users[1].Email, users[1].Phone); err != nil {
		loghelper.Error.Println("UserRegister() Fail: ", re, err)
	}
	if re, err := UserRegister(users[2].Name, users[2].Password, users[2].Email, users[2].Phone); err != nil {
		loghelper.Error.Println("UserRegister() Fail: ", re, err)
	}
	if login := UserLogin(u.Name, u.Password); !login {
		loghelper.Error.Println("UserLogin() Fail: ", login)
	}
	CreateMeeting(meetings[0].Sponsor, meetings[0].Title, "2000-02-28/00:00", "2000-02-29/01:01", meetings[0].Participators)
	CreateMeeting(meetings[1].Sponsor, meetings[1].Title, "2000-02-29/01:01", "2001-02-28/01:01", meetings[1].Participators)
	CreateMeeting(meetings[2].Sponsor, meetings[2].Title, "2000-02-29/01:01", "2001-02-28/01:01", meetings[2].Participators)
}

func _reset() {
	UserLogout()
	DeleteUser(users[0].Name)
	DeleteUser(users[1].Name)
	DeleteUser(users[2].Name)
}

func TestUserLogout(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Logout", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserLogout(); got != tt.want {
				t.Errorf("UserLogout() = %v, want %v", got, tt.want)
			}
		})
	}
	_reset()
}

func TestGetCurUser(t *testing.T) {
	tests := []struct {
		name  string
		want  entity.User
		want1 bool
	}{
		{"Get login", users[0], true},
		{"Get no login", entity.User{}, false},
	}
	_testLogin(&tests[0].want)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetCurUser()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurUser() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetCurUser() got1 = %v, want %v", got1, tt.want1)
			}
			UserLogout()
		})
	}
	_reset()
}

func TestUserLogin(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Login error", args{users[0].Name, users[1].Password}, false},
		{"Login succ", args{users[0].Name, users[0].Password}, true},
	}
	_testLogin(&users[0])
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserLogin(tt.args.username, tt.args.password); got != tt.want {
				t.Errorf("UserLogin() = %v, want %v", got, tt.want)
			}
		})
		UserLogout()
	}
	_reset()
}

func TestUserRegister(t *testing.T) {
	type args struct {
		username string
		password string
		email    string
		phone    string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Register new", args{users[1].Name, users[1].Password, users[1].Email, users[1].Phone}, true, false},
		{"Register exit", args{users[1].Name, users[1].Password, users[1].Email, users[1].Phone}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UserRegister(tt.args.username, tt.args.password, tt.args.email, tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserRegister() = %v, want %v", got, tt.want)
			}
		})
	}
	_reset()
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Delete exit", args{users[0].Name}, true},
		{"Delete not exit", args{""}, true},
	}
	_testLogin(&users[0])
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteUser(tt.args.username); got != tt.want {
				t.Errorf("DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
	_reset()
}

func TestListAllUser(t *testing.T) {
	exitUsers := ListAllUser()
	tests := []struct {
		name string
		want []entity.User
	}{
		{"List all", append(exitUsers, users[0], users[1], users[2])},
	}
	_testLogin(&users[0])
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListAllUser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllUser() = %v, want %v", got, tt.want)
			}
		})
	}
	_reset()
}

func TestCreateMeeting(t *testing.T) {
	type args struct {
		username     string
		title        string
		startDate    string
		endDate      string
		participator []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"CM new", args{users[0].Name, meetings[0].Title, "0000-00-00/00:00", "0001-00-00/00:00", nil}, true},
		{"CM exit", args{users[0].Name, meetings[0].Title, "0000-00-00/00:00", "0001-00-00/00:00", nil}, false},
	}
	_testLogin(&users[0])
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMeeting(tt.args.username, tt.args.title, tt.args.startDate, tt.args.endDate, tt.args.participator); got != tt.want {
				t.Errorf("CreateMeeting() = %v, want %v", got, tt.want)
			}
		})
	}
	_reset()
}

func TestQueryMeeting(t *testing.T) {
	type args struct {
		username  string
		startDate string
		endDate   string
	}
	tests := []struct {
		name  string
		args  args
		want  []entity.Meeting
		want1 bool
	}{
		{"QM not exit", args{users[0].Name, "1000-01-01/00:00", "2100-01-01/00:00"}, meetings[0:2], true},
		{"QM exit", args{users[0].Name, "2100-01-01/00:00", "2100-01-01/00:00"}, nil, true},
	}
	_testLogin(&users[0])
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := QueryMeeting(tt.args.username, tt.args.startDate, tt.args.endDate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryMeeting() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("QueryMeeting() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
	_reset()
}

func TestDeleteMeeting(t *testing.T) {
	type args struct {
		username string
		title    string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"DM", args{users[0].Name, meetings[0].Title}, 1},
		{"DM", args{users[0].Name, meetings[0].Title}, 0},
	}
	_testLogin(&users[0])
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteMeeting(tt.args.username, tt.args.title); got != tt.want {
				t.Errorf("DeleteMeeting() = %v, want %v", got, tt.want)
			}
		})
	}
	_reset()
}

func TestQuitMeeting(t *testing.T) {
	type args struct {
		username string
		title    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Quit exit", args{users[1].Name, meetings[0].Title}, true},
		{"Quit not exit", args{users[1].Name, meetings[0].Title}, false},
	}
	_testLogin(&users[0])
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuitMeeting(tt.args.username, tt.args.title); got != tt.want {
				t.Errorf("QuitMeeting() = %v, want %v", got, tt.want)
			}
		})
	}
	_reset()
}

func TestClearMeeting(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{"C 1", args{users[0].Name}, 2, true},
		{"C clear", args{users[0].Name}, 0, true},
		{"C 2", args{users[1].Name}, 1, true},
		{"C 3", args{users[2].Name}, 0, true},
	}
	_testLogin(&users[0])
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ClearMeeting(tt.args.username)
			if got != tt.want {
				t.Errorf("ClearMeeting() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ClearMeeting() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
	_reset()
}

func TestAddMeetingParticipator(t *testing.T) {
	type args struct {
		username      string
		title         string
		participators []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"AMP not exit", args{users[0].Name, meetings[1].Title, []string{users[1].Name, users[2].Name}}, true},
		{"AMP exit", args{users[0].Name, meetings[1].Title, []string{users[1].Name, users[2].Name}}, false},
	}
	_testLogin(&users[0])
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddMeetingParticipator(tt.args.username, tt.args.title, tt.args.participators); got != tt.want {
				t.Errorf("AddMeetingParticipator() = %v, want %v", got, tt.want)
			}
		})
	}
	_reset()
}

func TestRemoveMeetingParticipator(t *testing.T) {
	type args struct {
		username      string
		title         string
		participators []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"RMP exit", args{users[0].Name, meetings[0].Title, []string{users[1].Name, users[2].Name}}, true},
		{"RMP not exit", args{users[0].Name, meetings[0].Title, []string{users[1].Name, users[2].Name}}, false},
	}
	_testLogin(&users[0])
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveMeetingParticipator(tt.args.username, tt.args.title, tt.args.participators); got != tt.want {
				t.Errorf("RemoveMeetingParticipator() = %v, want %v", got, tt.want)
			}
		})
	}
	_reset()
}
