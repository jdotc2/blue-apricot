<script>
  import { onMount, createEventDispatcher } from "svelte";
  import { scale } from "svelte/transition";
  import { quintOut } from "svelte/easing";
  import { crossfade } from "svelte/transition";
  import { flip } from "svelte/animate";

  let types = [];
  let superETypes = [];
  let noETypes = [];
  let notETypes = [];
  let typeNotes = "";
  let typeName = "Types";
  let hovering;
  let visible = false;
  let notesVisible = false;
  let activeType = "";
  let checked = false;

  $: type = typeName;
  $: activeT = activeType;
  $: superEffective = superETypes;
  $: notEffective = notETypes;
  $: noEffect = noETypes;
  $: notes = typeNotes;

  onMount(async () => {
    const res = await fetch("http://localhost:4747/types");
    let typeRes = await res.json();
    types = await typeRes.data;
    console.log(types);
  });

  function enter(name, event) {
    hovering = true;
    typeName = name;
  }

  function leave() {
    hovering = false;
    typeName = "Types";
  }

  function refresh() {
    superETypes = [];
    noETypes = [];
    notETypes = [];
  }

  function handleEffectData(type, event) {
    refresh();
    typeNotes = "";
    activeType = type.name;
    if (type.veryEffective !== null) {
      type.veryEffective.forEach(t => superETypes.push(t));
      superEffective = superETypes;
    } else {
      superETypes.push("Unknown");
    }

    if (type.notEffective !== null) {
      type.notEffective.forEach(t => notETypes.push(t));
      notETypes = type.notEffective;
    } else {
      notETypes.push("Unknown");
    }

    if (type.noEffect !== null) {
      type.noEffect.forEach(t => noETypes.push(t));
      noETypes = type.noEffect;
    } else {
      noETypes.push("Unknown");
    }

    if (type.notes !== null) {
      typeNotes = "";
      typeNotes = type.notes;
    }

    visible = true;
  }

  function typewriter(node, { speed = 30 }) {
    const valid =
      node.childNodes.length === 1 &&
      node.childNodes[0].nodeType === Node.TEXT_NODE;

    if (!valid) {
      throw new Error(
        `This transition only works on elements with a single text node child`
      );
    }

    const text = node.textContent;
    const duration = text.length * speed;

    return {
      duration,
      tick: t => {
        const i = ~~(text.length * t);
        node.textContent = text.slice(0, i);
      }
    };
  }
</script>

