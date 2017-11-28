<template>
  <div>
    <h1>{{ list.name }}</h1>
    <ul>
      <li v-for="item in list.list">
        <sidebar-item v-if="item.href" :item="item"></sidebar-item>
        <sidebar-list v-else-if="item.list" :sub="true" :list="item"></sidebar-list>
      </li>
    </ul>
  </div>
</template>

<script>
  import SidebarItem from './Item.vue';

  export default {
    name: 'sidebar-list',

    props: {
      list: {
        type: Object
      }
    },

    components: {
      'sidebar-item': SidebarItem,
    },
  }
</script>
