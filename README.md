# study-mytwitch
### This is small project to study about web structure :)
You need to change host, if you want to execute project. (svelte fetch, nginx conf (server_name, not proxy)) <br>
When you compose up, "/share/..." directories are automatically made. <br>
You can refer "docker-compose.yml" file for ssl location.
<br>


# feature
- adaptive live video streaming
- websocket chat
- login/out & signup (password hashed by bcrypt)
- reverse proxy by nginx
<br>

# skill
- docker-compose
- svelte
- fiber (golang)
- postgresql (pgadmin4)
<br>

# structure

![structure image](./README/structure.svg) <br>
\* There is only table USERS which has id and password columns.

<details>
<summary>site images</summary>

![main page](./README/main.png)
![login&signup page](./README/login&signup.png)
![watch page](./README/watch.png)
</details>
<br>

# next
- reduce live streaming delay through hardware accelation
- use normal schema (data models)
- use strong authentication and authorize
- use redis (for jwt, cache...)
- centralize log
- service by k8s
- video spreads by http3
<br>
