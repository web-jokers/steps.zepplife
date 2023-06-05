package env

import "github.com/ns-cn/goter"

const (
	COM_17BUSHU = "17bushu"
	COM_SHUABU  = "shuabu"
	API         = "api"
)

var (
	CfgFile        = goter.NewCmdFlagString("istep.json", "load", "l", "加载的json配置文件")
	Cfg17BushuHost = goter.NewCmdFlagString("http://118.195.237.33/", "17bushu-host", "H", "使用17bushu方式时配置的服务器")
	CfgShuabuHost  = goter.NewCmdFlagString("https://shuabu.org/", "shuabu-host", "S", "使用shuabu方式时配置的服务器")
	CfgWay         = goter.NewCmdFlagString(COM_17BUSHU, "way", "W", "指定使用的方式，支持（17bushu、shuabu、api）三种方式")
)
