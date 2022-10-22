<template>
<!--  轮播图-->
  <div class="slide">
    <el-carousel  >
      <el-carousel-item v-for="item in images" :key="item">
        <img :src="item" width="100%" height="100%">
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {
  name: "SlideView",
  data(){
    return {
      images:[
        "https://tva1.sinaimg.cn/large/e6c9d24egy1h1vbkq9vebj20u00gvq5p.jpg",
          "https://tva1.sinaimg.cn/large/e6c9d24egy1h1vbl3t7zcj20tq0gqjtu.jpg",
          "https://tva1.sinaimg.cn/large/e6c9d24egy1h1vbld4hy6j20xq0iz0vl.jpg",
          "https://tva1.sinaimg.cn/large/e6c9d24egy1h1vbll569bj20tz0gu0vh.jpg"
      ]
    }
  }
}
</script>

<style scoped>
  .slide{
    width: 666px;
    height: 288px;
    /*border: 1px solid gray;*/
  }
</style>
