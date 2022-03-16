import axios, { AxiosResponse } from "axios";
import _ from "lodash";
import { ActionContext } from "vuex";
import { gameGetters } from "..";
import store, { GuessState } from "../store";
import { getters, mutations } from "./module";
import { GameState } from "./index";
import { defaultGuessResult } from "../../util/knowledge";

export type GameContext = ActionContext<GameState, GuessState>;

export const actions = {
  async startGame(
    context: GameContext,
    args: { language: string; maxGuessCount: number; letterCount: number }
  ): Promise<void> {
    try {
      const response = await axios.post(`/game/start`, {
        language: args.language,
        maxGuessCount: args.maxGuessCount,
        letterCount: args.letterCount,
      });
      mutations.setGameId(context, response.data.gameId);
      mutations.setGuessResult(
        context,
        defaultGuessResult(
          response.data.game,
          args.maxGuessCount,
          args.letterCount
        )
      );
    } catch (error) {
      console.error(error);
      mutations.setGameId(context, -1);
    }
  },
  async guessWord(
    context: GameContext,
    args: {
      word: string;
      gameId: number;
      maxGuessCount: number;
      letterCount: number;
    }
  ): Promise<void> {
    try {
      const response = await axios.post(`/game/guess`, {
        word: args.word,
        gameId: args.gameId,
      });
      mutations.setGuessResult(
        context,
        defaultGuessResult(response.data, args.maxGuessCount, args.letterCount)
      );
      mutations.setCurrentGuess(context, getters.getCurrentGuess(context) + 1);
    } catch (error) {
      console.log(error);
    }
  },
};
