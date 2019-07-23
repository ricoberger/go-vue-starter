<template>
  <v-container fluid fill-height>
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md4>
        <v-card class="elevation-12">
          <v-toolbar dark color="primary">
            <v-toolbar-title>Forgot password</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>
          <v-card-text>
            <v-form ref="form">
              <EmailTextField v-model="email" />
              <Alert v-model="success" type="success" />
              <Alert v-model="error" type="error" />
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" :loading="loading" :disabled="loading" @click="resetPassword">Reset password</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import Alert from '../components/Alert.vue';
import EmailTextField from '../components/EmailTextField.vue';
import { API_ENDPOINT } from '../constants';

@Component({
  components: {
    Alert,
    EmailTextField,
  },
})
export default class ForgotPassword extends Vue {
  private email: string = '';
  private loading: boolean = false;
  private success: string = '';
  private error: string = '';

  private async resetPassword() {
    if ((this.$refs.form as HTMLFormElement).validate()) {
      this.error = '';
      this.success = '';
      this.loading = true;

      try {
        const response = await fetch(API_ENDPOINT + '/api/v1/account/password', {
          body: JSON.stringify({
            email: this.email,
          }),
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
          },
          method: 'post',
        });

        const json = await response.json();

        if (response.status >= 200 && response.status < 300) {
          this.success = 'We have send you an email to reset your password.';
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
}
</script>
