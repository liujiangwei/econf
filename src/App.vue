<template>
  <a-layout id="app">
    <a-layout-header :style="{ position: 'fixed', zIndex: 1, width: '100%' }">
      <div class="logo" />
      <a-menu
        theme="dark"
        mode="horizontal"
        :defaultSelectedKeys="['2']"
        :style="{ lineHeight: '64px' }"
      >
        <a-menu-item key="1">Etcd Keys</a-menu-item>
        <a-menu-item key="2">Config Center</a-menu-item>
        <a-menu-item key="3">Service Register</a-menu-item>
      </a-menu>
    </a-layout-header>
    <a-layout-content :style="{ padding: '0 50px', marginTop: '64px' }">
      <!--
      <a-breadcrumb :style="{ margin: '16px 0' }">
        <a-breadcrumb-item>Home</a-breadcrumb-item>
        <a-breadcrumb-item>List</a-breadcrumb-item>
        <a-breadcrumb-item>App</a-breadcrumb-item>
      </a-breadcrumb>
      <div :style="{ background: '#fff', padding: '24px', minHeight: '380px' }">Content</div>
      -->
      <router-view :style="{ background: '#fff', padding: '24px', minHeight: '380px' }"></router-view>
    </a-layout-content>
    <a-layout-footer :style="{ textAlign: 'center' }">Micro Service Manage  r ©2018 Created by liujiangwei</a-layout-footer>
  </a-layout>
</template>

<script>
import HelloWorld from "./components/HelloWorld.vue";
import TopMenu from "./components/TopMenu.vue";
export default {
  name: "app",
  components: {
    TopMenu
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
.nt-layout-content {
  padding: 10px, 50px;
}
</style>
