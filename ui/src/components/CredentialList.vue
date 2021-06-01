<template>
  <div class="credentials-container">
    <div v-for="cred in credentials" :key="cred.id">
      {{cred.title}}
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    credentials() {
      return this.$store.state.credentials;
    }
  }
}
</script>

<style scoped>

</style>
