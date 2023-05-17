package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"istep/env"
	"log"
	"strconv"
	"time"
)

func Client(keeper chan int, commands chan ChangeCommand, showClient bool) {
	keeper <- 1
	var ctx context.Context
	var cancel context.CancelFunc
	defer func() {
		<-keeper
	}()
	if showClient {
		// 创建一个chrome配置实例
		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
		)
		ctx, cancel = chromedp.NewExecAllocator(context.Background(), opts...)
		defer cancel()
		ctx, cancel = chromedp.NewContext(ctx)
		defer cancel()
	} else {
		ctx, cancel = chromedp.NewContext(context.Background())
		defer cancel()
	}
	// 访问网站并等待页面加载完成
	if err := chromedp.Run(ctx,
		//chromedp.Emulate(device.IPadPro11),
		chromedp.Navigate(env.CfgHost.Value),
		chromedp.WaitVisible("submitBtn", chromedp.ByID),
	); err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
	}()
	for {
		cmd := <-commands
		changeStepNumberWithApi(cmd)
	}
}

func changeStepNumber(ctx context.Context, cmd ChangeCommand) {
	defer func() {
		wg.Done()
	}()
	fmt.Printf("【修改步数任务】: 为%s修改步数为%d", cmd.UserID, cmd.StepNumber)
	err := chromedp.Run(ctx,
		chromedp.Navigate(env.CfgHost.Value),
		chromedp.WaitVisible("submitBtn", chromedp.ByID),
		chromedp.SendKeys("#phone", cmd.UserID),
		chromedp.SendKeys("#password", cmd.Password, chromedp.ByID),
		chromedp.SendKeys("#steps", strconv.Itoa(cmd.StepNumber)),
		chromedp.Click("#submitBtn", chromedp.ByID),
		chromedp.Sleep(5*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var failMsg string
			_ = chromedp.Run(ctx, chromedp.Text("#failMsg", &failMsg, chromedp.ByID))
			if failMsg == "" {
				fmt.Print("\t修改步数成功\n")
			} else {
				fmt.Printf("\t修改步数失败: %s(请注意密码编码)\n", failMsg)
			}
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
}
