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
          <button :disabled="!isChanged">
            Сохранйть
          </button>
      </div>
      <div class="agree-btn">
        <router-link to="/study">
          <button>
            Проитй тест
          </button>
        </router-link>
      </div>
    </div>
    
  </div>
</template>

<script>
import CalendarSwipe from './CalendarSwipe.vue';

export default {
  name: 'MainContent',
  components: [CalendarSwipe],
  data() {
    return {
      isChanged: false,
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
</style>