<script>
    import {onMount} from "svelte";
    import icon_file from './assets/images/icon_file.png'
    import icon_ok from './assets/images/icon_ok.png'
    import {OpenFolder, Print} from "../wailsjs/go/main/App.js";

    export let outPath
    export let channelParam
    export let canClick
    //0：未选中，1：待打包，2：打包中，3：打包完成，4：打包失败

    export let handleItemClick = () => {}


    function handleFileClick(id) {
        OpenFolder(outPath)
    }

    //打包状态：未开始1，打包中2，打包成功0，打包失败-1
    onMount(async () => {

    })
</script>
<div class="info-layout" on:click={canClick ? handleItemClick : undefined}>
    <input class="info-checkbox" type="checkbox" disabled={!canClick} checked={channelParam.isChecked ? "checked" : ""} on:click|stopPropagation on:change={canClick ? handleItemClick : undefined}/>
    <span class="span">{channelParam.channelDesc}</span>
    <span class="span">{channelParam.channelId}</span>
    <span class="span">{channelParam.packageName}</span>
    <span class="span">{channelParam.version}</span>
    {#if channelParam.status === 0}
        <span></span>
    {:else if channelParam.status === 1}
        <span>待打包</span>
    {:else if channelParam.status === 2 || channelParam.status === 3}
        <div class="progress-container">
            <!--{#if channelParam.progress > 0}-->
<!--                <progress value={channelParam.progress} max="100">{channelParam.progress}%</progress>-->
<!--                <span class="progress-text">{channelParam.progress}%</span>-->
<!--            {/if}-->
            <progress value={channelParam.progress} max="100" class:running={channelParam.status === 2} class:completed={channelParam.status === 3}>{channelParam.progress}</progress>
            {#if channelParam.status === 3}
                <img style="height: 16px;" src={icon_ok}/>
            {:else}
                <span class="progress-text">{channelParam.progress}</span>
            {/if}
        </div>
<!--        <div class="progress-container">-->
<!--            <progress value="50" max="100" class:running={channelParam.status === 2} class:completed={channelParam.status === 3}>50%</progress>-->
<!--            {#if channelParam.status === 3}-->
<!--                <img style="height: 16px;" src={icon_ok}/>-->
<!--                {:else}-->
<!--                <span class="progress-text">50%</span>-->
<!--                {/if}-->
<!--        </div>-->

    {:else if channelParam.status === 4}
        <span class="progress-failed">打包失败</span>
        {/if}

    <div>
        <img class="img-file" src={icon_file} on:click|stopPropagation={handleFileClick(channelParam.channelId)} alt="打开文件"/>
    </div>
</div>
<style>
    .info-layout {
        display: inline-grid;
        grid-auto-flow: column;
        grid-template-columns: 0.8fr 1fr 1fr 2.4fr 1fr 1.8fr 1.2fr;
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
        position: relative;
        gap: 8px;
        min-width: 150px;
    }
    progress {
        top: 50%;
        transform: translateY(-50%);
        height: 8px;
        border-radius: 4px;
        margin-right: 4px;
    }

    progress::-webkit-progress-bar {
        background-color: #b1b1b1;
        border-radius: 4px;
    }

    progress.running::-webkit-progress-value {
        background-color: #5498f1;
        border-radius: 4px;
    }
    progress.completed::-webkit-progress-value {
        background-color: #53c746;
        border-radius: 4px;
    }
    .progress-text {
        font-size: 0.8em;
        position: absolute;
        top: 50%;
        transform: translateY(-50%);
    }
    .progress-failed {
        color: #f21a1a;
    }
    .img-file {
        align-content: center;
        width: 20px;
        height: 20px;
    }
</style>