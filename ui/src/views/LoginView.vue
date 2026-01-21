<template>
  <div class="login-container">
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>Login</span>
      </div>
      <el-form ref="loginForm" :model="loginForm" :rules="loginRules" label-width="80px">
        <el-form-item label="Username" prop="username">
          <el-input v-model="loginForm.username" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="Password" prop="password">
          <el-input type="password" v-model="loginForm.password" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin">Login</el-button>
          <el-button @click="$router.push('/register')">Register</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { login } from '@/api/auth'
import { setToken } from '@/utils/auth'

export default {
  name: 'LoginView',
  data() {
    return {
      loginForm: {
        username: '',
        password: ''
      },
      loginRules: {
        username: [{ required: true, message: 'Please input username', trigger: 'blur' }],
        password: [{ required: true, message: 'Please input password', trigger: 'blur' }]
      }
    }
  },
  methods: {
    handleLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          login(this.loginForm).then(token => {
            setToken(token)
            this.$message.success('Login successful')
            this.$router.push('/')
          }).catch(() => {})
        } else {
          return false
        }
      })
    }
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f0f2f5;
}
.box-card {
  width: 400px;
}
</style>
