<script>
  import { LottiePlayer } from "@lottiefiles/svelte-lottie-player";

  let weatherSprite = undefined;
</script>

<style>
  .contact-card {
    display: grid;
    grid-template-columns: auto;
    grid-gap: 2px;
    font-family: "Raleway", sans-serif;
    width: 95px;
    /* border: 1px solid #aaa;
    border-radius: 2px; */
    box-shadow: 2px 2px 8px rgba(0, 0, 0, 0.1);

    justify-content: center;
    align-items: center;
  }

  /* .contact-card > div {
    text-align: center;
    padding: 20px 0;
  } */

  .sprite {
    padding: 15px 0;
    justify-content: center;
    align-items: center;
  }

  .header {
    font-family: "Raleway", sans-serif;
    font-size: 12px;
    border-bottom: 1px solid #ff3e00;
    padding: 10px 0;
  }

  .date {
    font-size: 13px;
    font-weight: bold;
    padding: 10px 0;
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

  <article class="contact-card">
    <div class="sprite">
      <slot {weatherSprite} />
    </div>

    <div class="header">
      <slot name="name">
        <span class="missing">Unknown name</span>
      </slot>
    </div>

    <div class="date">
      <slot name="date">
        <span class="missing">Unknown address</span>
      </slot>
    </div>

  </article>
{/if}
