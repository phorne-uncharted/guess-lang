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
  setLetterCount(state: GameState, letterCount: number) {
    if (!letterCount) {
      return;
    }
    state.letterCount = letterCount;
  },
  setGuessCount(state: GameState, guessCount: number) {
    if (!guessCount) {
      return;
    }
    state.guessCount = guessCount;
  },
  resetState(state: GameState) {
    Object.assign(state, defaultState());
  },
};
