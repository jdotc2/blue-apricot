<script>
  import Navbar from "../Navbar.svelte";
  import Grid from "svelte-grid";
  import gridHelp from "svelte-grid/build/helper/index.mjs";

  const id = () => "_" + "navbar";
  let cols = 10;
  let adjustAfterRemove = false;
  let items = [gridHelp.item({ x: 0, y: 0, w: 5, h: 2, id: id(), component: Navbar })];
  let breakpoints = [[1100, 5], [800, 4], [530, 1]];
</script>

<Grid {breakpoints} bind:items={items} {cols} let:item rowHeight={100} gap={10}>
  <div class="content">
  </div>
</Grid>
