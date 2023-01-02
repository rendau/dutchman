export function get (ctx) {
  return this.$api.get('config').then(resp => ctx.commit('assign', resp.data))
}

export function set (ctx, data) {
  return this.$api.put('config', data).then(resp => ctx.commit('assign', resp.data))
}
