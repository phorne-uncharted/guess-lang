<template>
  <div class="container-fluid d-flex join-view">
    <settings-modal />
    <div class="">
      <div v-if="haveResults">
        <div class="flex-row">
          <div v-for="check in knowledge.knowledge.results">
            <div>
              <guess-letter :result="check" />
            </div>
          </div>
        </div>
        <div>
          <div
            v-for="(res, letter) in knowledge.knowledge.letters"
            class="letter"
          >
            <div>
              <letter :source-char="parseInt(letter)" :result="res" />
            </div>
          </div>
        </div>
      </div>
      <form ref="guessInputForm">
        <b-form-group label-for="guess-input">
          <b-form-input id="guess-input" v-model="guess" />
        </b-form-group>
      </form>
      <b-button variant="primary" @click="guessWord" :disabled="isGuessing">
        <b-spinner v-if="isGuessing" small />
        <span v-else>guess</span>
      </b-button>
      <b-button-group>
        <b-button variant="primary" @click="startGame" :disabled="isGuessing">
          <b-spinner v-if="isGuessing" small />
          <span v-else>restart</span>
        </b-button>
        <b-button v-b-modal.settings variant="success" :disabled="isGuessing">
          <i class="fa fa-cog" aria-hidden="true" />
        </b-button>
      </b-button-group>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Letter from "../components/Letter.vue";
import GuessLetter from "../components/GuessLetter.vue";
import SettingsModal from "../components/SettingsModal.vue";
import { CheckResult } from "../store/game/index";
import { actions, getters } from "../store/game/module";

export default Vue.extend({
  name: "game",

  components: {
    GuessLetter,
    Letter,
    SettingsModal,
  },

  data() {
    return {
      guess: "",
      isGuessing: false,
      gameId: -1,
      knowledge: null,
    };
  },

  computed: {
    haveResults(): boolean {
      return this.knowledge != null;
    },
    knowledgeResults(): CheckResult[] {
      return this.knowledge.results;
    },
  },

  methods: {
    async guessWord() {
      this.isGuessing = true;
      await actions.guessWord(this.$store, {
        word: this.guess,
        gameId: this.gameId,
      });
      this.knowledge = getters.getGuessResult(this.$store);
      this.guess = "";
      this.isGuessing = false;
    },
    async startGame() {
      this.isGuessing = true;
      await actions.startGame(this.$store, {
        language: "fr",
        maxGuessCount: 15,
        letterCount: getters.getLetterCount(this.$store),
      });
      this.gameId = getters.getGameId(this.$store);
      this.knowledge = null;
      this.isGuessing = false;
    },
  },

  async beforeMount() {
    this.startGame();
  },
});
</script>

<style>
.header-label {
  padding: 1rem 0 0.5rem 0;
  font-weight: bold;
}

.letter {
  display: table-cell;
}
</style>
