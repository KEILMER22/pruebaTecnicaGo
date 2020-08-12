<template>
    <div id="clientPage">
        <img alt="Vue logo" src="../assets/logo.png">
        <p></p>  
        <div>
            <input v-model="id" placeholder="ingrese el id">
            <button  type="info" @click="searchClient">Listar Compradores</button>
        </div>
      <table class=" zui-table center" style="display: inline-block; border: 1px solid; ">
        <thead>
         <tr class="text-center">
            <th>Usuarios con la misma Ip</th>
        </tr>
        </thead>
        <tbody>
         <tr  class="text-center" v-for="buyer in buyerIp" :key="buyer">
            <td>{{buyer.item}}</td>
         </tr>
        </tbody>
      </table>
            <table class=" zui-table center" style="display: inline-block; border: 1px solid; ">
        <thead>
         <tr class="text-center">
            <th>Compras del usuario</th>
        </tr>
        </thead>
        <tbody>
         <tr  class="text-center" v-for="buy in buyerProd" :key="buy">
            <td>{{buy.item}}</td>
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
  buyerIp:null ,
  buyerProd:null,
  props: {
  },
  data(){
    return{
      buyerIp:null ,
      buyerProd:null,
    }
  },
  methods:{
    getClients(){
      console.log("getClients")
      axios.get('http://localhost:8090/list_buyers')
        .then(response => {
          this.BuyerNames = response.data.people
          console.log(this.BuyerNames)
        })
        .catch( e=> console.log(e))
    },
    searchClient(){
      console.log(this.id)
      axios.get('http://localhost:8090/buyer_info/'+this.id)
        .then(response => {
          this.buyerIp = response.data.sameIpBuyers
          this.buyerProd = response.data.purchaseHistory
        })
        .catch( e=> console.log(e))
    }
  },
  components: {
  },
}

</script>

<style>
#clientPage {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.myButton {
	box-shadow:inset 0px 1px 0px 0px #ffffff;
	background:linear-gradient(to bottom, #f9f9f9 5%, #e9e9e9 100%);
	background-color:#f9f9f9;
	border-radius:6px;
	border:1px solid #dcdcdc;
	display:inline-block;
	cursor:pointer;
	color:#666666;
	font-family:Arial;
	font-size:15px;
	font-weight:bold;
	padding:6px 24px;
	text-decoration:none;
	text-shadow:0px 1px 0px #ffffff;
  margin-bottom: 20px;
}
.myButton:hover {
	background:linear-gradient(to bottom, #e9e9e9 5%, #f9f9f9 100%);
	background-color:#e9e9e9;
}
.myButton:active {
	position:relative;
	top:1px;
}
div{
  height: 10em;
  justify-content: center;
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
