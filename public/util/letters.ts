const KEY_LAYOUT = {
  100: { row: 3, col: 0, key: "Del", index: 100 },
  101: { row: 3, col: 8, key: "Enter", index: 101 },
  65: { row: 2, col: 0, key: "A", index: 65 },
  66: { row: 3, col: 5, key: "B", index: 66 },
  67: { row: 3, col: 3, key: "C", index: 67 },
  68: { row: 2, col: 2, key: "D", index: 68 },
  69: { row: 1, col: 2, key: "E", index: 69 },
  70: { row: 2, col: 3, key: "F", index: 70 },
  71: { row: 2, col: 4, key: "G", index: 71 },
  72: { row: 2, col: 5, key: "H", index: 72 },
  73: { row: 1, col: 7, key: "I", index: 73 },
  74: { row: 2, col: 6, key: "J", index: 74 },
  75: { row: 2, col: 7, key: "K", index: 75 },
  76: { row: 2, col: 8, key: "L", index: 76 },
  77: { row: 3, col: 7, key: "M", index: 77 },
  78: { row: 3, col: 6, key: "N", index: 78 },
  79: { row: 1, col: 8, key: "O", index: 79 },
  80: { row: 1, col: 9, key: "P", index: 80 },
  81: { row: 1, col: 0, key: "Q", index: 81 },
  82: { row: 1, col: 3, key: "R", index: 82 },
  83: { row: 2, col: 1, key: "S", index: 83 },
  84: { row: 1, col: 4, key: "T", index: 84 },
  85: { row: 1, col: 6, key: "U", index: 85 },
  86: { row: 3, col: 4, key: "V", index: 86 },
  87: { row: 1, col: 1, key: "W", index: 87 },
  88: { row: 3, col: 2, key: "X", index: 88 },
  89: { row: 1, col: 5, key: "Y", index: 89 },
  90: { row: 3, col: 1, key: "Z", index: 90 },
};

export interface KeyInfo {
  row: number;
  col: number;
  key: string;
  index: number;
}

export function getLetterInfo(letter: number): KeyInfo {
  return KEY_LAYOUT[letter];
}
