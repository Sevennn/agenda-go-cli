package entity

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"reflect"
	"testing"
)

var users = []User{
	{"u1", "up", "u1@q", "123"},
	{"u2", "ua", "u2@q", "456"},
	{"u3", "ua", "u3@q", "789"},
}

func TestCreateUser(t *testing.T) {
	type args struct {
		u *User
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{"CU u1:", args{&users[0]}, users[:1]},
		{"CU u2:", args{&users[1]}, users[:2]},
		{"CU u3:", args{&users[2]}, users[:]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateUser(tt.args.u)
			if got, _ := GetData("", ""); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryUser(t *testing.T) {
	type args struct {
		filter UserFilter
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{"QU u1",
			args{func (u *User) bool {
				return u.Name == "u1"
			}},
			users[0:1],
		},
		{"QU u0",
			args{func (u *User) bool {
				return u.Name == "u0"
			}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryUser(tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		filter   UserFilter
		switcher func(*User)
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"UU u1",
			args{func (u *User) bool {
				return u.Name == "u1"
			},
			func (u *User) {
				u.Phone = "321"
			},},
			1,
		},
		{"UU Phone123",
			args{func (u *User) bool {
				return u.Phone == "123"
			},
			func (u *User) {
				u.Phone = "Worry"
			},},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateUser(tt.args.filter, tt.args.switcher); got != tt.want {
				uData, _ := GetData("", "")
				t.Errorf("UpdateUser() = %v, want %v\nUser:%v", got, tt.want, uData)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		filter UserFilter
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"DU u1",
			args{func (u *User) bool {
				return u.Name == "u1"
			}},
			1,
		},
		{"DU u1",
			args{func (u *User) bool {
				return u.Name == "u1"
			}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteUser(tt.args.filter); got != tt.want {
				t.Errorf("DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateMeeting(t *testing.T) {
	type args struct {
		m *Meeting
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateMeeting(tt.args.m)
		})
	}
}

func TestQueryMeeting(t *testing.T) {
	type args struct {
		filter MeetingFilter
	}
	tests := []struct {
		name string
		args args
		want []Meeting
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryMeeting(tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryMeeting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateMeeting(t *testing.T) {
	type args struct {
		filter   MeetingFilter
		switcher func(*Meeting)
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateMeeting(tt.args.filter, tt.args.switcher); got != tt.want {
				t.Errorf("UpdateMeeting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteMeeting(t *testing.T) {
	type args struct {
		filter MeetingFilter
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteMeeting(tt.args.filter); got != tt.want {
				t.Errorf("DeleteMeeting() = %v, want %v", got, tt.want)
			}
		})
	}
}


const uFile = "u.test"
const mFile = "m.test"

func TestGetData(t *testing.T) {
	//str := `[{Name, Password, Email, Phone]`
	cases := []struct {
		uData, mData string
		uWant, mWant string
	}{
		{"", "null", "null", "null"},
		{`[{"Name":"n", "Password":"p", "Email":"q@q.com", "Phone":"123"},
			{"Name":"a", "Password":"a", "Email":"a@a.com", "Phone":"456"},
			{"Name":"m", "Password":"m"}
			]`,
			`[{"Sponsor":"n", "Participators":["a","b"], "StartDate":{"Year":2017, "Month":10, "Day":21, "Hour":7,"Minute":36}, "EndDate":{"Year":2017, "Month":10, "Day":22, "Hour":8,"Minute":0}, "Tittle":"t"},
			{"Sponsor":"a", "Participators":[], "StartDate":{"Year":2017, "Month":10, "Day":21, "Hour":7,"Minute":36}, "EndDate":{"Year":2017, "Month":10, "Day":22, "Hour":8,"Minute":0}, "Tittle":"f"},
			{"Sponsor":"n", "Participators":["n","b"], "StartDate":{"Year":2017, "Month":10, "Day":21, "Hour":7,"Minute":36}, "EndDate":{"Year":2017, "Month":10, "Day":22, "Hour":8,"Minute":0}, "Tittle":"g"}
			]`,
			`[{"Name":"n","Password":"p","Email":"q@q.com","Phone":"123"},{"Name":"a","Password":"a","Email":"a@a.com","Phone":"456"},{"Name":"m","Password":"m","Email":"","Phone":""}]`,
			`[{"Sponsor":"n","Participators":["a","b"],"StartDate":{"Year":2017,"Month":10,"Day":21,"Hour":7,"Minute":36},"EndDate":{"Year":2017,"Month":10,"Day":22,"Hour":8,"Minute":0},"Tittle":"t"},{"Sponsor":"a","Participators":[],"StartDate":{"Year":2017,"Month":10,"Day":21,"Hour":7,"Minute":36},"EndDate":{"Year":2017,"Month":10,"Day":22,"Hour":8,"Minute":0},"Tittle":"f"},{"Sponsor":"n","Participators":["n","b"],"StartDate":{"Year":2017,"Month":10,"Day":21,"Hour":7,"Minute":36},"EndDate":{"Year":2017,"Month":10,"Day":22,"Hour":8,"Minute":0},"Tittle":"g"}]`,
		},
	}

	for i, c := range cases {
		createTestFile(uFile, &c.uData)
		createTestFile(mFile, &c.mData)
		uData, mData := GetData(uFile, mFile)
		u := encodeJSON(uData)
		m := encodeJSON(mData)
		if u != c.uWant {
			t.Errorf("Test:%v: Sync User() == %v, want %v", i, u, c.uWant)
		}
		if m != c.mWant {
			t.Errorf("Test:%v: Sync Met() == %v, want %v", i, m, c.mWant)
		}
	}

	os.Remove(uFile)
	os.Remove(mFile)
}

func createTestFile(path string, data *string) {
	file, err := os.Create(path)
	if err != nil {
		log.Printf("Create file %q error.\n", path)
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(*data); err != nil {
		log.Printf("Write file %q file.\n", path)
		panic(err)
	}
	if err := writer.Flush(); err != nil {
		panic(err)
	}
}

func encodeJSON(intf interface{}) string {
	data, err := json.Marshal(intf)
	if err != nil {
		log.Printf("Can not encode Interface: %v\n", err)
		return ""
	}
	return string(data[:])
}
