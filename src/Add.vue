<template>
    <div id="add">
        <TheHeader />
        <div class="add-content">
            <TheMain />
        </div>
        <TheFooter />
    </div>
</template>

<script>
import TheHeader from './components/TheHeader.vue';
import TheMain from './components/AddUniversity.vue';
import TheFooter from './components/TheFooter.vue';

export default {
  name: 'theAdd',
  components: {
    TheHeader,
    TheMain,
    TheFooter
  },
};
</script>

<style scoped>
.add-content {
  padding-top: 120px;   
  padding-bottom: 60px; 
  min-height: 60vh;     
}
</style>