import constants from './constants'

export function initToken (state) {
  state.token.access = this.$ls.get(constants.LsKeyTokenAccess)
  state.token.refresh = this.$ls.get(constants.LsKeyTokenRefresh)
}

export function setTokenAccess (state, value) {
  state.token.access = value
  this.$ls.set(constants.LsKeyTokenAccess, value || '')
}

export function setTokenRefresh (state, value) {
  state.token.refresh = value
  this.$ls.set(constants.LsKeyTokenRefresh, value || '')
}

export function setToken (state, { access, refresh }) {
  state.token.access = access
  this.$ls.set(constants.LsKeyTokenAccess, access || '')
  state.token.refresh = refresh
  this.$ls.set(constants.LsKeyTokenRefresh, refresh || '')
}

export function setCtx (state, value) {
  state.ctx = value
}
