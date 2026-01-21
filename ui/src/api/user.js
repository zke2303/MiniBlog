import request from '@/utils/request'

export function getProfile() {
  return request({
    url: '/users/profile',
    method: 'get'
  })
}
