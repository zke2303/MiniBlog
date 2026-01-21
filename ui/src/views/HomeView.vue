<template>
  <AppLayout>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>Dashboard</span>
        <el-button style="float: right; padding: 3px 0" type="text" @click="$router.push('/create-blog')">Create Blog</el-button>
      </div>
      <div v-if="user" class="text item">
        <p><strong>Username:</strong> {{ user.username }}</p>
        <p><strong>ID:</strong> {{ user.id }}</p>
        <p><strong>Joined:</strong> {{ formatTime(user.create_time) }}</p>
      </div>
      <div v-else>
        Loading profile...
      </div>
    </el-card>
  </AppLayout>
</template>

<script>
import AppLayout from '@/components/layout/AppLayout.vue'
import { getProfile } from '@/api/user'

export default {
  name: 'HomeView',
  components: { AppLayout },
  data() {
    return {
      user: null
    }
  },
  created() {
    this.fetchProfile()
  },
  methods: {
    fetchProfile() {
      getProfile().then(data => {
        this.user = data
      }).catch(err => {
        console.error(err)
      })
    },
    formatTime(time) {
      return new Date(time).toLocaleString()
    }
  }
}
</script>

<style scoped>
.item {
  margin-bottom: 18px;
}
</style>
