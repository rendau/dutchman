import { boot } from 'quasar/wrappers'
import axios from 'axios'
import _ from 'lodash'
import { cns } from './cns'
import { ctx as i18nCtx } from './i18n'

const RTO = 20000

let _router = null
let _store = null

function irReqToken (config) {
  if (config.anonym) {
    delete config.anonym
  } else {
    let token = _store.getters['profile/authTokenAccess']
    if (token) {
      config.headers['Authorization'] = 'Bearer ' + token
    }
  }
  return config
}

function irRepSuccess (resp) {
  return resp
}

function irRepError (err) {
  if (!err) {
    err = {}
  }
  if (err.response?.status === 401) {
    err.response.data = { error_code: cns.ErrNotAuthorized }
  }
  return Promise.reject({
    config: err.config,
    request: err.request,
    response: err.response,
    data: transformErr(err.response?.data),
  })
}

function irRepNotAuthorized (err, axiosInstance) {
  if (err.data.code === cns.ErrNotAuthorized) {
    let oConfig = err.config
    let nr401 = oConfig.nr401 === true
    let nfa = oConfig.nfa === true
    let noLogout = oConfig.nl === true

    if (axiosInstance && !nfa) {
      return _store.dispatch('profile/refreshAccessToken').then(aToken => {
        oConfig.nfa = true
        return axiosInstance.request(oConfig)
      }, sErr => {
        if (sErr.response?.status !== 401) {
          console.error('fail to refresh access token', sErr)
        }
        if (!noLogout) {
          return logout(err, nr401).then(() => Promise.reject(err))
        } else {
          return Promise.reject(err)
        }
      })
    } else {
      if (!noLogout) {
        return logout(err, nr401).then(() => Promise.reject(err))
      } else {
        return Promise.reject(err)
      }
    }
  }

  return Promise.reject(err)
}

function logout (err, noRedirect = false) {
  return _store.dispatch('profile/logout').then(() => {
    if (!noRedirect) {
      console.log('not authed redirect')
      _router.push({ name: cns.authRouteName })
    }
  }, () => {})
}

function transformErr (data) { // returns custom (for app) error
  const apiErrAttr = 'error_code'
  if (data && data[apiErrAttr]) {
    let res = _.assign(
      _.omit(data, [apiErrAttr]),
      { code: data[apiErrAttr] },
      { desc: i18nCtx.i18n.global.t(`api_err.${data[apiErrAttr]}`) || i18nCtx.i18n.global.t(cns.ErrSystem) },
    )

    if (res.fields && _.isObject(res.fields)) {
      res.fields = _.mapValues(res.fields, eCode => i18nCtx.i18n.global.t(`api_err.${eCode}`))
    }

    return res
  }
  return {
    code: cns.ErrSystem,
    desc: i18nCtx.i18n.global.t(cns.ErrSystem),
  }
}

export default boot(({ app, store, router }) => {
  app.config.globalProperties.$axios = axios
  store.$axios = axios

  // api
  const api = axios.create({
    baseURL: cns.ApiUrl,
    responseType: 'json',
    timeout: RTO,
  })
  api.interceptors.request.use(irReqToken)
  api.interceptors.response.use(
    irRepSuccess,
    irRepError,
  )
  api.interceptors.response.use(
    undefined,
    err => irRepNotAuthorized(err, api),
  )
  app.config.globalProperties.$api = api
  store.$api = api

  _router = router
  _store = store
})
