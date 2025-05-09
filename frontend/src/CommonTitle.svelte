<script>
    export let onClose = () => {}
    export let title
    export let showClose = false
</script>
<div class="title-header">
    <h4 class="title-content">{title}</h4>
    {#if showClose}
        <button class="title-close" on:click={onClose()}>关闭</button>
    {/if}
</div>
<style>
    .title-header {
        display: flex;
        align-items: center;
        position: relative;
    }
    .title-content {
        flex: 1;
        text-align: center;
        margin: 0;
    }
    .title-close {
        margin-left: auto;
        padding: 4px 8px;
    }
</style>