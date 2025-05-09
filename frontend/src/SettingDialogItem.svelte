<script>
    export let itemTitle
    export let itemContent
    export let onClickItem = () => {}
</script>
<div class="item-layout">
    <div class="item-title">
        <span>{itemTitle}</span>
    </div>
    <div class="item-content">
        <div class="item-title">
        </div>
        <span class="item-span">{itemContent}</span>
        <button class="btn" on:click={onClickItem}>选择路径</button>
    </div>

</div>
<style>
    .item-layout {
        width: 80vh;
        text-align: start;
        margin-bottom: 10px;
    }
    .item-title {
        font-size: 12px;
        margin-bottom: 2px;
    }
    .item-content {
        display: flex;
        padding: 6px 10px;
        align-items: center;
        background-color: #f0f0f0;
        font-size: 14px;
        border: #1b83e7;
    }
    .item-content:hover{
        border-color: #1b83e7;
    }
    .item-span {
        flex: 1;
    }
    .btn {
        background: transparent;
        color: #1b83e7;
        border: none;
    }
</style>