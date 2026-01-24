<template>
  <div class="register-container">
    <div class="register-box">
      <div class="register-header">
        <h1 class="brand-logo">MiniBlog</h1>
        <p class="brand-slogan">加入我们，开启创作之旅</p>
      </div>
      
      <el-card class="register-card" shadow="always">
        <h2 class="form-title">创建账号</h2>
        <el-form ref="registerForm" :model="registerForm" :rules="registerRules">
          <el-form-item prop="username">
            <el-input 
              v-model="registerForm.username" 
              prefix-icon="el-icon-user" 
              placeholder="用户名"
              autocomplete="off"
            ></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input 
              type="password" 
              v-model="registerForm.password" 
              prefix-icon="el-icon-lock" 
              placeholder="密码"
              autocomplete="off"
            ></el-input>
          </el-form-item>
          <el-form-item prop="confirmPassword">
            <el-input 
              type="password" 
              v-model="registerForm.confirmPassword" 
              prefix-icon="el-icon-circle-check" 
              placeholder="确认密码"
              autocomplete="off"
              @keyup.enter.native="handleRegister"
            ></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" :loading="loading" class="register-submit" @click="handleRegister">注册</el-button>
          </el-form-item>
          <div class="form-footer">
            <span>已有账号? </span>
            <el-link type="primary" :underline="false" @click="$router.push('/login')">返回登录</el-link>
          </div>
        </el-form>
      </el-card>
    </div>
  </div>
</template>

<script>
import { register } from '@/api/auth'

export default {
  name: 'RegisterView',
  data() {
    const validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.registerForm.password) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      loading: false,
      registerForm: {
        username: '',
        password: '',
        confirmPassword: ''
      },
      registerRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, message: '密码长度不能小于 6 位', trigger: 'blur' }
        ],
        confirmPassword: [
          { required: true, validator: validatePass2, trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    handleRegister() {
      this.$refs.registerForm.validate(valid => {
        if (valid) {
          this.loading = true
          register({
            username: this.registerForm.username,
            password: this.registerForm.password
          }).then(() => {
            this.$message.success('注册成功，请登录')
            this.$router.push('/login')
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
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f7fa;
  background-image: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
}

.register-box {
  width: 100%;
  max-width: 400px;
  padding: 20px;
}

.register-header {
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

.register-card {
  border-radius: 12px;
  border: none;
}

.form-title {
  text-align: center;
  margin-bottom: 24px;
  color: #303133;
  font-weight: 500;
}

.register-submit {
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
