<template>
  <div>
    <letter-row :letters="row1" @letterclicked="letterClicked" />
    <letter-row :letters="row2" @letterclicked="letterClicked" />
    <letter-row :letters="row3" @letterclicked="letterClicked" />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { CharacterCompare } from "../store/game/index";
import LetterRow from "../components/LetterRow.vue";
import { getLetterInfo, KeyInfo } from "../util/letters";

export default Vue.extend({
  name: "Keyboard",

  components: {
    LetterRow,
  },

  props: {
    letters: {
      type: Object as () => any,
      default: null,
    },
  },
  computed: {
    row1(): CharacterCompare[] {
      const row = [];
      for (const [k, r] of Object.entries(this.letters)) {
        const info = getLetterInfo(+k);
        if (info.row == 1) {
          row[info.col] = {
            index: info.col,
            sourceChar: info.index,
            result: r,
            parsedChar: info.key,
          };
        }
      }
      return row;
    },
    row2(): CharacterCompare[] {
      const row = [];
      for (const [k, r] of Object.entries(this.letters)) {
        const info = getLetterInfo(+k);
        if (info.row == 2) {
          row[info.col] = {
            index: info.col,
            sourceChar: info.index,
            result: r,
            parsedChar: info.key,
          };
        }
      }
      return row;
    },
    row3(): CharacterCompare[] {
      const row = [];
      for (const [k, r] of Object.entries(this.letters)) {
        const info = getLetterInfo(+k);
        if (info.row == 3) {
          row[info.col] = {
            index: info.col,
            sourceChar: info.index,
            result: r,
            parsedChar: info.key,
          };
        }
      }
      return row;
    },
  },

  methods: {
    letterClicked(letter: string) {
      this.$emit("letterclicked", letter);
    },
  },
});
</script>
