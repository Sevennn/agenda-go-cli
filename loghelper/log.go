package loghelper

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"runtime"
)

var (
	// Info : Discard
	Info    *log.Logger
	// Warning : Stdout
	Warning *log.Logger
	// Error : Stderr
	Error   *log.Logger

	// GoPath : GoPath
	GoPath string

)

var errlog *os.File

func set(
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func init() {
	errlog = getErrLogFile()
	set(ioutil.Discard, os.Stdout, errlog)

	// Info.Println("Special Information")
	// Warning.Println("There is something you need to know about")
	Error.Println("Start up")
}

// Free : close log file
func Free()  {
	errlog.Close()
}

func getErrLogFile() *os.File  {

	if sP := GetGOPATH(); sP != nil {
		GoPath = *sP
	} else {
		log.Fatalf("data file not ecist\n")
		os.Exit(1)
	}
	logPath := filepath.Join(GoPath, "/src/agenda-go-cli/data/error.log")

	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v\n", err)
	}
	return file;
	// defer file.Close()
}

// GetGOPATH : get GOPATH
func GetGOPATH() *string {
	var sp string
	if runtime.GOOS == "windows" {
		sp = ";"
	} else {
		sp = ":"
	}
	goPath := strings.Split(os.Getenv("GOPATH"), sp)
	for _, v := range goPath {
		if _, err := os.Stat(filepath.Join(v, "/src/agenda-go-cli/data/meetinginfo")); !os.IsNotExist(err) {
			return &v
		}
	}
	return nil
}
