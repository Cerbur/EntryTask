<template>
  <el-container class="login_container">
    <el-main>
      <div class="avatar_box">
        <img src="../assets/logo.png" alt="">
      </div>

      <div class="login_box">
          <el-form :model="loginForm" ref="loginForm" class="login_form">
            <el-form-item label="" prop="username" :rules="[{ required: true, message: '请填写正确的用户名', trigger: 'blur' }, { required: true, pattern: /.{6,50}/, message: 'username至少为6位', trigger: 'blur' }]">
              <el-input v-model="loginForm.username" placeholder="账号" maxlength='16' prefix-icon="el-icon-user" ></el-input>
            </el-form-item>

            <el-form-item label="" prop="password" :rules="[{ required: true, message: '请填写密码', trigger: 'blur' }, { required: true, pattern: /.{6,50}/ , message: '密码长度至少为6位', trigger: 'blur' }]">
              <el-input v-model="loginForm.password" placeholder="密码" prefix-icon="el-icon-edit" type="password"></el-input>
            </el-form-item>

            <el-form-item class="btns">
              <el-button type="primary" :disabled="loginButton" :loading="loginButton" @click="submitForm">登录</el-button>
            </el-form-item>

          </el-form>
        </div>
    </el-main>
  </el-container>
</template>
<script>
import {loginRequest} from "../utils/api";
import {Message} from "element-ui";
export default {
  data() {
    return {
      labelPosition: "right",
      form: {},
      loginForm: {username: "", password: ""},
      loginButton:false,
    }
  },
  methods: {
    doLogin(phone,password) {
      loginRequest(phone,password).then((res) => {
        if (res.data.code === 1) {
          Message.success("登录成功")
          localStorage.setItem("eToken",res.data.data);
          
          this.$router.replace({ name: "home" });
        }
      })
    },
    submitForm() {
      this.loginButton = true;
      this.$refs.loginForm.validate((v) => {
          if(!v) {
            return false;
          }
          this.doLogin(this.loginForm.username,this.loginForm.password)
      })
      this.loginButton = false;
    },
  }
}
</script>
<style lang="less" scoped>
.login_container {
  background-color: #2b4b6b;
  height: 100vh;
  width: 100vw;
}

.avatar_box {
  height: 100px;
  width:  100px;
  border: 1px solid #eee;
  border-radius: 50%;
  padding: 10px;
  box-shadow: 0 0 10px #ddd;
  position: absolute;
  top: 15%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: #fff;

  img {
    width: 100%;
    height: 100%;
    border-radius: 50%;
    background-color: #eee;
  }
}

.login_box {
  width: 80vw;
  height: 40vh;
  background-color: #fff;
  border-radius: 5px;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);


  //
  //.box-card {
  //  width: 480px;
  //}

  .login_form {
    height: 80vw;
    position: absolute;
    bottom: 20vh;
    top: 10vh;
    width: 100%;
    padding: 0px 20px;
    box-sizing: border-box;
  }
  .btns {
    display: flex;
    justify-content: flex-end;
  }
}

</style>