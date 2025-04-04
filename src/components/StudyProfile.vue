<template>
  <div class='main-content'>
    <div class='vuz'>
        <p>КФУ</p>
    </div>
    <div class="first-article">
        
        <img src="../assets/кфу1.png" alt="">
        <p>Lorem ipsum dolor sit amet,<br>
        consectetur adipiscing elit. Mauris 
        viverra non nibh eu scelerisque. 
        Phasellus sodales justo purus, 
        quis placerat diam convallis eu. 
        Nulla facilisi. Praesent maximus 
        lacus quis massa luctus cursus. 
        Nullam ac rutrum ipsum, viverra egestas ipsum.</p>
    </div>
    <div class="second-article">
        <p>Lorem ipsum dolor sit amet,<br>
        consectetur adipiscing elit. Mauris 
        viverra non nibh eu scelerisque. 
        Phasellus sodales justo purus, 
        quis placerat diam convallis eu. 
        Nulla facilisi. Praesent maximus 
        lacus quis massa luctus cursus. 
        Nullam ac rutrum ipsum, viverra egestas ipsum.</p>
        <img src="../assets/кфу2.png" alt="">
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
        <input type="text" class="name-window">
    </div>
    <div class="input-name">
        <p>номер телефона</p>
        <input type="text" class="name-window">
    </div>
    <div class="input-name">
        <p>email</p>
        <input type="text" class="name-window">
    </div>
    <div class="ready-btn">
      <button>
              Готово!
      </button>
    </div>
  </div>
</template>

<script>
import CalendarSwipe from './CalendarSwipe.vue';

export default {
  name: 'MainContent',
  components: [CalendarSwipe],
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

.vuz p {
  color: #fff;
  font-family: 'IF Kica';
  font-size: 60px;
  margin: 20px;
}

.first-article {
    text-align: left;
    display: flex;
}

.first-article img {
    width: 600px;
    height: auto;
    margin: 10px 600px 15px 100px;
}

.first-article p {
    margin-top: 50px;
    color: #fff;
    font-family: 'IF Kica';
    font-size: 30px;
    width: 800px;
}

.second-article {
    text-align: right;
    display: flex;
}

.second-article img {
    width: 600px;
    height: auto;
    margin: 10px 100px 15px 600px;
}

.second-article p {
    margin-left:50px;
    margin-top: 50px;
    color: #fff;
    font-family: 'IF Kica';
    font-size: 30px;
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
}

</style>