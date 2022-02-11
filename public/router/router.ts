import Vue from "vue";
import VueRouter from "vue-router";
import Game from "../views/Game.vue";

Vue.use(VueRouter);

const router = new VueRouter({
  routes: [
    { path: "/", redirect: "/game" },
    { path: "/game", component: Game },
  ],
});

export default router;
