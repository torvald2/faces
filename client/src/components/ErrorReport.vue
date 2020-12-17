<template lang="pug">
b-container(fluid)
  b-row
    b-col
      h1.text-center.mt-2 Отчет об ошибках
  b-row
    b-col 
      b-form(inline)
        label(for="date-from") Дата с
        b-form-datepicker(id="date-from" v-model="date_from").ml-2
        label(for="date-to").ml-2 Дата до
        b-form-datepicker(id="date-to" v-model="date_to").ml-2
        b-button(variant="primary" @click="load_data").ml-2 Сформировать    
  b-row.mt-2
    b-col
      b-form(inline)
        label(for="filter-error-type") Тип ошибки
        b-form-select(v-model="error_type_filter" :options="filter_options" id='filter-error-type').ml-3
  b-row
    b-col
      b-alert(
      variant="danger"
      dismissible
      fade
      :show="is_error"
      @dismissed="is_error=false"
      ) Ошибка загрузки данных: {{error_text}}
  b-row
    b-col 
      b-table(:items="items" :busy="isBusy" :fields="table_fields" class="mt-3" outlined)
        template(#table-busy)
          .text-center.text-danger.my-2
            b-spinner.align-middle
            | Loading...
        template(#cell(current_image)="row")
          b-img(v-bind="image_props" :src="format_current_image_url(row.item)" fluid)
        template(#cell(new_image)="row")
          div(v-if="row.item.recognized_users")
            b-img( v-bind="image_props" :src="`/images/${row.item.recognized_users[0]}`")
        template(#cell(face_distance)="row")
          distance(v-if="row.item.recognized_users",
                   :first_user_id="row.item.user_id"
                   :second_user_id="row.item.recognized_users[0]"
           )
            
        
        
</template>

<script>
import distance from '../components/Distance.vue'

export default {
  name: 'RequestPicture',
  components:{distance},
  data(){
    return{
      isBusy: false,
      date_from: '',
      date_to: '',
      error_type_filter: '',
      loaded_data: [],
      is_error:false,
      error_text:'',
      filter_options:[
        {value:'', text:'Все ошибки'},
        {value:'not_me', text:'Не верное определение лица'},
        {value:'not_recognized',text:'Лицо не расспознано'},
        {value:'no_face', text:'Лицо не найдено'},
        {value:'user_not_found', text:'Пользователь не найден'}
      ],
      table_fields:[
        { key: 'recognize_time', label: 'Дата и время операции'},
        {key:'error_type', label:'Тип ошибки'},
        {key:'current_image', label:'Текущее изображение'},
        {key:'new_image', label:'Совпавшее лицо'},
        {key:'face_distance',label:'Расстояние'}
      ],
      image_props: {  width: 120, height: 120, class: 'm1' }
    }
  },  
  methods: {
    format_current_image_url(item){
      if (item.error_type === "not_me"){
        return `/images/${item.user_id}`
      } else {
        return `/imagesrequest/${item.record_id}`

      }
    },
    async load_data(){
      this.isBusy = true
      try {
        const resp =  await fetch(`/api/badrequest?start=${this.date_from}&end=${this.date_to}`)
        const response_data = await resp.json()
        const res = response_data.data
        this.loaded_data = res
        this.isBusy = false
      } catch(err){
         this.error_text = err
         this.is_error = true
         this.isBusy = false
      }
    },
    async get_distance(profile1, profile2){
      try {
      const resp = await fetch(`/api/distance?first_id=${profile1}&second_id=${profile2}`)
      const data = await resp.json()
      return data.data.distance
      } catch {
        return "ERROR"
      }

    }
  },
computed:{
  items(){
    if (this.error_type_filter === ''){
      return this.loaded_data
    }
    else {
      return this.loaded_data.filter(obj => obj.error_type === this.error_type_filter)
    }
  }
}
}
</script>
