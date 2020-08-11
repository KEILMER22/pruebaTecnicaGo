import Vue from 'vue'
import index from './pages/index.vue'
import VueRouter from 'vue-router'
import clientPage from './pages/clientPage.vue'
import home from './pages/home.vue'

Vue.config.productionTip = false
Vue.use(VueRouter);

const routes = [
    { path: '/clientPage', component: clientPage },
    { path: '/', component: home }
];
const router = new VueRouter({
    routes,
    mode: 'history'
});

new Vue({
    render: h => h(index),
    router
}).$mount('#app')