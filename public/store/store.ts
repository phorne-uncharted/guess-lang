import Vue from "vue";
import Vuex, { Store } from "vuex";
import { Route } from "vue-router";
import { gameModule } from "./game/module";
import { GameState } from "./game/index";

Vue.use(Vuex);

export interface GuessState {
  gameModule: GameState;
}

const store = new Store<GuessState>({
  modules: {
    gameModule,
  },
  strict: true,
});

export default store;
