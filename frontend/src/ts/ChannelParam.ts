export interface ChannelParam {
    appName: string
    channelId: string
    channelName: string
    channelDesc: string
    version: string
    //0：未选中，1：待打包，2：打包中，3：打包完成，4：打包失败
    status: number
    progress: number
    isChecked: boolean
    packageName: string

}