<script>
  export let name;
	import { onMount } from "svelte";
	import axios from 'axios'
	import Grid from 'svelte-grid'

	let image;

	onMount(async () => {
		try {
			let res =	await axios.get("http://127.0.0.1:4747/gen8")
			image = res.data.data[0].Images.main
			console.log(image)
		} catch (e) {
			console.error(e)
		}
	});
	
</script>

<svelte:head>
	<title>Blue Apricot</title>
	<link href="https://fonts.googleapis.com/css2?family=Montserrat+Subrayada&display=swap" rel="stylesheet">
</svelte:head>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

	.blue {
		font-family: 'Montserrat Subrayada', sans-serif;
		color: #4d90fa;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
	}

  h1 {
		font-family: 'Montserrat Subrayada', sans-serif;
    color: #6d7075;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
		display: inline-block
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>

<main>
	<div class="baHeader">
		<div class="container">
			<img src='img/blue-apricot.png' alt='' />
			<h1 class="blue">Blue</h1> <h1> Apricot</h1>
		</div>
	</div>
	<img src={image} alt='' />
</main>
