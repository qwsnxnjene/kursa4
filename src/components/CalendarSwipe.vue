<template>
    <div class='calendar-swipe' @mouseenter="isExpanded = true" @mouseleave="isExpanded = false">
        <div class="date">{{ formattedDate }}</div>
        <img v-if="!isExpanded" class="cal" src="../assets/calendar.svg" />
        

        <transition name="slide">
          <div v-if="isExpanded" class="expanded-content">
            <p>20.06.25<br>начало приёма документов</p>
            <p>25.07.25<br>окончание приёма документов(без ДВИ)</p>
            <p>27.07.25<br>публикация конкурсных списков</p>
            <p>05.08.25<br>последний день приёма согласий</p>
            <p>06.08.25 - 07.08.25<br>публикация приказов о зачислении</p>
          </div>
        </transition>
        <div class="time">{{ formattedTime }}</div>
    </div>
</template>

<script>
export default {
    name: "CalendarSwipe",
    data() {
        return {
            isExpanded: false,
            updateInterval: null,
            now: new Date(),
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
  
}
</script>

<style scoped>
.calendar-swipe {
  border-radius: 30px 0 0 30px;
  background-color: #6a1b9a;
  height: fit-content;
  padding: 5px;
  margin-top: 15%;
  box-shadow: 0 4px 4px #000;
  cursor: pointer;
}

.calendar-swipe:hover {
  width: 300px; 
  height: fit-content;
}

.expanded-content {
  height: fit-content;
  padding-left: 5px;
  margin-left: 5px;
  border-left: 1px solid rgba(255,255,255,0.3);
  font-family: 'LC Web';
  font-size: 30px;
  text-align: left;
  color: #fff;
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
</style>