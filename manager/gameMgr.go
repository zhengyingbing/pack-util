package manager

import (
	"changeme/api"
	"github.com/zhengyingbing/common-utils/common/utils"
	"log"
)

type GameInfo struct {
	Id           int            `json:"gameId"`       // 游戏id
	Name         string         `json:"gameName"`     // 游戏名
	BundleId     string         `json:"bundleId"`     // 游戏包名
	SdkVersion   string         `json:"sdkVersion"`   // sdk版本
	Architecture string         `json:"architecture"` // 游戏支持架构
	ChannelInfos []*ChannelInfo // 渠道信息
}

type ChannelInfo struct {
	Id          int            `json:"channelId"`   // 渠道id
	Name        string         `json:"channelName"` // 渠道名
	Flag        string         `json:"flag"`        // 渠道
	Icon        string         `json:"icon"`        // icon url
	Sign        string         `json:"sign"`        // 签名文件路径
	DynamicCfg  map[string]any `json:"dynamicCfg"`  // 渠道动态配置
	Description string         `json:"description"` // 渠道描述
	OptionInfo  map[string]any `json:"optionInfo"`  // 扩展配置
}

func init() {
	gameMgr = GameMgr{}
}

var gameMgr GameMgr

type GameMgr struct {
	gameInfos []*GameInfo
}

func GameMgrInstance() *GameMgr {
	return &gameMgr
}

func (gameMgr *GameMgr) initGameInfo(data interface{}) {
	log.Println("7", utils.Struct2Json(data))
	_ = utils.Json2Struct(utils.Struct2Json(data), &gameMgr.gameInfos)
}

func (gameMgr *GameMgr) AllGameInfo() []*GameInfo {
	log.Println("开始获取产品信息！！")
	games := api.AllGames()
	gameMgr.initGameInfo(games)
	log.Println("获取产品结果：", gameMgr.gameInfos)
	return gameMgr.gameInfos
}
