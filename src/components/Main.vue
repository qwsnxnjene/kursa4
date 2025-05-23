<template>
  <div class='main-content'>
    <div class='telegram-swipe'>
        <img class="cal" src="../assets/teleg.svg" />
    </div>
    <div class='search'>
      <p>Пойск<br>по ВУЗам</p>
      <input @input="performSearch" v-model="searchQuery" type="text" class="search-window">
      <div v-if="results.length > 0" class="results-container">
      <div 
        v-for="(item, index) in results" 
        :key="index"
        class="result-item"
      >
        <router-link :to="'/universities/' + item.id" class="router">
          {{ item.name }}
        </router-link>
      </div>
    </div>
    </div>
    
    <CalendarSwipe></CalendarSwipe>
  </div>
</template>

<script>
import axios from 'axios';
import CalendarSwipe from './CalendarSwipe.vue';

export default {
  name: 'MainContent',
  components: {CalendarSwipe},
  data() {
    return {
      searchQuery: "",
      results: [],
      debounceTimer: null,
      citySelected: "kazan"
    }
  },
  beforeDestroy() {
    clearInterval(this.updateInterval)
  },

  computed: {
    currentCity() {
      return this.$store.state.selectedCity
    }
  },

  methods: {
    async performSearch() {
      console.log("Отправка запроса с query:", this.searchQuery)
      
      if (this.searchQuery === "") {
        this.results = []
        return
      }

      try {
        const response = await axios.get('/api/search', {
          params: {
            query: this.searchQuery,
            city: this.$store.state.selectedCity
          }
        })
        console.log(response.data)
        this.results = this.sortResults(response.data, this.searchQuery)
        
      } catch (error) {
        console.error('Ошибка поиска', error)
      } 
    },

    sortResults(results, query) {
      const lowerQuery = query.toLowerCase();
      
      return results.sort((a, b) => {
        // Приоритет 1: начинается с запроса
        const aStartsWith = a.name.toLowerCase().startsWith(lowerQuery);
        const bStartsWith = b.name.toLowerCase().startsWith(lowerQuery);
        
        if (aStartsWith && !bStartsWith) return -1;
        if (!aStartsWith && bStartsWith) return 1;
        
        // Приоритет 2: содержит запрос
        const aContains = a.name.toLowerCase().includes(lowerQuery);
        const bContains = b.name.toLowerCase().includes(lowerQuery);
        
        if (aContains && !bContains) return -1;
        if (!aContains && bContains) return 1;
        
        // Если одинаковый приоритет - сортируем по алфавиту
        return a.name.localeCompare(b.name);
      }).slice(0, 3); // Берем только 3 результата
    },
  },

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
  justify-content: space-between;
  background-color: #1a1a1a;
  padding-top: 106px;
  height: 100%;
  padding-bottom: 17%;
}

.search {
  padding-top: 15%;
}

.search p {
  font-family: 'IF Kica';
  font-size: 60px;
  color: #ffffff;
  margin-bottom: 2%;
}

.search .search-window {
  border-radius: 30px;
  border: none;
  width: 500px;
  height: 65px;
  box-shadow: 0 4px 4px #6A1B9A;
  font-size: 30px;
  font-family: 'LC Web';
  box-sizing: border-box;
  padding-left: 10px;
}

.telegram-swipe {
  border-radius: 0 30px 30px 0;
  background-color: #2e2e2e;
  box-shadow: 0 4px 4px #000;
  padding: 5px;
  margin-top: 18%;
  height: fit-content;
}

.telegram-swipe img {
  margin: 100px 0;
  width: 80px;
  height: 80px;
}

.result-item {
   margin: 20px;
  padding: 8px;
  
}

.result-item .router {
  color: #fff;
  font-family: 'LC Web';
  padding: 8px;
  border: #6a1b9a solid;
  border-radius: 13px;
  text-decoration: none;
}
</style>