<template>

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
        Axios.get('/api/getFeeds')
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
