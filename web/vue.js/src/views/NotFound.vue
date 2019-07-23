<template>
  <v-container fluid fill-height>
    <v-layout align-center justify-center>
      <v-flex xs12 sm10 md8 class="center-text">
        <h1>That happens not to be a page</h1>
        <v-divider></v-divider>
        <p>Rewind and try another one</p>
        <p>
          <v-btn color="primary" @click="goBack"><v-icon left dark>arrow_back</v-icon> Back</v-btn>
          <v-btn color="primary" @click="goHome">Home <v-icon right dark>home</v-icon></v-btn>
        </p>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class NotFound extends Vue {
  private goHome() {
    this.$router.replace('/');
  }

  private goBack() {
    this.$router.back();
  }
}
</script>

<style scoped>
.center-text {
  text-align: center;
}
</style>
