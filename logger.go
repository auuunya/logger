package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var logging *Logger

// Errorf 默认日志对象方法，记录一条错误日志，需要先初始化
func Errorf(format string, v ...interface{}) {
	logging.Errorf(format, v...)
}

// Error 默认日志对象方法，记录一条消息日志，需要先初始化
func Error(args ...interface{}) {
	logging.Error(args...)
}

// Infof 默认日志对象方法，记录一条消息日志，需要先初始化
func Infof(format string, v ...interface{}) {
	logging.Infof(format, v...)
}

// Info 默认日志对象方法，记录一条消息日志，需要先初始化
func Info(args ...interface{}) {
	logging.Info(args...)
}

// Debugf 默认日志对象方法，记录一条消息日志，需要先初始化
func Debugf(format string, v ...interface{}) {
	logging.Debugf(format, v...)
}

// Debug 默认日志对象方法，记录一条调试日志，需要先初始化
func Debug(args ...interface{}) {
	logging.Debug(args...)
}

//Waring
func Warningf(format string, v ...interface{}) {
	logging.Warningf(format, v...)
}

// Waringf
func Warning(args ...interface{}) {
	logging.Warning(args...)
}

type Level int16

// 定义日志级别
const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

//定义日志结构体
type Logger struct {
	level     Level
	file      *os.File
	loggerf   *log.Logger
	writeFile bool
	path      string
	logname   string
	today     string
	switchLog bool
	closed    bool
}

// 获取日志级别
func getLevel(level Level) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarningLevel:
		return "WARNING"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

//等级设置
func SetLevel(level string) Level {
	switch strings.ToLower(level) {
	case "DEBUG":
		return DebugLevel
	case "INFO":
		return InfoLevel
	case "warning":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	default:
		return DebugLevel
	}
}

/**
 * @Description:判断路径是否存在
 * @param path 路径，type: 字符串
 */

func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	return f, err == nil || os.IsExist(err)
}

/**
 * @Description:判断所给路径是否为文件
 * @param path 文件，type: 字符串
 */

func IsFile(path string) (os.FileInfo, bool) {
	f, flag := IsExists(path)
	return f, flag && !f.IsDir()
}

