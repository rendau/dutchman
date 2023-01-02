import { boot } from 'quasar/wrappers'
import _ from 'lodash'

export default boot(async ({ app }) => {
  app.config.globalProperties.$_ = _
})
