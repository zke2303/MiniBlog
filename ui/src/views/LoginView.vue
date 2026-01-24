<template>
  <div class="login-container">
    <div class="login-box">
      <div class="login-header">
        <h1 class="brand-logo">MiniBlog</h1>
        <p class="brand-slogan">记录生活，分享快乐</p>
      </div>
      
      <el-card class="login-card" shadow="always">
        <h2 class="form-title">欢迎回来</h2>
        <el-form ref="loginForm" :model="loginForm" :rules="loginRules">
          <el-form-item prop="username">
            <el-input 
              v-model="loginForm.username" 
              prefix-icon="el-icon-user" 
              placeholder="用户名"
              autocomplete="off"
            ></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input 
              type="password" 
              v-model="loginForm.password" 
              prefix-icon="el-icon-lock" 
              placeholder="密码"
              autocomplete="off"
              @keyup.enter.native="handleLogin"
            ></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" :loading="loading" class="login-submit" @click="handleLogin">登录</el-button>
          </el-form-item>
          <div class="form-footer">
            <span>还没有账号? </span>
            <el-link type="primary" :underline="false" @click="$router.push('/register')">立即注册</el-link>
          </div>
        </el-form>
      </el-card>
    </div>
  </div>
</template>

<script>
import { login } from '@/api/auth'
import { setToken } from '@/utils/auth'

export default {
  name: 'LoginView',
  data() {
    return {
      loading: false,
      loginForm: {
        username: '',
        password: ''
      },
      loginRules: {
        username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
        password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
      }
    }
  },
  methods: {
    handleLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true
          login(this.loginForm).then(token => {
            setToken(token)
            this.$message.success('登录成功')
            this.$router.push('/')
          }).catch(() => {
            // Error handled by interceptor
          }).finally(() => {
            this.loading = false
          })
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
  min-height: 100vh;
  background-color: #f5f7fa;
  background-image: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
}

.login-box {
  width: 100%;
  max-width: 400px;
  padding: 20px;
}

.login-header {
  text-align: center;
  margin-bottom: 30px;
}

.brand-logo {
  font-size: 42px;
  color: #409EFF;
  margin: 0;
  font-weight: bold;
  letter-spacing: 2px;
}

.brand-slogan {
  color: #909399;
  margin-top: 8px;
  font-size: 16px;
}

.login-card {
  border-radius: 12px;
  border: none;
}

.form-title {
  text-align: center;
  margin-bottom: 24px;
  color: #303133;
  font-weight: 500;
}

.login-submit {
  width: 100%;
  padding: 12px;
  font-size: 16px;
}

.form-footer {
  text-align: center;
  font-size: 14px;
  color: #606266;
  margin-top: 16px;
}
</style>