<style type="text/scss">
  .main-grid {
    display: grid;
    grid-template-columns: 100%;
    grid-gap: 10px;
    align-items: center;
  }

  .type-grid {
    display: grid;
    grid-template-columns: auto auto auto auto auto auto auto auto auto;
    grid-gap: 20px;
    padding: 10px;
    justify-content: center;
    align-items: center;
    margin-bottom: 5%;
  }

  .header {
    grid-column: 1 / span 9;
    font-family: "Raleway", sans-serif;
    font-size: 28px;
    border-bottom: 1px solid #6d7075;
    align-items: center;
    text-align: center;
    padding-bottom: 2%;
  }

  .icon {
    border-radius: 100%;
    height: 40px;
    width: 40px;
    margin: auto;
    transition: 200ms all;
  }

  .icon img {
    height: 60%;
    width: 60%;
    margin: 20%;
  }

  .icon:hover {
    filter: saturate(200%);
    transform: scale(1.1);
    cursor: pointer;
  }

  .icon-effect {
    border-radius: 100%;
    height: 35px;
    width: 35px;
    margin: auto;
    transition: 200ms all;
  }

  .icon-effect img {
    height: 60%;
    width: 60%;
    margin: 20%;
  }

  .Bug {
    background: #92bc2c;
    box-shadow: 0 0 20px #92bc2c;
  }

  .Dark {
    background: #595761;
    box-shadow: 0 0 20px #595761;
  }

  .Dragon {
    background: #0c69c8;
    box-shadow: 0 0 20px #0c69c8;
  }

  .Electric {
    background: #f2d94e;
    box-shadow: 0 0 20px #f2d94e;
  }

  .Fire {
    background: #fba54c;
    box-shadow: 0 0 20px #fba54c;
  }

  .Fairy {
    background: #ee90e6;
    box-shadow: 0 0 20px #ee90e6;
  }

  .Fighting {
    background: #d3425f;
    box-shadow: 0 0 20px #d3425f;
  }

  .Flying {
    background: #a1bbec;
    box-shadow: 0 0 20px #a1bbec;
  }

  .Ghost {
    background: #5f6dbc;
    box-shadow: 0 0 20px #5f6dbc;
  }

  .Grass {
    background: #5fbd58;
    box-shadow: 0 0 20px #5fbd58;
  }

  .Ground {
    background: #da7c4d;
    box-shadow: 0 0 20px #da7c4d;
  }

  .Ice {
    background: #75d0c1;
    box-shadow: 0 0 20px #75d0c1;
  }

  .Normal {
    background: #a0a29f;
    box-shadow: 0 0 20px #a0a29f;
  }

  .Poison {
    background: #b763cf;
    box-shadow: 0 0 20px #b763cf;
  }

  .Psychic {
    background: #fa8581;
    box-shadow: 0 0 20px #fa8581;
  }

  .Rock {
    background: #c9bb8a;
    box-shadow: 0 0 20px #c9bb8a;
  }

  .Steel {
    background: #5695a3;
    box-shadow: 0 0 20px #5695a3;
  }

  .Water {
    background: #539ddf;
    box-shadow: 0 0 20px #539ddf;
  }

  .active-type-grid {
    display: grid;
    grid-template: 25% 75%;
    grid-gap: 15px;
    padding-top: 10px;
  }

  .active-type-container {
    position: relative;
    grid-column: 1;
    grid-row: 1 / span 4;
    justify-content: center;
    align-items: center;
  }

  .icon-active {
    border-radius: 100%;
    height: 100px;
    width: 100px;
    transition: 200ms all;
    position: absolute;
    left: 50%;
    top: 50%;
    margin: 0;
    transform: translate(-50%, -50%);
  }

  .icon-active img {
    height: 60%;
    width: 60%;
    margin: 20%;
  }

  .active-type-title {
    font-family: "Raleway", sans-serif;
    font-size: 26px;
    margin-top: 5%;
  }

  .s-effect-header {
    font-family: "Saira Condensed", sans-serif;
    color: white;
    font-size: 16px;
    text-align: left;
    padding: 1%;
    position: relative;
    background: #10c269;
    &:focus {
      outline: 0;
    }
    &:before {
      content: "";
      display: block;
      position: absolute;
      background: rgba(255, 255, 255, 0.5);
      width: 60px;
      height: 100%;
      left: 0;
      top: 0;
      opacity: 0.5;
      filter: blur(30px);
      transform: translateX(-100px) skewX(-15deg);
    }
    &:after {
      content: "";
      display: block;
      position: absolute;
      background: rgba(255, 255, 255, 0.2);
      width: 30px;
      height: 100%;
      left: 50px;
      top: 0;
      opacity: 0;
      filter: blur(5px);
      transform: translateX(-100px) skewX(-15deg);
    }
    &:hover {
      background: #10c269;
      cursor: pointer;
      &:before {
        transform: translateX(300px) skewX(-15deg);
        opacity: 0.6;
        transition: 0.7s;
      }
      &:after {
        transform: translateX(300px) skewX(-15deg);
        opacity: 1;
        transition: 0.7s;
      }
    }
  }

  .not-effect-header {
    font-family: "Saira Condensed", sans-serif;
    color: white;
    font-size: 16px;
    text-align: left;
    padding: 1%;
    position: relative;
    background: #f44336;
    &:focus {
      outline: 0;
    }
    &:before {
      content: "";
      display: block;
      position: absolute;
      background: rgba(255, 255, 255, 0.5);
      width: 60px;
      height: 100%;
      left: 0;
      top: 0;
      opacity: 0.5;
      filter: blur(30px);
      transform: translateX(-100px) skewX(-15deg);
    }
    &:after {
      content: "";
      display: block;
      position: absolute;
      background: rgba(255, 255, 255, 0.2);
      width: 30px;
      height: 100%;
      left: 50px;
      top: 0;
      opacity: 0;
      filter: blur(5px);
      transform: translateX(-100px) skewX(-15deg);
    }
    &:hover {
      background: #f44336;
      cursor: pointer;
      &:before {
        transform: translateX(300px) skewX(-15deg);
        opacity: 0.6;
        transition: 0.7s;
      }
      &:after {
        transform: translateX(300px) skewX(-15deg);
        opacity: 1;
        transition: 0.7s;
      }
    }
  }

  .no-effect-header {
    font-family: "Saira Condensed", sans-serif;
    color: white;
    font-size: 16px;
    text-align: left;
    padding: 1%;
    background: #7d7979;
    position: relative;
    &:focus {
      outline: 0;
    }
    &:before {
      content: "";
      display: block;
      position: absolute;
      background: rgba(255, 255, 255, 0.5);
      width: 60px;
      height: 100%;
      left: 0;
      top: 0;
      opacity: 0.5;
      filter: blur(30px);
      transform: translateX(-100px) skewX(-15deg);
    }
    &:after {
      content: "";
      display: block;
      position: absolute;
      background: rgba(255, 255, 255, 0.2);
      width: 30px;
      height: 100%;
      left: 50px;
      top: 0;
      opacity: 0;
      filter: blur(5px);
      transform: translateX(-100px) skewX(-15deg);
    }
    &:hover {
      background: #7d7979;
      cursor: pointer;
      &:before {
        transform: translateX(300px) skewX(-15deg);
        opacity: 0.6;
        transition: 0.7s;
      }
      &:after {
        transform: translateX(300px) skewX(-15deg);
        opacity: 1;
        transition: 0.7s;
      }
    }
  }

  .effect-container {
    display: inline-grid;
    padding-top: 2%;
  }

  .super-e {
    grid-column: 2;
    grid-row: 1;
    border: 1px solid #10c269;
    border-radius: 8px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
    width: 100%;
  }

  .not-e {
    grid-column: 2;
    grid-row: 2;
    border: 1px solid #f44336;
    border-radius: 8px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
    width: 100%;
  }

  .no-e {
    grid-column: 2;
    grid-row: 3;
    border: 1px solid #7d7979;
    border-radius: 8px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
    width: 100%;
  }

  .notes {
    grid-column: 2;
    grid-row: 4;
    text-align: center;
    margin: auto;
  }

  .notes-container {
    width: 60%;
  }

  p {
    font-family: "Raleway", sans-serif;
    font-size: 13px;
    font-weight: bolder;
    text-align: center;
  }
