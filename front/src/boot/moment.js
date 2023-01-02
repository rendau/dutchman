import { boot } from 'quasar/wrappers'
import moment from 'moment'
import 'moment/locale/ru'

export default boot(async ({ app }) => {
  app.config.globalProperties.$m = moment
})
