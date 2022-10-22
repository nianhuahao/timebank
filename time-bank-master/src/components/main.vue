<template>
  <div id="app">
    <HeaderTop></HeaderTop>
    <NavBar></NavBar>
    <div class="main">
      <SideBar class="son1"></SideBar>
      <div class="middle">
        <SlideView class="son2"></SlideView>
        <NoticeView class="son3"></NoticeView>
      </div>
      <div class="right">
        <PlatData></PlatData>
      </div>
    </div>
  </div>
</template>

<script>
import HeaderTop from '@/components/home/HeaderTop.vue'
import NavBar from "@/components/home/Nav-Bar";
import SideBar from "@/components/home/SideBar";
import SlideView from "@/components/home/SlideView";
import NoticeView from "@/components/home/NoticeView";
import PlatData from "@/components/home/PlatData";
export default {
  name: 'main',
  components: {
    HeaderTop,
    NavBar,
    SideBar,
    SlideView,
    NoticeView,
    PlatData
  }
}
</script>

<style>
.main{
  width: 1100px;
  height: 1000px;
  margin: 0 auto;
}
.son1{
  float: left;
}
.middle{
  width: 666px;
  height: 600px;
  float: left;
  padding-left: 15px;
}
.son2{
  float: left;
}
.son3{
  float: left;
}
.right{
  float: right;
}
</style>
