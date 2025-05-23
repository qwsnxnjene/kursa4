import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    selectedCity: 'kazan',
    isAuthenticated: false,
    authToken: '',
    userName: '',
    userEmail: '',
    userSurname: '',
    userRole: '',
  },
  mutations: {
    setCity(state, city) {
      state.selectedCity = city
    },
    setAuthentication(state, value) {
      state.isAuthenticated = value
    },
    setAuthToken(state, token) {
      state.authToken = token
    },
    setUserName(state, name) {
      state.userName = name
    },
    setUserEmail(state, email) {
      state.userEmail = email
    },
    setUserSurname(state, surname) {
      state.userSurname = surname
    },
    setUserRole(state, role) {
      state.userRole = role
    }
  },
  getters: {
    currentCity: state => state.selectedCity,
    isAuthenticated: state => state.isAuthenticated,
    authToken: state => state.authToken,
    userName: state => state.userName,
    userEmail: state => state.userEmail,
    userSurname: state => state.userSurname,
    userRole: state => state.userRole,
  },
  plugins: [
    createPersistedState({
      key: 'uni-app',
      paths: ['selectedCity', 'isAuthenticated', 'authToken', 'userName', 'userEmail', 'userSurname', 'userRole']
    })
  ]
})