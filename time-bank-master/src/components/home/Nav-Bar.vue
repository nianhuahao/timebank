<template>
  <div class="nav" >
    <el-menu
        :default-active="this.$route.path"
        class="el-menu-demo"
        router
        mode="horizontal"
        @select="handleSelect"
        background-color="#ec7f30"
        text-color="#ffffff"
        style="color:#ffffff;"
    >
      <el-menu-item v-for="(item,i) in navList" :key="i" :index="item.name">
        {{ item.navItem }}
      </el-menu-item>
    </el-menu>

  </div>
</template>

<script>
export default {
  name: "NavBar",
  data(){
    return {
      activeIndex:'1',
      dialogFormVisible1: false,
      form: {
        servicingType: '',
        servicingOlderID: '',
        servicingCount: '',
      },
      navList:[
        {name:'/main',navItem:'首页'},
        {name:'/SearchServicing',navItem:'未结束的服务'},
        {name:'/SearchServed',navItem:'已结束的服务'},
        {name:'/SearchRechargeList',navItem:'充值历史'},
        {name:'/SearchSpecialTxList',navItem:'转账历史'},
      ],
      formLabelWidth: '120px'
    }
  },
  methods:{
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
    cancel(){
      this.dialogFormVisible1 = false
    },
  }
}
</script>

<style scoped>
.nav{
  width: 1100px;
  margin: 0px auto;
  height: 44px;
  margin-bottom: 40px;
}
.datafrom{
  width: 1000px;
  margin: 0px auto;
}
</style>
