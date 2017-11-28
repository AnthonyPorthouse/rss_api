<template>
  <sidebar-list class="Sidebar" :list="getSidebar"></sidebar-list>
</template>

<style lang="scss">
  .Sidebar {
    grid-column: 1 / span 3;
  }

</style>

<script>
  import { mapGetters, mapActions } from 'vuex';
  import api from "../../api/api";
  import SidebarList from './List.vue';

  export default {
    components: {
      'sidebar-list': SidebarList,
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
    },

    beforeMount() {
      this.getSidebarItems();
    }
  }
</script>
