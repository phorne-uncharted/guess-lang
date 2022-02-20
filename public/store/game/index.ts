export interface CharacterCompare {
  index: number;
  sourceChar: number;
  result: number;
  parsedChar: string;
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
  done: boolean;
  solved: boolean;
  target: string;
}

export interface GameState {
  guessResult: GuessResult;
  gameId: number;
  letterCount: number;
  guessCount: number;
}

export const defaultState = (): GameState => {
  return {
    guessResult: null,
    gameId: -1,
    letterCount: 5,
    guessCount: 6,
  };
};

export const state: GameState = defaultState();
