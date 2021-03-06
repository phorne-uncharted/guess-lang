<template>
  <div class="container-fluid d-flex game">
    <settings-modal />
    <div class="">
      <div class="title">Guess The Word ({{ letterCount }} Letters)</div>
      <div v-if="haveResults">
        <div class="flex-row">
          <div v-for="check in knowledge.knowledge.results">
            <div>
              <guess-letter :result="check" />
            </div>
          </div>
          <div v-if="solved">
            <span class="checkmark">
              <div class="checkmark_stem"></div>
              <div class="checkmark_kick"></div>
            </span>
          </div>
          <div v-if="failed">
            {{ target }}
            <div class="failure"></div>
          </div>
        </div>
        <div class="alphabet">
          <keyboard
            :letters="knowledge.knowledge.letters"
            @letterclicked="letterClicked"
            @deleteclicked="deleteClicked"
            @enterclicked="enterClicked"
          />
        </div>
      </div>
      <!--<form ref="guessInputForm">
        <b-form-group label-for="guess-input">
          <b-form-input
            id="guess-input"
            v-model="guess"
            @keyup.enter="guessWord"
          />
        </b-form-group>
      </form> -->
      <b-button variant="primary" @click="guessWord" :disabled="!canGuess">
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
import Keyboard from "../components/Keyboard.vue";
import SettingsModal from "../components/SettingsModal.vue";
import { GuessResult, CheckResult } from "../store/game/index";
import { actions, getters, mutations } from "../store/game/module";
import { defaultGuessResult } from "../util/knowledge";

export default Vue.extend({
  name: "game",

  components: {
    GuessLetter,
    Letter,
    SettingsModal,
    Keyboard,
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
    solved(): boolean {
      return this.haveResults && getters.getGuessResult(this.$store).solved;
    },
    canGuess(): boolean {
      return (
        !this.isGuessing &&
        ((this.haveResults && !getters.getGuessResult(this.$store).done) ||
          !this.haveResults)
      );
    },
    failed(): boolean {
      return (
        this.haveResults &&
        getters.getGuessResult(this.$store).done &&
        !this.solved
      );
    },
    target(): string {
      if (this.failed) {
        return getters.getGuessResult(this.$store).target;
      }

      return "";
    },
    letterCount(): number {
      return getters.getLetterCount(this.$store);
    },
  },

  methods: {
    async guessWord() {
      this.isGuessing = true;
      await actions.guessWord(this.$store, {
        word: this.guess,
        gameId: this.gameId,
        maxGuessCount: getters.getGuessCount(this.$store),
        letterCount: getters.getLetterCount(this.$store),
      });
      this.knowledge = getters.getGuessResult(this.$store);
      const currentGuess = getters.getCurrentGuess(this.$store);
      if (
        currentGuess < getters.getGuessCount(this.$store) &&
        this.knowledge.knowledge.results[currentGuess].word[0] == " "
      ) {
        this.guess = "";
      }
      this.isGuessing = false;
    },

    async startGame() {
      this.isGuessing = true;
      this.guess = "";
      await actions.startGame(this.$store, {
        language: "fr",
        maxGuessCount: getters.getGuessCount(this.$store),
        letterCount: getters.getLetterCount(this.$store),
      });
      this.gameId = getters.getGameId(this.$store);

      this.knowledge = getters.getGuessResult(this.$store);
      mutations.setCurrentGuess(this.$store, 0);
      this.isGuessing = false;
    },

    letterClicked(letter: string) {
      if (this.guess.length < getters.getLetterCount(this.$store)) {
        this.guess = this.guess + letter;
        mutations.updateGuess(this.$store, {
          guessIndex: getters.getCurrentGuess(this.$store),
          letterIndex: this.guess.length - 1,
          letter: letter,
        });
        this.knowledge = getters.getGuessResult(this.$store);
      }
    },

    deleteClicked() {
      this.guess = this.guess.substr(0, this.guess.length - 1);
      mutations.deleteLetter(this.$store, {
        guessIndex: getters.getCurrentGuess(this.$store),
        letterIndex: this.guess.length,
      });
      this.knowledge = getters.getGuessResult(this.$store);
    },

    enterClicked() {
      this.guessWord();
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

.game {
  justify-content: center;
}

.letter {
  display: table-cell;
  min-width: 35px;
  min-height: 45px;
}

.alphabet {
  margin-top: 20px;
  margin-bottom: 20px;
}

.checkmark {
  display: inline-block;
  width: 66px;
  height: 66px;
  -ms-transform: rotate(45deg); /* IE 9 */
  -webkit-transform: rotate(45deg); /* Chrome, Safari, Opera */
  transform: rotate(45deg);
}

.checkmark_stem {
  position: absolute;
  width: 9px;
  height: 27px;
  background-color: #6aaa64;
  left: 33px;
  top: 18px;
}

.checkmark_kick {
  position: absolute;
  width: 9px;
  height: 9px;
  background-color: #6aaa64;
  left: 24px;
  top: 36px;
}

.failure {
  height: 100px;
  width: 100px;
  border-radius: 5px;
  position: relative;
  &:after {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    content: "\274c";
    font-size: 60px;
    color: #fff;
    line-height: 100px;
    text-align: center;
  }
}
</style>
