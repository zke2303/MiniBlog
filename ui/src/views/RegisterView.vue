<template>
  <div class="register-container">
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>Register</span>
      </div>
      <el-form ref="registerForm" :model="registerForm" :rules="registerRules" label-width="80px">
        <el-form-item label="Username" prop="username">
          <el-input v-model="registerForm.username" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="Password" prop="password">
          <el-input type="password" v-model="registerForm.password" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleRegister">Register</el-button>
          <el-button @click="$router.push('/login')">Back to Login</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { register } from '@/api/auth'

export default {
  name: 'RegisterView',
  data() {
    return {
      registerForm: {
        username: '',
        password: ''
      },
      registerRules: {
        username: [{ required: true, message: 'Please input username', trigger: 'blur' }],
        password: [{ required: true, message: 'Please input password', trigger: 'blur' }]
      }
    }
  },
  methods: {
    handleRegister() {
      this.$refs.registerForm.validate(valid => {
        if (valid) {
          register(this.registerForm).then(() => {
            this.$message.success('Registration successful. Please login.')
            this.$router.push('/login')
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
.register-container {
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
