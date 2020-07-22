<script>
  import { onMount, createEventDispatcher } from "svelte";
  import { LottiePlayer } from "@lottiefiles/svelte-lottie-player";
  import { scale } from "svelte/transition";
  import { cubicOut } from "svelte/easing";

  import WeatherCard from "../cards/WeatherCard.svelte";

  let allWeather = [];

  const id = () => "_weather";

  onMount(async () => {
    const res = await fetch("http://localhost:4747/weather");
    let weathers = await res.json();
    allWeather = await weathers.data;
    console.log(allWeather);
  });

</script>

<style>
  .weather-grid {
    display: inline-grid;
    grid-template-columns: auto auto auto auto auto auto auto auto auto;
    grid-gap: 2px;
    height: auto;
  }
</style>

<div  transition:scale={{ duration: 500, delay: 500, opacity: 0.5, start: 0.5, easing: cubicOut }}>

  {#each allWeather as weather}
    <div class="weather-grid" >
      <WeatherCard sprite={weather.sprite}>
        <span slot="name">{weather.name}</span>
        <span slot="date">{weather.date}</span>
      </WeatherCard>
    </div>
  {:else}
    <p>loading...</p>
  {/each}

</div>
