<template>
  <div v-bind:class="getClass()">
    {{ displayCharacter() }}
  </div>
</template>

<script lang="ts">
import _ from "lodash";
import Vue from "vue";

export default Vue.extend({
  name: "letter",

  props: {
    sourceChar: {
      type: Number,
      default: 0,
    },
    result: {
      type: Number,
      default: 0,
    },
    barNotMatched: {
      type: Boolean,
      default: true,
    },
  },

  methods: {
    displayCharacter(): string {
      return String.fromCharCode(this.sourceChar);
    },
    getClass(): string {
      if (this.result == 1 && this.barNotMatched) {
        return "not-matched letter";
      } else if (this.result == 2) {
        return "in-word letter";
      } else if (this.result == 3) {
        return "at-place letter";
      }

      return "letter";
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
  text-decoration: line-through;
}
.letter {
  padding: 10px;
  font-weight: bold;
}
</style>
