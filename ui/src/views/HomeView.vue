<template>
  <AppLayout>
    <div class="home-container">
      <el-row :gutter="20">
        <!-- Main Content: Blog List -->
        <el-col :xs="24" :sm="24" :md="17" :lg="18">
          <div v-loading="loading" class="blog-list">
            <template v-if="blogs && blogs.length > 0">
              <el-card 
                v-for="blog in blogs" 
                :key="blog.id" 
                class="blog-item-card" 
                shadow="hover"
                @click.native="$router.push(`/blog/${blog.id}`)"
              >
                <div class="blog-item-content">
                  <h2 class="blog-title">{{ blog.title }}</h2>
                  <p class="blog-summary">{{ getSummary(blog.content) }}</p>
                  <div class="blog-footer">
                    <div class="meta-info">
                      <span class="meta-item">
                        <i class="el-icon-user"></i> {{ blog.author_name || 'Anonymous' }}
                      </span>
                      <span class="meta-item">
                        <i class="el-icon-time"></i> {{ formatTime(blog.create_time) }}
                      </span>
                    </div>
                    <el-button type="text" class="read-more">阅读全文 <i class="el-icon-arrow-right"></i></el-button>
                  </div>
                </div>
              </el-card>

              <!-- Pagination -->
              <div class="pagination-container">
                <el-pagination
                  background
                  layout="prev, pager, next"
                  :total="total"
                  :page-size="query.pageSize"
                  :current-page.sync="query.page"
                  @current-change="fetchBlogs"
                >
                </el-pagination>
              </div>
            </template>
            
            <el-empty v-else-if="!loading" description="暂无文章"></el-empty>
          </div>
        </el-col>

        <!-- Sidebar -->
        <el-col :xs="24" :sm="24" :md="7" :lg="6" class="hidden-sm-and-down">
          <el-card class="profile-card" shadow="never">
            <div slot="header" class="clearfix">
              <span>个人中心</span>
            </div>
            <div v-if="user" class="user-info">
              <div class="user-header">
                <el-avatar :size="64" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
                <h3>{{ user.username }}</h3>
                <p class="user-bio">记录生活，分享快乐。</p>
              </div>
              <el-divider></el-divider>
              <div class="user-stats">
                <div class="stat-item">
                  <div class="stat-value">{{ total }}</div>
                  <div class="stat-label">文章</div>
                </div>
                <div class="stat-item">
                  <div class="stat-value">0</div>
                  <div class="stat-label">粉丝</div>
                </div>
                <div class="stat-item">
                  <div class="stat-value">0</div>
                  <div class="stat-label">获赞</div>
                </div>
              </div>
              <el-button type="primary" style="width: 100%; margin-top: 20px" icon="el-icon-edit" @click="$router.push('/create-blog')">
                发布新博文
              </el-button>
            </div>
            <div v-else class="login-prompt">
              <p>登录以查看更多个人信息</p>
              <el-button type="primary" size="small" @click="$router.push('/login')">去登录</el-button>
            </div>
          </el-card>

          <el-card class="tags-card" shadow="never" style="margin-top: 20px">
            <div slot="header" class="clearfix">
              <span>热门话题</span>
            </div>
            <div class="tags-cloud">
              <el-tag size="small" type="info">Golang</el-tag>
              <el-tag size="small" type="info">Vue.js</el-tag>
              <el-tag size="small" type="info">Backend</el-tag>
              <el-tag size="small" type="info">Frontend</el-tag>
              <el-tag size="small" type="info">Database</el-tag>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </AppLayout>
</template>

<script>
import AppLayout from '@/components/layout/AppLayout.vue'
import { getProfile } from '@/api/user'
import { listBlogs } from '@/api/blog'

export default {
  name: 'HomeView',
  components: { AppLayout },
  data() {
    return {
      user: null,
      blogs: [],
      total: 0,
      loading: false,
      query: {
        page: 1,
        pageSize: 10,
        keyword: ''
      }
    }
  },
  created() {
    this.fetchProfile()
    this.fetchBlogs()
  },
  methods: {
    fetchProfile() {
      getProfile().then(data => {
        this.user = data
      }).catch(err => {
        console.error('Failed to fetch profile:', err)
      })
    },
    fetchBlogs() {
      this.loading = true
      listBlogs(this.query).then(res => {
        // Backend might return directly the list or an object with total
        if (Array.isArray(res)) {
          this.blogs = res
          this.total = res.length
        } else if (res && res.list) {
          this.blogs = res.list
          this.total = res.total
        }
      }).catch(err => {
        console.error('Failed to fetch blogs:', err)
        this.$message.error('获取博文列表失败')
      }).finally(() => {
        this.loading = false
      })
    },
    formatTime(time) {
      if (!time) return ''
      const date = new Date(time)
      return `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()}`
    },
    getSummary(content) {
      if (!content) return ''
      return content.length > 120 ? content.substring(0, 120) + '...' : content
    }
  }
}
</script>

<style scoped>
.home-container {
  padding: 10px 0;
}

.blog-item-card {
  margin-bottom: 16px;
  cursor: pointer;
  border-radius: 8px;
  border: none;
  transition: all 0.3s;
}

.blog-item-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0,0,0,0.1) !important;
}

.blog-title {
  margin: 0 0 12px 0;
  font-size: 20px;
  color: #303133;
  font-weight: 600;
}

.blog-summary {
  color: #606266;
  font-size: 14px;
  line-height: 1.6;
  margin-bottom: 16px;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 3;
  overflow: hidden;
}

.blog-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.meta-info {
  display: flex;
  color: #909399;
  font-size: 13px;
}

.meta-item {
  margin-right: 16px;
  display: flex;
  align-items: center;
}

.meta-item i {
  margin-right: 4px;
}

.read-more {
  padding: 0;
  font-size: 14px;
}

.pagination-container {
  margin-top: 30px;
  display: flex;
  justify-content: center;
}

/* Sidebar styles */
.profile-card {
  border-radius: 8px;
  text-align: center;
}

.user-header h3 {
  margin: 12px 0 4px 0;
  font-size: 18px;
}

.user-bio {
  color: #909399;
  font-size: 13px;
  margin-bottom: 0;
}

.user-stats {
  display: flex;
  justify-content: space-around;
}

.stat-item {
  text-align: center;
}

.stat-value {
  font-weight: bold;
  font-size: 18px;
  color: #303133;
}

.stat-label {
  font-size: 12px;
  color: #909399;
}

.tags-cloud {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.login-prompt {
  padding: 20px 0;
}
</style>
