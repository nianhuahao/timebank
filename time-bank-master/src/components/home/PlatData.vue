<template>
  <div class="data">
    <div class="data-top">
      <h3 style="border-bottom: 1px solid #dddddd;text-align: center">平台数据</h3>
      <p style="font-size: 12px;padding: 8px 0px 4px 8px">现有会员数：<span>{{platdata.VipNum}}</span> 团体会员：<span>{{platdata.GroupVip}}</span> </p>
      <p style="font-size: 12px;padding: 4px 0px 8px 8px">需求总数：<span style="margin-right: 10px;">{{platdata.NeedSum}}</span> 对接按成：<span>{{platdata.Succeed}}</span> </p>
    </div>
    <el-tabs v-model="activeName" type="card" @tab-click="handleClick" class="tabs">
      <el-tab-pane label="服务" name="first" style="">
        <el-button type="warning" @click="dialogFormVisible1 = true" class="server_bt" style="padding: 0 20px">结束服务</el-button>
        <el-button type="warning" @click="dialogFormVisible2 = true" class="server_bt" style="padding: 0 20px">取消服务</el-button>
        <el-button type="warning" @click="dialogFormVisible3 = true" class="server_bt" style="padding: 0 20px">接受服务</el-button>
        <el-button type="warning" @click="dialogFormVisible4 = true" class="server_bt" style="padding: 0 20px">发布服务</el-button>
      </el-tab-pane>
      <el-tab-pane label="工分" name="second">
        <el-button type="warning" @click="dialogFormVisible5 = true" class="server_bt" style="padding: 0 20px">工分转移</el-button>
        <el-button type="warning" @click="dialogFormVisible6 = true" class="server_bt" style="padding: 0 20px">工分继承</el-button>
        <el-button type="warning" @click="dialogFormVisible7 = true" class="server_bt" style="padding: 0 20px">工分充值</el-button>
      </el-tab-pane>
    </el-tabs>
    <div class="info">
      <router-link to="/vipinfo" class="info-bt" style="padding: 0 20px"><img :src="info.photo" alt="" style="border-radius: 50%; width: 52px;height: 52px;float: left"></router-link>
      <p style="width: 190px;text-align: center;height:52px;font-size: 20px;">杨大爷</p>
      <button style="" @click="login">登入</button>
      <button @click="register">注册</button>
      <div style="border-bottom: 1px solid #dddddd;"></div>
      <div style="width: 220px;margin: 0 auto;">
        <li v-for="(item,index) in images" :key="index">
          <img :src="item" alt="">
        </li>
      </div>
    </div>
<!--    结束服务-->
    <el-dialog title="结束服务" :visible.sync="dialogFormVisible1" class="datafrom">
      <el-form :model="form1">
        <el-form-item label="服务类型" :label-width="formLabelWidth">
          <el-input v-model="form1.TxType" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="被服务人ID" :label-width="formLabelWidth">
          <el-input v-model="form1.CustID" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="服务人ID" :label-width="formLabelWidth">
          <el-input v-model="form1.ServeID" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="评论" :label-width="formLabelWidth">
          <el-input v-model="form1.ServeComment" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel1">取 消</el-button>
        <el-button type="primary" @click="submit1">确 定</el-button>
      </div>
    </el-dialog>
<!--取消服务-->
    <el-dialog title="取消服务" :visible.sync="dialogFormVisible2" class="datafrom">
      <el-form :model="form2">
        <el-form-item label="服务类型" :label-width="formLabelWidth">
          <el-input v-model="form2.servicingType" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="被服务人ID" :label-width="formLabelWidth">
          <el-input v-model="form2.servicingOlderID" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel2">取 消</el-button>
        <el-button type="primary" @click="submit2">确 定</el-button>
      </div>
    </el-dialog>
<!--接受服务-->
    <el-dialog title="接受服务" :visible.sync="dialogFormVisible3" class="datafrom">
      <el-form :model="form3">
        <el-form-item label="服务类型" :label-width="formLabelWidth">
          <el-input v-model="form3.servicingType" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="接收人ID" :label-width="formLabelWidth">
          <el-input v-model="form3.servicingVolID" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="被服务人ID" :label-width="formLabelWidth">
          <el-input v-model="form3.servicingOlderID" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel3">取 消</el-button>
        <el-button type="primary" @click="submit3">确 定</el-button>
      </div>
    </el-dialog>
