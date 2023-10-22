package utils

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var purple = "\033[35m"

var loggerInstance *logger = InitLogger()

type logger struct {
	DataFile    *os.File
	InfoFile    *os.File
	WarningFile *os.File
	ErrorFile   *os.File
	FatalFile   *os.File
}

func formatCurrentTimeStamp() string {
	return time.Now().Format(time.RFC3339)
}

func colorMessage(msg string, color string) string {
	return fmt.Sprintf("%v %v %v", color, msg, reset)
}

func creteLogFile(dirname, name string) *os.File {
	formattedName := fmt.Sprintf("%v/%v-%v-log.txt", dirname, name, time.Now().Local().Format("2006-01-02"))
	file, err := os.OpenFile(formattedName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error when tying to create file %v", err)
		os.Exit(1)
	}
	return file
}

func checkAndCreateLogDir() error {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		return os.Mkdir("logs", os.ModePerm)
	}
	return nil
}

func InitLogger() *logger {
	if err := checkAndCreateLogDir(); err != nil {
		fmt.Println("Can't create log dir.")
		os.Exit(1)
	}

	lg := &logger{
		DataFile:    creteLogFile("logs", "data"),
		InfoFile:    creteLogFile("logs", "info"),
		WarningFile: creteLogFile("logs", "warning"),
		ErrorFile:   creteLogFile("logs", "error"),
		FatalFile:   creteLogFile("logs", "fatal"),
	}
	return lg
}

func CleanupLogFiles() {
	os.Remove(loggerInstance.DataFile.Name())
	os.Remove(loggerInstance.InfoFile.Name())
	os.Remove(loggerInstance.WarningFile.Name())
	os.Remove(loggerInstance.ErrorFile.Name())
	os.Remove(loggerInstance.FatalFile.Name())
	os.RemoveAll("logs")
}

func fmtArgs(args ...interface{}) string {
	formatted := ""
	for _, arg := range args {
		formatted += fmt.Sprintf("%v ", arg)
	}
	return formatted
}

func LogInfo(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("[%v][%v:%v][%v] %v\n", "INFO", runtime.FuncForPC(pc).Name(), line, formatCurrentTimeStamp(), fmtArgs(args...))
	fmt.Println(colorMessage(msg, green))
	loggerInstance.InfoFile.WriteString(msg)
}

func LogInfoF(str string, args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("[%v][%v:%v][%v] %v\n", "INFO", runtime.FuncForPC(pc).Name(), line, formatCurrentTimeStamp(), fmt.Sprintf(str, args...))
	fmt.Println(colorMessage(msg, green))
	loggerInstance.InfoFile.WriteString(msg)
}

func LogError(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("[%v][%v:%v][%v] %v\n", "ERROR", runtime.FuncForPC(pc).Name(), line, formatCurrentTimeStamp(), fmtArgs(args...))
	fmt.Println(colorMessage(msg, red))
	loggerInstance.ErrorFile.WriteString(msg)
}

func LogErrorF(str string, args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("[%v][%v:%v][%v] %v\n", "ERROR", runtime.FuncForPC(pc).Name(), line, formatCurrentTimeStamp(), fmt.Sprintf(str, args...))
	fmt.Println(colorMessage(msg, red))
	loggerInstance.ErrorFile.WriteString(msg)
}

func LogWarn(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("[%v][%v:%v][%v] %v\n", "WARNING", runtime.FuncForPC(pc).Name(), line, formatCurrentTimeStamp(), fmtArgs(args...))
	fmt.Println(colorMessage(msg, yellow))
	loggerInstance.WarningFile.WriteString(msg)
}

func LogWarnF(str string, args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("[%v][%v:%v][%v] %v\n", "WARNING", runtime.FuncForPC(pc).Name(), line, formatCurrentTimeStamp(), fmt.Sprintf(str, args...))
	fmt.Println(colorMessage(msg, yellow))
	loggerInstance.WarningFile.WriteString(msg)
}

func LogData(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("[%v][%v:%v][%v] %v\n", "DATA", runtime.FuncForPC(pc).Name(), line, formatCurrentTimeStamp(), fmtArgs(args...))
	fmt.Println(colorMessage(msg, blue))
	loggerInstance.DataFile.WriteString(msg)

}

func LogDataF(str string, args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("[%v][%v:%v][%v] %v\n", "DATA", runtime.FuncForPC(pc).Name(), line, formatCurrentTimeStamp(), fmt.Sprintf(str, args...))
	fmt.Println(colorMessage(msg, blue))
	loggerInstance.DataFile.WriteString(msg)
}

func LogFatal(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("[%v][%v:%v][%v] %v\n", "FATAL", runtime.FuncForPC(pc).Name(), line, formatCurrentTimeStamp(), fmtArgs(args...))
	fmt.Println(colorMessage(msg, purple))
	loggerInstance.FatalFile.WriteString(msg)
	os.Exit(1)
}

func LogFatalF(str string, args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("[%v][%v:%v][%v] %v\n", "FATAL", runtime.FuncForPC(pc).Name(), line, formatCurrentTimeStamp(), fmt.Sprintf(str, args...))
	fmt.Println(colorMessage(msg, purple))
	loggerInstance.FatalFile.WriteString(msg)
	os.Exit(1)
}