//时间转字符串
func formatTime(t *time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// 打开日志文件
func NewLogger(level, filepath, logname string, outfile, switchLog bool) {
	logging = &Logger{
		level:     SetLevel(level),
		path:      filepath,
		logname:   logname,
		writeFile: outfile,
		switchLog: switchLog,
		today:     time.Now().Format("2006-01-02"),
	}
	logging.Init()
}

// 关闭日志文件
func (logging *Logger) CloseLogger() {
	logging.file.Close()
}

func (llog *Logger) Init() {
	llog.loggerf = log.New(os.Stdout, "Logger_Out_", log.Llongfile|log.Ltime|log.Ldate)
	if llog.writeFile {
		if llog.path != "" {
			if err := llog.SetLogFile(); err == nil {
				llog.SetWrite()
				logging.closed = false
			}
		}
	}
	if llog.switchLog {
		go logging.logWorker()
	}
}

// 控制文件切换,协程函数
func (llog *Logger) logWorker() {
	for llog.closed == false {
		nowDate := time.Now().Format("2006-01-02")
		if nowDate != llog.today {
			llog.rotate()
		}
	}
}

// 文件切换函数
func (llog *Logger) rotate() {
	fmt.Println("rotate run")
	defer func() {
		rec := recover()
		if rec != nil {
			llog.loggerf.Printf("recover error: %v", rec)
		}
	}()
	logname := filepath.Join(llog.path, llog.logname)
	_, err := os.Stat(logname)
	if err == nil {
		llog.CloseLogger()
		err = os.Rename(logname, logname+"."+time.Now().Add(-1*time.Hour*24).Format("2006-01-02"))
		if err != nil {
			llog.loggerf.Printf("rename log error: %v", err)
		}
	}
	if llog.logname != "" {
		nextfile, err := os.OpenFile(filepath.Join(llog.path, llog.logname),
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			llog.loggerf.Printf("open file error: %v", err)
		}
		llog.file = nextfile
		llog.SetWrite()
	}
	llog.today = time.Now().Format("2006-01-02")
}

//创建日志文件
func (llog *Logger) SetLogPath(path string) {
	llog.path = path
}

// 创建日志文件
func (llog *Logger) SetLogFile() error {
	pathList := strings.Split(llog.path, "/")
	_, filebool := IsFile(filepath.Join(llog.path, llog.logname))
	filePathTemp := ""
	if filebool {
		loggerfile, err := os.OpenFile(llog.path, os.O_APPEND, 0666)
		if err != nil {
			return errors.New(fmt.Sprintln("Open File Error:", err))
		}
		llog.file = loggerfile
	} else {
		for _, v := range pathList {
			filePathTemp += v
			_, isExistsPath := IsExists(filePathTemp)
			if !isExistsPath {
				os.Mkdir(filePathTemp, os.ModePerm)
			}
			filePathTemp += "/"
		}
		loggerfile := filepath.Join(filePathTemp, llog.logname)
		logfile, err := os.Create(loggerfile)
		if err != nil {
			return errors.New(fmt.Sprintln("Create File Error:", err))
		}
		llog.file = logfile
	}
	return nil
}

// 设置输出方式
func (llog *Logger) SetPut(boolean bool) {
	llog.writeFile = boolean
}

// 判断输出方式
func (llog *Logger) SetWrite() {
	if llog.writeFile == true {
		llog.loggerf.SetOutput(llog.file)
	} else {
		llog.loggerf.SetOutput(os.Stdout)
	}
}

//按照格式打印日志
func (llog *Logger) Debugf(format string, v ...interface{}) {
	if llog.level > DebugLevel {
		//log里的函数都自带有mutex
		//这里获取一次锁
		return
	}
	llog.loggerf.Printf("[DEBUG] "+format, v...)
}

//Infof
func (llog *Logger) Infof(format string, v ...interface{}) {
	if llog.level > InfoLevel {
		return
	}
	llog.loggerf.Printf("[INFO] "+format, v...)
}

//Warningf
func (llog *Logger) Warningf(format string, v ...interface{}) {
	if llog.level > WarningLevel {
		return
	}
	llog.loggerf.Printf("[WARNING] "+format, v...)
}

//Errorf
func (llog *Logger) Errorf(format string, v ...interface{}) {
	if llog.level > ErrorLevel {
		return
	}
	llog.loggerf.Printf("[ERROR] "+format, v...)
}

func (llog *Logger) Fatalf(format string, v ...interface{}) {
	if llog.level > FatalLevel {
		return
	}
	llog.loggerf.Fatalf("[FATAL] "+format, v...)
}

//标准输出
//Println存在输出会换行
//这里采用Print输出方式
func (llog *Logger) Debug(v ...interface{}) {
	if llog.level > DebugLevel {
		return
	}
	llog.loggerf.Print("[DEBUG] " + fmt.Sprintln(v...))
}

//Info
func (llog *Logger) Info(v ...interface{}) {
	if llog.level > InfoLevel {
		return
	}

	llog.loggerf.Print("[INFO] " + fmt.Sprintln(v...))
}

//Warning
func (llog *Logger) Warning(v ...interface{}) {
	if llog.level > WarningLevel {
		return
	}
	llog.loggerf.Print("[WARNING] " + fmt.Sprintln(v...))
}

//Error
func (llog *Logger) Error(v ...interface{}) {
	if llog.level > ErrorLevel {
		return
	}
	llog.loggerf.Print("[ERROR] " + fmt.Sprintln(v...))
}

func (llog *Logger) Fatal(v ...interface{}) {
	if llog.level > FatalLevel {
		return
	}
	llog.loggerf.Fatal("[FATAL] " + fmt.Sprintln(v...))
}
