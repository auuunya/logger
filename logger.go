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

var llog Logger

func init() {
	// 日志设置
	llog := NewLogger()
	llog.SetLevel("error")
	llog.SetLogPath(fmt.Sprintf("logs/log_%s.log", time.Now().Format("20060102")))
}

//定义日志级别
const (
	LevelError   = iota // 0
	LevelWarning        // 1
	LevelInfo           // 2
	LevelDebug          // 3
)

//定义日志结构体
type Logger struct {
	level     int
	file      *os.File
	loggerf   *log.Logger
	writeFile bool
	path      string
}

// Errorf 默认日志对象方法，记录一条错误日志，需要先初始化
func Errorf(format string, v ...interface{}) {
	llog.Errorf(format, v...)
}

// Error 默认日志对象方法，记录一条消息日志，需要先初始化
func Error(args ...interface{}) {
	llog.Error(args...)
}

// Infof 默认日志对象方法，记录一条消息日志，需要先初始化
func Infof(format string, v ...interface{}) {
	llog.Infof(format, v...)
}

// Info 默认日志对象方法，记录一条消息日志，需要先初始化
func Info(args ...interface{}) {
	fmt.Println("args：", args)
	llog.Info(args...)
}

// Debugf 默认日志对象方法，记录一条消息日志，需要先初始化
func Debugf(format string, v ...interface{}) {
	llog.Debugf(format, v...)
}

// Debug 默认日志对象方法，记录一条调试日志，需要先初始化
func Debug(args ...interface{}) {
	llog.Debug(args...)
}

//Waring
func Waring(format string, v ...interface{}) {
	llog.Warningf(format, v...)
}

// Waringf
func Waringf(args ...interface{}) {
	llog.Warning(args...)
}

// logging 接口
type logging interface {
	Init(path, logname string, level int) error
	fileWrite(time time.Time, msg interface{}) error
	SetLevel(level string)
	GetLevel() int
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

// 获取等级
func (llog *Logger) GetLevel() int {
	return llog.level
}

//等级设置
func (llog *Logger) SetLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		llog.level = LevelDebug
	case "info":
		llog.level = LevelInfo
	case "warning":
		llog.level = LevelWarning
	case "error":
		llog.level = LevelError
	}
}

//时间转字符串
func formatTime(t *time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func NewLogger() *Logger {
	loger := &Logger{
		writeFile: false,
	}
	loger.SetLevel("debug")
	loger.loggerf = log.New(os.Stdout, "Logger_Stdout_", log.Llongfile|log.Ltime|log.Ldate)
	return loger
}

//创建日志文件
//制定创建路径和日志名
func (llog *Logger) SetLogPath(path string) {
	llog.path = path
	llog.SetLogFile()
	llog.SetPut(true)
}

// 创建日志文件
func (llog *Logger) SetLogFile() error {
	pathList := strings.Split(llog.path, "/")
	pathLen := len(pathList)
	filePrefix := pathList[pathLen-1]
	_, filebool := IsFile(llog.path)
	filePathTemp := ""
	if filebool {
		logname, err := os.OpenFile(llog.path, os.O_APPEND, 0666)
		if err != nil {
			return errors.New(fmt.Sprintln("Open File Error:", err))
		}
		llog.file = logname
	} else {
		for _, v := range pathList[:pathLen-1] {
			filePathTemp += v
			_, isExistsPath := IsExists(filePathTemp)
			if !isExistsPath {
				os.Mkdir(filePathTemp, os.ModePerm)
			}
			filePathTemp += "/"
		}
		logname := filepath.Join(filePathTemp, filePrefix)
		logfile, err := os.Create(logname)
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
	llog.SetWrite()
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
	if llog.level > LevelDebug {
		//log里的函数都自带有mutex
		//这里获取一次锁
		return
	}
	llog.loggerf.Printf("[DEBUG] "+format, v...)
}

//Infof
func (llog *Logger) Infof(format string, v ...interface{}) {
	if llog.level > LevelInfo {
		return
	}
	llog.loggerf.Printf("[INFO] "+format, v...)
}

//Warningf
func (llog *Logger) Warningf(format string, v ...interface{}) {
	if llog.level > LevelWarning {
		return
	}
	llog.loggerf.Printf("[WARNING] "+format, v...)
}

//Errorf
func (llog *Logger) Errorf(format string, v ...interface{}) {
	if llog.level > LevelError {
		return
	}
	llog.loggerf.Printf("[ERROR] "+format, v...)
}

//标准输出
//Println存在输出会换行
//这里采用Print输出方式
func (llog *Logger) Debug(v ...interface{}) {
	if llog.level > LevelDebug {
		return
	}
	llog.loggerf.Print("[DEBUG] " + fmt.Sprintln(v...))
}

//Info
func (llog *Logger) Info(v ...interface{}) {
	fmt.Println("111111", llog.level > LevelInfo)
	if llog.level > LevelInfo {
		return
	}

	llog.loggerf.Print("[INFO] " + fmt.Sprintln(v...))
}

//Warning
func (llog *Logger) Warning(v ...interface{}) {
	if llog.level > LevelWarning {
		return
	}
	llog.loggerf.Print("[WARNING] " + fmt.Sprintln(v...))
}

//Error
func (llog *Logger) Error(v ...interface{}) {
	if llog.level > LevelError {
		return
	}
	llog.loggerf.Print("[ERROR] " + fmt.Sprintln(v...))
}
