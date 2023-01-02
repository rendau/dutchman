import { boot } from 'quasar/wrappers'

const cns = {
  ApiUrl: 'https://google.kz',
  AccountApiUrl: 'https://google.kz',
  AccountAuthUrl: 'https://google.kz',
  AccountCabinetUrl: 'https://google.kz',

  LocaleRu: 'ru-ru',

  DefaultLocale: 'ru-ru',

  StorageTokenKey: 'auth.token',

  ErrSystem: 'system_error',
  ServiceNotAvailable: 'service_not_available',
  ErrNotAuthorized: 'not_authorized',
  PermissionDenied: 'permission_denied',
  ErrFormValidate: 'form_validate',

  authRouteName: 'auth',
}

cns.ApiUrl = window.API_URL || cns.ApiUrl
cns.AccountApiUrl = window.ACCOUNT_API_URL || cns.AccountApiUrl
cns.AccountAuthUrl = window.ACCOUNT_AUTH_URL || cns.AccountAuthUrl
cns.AccountCabinetUrl = window.ACCOUNT_CABINET_URL || cns.AccountCabinetUrl

if (process.env.API_URL) {
  cns.ApiUrl = process.env.API_URL
}

if (process.env.ACCOUNT_API_URL) {
  cns.AccountApiUrl = process.env.ACCOUNT_API_URL
}

if (process.env.ACCOUNT_AUTH_URL) {
  cns.AccountAuthUrl = process.env.ACCOUNT_AUTH_URL
}

if (process.env.ACCOUNT_CABINET_URL) {
  cns.AccountCabinetUrl = process.env.ACCOUNT_CABINET_URL
}

export default boot(async ({ app, store }) => {
  console.log('ApiUrl:', cns.ApiUrl)
  console.log('AccountApiUrl:', cns.AccountApiUrl)
  console.log('AccountAuthUrl:', cns.AccountAuthUrl)
  console.log('AccountCabinetUrl:', cns.AccountCabinetUrl)
  app.config.globalProperties.$cns = cns
  store.$cns = cns
})

export { cns }
