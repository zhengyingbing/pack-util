<script>
    import { navigate } from "svelte-navigator";
    import {ResizeWindow} from "../wailsjs/go/main/App.js";
    import icon_logo from './assets/images/icon_logo4.gif'
    // import {setupWindowDrag} from './utils/windowDrag.js'
    // import {onMount} from "svelte";
    //
    // // let dragElement
    // onMount(async () => {
    //     const el = document.getElementById('login-container')
    //     setupWindowDrag(el)
    // })

    function Login() {
        ResizeWindow(1024, 768)
        navigate("/app")
    }
</script>
<div class="login" id="login-container">
    <img alt="logo" src="{icon_logo}" style="height: 200px;"/>
    <button class="login-btn" on:click={Login}>登录</button>
</div>
<style>
    :global(html), :global(body), :global(#app) {
        margin: 0;
        width: 100%;
        height: 100%;
        overflow: hidden;
    }
    .login {
        width: 100vw;
        height: 100vh;
        background: #ffffff;
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 20px;
    }
    .login-btn {
        background: transparent url('./assets/images/icon_login.png');
        background-size: cover;
        border-width: 0;
        width: 200px;
        height: 60px;
    }
</style>