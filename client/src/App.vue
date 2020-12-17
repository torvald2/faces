<template lang="pug">
div
  b-navbar(toggleable="lg" type="dark" variant="info")
    b-navbar-brand Отчеты системы регистрации прихода/ухода
    b-navbar-nav Отчет об ошибках
  b-container
    report
</template>

<script>
import report from './components/ErrorReport.vue'


export default {
  name: 'App',
  components: {
    report,
  },
 

}
</script>