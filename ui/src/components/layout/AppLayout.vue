<template>
  <el-container class="app-wrapper">
    <el-header height="60px" class="app-header">
      <div class="header-content">
        <!-- Brand -->
        <div class="brand">
          <span class="logo">MiniBlog</span>
        </div>

        <!-- Search -->
        <div class="search-section">
          <el-input
            v-model="input"
            placeholder="搜索文章..."
            prefix-icon="el-icon-search"
            size="small"
            class="search-input"
          ></el-input>
        </div>

        <!-- Actions -->
        <div class="actions">
          <el-tooltip content="文章列表" placement="bottom">
            <i class="el-icon-s-order action-icon"></i>
          </el-tooltip>
          <el-tooltip content="消息通知" placement="bottom">
            <i class="el-icon-message-solid action-icon"></i>
          </el-tooltip>
          <el-dropdown trigger="click" @command="handleCommand">
            <span class="el-dropdown-link">
              <el-avatar size="small" :src="avatarUrl"></el-avatar>
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item command="profile">个人中心</el-dropdown-item>
              <el-dropdown-item command="create">写博文</el-dropdown-item>
              <el-dropdown-item divided command="logout">退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </div>
      </div>
    </el-header>

    <el-main class="app-main">
      <div class="main-container">
        <!-- Content Area -->
         <slot></slot>
      </div>
    </el-main>
  </el-container>
</template>

<script>
import { removeToken } from '@/utils/auth'

export default {
  name: 'AppLayout',
  data() {
    return {
      input: '',
      avatarUrl: 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'
    }
  },
  methods: {
    handleCommand(command) {
      if (command === 'logout') {
        removeToken()
        this.$router.push('/login')
      } else if (command === 'create') {
        this.$router.push('/create-blog')
      }
    }
  }
}
</script>

<style scoped>
.app-wrapper {
  min-height: 100vh;
  background-color: #f5f7fa;
}

.app-header {
  background-color: #ffffff;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.05);
  position: sticky;
  top: 0;
  z-index: 100;
  padding: 0;
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.brand .logo {
  font-size: 24px;
  font-weight: bold;
  color: #409EFF;
  cursor: pointer;
}

.search-section {
  flex: 1;
  max-width: 400px;
  margin: 0 40px;
}

.actions {
  display: flex;
  align-items: center;
}

/* Use margin for spacing instead of gap for better compatibility */
.action-icon {
  font-size: 20px;
  color: #606266;
  cursor: pointer;
  transition: color 0.3s;
  margin-right: 24px;
}

.action-icon:last-child {
  margin-right: 0;
}

.action-icon:hover {
  color: #409EFF;
}

.el-dropdown-link {
  cursor: pointer;
  display: flex;
  align-items: center;
}

.app-main {
  padding: 20px;
  display: flex;
  justify-content: center;
}

.main-container {
  width: 100%;
  max-width: 1200px;
}
</style>
