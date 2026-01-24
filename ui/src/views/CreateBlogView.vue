<template>
  <AppLayout>
    <div class="create-blog-container">
      <el-page-header @back="$router.push('/')" content="发布新博文" class="page-header"></el-page-header>
      
      <el-card class="form-card" shadow="never">
        <el-form ref="blogForm" :model="blogForm" :rules="blogRules" label-position="top">
          <el-form-item label="文章标题" prop="title">
            <el-input 
              v-model="blogForm.title" 
              placeholder="请输入标题 (1-30个字符)"
              maxlength="30"
              show-word-limit
            ></el-input>
          </el-form-item>
          
          <el-form-item label="正文内容" prop="content">
            <el-input 
              type="textarea" 
              v-model="blogForm.content" 
              :rows="15"
              placeholder="撰写你的故事... (最少20个字符)"
            ></el-input>
          </el-form-item>
          
          <el-form-item class="form-actions">
            <el-button type="primary" :loading="submitting" @click="handleCreate" icon="el-icon-check">立即发布</el-button>
            <el-button @click="$router.push('/')">取消</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </AppLayout>
</template>

<script>
import AppLayout from '@/components/layout/AppLayout.vue'
import { createBlog } from '@/api/blog'

export default {
  name: 'CreateBlogView',
  components: { AppLayout },
  data() {
    return {
      submitting: false,
      blogForm: {
        title: '',
        content: ''
      },
      blogRules: {
        title: [
          { required: true, message: '请输入文章标题', trigger: 'blur' },
          { min: 1, max: 30, message: '长度在 1 到 30 个字符', trigger: 'blur' }
        ],
        content: [
          { required: true, message: '请输入文章内容', trigger: 'blur' },
          { min: 20, message: '内容最少需要 20 个字符', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    handleCreate() {
      this.$refs.blogForm.validate(valid => {
        if (valid) {
          this.submitting = true
          createBlog(this.blogForm).then(() => {
            this.$message.success('博文发布成功！')
            this.$router.push('/')
          }).catch(err => {
            console.error(err)
            this.$message.error('发布失败，请稍后重试')
          }).finally(() => {
            this.submitting = false
          })
        }
      })
    }
  }
}
</script>

<style scoped>
.create-blog-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px 0;
}

.page-header {
  margin-bottom: 24px;
}

.form-card {
  border-radius: 8px;
  padding: 10px 20px;
}

.form-actions {
  margin-top: 30px;
  border-top: 1px solid #EBEEF2;
  padding-top: 20px;
}
</style>
