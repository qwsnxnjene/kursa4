<template>
  <div class='main-content'>
    <div class="profile-text">
      <p>Профйль</p>
      <img src="../assets/profile.svg" alt="Аватар профиля">
      <div class="input-name">
        <p>ймя</p>
        <input @input="inputChanged" type="text" class="name-window" v-model="form.name">
        <!-- <img src="../assets/ok_btn.svg" alt=""> -->
        
      </div>
      <div class="input-surname">
          <p>фамйлйя</p>
          <input @input="inputChanged" type="text" class="surname-window" v-model="form.surname">
          <!-- <img src="../assets/ok_btn.svg" alt=""> -->
        </div>
      <div class="save-btn">
          <button :disabled="!isChanged" @click="saveChanges">
            Сохранйть
          </button>
      </div>
      <div class="agree-btn">
          <button @click="showTestMessage">
            Проитй тест
          </button>
      </div>
    </div>
    <transition name="popup-fade">
      <div v-if="showPopup" class="popup-success">
        Измененйя успешно сохранены
      </div>
    </transition>
    <transition name="popup-fade">
      <div v-if="showTestPopup" class="popup-success">
        Тест пока в разработке...
      </div>
    </transition>
  </div>
</template>

<script>
import CalendarSwipe from './CalendarSwipe.vue';
import axios from 'axios';

export default {
  name: 'MainContent',
  components: [CalendarSwipe],
  data() {
    return {
      isChanged: false,
      showPopup: false,
      showTestPopup: false,
      form: {
        name: "",
        surname: ""
      }
    }
  },
  methods: {
    inputChanged() {
      if (this.form.name == "" || this.form.surname == "") {
        this.isChanged = false
      } else {
        this.isChanged = true
      }
    },
    async saveChanges() {
      this.$store.commit('setUserName', this.form.name);
      this.$store.commit('setUserSurname', this.form.surname);
      try {
        const response = await axios.post('/api/save', {
          name: this.form.name,
          surname: this.form.surname,
          email: this.$store.state.userEmail
        })
        console.log(response.data);
      } catch (error) {
        console.error(error);
      }
      this.showPopup = true;
      setTimeout(() => {
        this.showPopup = false;
      }, 1000);
      this.isChanged = false;
    },
    showTestMessage() {
      this.showTestPopup = true;
      setTimeout(() => {
        this.showTestPopup = false;
      }, 1000);
    }
  },
  computed: {
    userName() {
      return this.$store.state.userName;
    },
    userSurname() {
      return this.$store.state.userSurname;
    }
  },
  created() {
    this.form.name = this.userName;
    this.form.surname = this.userSurname;
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
  display: flex;
  justify-content: center;
  width: auto;
  
  background-color: #1a1a1a;
  padding-top: 106px;
  height: 100%;
  padding-bottom: 18%;
}

.profile-text p {
  color: #fff;
  font-family: 'IF Kica';
  font-size: 60px;
  margin: 20px;
}

.profile-text img {
  height: 350px;
}

.input-name p {
  font-size: 45px;
  text-align: left;
  margin-bottom: 5px;
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

/* .input-name img {
  height: 60px;
  margin-left: 20px;
} */

.input-surname p {
  font-size: 45px;
  text-align: left;
  margin-bottom: 5px;
}

.input-surname input {
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

/* .input-surname img {
  height: 60px;
  margin-left: 20px;
} */

.agree-btn button{
  font-size: 40px;
  color: #fff;
  font-family: 'IF Kica';
  background-color: #6A1B9A;
  border: none;
  border-radius: 16px;
  width: 410px;
  margin-right: -10px;
  cursor: pointer;
}

.agree-btn {
  margin-top: 40px;
}

.save-btn button{
  font-size: 40px;
  color: #fff;
  font-family: 'IF Kica';
  background-color: #2e2e2e;
  border: 1px solid;
  border-color: #6A1B9A;
  border-radius: 16px;
  width: 410px;
  margin-right: -10px;
  cursor: pointer;
}

.save-btn {
  margin-top: 40px;
}

.save-btn button:disabled {
  background-color: #535353;
  color: #b1b1b1;
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