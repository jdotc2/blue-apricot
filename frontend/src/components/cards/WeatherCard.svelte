<script>
  let weatherSprite = undefined;
  let hovering;

  function enter() {
    hovering = true;
  }

  function leave() {
    hovering = false;
  }
</script>

<style>
  .contact-card {
    display: grid;
    grid-template-columns: auto;
    grid-gap: 2px;
    font-family: "Raleway", sans-serif;
    width: 110px;
    /* border: 1px solid #aaa; */
    border-radius: 25% 25%;
    box-shadow: 2px 2px 8px rgba(0, 0, 0, 0.5);
    justify-content: center;
    align-items: center;
    padding: 1px;
  }

  .contact-card > div {
    text-align: center;
    padding: 10px 0;
  }

  .header {
    font-family: "Raleway", sans-serif;
    font-size: 14px;
    border-bottom: 1px solid #6d7075;
    justify-content: center;
    align-items: center;
    display: grid;
    grid-template-columns: auto auto;
    grid-gap: 5px;

  }

  .date {
    font-size: 13px;
    font-weight: bold;
    padding: 10px 10px;
  }

  .missing {
    color: #999;
  }
</style>

{#if weatherSprite !== undefined}

  <article class="contact-card">
    <div class="sprite">
      <slot {weatherSprite} />
    </div>
  </article>
{:else}

  <div class="contact-card" on:mouseenter={enter} on:mouseleave={leave}>
    <div class="header">
      <div class="sprite" >
        <slot {weatherSprite} hovering={hovering}/>
      </div>
      
      <div class="sprite">
        <slot name="name">
          <span class="missing">Unknown name</span>
        </slot>
      </div>
    </div>

    <div class="date">
      <slot name="date">
        <span class="missing">Unknown address</span>
      </slot>
    </div>
  </div>
{/if}
