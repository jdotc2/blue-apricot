<script>
import { onMount, createEventDispatcher } from 'svelte'
import { LottiePlayer } from "@lottiefiles/svelte-lottie-player";
import { scale } from 'svelte/transition';
import { cubicOut } from 'svelte/easing';
import axios from 'axios'

import WeatherCard from '../cards/WeatherCard.svelte'

let allWeather = [];
export let weather = { visible: false };

const id = () => "_weather";

onMount(async () => {
  const res = await fetch('http://localhost:4747/weather')
  let weathers = await res.json()
  allWeather = await weathers.data
  console.log(allWeather)
})

const dispatch = createEventDispatcher();

const toggle = () => dispatch('toggle')

const toggleWeather = () => { 
  if (weather.visible) {
    weather.visible = !weather.visible;
    dispatch('toggle');
  } else {
    weather.visible = !weather.visible;
    dispatch('toggle');
    }
  }
</script>

<div class="photos" transition:scale="{{duration: 500, delay: 500, opacity: 0.5, start: 0.5, easing: cubicOut}}">
  
  {#each allWeather as weather}
    <div class="weather-grid">
      <WeatherCard>
        <LottiePlayer
          src="{weather.sprite}"
          loop={true}
          autoplay={true}
          renderer="svg"
          background="transparent"
          height={30}
          width={50}
          controls={false}
          controlsLayout={null} />
        <span slot="name">{weather.name}</span>
        <span slot="date">{weather.date}</span>
      </WeatherCard>
    </div>
  {:else}
  <p>loading...</p>
  {/each}

</div>


<style>

.weather-grid {
  display: inline-grid;
  grid-template-columns: auto auto auto auto auto auto auto auto auto;
  grid-gap: 2px;
  height: auto;
}

.lottie-container {
  width: 100px;
}

</style>