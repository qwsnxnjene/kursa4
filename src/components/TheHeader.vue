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
    <div class="avatar" v-if="isAuthenticated">
      <router-link to="/profile">
        <img src="../assets/profile.svg" alt="Аватар">
      </router-link>
    </div>
    <router-link v-if="isAuthenticated && $store.state.userRole === 'admin'" to="/add-university" class="register-btn">Добавйть ВУЗ</router-link>
    <button class="login-btn" v-if="!isAuthenticated" @click="showLogin = true">Вход</button>
    <button class="register-btn" @click="showRegister = true" v-if="!isAuthenticated">Регйстрацйя</button>
    <button class="register-btn" v-else @click="logout">Выитй</button>
    <transition name="popup-fade">
      <div v-if="showRegister" class="register-modal">
        <div class="register-content">
          <h2>Регйстрацйя</h2>
          <form @submit.prevent="register">
            <label>
              Имя:
              <input v-model="registerForm.name" required>
            </label>
            <label>
              Email:
              <input v-model="registerForm.email" type="email" required>
            </label>
            <label>
              Пароль:
              <input v-model="registerForm.password" type="password" required>
            </label>
            <label>
              Роль:
              <select v-model="registerForm.role" required>
                <option value="student">Студент</option>
                <option value="admin">Представитель университета</option>
              </select>
            </label>
            <div class="register-actions">
              <button type="submit">Зарегйстрйроваться</button>
              <button type="button" @click="showRegister = false">Отмена</button>
            </div>
          </form>
        </div>
      </div>
    </transition>
    <transition name="popup-fade">
      <div v-if="registerMessage" class="popup-success">
        {{ registerMessage }}
      </div>
    </transition>
    <transition name="popup-fade">
      <div v-if="showLogin" class="register-modal">
        <div class="register-content">
          <h2>Вход</h2>
          <form @submit.prevent="login">
            <label>
              Email:
              <input v-model="loginForm.email" type="email" required>
            </label>
            <label>
              Пароль:
              <input v-model="loginForm.password" type="password" required>
            </label>
            <div class="register-actions">
              <button type="submit">Воитй</button>
              <button type="button" @click="showLogin = false">Отмена</button>
            </div>
          </form>
        </div>
      </div>
    </transition>
  </header>
</template>

