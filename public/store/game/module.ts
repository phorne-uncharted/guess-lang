import { Module } from "vuex";
import { getStoreAccessors } from "vuex-typescript";
import { GuessState } from "../store";
import { actions as moduleActions } from "./actions";
import { getters as moduleGetters } from "./getters";
import { GameState, state } from "./index";
import { mutations as moduleMutations } from "./mutations";

export const gameModule: Module<GameState, GuessState> = {
  getters: moduleGetters,
  actions: moduleActions,
  mutations: moduleMutations,
  state: state,
};

const { commit, read, dispatch } = getStoreAccessors<GameState, GuessState>(
  null
);

// Typed getters
export const getters = {
  getGameId: read(moduleGetters.getGameId),
  getLetterCount: read(moduleGetters.getLetterCount),
  getGuessCount: read(moduleGetters.getGuessCount),
  getGuessResult: read(moduleGetters.getGuessResult),
};

// Typed actions
export const actions = {
  startGame: dispatch(moduleActions.startGame),
  guessWord: dispatch(moduleActions.guessWord),
};

// Typed mutations
export const mutations = {
  setGuessResult: commit(moduleMutations.setGuessResult),
  setGameId: commit(moduleMutations.setGameId),
  setLetterCount: commit(moduleMutations.setLetterCount),
  setGuessCount: commit(moduleMutations.setGuessCount),
};
