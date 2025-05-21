<template>
    <div class='calendar-swipe' @mouseenter="isExpanded = true; isImg = false" @mouseleave="showImage">
        <div class="date">{{ formattedDate }}</div>
        <img v-if="isImg" class="cal" src="../assets/calendar.svg" />
        
        <div
          class="expanded-content"
          :class="{ expanded: isExpanded }"
        >
          <p>20.06.25<br>начало приёма документов</p>
          <p>25.07.25<br>окончание приёма документов(без ДВИ)</p>
          <p>27.07.25<br>публикация конкурсных списков</p>
          <p>05.08.25<br>последний день приёма согласий</p>
          <p>06.08.25 - 07.08.25<br>публикация приказов о зачислении</p>
        </div>
        <div class="time">{{ formattedTime }}</div>
    </div>
</template>

<script>
export default {
    name: "CalendarSwipe",
    data() {
        return {
            isImg: true,
            isExpanded: false,
            updateInterval: null,
            now: new Date(),
        }
    },

    methods: {
      showImage() {
        this.isExpanded = false
        setTimeout(() => {this.isImg = true}, 400)
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
  min-width: 160px;
  width: 160px;
  max-width: 480px;
  padding: 5px;
  margin-top: 10%;
  box-shadow: 0 4px 4px #000;
  cursor: pointer;
  transition: width 0.4s cubic-bezier(.4,0,.2,1);
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  height: 600px;
  overflow: hidden; /* добавьте, чтобы текст не вылезал */
}

.calendar-swipe:hover {
  width: 480px;
}

.calendar-swipe .date {
  font-family: 'LC Web';
  font-size: 55px;
  color: #ffffff;
  margin-bottom: 10px;
  margin-top: 10px;
}

.calendar-swipe .time {
  font-family: 'LC Web';
  font-size: 50px;
  color: #ffffff;
  margin-top: 10px;
  margin-bottom: 10px;
}

.calendar-swipe img {
  width: 69px;
  height: 69px;
  margin: auto 100px;
  flex-shrink: 0;
}

.expanded-content {
  width: 100%;
  opacity: 0;
  transition: opacity 0.4s cubic-bezier(.4,0,.2,1);
  overflow: hidden;
  padding-left: 0;
  margin-left: 0;
  border-left: none;
  font-family: 'LC Web';
  font-size: 25px;
  text-align: left;
  color: #fff;
  pointer-events: none;
  /* удалено: position, left, top, background, border-radius, box-shadow, z-index, height */
}

.expanded-content.expanded {
  opacity: 1;
  padding-left: 20px;
  margin-left: 0;
  border-left: 1px solid rgba(255,255,255,0.3);
  pointer-events: auto;
}
</style>