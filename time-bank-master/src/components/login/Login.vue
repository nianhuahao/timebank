<template>
  <div class="login_container">
   <div class="box1">
     <h2>欢迎登入!</h2>
   </div>
    <div class="login_box">
         <el-form ref="loginFormRef" :model="loginForm" :rules="loginFormRules" label-width="0px" class="login_form">
           <!-- 用户名 -->
           <el-form-item prop="username" >
             <span>用户名:</span>
             <el-input v-model="loginForm.username" prefix-icon="iconfont icon-user" placeholder="帐号" style="width: 250px;"></el-input>
           </el-form-item>
           <!-- 密码 -->
           <el-form-item prop="password">
             <span>密码:</span>
             <el-input v-model="loginForm.password" prefix-icon="iconfont icon-3702mima" type="password" style="width: 250px;float: right;"></el-input>
           </el-form-item>
           <!-- 按钮区域 -->
           <el-form-item class="btns">
             <el-button type="primary" @click="login">登录</el-button>
             <el-button type="info" @click="resetLoginForm">重置</el-button>
           </el-form-item>
         </el-form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      // 这是登录表单的数据绑定对象
      loginForm: {
        username: 'admin',
        password: '123456'
      },
      // 这是表单的验证规则对象
      loginFormRules: {
        // 验证用户名是否合法
        username: [
          { required: true, message: '请输入登录名称', trigger: 'blur' },
          { min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' }
        ],
        // 验证密码是否合法
        password: [
          { required: true, message: '请输入登录密码', trigger: 'blur' },
          { min: 6, max: 15, message: '长度在 6 到 15 个字符', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    // 点击重置按钮，重置登录表单
    resetLoginForm() {
      // console.log(this);
      this.$refs.loginFormRef.resetFields()
    },
    login() {
      this.$refs.loginFormRef.validate(async valid => {
        // if (!valid) return
        // const { data: res } = await this.$http.post('login', this.loginForm)
        // if (res.meta.status !== 200) return this.$message.error('登录失败！')
        // this.$message.success('登录成功')
        // // 1. 将登录成功之后的 token，保存到客户端的 sessionStorage 中
        // //   1.1 项目中出了登录之外的其他API接口，必须在登录之后才能访问
        // //   1.2 token 只应在当前网站打开期间生效，所以将 token 保存在 sessionStorage 中
        // window.sessionStorage.setItem('token', res.data.token)
        // 2. 通过编程式导航跳转到后台主页，路由地址是 /home
        this.$router.push('/main')
      })
    }
  }
}
</script>

<style lang="less" scoped>
.login_container {
  height: 100%;
}
.box1{
  width: 1100px;
  height: 83px;
  padding: 15px 0px;
  position: absolute;
  left: 50%;
  transform: translate(-50%);
}
.login_box{
  width: 100%;
  height: 550px;
  position: absolute;
  top:113px;
  background-image: url("https://www.nstimebank.com/timebank/images/dl_bj.jpg");
}
.login_form{
  position: absolute;

  left: 800px;
  top:100px;
  background-color: #ffffff;
  border-radius: 10px;
  padding-top: 10px;
  width: 300px;
  height: 200px;
}
.btns{
  display: flex;
  justify-content: center;
}
</style>
