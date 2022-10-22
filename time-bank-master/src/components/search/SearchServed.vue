<template>
  <div>
    <HeaderTop></HeaderTop>
    <Navbar></Navbar>
    <div class="search-input">
      <h2 style="margin: 20px 0px">搜索已经结束的服务</h2>
      <el-input
          placeholder="请输入交易id"
          v-model="form.txID"
          style="width: 200px;"
      >
      </el-input>
      <el-input
          placeholder="请输入被服务人ID"
          v-model="form.elderID"
          style="width: 200px;"
      >
      </el-input>
      <el-input
          placeholder="请输入志愿者ID"
          v-model="form.volunteerID"
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
              prop="CustID"
              label="服务ID"
              width="180">
          </el-table-column>
          <el-table-column
              prop="EndTime"
              label="结束时间"
              width="180">
          </el-table-column>
          <el-table-column
              prop="ServeComment"
              label="评论">
          </el-table-column>

          <el-table-column
              prop="ServeID"
              label="服务ID">
          </el-table-column>

          <el-table-column
              prop="TxID"
              label="TxID">
          </el-table-column>

          <el-table-column
              prop="TxType"
              label="类型">
          </el-table-column>

          <el-table-column
              prop="WorkHours"
              label="工作时间">
          </el-table-column>

          <el-table-column
              prop="WorkValue"
              label="工作价值">
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>
</template>

<script>
import HeaderTop from '@/components/home/HeaderTop.vue'
// import NavBar from '@/components/home/Nav-Bar';
import Navbar from '@/components/home/Nav-Bar'
export default {
  components: {
    HeaderTop,
    // NavBar
    Navbar
  },
  name: "SearchServed",
  data(){
    return{
      form: {
        txID: '',
        elderID: '',
        volunteerID: ''
      },
      ServeData:[
        {
          CustID: '',
          EndTime: '',
          ServeComment: '',
          ServeID: '',
          TxID: '',
          TxType: '',
          WorkHours: '',
          WorkValue: ''
        },
      ]
    }
  },
  methods:{
    serch(){
      if(this.form.txID == ''){
        this.$message({
          message: '请输入交易id',
          type: 'warning'
        });
        return;
      }
      if(this.form.elderID == ''){
        this.$message({
          message: '请输入被服务人ID',
          type: 'warning'
        });
        return;
      }
      if(this.form.volunteerID == ''){
        this.$message({
          message: '请输入志愿者ID',
          type: 'warning'
        });
        return;
      }
      this.$http.post('api/v1/TxList',this.form).then(res => {
        console.log(res)
        this.ServeData=res.data.Data
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
