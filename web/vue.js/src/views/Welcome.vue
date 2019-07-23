<template>
  <v-container fluid fill-height>
    <v-layout align-content-center justify-center>
      <v-flex xs12 sm10 md8 class="center-text">
        <v-layout row wrap>
          <v-flex xs12>
            <h1>Welcome to the Go &#9829; Vue Starter</h1>
          </v-flex>
        </v-layout>

        <v-layout row wrap>
          <v-flex xs12>
            <p>Hi {{user.name}} and welcome to the Go &#9829; Vue Starter. Befor you start please verify your email address. If you have not get an verification email you can use the button below to resend the verification email.</p>
            <LoremIpsum />
          </v-flex>
        </v-layout>

        <v-layout row wrap>
          <v-flex xs12>
            <Alert v-model="success" type="success" />
            <Alert v-model="error" type="error" />
            <v-btn v-if="user.id !== undefined" color="primary" :loading="loading" :disabled="loading" @click="resendVerificationEmail">Resend verification email</v-btn>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { State } from 'vuex-class';

import Alert from '../components/Alert.vue';
import LoremIpsum from '../components/LoremIpsum.vue';
import { API_ENDPOINT } from '../constants';
import { IUserState } from '../store/modules/user';

@Component({
  components: {
    Alert,
    LoremIpsum,
  },
})
export default class Welcome extends Vue {
  @State('user') private user!: IUserState;

  private loading: boolean = false;
  private error: string = '';
  private success: string = '';

  private async resendVerificationEmail() {
    this.error = '';
    this.success = '';
    this.loading = true;

    try {
      const response = await fetch(API_ENDPOINT + '/api/v1/account/email', {
        body: JSON.stringify({
          id: this.user.id,
        }),
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        },
        method: 'post',
      });

      const json = await response.json();

      if (response.status >= 200 && response.status < 300) {
        this.success = 'Verification email was sent.';
        this.loading = false;
      } else {
        if (json.error) {
          this.error = json.message;
          this.loading = false;
        }
      }
    } catch (e) {
      this.error = 'An error occured';
      this.loading = false;
    }
  }
}
</script>

<style scoped>
.center-text {
  text-align: center;
}
</style>
