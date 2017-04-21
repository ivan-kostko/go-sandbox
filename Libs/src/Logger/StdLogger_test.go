package Logger

import (
	"testing"
)

func ExampleStdLogger(t *testing.T) {

	t.SkipNow()

	l := GetStdTerminalLogger()
	l.Alert("TestAlert")
	l.Alertf("TestAlertf %v", "Extra")
	l.Emergency("TestEmergency")
	l.Emergencyf("TestEmergencyf %v", "Extra")
	l.Critical("TestCritical")
	l.Criticalf("TestCriticalf %v", "Extra")
	l.Error("TestError")
	l.Errorf("TestErrorf %v", "Extra")
	l.Warning("TestWarning")
	l.Warningf("TestWarningf %v", "Extra")
	l.Notice("TestNotice")
	l.Noticef("TestNoticef %v", "Extra")
	l.Info("TestInfo")
	l.Infof("TestInfof %v", "Extra")
	l.Debug("TestDebug")
	l.Debugf("TestDebugf %v", "Extra")
	l.Log(None, "TestLog")
	l.Logf(None, "TestLogf %v", "Extra")
}
