<script>
  import { onMount } from "svelte";
  import axios from "axios";
  import MainGrid from "./components/grid/MainGrid.svelte";

  let weatherImage;

  onMount(async () => {
    try {
      let res = await axios.get("http://127.0.0.1:4747/weather");
      weatherImage = res.data.data[0].sprite;
      console.log(weatherImage);
    } catch (e) {
      console.error(e);
    }
  });
</script>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    /* margin: 0 auto; */
  }

  .blue {
    font-family: "Varela Round", sans-serif;
    color: #4d90fa;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  }

  h1 {
    font-family: "Varela Round", sans-serif;
    color: #6d7075;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
    display: inline-block;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>

<svelte:head>
  <title>Blue Apricot</title>
  <link
    href="https://fonts.googleapis.com/css2?family=Montserrat+Subrayada&display=swap"
    rel="stylesheet" />
  <link
    href="https://fonts.googleapis.com/css2?family=Raleway:wght@300&display=swap"
    rel="stylesheet" />
  <link
    href="https://fonts.googleapis.com/css2?family=Questrial&family=Raleway:wght@300&display=swap"
    rel="stylesheet" />
  <link
    href="https://fonts.googleapis.com/css2?family=Questrial&family=Raleway:wght@300&family=Varela+Round&display=swap"
    rel="stylesheet" />
  <link
    href="https://fonts.googleapis.com/css2?family=Saira+Condensed&display=swap"
    rel="stylesheet" />
</svelte:head>
<main>
  <div class="baHeader">
    <div class="container">
      <h1 class="blue">Blue</h1>
      <h1>Apric</h1>
      <img src="img/blue-apricot.png" alt="" />
      <h1>t</h1>
    </div>
  </div>
  <MainGrid />
</main>
