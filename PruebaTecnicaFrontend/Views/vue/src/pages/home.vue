<template>
  <div id="home">
    
    <img alt="Vue logo" src="../assets/logo.png">
    
    <div>
      <a href="clientPage" class="myButton">Consultar Cliente</a>
    </div>

    <input-date> </input-date>
    
    <BuyersList> </BuyersList> 

  </div>
</template>

<script>
import inputDate from '../components/inputDate.vue'
import BuyersList from '../components/buyersList.vue'
export default {
  name: 'home',
  components: {
     inputDate,
     BuyersList
  }

}


</script>

<style>
#home {
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
</style>
