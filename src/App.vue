<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/profile">About</router-link>
      <router-link to="/study">Study</router-link>
    </div>
    <router-view/>
  </div>
</template>