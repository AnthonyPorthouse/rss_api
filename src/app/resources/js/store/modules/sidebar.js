import * as types from '../mutation-types';
import api from '../../api/api';

const state = {
  items: [],
};

const getters = {
  getSidebar: state => state.items,
};

const actions = {
  getSidebarItems({ commit }) {
    api.getSidebar((sidebar) => {
      commit(types.SET_SIDEBAR, sidebar);
    });
  },
};

const mutations = {
  [types.SET_SIDEBAR](state, items) {
    state.items = items;
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
