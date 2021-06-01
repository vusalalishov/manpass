<template>
  <div v-for="cred in credentials" key="cred.id">
    {{cred.title}}
  </div>
</template>

<script>
export default {
  computed: {
    credentials: () => {
      return this.$state.credentials;
    }
  }
}
</script>

<style scoped>

</style>