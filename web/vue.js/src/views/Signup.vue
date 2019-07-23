<template>
  <v-container fluid fill-height>
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md4>
        <v-card class="elevation-12">
          <v-toolbar dark color="primary">
            <v-toolbar-title>Sign up</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>
          <v-card-text>
            <v-form ref="form">
              <v-text-field v-model="name" :rules="nameRules" prepend-icon="person" name="name" label="Name" type="text"></v-text-field>
              <EmailTextField v-model="email" />
              <PasswordTextField v-model="password" />
              <Alert v-model="error" type="error" />
            </v-form>
          </v-card-text>
          <v-card-actions>
            <router-link to="/login">Already have an account?</router-link>
            <v-spacer></v-spacer>
            <v-btn color="primary" :loading="loading" :disabled="loading" @click="doSignup">Sign up</v-btn>
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
export default class Signup extends Vue {
  @Action('signup', { namespace: 'user' }) private signup: any;

  private name: string = '';
  private email: string = '';
  private password: string = '';
  private loading: boolean = false;
  private error: string = '';

  private nameRules = [
    (v: string) => !!v || 'Name is required',
  ];

  private doSignup() {
    if ((this.$refs.form as HTMLFormElement).validate()) {
      this.loading = true;

      this.signup({name: this.name, email: this.email, password: this.password}).then(() => {
        this.loading = false;
        this.$router.push({ path: '/welcome' });
      }).catch((err: any) => {
        this.error = err.message;
        this.loading = false;
      });
    }
  }
}
</script>
