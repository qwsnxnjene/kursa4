<template>
    <div id="add">
        <TheHeader />
        <div class="test-content">
            <TheMain />
        </div>
        <TheFooter />
    </div>
</template>

<script>
import TheHeader from './components/TheHeader.vue';
import TheMain from './components/StudentTest.vue';
import TheFooter from './components/TheFooter.vue';

export default {
  name: 'theTest',
  components: {
    TheHeader,
    TheMain,
    TheFooter
  },
};
</script>

<style scoped>
.test-content {
  padding-top: 120px;   
  padding-bottom: 60px; 
  min-height: 60vh;     
}
</style>