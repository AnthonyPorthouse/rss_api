<template>
  <sidebar-list :list="getSidebar"></sidebar-list>
</template>

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
