<script></script>

<div class="grid-container">
  <button class="pokedex-button">Pokedex</button>
  <button class="trade-button">Trade</button>
  <button class="map-button">Map</button>
  <button class="settings-button">Settings</button>
</div>

<style>
.grid-container {
  display: grid;
  grid-template-columns: 12% 12% 12% 12%;
  grid-gap: 10px;
  padding: 10px;
  justify-content: center;
  align-items: center;
  padding-bottom: 1%;
}

.grid-container > div {
  text-align: center;
  padding: 20px 0;
}

.pokedex-button {
  font-family: 'Raleway', sans-serif;
  background-color: white;
  border: 2px solid #f44336;
  color: #6d7075;
  padding: 10px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  cursor: pointer;
  -webkit-transition-duration: 0.4s; /* Safari */
  transition-duration: 0.4s;
  border-radius: 8px;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}

.pokedex-button:hover {
  background-color: #D82B2B; 
  color: white; 
}

.pokedex-button:active {
  box-shadow: 0 0px 0px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  transition-duration: 0.1s;
  transform: translateY(1px);
}

.trade-button {
  font-family: 'Raleway', sans-serif;
  background-color: white;
  border: 2px solid #4d90fa;
  color: #6d7075;
  padding: 10px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  cursor: pointer;
  -webkit-transition-duration: 0.4s; /* Safari */
  transition-duration: 0.4s;
  border-radius: 8px;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}

.trade-button:hover {
  background-color: #4d90fa; 
  color: white; 
}

.trade-button:active {
  box-shadow: 0 0px 0px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  transition-duration: 0.1s;
  transform: translateY(1px);
}

.map-button {
  font-family: 'Raleway', sans-serif;
  background-color: white;
  border: 2px solid #43AC5B;
  color: #6d7075;
  padding: 10px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  cursor: pointer;
  -webkit-transition-duration: 0.4s; /* Safari */
  transition-duration: 0.4s;
  border-radius: 8px;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}

.map-button:hover {
  background-color: #43AC5B; 
  color: white; 
}

.map-button:active {
  box-shadow: 0 0px 0px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  transition-duration: 0.1s;
  transform: translateY(1px);
}

.settings-button {
  font-family: 'Raleway', sans-serif;
  background-color: white;
  border: 2px solid #6d7075;
  color: #6d7075;
  padding: 10px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  cursor: pointer;
  -webkit-transition-duration: 0.4s; /* Safari */
  transition-duration: 0.4s;
  border-radius: 8px;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}

.settings-button:hover {
  background-color: #6d7075; 
  color: white; 
}

.settings-button:active {
  box-shadow: 0 0px 0px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  transition-duration: 0.1s;
  transform: translateY(1px);
}
</style>