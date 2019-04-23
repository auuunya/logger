package main

import (
	"fmt"
	"logger"
)

func main() {
	// loger, err := logger.Init("./log_test", "test.log", 3)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// loger := logger.NewLogger()
	loger, _ := logger.Init("./log_test", "logs.log", "error")
	loger.SetPut(false)
	loger.SetWrite()
	loger.SetLevel("error")
	fmt.Println("getlevel:", loger.GetLevel())
	loger.Debugf("the Debugf test %s", "DEBUGF")
	loger.Infof("the Infof test %s", "INFOF")
	loger.Warningf("test Warningf test %s", "WARNINGF")
	loger.Errorf("test Errorf test %s", "ERRORF")

	loger.Debug("the Debug test" + "DEBUG")
	loger.Info("the Info test" + "INFO")
	loger.Warning("test Warning test" + "WARNING")
	loger.Error("test Error test" + "ERROR")
}
