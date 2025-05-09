export default class ChannelParam {

    private _id: number
    private _channelId: string
    private _channelName: string
    private _version: string
    //打包状态：未开始1，打包中2，打包成功0，打包失败-1
    private _status: number
    private _progress: number
    private _totalTime: string


    constructor(channelId: string, channelName: string, version: string, status: number, progress: number) {
        this._channelId = channelId;
        this._channelName = channelName;
        this._version = version;
        this._status = status;
        this._progress = progress;
    }

    get id(): number {
        return this._id;
    }

    set id(value: number) {
        this._id = value;
    }

    get channelId(): string {
        return this._channelId;
    }

    set channelId(value: string) {
        this._channelId = value;
    }

    get channelName(): string {
        return this._channelName;
    }

    set channelName(value: string) {
        this._channelName = value;
    }

    get version(): string {
        return this._version;
    }

    set version(value: string) {
        this._version = value;
    }

    get status(): number {
        return this._status;
    }

    set status(value: number) {
        this._status = value;
    }

    get progress(): number {
        return this._progress;
    }

    set progress(value: number) {
        this._progress = value;
    }

    get totalTime(): string {
        return this._totalTime;
    }

    set totalTime(value: string) {
        this._totalTime = value;
    }


}