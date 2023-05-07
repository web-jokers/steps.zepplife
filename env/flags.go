package env

import "github.com/ns-cn/goter"

var (
	CfgFile = goter.NewCmdFlagString("istep.json", "load", "l", "config file")
)
