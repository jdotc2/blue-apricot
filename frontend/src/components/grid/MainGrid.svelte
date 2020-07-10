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

  const id = () =>
    "_" +
    Math.random()
      .toString(36)
      .substr(2, 9);

  let weatherGridItem;
  let isWeatherVisible = false;
  let isMapVisible = false;
  let isPokedexVisible = false;
  let weatherGrid = component => {
    return gridHelp.item({
      x: 0,
      y: 0,
      w: 10,
      h: 1,
      id: id(),
      component: component
    });
  };

  let adjustAfterRemove = false;

  let cols = 5;

  let items = [];
  // Apply breakpoints
  let breakpoints = [[1100, 5], [800, 4], [530, 1]];

  function handleWeatherClick(item, event) {
    if (!isWeatherVisible) {
      isWeatherVisible = !isWeatherVisible;
      console.log(isWeatherVisible);
      return gridHelp.item({
        x: 0,
        y: 0,
        w: 3,
        h: 1,
        id: id(),
        component: WeatherGrid
      });
    } else {
      isWeatherVisible = !isWeatherVisible;
      let newItem = gridHelp.item({
        x: 0,
        y: 0,
        w: 4,
        h: 7,
        id: id(),
        component: WeatherGrid
      });
      remove(item, event);
      return (items = gridHelp.appendItem(newItem, items, cols));
    }
  }

  function toggleWeather(event) {
    if (!isWeatherVisible) {
      isWeatherVisible = !isWeatherVisible;
      let newWeatherGridItem = gridHelp.item({
        x: 0,
        y: 0,
        w: 10,
        h: 2,
        id: id(),
        component: WeatherGrid
      });
      weatherGridItem = newWeatherGridItem;
      console.log(weatherGridItem);
      items = gridHelp.appendItem(newWeatherGridItem, items, cols);
      console.log(items);
      console.log(isWeatherVisible);
      return items;
    } else {
      isWeatherVisible = !isWeatherVisible;
      console.log(isWeatherVisible);
      remove(weatherGridItem, event);
      return;
    }
  }

  // WORKING
  function remove(item, event) {
    items = items.filter(value => value.id !== item.id);
    if (adjustAfterRemove) {
      items = gridHelp.resizeItems(items, cols);
    }
    isWeatherVisible = !isWeatherVisible;
  }
</script>

<style>
  .content {
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
<Widgetbar running={isWeatherVisible} on:toggle={toggleWeather} />
<Grid bind:items let:item cols={10} rowHeight={100} gap={10}>
  <div class="content" transition:fade={{ delay: 250, duration: 300 }}>
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

    <svelte:component
      this={item.component}
      visible={isWeatherVisible}
      on:toggle={handleWeatherClick} />
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
