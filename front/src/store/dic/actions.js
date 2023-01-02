export function get (ctx) {
  return Promise.all([
    getDic.call(this, ctx),
  ])
}

export function getDic (ctx) {
  return this.$api.get('dic').then(resp => ctx.commit('assign', resp.data))
}
