import {TOKEN} from './constants'

const parse = (res) => {
  // res.status === 200 || res.status === 0
  return res.ok ? res.json() : res.text().then(err => { throw err })
}

export const options = (method) => {
  return {
    method: method,
    mode: 'cors',
    credentials: 'include',
    headers: {
      'Authorization': `BEARER ${window.sessionStorage.getItem(TOKEN)}`
    }
  }
}

const api = (path) => `/api${path}`

export const get = (path) => {
  return fetch(api(path), options('get')).then(parse)
}

export const _delete = (path) => {
  return fetch(api(path), options('delete')).then(parse)
}

export const post = (path, body) => {
  var data = options('post')
  data.body = JSON.stringify(body)
  // https://github.github.io/fetch/#options
  return fetch(api(path), data).then(parse)
}
