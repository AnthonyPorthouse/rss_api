import Axios from 'axios';

export default {
  getStatus(callback) {
    return Axios.get("http://api.rss.dev/status")
      .then((response) => {
        callback(response.data);
      });
  },

  getFeeds(callback) {
    return Axios.get('http://api.rss.dev/feeds')
      .then((response) => {
        callback(response.data);
      })
      .catch((response) => {
        throw response;
      });
  },
};
