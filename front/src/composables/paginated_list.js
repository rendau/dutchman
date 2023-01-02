import { computed, ref } from 'vue'
import { cns } from 'boot/cns'
import _ from 'lodash'
import { useStore } from 'vuex'
import { useQuasar } from 'quasar'

export default function (action, params = ref({})) {
  const store = useStore()
  const $q = useQuasar()

  const loading = ref(false)
  const paginationParams = ref({
    page: 0,
    page_size: cns.DefaultPageSize,
    with_total_count: true,
  })
  const data = ref({
    total_count: 0,
    results: [],
  })

  const results = computed(() => (data.value?.results || []))

  const fetch = async () => {
    loading.value = true
    let pars = _.assign({}, params.value, paginationParams.value)
    return store.dispatch(action, pars).then(resp => {
      data.value = resp.data
    }, err => {
      $q.notify({ type: 'negative', message: err.data.desc })
    }).then(() => {
      loading.value = false
    })
  }

  const refresh = async (restart = false) => {
    if (restart) {
      paginationParams.value.page = 0
    }
    return fetch()
  }

  const onPage = page => {
    paginationParams.value.page = page
    fetch().catch(() => {})
  }

  return {
    loading,
    paginationParams,
    data,
    results,
    fetch,
    refresh,
    onPage,
  }
}
