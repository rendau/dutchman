let _appStartPrResolver = null
let _appStartPr = new Promise(resolve => {
  _appStartPrResolver = resolve
})

export function appStartPr () {
  return _appStartPr
}

export function resolveAppStartPr () { // must be called by App.vue
  if (_appStartPrResolver) {
    _appStartPrResolver()
    _appStartPrResolver = null
  }
}
