<script>
    export let buttonName
    export let onClick = () => {}
</script>
<div>
    <button class="common-button" on:click={onClick}>{buttonName}</button>
</div>
<style>
    .common-button {
        background: #1b83e7;
        color: white;
        padding: 6px 10px;
        border: none;
        border-radius: 4px;
    }
    .common-button :hover {
        background: #0d5194;
    }
    .common-button:active {
        background: #0d5194;
    }
</style>