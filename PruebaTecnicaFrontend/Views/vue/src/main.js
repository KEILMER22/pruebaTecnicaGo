import vue from 'vue'
import index from './pages/index.vue'
import vueRouter from 'vue-router'
import clientPage from './pages/clientPage.vue'
import home from './pages/home.vue'
import VueMaterial from 'vue-material'
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css'

vue.use(VueMaterial)
vue.config.productionTip = false
vue.use(vueRouter);

const routes = [
    { path: '/clientPage', component: clientPage },
    { path: '/', component: home }
];
const router = new vueRouter({
    routes,
    mode: 'history'
});

new vue({
    render: h => h(index),
    router
}).$mount('#app')