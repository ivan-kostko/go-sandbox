package Logger

import (
	"fmt"
)

func GetStdTerminalLogger() ILogger {
	return GetNewLogContainer(stdLogToTerminal)
}

func stdLogToTerminal(level Level, args ...interface{}) {
	fmt.Println(level.String(), args)
}
