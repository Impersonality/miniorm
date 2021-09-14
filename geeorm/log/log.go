package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	errLog  = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog = log.New(os.Stdout, "\033[34m[info]\033[0m ", log.LstdFlags|log.Lshortfile)
	mu      sync.Mutex
	loggers = []*log.Logger{errLog, infoLog}
)

var (
	Error  = errLog.Println
	Errorf = errLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, log := range loggers {
		log.SetOutput(os.Stdout)
	}

	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
	if ErrorLevel < level {
		errLog.SetOutput(ioutil.Discard)
	}
}
