package entity

import (
	"os"
	"io"
	"bufio"
	"path/filepath"
	"errors"
	"log"
	"encoding/json"
	"agenda-go-cli/loghelper"
)

// UserFilter : UserFilter types take an *User and return a bool value.
type UserFilter func (*User) bool
// MeetingFilter : MeetingFilter types take an *User and return a bool value.
type MeetingFilter func (*Meeting) bool

var userinfoPath = "/src/agenda-go-cli/data/userinfo"
var metinfoPath = "/src/agenda-go-cli/data/meetinginfo"
var curUserPath = "/src/agenda-go-cli/data/curUser.txt"

var curUserName *string;

var dirty bool

var uData []User
var mData []Meeting

var errLog *log.Logger

func init()  {
	errLog = loghelper.Error
	dirty = false
	userinfoPath = filepath.Join(loghelper.GoPath, userinfoPath)
	metinfoPath = filepath.Join(loghelper.GoPath, metinfoPath)
	curUserPath = filepath.Join(loghelper.GoPath, curUserPath)
	if err := readFromFile(); err != nil {
		errLog.Println("readFromFile fail:", err)
	}
}

// Logout : logout
func Logout() error {
	curUserName = nil
	return Sync()
}

// Sync : sync file
func Sync() error {
	if err := writeToFile(); err != nil {
		errLog.Println("writeToFile fail:", err)
		return err
	}
	return nil
}


// CreateUser : create a user
// @param a user object
func CreateUser(v *User) {
	uData = append(uData, *v)
	dirty = true
}

// QueryUser : query users
// @param a lambda function as the filter
// @return a list of fitted users
func QueryUser(filter UserFilter) []User {
	var user []User
	for _, v := range uData {
		if filter(&v) {
			user = append(user, v)
		}
	}
	return user
}

// UpdateUser : update users
// @param a lambda function as the filter
// @param a lambda function as the method to update the user
// @return the number of updated users
func UpdateUser(filter UserFilter, switcher func (*User)) int {
	count := 0
	for i := 0; i < len(uData); i++ {
		if v := &uData[i]; filter(v) {
			switcher(v)
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

// DeleteUser : delete users
// @param a lambda function as the filter
// @return the number of deleted users
func DeleteUser(filter UserFilter) int {
	count := 0
	for i, v := range uData {
		if filter(&v) {
			uData[i] = uData[len(uData) - 1]
			uData = uData[:len(uData) - 1]
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

// CreateMeeting : create a meeting
// @param a meeting object
func CreateMeeting(v *Meeting) {
	mData = append(mData, *v)
	dirty = true
}

// QueryMeeting : query meetings
// @param a lambda function as the filter
// @return a list of fitted meetings
func QueryMeeting(filter MeetingFilter) []Meeting {
	var met []Meeting
	for _, v := range mData {
		if filter(&v) {
			met = append(met, v)
		}
	}
	return met;
}

// UpdateMeeting : update meetings
// @param a lambda function as the filter
// @param a lambda function as the method to update the meeting
// @return the number of updated meetings
func UpdateMeeting(filter MeetingFilter, switcher func (*Meeting)) int {
	count := 0
	for i := 0; i < len(mData); i++ {
		if v := &mData[i]; filter(v) {
			switcher(v)
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

// DeleteMeeting : delete meetings
// @param a lambda function as the filter
// @return the number of deleted meetings
func DeleteMeeting(filter MeetingFilter) int {
	count := 0
	for i, v := range mData {
		if filter(&v) {
			mData[i] = mData[len(mData) - 1]
			mData = mData[:len(mData) - 1]
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

// GetCurUser : get current user
// @return the current user
// @return error if current user does not exist
func GetCurUser() (User, error) {
	if curUserName == nil {
		return User{}, errors.New("Current user does not exist")
	}
	for _, v := range uData {
		if v.Name == *curUserName {
			return v, nil
		}
	}
	return User{}, errors.New("Current user does not exist")
}

// SetCurUser : get current user
// @param current user
func SetCurUser(u *User) {
	if u == nil {
		curUserName = nil
		return
	}
	if (curUserName == nil) {
		p := u.Name
		curUserName = &p
	} else {
		*curUserName = u.Name
	}
}

// readFromFile : read file content into memory
// @return if fail, error will be returned
func readFromFile() error {
	var e []error
	str, err1 := readString(curUserPath)
	if err1 != nil {
		e = append(e, err1)
	}
	curUserName = str
	if err := readUser(); err != nil {
		e = append(e, err)
	}
	if err := readMet(); err != nil {
		e = append(e, err)
	}
	if len(e) == 0 {
		return nil
	}
	er := e[0]
	for i := 1; i < len(e); i++ {
		er = errors.New(er.Error() + e[i].Error())
	}
	return er
}

// writeToFile : write file content from memory
// @return if fail, error will be returned
func writeToFile() error {
	var e []error
	if err := writeString(curUserPath, curUserName); err != nil {
		e = append(e, err)
	}
	if dirty {
		if err := writeJSON(userinfoPath, uData); err != nil {
			e = append(e, err)
		}
		if err := writeJSON(metinfoPath, mData); err != nil {
			e = append(e, err)
		}
	}
	if len(e) == 0 {
		return nil
	}
	er := e[0]
	for i := 1; i < len(e); i++ {
		er = errors.New(er.Error() + e[i].Error())
	}
	return er
}

func readUser() error {
	file, err := os.Open(userinfoPath);
	if err != nil {
		errLog.Println("Open File Fail:", userinfoPath, err)
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	switch err := dec.Decode(&uData); err {
	case nil, io.EOF:
		return nil
	default:
		errLog.Println("Decode User Fail:", err)
		return err
	}
}

func readMet() error {
	file, err := os.Open(metinfoPath);
	if err != nil {
		errLog.Println("Open File Fail:", metinfoPath, err)
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	switch err := dec.Decode(&mData); err {
	case nil, io.EOF:
		return nil
	default:
		errLog.Println("Decode Met Fail:", err)
		return err
	}
}

func writeJSON(fpath string, data interface{}) error {
	file, err := os.Create(fpath);
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)

	if err := enc.Encode(&data); err != nil {
		errLog.Println("writeJSON:", err)
		return err
	}
	return nil
}

func writeString(path string, data *string) error {
	file, err := os.Create(path)
	if err != nil {
		loghelper.Error.Println("Create file error:", path)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if data != nil {
		if _, err := writer.WriteString(*data); err != nil {
			loghelper.Error.Println("Write file fail:", path)
			return err
		}
	}
	if err := writer.Flush(); err != nil {
		loghelper.Error.Println("Flush file fail:", path)
		return err
	}
	return nil
}

func readString(path string) (*string, error) {
	file, err := os.Open(path)
	if err != nil {
		loghelper.Error.Println("Open file error:", path)
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n');
	if err != nil && err != io.EOF {
		loghelper.Error.Println("Read file fail:", path)
		return nil, err
	}
	return &str, nil
}
