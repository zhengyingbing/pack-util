<script>
    import App from "./App.svelte";
    import Login from "./Login.svelte";
    import {Route, Router} from "svelte-navigator";

</script>
<Router>
    <Route path="/" component="{Login}"/>
    <Route path="/app" component="{App}"/>
</Router>