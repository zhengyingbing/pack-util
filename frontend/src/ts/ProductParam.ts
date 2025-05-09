export default class ChannelParam {

    private _id: number
    private _productId: string
    private _productName: string


    constructor(productId: string, productName: string) {
        this._productId = productId;
        this._productName = productName;
    }

    get id(): number {
        return this._id;
    }

    set id(value: number) {
        this._id = value;
    }

    get productId(): string {
        return this._productId;
    }

    set productId(value: string) {
        this._productId = value;
    }

    get productName(): string {
        return this._productName;
    }

    set productName(value: string) {
        this._productName = value;
    }
}