<script>
  import { onMount, createEventDispatcher } from "svelte";
  import { scale } from "svelte/transition";
  import { quintOut } from "svelte/easing";
  import { crossfade } from "svelte/transition";
  import { flip } from "svelte/animate";

  let pokeballs = [];
  let pokeballName = "Pokeballs";
  let uid = pokeballs.length + 1;
  let hovering;
  let checked = false;

  $: name = pokeballName;

  const id = () => "_pokeball";

  onMount(async () => {
    const res = await fetch("http://localhost:4747/pokeballs");
    let pokeballsRes = await res.json();
    pokeballs = await pokeballsRes.data;
    console.log(pokeballs);
  });

  const [send, receive] = crossfade({
    fallback(node, params) {
      const style = getComputedStyle(node);
      const transform = style.transform === "none" ? "" : style.transform;

      return {
        duration: 100,
        easing: quintOut,
        css: t => `
					transform: ${transform} scale(${t});
					opacity: ${t}
				`
      };
    }
  });

  function enter(name, event) {
    hovering = true;
    pokeballName = name;
  }

  function leave() {
    hovering = false;
    pokeballName = "Pokeballs";
  }

  function add(input) {
    const pokeball = {
      id: uid++,
      done: false,
      description: input.value
    };

    pokeballs = [pokeball, ...pokeballs];
    input.value = "";
  }

  function remove(pokeball) {
    pokeballs = pokeballs.filter(t => t !== pokeball);
  }
</script>

<style>
  .board {
    display: grid;
    grid-template-columns: auto;
    grid-gap: 10px;
    padding: 10px;
    padding-bottom: 1%;
  }

  .top {
    float: left;
    box-sizing: border-box;
  }

  .top-grid {
    display: inline-grid;
    grid-template-columns: auto auto auto auto auto auto auto auto auto auto auto auto auto;
    grid-gap: 10px;
    justify-content: center;
    align-items: center;
    padding-top: 3%;
  }

  .bottom {
    float: left;
    width: 90%;
    padding: 0 1em 0 0;
  }

  .bottom-grid {
    display: grid;
    grid-template-columns: auto auto;
    grid-gap: 5px;
    justify-content: center;
    padding-top: 5%;
  }

  .rg-img-item {
    grid-column: 1;
    grid-row: 1;
    width: auto;
    justify-content: center;
    align-items: center;
    padding-right: 1%;
  }

  .rg-data-container {
    grid-column: 2 / span 3;
    grid-row: 1;
    text-align: left;
    font-family: "Raleway", sans-serif;
  }

  .title {
    font-family: "Raleway", sans-serif;
    font-size: 24px;
    border-bottom: 1px solid #6d7075;
    font-weight: bolder;
  }

  .meta-data {
    font-family: "Raleway", sans-serif;
    font-size: 12px;
    padding-top: 2%;
    padding-bottom: 2%;
  }

  .rg-location-container {
    grid-column: 2 / span 2;
    grid-row: 2;
    text-align: left;
    font-family: "Raleway", sans-serif;
  }

  .sub-header {
    font-family: "Raleway", sans-serif;
    font-size: 16px;
    border-bottom: 1px solid #6d7075;
    font-weight: bolder;
  }

  .rg-cost-container {
    grid-column: 4;
    grid-row: 2;
    text-align: left;
    font-family: "Raleway", sans-serif;
  }

  .rg-npc-container {
    grid-column: 2 / span 2;
    grid-row: 3;
    text-align: left;
    font-family: "Raleway", sans-serif;
  }

  .rg-npc-desc-container {
    grid-column: 4;
    grid-row: 3;
    text-align: left;
    font-family: "Raleway", sans-serif;
  }

  .active-pokeball-name {
    font-family: "Raleway", sans-serif;
    font-size: 2em;
    border-bottom: 1px solid #6d7075;
    border-width: 90%;
  }

  label {
    top: 0;
    left: 0;
    display: block;
    font-size: 1em;
    margin: 0 auto 0.1em auto;
    user-select: none;
  }

  .active {
    background-color: #ff3e00;
    color: white;
  }

  label:hover {
    opacity: 1;
    transform: scale(1.2, 1.2);
    color: white;
  }

  img {
    width: 30px;
  }
</style>

<div class="board">

  <div class="top">
    <div class="active-pokeball-name">{name}</div>
    <div class="top-grid">
      {#each pokeballs.filter(t => !t.done) as pokeball (pokeball.id)}
        <label
          in:receive={{ key: pokeball.id }}
          out:send={{ key: pokeball.id }}
          animate:flip>
          <div on:mouseenter={enter(pokeball.name)} on:mouseleave={leave}>
            <div let:hovering={active}>
              <div class:active>
                {#if active}
                  <div class="active">
                    <img src={pokeball.icon} alt={pokeball.name} />
                    <input
                      style="width: 0%;"
                      type="checkbox"
                      bind:checked={pokeball.done} />
                  </div>
                {:else}
                  <div class="hidden">
                    <img src={pokeball.icon} alt={pokeball.name} />
                    <input
                      style="width: 0%; position: absolute;"
                      type="checkbox"
                      bind:checked={pokeball.done} />
                  </div>
                {/if}
              </div>
            </div>
          </div>
        </label>
      {/each}
    </div>
  </div>

  <div class="bottom">
    <div class="bottom-grid">
      <div class="rg-img-item">

        {#each pokeballs.filter(t => t.done) as pokeball (pokeball.id)}
          <label
            in:receive={{ key: pokeball.id }}
            out:send={{ key: pokeball.id }}
            animate:flip>
            <img style="width: 70px;" src={pokeball.icon} alt={pokeball.name} />
            <input
              style="width: 0%;"
              type="checkbox"
              bind:checked={pokeball.done} />
          </label>
        {/each}
      </div>
      {#each pokeballs.filter(t => t.done) as pokeball (pokeball.id)}
        <div class="rg-data-container">
          <div class="title">{pokeball.name}</div>
          <div class="meta-data">{pokeball.effect}</div>
        </div>
        <div class="rg-location-container">
          <div class="sub-header">
            Location - {pokeball.BuyLocation.location}
          </div>
          <div class="meta-data">
            {pokeball.BuyLocation.name} - {pokeball.BuyLocation.description}
          </div>
        </div>
        <div class="rg-cost-container">
          <div class="sub-header">Cost</div>
          <div class="meta-data">{pokeball.BuyLocation.cost}</div>
        </div>
        <div class="rg-npc-container">
          <div class="sub-header">NPC - {pokeball.NPCLocation.location}</div>
          <div class="meta-data">{pokeball.NPCLocation.name}</div>
        </div>
        <div class="rg-npc-desc-container">
          <div class="sub-header">Description</div>
          <div class="meta-data">{pokeball.NPCLocation.description}</div>
        </div>
      {/each}
    </div>
  </div>
</div>

<!-- </div> -->
