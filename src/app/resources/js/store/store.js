import Vue from 'vue';
import Vuex from 'vuex';

import sidebar from './modules/sidebar';
import feeds from './modules/feeds';

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

export default new Vuex.Store({
  modules: {
    sidebar,
    feeds,
  },
  strict: debug,
});
