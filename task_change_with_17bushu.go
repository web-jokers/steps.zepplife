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

func changeStepNumberWith17bushu(ctx context.Context, cmd ChangeCommand) {
	defer func() {
		wg.Done()
	}()
	fmt.Printf("【修改步数任务】: 为%s修改步数为%d", cmd.UserID, cmd.StepNumber)
	err := chromedp.Run(ctx,
		chromedp.Navigate(env.Cfg17BushuHost.Value),
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
