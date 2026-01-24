<template>
  <AppLayout>
    <div class="blog-detail-container">
      <el-page-header @back="$router.push('/')" content="文章详情" class="page-header"></el-page-header>
      
      <el-card v-loading="loading" class="blog-card" shadow="never">
        <template v-if="blog">
          <h1 class="blog-title">{{ blog.title }}</h1>
          
          <div class="blog-meta">
            <span class="meta-item">
              <i class="el-icon-user"></i> {{ blog.author_name || 'Anonymous' }}
            </span>
            <span class="meta-item">
              <i class="el-icon-time"></i> {{ formatTime(blog.create_time) }}
            </span>
            <el-tag size="mini" type="info" class="meta-tag" v-if="blog.category">{{ blog.category }}</el-tag>
            
            <div class="actions" v-if="isAuthor">
              <el-button type="text" icon="el-icon-delete" class="delete-btn" @click="handleDelete">删除</el-button>
            </div>
          </div>
          
          <el-divider></el-divider>
          
          <div class="blog-content">
            {{ blog.content }}
          </div>
        </template>
        
        <el-empty v-else-if="!loading" description="文章不存在或已删除"></el-empty>
      </el-card>
    </div>
  </AppLayout>
</template>

<script>
import AppLayout from '@/components/layout/AppLayout.vue'
import { getBlogDetail, deleteBlog } from '@/api/blog'
import { getProfile } from '@/api/user'

export default {
  name: 'BlogDetailView',
  components: { AppLayout },
  data() {
    return {
      blog: null,
      loading: true,
      currentUser: null
    }
  },
  computed: {
    isAuthor() {
      return this.blog && this.currentUser && (this.blog.author_id === this.currentUser.id)
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    async fetchData() {
      this.loading = true
      const id = this.$route.params.id
      try {
        const [blogData, userData] = await Promise.all([
          getBlogDetail(id),
          getProfile().catch(() => null)
        ])
        this.blog = blogData
        this.currentUser = userData
      } catch (err) {
        console.error(err)
        this.$message.error('获取博文详情失败')
      } finally {
        this.loading = false
      }
    },
    formatTime(time) {
      if (!time) return ''
      return new Date(time).toLocaleString()
    },
    handleDelete() {
      this.$confirm('此操作将永久删除该博文, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        try {
          await deleteBlog(this.blog.id)
          this.$message.success('删除成功')
          this.$router.push('/')
        } catch (err) {
          this.$message.error('删除失败')
        }
      })
    }
  }
}
</script>

<script>
// Mocking author_name if not present in API for UI demo
</script>

<style scoped>
.blog-detail-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px 0;
}

.page-header {
  margin-bottom: 24px;
}

.blog-card {
  border-radius: 8px;
  padding: 20px;
}

.blog-title {
  font-size: 32px;
  font-weight: 700;
  color: #303133;
  margin-bottom: 20px;
  line-height: 1.4;
}

.blog-meta {
  display: flex;
  align-items: center;
  color: #909399;
  font-size: 14px;
  margin-bottom: 10px;
}

.meta-item {
  margin-right: 20px;
  display: flex;
  align-items: center;
}

.meta-item i {
  margin-right: 4px;
}

.meta-tag {
  margin-right: 10px;
}

.actions {
  margin-left: auto;
}

.delete-btn {
  color: #F56C6C;
}

.blog-content {
  font-size: 16px;
  line-height: 1.8;
  color: #303133;
  white-space: pre-wrap;
  min-height: 300px;
}

@media (max-width: 768px) {
  .blog-title {
    font-size: 24px;
  }
  
  .blog-meta {
    flex-wrap: wrap;
  }
  
  .meta-item {
    margin-bottom: 10px;
  }
}
</style>
