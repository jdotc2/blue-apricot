<style>.content {
  width: 100%;
  height: 100%;
  color: black;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: large;
}
:global(.svlt-grid-shadow) {
  background: pink;
}
:global(.svlt-grid-container) {
  background: #eee;
}
</style>

<script>
import Grid from "svelte-grid";
import gridHelp from "svelte-grid/build/helper";
import map from "lodash.map";

const id = () =>
  "_" +
  Math.random()
    .toString(36)
    .substr(2, 9);

function generateLayout(col) {
  return map(new Array(20), function(item, i) {
    const y = Math.ceil(Math.random() * 4) + 1;
    return gridHelp.item({
      x: (i * 2) % col,
      y: Math.floor(i / 6) * y,
      w: 2,
      h: y,
      id: id()
    });
  });
}

const randomNumberInRange = (min, max) => Math.random() * (max - min) + min;

let adjustAfterRemove = false;

let cols = 10;
// Just generate messy layout
let layout = generateLayout(cols);
// Helper function which normalize. you need to pass items and columns
let items = gridHelp.resizeItems(layout, cols);
// Apply breakpoints
let breakpoints = [[1100, 5], [800, 4], [530, 1]];
</script>

<h1>Responsive</h1>

<Grid {breakpoints} {items} bind:items={items} {cols} let:item={item} rowHeight={100} gap={10}>
	<div class=content style="background: #ccc; border: 1px solid black;">
    {item.id}
  </div>
</Grid>