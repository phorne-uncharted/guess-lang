<template>
  <div>
    <div v-for="l in letters" class="letter">
      <div>
        <letter :source-char="l" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Letter from "../components/Letter.vue";
import { CharacterCompare } from "../store/game/index";

export default Vue.extend({
  name: "LetterRow",

  components: {
    Letter,
  },

  props: {
    letters: {
      type: Object as () => CharacterCompare[],
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
