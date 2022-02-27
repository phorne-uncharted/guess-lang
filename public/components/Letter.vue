<template>
  <b-button v-bind:class="getClass()" @click="letterClicked">
    {{ displayCharacter() }}
  </b-button>
</template>

<script lang="ts">
import _ from "lodash";
import Vue from "vue";
import { CharacterCompare } from "../store/game/index";

export default Vue.extend({
  name: "letter",

  props: {
    sourceChar: {
      type: Object as () => CharacterCompare,
      default: null,
    },
    barNotMatched: {
      type: Boolean,
      default: true,
    },
  },

  methods: {
    displayCharacter(): string {
      return this.sourceChar.parsedChar;
    },
    getClass(): string {
      if (this.sourceChar.result == 1 && this.barNotMatched) {
        return "not-matched letter";
      } else if (this.sourceChar.result == 2) {
        return "in-word letter";
      } else if (this.sourceChar.result == 3) {
        return "at-place letter";
      }

      return "letter";
    },

    letterClicked() {
      this.$emit("letterclicked", this.sourceChar.parsedChar);
    },
  },
});
</script>

<style scoped>
.at-place {
  background-color: #6aaa64;
}
.in-word {
  background-color: #c9b458;
}
.not-matched {
  background-color: #787c7e;
}
.letter {
  padding: 10px;
  font-weight: bold;
  min-width: 35px;
  min-height: 45px;
  margin: 3px;
}
</style>
