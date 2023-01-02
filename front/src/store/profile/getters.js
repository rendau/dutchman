export function authTokenAccess (state) {
  return state.token.access
}

export function authTokenRefresh (state) {
  return state.token.refresh
}

export function isAuthed (state, getters) {
  return !!state.ctx
}

export function id (state) {
  return (state.ctx || {}).id || 0
}
