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
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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

type ProgressImpl struct {
}

func (ProgressImpl) Progress(channelId string, num int) {
	log.Println("当前进度", strconv.Itoa(num)+"%")
}
func (a *App) GetWindowPosition() map[string]int {
	x, y := runtime.WindowGetPosition(a.ctx)
	print("获取当前坐标：", x, y)
	return map[string]int{"x": x, "y": y}
}
func (a *App) Start(name string) error {
	println("开始打包~")
	return nil
}

func (a *App) Print(msg string) {
	println("js日志：", msg)
}

// main.go
type Config struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (a *App) OpenDirectoryDialog() string {
	// 使用系统原生对话框
	path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择文件夹",
	})
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return ""
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

func (a *App) StartPack(params *packParams) error {
	fmt.Printf("%+v\n", params)
	buildRootPath := filepath.Join(params.RootPath, "build")
	buildPath := filepath.Join(buildRootPath, params.ProductId+"_"+params.ChannelId)
	if !utils.Exist(buildPath) {
		utils.CreateDir(buildPath)
	}
	cfg := make(map[string]string)
	cfg[models.AppName] = params.ProductName
	cfg[models.IconName] = "ic_launcher.png"
	cfg[models.TargetSdkVersion] = "30"
	cfg[models.DexMethodCounters] = "60000"
	cfg[models.BundleId] = params.PackageName
	cfg[models.Orientation] = "sensorPortrait"
	cfg[models.SignVersion] = "2"
	cfg[models.KeystoreAlias] = "aygd3"
	cfg[models.KeystorePass] = "aygd3123"
	cfg[models.KeyPass] = "aygd3123"
	cfg["appId"] = "614371"
	models.SetServerDynamic(params.ChannelId, cfg)

	utils.Remove(buildRootPath)
	utils.Copy(filepath.Join(params.RootPath, "config", params.ChannelId, "access.config"), filepath.Join(buildPath, "access.config"), true)
	utils.Copy(filepath.Join(params.RootPath, "config", params.ChannelId, "ic_launcher.png"), filepath.Join(buildPath, "ic_launcher.png"), true)
	utils.Copy(filepath.Join(params.RootPath, "config", params.ChannelId, "game.keystore"), filepath.Join(buildPath, "game.keystore"), true)
	println("路径=============：", buildRootPath)

	preParams := models.PreParams{
		JavaHome:    params.JavaPath,
		AndroidHome: params.AndroidPath,
		RootPath:    params.RootPath,
		ChannelName: params.ChannelName,
		ChannelId:   params.ChannelId,
		ProductId:   params.ProductId,
		ApkName:     params.ApkName,
		OutPutPath:  params.OutputPath,
		PackageName: params.PackageName,
		//HomePath:     buildRootPath,
		ApkPath: params.ApkPath,
		//ExpandPath:   filepath.Join(params.RootPath, "sdk", "expand"),
		KeystoreName: "game.keystore",
	}
	packaging.Execute(&preParams, &ProgressImpl{}, &models.LogImpl{})
	return nil
}
