<template>
  <div>
    <div v-for="res in result.comparison" class="letter">
      <letter
        :source-char="parseCharacterCompare(res)"
        :barNotMatched="false"
      />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Letter from "../components/Letter.vue";
import { CheckResult } from "../store/game/index";
import { actions, getters } from "../store/game/module";
import { CharacterCompare } from "../store/game/index";

export default Vue.extend({
  name: "GuessLetter",

  components: {
    Letter,
  },

  props: {
    result: {
      type: Object as () => CheckResult,
      default: null,
    },
  },
  methods: {
    parseCharacterCompare(res: CharacterCompare): CharacterCompare {
      res.parsedChar = String.fromCharCode(res.sourceChar);

      return res;
    },
  },
});
</script>

<style>
.letter {
  display: table-cell;
}
</style>
