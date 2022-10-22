<template>
  <div class="register">
    <el-form ref="Registerform" :model="info" label-width="80px" class="res_box">
      <h1 style="text-align: center;margin-bottom: 10px;">用户注册</h1>
      <el-form-item label="能力">
        <el-input v-model="info.ability"></el-input>
      </el-form-item>
      <el-form-item label="用户姓名">
        <el-input v-model="info.userName"></el-input>
      </el-form-item>
      <el-form-item label="身份证号">
        <el-input v-model="info.userIdentification"></el-input>
      </el-form-item>
      <el-form-item label="性别">
        <el-input v-model="info.sex"></el-input>
      </el-form-item>
      <el-form-item label="出生日期">
        <el-input v-model="info.birthday"></el-input>
      </el-form-item>
      <el-form-item label="地址">
        <el-input v-model="info.address"></el-input>
      </el-form-item>
      <el-form-item label="邮政编码">
        <el-input v-model="info.postcode"></el-input>
      </el-form-item>
      <el-form-item label="推荐人">
        <el-input v-model="info.recommendID"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
        <el-button @click="cancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: "Register",
  data() {
    return {
      info: {
        ability:'',
        userName: '',
        userIdentification: '',
        sex: '',
        birthday: '',
        address: '',
        postcode:'',
        recommendID:''
      }
    }
  },
  methods: {

    onSubmit() {
      this.$refs.Registerform.validate(async valid => {
        if(!valid) return
        this.$http.post('api/v1/CreateUsers',this.info).then(res => {
          console.log(res)
          this.$message.success('登录成功')
          this.$router.push('/main')
        }).catch(err => {
          console.log(err,"请求出错")
        })
      })
    },
    cancel(){
      this.$router.push('/main');
    }
  }
}
</script>

<style scoped>
.register{
  margin-top: 50px;
}
.res_box{
  width: 480px;
  margin: 0 auto;
  height: 100%;
  padding: 0px 20px 10px 0px;
  background-color: rgb(255, 255, 255);
  border-radius: 32px;
  text-align: center;
  min-height: initial;
  box-shadow: rgb(0 0 0 / 45%) 0px 2px 10px;
}


</style>
