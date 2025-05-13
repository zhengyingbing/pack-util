<script>
  import {onMount} from 'svelte'
  import Dialog from './SettingDialog.svelte'
  import DialogItem from './SettingDialogItem.svelte'
  import CommonTitle from "./CommonTitle.svelte";
  import icon_menu from './assets/images/icon_menu.png'
  import icon_logo from './assets/images/icon_logo.png'
  import icon_exit from './assets/images/icon_exit.png'
  import icon_minimize from './assets/images/icon_minimize.png'
  import {
    LoadConfig,
    OpenDirectoryDialog,
    OpenFolder,
    Print,
    SaveConfig,
    SelectApk,
    StartPack
  } from "../wailsjs/go/main/App.js"
  import ToolBarContainer from "./ToolBarContainer.svelte";
  import PackInfoItem from "./PackInfoItem.svelte";
  import {channelParamsStore, getPackParamsData, getProductData} from "./ts/HttpUtils.ts";
  import {productParamsStore} from "./ts/HttpUtils.ts";

  const OUTPUT_PATH = "output_path";
  const JAVA_PATH = "java_path";
  const ANDROID_PATH = "android_path";
  const BUILD_PATH = "build_path";

  let openMenuId = 0
  let selectedMenu = null
  let selectedItem = null

  let javaPath, androidPath, buildPath, outputDirPath
  let folderPath = ''; // 存储选择的文件夹路径
  let isOpenSettingDialog = false

  //选择产品相关
  let isOpenProductList = false
  // let gameName = '请选择游戏'

  //是否正在打包
  let isPackaging

  //母包路径
  let gameApkPath = ""

  let isDebugMode = false
  let menus = [
    {id: 1, name:"文件", icon:""},
    {id: 2, name:"重新签名", icon:"icon_menu", items: []},
    {id: 3, name:"母包审核", icon:"icon_menu", items: []},
    {id: 4, name:"网络代理", icon:"icon_menu", items: []},
    {id: 5, name:"静态检测", icon:"icon_menu", items: []},
    {id: 6, name:"自动化检", icon:"icon_menu", items: []},
    {id: 8, name:"设置", icon:"icon_menu", items: ["清空缓存", "路径设置"]},
    {id: 9, name:"帮助", icon:"icon_menu", items: ["更新记录", "关于"]},
    {id: 10, name:"个人中心", icon:"icon_menu", items: ["用户信息", "注销"]},
  ]

  let productId = ""
  let productName = ""

  onMount(async () => {
    outputDirPath = (await LoadConfig(OUTPUT_PATH)).value
    javaPath = (await LoadConfig(JAVA_PATH)).value
    androidPath = (await LoadConfig(ANDROID_PATH)).value
    buildPath = (await LoadConfig(BUILD_PATH)).value
  })

  //获取渠道数据
  async function addData() {

  }
  function toggleMaximise() {
    window.runtime.WindowMinimise();
  }
  function closeApp() {
    window.runtime.Quit();
  }
  function toggleDropdown() {
    isOpenProductList = !isOpenProductList
    getProductData()
  }

  function selectGame(id, name) {
    productId = id
    productName = name
    isOpenProductList = false
    getPackParamsData(productId)
  }

  function handleClickOutside(event) {
    if (!event.target.closest('.dropdown-container')) {
      isOpenProductList = false
      openMenuId = 0
    }
  }

  // const onProgress: ProgressCallback = (channelId, num) => {
  //   Print("打包进度：" + num)
  //   channelParamsStore.update(items =>
  //           items.map(item =>
  //                   item.channelId === channelId ? {...item, progress: num} : item
  //           ))
  // }

  async function startPack() {
    Print("一共选中了" + selectedItems.length + "项.")

    if(buildPath === "" || javaPath === "" || androidPath === ""){
      alert("请检查路径配置是否正确")
      return
    }
    if (gameApkPath === "") {
      alert("请先选择母包")
      return
    }
    if (selectedItems.length > 0) {
      for (const item of selectedItems) {
        const params = {
          JavaPath:     javaPath,
          AndroidPath:  androidPath,
          RootPath: buildPath,
          ApkPath:     gameApkPath,
          OutputPath: outputDirPath,
          ProductName: productName,
          ProductId: productId,
          ApkName: productName + "_" + item.channelName + "_" + productId + "_" + item.channelId + "_" + item.version,
          ChannelName:      item.channelName,
          ChannelId:    item.channelId,
          PackageName: item.packageName
        }
        StartPack(params)
      }
      isPackaging = true
      channelParamsStore.update(items =>
              items.map(item => item.isChecked ? {...item, statusContent: "打包中"} : item)
      )
    } else {
      alert("请先选择渠道")
    }

    // isPackaging = true
    // await new Promise(resolve => setTimeout(resolve, 2000))
    // isPackaging = false
    // channelParamsStore.update(items =>
    //   items.map(item => item.isChecked ? {...item, statusContent: "打包完成", progress: 100} : item)
    // )
  }

  //选择母包
  async function selectGameApk(){
    gameApkPath = await SelectApk()
  }

  function selectMenu(menu) {
    Print("点击了" + menu.name + ", " + menu.id + ", " + menu.items)
    switch (menu.id) {
      case 1:
        alert("特殊处理" + menu.name)
        break;
      case 7:
        alert("当前是开发模式，请谨慎打包")
        break;
      default:
        openMenuId = menu.id
    }
    selectedMenu = menu
  }

  async function selectMenuItem(item) {
    // Print(item)
    selectedItem = item
    openMenuId = 0
    switch (item) {
      case "路径设置":
        outputDirPath = (await LoadConfig(OUTPUT_PATH)).value
        javaPath = (await LoadConfig(JAVA_PATH)).value
        androidPath = (await LoadConfig(ANDROID_PATH)).value
        buildPath = (await LoadConfig(BUILD_PATH)).value
        isOpenSettingDialog = true
        break;
      case "清空缓存":
        isPackaging = false
        break;
      case "更新记录":
        alert("这是最新")
        break;
      case "关于":
        alert("这是个打包工具")
        break;
    }
  }

  function closeDialog() {
    isOpenSettingDialog = false
  }
  //设置输出路径
  async function setOutputPath() {
    outputDirPath = await OpenDirectoryDialog()
    Print("outputDirPath = " + outputDirPath)
  }
  async function setAndroidPath() {
    androidPath = await OpenDirectoryDialog()
  }
  async function setJavaPath() {
    javaPath = await OpenDirectoryDialog()
  }
  async function setBuildPath() {
    buildPath = await OpenDirectoryDialog()
  }
  function savePath() {
    SaveConfig({"Key": JAVA_PATH, "Value": javaPath})
    SaveConfig({"Key": ANDROID_PATH, "Value": androidPath})
    SaveConfig({"Key": BUILD_PATH, "Value": buildPath})
    SaveConfig({"Key": OUTPUT_PATH, "Value": outputDirPath})
    isOpenSettingDialog = false
  }
  // 打开已保存的文件夹
  async function openOutputDir() {
    OpenFolder(outputDirPath)
  }

  function toggleAll(checked) {
    channelParamsStore.update(items => {
      const newItems = items.map(item => ({...item, isChecked: checked}));
      return newItems;
    })
  }

  function toggleItem(channelId) {
    Print("选中..." + channelId)
    channelParamsStore.update(items =>
      items.map(item =>
        item.channelId === channelId ? {...item, isChecked: !item.isChecked} : item
      )
    )
  }

  $: allChecked = $channelParamsStore.every(item => item.isChecked)

  //selectedItems是派生变量不是Store，没有update方法，需要直接修改channelParamsStore
  $: selectedItems = $channelParamsStore.filter(item => item.isChecked)