<!--    发布服务-->
    <el-dialog title="发布服务" :visible.sync="dialogFormVisible4" class="datafrom">
      <el-form :model="form4">
        <el-form-item label="服务类型" :label-width="formLabelWidth">
          <el-input v-model="form4.ServicingType" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="被服务人ID" :label-width="formLabelWidth">
          <el-input v-model="form4.ServicingOlderID" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="服务数量" :label-width="formLabelWidth">
          <el-input v-model="form4.ServicingCount" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel4">取 消</el-button>
        <el-button type="primary" @click="submit4">确 定</el-button>
      </div>
    </el-dialog>
<!--    公分转移-->
    <el-dialog title="公分转移" :visible.sync="dialogFormVisible5" class="datafrom">
      <el-form :model="gongfenfrom1">
        <el-form-item label="发送人ID" :label-width="formLabelWidth">
          <el-input v-model="gongfenfrom1.fromAssetId" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="接收人ID" :label-width="formLabelWidth">
          <el-input v-model="gongfenfrom1.toAssetId" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="工分数" :label-width="formLabelWidth">
          <el-input v-model="gongfenfrom1.assetValue" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancle5">取 消</el-button>
        <el-button type="primary" @click="submit5">确 定</el-button>
      </div>
    </el-dialog>
<!--    公分继承-->
    <el-dialog title="公分继承" :visible.sync="dialogFormVisible6" class="datafrom">
      <el-form :model="gongfenfrom2">
        <el-form-item label="继承人ID" :label-width="formLabelWidth">
          <el-input v-model="gongfenfrom2.fromAssetId" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="被继承人ID" :label-width="formLabelWidth">
          <el-input v-model="gongfenfrom2.toAssetId" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancle6">取 消</el-button>
        <el-button type="primary" @click="submit6">确 定</el-button>
      </div>
    </el-dialog>
