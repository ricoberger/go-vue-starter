<template>
  <v-container fluid fill-height>
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md4>
        <v-card class="elevation-12">
          <v-toolbar dark color="primary">
            <v-toolbar-title>Login</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>
          <v-card-text>
            <v-form ref="form">
              <EmailTextField v-model="email" />
              <PasswordTextField v-model="password" />
              <Alert v-model="error" type="error" />
            </v-form>
          </v-card-text>
          <v-card-actions>
            <router-link to="/forgot-password">Forgotton your password?</router-link>
            <v-spacer></v-spacer>
            <v-btn color="primary" :loading="loading" :disabled="loading" @click="doLogin">Login</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Action } from 'vuex-class';

import Alert from '../components/Alert.vue';
import EmailTextField from '../components/EmailTextField.vue';
import PasswordTextField from '../components/PasswordTextField.vue';

@Component({
  components: {
    Alert,
    EmailTextField,
    PasswordTextField,
  },
})
export default class Login extends Vue {
  @Action('login', { namespace: 'user' }) private login: any;

  private email: string = '';
  private password: string = '';
  private loading: boolean = false;
  private error: string = '';

  private doLogin() {
    if ((this.$refs.form as HTMLFormElement).validate()) {
      this.loading = true;

      this.login({email: this.email, password: this.password}).then(() => {
        this.loading = false;
        if (this.$route.query.next) {
          this.$router.replace({ path: this.$route.query.next as string });
        } else {
          this.$router.replace({ path: '/' });
        }
      }).catch((err: any) => {
        this.error = err.message;
        this.loading = false;
      });
    }
  }
}
</script>
