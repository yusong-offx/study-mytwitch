<script>
  import {onMount} from 'svelte';
  import Chat from './Chat.svelte'

  onMount (()=>{
    let url = "https://yusong-offx.link/api/streamer/video/adaptive.mpd";
    let player = dashjs.MediaPlayer().create();
    let player_component = document.querySelector("#videoPlayer")
  
    player.initialize(player_component, url, true)
  })

</script>

<div class="container">
  <div>
    <video autoplay id="videoPlayer" controls>
      <track kind="captions" />
    </video>
  </div>

  <Chat></Chat>
</div>

<style>
.container {
  box-sizing: border-box;
  height: 100vh;
  width: 95vw;
  padding: 10vh 10vw;
  display: grid;
  grid-template-columns: 8fr 2fr;
  column-gap: 1vw;
}

#videoPlayer {
  position: relative;
  width: 100%;
}
</style>