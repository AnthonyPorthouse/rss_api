<template>
  <div>
    <h1>{{ message }}</h1>
    <sidebar></sidebar>
    <feeds></feeds>
  </div>

</template>

<script>
  import Feeds from './Feeds.vue';
  import Sidebar from './Sidebar/Sidebar.vue';
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
      'sidebar': Sidebar,
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
