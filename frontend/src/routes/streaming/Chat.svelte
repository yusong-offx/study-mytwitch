<script>
    import { onMount } from 'svelte'
    import { count } from '../store'

    let conn;
    let msg;
    let chatroom;
    let websocket_url;

    count.subscribe(value => {
        websocket_url = "wss://yusong-offx.link/api/ws/".concat(value)
    })

    function appendMessage(data) {
        let item = document.createElement("div");
        item.innerHTML = data;

        let doScroll = chatroom.scrollTop > chatroom.scrollHeight - chatroom.clientHeight - 1;
        chatroom.appendChild(item);
        if (doScroll) {
            chatroom.scrollTop = chatroom.scrollHeight - chatroom.clientHeight;
        }
    }

    function submitOnEnter(event){
        if(event.which === 13){
            event.target.form.dispatchEvent(new Event("submit", {cancelable: true}));
            event.preventDefault();
        }
    }

    async function onSubmitHandler(e) {
        msg = e.target[0].value
        e.target[0].value = ""
        if (!conn) {
            return false;
        }
        if (!msg) {
            return false;
        }
        conn.send(msg);
        return false
    }

    onMount(()=>{

        chatroom = document.getElementById("chatroom")
        document.getElementById("usermsg").addEventListener("keypress", submitOnEnter);

        if (window["WebSocket"]) {
            conn = new WebSocket(websocket_url);
            conn.onclose = function (evt) {
                appendMessage("Connection closed.");
            };
            conn.onmessage = function (evt) {
                appendMessage(evt.data)
            };
        } else {
            appendMessage("Your browser does not support WebSockets.");
        }
    })
</script>

<div class="container">
    <div id="chatroom">
    </div>
    <form on:submit|preventDefault={onSubmitHandler}>
        <textarea id="usermsg" cols="30" rows="5"></textarea>
    </form>
</div>

<style>
    .container {
        display:grid;
        flex-direction:column;
        grid-template-rows: 9fr 1fr;
        overflow: hidden;
        /* background-color: aqua; */
        width: 100%;
        height: 85%;
    }
    .container div {
        width: 100%;
        word-break: break-all;
    }
    #chatroom {
        height: 100%;
        width: 100%;
        overflow: auto;
    }
    textarea {
        width: 95%;
        resize: none;
    }
</style>