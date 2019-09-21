<template>
  <div>
    <a-menu v-model="current" mode="horizontal">
      <a-menu-item key="mail">
        <a-icon type="mail" />Etcd Keys
      </a-menu-item>
      <a-menu-item key="app">
        <a-icon type="appstore" />Config Conter
      </a-menu-item>
      <a-menu-item>
                <a-icon type="appstore" />Service Register
      </a-menu-item>
    </a-menu>
  </div>
</template>
<script>
export default {
  name: "TopMenu",
  data() {
    return {
      current: ["mail"]
    };
  }
};
</script>
