package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// 嵌入资源
//
//go:embed all:frontend/dist
var assets embed.FS

// App struct
type App struct {
	ctx context.Context
}

func (a *App) Login(username, password string) bool {
	return username == "admin" && password == "123123"
}

func (a *App) ResizeWindow(width, height int) {
	runtime.WindowSetSize(a.ctx, width, height)
}

func main() {
	app := &App{}
	//创建wails应用
	err := wails.Run(&options.App{
		Title:     "打包工具",
		Width:     400,
		Height:    600,
		MinWidth:  400,
		MinHeight: 600,
		MaxWidth:  1080,
		MaxHeight: 720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
		Frameless:     true, //禁用系统标题栏
		DisableResize: true, //尺寸保持不可变
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

// 捕获系统信号实现优雅退出
//func init() {
//	c := make(chan os.Signal, 1)
//	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
//	go func() {
//		<-c
//		if logSystem != nil {
//			logSystem.Shutdown()
//		}
//		os.Exit(0)
//	}()
//}