<script>
import axios from 'axios';

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: 'Header',
  data() {
    return {
      showRegister: false,
      showLogin: false,
      registerForm: {
        name: '',
        email: '',
        password: '',
        role: 'student'
      },
      loginForm: {
        email: '',
        password: ''
      },
      registerMessage: ''
    }
  },
  computed: {
    selectedCity: {
      get() {
        return this.$store.state.selectedCity
      },
      set(value) {
        this.$store.commit('setCity', value)
      }
    },
    isAuthenticated() {
      return this.$store.state.isAuthenticated
    } 
  },
  methods: {
    updateCity() {
      this.$emit('city-changed', this.selectedCity)
    },
    async register() {
      
      try {
      const response = await axios.post('/api/signin', {
        name: this.registerForm.name,
        email: this.registerForm.email,
        password: this.registerForm.password,
        role: this.registerForm.role
      });
      console.log(response.data);
      this.showRegister = false;
      this.registerForm = { name: '', email: '', password: '', role: 'student' };
      this.registerMessage = response.data.message || 'Регйстрацйя успешна!';
      this.$store.commit('setAuthentication', true);
      this.$store.commit("setAuthToken", response.data.token);
      this.$store.commit("setUserName", response.data.name);
      this.$store.commit("setUserRole", response.data.role);
      this.$store.commit("setUserEmail", response.data.email);
    } catch (error) {
      this.registerMessage = error.response?.data?.error || error.response?.data?.message || 'Ошйбка регйстрацйй';
    }
    setTimeout(() => {
        this.registerMessage = '';
      }, 3000);
    },
    logout() {
      this.$store.commit('setAuthentication', false);
      this.$store.commit("setAuthToken", '');
      this.$store.commit("setUserName", "");
      this.$store.commit("setUserRole", '');
      this.$store.commit("setUserEmail", '');
      this.$store.commit("setUserSurname", '');
      this.$store.commit("setCity", 'kazan');
      if (this.$route.path !== '/') {
        this.$router.push('/');
      }
    },
    async login() {
      try {
        const response = await axios.post('/api/login', {
          email: this.loginForm.email,
          password: this.loginForm.password
        });
        console.log(response.data);
        this.showLogin = false;
        this.$store.commit("setUserEmail", response.data.email)
        this.loginForm = { email: '', password: '' };
        this.$store.commit('setAuthentication', true);
        this.$store.commit("setAuthToken", response.data.token);
        this.$store.commit("setUserName", response.data.name);
        this.$store.commit("setUserRole", response.data.role);
        this.$store.commit("setUserSurname", response.data.surname);
      } catch (error) {
        console.error(error);
        this.registerMessage = error.response?.data?.error || error.response?.data?.message || 'Ошйбка входа';
      }
      setTimeout(() => {
        this.registerMessage = '';
      }, 3000);
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
  width: 550px;
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
  margin-left: 120px;
  font-family: "IF Kica", Helvetica;
  font-weight: 400;
  color: #ffffff;
  font-size: 40px;
  letter-spacing: 0;
  white-space: nowrap;
}
.city-selector {
  margin-left: 60%;
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

.register-btn {
  margin-left: 15px;
  margin-right: 20px;
  font-size: 22px;
  background: #6A1B9A;
  color: #fff;
  border: none;
  border-radius: 12px;
  padding: 10px 24px;
  cursor: pointer;
  font-family: 'IF Kica';
}

.register-modal {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(0,0,0,0.5);
  z-index: 9999;
  display: flex;
  align-items: center;
  justify-content: center;
}

.register-content {
  background: #2e2e2e;
  padding: 40px 30px;
  border-radius: 20px;
  min-width: 350px;
  color: #fff;
  font-family: 'LC Web';
  box-shadow: 0 4px 24px #6A1B9A99;
}

.register-content h2 {
  font-family: 'IF Kica';
  font-size: 32px;
  margin-bottom: 20px;
}

.register-content label {
  display: block;
  margin-bottom: 16px;
  font-size: 22px;
}

.register-content input,
.register-content select {
  width: 90%;
  padding: 8px 10px;
  border-radius: 8px;
  border: none;
  margin-top: 4px;
  font-size: 18px;
  font-family: 'LC Web';
}

.register-actions {
  display: flex;
  justify-content: space-between;
  margin-top: 24px;
}

.register-actions button {
  font-size: 18px;
  border-radius: 8px;
  border: none;
  padding: 8px 18px;
  font-family: 'IF Kica';
  cursor: pointer;
}

.register-actions button[type="submit"] {
  background: #6A1B9A;
  color: #fff;
}

.register-actions button[type="button"] {
  background: #444;
  color: #fff;
}

.popup-success {
  position: fixed;
  top: 40px;
  left: 50%;
  transform: translateX(-50%);
  background: #6A1B9A;
  color: #fff;
  font-family: 'IF Kica';
  font-size: 28px;
  padding: 18px 40px;
  border-radius: 20px;
  box-shadow: 0 4px 24px #6A1B9A99;
  z-index: 10000;
  animation: fadeInOut 5s;
}

.login-btn {
  margin-left: 0;
  margin-right: 10px;
  font-size: 22px;
  background: #fff;
  color: #6A1B9A;
  border: 2px solid #6A1B9A;
  border-radius: 12px;
  padding: 10px 24px;
  cursor: pointer;
  font-family: 'IF Kica';
  transition: background 0.2s, color 0.2s;
}
.login-btn:hover {
  background: #6A1B9A;
  color: #fff;
}

@keyframes fadeInOut {
  0% { opacity: 0; }
  10% { opacity: 1; }
  90% { opacity: 1; }
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