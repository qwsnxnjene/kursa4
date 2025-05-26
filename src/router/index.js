import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../Start.vue'
import Profile from '../Profile.vue'
import Study from '../Study.vue'
import AddUniversity from '../Add.vue'
import StudentTest from '../Test.vue'

Vue.use(VueRouter) // Подключаем плагин

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/profile',
    name: 'Profile',
    component: Profile
  },
  {
    path: '/study',
    name: 'Study',
    component: Study
  },
  {
    path: '/universities/:id',
    component: Study,
    props: true
  },
  {
    path: '/add-university',
    component: AddUniversity,
  },
  {
    path: '/student-test',
    component: StudentTest
  }
]

const router = new VueRouter({
  mode: 'history', // или 'hash' если нужно
  base: process.env.BASE_URL,
  routes
})

export default router