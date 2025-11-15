<script lang="ts">
  import {Greet, ListRifles} from '../wailsjs/go/main/App.js'
  import type {ports} from '../wailsjs/go/models';

  let resultText: string = "Please enter your name below 👇"
  let name: string
  let rifles: ports.Rifle[];

  function greet(): void {
    Greet(name).then(result => resultText = result)
  }

  function listRifles(): void {
    ListRifles().then(result => rifles = result)
  }
</script>

<main>
  <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={name} class="input" id="name" type="text"/>
    <button class="btn" on:click={greet}>Submit</button>
  </div>
  <div class="input-box">
    <button class="btn" on:click={listRifles}>List Rifles</button>
  </div>

  {#if rifles}
    <table>
      <thead>
      <tr>
        <th>Name</th>
        <th>Sight Height</th>
        <th>Barrel Twist</th>
        <th>Zero Range</th>
      </tr>
      </thead>
      <tbody>
      {#each rifles as rifle}
        <tr>
          <td>{rifle.name}</td>
          <td>{rifle.sight_height}</td>
          <td>{rifle.barrel_twist}</td>
          <td>{rifle.zero_range}</td>
        </tr>
      {/each}
      </tbody>
    </table>
  {/if}
</main>

<style>

    table {
        width: 100%;
        border-collapse: collapse;
        margin-top: 20px;
    }

  th, td {
    border: 1px solid #ddd;
    padding: 8px;
  }

  th {
    background-color: #f2f2f2;
  }

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 100px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>
