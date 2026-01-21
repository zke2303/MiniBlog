import request from '@/utils/request'

export function createBlog(data) {
  return request({
    url: '/blogs/',
    method: 'post',
    data
  })
}
