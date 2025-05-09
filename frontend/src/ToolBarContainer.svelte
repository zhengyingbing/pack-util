<script>

    //拖拽窗口相关
    import {Print} from "../wailsjs/go/main/App.js";
    import {onMount} from "svelte";

    let isDragging = false
    let startX, startY=0
    let lastPos = { x: 0, y: 0 };

    onMount(async () => {
        const pos = await window.runtime.WindowGetPosition();
        lastPos = {x: pos.x, y: pos.y}
    })

    function startDrag(e) {
        isDragging = true
        startX = e.screenX
        startY = e.screenY
        window.addEventListener('mousemove', handleDrag)
        window.addEventListener('mouseup', stopDrag)
    }
    //拖拽事件处理
    async function handleDrag(e) {
        if (!isDragging) return;

        const dx = e.screenX - startX;
        const dy = e.screenY - startY;
        requestAnimationFrame(() => {
            setWindowPosition(dx, dy)
        })

        startX = e.screenX;
        startY = e.screenY;
    }
    //刷新窗口坐标
    async function setWindowPosition(dx, dy) {
        try {
            // Print("dx,dy = " + dx + ", " + dy)
            // Print("lastPos.x,lastPos.y = " + lastPos.x + ", " + lastPos.y)
            // 双重验证数值
            const newX = lastPos.x + dx
            const newY = lastPos.y + dy
            window.runtime.WindowSetPosition(newX, newY, {animate: true});
            lastPos = {x: newX, y: newY}
        } catch (err) {
            Print("获取窗口位置失败:" + err);
        }
    }

    function stopDrag(e) {
        isDragging = false
        window.removeEventListener('mousemove', handleDrag)
        window.removeEventListener('mouseup', stopDrag)
    }
</script>
<div id="title-layout" style="cursor: {isDragging?'grabbing':'grab'};" on:mousedown={startDrag}>
    <slot/>
</div>
<style>
    #title-layout {
        height: auto;
        z-index: 1;
        background: #c9dbeb;
        display: flex;
        flex-direction: row;
        align-items: center;
        user-select: none;
    }
</style>