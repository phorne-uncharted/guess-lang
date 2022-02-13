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
    };
  },

  methods: {
    handleOk() {
      mutations.setLetterCount(this.$store, this.letterCount);
    },
  },
});
</script>
