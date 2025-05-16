import {ProductParam} from "./ProductParam";
import {get, writable} from "svelte/store";
import {Print} from "../../wailsjs/go/main/App";
import { produce } from 'immer';
import {ChannelParam} from "./ChannelParam";
export let channelParamsStore = writable<ChannelParam[]>([])
export let productParamsStore = writable<ProductParam[]>([])

initProductData()

export function getPackParamsData(productId) {
    channelParamsStore.set([])
    if (productId === "3015") {
        const a = {appName: "", channelId: "10219", channelName: "hoolai", channelDesc: "hoolai_万达", packageName: "com.wanda.sf3.wd", version: "1.0.5.5", isChecked: false, status: 0, progress: 0}
        const b = {appName: "", channelId: "10221", channelName: "xiaomi", channelDesc: "xiaomi_小米", packageName: "com.hahd.aygd3.mi", version: "3.4.3", isChecked: false, status: 0, progress: 0}
        const c = {appName: "", channelId: "10222", channelName: "huawei", channelDesc: "huawei_华为", packageName: "com.wanda.sf3.huawei", version: "6.13.0.300", isChecked: false, status: 0, progress: 0}
        const d = {appName: "", channelId: "10223", channelName: "oppo", channelDesc: "oppo_OPPO", packageName: "com.wanda.sf3.nearme.gamecenter", version: "1.0.4", isChecked: false, status: 0, progress: 0}
        const e = {appName: "", channelId: "10224", channelName: "vivo", channelDesc: "vivo_VIVO", packageName: "com.wanda.sf3.vivo", version: "4.7.8.0", isChecked: false, status: 0, progress: 0}
        let initialData1 = [a,b,c,d,e]
        channelParamsStore.set(initialData1)
    }

    return [...get(channelParamsStore)]
}

export function getProductData() {
    return [...get(productParamsStore)]
}

function initProductData() {
    const a1 = {productId:"11", productName:"demo测试"}
    const a2 = {productId:"12", productName:"圣斗士"}
    const a3 = {productId:"3015", productName:"暗影格斗3"}
    const a4 = {productId:"14", productName:"胡莱三国2"}
    const a5 = {productId:"15", productName:"秦时明月：沧海"}
    const a6 = {productId:"16", productName:"英雄传说·闪之轨迹之我的弟弟是真命天子"}
    let initialData1 = [a1, a2, a3, a4, a5, a6]
    productParamsStore.set(initialData1)
    Print("获取产品信息:" + JSON.stringify(get(productParamsStore), null, 2))
}

export const packParamsActions = {
    add:(channelId, channelName, channelDesc, version) => {
        channelParamsStore.update(params => {
            const p = {appName: "", channelId: channelId, channelName: channelName, channelDesc: channelDesc, version: version, isChecked: false, status: 0, progress: 0, packageName: ""}
            return [...params, p]
        })
    },
    remove:(channelId) => {
        channelParamsStore.update(params => params.filter(u => u.channelId !== channelId))
    },
    updateChecked: (channelId: string, checked: string) => {
        channelParamsStore.update(params =>
            produce(params, draft => {
                const target = draft.find(p => p.channelId === channelId);
                if (target) {
                    Print("更新状态====" + target.channelName + " - " + channelId + "...." + checked)
                    target.checked = checked
                }
            })
        );

    },
    get: (cid) => {
        [...get(channelParamsStore)].filter(param => param.channelId === cid)
    }

}