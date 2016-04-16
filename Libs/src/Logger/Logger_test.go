package Logger

import (
	"reflect"
	"testing"
)

func TestLevelStringer(t *testing.T) {
	if None.String() != "None" {
		t.Errorf("None.String() returned %v while expected None", None.String())
	}
	if Emergency.String() != "Emergency" {
		t.Errorf("Emergency.String() returned %v while expected Emergency", Emergency.String())
	}
	if Critical.String() != "Critical" {
		t.Errorf("Critical.String() returned %v while expected Critical", Critical.String())
	}
	if Error.String() != "Error" {
		t.Errorf("Error.String() returned %v while expected Error", Error.String())
	}
	if Warning.String() != "Warning" {
		t.Errorf("Warning.String() returned %v while expected Warning", Warning.String())
	}
	if Notice.String() != "Notice" {
		t.Errorf("Notice.String() returned %v while expected Notice", Notice.String())
	}
	if Info.String() != "Info" {
		t.Errorf("Info.String() returned %v while expected Info", Info.String())
	}
	if Debug.String() != "Debug" {
		t.Errorf("Debug.String() returned %v while expected Debug", Debug.String())
	}
}

func TestLoggerContainerPrintLikeFunctions(t *testing.T) {
	originalArgs := []interface{}{"Arg1", "Arg2", "Arg3", 4}
	expectedArgs := []interface{}{"Arg1", "Arg2", "Arg3", 4}
	var actualArgs []interface{}

	var expectedLevel Level
	var actualLevel Level

	ilf := func(level Level, args ...interface{}) { actualLevel = level; actualArgs = args }

	l := GetNewLogAdapter(ilf)

	// Just Print like functions

	expectedLevel = None
	l.Log(None, originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Log passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Emergency
	l.Emergency(originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Emergency passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Alert
	l.Alert(originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Emergency passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Critical
	l.Critical(originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Critical passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Error
	l.Error(originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Error passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Warning
	l.Warning(originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Warning passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Notice
	l.Notice(originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Notice passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Info
	l.Info(originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Info passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Debug
	l.Debug(originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Debug passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}
}

func TestLoggerContainerPrintFLikeFunctions(t *testing.T) {
	originalArgs := []interface{}{"Arg1", "Arg2", "Arg3", 4}
	formatString := " %s %s %s %v "
	expectedArgs := []interface{}{" Arg1 Arg2 Arg3 4 "}
	var actualArgs []interface{}

	var expectedLevel Level
	var actualLevel Level

	ilf := func(level Level, args ...interface{}) { actualLevel = level; actualArgs = args }

	l := GetNewLogAdapter(ilf)

	// Just PrintF like functions

	expectedLevel = None
	l.Logf(None, formatString, originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Logf passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Emergency
	l.Emergencyf(formatString, originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Emergency passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Alert
	l.Alertf(formatString, originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Emergency passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Critical
	l.Criticalf(formatString, originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Critical passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Error
	l.Errorf(formatString, originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Error passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Warning
	l.Warningf(formatString, originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Warning passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Notice
	l.Noticef(formatString, originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Notice passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Info
	l.Infof(formatString, originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Info passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}

	expectedLevel = Debug
	l.Debugf(formatString, originalArgs...)
	if actualLevel != expectedLevel || !reflect.DeepEqual(actualArgs, expectedArgs) {
		t.Errorf("l.Debug passed %s %v when expected %s %v", actualLevel, actualArgs, expectedLevel, expectedArgs)
	}
}
