<template>
  <div>
    <div v-for="l in letters" class="letter">
      <div>
        <letter
          :source-char="l"
          @letterclicked="letterClicked"
          @deleteclicked="deleteClicked"
          @enterclicked="enterClicked"
        />
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

  methods: {
    letterClicked(letter: string) {
      this.$emit("letterclicked", letter);
    },
    deleteClicked() {
      this.$emit("deleteclicked");
    },
    enterClicked() {
      this.$emit("enterclicked");
    },
  },
});
</script>

<style>
.letter {
  display: table-cell;
}
</style>
