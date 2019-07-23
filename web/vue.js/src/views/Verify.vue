<template>
  <v-container fluid fill-height>
    <v-layout align-content-center justify-center>
      <v-flex xs12 sm10 md8 class="center-text">
        <h1>Account verification</h1>
        <p>Irgendwas</p>
        <Alert v-model="success" type="success" />
        <Alert v-model="error" type="error" />
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import Alert from '../components/Alert.vue';
import { API_ENDPOINT } from '../constants';

@Component({
  components: {
    Alert,
  },
})
export default class Verify extends Vue {
  private error: string = '';
  private success: string = '';

  private mounted() {
    this.verify();
  }

  private async verify() {
    this.error = '';
    this.success = '';

    try {
      const response = await fetch(API_ENDPOINT + '/api/v1/account/email/' + this.$route.params.id + '/' + this.$route.params.token, {
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        },
        method: 'get',
      });

      const json = await response.json();

      if (response.status >= 200 && response.status < 300) {
        this.success = 'Account verifyed.';
      } else {
        if (json.error) {
          this.error = json.message;
        }
      }
    } catch (e) {
      this.error = 'An error occured';
    }
  }
}
</script>

<style scoped>
.center-text {
  text-align: center;
}
</style>
