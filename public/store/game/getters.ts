import { GameState, GuessResult } from "./index";

export const getters = {
  getGuessResult(state: GameState): GuessResult {
    return state.guessResult;
  },
  getGameId(state: GameState): number {
    return state.gameId;
  },
  getLetterCount(state: GameState): number {
    return state.letterCount;
  },
  getGuessCount(state: GameState): number {
    return state.guessCount;
  },
};
