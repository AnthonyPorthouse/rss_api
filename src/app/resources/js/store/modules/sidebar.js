import * as types from "../mutation-types";

const state = {
  items: [],
};

const getters = {
  getItems: (state) => {
    return state.items;
  }
};

const actions = {
  getSidebarItems({ commit }) {
    commit(types.GET_SIDEBAR_ITEMS)
  },
};

const mutations = {
  [types.GET_SIDEBAR_ITEMS] (state, { items }) {
    state.items = items;
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
