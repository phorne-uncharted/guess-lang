import { GuessResult } from "../store/game/index";

export function defaultGuessResult(
  guessResult: GuessResult,
  guessCount: number,
  letterCount: number
): GuessResult {
  if (!guessResult.knowledge.results) {
    guessResult.knowledge.results = [];
  }

  for (var i = 0; i < guessCount; i++) {
    if (
      guessResult.knowledge.results.length < i ||
      !guessResult.knowledge.results[i]
    ) {
      const cr = { word: "", comparison: [] };
      for (var j = 0; j < letterCount; j++) {
        cr.word = cr.word + " ";
        cr.comparison.push({
          index: j,
          sourceChar: 32,
          result: 0,
          parsedChar: " ",
        });
      }
      if (guessResult.knowledge.results.length < i) {
        guessResult.knowledge.results.push(cr);
      } else {
        guessResult.knowledge.results[i] = cr;
      }
    }
  }

  return guessResult;
}
