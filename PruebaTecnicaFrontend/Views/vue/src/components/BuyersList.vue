<template>

    <div>
        <button  type="info" @click="getClients">Listar Compradores</button>
        
      <table class=" zui-table center">
        <thead>
         <tr class="text-center">
            <th>Codigo</th>
            <th>Nombre</th>
            <th>Edad</th>
        </tr>
        </thead>
        <tbody>
         <tr  class="text-center" v-for="buyer in BuyerNames" :key="buyer">
            <td>{{buyer.id}}</td>
            <td>{{buyer.name}}</td>
            <td>{{buyer.age}}</td>
         </tr>
        </tbody>
      </table>

    </div>
  
</template>

<script>

//import moment from 'moment';
import axios from 'axios'

export default {
  name: 'BuyersList',
  props: {
  },
  data(){
    return{
      BuyerNames:null
    }
  },
  methods:{
    getClients(){
      console.log("getClients")
      axios.get('http://localhost:8090/list_buyers')
        .then(response => {
          this.BuyerNames = response.data.people
          console.log(response)
        })
        .catch( e=> console.log(e))
    }
  },
  components: {
  },
}

</script>
@import'~bootstrap/dist/css/bootstrap.css'
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
div{
  height: 10em;
  justify-content: center;
}

button{
  margin-top:60px;
}

table.center {
  margin-left:auto; 
  margin-right:auto;
  margin-top:5px;
}

.zui-table {
    border: solid 1px #DDEEEE;
    border-collapse: collapse;
    border-spacing: 0;
    font: normal 13px Arial, sans-serif;
}
.zui-table thead th {
    background-color: #DDEFEF;
    border: solid 1px #DDEEEE;
    color: #336B6B;
    padding: 10px;
    text-align: center;
    text-shadow: 1px 1px 1px #fff;
}
.zui-table tbody td {
    border: solid 1px #DDEEEE;
    color: #333;
    padding: 10px;
    text-shadow: 1px 1px 1px #fff;
}




</style>
