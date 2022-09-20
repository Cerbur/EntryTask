<template>
  <el-container class="login_container">
    <el-main>
      <el-dialog
        title="更新头像"
        :visible.sync="dialogVisible"
        width="30%">
          <el-upload  ref="upload" action="" :multiple="false" :pic-width="250" :pic-height="90"
              :auto-upload="false" :limit="1" :http-request="requestFile" >
            <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
          </el-upload>
          <span slot="footer" class="dialog-footer">
            <el-button @click="dialogVisible = false">取 消</el-button>
            <el-button type="primary" :disabled="updateButton" :loading="updateButton" @click="clickImgUpdate">确 定</el-button>
          </span>
      </el-dialog>
      <div class="avatar_box">
        <img :src="loginForm.picture" alt="" @click="clickImg">
      </div>
      <div class="login_box">

          <el-form :model="loginForm" ref="loginForm" class="login_form">
              username: <el-input v-model="loginForm.username" placeholder="账号" maxlength='20' prefix-icon="el-icon-user" :disabled="true"></el-input>
              nickname: <el-input v-model="loginForm.nickname" placeholder="nickname" minlength="6" prefix-icon="el-icon-edit"></el-input>
            <el-form-item class="btns">
              <el-button type="primary" :disabled="loginButton" :loading="loginButton" @click="putNickname">更新</el-button>
            </el-form-item>

          </el-form>
      </div>
    </el-main>
  </el-container>
</template>
<script>
import {getProfileRequest} from "../utils/api";
import {putProfileNicknameRequest} from "../utils/api";
import {Message} from "element-ui";
import axios from "../utils/http";
export default {
  data() {
    return {
      labelPosition: "right",
      form: {},
      loginForm: {username: "", nickname: "",picture:""},
      file:'',
      loginButton:false,
      dialogVisible:false,
      updateButton:false,
    }
  },
  
  mounted() {
    if (!localStorage.getItem("eToken")) {
      this.$router.replace({name:"login"});
    }
    console.log()
    this.getProfile();
  },
  methods: {
    requestFile(param) {
        let form = new FormData()    // FormData 对象
        form.append('file', param.file)    // 文件对象
        let config = {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        }
        let url = `/api/profile/picture`
        axios.put(url,form,config).then((res)=>{
            if (res.data.code == 1) {
              let data = res.data.data
              this.loginForm.picture = "http://127.0.0.1:8888/api/img?img="+data.profile_picture
              Message.success("更换成功")
            }
        })
    },

    clickImgUpdate(){
      this.updateButton = true
      this.$refs.upload.submit();
      this.updateButton = false
      this.dialogVisible = false
    },
    clickImg() {
      this.dialogVisible = true
    },
    getProfile() {
      getProfileRequest().then((res) =>{
        if (res.data.code === 1) {
          let data = res.data.data
          Message.success("获取成功")
          this.loginForm.username = data.username
          this.loginForm.nickname = data.nickname
          this.loginForm.picture = "http://127.0.0.1:8888/api/img?img="+data.profile_picture
          console.log(data)
        } else {
          Message.error("获取失败")
          this.$router.replace({name: "login"});
        }
      })
    },
    putNickname() {
      putProfileNicknameRequest(this.loginForm.nickname).then((res) =>{
        if (res.data.code === 1) {
          let data = res.data.data
          Message.success("更新成功")
          this.loginForm.nickname = data.nickname
        } else {
          Message.error("更新失败")
          this.$router.replace({name: "login"});
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