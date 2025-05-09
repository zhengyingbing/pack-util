import ChannelParam from "./ChannelParam";
import ProductParam from "./ProductParam";
import {get, writable} from "svelte/store";
import {Print} from "../../wailsjs/go/main/App";

export const packParamsStore = writable<ChannelParam[]>([])
export const productParamsStore = writable<ProductParam[]>([])

initProductData()

export function getPackParamsData(productId): ChannelParam[] {
    packParamsStore.set([])

    const a = new ChannelParam(productId + "010","hoolai_万达","1.5.5", 1, 0)
    const b = new ChannelParam(productId + "011", "xiaomi_小米", "2.4.3", 1, 0);
    const c = new ChannelParam(productId + "012", "huawei_华为", "6.300.9.16", 1, 0);
    const d = new ChannelParam(productId + "013", "oppo_OPPO", "4.8.0.8", 1, 0);
    const e = new ChannelParam(productId + "014", "vivo_VIVO", "1.5.5.0", 1, 0);
    const f = new ChannelParam(productId + "015", "uc_UC", "5.0.8", 1, 0);
    const g = new ChannelParam(productId + "016", "4399_4399", "7.0.7", 1, 0);
    const h = new ChannelParam(productId + "017", "hoolai_好游快爆", "4.9.0", 1, 0);
    let initialData1 = [a,b,c,d,e,f,g,h]
    packParamsStore.set(initialData1)
    return [...get(packParamsStore)]
}

export function getProductData(): ProductParam[] {
    return [...get(productParamsStore)]
}

function initProductData() {
    const a1 = new ProductParam("11", "demo测试")
    const a2 = new ProductParam("12", "圣斗士")
    const a3 = new ProductParam("13", "暗影格斗3")
    const a4 = new ProductParam("14", "胡莱三国2")
    const a5 = new ProductParam("15", "秦时明月：沧海")
    const a6 = new ProductParam("16", "英雄传说·闪之轨迹之我的弟弟是真命天子")
    let initialData1 = [a1, a2, a3, a4, a5, a6]
    productParamsStore.set(initialData1)
    Print("获取产品信息:" + JSON.stringify(get(productParamsStore), null, 2))
}

export const packParamsActions = {
    add:(channelId, channelName, version) => {
        packParamsStore.update(params => {
            const p = new ChannelParam(channelId,channelName,version, 1, 0)
            p.channelId = channelId
            p.channelName = channelName
            p.version = version
            return [...params, p]
        })
    },
    remove:(channelId) => {
        packParamsStore.update(params => params.filter(u => u.channelId !== channelId))
    },
    update: (channelId, channelName, version) => {
        const param = new ChannelParam(channelId,channelName,version, 1, 0)
        param.channelId = channelId
        param.version = version
        packParamsStore.update(params => params.map(u => u.channelId === channelId ? param : u))
    },
}