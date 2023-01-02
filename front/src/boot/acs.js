import Spn from 'components/acs/Spn'
import Spinner from 'components/acs/Spinner'
import Icn from 'components/acs/Icn'
import AdaptivePage from 'components/acs/AdaptivePage'
import InputDate from 'components/acs/InputDate'
import InputPeriod from 'components/acs/InputPeriod'
import PageTitle from 'components/acs/PageTitle'
import PageToolbar from 'components/acs/PageToolbar'
import TablePagination from 'components/acs/TablePagination'
import TableTools from 'components/acs/TableTools'
import TableTotalCounter from 'components/acs/TableTotalCounter'
import TrNoRows from 'components/acs/TrNoRows'

export default ({ app }) => {
  app.component('ac-spn', Spn)
  app.component('ac-spinner', Spinner)
  app.component('ac-icn', Icn)
  app.component('ac-adaptive-page', AdaptivePage)
  app.component('ac-input-date', InputDate)
  app.component('ac-input-period', InputPeriod)
  app.component('ac-page-title', PageTitle)
  app.component('ac-page-toolbar', PageToolbar)
  app.component('ac-table-pagination', TablePagination)
  app.component('ac-table-tools', TableTools)
  app.component('ac-table-total-counter', TableTotalCounter)
  app.component('ac-tr-no-rows', TrNoRows)

}
