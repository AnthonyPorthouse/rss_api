<template>
  <div class="App">

    <menu-bar></menu-bar>
    <sidebar></sidebar>
    <feeds></feeds>
  </div>

</template>

<style lang="scss">
  .App {
    display: grid;
    grid-template-columns: repeat(12, 1fr);
    grid-gap: .5rem;
  }
</style>

<script>
  import Menu from './Menu/Menu.vue';
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
      'menu-bar': Menu,
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
