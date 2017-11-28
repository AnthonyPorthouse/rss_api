<template>
  <div>
    <h1>{{ message }}</h1>
    <sidebar :list="getSidebar"></sidebar>
    <feeds></feeds>
  </div>

</template>

<script>
  import { mapGetters, mapActions } from 'vuex';

  import Feeds from './Feeds.vue';
  import Sidebar from './Sidebar/List.vue';
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

    computed: {
      ...mapGetters([
        'getSidebar',
      ]),
    },

    methods: {
      ...mapActions([
        'getSidebarItems',
      ]),

      getStatus() {
        api.getStatus((status) => {
          this.status = status;
        })
      },
    },

    beforeMount() {
      this.getStatus();
      this.getSidebarItems();
    },
  }
</script>
