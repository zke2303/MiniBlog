<template>
  <div class="create-blog-container">
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>Create New Blog Post</span>
        <el-button style="float: right; padding: 3px 0" type="text" @click="$router.push('/')">Cancel</el-button>
      </div>
      <el-form ref="blogForm" :model="blogForm" :rules="blogRules" label-width="80px">
        <el-form-item label="Title" prop="title">
          <el-input v-model="blogForm.title" placeholder="Enter title (1-30 chars)"></el-input>
        </el-form-item>
        <el-form-item label="Content" prop="content">
          <el-input 
            type="textarea" 
            v-model="blogForm.content" 
            :rows="10"
            placeholder="Enter content (min 20 chars)"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleCreate">Publish</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { createBlog } from '@/api/blog'

export default {
  name: 'CreateBlogView',
  data() {
    return {
      blogForm: {
        title: '',
        content: ''
      },
      blogRules: {
        title: [
          { required: true, message: 'Please input title', trigger: 'blur' },
          { min: 1, max: 30, message: 'Length should be 1 to 30', trigger: 'blur' }
        ],
        content: [
          { required: true, message: 'Please input content', trigger: 'blur' },
          { min: 20, message: 'Length should be at least 20', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    handleCreate() {
      this.$refs.blogForm.validate(valid => {
        if (valid) {
          createBlog(this.blogForm).then(() => {
            this.$message.success('Blog published successfully')
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
.create-blog-container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}
</style>
