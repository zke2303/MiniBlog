import request from '@/utils/request'

export function createBlog(data) {
  return request({
    url: '/blogs/',
    method: 'post',
    data
  })
}

export function listBlogs(params) {
  return request({
    url: '/blogs',
    method: 'get',
    params
  })
}

export function getBlogDetail(id) {
  return request({
    url: `/blogs/${id}`,
    method: 'get'
  })
}

export function deleteBlog(id) {
  return request({
    url: `/blogs/${id}`,
    method: 'delete'
  })
}
