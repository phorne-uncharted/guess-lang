import _ from "lodash";
import Vue from "vue";
import { defaultState, GameState, GuessResult } from "./index";

export const mutations = {
  setGuessResult(state: GameState, guessResult: GuessResult) {
    if (!guessResult) {
      return;
    }
    state.guessResult = guessResult;
  },
  setGameId(state: GameState, gameId: number) {
    if (!gameId) {
      return;
    }
    state.gameId = gameId;
  },
  resetState(state: GameState) {
    Object.assign(state, defaultState());
  },
};
