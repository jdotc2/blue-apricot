<script>
  import { LottiePlayer } from "@lottiefiles/svelte-lottie-player";
  import { createEventDispatcher } from "svelte";

  export let isWeatherRunning = false;
  export let isPokeEggRunning = false;
  export let isPokeTypeRunning = false;
  export let isPokeballRunning = false;

  const dispatch = createEventDispatcher();


  const toggleWeatherGrid = () => {
    isWeatherRunning = !isWeatherRunning;
    dispatch("toggleWeather");
  };

  const togglePokeEggGrid = () => {
    isPokeEggRunning = !isPokeEggRunning
    dispatch("toggleEgg");
  }

  const togglePokeTypeGrid = () => {
    isPokeTypeRunning = !isPokeTypeRunning;
    dispatch("togglePokeType");
  }

  const togglePokeballGrid = () => {
    isPokeballRunning = !isPokeballRunning;
    dispatch("togglePokeball");
  }
</script>

<style type="text/scss">
  .grid-container {
    display: grid;
    grid-template-columns: auto auto auto;
    grid-gap: 30px;
    padding: 2px;
    justify-content: center;
    align-items: center;
    padding-bottom: 3%;
    text-align: center;
  }

  .grid-container > div {
    text-align: center;
    padding: 20px 0;
    display: block;
    position: relative;
    width: 100%;
    border-collapse: collapse;
    text-decoration: none;
    color: #fff;
    padding: 25%;
    border-radius: 50%;
  }

  .box {
    display: block;
    position: relative;
    width: 100%;
    height: auto;
    border-collapse: collapse;
    text-decoration: none;
    color: #fff;
    padding: 25%;
    border-radius: 50%;
  }

  .box:before {
    transition: all 0.4s;
    position: absolute;
    content: "";
    width: 100%;
    height: 50%;
    left: 50%;
    margin-left: -50%;
    top: 25%;
    border-color: rgba(200, 200, 200, 0);
    border-style: solid;
    border-width: 0 1px;
    border-radius: 50%;
  }

  .box:after {
    transition: all 0.4s;
    position: absolute;
    content: "";
    width: 50%;
    height: 100%;
    left: 50%;
    margin-left: -25%;
    top: 0;
    border-color: rgba(200, 200, 200, 0);
    border-style: solid;
    border-width: 1px 0;
    border-radius: 50%;
  }

  .box:hover {
    color: #4099ff;

    &:after {
      width: 100%;
      margin-left: -50%;
      border-color: rgba(90, 90, 90, 1);
      border-radius: 50%;
    }

    &:before {
      height: 100%;
      top: 0%;
      border-color: rgba(90, 90, 90, 1);
      border-radius: 50%;
    }
  }

  .box:active {
    box-shadow: 0 0px 0px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    transition-duration: 4.1s;
    transform: translateY(1px);
  }

  img {
    width: 35px;
  }
</style>

<div class="grid-container">
  <div class="box" on:click={toggleWeatherGrid}>
    <img src="https://img.icons8.com/color/48/000000/chance-of-storm.png" />
  </div>
  <!-- <div class="box" on:click={togglePokeEggGrid}>
    <img src="https://img.icons8.com/color/48/000000/egg-pokemon.png" />
  </div> -->
  <div class="box" on:click={togglePokeTypeGrid}>
    <img src="https://vignette.wikia.nocookie.net/pokemongo/images/4/43/Fairy.png/revision/latest/scale-to-width-down/340?cb=20161013132820" />
  </div> 
  <div class="box" on:click={togglePokeballGrid}>
    <img src="https://cdn.iconscout.com/icon/premium/png-256-thumb/pokeball-536029.png" />
  </div>
</div>

<!-- {#if !isWeatherRunning}

{:else}
  <div class="grid-container">
    <div class="box" on:click={toggleRunning}>
      <img src="https://img.icons8.com/color/48/000000/chance-of-storm.png" />
    </div>
    <div class="box" on:click={toggleRunning}>
      <img src="https://img.icons8.com/color/48/000000/egg-pokemon.png" />
    </div>
    <div class="box" on:click={toggleRunning}>
      <img src="https://vignette.wikia.nocookie.net/pokemongo/images/4/43/Fairy.png/revision/latest/scale-to-width-down/340?cb=20161013132820" />
    </div>
    <div class="box" on:click={toggleRunning}>
      <img src="https://cdn.iconscout.com/icon/premium/png-256-thumb/pokeball-536029.png" />
    </div>
  </div>
{/if} -->
