<template>
  <div id="profile">
    <TheHeader />
    <TheMain />
    <TheFooter />
  </div>
</template>

<script>
import TheHeader from './components/TheHeader.vue';
import TheMain from './components/MainProfile.vue';
import TheFooter from './components/TheFooter.vue';

export default {
  name: 'theProfile',
  components: {
    TheHeader,
    TheMain,
    TheFooter
  },
};
</script>
<style>
body {
  background-color: #1a1a1a;
  margin: 0;
  width: 100%;
  height: 100%;
}
#profile {
  width: 100%;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
