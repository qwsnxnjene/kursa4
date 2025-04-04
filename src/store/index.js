import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    selectedCity: 'kazan'
  },
  mutations: {
    setCity(state, city) {
      state.selectedCity = city
    }
  },
  getters: {
    currentCity: state => state.selectedCity
  },
  plugins: [
    createPersistedState({
      key: 'uni-app',
      paths: ['selectedCity']
    })
  ]
})