<template>
  <b-modal
    id="settings"
    title="Game Creation Settings"
    @ok="handleOk"
    @show="onShow"
  >
    <b-form-group
      id="letter-count"
      label="Number of letters in the word:"
      label-for="letter-count-spinner"
      description="Sets the number of letters in the word."
    >
      <b-form-spinbutton
        id="letter-count-spinner"
        v-model="letterCount"
        inline
        min="3"
        max="20"
      />
    </b-form-group>
    <b-form-group
      id="guess-count"
      label="Number of guesses allowed:"
      label-for="guess-count-spinner"
      description="Sets the number of guesses allowed."
    >
      <b-form-spinbutton
        id="guess-count-spinner"
        v-model="guessCount"
        inline
        min="1"
        max="15"
      />
    </b-form-group>
  </b-modal>
</template>

<script lang="ts">
import Vue from "vue";
import { getters, mutations } from "../store/game/module";

export default Vue.extend({
  name: "SettingsModal",

  data() {
    return {
      letterCount: getters.getLetterCount(this.$store) || 5,
      guessCount: getters.getGuessCount(this.$store) || 6,
    };
  },

  methods: {
    handleOk() {
      mutations.setLetterCount(this.$store, this.letterCount);
      mutations.setGuessCount(this.$store, this.guessCount);
    },
  },
});
</script>
