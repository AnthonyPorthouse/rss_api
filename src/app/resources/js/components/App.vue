<template>
  <div>
    <h1>{{ message }}</h1>
    <feeds></feeds>
  </div>

</template>

<script>
  import Axios from 'axios';

  import Feeds from './Feeds.vue';
  import api from "../api/api";

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
        api.getStatus((status) => {
          this.status = status;
        })
      },
    },

    beforeMount() {
      this.getStatus();
    },
  }
</script>
