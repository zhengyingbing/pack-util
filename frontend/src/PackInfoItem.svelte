<script>
    import {onMount} from "svelte";
    import CommonButton from "./CommonButton.svelte";
    import icon_file from './assets/images/icon_file.png'
    import {OpenFolder, Print} from "../wailsjs/go/main/App.js";
    export let isChecked = false
    export let outPath
    export let channelName, channelId, version, totalTime

    let status = 1;
    export let progress
    // export let handleItemClick = (id) => {}
    // export let handleFileClick = (id) => {}
    function handleItemClick(id) {
        isChecked = !isChecked
    }
    function handleFileClick(id) {
        OpenFolder(outPath)
    }
    export let showProgress = false
    export let statusContent = ""

    //打包状态：未开始1，打包中2，打包成功0，打包失败-1
    export function updateStatus(packStatus) {
        Print("更新打包状态" + packStatus)
        status = packStatus
        if(status === 1){
            status = ""
        } else if(status === 2){
            showProgress = true
            statusContent = "打包中"
        } else if(status === 0){
            statusContent = "打包成功"
        } else if(status === -1){
            statusContent = "打包失败"
        }
    }
    onMount(async () => {

    })
</script>
<div class="info-layout" on:click={handleItemClick(channelId)}>
    <input class="info-checkbox" type="checkbox" checked={isChecked}/>
    <span class="span">{channelName}</span>
    <span class="span">{channelId}</span>
    <span class="span">{version}</span>
    <div class="progress-container">
        {#if showProgress}
            <progress value={progress} max="100">{progress}%</progress>
            <span class="progress-text">{progress}%</span>
            {/if}

    </div>

    <span class="status-content">{statusContent}</span>

    <span class="span">{totalTime}</span>
    <div>
        <img class="img-file" src={icon_file} on:click={handleFileClick(channelId)} alt="打开文件"/>
    </div>
</div>
<style>
    .info-layout {
        display: inline-grid;
        grid-auto-flow: column;
        grid-template-columns: 1fr 1fr 1fr 1fr 2fr 1fr 1fr 1fr;
        align-items: center;
        background: #ffffff;
        width: 100%;
        height: 50px;
        color: #333333;
        font-size: 14px;
        border-bottom: 1px solid #eeeeee;
    }
    .info-layout:hover {
        background: #efefef;
    }
    .info-layout .span {
        align-content: center;
    }
    .info-checkbox {
        width: 20px;
        height: 20px;
        margin-left: 40px;
    }
    /* 进度条样式 */
    .progress-container {
        align-items: center;
        gap: 8px;
        min-width: 150px;
    }
    progress {
        flex: 1;
        height: 8px;
        border-radius: 4px;
    }

    progress::-webkit-progress-bar {
        background-color: #b1b1b1;
        border-radius: 4px;
    }

    progress::-webkit-progress-value {
        background-color: #5498f1;
        border-radius: 4px;
    }
    .progress-text {
        width: 40px;
        text-align: right;
        font-size: 0.8em;
    }
    .img-file {
        align-content: center;
        width: 20px;
        height: 20px;
    }
</style>