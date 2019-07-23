<template>
  <v-container fluid fill-height>
    <v-layout align-content-center justify-center>
      <v-flex xs12 sm10 md8>
        <v-layout row wrap>
          <v-flex xs12>
            <h1>Profile</h1>
          </v-flex>
        </v-layout>

        <v-layout row wrap>
          <v-flex xs12>
            <v-form ref="form">
              <v-text-field v-model="name" :rules="nameRules" prepend-icon="person" name="name" label="Name" type="text"></v-text-field>
              <EmailTextField v-model="email" />
              <Alert v-model="success" type="success" />
              <Alert v-model="error" type="error" />
              <Alert v-model="info" type="info" />
            </v-form>
          </v-flex>
        </v-layout>

        <v-layout row wrap>
          <v-flex xs12 class="right-text">
            <v-btn color="primary" :loading="loading" :disabled="loading" @click="saveProfile">Save profile</v-btn>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Action, State } from 'vuex-class';

import Alert from '../components/Alert.vue';
import EmailTextField from '../components/EmailTextField.vue';
import { IUserState } from '../store/modules/user';

@Component({
  components: {
    Alert,
    EmailTextField,
  },
})
export default class Profile extends Vue {
  @State('user') private user!: IUserState;
  @Action('save', { namespace: 'user' }) private save: any;

  private name: string = '';
  private email: string = '';
  private loading: boolean = false;
  private success: string = '';
  private error: string = '';
  private info: string = '';

  private nameRules = [
    (v: string) => !!v || 'Name is required',
  ];

  private mounted() {
    this.name = this.user.name ? this.user.name : '';
    this.email = this.user.email ? this.user.email : '';
  }

  private saveProfile() {
    let tmpInfo = '';
    if (this.email !== this.user.email) {
      tmpInfo = 'Please verify your new email address';
    }

    if ((this.$refs.form as HTMLFormElement).validate()) {
      this.loading = true;

      this.save({name: this.name, email: this.email}).then(() => {
        this.success = 'Changes were saved';
        this.info = tmpInfo;
        this.loading = false;
      }).catch((err: any) => {
        this.error = err.message;
        this.loading = false;
      });
    }
  }
}
</script>

<style scoped>
.right-text {
  text-align: right;
}
</style>
