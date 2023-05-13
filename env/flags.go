package env

import "github.com/ns-cn/goter"

var (
	CfgFile = goter.NewCmdFlagString("istep.json", "load", "l", "config file")
	CfgHost = goter.NewCmdFlagString("http://118.195.237.33/", "host", "H", "服务器")
)
