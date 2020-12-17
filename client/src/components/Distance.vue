<template lang="pug">
strong
  | {{distance}}
</template>

<script>

export default {
  name: 'distance',
  props: ['first_user_id', 'second_user_id'],
  data(){
    return {
      distance:0,
    }
  },
 
  async mounted(){
    try {
      const resp = await fetch(`/api/distance?first_id=${this.first_user_id}&second_id=${this.second_user_id}`)
      const data = await resp.json()
      this.distance= data.data.distance
    } catch {
        return this.distance = "error"
      }

  }

}
</script>
