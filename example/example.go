package example

import "logger"

func Handle() {
	logger.Debugf("the Debugf test %s", "DEBUGF")
	logger.Infof("the Infof test %s", "INFOF")
	logger.Warningf("test Warningf test %s", "WARNINGF")
	logger.Errorf("test Errorf test %s", "ERRORF")

	logger.Debug("the Debug test" + "DEBUG")
	logger.Info("the Info test" + "INFO")
	logger.Warning("test Warning test" + "WARNING")
	logger.Error("test Error test" + "ERROR")
}
