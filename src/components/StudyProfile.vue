<template>
  <div class='main-content'>
    <div class='vuz'>
        <a :href="uni.url">{{ uni.short_name }}</a>
    </div>
    <div class="first-article">
        
        <img :src="uni.img1" alt="">
        <p>{{ uni.article1 }}</p>
    </div>
    <div class="second-article">
        <p>{{ uni.article2 }}</p>
        <img :src="uni.img2" alt="">
    </div>
    <div class="info">
        <p>Информацйя о поступленйй</p>
    </div>
    <div class="info-small">
        <p>
            Заполните форму и наш менеджер проконсультирует вас абсолютно бесплатно! 
            Ответим на все вопросы, поможем определиться 
            с направлением и будем поддерживать вплоть до поступления!
        </p>
    </div>
    <div class="input-name">
        <p>ймя</p>
        <input @input="isChanged=true" type="text" class="name-window">
    </div>
    <div class="input-name">
        <p>номер телефона</p>
        <input @input="isChanged=true" type="text" class="name-window">
    </div>
    <div class="input-name">
        <p>email</p>
        <input @input="isChanged=true" type="text" class="name-window">
    </div>
    <div class="ready-btn">
      <button :disabled="!isChanged" @click="submitForm">
              Готово!
      </button>
    </div>
    <transition name="popup-fade">
      <div v-if="showPopup" class="popup-success">
        Заявка прйнята!
      </div>
    </transition>
  </div>
</template>

<script>
import axios from 'axios';
export default {

  name: 'MainContent',
  data() {
    return {
      isChanged: false,
      showPopup: false,
      uni: {
        name: '',
        short_name: '',
        img1: '',
        img2: '',
        article1: '',
        article2: '',
        url: ''
      }
    }
  },
  methods: {
    async fetchUniversity() {
      try {
        const response = await axios.get(`/api/universities/${this.$route.params.id}`, {
          params: {
            city: this.$store.state.selectedCity
          }
      })
        this.uni = response.data
        console.log(this.uni)
      } catch (error) {
        console.error('error')
      }
    },
    submitForm() {
      this.showPopup = true;
      setTimeout(() => {
        this.showPopup = false;
      }, 2000);
      this.isChanged = false;
    }
  },
  created() {
    this.fetchUniversity()
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

.main-content {
  display: flexbox;
  justify-content: center;
  width: auto;
  
  background-color: #1a1a1a;
  padding-top: 106px;
  height: 100%;
  padding-bottom: 18%;
}

.vuz a {
  color: #fff;
  font-family: 'IF Kica';
  font-size: 60px;
  margin: 20px;
  text-decoration: none;
}

.first-article {
    text-align: left;
    display: flex;
}

.first-article img {
    width: 600px;
    height: 400px;
    margin: 10px 600px 15px 300px;
}

.first-article p {
    margin-top: 50px;
    color: #fff;
    font-family: 'LC Web';
    font-size: 50px;
    width: 800px;
}

.second-article {
    text-align: right;
    display: flex;
}

.second-article img {
    width: 600px;
    height: 400px;
    margin: 10px 100px 15px 600px;
}

.second-article p {
    margin-left:100px;
    margin-top: 50px;
    color: #fff;
    font-family: 'LC Web';
    font-size: 50px;
    width: 800px;
}

.info p {
    color: #fff;
  font-family: 'IF Kica';
  font-size: 60px;
  margin: 20px;
  margin-top:100px;
}

.info-small p {
    font-family: 'LC Web';
    color: #fff;
    font-size: 20px;
}

.input-name {
  text-align: left;
  padding-left: 450px;
}

.input-name p {
  font-size: 45px;
  text-align: left;
  margin-bottom: 5px;
  color: #fff;
  font-family: 'IF Kica';
}

.input-name input {
  margin-right: -10px;
  border-radius: 16px;
  border: none;
  width: 410px;
  height: 60px;
  box-shadow: 0 4px 4px #6A1B9A;
  font-size: 30px;
  font-family: 'LC Web';
  box-sizing: border-box;
  padding-left: 10px;
}

.ready-btn {
  text-align: left;
  padding-left: 450px;
  margin-top:  60px;
}

button {
  font-size: 40px;
  color: #fff;
  font-family: 'IF Kica';
  background-color: #6A1B9A;
  border: none;
  border-radius: 16px;
  width: 410px;
  cursor: pointer;
}

button:disabled {
  background-color: #4A126B; /* Промежуточный тёмный оттенок */
  color: #C9A3E8;
  cursor: not-allowed;
  /* Можно добавить эффект "выцветания" */
  box-shadow: inset 0 0 10px rgba(0,0,0,0.2);
}

.popup-success {
  position: fixed;
  top: 40px;
  left: 50%;
  transform: translateX(-50%);
  background: #6A1B9A;
  color: #fff;
  font-family: 'IF Kica';
  font-size: 36px;
  padding: 24px 60px;
  border-radius: 20px;
  box-shadow: 0 4px 24px #6A1B9A99;
  z-index: 1000;
  animation: fadeInOut 5s;
}

@keyframes fadeInOut {
  0% { opacity: 0; }
  10% { opacity: 1; }
  80% { opacity: 0.9; }
  90% { opacity: 0.4; }
  100% { opacity: 0; }
}

.popup-fade-enter-active, .popup-fade-leave-active {
  transition: opacity 1s;
}
.popup-fade-enter-from, .popup-fade-leave-to {
  opacity: 0;
}
.popup-fade-enter-to, .popup-fade-leave-from {
  opacity: 1;
}

</style>