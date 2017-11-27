import * as types from '../mutation-types';
import Api from '../../api/api';

const state = {
  feeds: [],
};

const getters = {
  getFeeds: state => state.feeds,
};

const actions = {
  getAllFeeds({ commit }) {
    Api.getFeeds((feeds) => {
      commit(types.SET_FEEDS, { feeds });
    });
  },
};

const mutations = {
  [types.SET_FEEDS](state, { feeds }) {
    state.feeds = feeds;
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
