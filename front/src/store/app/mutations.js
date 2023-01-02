export function addLoading (state) {
  state.loading += 1
}

export function removeLoading (state) {
  if (state.loading > 0) {
    state.loading -= 1
  }
}

export function setAppError (state, value) {
  state.appError = value
}

export function setDrawer (state, value) {
  state.drawer = value
}

export function toggleDrawer (state) {
  state.drawer = !state.drawer
}
