package entity

import (
	"os"
	"io"
	"path/filepath"
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

var dirty bool

var uData []User
var mData []Meeting

var errLog *log.Logger

func init()  {
	errLog = loghelper.Error
	dirty = false
	birDir := os.Getenv("GOPATH")
	userinfoPath = filepath.Join(birDir, userinfoPath)
	metinfoPath = filepath.Join(birDir, metinfoPath)
	if err := readFromFile(); err != nil {
		errLog.Println("readFromFile fail:", err)
	}
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

// GetData : for test
// @param : test file path
func GetData(uPath, mPath string) ([]User, []Meeting) {
	if len(uPath) == 0 || len(mPath) == 0 {
		return uData, mData
	}
	userinfoPath = uPath
	metinfoPath = mPath
	uData = nil
	mData = nil
	if err := readFromFile(); err != nil {
		errLog.Println("readFromFile fail:", err)
	}
	if err := writeToFile(); err != nil {
		errLog.Println("writeToFile fail:", err)
	}
	// errLog.Println(uData)
	// errLog.Println(mData)
	return uData, mData
}

// readFromFile : read file content into memory
// @return if fail, error will be returned
func readFromFile() error {
	if err := readUser(); err != nil {
		return err
	}
	return readMet()
}

// writeToFile : write file content from memory
// @return if fail, error will be returned
func writeToFile() error {
	// dirty = true
	if !dirty {
		return nil
	}
	if err := writeJSON(userinfoPath, uData); err != nil {
		return err;
	}
	return writeJSON(metinfoPath, mData)
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
