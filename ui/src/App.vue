<template>
  <div>
    <TopPanel />
  </div>
  <SearchBox />
  <CredentialList />
</template>

<script>
import SearchBox from "./components/SearchBox";
import TopPanel from "./components/TopPanel"
import CredentialList from "./components/CredentialList";

export default {
  name: "App",
  components: {
    CredentialList,
    SearchBox,
    TopPanel
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-left: 20%;
  margin-right: 20%;
}
</style>
