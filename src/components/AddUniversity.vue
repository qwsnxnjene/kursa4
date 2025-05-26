<template>
  <div class="add-university">
    <h2>Добавйть ВУЗ</h2>
    <form @submit.prevent="submitForm">
      <label>
        Полное название:
        <input v-model="form.full_name" required />
      </label>
      <label>
        Короткое название:
        <input v-model="form.short_name" required />
      </label>
      <label>
        Город:
        <select v-model="form.city" required>
          <option value="kazan">Казань</option>
          <option value="moscow">Москва</option>
          <option value="petersburg">Санкт-Петербург</option>
        </select>
      </label>
      <label>
        Первый абзац:
        <textarea v-model="form.article1" required></textarea>
      </label>
      <label>
        Второй абзац:
        <textarea v-model="form.article2" required></textarea>
      </label>
      <label>
        Изображение 1 (URL):
        <input v-model="form.img1" required />
      </label>
      <label>
        Изображение 2 (URL):
        <input v-model="form.img2" required />
      </label>
      <label>
        Ссылка на сайт:
        <input v-model="form.url" type="url" required />
      </label>
      <button type="submit" @click="submitForm">Добавйть</button>
    </form>
    <div v-if="message" class="message">{{ message }}</div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      form: {
        full_name: '',
        short_name: '',
        city: 'kazan',
        article1: '',
        article2: '',
        img1: '',
        img2: '',
        url: ''
      },
      message: ''
    }
  },
  computed: {
    userRole() {
      return this.$store.state.userRole
    }
  },
  created() {
    // Если не представитель вуза — редирект на главную
    if (this.userRole !== 'admin') {
      this.$router.replace('/');
    }
  },
  methods: {
    async submitForm() {
      try {
        const response = await axios.post('/api/adduniversity', this.form);
        console.log(response.data)
        this.message = 'ВУЗ успешно добавлен!';
        this.form = {
          full_name: '',
          short_name: '',
          city: 'kazan',
          article1: '',
          article2: '',
          img1: '',
          img2: '',
          url: ''
        };
      } catch (error) {
        this.message = error.response?.data?.error | 'Ошибка при добавлении ВУЗа';
      }
      setTimeout(() => this.message = '', 4000);
    }
  }
}
</script>

<style scoped>
.add-university {
  max-width: 600px;
  margin: 40px auto;
  background: #2e2e2e;
  color: #fff;
  border-radius: 20px;
  padding: 32px 40px;
  font-family: 'LC Web';
  box-shadow: 0 4px 24px #6A1B9A99;
}
.add-university h2 {
  font-family: 'IF Kica';
  font-size: 32px;
  margin-bottom: 24px;
}
.add-university label {
  display: block;
  margin-bottom: 18px;
  font-size: 20px;
}
.add-university input,
.add-university textarea,
.add-university select {
  width: 100%;
  padding: 8px 10px;
  border-radius: 8px;
  border: none;
  margin-top: 4px;
  font-size: 18px;
  font-family: 'LC Web';
  background: #444;
  color: #fff;
}
.add-university button {
  margin-top: 18px;
  font-size: 20px;
  border-radius: 8px;
  border: none;
  padding: 10px 28px;
  font-family: 'IF Kica';
  background: #6A1B9A;
  color: #fff;
  cursor: pointer;
}
.message {
  margin-top: 20px;
  color: #fff;
  font-size: 22px;
}
</style>