<script>

    //拖拽窗口相关
    import {Print} from "../wailsjs/go/main/App.js";
    import {onMount} from "svelte";
    import {WindowGetPosition} from "../wailsjs/runtime/runtime.js";

    let isDragging = false
    let startX=0, startY=0
    let lastPos = { x: 0, y: 0 };

    onMount(async () => {
        const pos = await window.runtime.WindowGetPosition();
        lastPos = {x: pos.x, y: pos.y}
    })

    function startDrag(e) {
        e.preventDefault();
        e.stopPropagation();

        isDragging = true
        startX = e.screenX
        startY = e.screenY

        WindowGetPosition().then(pos => {
            lastPos = {x: pos.x, y: pos.y}
            startX = e.screenX;
            startY = e.screenY;
        }).catch(err => {

        })
        document.addEventListener('mousemove', handleDrag)
        document.addEventListener('mouseup', stopDrag)

        document.body.style.userSelect = 'none'
    }
    //拖拽事件处理
    async function handleDrag(e) {
        if (!isDragging) return;

        e.preventDefault()

        const dx = e.screenX - startX;
        const dy = e.screenY - startY;

        // const newX = lastPos.x + dx
        // const newY = lastPos.y + dy
        requestAnimationFrame(() => {
            setWindowPosition(dx, dy)
        })


    }
    //刷新窗口坐标
    async function setWindowPosition(dx, dy) {
        try {
            // Print("dx,dy = " + dx + ", " + dy)
            // Print("lastPos.x,lastPos.y = " + lastPos.x + ", " + lastPos.y)
            // 双重验证数值
            const newX = lastPos.x + dx
            const newY = lastPos.y + dy
            Print(`移动偏移: dx=${dx}, dy=${dy}`);
            Print(`新位置: x=${newX}, y=${newY}`);
            await window.runtime.WindowSetPosition(newX, newY);
            // lastPos = {x: newX, y: newY}
        } catch (err) {
            Print("获取窗口位置失败:" + err);
        }
    }

    function stopDrag(e) {
        isDragging = false
        document.removeEventListener('mousemove', handleDrag)
        document.removeEventListener('mouseup', stopDrag)
        document.body.style.userSelect = ''
        WindowGetPosition().then(pos => {
            lastPos = {x: pos.x, y: pos.y}
        }).catch(err => {

        })
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