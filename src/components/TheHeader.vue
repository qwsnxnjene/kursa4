<template>
  <header class="header">
    <div class="logo">
      <router-link to="/">
        <img src="../assets/logo.svg" alt="Логотип">
      </router-link>
      <p class="logo_text">Цйфровои помощнйк</p>
    </div>
      <div class="city-selector">
        <select v-model="selectedCity" class="custom-select" @change="updateCity">
          <option value="kazan">Казань</option>
          <option value="moscow">Москва</option>
          <option value="petersburg">Санкт-Петербург</option>
        </select>
      </div>
    <div class="avatar">
      <router-link to="/profile">
        <img src="../assets/profile.svg" alt="Аватар">
      </router-link>
    </div>
  </header>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: 'Header',
  // data() {
  //   return {
  //     selectedCity: "kazan"
  //   }
  // },
  computed: {
    selectedCity: {
      get() {
        return this.$store.state.selectedCity
      },
      set(value) {
        this.$store.commit('setCity', value)
      }
    }
  },
  methods: {
    updateCity() {
      this.$emit('city-changed', this.selectedCity)
    }
  }
};
</script>

<style scoped>
@font-face {
	font-family: 'IF Kica'; 
	src: url(../assets/fonts/IFKica-Regular.ttf); 
}

@font-face {
  font-family: 'LC Web';
  src: url(../assets/fonts/Platform-Regular-Web.otf);
}

.header {
  display: flex;
  align-items: center;
  padding: 0px 10px;
  position: fixed;
  height: 106px;
  width: 100%;
  top: 0px;
  left: 0px;
  background-color: #2e2e2e;
  box-shadow: 0 2px 4px #6A1B9A;
}

.logo {
  position: relative;
  width: 839px;
  height: 100px;
}

.logo img {
  position: absolute;
  width: 100px;
  height: 100px;
  top: 0;
  left: 0;
}

.logo .logo_text {
  position: absolute;
  top: -15px;
  left: 80px;
  font-family: "IF Kica", Helvetica;
  font-weight: 400;
  color: #ffffff;
  font-size: 40px;
  letter-spacing: 0;
  width: 150%;
}
.city-selector {
  margin-left: 50%;
  width: fit-content;
}

.city-selector .custom-select {
  cursor: pointer;
  width: fit-content;
  appearance: none;       
  border: none;
  text-align: center;
  outline: none;
  background: transparent;
  -webkit-appearance: none; 
  -moz-appearance: none;   
  height: 100px;
  font: 35px 'LC Web';
  color: #ffffff;
  background-color: #2e2e2e;
}

.avatar {
  margin-left: -10px;
  padding: 0px 15px;
}
</style>