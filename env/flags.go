package env

import "github.com/ns-cn/goter"

var (
	CfgFile = goter.NewCmdFlagString("ibushu.json", "load", "l", "config file")
)
