import Axios from 'axios';

export default {
  getStatus(callback) {
    return Axios.get('http://api.rss.dev/status')
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

  getSidebar(callback) {
    callback({
      name: 'All',
      list: [
        {
          name: 'link 1',
          href: '#',
        },
        {
          name: 'link 2',
          href: '#',
        },
        {
          name: 'sub folder',
          list: [
            {
              name: 'link 3',
              href: '#',
            },
          ],
        },
      ],
    });
  },
};
