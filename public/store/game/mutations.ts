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
  setCurrentGuess(state: GameState, guess: number) {
    state.currentGuess = guess;
  },

  updateGuess(
    state: GameState,
    args: { guessIndex: number; letterIndex: number; letter: string }
  ) {
    const code = args.letter.charCodeAt(0);
    Vue.set(
      state.guessResult.knowledge.results[args.guessIndex],
      "word",
      state.guessResult.knowledge.results[args.guessIndex].word.substr(
        0,
        args.letterIndex
      ) +
        args.letter +
        state.guessResult.knowledge.results[args.guessIndex].word.substr(
          args.letterIndex + 1
        )
    );
    Vue.set(
      state.guessResult.knowledge.results[args.guessIndex].comparison,
      args.letterIndex,
      {
        sourceChar: code,
        index: args.letterIndex,
        result: 0,
        parsedChar: args.letter,
      }
    );

    Vue.set(
      state.guessResult.knowledge.results,
      args.guessIndex,
      state.guessResult.knowledge.results[args.guessIndex]
    );
  },

  deleteLetter(
    state: GameState,
    args: { guessIndex: number; letterIndex: number }
  ) {
    Vue.set(
      state.guessResult.knowledge.results[args.guessIndex],
      "word",
      state.guessResult.knowledge.results[args.guessIndex].word.substr(
        0,
        args.letterIndex
      )
    );
    Vue.set(
      state.guessResult.knowledge.results[args.guessIndex].comparison,
      args.letterIndex,
      {
        sourceChar: 32,
        index: args.letterIndex,
        result: 0,
        parsedChar: " ",
      }
    );

    Vue.set(
      state.guessResult.knowledge.results,
      args.guessIndex,
      state.guessResult.knowledge.results[args.guessIndex]
    );
  },

  resetState(state: GameState) {
    Object.assign(state, defaultState());
  },
};
