<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <!--<input-date> </input-date>
    -->
    <BuyersList>

    </BuyersList>

  </div>
</template>

<script>
//import HelloWorld from './components/HelloWorld.vue'
//import inputDate from './components/inputDate.vue'
import BuyersList from './components/BuyersList.vue'
export default {
  name: 'App',
  components: {
    //HelloWorld,
    //inputDate,
    BuyersList
  }

}

    //<HelloWorld msg="Bienvenido a la prueba tecnica"/>

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
