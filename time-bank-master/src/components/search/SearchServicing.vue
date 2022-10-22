<template>
  <div>
    <HeaderTop></HeaderTop>
    <Navbar></Navbar>
    <div class="search-input">
      <h2 style="margin: 20px 0px">搜索未结束的服务</h2>
      <el-input
          placeholder="请输入被服务人id"
          v-model="form.elderID"
          style="width: 200px;"
      >
      </el-input>
      <el-input
          placeholder="请输入服务类型"
          v-model="form.jobID"
          style="width: 200px;"
      >
      </el-input>
      <el-button type="primary" icon="el-icon-search" @click="serch">搜索</el-button>
      <div style="margin: 20px 0px">
        <h2>服务信息</h2>
        <el-table
            :data="ServeData"
            stripe
            style="width: 100%">
          <el-table-column
              prop="ServicingCount"
              label="服务数量"
              width="180">
          </el-table-column>
          <el-table-column
              prop="ServicingOlderID"
              label="被服务人id"
              width="180">
          </el-table-column>
          <el-table-column
              prop="ServicingState"
              label="服务状态">
          </el-table-column>

          <el-table-column
              prop="ServicingType"
              label="服务类型">
          </el-table-column>

          <el-table-column
              prop="ServicingValue"
              label="服务价值">
          </el-table-column>

          <el-table-column
              prop="ServicingVolID"
              label="服务ID">
          </el-table-column>

          <el-table-column
              prop="StartTime"
              label="开始时间">
          </el-table-column>

        </el-table>
      </div>
    </div>
  </div>
</template>

<script>
import HeaderTop from '@/components/home/HeaderTop.vue'
import Navbar from '@/components/home/Nav-Bar'
export default {
  components: {
    HeaderTop,
    Navbar
  },
  name: "SearchServicing",
  data(){
    return{
      form: {
        elderID: '',
        jobID: ''
      },
      ServeData:[
        {
          ServicingCount: '',
          ServicingOlderID: '',
          ServicingState: '',
          ServicingType: '',
          ServicingValue: '',
          ServicingVolID: '',
          StartTime: ''
        }
      ]
    }
  },
  methods:{
    serch(){
      if(this.form.elderID == '' || this.form.jobID == '') {
        this.$message({
          message: '请输入被服务人id和服务类型',
          type: 'warning'
        })
        return
      }
      this.$http.post('api/v1/ServicingStatusList',this.form).then(res => {
        console.log(res)
        this.ServeData = res.data.Data
        this.$message.success('查询成功')
      }).catch(err => {
        console.log(err)
        this.$message.error('查询出错')
      })
    }
  }
}
</script>

<style scoped>

.search-input{
  width: 1100px;
  height: 990px;
  margin: 0 auto;
  background: url('https://tva1.sinaimg.cn/large/e6c9d24egy1h1vd0r8t7hj20mf0iyjsq.jpg') no-repeat;
  background-size: 100% 990px;
}

</style>
