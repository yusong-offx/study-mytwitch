<script>
    function onSubmitHandler(e) {
        const data = {
            user_id: e.target[0].value,
            user_password: e.target[1].value,
        }

        fetch(
        'https://yusong-offx.link/api/login',
        {
            method: 'POST',
            cache: "no-cache",
            credentials: 'include',
            redirect: 'follow',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data)
        })
        .then((resp) => {
            if (resp.ok === false) {
                window.alert("check ID and PASSWORD again")
                throw 'something wrong'
            }
            window.location.replace("/")
        })
        .catch((err) => {
            console.log(err)
        })
    }
</script>

<div>
    <form on:submit|preventDefault={onSubmitHandler}>
        <label for="user-id">ID</label>
        <input id="user-id" type="text">
        <label for="user-password">PASSWORD</label>
        <input id="user-password" type="password">
        <button type="submit">login</button>
    </form>
</div>

<style>
    div {
        position: absolute;
        top: 40vh; left: 40vw;
        height: 12vh; width: 20vw;
        border-style: solid;
        padding: 1vh;
    }
    form {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        column-gap: 1vw;
        row-gap: 1vh;
    }
    label {
        text-align: center;
    }
    input {
        height: 2vh;
        width: 10vw;
    }
    button {
        justify-self: center;
        grid-column: 1 / 3;
        text-align: center;
        width: 100px;
        height: 23px;
    }
</style>