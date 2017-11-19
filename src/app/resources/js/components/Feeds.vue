<template>
  <div class="Feeds">

    <div v-for="feed in feeds" :key="feed.id">
      <h2>{{ feed.id }}</h2>
      <p>{{ feed.url }}</p>
    </div>

  </div>

</template>

<script>
  import Axios from 'axios'

  export default {
    data() {
      return {
        'feeds': [],
      }
    },

    methods: {
      getFeeds() {
        Axios.get('http://api.rss.dev/feeds')
          .then((response) => {
            this.feeds = response.data;
          })
          .catch((response) => {
            this.feeds = [];
            console.error(response.data);
          });
      }
    },

    beforeMount() {
      this.getFeeds();
    }
  }
</script>