<!--    公分充值-->
    <el-dialog title="公分充值" :visible.sync="dialogFormVisible7" class="datafrom">
      <el-form :model="gongfenfrom3">
        <el-form-item label="充值人ID" :label-width="formLabelWidth">
          <el-input v-model="gongfenfrom3.toUserID" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="充值数量" :label-width="formLabelWidth">
          <el-input v-model="gongfenfrom3.rechargeValue" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancle7">取 消</el-button>
        <el-button type="primary" @click="submit7">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "PlatData",
  data(){
    return{
      activeName: 'first',
      platdata:{
        VipNum:'118262',
        GroupVip:'3757',
        NeedSum:'173378',
        Succeed:'106461',
      },
      info:{
        photo:"https://www.nstimebank.com/images/tx.jpg"
      },
      images:[

          "https://www.nstimebank.com/images/sjzx.jpg",
          "https://www.nstimebank.com/images/sjyl.jpg"
      ],
      dialogFormVisible1: false,
      dialogFormVisible2: false,
      dialogFormVisible3: false,
      dialogFormVisible4: false,
      form1: {
        TxType: '',
        CustID: '',
        ServeID: '',
        ServeComment: '',
      },
      form2: {
        servicingType: '',
        servicingOlderID: '',
      },
      form3: {
        servicingType: '',
        servicingVolID: '',
        servicingOlderID: '',
      },
      form4: {
        ServicingType: '',
        ServicingOlderID: '',
        ServicingCount: '',
      },
      dialogFormVisible5: false,
      dialogFormVisible6: false,
      dialogFormVisible7: false,
      gongfenfrom1:{
        fromAssetId:'',
        toAssetId:'',
        assetValue:'',
      },
      gongfenfrom2:{
        fromAssetId:'',
        toAssetId:'',

      },
      gongfenfrom3:{
        toUserID:'',
        rechargeValue:'',
      },
      formLabelWidth: '120px'
    }
  },
  methods:{
    login() {
      this.$router.push('/Login')
    },
    register(){
      this.$router.push('/Register')
    },
    handleClick(){

    },
    activeName(){

    },
    //提交结束服务的信息
    submit1() {
      this.dialogFormVisible1 = false
      this.$http.post('api/v1/DoneServicing',this.form1).then(res => {
        console.log(res)
        const h = this.$createElement;
        this.$notify({
          title: '您的交易ID为',
          message: h('i', { style: 'color: teal'}, res.data.Data.TxID),
        });
        this.$message.success('结束服务成功')
      }).catch(err => {
        console.log(err,"添加失败，请求出错")
      })
    },
    //提交取消服务的信息
    submit2() {
      this.dialogFormVisible2 = false
      this.$http.post('api/v1/CloseServicing',this.form2).then(res => {
        console.log(res)
        this.$message.success('取消服务成功')
      }).catch(err => {
        console.log(err,"添加失败，请求出错")
      })
    },
    //提交接受服务的信息
    submit3() {
      this.dialogFormVisible3 = false
      this.$http.post('api/v1/AcceptServicing',this.form3).then(res => {
        console.log(res)
        this.$message.success('接受服务成功')
      }).catch(err => {
        console.log(err,"添加失败，请求出错")
      })
    },
    //提交发布服务的信息
    submit4() {
      this.dialogFormVisible4 = false
      this.$http.post('api/v1/CreateServicing',this.form4).then(res => {
        console.log(res)
        this.$message.success('发布服务成功')
      }).catch(err => {
        console.log(err,"添加失败，请求出错")
      })
    },
    //提交公分转移的信息
    submit5() {
      this.dialogFormVisible5 = false
      this.$http.post('api/v1/TransferAsset',this.gongfenfrom1).then(res => {
        console.log(res)
        const h = this.$createElement;
        this.$notify({
          title: '您的交易ID为',
          message: h('i', { style: 'color: teal'}, res.data.Data.TransferID),
        });
        this.$message.success('转移工分成功')
      }).catch(err => {
        console.log(err,"添加失败，请求出错")
      })
    },
    //提交公分继承的信息
    submit6() {
      this.dialogFormVisible6 = false
      this.$http.post('api/v1/InheritAsset',this.gongfenfrom2).then(res => {
        console.log(res)
        const h = this.$createElement;
        this.$notify({
          title: '您的交易ID为',
          message: h('i', { style: 'color: teal'}, res.data.Data.TransferID),
        });
        this.$message.success('继承成功')
      }).catch(err => {
        console.log(err,"添加失败，请求出错")
      })
    },
    //提交公分充值的信息
    submit7() {
      this.dialogFormVisible7 = false
      this.$http.post('api/v1/RechargeAsset',this.gongfenfrom3).then(res => {
        console.log(res)
        const h = this.$createElement;
        this.$notify({
          title: '您的交易ID为',
          message: h('i', { style: 'color: teal'}, res.data.Data.RechargeID),
        });
        this.$message.success('添加工分成功')
      }).catch(err => {
        console.log(err,"添加失败，请求出错")
      })
    },
      //取消结束服务的弹窗
    cancel1(){
      this.dialogFormVisible1 = false
    },
    //取消取消服务的弹窗
    cancel2(){
      this.dialogFormVisible2 = false
    },
    //取消接受服务的弹窗
    cancel3(){
      this.dialogFormVisible3 = false
    },
    //取消发布服务的弹窗
    cancel4(){
      this.dialogFormVisible4 = false
    },
    //取消公分转移的弹窗
    cancle5(){
      this.dialogFormVisible5 = false
    },
    //取消公分继承的弹窗
    cancle6(){
      this.dialogFormVisible6 = false
    },
    //取消公分充值的弹窗
    cancle7(){
      this.dialogFormVisible7 = false
    },
  }
}
</script>

<style>
.data-top{
  width: 240px;
  height: 87px;
  border: 1px solid #dddddd;
}
.tabs{
  width: 240px;
  height: 150px;
  margin-top: 10px;
  border: 1px solid #dddddd;
}
.info{
  width: 240px;
  height: 350px;
  margin-top: 10px;
  padding: 10px;
  border: 1px solid #dddddd;
}
.data-top span{
  color: #ec7f30;
}
.server_bt{
  border-radius: 3px;
  background: #fd7700;
  float: left;
  border-bottom: 1px solid #ff4200;
  width: 88px;
  height: 26px;
  color: #fff;
  border-left: none;
  border-top: none;
  border-right: none;
  margin: 10px 6px;
  cursor: pointer;
}
.info button{
  border-radius: 3px;
  background: #fd7700;
  border-bottom: 1px solid #ff4200;
  width: 88px;
  height: 26px;
  color: #fff;
  border-left: none;
  border-top: none;
  border-right: none;
  margin: 10px 6px;
  cursor: pointer;
}
li{
  list-style: none;
}
</style>
