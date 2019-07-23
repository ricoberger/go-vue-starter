<template>
  <v-toolbar app color="primary" :clipped-left="$vuetify.breakpoint.lgAndUp" dark>
    <v-toolbar-title class="headline text-uppercase header" @click="goTo('/')">
      <v-toolbar-side-icon v-if="isAuthenticated === true" @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <span v-bind:class="{ 'hidden-xs-only': isAuthenticated }">Go &#9829; Vue</span>
      <span class="font-weight-light " v-bind:class="{ 'hidden-xs-only': isAuthenticated}"> STARTER</span>
    </v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn v-if="isAuthenticated === false" class="hidden-xs-only" to="/login" flat>Login</v-btn>
    <v-btn v-if="isAuthenticated === false" class="hidden-xs-only" to="/signup" flat>Sign up</v-btn>

    <v-btn v-if="isAuthenticated === true" @click="goTo('/')" offset-y icon>
      <v-icon>apps</v-icon>
    </v-btn>
    <v-btn v-if="isAuthenticated === true" @click="goTo('/')" offset-y icon>
      <v-icon>notifications</v-icon>
    </v-btn>
  </v-toolbar>
</template>

<script lang="ts">
import { Component, Emit, Prop, Vue } from 'vue-property-decorator';
import { Action, Getter } from 'vuex-class';

@Component
export default class Header extends Vue {
  @Prop(Boolean) public value!: boolean;

  @Action('logout', { namespace: 'user' }) private logout: any;
  @Getter('isAuthenticated', { namespace: 'user' }) private isAuthenticated!: boolean;

  private get drawer(): boolean {
    return this.value;
  }

  private set drawer(newVal: boolean) {
    this.changeValue(newVal);
  }

  @Emit('input')
  public changeValue(val: boolean) { /* */ }

  private goTo(path: string) {
    this.$router.push({ path });
  }

  private doLogout() {
    this.drawer = false;
    this.logout();
    this.$router.push({ path: '/' });
  }
}
</script>

<style scoped>
.header {
  cursor: pointer;
}
</style>
