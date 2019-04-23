package logger

import (
	"fmt"
	"testing"
)

var loger = &logger{}

func Test_logger_Init(t *testing.T) {
	err := loger.Init("./log", "logs.log", 1)
	fmt.Println(loger.level)
	if err != nil {
		t.Errorf("Loger Init Error %v", err)
	} else {
		t.Errorf("Logger Init Success")
	}
}

func Test_logger_Debugf(t *testing.T) {
	_ = loger.Init("./log", "logs.log", 1)
	fmt.Println("before:", loger.level)
	loger.setLevel(3)
	fmt.Println("after:", loger.level)
	loger.Debugf("The Debugf %s", "test")
}

func Test_logger_Debug(t *testing.T) {
	_ = loger.Init("./log", "logs.log", 1)
	fmt.Printf("%v", loger)
	fmt.Println("before:", loger.level)
	loger.setLevel(3)
	fmt.Println("after:", loger.level)
	loger.Debug("The Debugf test")
}
