<template>
  <div id="app">
    <TheHeader />
    <TheMain />
    <TheFooter />
  </div>
</template>

<script>
import TheHeader from './components/TheHeader.vue';
import TheMain from './components/Main.vue';
import TheFooter from './components/TheFooter.vue'

export default {
  name: 'theStart',
  components: {
    TheHeader,
    TheMain,
    TheFooter
  },
};
</script>
<style>
body {
  margin: 0;
  width: 100%;
  height: 100%;
}
#app {
  width: 100%;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
