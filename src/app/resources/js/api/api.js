import Axios from 'axios';

export default {
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
