package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sqweek/dialog"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/zhengyingbing/common-utils/common/utils"
	"github.com/zhengyingbing/common-utils/packaging"
	"github.com/zhengyingbing/common-utils/packaging/models"
	util2 "github.com/zhengyingbing/common-utils/packaging/utils"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetWindowPosition() map[string]int {
	x, y := runtime.WindowGetPosition(a.ctx)
	print("获取当前坐标：", x, y)
	return map[string]int{"x": x, "y": y}
}

func (a *App) Clear(buildPath string) {
	utils.Remove(filepath.Join(buildPath, "build"))
}

func (a *App) Print(msg string) {
	println("js日志：", msg)
}

// main.go
type Config struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (a *App) OpenDirectoryDialog(p string) string {
	// 使用系统原生对话框
	path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择文件夹",
	})
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return p
	}
	return path
}

func (a *App) SelectApk() string {
	path, err := dialog.File().Title("选择母包").Filter(".apk", "apk").Load()
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return ""
	}
	return path
}

func (a *App) SaveConfig(config Config) error {
	// 保存到本地文件 (示例路径)
	//filePath := filepath.Join(runtime.ApplicationDataDirectory(a.ctx), "config.json")
	path, _ := os.Getwd()
	filePath := filepath.Join(path, "config.json")
	ab, _ := os.ReadFile(filePath)
	cfg := make(map[string]string, 0)
	json.Unmarshal(ab, &cfg)
	println("保存路径", config.Value)
	cfg[config.Key] = config.Value
	jsonData, _ := json.Marshal(cfg)
	return os.WriteFile(filePath, jsonData, 0644)
}

func (a *App) LoadConfig(key string) Config {
	path, _ := os.Getwd()
	filePath := filepath.Join(path, "config.json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}
	}

	var result map[string]string
	json.Unmarshal(data, &result)
	return Config{
		Key:   key,
		Value: result[key],
	}
}

func (a *App) OpenFolder(path string) {
	// 跨平台打开文件夹
	println("打开路径", path)
	switch utils.CurrentOsType() {
	case "windows":
		exec.Command("explorer", path).Start()
	case "darwin":
		exec.Command("open", path).Start()
	case "linux":
		exec.Command("xdg-open", path).Start()
	}
}

type ProgressImpl struct {
	ctx context.Context
}

func (p *ProgressImpl) Progress(channelId string, num int) {
	runtime.EventsEmit(p.ctx, "progress", map[string]interface{}{
		"channelId": channelId,
		"progress":  num,
	})
	log.Println("当前进度", strconv.Itoa(num)+"%")
}

func (a *App) Start(productParam models.ProductParam, channelParams []models.ChannelParam) error {
	util2.Write("您共选中的渠道数：", len(channelParams))
	packaging.Preparation(productParam, channelParams, &ProgressImpl{ctx: a.ctx})
	var wg sync.WaitGroup
	errChan := make(chan error, len(channelParams)) // 带缓冲的错误通道
	for _, channelParam := range channelParams {
		wg.Add(1)
		go func(cp models.ChannelParam) {
			defer wg.Done()
			buildPath := filepath.Join(productParam.RootPath, "build", productParam.ProductId+"_"+cp.ChannelId)
			apkName := strings.Join(productParam.ApkName, "_")
			apkName = strings.Replace(apkName, "channelId", cp.ChannelId, -1)
			apkName = strings.Replace(apkName, "channelName", cp.ChannelName, -1) + ".apk"
			preParams := models.PreParams{
				JavaHome:     productParam.JavaPath,
				AndroidHome:  productParam.AndroidPath,
				RootPath:     productParam.RootPath,
				BuildPath:    buildPath,
				ChannelName:  cp.ChannelName,
				ChannelId:    cp.ChannelId,
				ProductId:    productParam.ProductId,
				ProductName:  productParam.ProductName,
				ApkName:      apkName,
				OutPutPath:   productParam.OutputPath,
				PackageName:  cp.PackageName,
				ApkPath:      productParam.ApkPath,
				KeystoreName: "game.keystore",
			}
			//packaging.Execute(&preParams, &ProgressImpl{ctx: a.ctx}, &models.LogImpl{})
			logger := util2.Init(channelParam.ChannelId, productParam.RootPath)
			//defer logger.Shutdown()
			packaging.Execute(&preParams, &ProgressImpl{ctx: a.ctx}, logger)
		}(channelParam)

	}
	go func() {
		wg.Wait()
		close(errChan)
	}()
	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}
