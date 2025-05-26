<template>
  <div class="test-content">
    <h2>Профорйентацйонныи тест</h2>
    <div v-if="current < questions.length">
      <div class="question">
        <p>{{ current + 1 }}. {{ questions[current].text }}</p>
        <div class="answers">
          <button
            v-for="(answer, idx) in questions[current].answers"
            :key="idx"
            @click="choose(answer.type)"
          >
            {{ answer.text }}
          </button>
        </div>
      </div>
    </div>
    <div v-else class="result">
      <h3>Вам могут подойти направления:</h3>
      <ul>
        <li v-for="(dir, idx) in resultDirections" :key="idx">{{ dir }}</li>
      </ul>
      <button @click="restart">Проитй ещё раз</button>
    </div>
  </div>
</template>

<script>
const directions = {
  tech: 'Технические специальности (ИТ, инженерия, математика)',
  hum: 'Гуманитарные специальности (филология, история, журналистика)',
  med: 'Медицинские специальности (медицина, фармация, биология)',
  econ: 'Экономические специальности (экономика, менеджмент, финансы)',
  art: 'Творческие специальности (дизайн, искусство, архитектура)'
};

export default {
  name: 'StudentTest',
  data() {
    return {
      current: 0,
      scores: { tech: 0, hum: 0, med: 0, econ: 0, art: 0 },
      questions: [
        {
          text: 'Что вам больше нравится?',
          answers: [
            { text: 'Решать задачи, программировать', type: 'tech' },
            { text: 'Читать книги, писать тексты', type: 'hum' },
            { text: 'Изучать биологию, помогать людям', type: 'med' },
            { text: 'Планировать, считать деньги', type: 'econ' },
            { text: 'Рисовать, создавать что-то новое', type: 'art' }
          ]
        },
        {
          text: 'Какой предмет вам больше всего нравился в школе?',
          answers: [
            { text: 'Математика/Информатика', type: 'tech' },
            { text: 'Литература/История', type: 'hum' },
            { text: 'Биология/Химия', type: 'med' },
            { text: 'Обществознание/Экономика', type: 'econ' },
            { text: 'ИЗО/Музыка', type: 'art' }
          ]
        },
        {
          text: 'Что вы предпочли бы делать в свободное время?',
          answers: [
            { text: 'Собирать конструкторы, разбираться в технике', type: 'tech' },
            { text: 'Писать рассказы, вести блог', type: 'hum' },
            { text: 'Участвовать в волонтёрских акциях', type: 'med' },
            { text: 'Играть в экономические игры', type: 'econ' },
            { text: 'Рисовать, заниматься творчеством', type: 'art' }
          ]
        },
        {
          text: 'Что вам ближе?',
          answers: [
            { text: 'Логика и точные науки', type: 'tech' },
            { text: 'Общение с людьми, языки', type: 'hum' },
            { text: 'Забота о здоровье', type: 'med' },
            { text: 'Анализировать и планировать', type: 'econ' },
            { text: 'Самовыражение через творчество', type: 'art' }
          ]
        },
        {
          text: 'Какую профессию вы бы выбрали?',
          answers: [
            { text: 'Инженер/Программист', type: 'tech' },
            { text: 'Журналист/Учитель', type: 'hum' },
            { text: 'Врач/Фармацевт', type: 'med' },
            { text: 'Экономист/Менеджер', type: 'econ' },
            { text: 'Дизайнер/Архитектор', type: 'art' }
          ]
        },
        {
          text: 'Что вам интереснее?',
          answers: [
            { text: 'Технологии будущего', type: 'tech' },
            { text: 'Культура и общество', type: 'hum' },
            { text: 'Наука о жизни', type: 'med' },
            { text: 'Бизнес и финансы', type: 'econ' },
            { text: 'Искусство и дизайн', type: 'art' }
          ]
        },
        {
          text: 'Как вы относитесь к работе с людьми?',
          answers: [
            { text: 'Лучше с техникой', type: 'tech' },
            { text: 'Люблю общаться', type: 'hum' },
            { text: 'Хочу помогать людям', type: 'med' },
            { text: 'Интересно управлять коллективом', type: 'econ' },
            { text: 'Вдохновляет работать в команде творцов', type: 'art' }
          ]
        },
        {
          text: 'Что для вас важнее в работе?',
          answers: [
            { text: 'Решать сложные задачи', type: 'tech' },
            { text: 'Вдохновлять и обучать', type: 'hum' },
            { text: 'Приносить пользу обществу', type: 'med' },
            { text: 'Достигать целей и успеха', type: 'econ' },
            { text: 'Создавать что-то уникальное', type: 'art' }
          ]
        },
        {
          text: 'Какой формат работы вам ближе?',
          answers: [
            { text: 'Офис, лаборатория', type: 'tech' },
            { text: 'Редакция, школа', type: 'hum' },
            { text: 'Больница, лаборатория', type: 'med' },
            { text: 'Офис, бизнес-центр', type: 'econ' },
            { text: 'Студия, мастерская', type: 'art' }
          ]
        },
        {
          text: 'Что вы цените в профессии?',
          answers: [
            { text: 'Инновации и развитие', type: 'tech' },
            { text: 'Влияние на общество', type: 'hum' },
            { text: 'Возможность помогать', type: 'med' },
            { text: 'Стабильность и доход', type: 'econ' },
            { text: 'Свободу самовыражения', type: 'art' }
          ]
        },
        {
          text: 'Что вы чаще всего ищете в интернете?',
          answers: [
            { text: 'Новости технологий', type: 'tech' },
            { text: 'Статьи, книги, блоги', type: 'hum' },
            { text: 'Медицинские факты', type: 'med' },
            { text: 'Финансовые новости', type: 'econ' },
            { text: 'Идеи для творчества', type: 'art' }
          ]
        },
        {
          text: 'Какой проект вам интереснее всего?',
          answers: [
            { text: 'Разработка приложения', type: 'tech' },
            { text: 'Организация мероприятия', type: 'hum' },
            { text: 'Социальная акция', type: 'med' },
            { text: 'Стартап или бизнес', type: 'econ' },
            { text: 'Создание арт-объекта', type: 'art' }
          ]
        }
      ]
    }
  },
  computed: {
    resultDirections() {
      // Сортируем направления по количеству баллов и берём топ-3
      const sorted = Object.entries(this.scores)
        .sort((a, b) => b[1] - a[1])
        .slice(0, 3)
        // eslint-disable-next-line no-unused-vars
        .filter(([_, v]) => v > 0)
        .map(([k]) => directions[k]);
      return sorted.length ? sorted : ['Пожалуйста, ответьте на вопросы теста.'];
    }
  },
  methods: {
    choose(type) {
      this.scores[type]++;
      this.current++;
    },
    restart() {
      this.current = 0;
      this.scores = { tech: 0, hum: 0, med: 0, econ: 0, art: 0 };
    }
  }
}
</script>

<style scoped>
.test-content {
  max-width: 700px;
  margin: 40px auto;
  background: #2e2e2e;
  color: #fff;
  border-radius: 20px;
  padding: 32px 40px;
  font-family: 'LC Web';
  box-shadow: 0 4px 24px #6A1B9A99;
}
.test-content h2 {
  font-family: 'IF Kica';
  font-size: 32px;
  margin-bottom: 24px;
}
.question p {
  font-size: 24px;
  margin-bottom: 18px;
}
.answers {
  display: flex;
  flex-direction: column;
  gap: 14px;
}
.answers button {
  font-size: 20px;
  border-radius: 8px;
  border: none;
  padding: 10px 18px;
  font-family: 'LC Web';
  background: #6A1B9A;
  color: #fff;
  cursor: pointer;
  transition: background 0.2s;
}
.answers button:hover {
  background: #8e24aa;
}
.result h3 {
  font-size: 26px;
  margin-bottom: 18px;
}
.result ul {
  font-size: 22px;
  margin-bottom: 18px;
}
.result button {
  font-size: 18px;
  border-radius: 8px;
  border: none;
  padding: 8px 18px;
  font-family: 'IF Kica';
  background: #6A1B9A;
  color: #fff;
  cursor: pointer;
}
</style>