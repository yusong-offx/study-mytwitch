<script>
    import { onMount } from 'svelte'
    import { count } from "./store"

    let user_id = false
    let name = "world"
    onMount(async () => {
        await fetch("https://yusong-offx.link/api", {
            method: "GET",
            cache: "no-cache",
            credentials: "include",
            redirect: 'follow'
        })
        .then((resp) => resp.json())
        .then((data) => {
            user_id = data.user_id
            if (user_id) {
                name = user_id
                count.set(name)
            }
        })
    })

</script>

<svelte:head>
  <script src="https://cdn.dashjs.org/latest/dash.all.min.js" ></script>
</svelte:head>

<div>
    <h1 class="hello">Hello!</h1>
    <h1 class="world">{name} :)</h1>
    <nav>
        {#if user_id == false}
            <a href="/login">LOGIN</a>
            <a href="/signup">SIGNUP</a>
        {:else}
            <a href="/logout">LOGOUT</a>
        {/if}
            <a href="/streaming">WATCH</a>

    </nav>
</div>

<style>
    h1 {
        text-align: left;
        animation: fadeInDown 1s;
    }
    @keyframes fadeInDown {
        0% {
            opacity: 0;
            transform: translate3d(0, 150%, 0);
        }
        50% {
            transform: translate3d(0, -40%, 0);
        }
        to {
            opacity: 1;
            transform: translateZ(0);
        }
    }
    h1.hello::first-letter {
        font-size: 12vh;
    }
    h1.hello {
        font-size: 10vh;
    }
    h1.world {
        font-size: 10vh;
        padding-left: 0.3vw;
    }
    div {
        height: 100%;
        padding: 25vh 25vw;
    }
    a {
        font-size: 1.3vw;
        color:rgb(176, 130, 219);
        margin: 0.4vw;
        text-decoration: none;
        animation: fadeIN 1s;
    }
    @keyframes fadeIN {
        0% {
            opacity: 0;
        }
        to {
            opacity: 1;
        }
    }
    a:hover {
        color:wheat;
    }
</style>