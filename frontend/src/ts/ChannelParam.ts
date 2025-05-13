export interface ChannelParam {
    appName: string
    channelId: string
    channelName: string
    channelDesc: string
    version: string
    //打包状态：未开始1，打包中2，打包成功0，打包失败-1
    statusContent: string
    progress: number
    isChecked: boolean
    packageName: string

}