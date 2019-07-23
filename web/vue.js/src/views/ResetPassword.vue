<template>
  <v-container fluid fill-height>
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md4>
        <v-card class="elevation-12">
          <v-toolbar dark color="primary">
            <v-toolbar-title>Reset password</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>
          <v-card-text>
            <v-form ref="form">
              <PasswordTextField v-model="password" />
              <Alert v-model="error" type="error" />
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" :loading="loading" :disabled="loading" @click="setPassword">Set new password</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import Alert from '../components/Alert.vue';
import PasswordTextField from '../components/PasswordTextField.vue';
import { API_ENDPOINT } from '../constants';

@Component({
  components: {
    Alert,
    PasswordTextField,
  },
})
export default class ResetPassword extends Vue {
  private password: string = '';
  private loading: boolean = false;
  private error: string = '';

  private async setPassword() {
    if ((this.$refs.form as HTMLFormElement).validate()) {
      this.error = '';
      this.loading = true;

      try {
        const response = await fetch(API_ENDPOINT + '/api/v1/account/password', {
          body: JSON.stringify({
            id: this.$route.params.id,
            resetPasswordToken: this.$route.params.token,
            password: this.password,
          }),
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
          },
          method: 'put',
        });

        const json = await response.json();

        if (response.status >= 200 && response.status < 300) {
          this.loading = false;
          this.$router.replace('/login');
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
