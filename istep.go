package main

import (
	"encoding/json"
	"fmt"
	"github.com/ns-cn/goter"
	"github.com/spf13/cobra"
	"istep/env"
	"math/rand"
	"os"
	"sync"
	"time"
)

// 读取配置文件并解析为结构体数组
var configuration = new(Configuration)
var wg = sync.WaitGroup{}

func main() {
	root := goter.NewRootCmdWithAction("istep", "基于zepp life修改步数的程序", env.VERSION, func(command *cobra.Command, strings []string) {
		fmt.Printf("【istep】%v读取配置文件%s\n", time.Now(), env.CfgFile.Value)
		way := env.CfgWay.Value
		if way != env.COM_17BUSHU && way != env.COM_SHUABU && way != env.API {
			fmt.Println("不被支持的方式")
			_ = command.Help()
			return
		}
		// 读取配置文件
		data, err := os.ReadFile(env.CfgFile.Value)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, &configuration)
		if err != nil {
			panic(err)
		}
		fmt.Println(configuration.Users)
		keeper := make(chan int, 1)
		commands := make(chan ChangeCommand, 1)
		go func() {
			for {
				Client(keeper, commands, configuration.ShowClient)
			}
		}()
		for _, user := range configuration.Users {
			if user.Ignore {
				continue
			}
			wg.Add(1)
			// 根据步数的最大最小算一个随机结果
			stepNumber := rand.Intn(user.Steps.Max-user.Steps.Min+1) + user.Steps.Min
			commands <- ChangeCommand{
				UserID:     user.ID,
				Password:   user.getPassword(),
				StepNumber: stepNumber,
			}
		}
		wg.Wait()
	})
	root.Bind(&env.CfgFile, &env.Cfg17BushuHost, &env.CfgShuabuHost, &env.CfgWay)
	_ = root.Execute()
}
