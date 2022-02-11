<template>
  <div>
    <div v-for="res in result.comparison" class="letter">
      <letter
        :source-char="res.sourceChar"
        :result="res.result"
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
});
</script>

<style>
.letter {
  display: table-cell;
}
</style>
