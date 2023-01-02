import { boot } from 'quasar/wrappers'
import { createI18n } from 'vue-i18n/index'
import messages from 'src/i18n'
import { cns } from './cns'

const ctx = {
  i18n: null,
}

export default boot(({ app }) => {
  ctx.i18n = createI18n({
    locale: cns.DefaultLocale,
    fallbackLocale: cns.DefaultLocale,
    messages,
  })

  // Set i18n instance on app
  app.use(ctx.i18n)
})

export { ctx }
