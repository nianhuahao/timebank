<template>
  <div>
    <HeaderTop></HeaderTop>
    <Navbar></Navbar>
    <div class="search-input">
      <h2 style="margin: 20px 0px">转账历史</h2>
      <el-input
          placeholder="请输入被服务人id"
          v-model="form.specialTxID"
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
              prop="AssetValue"
              label="转账数量"
              width="180">
          </el-table-column>
          <el-table-column
              prop="FromAsset"
              label="转账人ID"
              width="180">
          </el-table-column>
          <el-table-column
              prop="ToAsset"
              label="收账人ID">
          </el-table-column>

          <el-table-column
              prop="TransferID"
              label="转账ID">
          </el-table-column>

          <el-table-column
              prop="TransferTime"
              label="转账时间">
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
  name: "SearchRechargeList",
  data(){
    return{
      form: {
        specialTxID: '',
      },
      ServeData:[
        {
          AssetValue: '',
          FormAsset: '',
          ToAsset: '',
          TransferID: '',
          TransferTime: '',
        },
      ]
    }
  },
  methods:{
    serch(){
      if(this.form.rechargeID == ''){
        this.$message({
          message: '请输入被服务人id',
          type: 'warning'
        });
        return;
      }
      this.$http.post('api/v1/SpecialTxList',this.form).then(res => {
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
  background: url('https://tva1.sinaimg.cn/large/e6c9d24egy1h1vdaror4wj20pf0kjgmq.jpg') no-repeat;
  background-size: 100% 100%;
}

</style>
