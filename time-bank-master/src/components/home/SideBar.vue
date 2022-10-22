<template>
  <div class="sidebar">
    <h5>服务分类</h5>
    <el-menu
        default-active="2"
        class="el-menu-vertical-demo"
        @open="handleOpen"
        @close="handleClose"
        background-color="#f7f7f7"
        text-color="#333333"
        active-text-color="black"
        style="color:#333333;"
    >
      <el-menu-item index="1">
        <span slot="title">做饭</span>
      </el-menu-item>
      <el-menu-item index="2">
        <span slot="title">游戏代练</span>
      </el-menu-item>
      <el-menu-item index="3">
        <span slot="title">洗衣服</span>
      </el-menu-item>
      <el-menu-item index="4">
        <span slot="title">洗被子</span>
      </el-menu-item>
      <el-menu-item index="5">
        <span slot="title">按摩</span>
      </el-menu-item>
      <el-menu-item index="6">
        <span slot="title">理发</span>
      </el-menu-item>
      <el-menu-item index="7">
        <span slot="title">陪伴就医</span>
      </el-menu-item>
      <el-menu-item index="8">
        <span slot="title">修理电器</span>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  data () {
    return {
    }
  },
  methods:{
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    }
  }
}
</script>

<style >
.sidebar{
  width: 170px;
  height: 578px;
  border: 2px solid #ff6624;
}
.sidebar h5{
  width: 169px;
  height: 40px;
  line-height: 40px;
  text-align: center;
  color: #ffffff;
  font-size: 16px;
  background-color: #ec7f30;
}
.el-submenu .el-menu-item{
  min-width: 0px;
}
</style>
