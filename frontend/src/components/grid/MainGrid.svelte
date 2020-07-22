<script>
  import axios from "axios";
  import Grid from "svelte-grid";
  import gridHelp from "svelte-grid/build/helper/index.mjs";
  import { LottiePlayer } from "@lottiefiles/svelte-lottie-player";
  import { fade } from "svelte/transition";
  import map from "lodash.map";
  import NavbarGrid from "./NavbarGrid.svelte";
  import Widgetbar from "../widgets/Widgetbar.svelte";
  import Navbar from "../Navbar.svelte";
  import WeatherGrid from "./WeatherGrid.svelte";
  import PokeTypeGrid from "./PokeTypeGrid.svelte";
  import PokeballGrid from "./PokeballGrid.svelte";

  const id = () =>
    "_" +
    Math.random()
      .toString(36)
      .substr(2, 9);

  let weatherGridItem;
  let pokeTypeGridItem;
  let pokeballGridItem;
  let isWeatherVisible = false;
  let isPokeTypeVisible = false;
  let isPokeballVisible = false;


  let adjustAfterRemove = false;

  let cols = 5;

  let layout = [];

  let items = gridHelp.resizeItems(layout, cols);

  
  // Apply breakpoints
  let breakpoints = [[1100, 5], [800, 4], [530, 1]];

  function toggleWeather(event) {
    if (!isWeatherVisible) {
      isWeatherVisible = !isWeatherVisible;
      let newWeatherGridItem = gridHelp.item({ x: 0, y: 0, w: 10, h: 2, id: id(), component: WeatherGrid });
      weatherGridItem = newWeatherGridItem;
      items = gridHelp.appendItem(newWeatherGridItem, items, cols);
      return items;
    } else {
      remove(weatherGridItem, event);
      isWeatherVisible = !isWeatherVisible;
    }
  }

  function togglePokeEgg(event) {
    if (!isPokeEggVisible) {
      isPokeEggVisible = !isPokeEggVisible;
      let newPokeEggGridItem = gridHelp.item({ x: 0, y: 0, w: 10, h: 2, id: id(), component: PokeEggGrid });
      pokeEggGridItem = newPokeEggGridItem;
      items = gridHelp.appendItem(newPokeEggGridItem, items, cols);
      return items;
    } else {
      remove(pokeEggGridItem, event);
    }
  }

  function togglePokeType(event) {
    if (!isPokeTypeVisible) {
      isPokeTypeVisible = !isPokeTypeVisible;
      let newPokeTypeGridItem = gridHelp.item({ x: 0, y: 0, w: 5, h: 6, id: id(), component: PokeTypeGrid });
      pokeTypeGridItem = newPokeTypeGridItem;
      items = gridHelp.appendItem(newPokeTypeGridItem, items, cols);
      return items;
    } else {
      remove(pokeTypeGridItem, event);
    }
  }

  function togglePokeball(event) {
    if (!isPokeballVisible) {
      isPokeballVisible = !isPokeballVisible;
      let newPokeballGridItem = gridHelp.item({ x: 0, y: 0, w: 5, h: 5, id: id(), component: PokeballGrid });
      pokeballGridItem = newPokeballGridItem;
      items = gridHelp.appendItem(newPokeballGridItem, items, cols);
      return items;
    } else {
      remove(pokeballGridItem, event);
    }
  }

  // WORKING
  function remove(item, event) {
    items = items.filter(value => value.id !== item.id);
    if (adjustAfterRemove) {
      items = gridHelp.resizeItems(items, cols);
    }
  }
</script>

<style>
  .content {
    z-index: 5;
    width: 100%;
    height: 100%;
    color: black;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: large;
    /* border-radius: 5px; */
    border-radius: 5px 30%;
    background-color: white;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.29);
  }

  .lock-grid-nav {
    z-index: -1;
    position: absolute;
    bottom: 0;
    left: 0;
    padding: 10px 0;
    margin: 3px;
    border-radius: 0 8px;
    cursor: pointer;
    background-image: linear-gradient(
      to bottom left,
      rgba(255, 255, 255, 0),
      rgba(255, 255, 255, 0.1),
      rgba(255, 255, 255, 0.4),
      rgba(255, 255, 255, 1),
      rgba(77, 144, 250, 0),
      rgba(77, 144, 250, 0.3),
      rgba(77, 144, 250, 0.6),
      rgba(77, 144, 250, 1)
    );
    width: 500px;
    justify-content: left;
    height: 15px;
    display: flex;
  }

  .lock {
    display: grid;
    grid-template-columns: auto auto;
    grid-gap: 10px;
    position: absolute;
    left: 0;
    top: 0;
    align-items: center;
    padding: 8px 8px;
  }

  .close-grid-nav {
    z-index: -1;
    position: absolute;
    top: 0;
    right: 0;
    padding: 10px 0;
    margin: 3px;
    border-radius: 0 8px;
    cursor: pointer;
    background-image: linear-gradient(
      to top right,
      rgba(255, 255, 255, 0),
      rgba(255, 255, 255, 0.1),
      rgba(255, 255, 255, 0.4),
      rgba(255, 255, 255, 1),
      rgba(77, 144, 250, 0),
      rgba(77, 144, 250, 0.3),
      rgba(77, 144, 250, 0.6),
      rgba(77, 144, 250, 1)
    );
    width: 400px;
    justify-content: right;
    height: 15px;
    display: flex;
  }

  .close {
    position: absolute;
    right: 0;
    top: 0;
    align-items: center;
    padding: 6px 6px;
  }
</style>

<Navbar />
<Widgetbar 
  isWeatherRunning={isWeatherVisible} 
  isPokeTypeRunning={isPokeTypeVisible}
  isPokeballRunning={isPokeballVisible}
  on:toggleWeather={toggleWeather}
  on:togglePokeType={togglePokeType}
  on:togglePokeball={togglePokeball} />
<Grid bind:items let:item cols={10} rowHeight={100} gap={10} useTransform={true} >
  <div class="content" transition:fade={{ delay: 250, duration: 300 }}>
  <svelte:component this={item.component} />
    <span on:click={remove.bind(null, item)} class="close-grid-nav">
      <div class="close">
        <LottiePlayer
          src={'https://maxst.icons8.com/vue-static/landings/animated-icons/icons/delete/delete.json'}
          loop={true}
          hover={true}
          renderer="svg"
          background="transparent"
          height={18}
          width={18}
          controls={false}
          controlsLayout={null} />
      </div>
    </span>
    <span on:click={remove.bind(null, item)} class="lock-grid-nav">
      <div class="lock">
        <LottiePlayer
          src={'https://maxst.icons8.com/vue-static/landings/animated-icons/icons/unlock/unlock.json'}
          loop={true}
          hover={true}
          renderer="svg"
          background="transparent"
          height={18}
          width={18}
          controls={false}
          controlsLayout={null} />
      </div>
    </span>
  </div>
</Grid>
