import { createStore } from 'vuex'
import { user } from "@/store/module/user"
import VuexPersistence from 'vuex-persist'

const vuexLocal = new VuexPersistence({
  storage: window.localStorage,
  modules: ['user']
})

export default createStore({
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    user,
  },
  plugins: [vuexLocal.plugin]
})
