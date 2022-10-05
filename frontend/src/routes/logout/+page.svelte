<script>
    import { count } from '../store'

    const doing = async () => {
        await fetch("https://yusong-offx.link/api/logout", {
            method: "GET",
            cache: "no-cache",
            credentials: "include",
            redirect: 'follow',
        })
        .then((resp) => {
            if (resp.ok === false) {
                throw 'something wrong'
            }
            window.location.replace("/")
        })
        .catch((err) => {
            console.log(err)
        })
        count.set("")
    }
    doing()
</script>