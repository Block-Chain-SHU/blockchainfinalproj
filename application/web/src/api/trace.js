import request from '@/utils/request'

export function uplink(data) {
  return request({
    url: '/uplink',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getTraceInfo
export function getTraceInfo(data) {
  return request({
    url: '/getTraceInfo',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getTraceList
export function getTraceList(data) {
  return request({
    url: '/getTraceList',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getAllTraceInfo
export function getAllTraceInfo(data) {
  return request({
    url: '/getAllTraceInfo',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getTraceHistory
export function getTraceHistory(data) {
  return request({
    url: '/getTraceHistory',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}
