<template>
  <div class="top">
    <div>
      <p style="float:left;margin: 0px;">欢迎来到时光元链！</p>
      <button style="float:right" onmouseover="this.style.color='#ec7f30';" onmouseout="this.style.color='#000000';">
        <router-link to="/main">首页</router-link>
      </button>
      <button style="float:right"
              onmouseover="this.style.color='#ec7f30';"
              onmouseout="this.style.color='#000000';"
      >
        <router-link :to="{name:'Login'}">登录</router-link>
      </button>
    </div>
    <div class="dav">
      <img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h1vbsfi4myj20u005kwf0.jpg" alt="" style="width: 100%;height: 180px;">
    </div>
  </div>
</template>

<script>
export default {
  name: 'HeaderTop',
  data(){
    return{

    }
  },
  methods:{
    change(){
      alert("点击事件");
    }
  }
}
</script>
<style>
.top {
  width: 1100px;
  margin: 0px auto;
}
button{
  margin: 0;
  padding: 0px 20px;
  border: 1px solid transparent;
  background: #ffffff;
  outline: none;
}
</style>