</style>

<div class="main-grid">
  <div class="type-grid">
    <div class="header">{type}</div>
    {#each types as type}
      <div
        on:mouseenter={enter(type.name)}
        on:mouseleave={leave}
        on:click={handleEffectData(type)}>
        <div class="select-icon">
          <div class={`icon ${type.name}`}>
            <img src={`img/types/${type.name}.svg`} alt={type.name} />
          </div>
        </div>
      </div>
    {/each}
  </div>

  {#if visible}
    <div class="active-type-grid">
      <div class="active-type-container">
        <div class={`icon-active ${activeType}`}>
          <img src={`img/types/${activeType}.svg`} alt={activeType} />
          <div class="active-type-title">{activeType}</div>

        </div>
      </div>

      <div class="super-e">
        <div class="s-effect-header">Super Effective</div>
        {#each superEffective as superE}
          <div class="effect-container">
            <img src={`img/type/${superE}.png`} alt={superE} />
          </div>
        {/each}
      </div>

      <div class="not-e">
        <div class="not-effect-header">Not Effective</div>
        {#each notEffective as notE}
          <div class="effect-container">
            <img src={`img/type/${notE}.png`} alt={notE} />
          </div>
        {/each}
      </div>

      <div class="no-e">
        <div class="no-effect-header">No Effect</div>
        {#each noEffect as noE}
          <div class="effect-container">
            <img src={`img/type/${noE}.png`} alt={noE} />
          </div>
        {/each}
      </div>

    </div>
      
        <p>{notes}</p>

  {/if}

</div>
