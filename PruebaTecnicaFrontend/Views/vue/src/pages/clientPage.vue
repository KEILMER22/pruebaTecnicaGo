<template>
    <div id="clientPage">
        <img alt="Vue logo" src="../assets/logo.png">
        <p></p>  
        <div>
            <input v-model="id" placeholder="ingrese el id">
            <button  type="info" @click="searchClient">Buscar comprador</button>
        </div>
        <div>
          <md-table md-card>
            <md-table-toolbar>
              <h1 class="md-title">Misma Ip</h1>
            </md-table-toolbar>
            <md-table-row>
              <md-table-head>ID</md-table-head>
              <md-table-head>Name</md-table-head>
              <md-table-head>Age</md-table-head>
            </md-table-row>
            <md-table-row  v-for="item in buyerIp" :key="item">
              <md-table-cell>{{item.id}}</md-table-cell>
              <md-table-cell>{{item.name}}</md-table-cell>
              <md-table-cell>{{item.age}}</md-table-cell>
            </md-table-row>
          </md-table>
          <md-table md-card>
            <md-table-toolbar>
              <h1 class="md-title">Historial de compra</h1>
            </md-table-toolbar>
            <md-table-row>
              <md-table-head>ID</md-table-head>
              <md-table-head>Name</md-table-head>
              <md-table-head>Price</md-table-head>
            </md-table-row>
            <md-table-row  v-for="prod in buyerProd" :key="prod">
              <md-table-cell>{{prod.id}}</md-table-cell>
              <md-table-cell>{{prod.name}}</md-table-cell>
              <md-table-cell>{{prod.price}}</md-table-cell>
            </md-table-row>
          </md-table>
          <md-table md-card>
            <md-table-toolbar>
              <h1 class="md-title">Recomendaciones</h1>
            </md-table-toolbar>
            <md-table-row>
              <md-table-head>ID</md-table-head>
              <md-table-head>Name</md-table-head>
              <md-table-head>Price</md-table-head>
            </md-table-row>
            <md-table-row  v-for="item in recomendations" :key="item">
              <md-table-cell>{{item.id}}</md-table-cell>
              <md-table-cell>{{item.name}}</md-table-cell>
              <md-table-cell>{{item.price}}</md-table-cell>
            </md-table-row>
          </md-table>
        </div>
    </div>
</template>

<script>

//import moment from 'moment';
import axios from 'axios'

export default {
  name: 'ClientPage',
  buyerIp:null ,
  buyerProd:null,
  recomendations:null,
  id:null,
  props: {
  },
  data(){
    return{
      buyerIp:null ,
      buyerProd:null,
      recomendations:null,
      id:null,
    }
  },
  methods:{
    searchClient(){
      console.log('http://localhost:8090/buyer_info/'+this.id)
      axios.get('http://localhost:8090/buyer_info/'+this.id)
        .then(response => {
          this.buyerIp = response.data.sameIpBuyers
          this.buyerProd = response.data.purchaseHistory
          this.recomendations = response.data.recomendations
          console.log(response)
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
  justify-content: center;
}


</style>
