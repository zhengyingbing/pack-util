<script>
    import CommonButton from "./CommonButton.svelte";

    export let isOpenDialog = false
    export let onClose = () => {}
    export let onSave = () => {}

    function onClose1() {
        isOpenDialog = false
    }

    function onSave1() {
        isOpenDialog = false
    }

    function handleDialogClick(e) {
        e.stopPropagation()
    }
</script>
{#if isOpenDialog}
    <div class="dialog-backdrop">
        <div class="dialog" on:click={handleDialogClick}>
            <div class="dialog-content">
                <slot/>
            </div>
            <div class="dialog-actions">
                <slot name="actions"/>
            </div>
            <div class="dialog-bottom">
                <CommonButton onClick={onClose} buttonName="取消"/>
                <div style="width: 10px"/>
                <CommonButton onClick={onSave} buttonName="保存设置"/>
            </div>
        </div>
    </div>
{/if}
<style>
    .dialog-backdrop {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background-color: rgba(0,0,0,0.2);
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 1000;
    }
    .dialog {
        background: white;
        border-radius: 8px;
        padding: 20px;
        min-width: 300px;
        box-shadow: 0 2px 6px rgb(155, 147, 147);
    }
    .dialog-actions {
        display: flex;
        justify-content: center;
        margin: 0;
    }
    .dialog-actions > * {
        width: 100%;
    }
    .dialog-bottom {
        display: flex;
        justify-content: flex-end;
    }

</style>