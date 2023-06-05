package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"istep/env"
	"log"
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
	if env.CfgWay.Value == env.COM_SHUABU {
		if err := chromedp.Run(ctx,
			//chromedp.Emulate(device.IPadPro11),
			chromedp.Navigate(env.CfgShuabuHost.Value),
			chromedp.WaitVisible("app", chromedp.ByID),
		); err != nil {
			log.Fatal(err)
		}
	} else if env.CfgWay.Value == env.COM_17BUSHU {
		if err := chromedp.Run(ctx,
			//chromedp.Emulate(device.IPadPro11),
			chromedp.Navigate(env.Cfg17BushuHost.Value),
			chromedp.WaitVisible("submitBtn", chromedp.ByID),
		); err != nil {
			log.Fatal(err)
		}
	}
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
	}()
	for {
		cmd := <-commands
		if env.CfgWay.Value == env.COM_SHUABU {
			changeStepNumberWithShuabu(ctx, cmd)
		} else if env.CfgWay.Value == env.API {
			changeStepNumberWithApi(cmd)
		} else {
			changeStepNumberWith17bushu(ctx, cmd)
		}
	}
}
