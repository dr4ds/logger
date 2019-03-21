package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

type LogType string
type LogLevel int

const (
	TypeDebug    LogType = "DEBUG"
	TypeSuccess  LogType = "SUCCESS"
	TypeInfo     LogType = "INFO"
	TypeWarning  LogType = "WARNING"
	TypeError    LogType = "ERROR"
	TypeCritical LogType = "CRITICAL"
)

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

type Logger struct {
	writer io.Writer
	mux    sync.Mutex
	buf    []byte
	level  LogLevel
}

func getTimeNow() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

func New(lvl LogLevel) *Logger {
	l := new(Logger)
	l.level = lvl

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0666)
		if err != nil {
			panic(err)
		}
	}

	f, err := os.Create("logs/" + strings.Replace(getTimeNow(), ":", "_", -1) + ".log")

	if err != nil {
		panic(err)
	}

	l.writer = io.MultiWriter(os.Stdout, f)

	return l
}

func (l *Logger) Log(s string, logType LogType) {
	l.mux.Lock()
	defer l.mux.Unlock()

	switch logType {
	case TypeDebug:
		if l.level > LevelDebug {
			return
		}
		color.Set(color.FgWhite)
	case TypeInfo:
		if l.level > LevelInfo {
			return
		}
		color.Set(color.FgCyan)
	case TypeSuccess:
		if l.level > LevelInfo {
			return
		}
		color.Set(color.FgGreen)
	case TypeWarning:
		if l.level > LevelWarning {
			return
		}
		color.Set(color.FgYellow)
	case TypeError:
		if l.level > LevelError {
			return
		}
		color.Set(color.FgRed)
	case TypeCritical:
		color.Set(color.FgBlack, color.BgRed)
	}
	defer color.Unset()

	l.buf = l.buf[:0]

	s = fmt.Sprintf("[%s] [%s]: %s", getTimeNow(), logType, s)
	l.buf = append(l.buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}

	l.writer.Write(l.buf)
}

func (l *Logger) SetLevel(lvl LogLevel) {
	l.level = lvl
}

func (l *Logger) Error(v ...interface{}) {
	l.Log(fmt.Sprint(v...), TypeError)
}

func (l *Logger) Errorln(v ...interface{}) {
	l.Log(fmt.Sprintln(v...), TypeError)
}

func (l *Logger) Errorf(s string, v ...interface{}) {
	l.Log(fmt.Sprintf(s, v...), TypeError)
}

func (l *Logger) Warning(v ...interface{}) {
	l.Log(fmt.Sprint(v...), TypeWarning)
}

func (l *Logger) Warningln(v ...interface{}) {
	l.Log(fmt.Sprintln(v...), TypeWarning)
}

func (l *Logger) Warningf(s string, v ...interface{}) {
	l.Log(fmt.Sprintf(s, v...), TypeWarning)
}

func (l *Logger) Info(v ...interface{}) {
	l.Log(fmt.Sprint(v...), TypeInfo)
}

func (l *Logger) Infoln(v ...interface{}) {
	l.Log(fmt.Sprintln(v...), TypeInfo)
}

func (l *Logger) Infof(s string, v ...interface{}) {
	l.Log(fmt.Sprintf(s, v...), TypeInfo)
}

func (l *Logger) Success(v ...interface{}) {
	l.Log(fmt.Sprint(v...), TypeSuccess)
}

func (l *Logger) Successln(v ...interface{}) {
	l.Log(fmt.Sprintln(v...), TypeSuccess)
}

func (l *Logger) Successf(s string, v ...interface{}) {
	l.Log(fmt.Sprintf(s, v...), TypeSuccess)
}

func (l *Logger) Debug(v ...interface{}) {
	l.Log(fmt.Sprint(v...), TypeDebug)
}

func (l *Logger) Debugln(v ...interface{}) {
	l.Log(fmt.Sprintln(v...), TypeDebug)
}

func (l *Logger) Debugf(s string, v ...interface{}) {
	l.Log(fmt.Sprintf(s, v...), TypeDebug)
}

func (l *Logger) Critical(v ...interface{}) {
	l.Log(fmt.Sprint(v...), TypeCritical)
	os.Exit(0)
}

func (l *Logger) Criticalln(v ...interface{}) {
	l.Log(fmt.Sprintln(v...), TypeCritical)
	os.Exit(0)
}

func (l *Logger) Criticalf(s string, v ...interface{}) {
	l.Log(fmt.Sprintf(s, v...), TypeCritical)
	os.Exit(0)
}
