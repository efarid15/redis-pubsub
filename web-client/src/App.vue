<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">

    <hr>
      {{message}}
    <hr>

    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',
  data() {
    return {
      message : ""
    }
  },
  created() {
    try {

      const ws = new WebSocket("ws://localhost:3000");
      ws.onmessage = ({data}) => {
        this.message = data;
        console.log( data );
      }

    } catch (err) {
        console.log(err);
    }
  },
  components: {
    HelloWorld
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
