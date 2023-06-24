import {createRouter, createWebHistory} from 'vue-router'

// router options
const routes = [
]

// create router
const router = createRouter({
    history: createWebHistory(),
    routes: routes
})

// router hook
router.beforeEach((to, from, next) => {
    console.log('beforeEach triggered');
    console.log('From:', from.path);
    console.log('To:', to.path);
    next();
});
router.afterEach((to, from) => {
    console.log('afterEach triggered');
    console.log('From:', from.path);
    console.log('To:', to.path);
});
  

// expose router
export default router