<template>
  <div class="Feeds">

    <div v-for="feed in getFeeds" :key="feed.id">
      <h2>{{ feed.id }}</h2>
      <p>{{ feed.url }}</p>
    </div>

  </div>
</template>

<style lang="scss">

  .Feeds {
    grid-column: span 9;
  }

</style>

<script>
  import { mapGetters, mapActions } from 'vuex';

  export default {
    data() {
      return {};
    },

    computed: {
      ...mapGetters([
        'getFeeds',
      ]),
    },

    methods: {
      ...mapActions([
        'getAllFeeds',
      ]),
    },

    beforeMount() {
      this.getAllFeeds();
    }
  }
</script>
