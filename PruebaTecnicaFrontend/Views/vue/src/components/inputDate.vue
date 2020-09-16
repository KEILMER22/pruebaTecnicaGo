<template>
    <div class="center">
      
      <datepicker  :format=dateFormat v-model="date" placeholder="Ingrese una fecha" name="fecha" @input="fechaSeleccionada" >
      </datepicker>
      <p> estado: {{dateLabel}}</p>
     
    </div>
  
</template>

<script>

import Datepicker from 'vuejs-datepicker';
import moment from 'moment';
import axios from 'axios'

export default {
  name: 'inputDate',
  endPointRoute: 'http://localhost:8090/load_data/', 
  props: {
    msg: String
  },
  data(){
    return{
      dateFormat : 'dd-MM-yyyy',
      date: new Date(),
      fecha: null,
      dateLabel: "sincronizando..."
    }
  },
  methods:{
    fechaSeleccionada: function() {
      this.dateLabel = "Sincronizando..."
      this.fecha = moment(String(this.date)).format('DD-MM-YYYY')
      this.endPointRoute = this.endPointRoute+String(this.fecha)
      axios.get(endPointRoute)
        .then(response => {
          console.log(response)
          this.dateLabel = "Sincronizado en: "+  this.fecha
        })
        .catch( e=> this.dateLabel = "Sincronizacion fallida "+ e)
    }    
  },
  components: {
        Datepicker
  },
  mounted(){
    this.fechaSeleccionada()

  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.center{
  margin-left:auto; 
  margin-right:auto;
}
div{
  text-align:center;
  justify-content: center;
  align-items: center;
}
datepicker {
  margin-right: 10px;
}
</style>