</script>
<svelte:window on:click={handleClickOutside}/>
<main>
  <Dialog isOpenDialog={isOpenSettingDialog} onClose="{closeDialog}"  onSave={savePath}>
    <CommonTitle title="路径设置"></CommonTitle>

    <div slot="actions">
      <DialogItem itemTitle="Java路径" itemContent={javaPath} onClickItem={setJavaPath}/>
      <DialogItem itemTitle="Android路径" itemContent={androidPath} onClickItem={setAndroidPath}/>
      <DialogItem itemTitle="编译路径" itemContent={buildPath} onClickItem={setBuildPath}/>
      <DialogItem itemTitle="输出路径" itemContent={outputDirPath} onClickItem={setOutputPath}/>
    </div>
  </Dialog>
  <div class="parent-layout" style="display: flex; flex-direction: column; height: 100vh">
    <ToolBarContainer>
      <img alt="logo" src="{icon_logo}" style="height: 24px; margin-left: 20px" on:dblclick={() => isDebugMode = !isDebugMode} on:click={addData}/>
      <div id="title-box">
        {#each menus as menu}
          <div class="dropdown-container" id="dropdown-menu">
            <div>
              <button class="menu-button {selectedMenu?.id === menu.id ? 'selected':''}" on:click={() => selectMenu(menu)}>{menu.name}</button>
              {#if menu.id === openMenuId}
                <div id="menu_list">
                  {#each menu.items as item}
                    <div class="dropdown-item" on:click={() => selectMenuItem(item)}>
                      {item}
                    </div>
                  {/each}
                </div>
              {/if}
            </div>
          </div>

        {/each}
      </div>

      <img class="toolbar-btn" alt="logo" src="{icon_minimize}" on:click={toggleMaximise}/>
      <img class="toolbar-btn" alt="logo" src="{icon_exit}" on:click={closeApp} />
    </ToolBarContainer>

    <div id="operate-box">
      <div class="dropdown-container" id="dropdown-product">
        <div style="height: fit-content;">
          <button class="dropdown-button" on:click={toggleDropdown}>
            {productName || '请选择游戏'}
          </button>
          {#if isOpenProductList}
            <div id="product_list">
              {#each $productParamsStore as game}
                <div class="dropdown-item" id="product_item" on:click={() => selectGame(game.productId, game.productName)}>
                  {game.productName}
                </div>
              {/each}
            </div>
          {/if}
        </div>
      </div>
      <span style="background: #ff4f4f; border-radius: 4px; color: white">{#if isDebugMode}开发者模式{/if}</span>
      <div class="input-box">
        <span id="gameApkName">{gameApkPath}</span>
        <button class="btn" on:click={selectGameApk}>选择母包</button>
        <button class="button-start" on:click={startPack} disabled={isPackaging}>{isPackaging ? "打包中" : "开始打包"}</button>
        <button class="btn" on:click={openOutputDir}>输出目录</button>
      </div>
    </div>

    <div id="center-box">

      <div class="center-header">
        <!--{#each listHeader as item}-->
          <!--{#if item.type === 1}-->
          <!--  <span class="span" id="center-type">{item.content}</span>-->
          <!--{/if}-->
          <!--{#if item.type === 4}-->
          <!--  <input class="checkbox" type="checkbox" id="myCheckbox" checked={isCheckAll}/>-->
          <!--{/if}-->
        <!--{/each}-->
        <input class="checkbox" type="checkbox" id="myCheckbox" checked={allChecked} on:change={() => toggleAll(!allChecked)}/>
        <span class="span">渠道名称</span>
        <span class="span1">渠道ID</span>
        <span class="span1">包名</span>
        <span class="span1">版本号</span>
        <span class="span">打包进度</span>
        <span class="span1">打包状态</span>
        <span class="span2">包文件</span>
      </div>
      {#each $channelParamsStore as item}
        <PackInfoItem channelParam={item} outPath={outputDirPath} handleItemClick={() => toggleItem(item.channelId)}>

        </PackInfoItem>
        {/each}

    </div>
  </div>


</main>

<style>

  .parent-layout {
    color: #333333;
    font-size: 14px;
  }
  .parent-layout:hover{
    background: #8b8b8b;
  }
  #title-box {
    flex: 1;
    padding-left: 40px;
    padding-right: 40px;
    display: flex;
    justify-items: center;
  }
  #dropdown-menu {
    position: relative;
    display: inline-block;
  }
  #menu_list {
    width: 100px;
    background: white;
    position: absolute;
    top:100%;
    left: 0;
    border-radius: 4px;
    z-index: 900; /*确保在最上层*/
    text-align: left;
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  }
  .toolbar-btn {
    width: 20px;
    height: 20px;
    padding: 6px;
    margin-right: 20px;
    border: transparent;
    border-radius: 4px;
  }
  .toolbar-btn:hover{
    background: #bdcede;
  }
  .menu-button {
    height: 40px;
    border: transparent;
    background: transparent;
    color: #161616;
    font-size: 14px;
  }

  .menu-button:hover {
    background: #bdcede;
    border-radius: 4px;
  }


  #operate-box {
    display: flex;
    background: #ffffff;
    padding-left: 20px;
    padding-right: 20px;
    border-bottom: #d9d9d9 solid 1px;
    gap: 10px;
    justify-content: left;
    align-items: center;
  }

  .dropdown-container {
    position: relative;
    display: inline-block;
    margin-top: 12px;
    margin-bottom: 12px;
  }

  #product_list {
    width: 100%;
    max-height: 200px;
    overflow-y: auto;
    background: white;
    position: absolute;
    top:100%;
    left: 0;
    border: 1px solid #ddd;
    border-radius: 4px;
    z-index: 1000; /*确保在最上层*/
    text-align: left;
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  }
  .dropdown-item {
    align-content: center;
    height: 36px;
    padding-left: 10px;
    border-top: #eae6e6 solid 1px;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    direction: ltr;
  }
  .dropdown-item:hover {
    background: #dcdcdc;
  }
  .dropdown-button {
    height: 36px;
    width: 200px;
    padding-left: 10px;
    background: #d9d9d9;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    text-align: left;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    direction: ltr;
  }
  .dropdown-button::after {
    content: '▼';
    position: absolute;
    right: 10px;
  }

  .input-box {
    display: inline-flex;
    justify-content: right;
    flex: 1;
    gap: 20px; /* 控制元素间距 */
    overflow: hidden; /* 防止内容溢出容器 */
  }
  #gameApkName {
    align-content: center;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    direction: ltr;
    text-align: right;
  }

  #operate-box .btn {
    background: #eeeeee;
    padding-left: 14px;
    padding-right: 14px;
    height: 40px;
    border: none;
    border-radius: 4px;
    flex-shrink: 0;
  }
  #operate-box .btn:hover {
    background: #dcdcdc;
  }
  .button-start {
    width: 100px;
    height: 40px;
    background: dodgerblue;
    color: #f0f0f0;
    border: none;
    border-radius: 4px;
    flex-shrink: 0;
  }
  .button-start:hover {
    background: #1b83e7;
  }
  .button-start:active {
    background: #1b83e7;
  }
  .button-start:disabled {
    background: #a2aab8;
  }


  #center-box {
    flex: 1;
    background: #dddddd;
    overflow-y: auto;

  }
  .center-header {
    width: 100%;
    height: 50px;
    display: inline-grid;
    grid-auto-flow: column;
    grid-template-columns: 0.8fr 1fr 1fr 2.4fr 1fr 1.8fr 1.2fr 0.8fr;
    align-items: center;
    font-size: 14px;
  }

  .center-header .checkbox {
    width: 20px;
    height: 20px;
    margin-left: 40px;
  }
  .center-header .span {
    align-content: center;
  }
  #center-type {
    padding-bottom: 10px;
    padding-top: 10px;
  }
</style>
