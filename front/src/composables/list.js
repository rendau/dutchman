import { ref } from 'vue'
import _ from 'lodash'
import { useStore } from 'vuex'
import { useQuasar } from 'quasar'

export default function (action, params = {}) {
  const store = useStore()
  const $q = useQuasar()

  const loading = ref(false)
  const results = ref([])

  const fetch = async () => {
    loading.value = true
    let pars = _.assign({}, params)
    return store.dispatch(action, pars).then(resp => {
      results.value = resp.data
    }, err => {
      $q.notify({ type: 'negative', message: err.data.desc })
    }).then(() => {
      loading.value = false
    })
  }

  const refresh = async () => {
    return fetch()
  }

  return {
    loading,
    results,
    fetch,
    refresh,
  }
}
