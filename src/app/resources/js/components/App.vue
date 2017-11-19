<template>
  <div>
    <h1>{{ message }}</h1>
    <feeds></feeds>
  </div>

</template>

<script>
  import Axios from 'axios';

  import Feeds from './Feeds.vue';

  export default {
    data: () => {
      return {
        message: 'Hello, World!',

        status: 'Unknown',
      }
    },

    components: {
      'feeds': Feeds,
    },

    methods: {
      getStatus() {
        Axios.get("/status")
          .then((response) => {
            this.status = response.data;
          });
      },
    },

    beforeMount() {
      this.getStatus();
    },
  }
</script>
