package Logger

import (
	"reflect"
	"testing"
)

func TestLoggerContainerPrintLikeFunctions(t *testing.T) {
	originalArgs := []interface{}{"Arg1", "Arg2", "Arg3", 4}
	expectedArgs := []interface{}{"Arg1", "Arg2", "Arg3", 4}
	var actualArgs []interface{}

	var expectedLevel Level
	var actualLevel Level

	ilf := func(level Level, args ...interface{}) { actualLevel = level; actualArgs = args }

	l := GetNewLogContainer(ilf)

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

	l := GetNewLogContainer(ilf)

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
