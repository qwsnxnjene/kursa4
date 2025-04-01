<template>
  <div class='main-content'>
    <div class='telegram-swipe'>
        <img class="cal" src="../assets/teleg.svg" />
    </div>
    <div class='search'>
      <p>Пойск<br>по ВУЗам</p>
      <input type="text" class="search-window">
    </div>
    <div class='calendar-swipe'>
        <div class="date">{{ formattedDate }}</div>
        <img class="cal" src="../assets/calendar.svg" />
        
        <div class="time">{{ formattedTime }}</div>
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
      now: new Date(),
      updateInterval: null
    }
  },
  computed: {
    formattedDate() {
      return this.now.toLocaleDateString('ru-RU', {
        day: '2-digit',
        month: '2-digit',
      })
    },
    formattedTime() {
      return this.now.toLocaleTimeString('ru-RU', {
        hour: '2-digit',
        minute: '2-digit',
      })
    }
  },
  mounted() {
    this.updateInterval = setInterval(() => {
      this.now = new Date()
    }, 1000)
  },
  beforeDestroy() {
    clearInterval(this.updateInterval)
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
  justify-content: space-between;
  background-color: #1a1a1a;
  padding-top: 106px;
  height: 100%;
  padding-bottom: 17%;
}

.search {
  position: relative;
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
}

.calendar-swipe {
  border-radius: 30px 0 0 30px;
  background-color: #6a1b9a;
  height: fit-content;
  padding: 5px;
  margin-top: 15%;
  box-shadow: 0 4px 4px #000;
}

.calendar-swipe .date {
  font-family: 'LC Web';
  height: fit-content;
  font-size: 55px;
  color: #ffffff;
}

.calendar-swipe .time {
  font-family: 'LC Web';
  height: fit-content;
  font-size: 50px;
  color: #ffffff;
}

.calendar-swipe img {
  margin-top: 100%;
  margin-bottom: 100%;
  width: 69px;
  height: 69px;
}

.telegram-swipe {
  border-radius: 0 30px 30px 0;
  background-color: #2e2e2e;
  width: 6%;
  box-shadow: 0 4px 4px #000;
  padding: 5px;
  margin-top: 15%;
  height: auto;
}

.telegram-swipe img {
  margin-top: 120%;
  width: 90px;
  height: 90px;
  padding-bottom: 5px;
}
</style>