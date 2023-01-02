import { boot } from 'quasar/wrappers'
import { cns } from './cns'

export default boot(async ({ store, router }) => {
  router.beforeEach((to, from, next) => {
    // if (process.env.NODE_ENV === 'production') {
    //   console.log('router beforeEach', from, to.fullPath)
    // }
    store.dispatch('app/appStartPr').then(() => {
      return store.dispatch('profile/refreshSinceAppStart')
    }).then(() => {
      if (to.matched.some(r => r.meta.requiresAuth)) {
        if (store.getters['profile/isAuthed']) {
          next()
        } else {
          next({ name: cns.authRouteName })
        }
      } else if (to.matched.some(r => r.meta.restrictAuth)) {
        if (store.getters['profile/isAuthed']) {
          next('/')
        } else {
          next()
        }
      } else {
        next()
      }
    }, () => {
      next()
    })
  })
})
