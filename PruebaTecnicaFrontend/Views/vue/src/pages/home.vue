<template>
  <div id="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <div>
      <a href="clientPage" class="myButton">Consultar Cliente</a>
      <inputDate></inputDate>
      <BuyersList> </BuyersList> 
    </div>

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
</style>
