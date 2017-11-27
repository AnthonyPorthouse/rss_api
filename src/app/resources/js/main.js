import Vue from 'vue';
import Store from './store/store';

import App from './components/App.vue';

/* eslint-disable no-new */
new Vue({
  el: '#app',
  store: Store,

  components: {
    app: App,
  },
});
