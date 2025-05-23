package api

type ApiType string

var Host = "https://service.wdyxgames.com/server_admin"

const (
	//Host = "http://center.hoolai.com/server_admin"
	//Host = "http://service.wdyxgames.com/server_admin"
	//Host = "http://10.1.65.16:4000"
	GlobalHost   = "https://global-service.wdyxgames.com/server_admin"
	InternalHost = "https://service.wdyxgames.com/server_admin"
	TestHost     = "http://center.hoolai.com/server_admin"

	PackageInfo             ApiType = "/public/package/packageToolsInfo" // 打包工具信息
	AuthLogin               ApiType = "/public/user/userAuth"            // 用户登录api
	Login                   ApiType = "/public/user/login"               // 用户登录api
	QRLogin                 ApiType = "/public/user/quickAuth"           // 二维码授权登录
	DownloadConfig          ApiType = "/channel/downloadCfg"             //下载渠道配置
	GameList                ApiType = "/package/gameList"                // 游戏列表api
	ChannelList             ApiType = "/package/channelList"             // 渠道列表api
	ChannelPackageCfg       ApiType = "/package/dynamicConfig"           // 渠道配置api
	ChannelOnlineSdkVersion ApiType = "/package/channelSdkVersion"       // 线上sdk版本信息
	ChannelDevSdkVersion    ApiType = "/package/channelDevSdkVersion"    // 开发sdk版本信息
	RsyncPackageChannel     ApiType = "/public/package/packageNotify"    // 同步打包信息
	FetchBaseApkInfo        ApiType = "/public/package/basePackageInfo"  // 查询母包信息
	ReviewBaseApk           ApiType = "/package/basePackageAudit"        // 审核母包
	AutoCheckConfig         ApiType = "/package/autoCheckConfig"         // 获取自动检测配置
)
