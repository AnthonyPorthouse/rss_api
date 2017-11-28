<template>
  <a :href="item.href">{{ item.name }}</a>
</template>

<script>
  export default {
    props: {
      item: Object
    },
  }
</script>
