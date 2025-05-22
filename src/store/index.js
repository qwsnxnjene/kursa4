import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    selectedCity: 'kazan',
    isAuthenticated: false
  },
  mutations: {
    setCity(state, city) {
      state.selectedCity = city
    },
    setAuthentication(state, value) {
      state.isAuthenticated = value
    }
  },
  getters: {
    currentCity: state => state.selectedCity,
    isAuthenticated: state => state.isAuthenticated
  },
  plugins: [
    createPersistedState({
      key: 'uni-app',
      paths: ['selectedCity', 'isAuthenticated'],
    })
  ]
})