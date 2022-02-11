export interface CharacterCompare {
  index: number;
  sourceChar: number;
  result: number;
}

export interface CheckResult {
  word: string;
  comparison: CharacterCompare[];
}

export interface Knowledge {
  results: CheckResult[];
  letters: any;
}

export interface GuessResult {
  check: CheckResult;
  knowledge: Knowledge;
}

export interface GameState {
  guessResult: GuessResult;
  gameId: number;
}

export const defaultState = (): GameState => {
  return {
    guessResult: null,
    gameId: -1,
  };
};

export const state: GameState = defaultState();
