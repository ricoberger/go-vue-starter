<template>
  <v-navigation-drawer :clipped="$vuetify.breakpoint.lgAndUp" v-model="drawer" fixed app disable-resize-watcher>
    <v-img :aspect-ratio="16/9" :src="require('../assets/sidebar.png')">
      <v-layout pa-2 column fill-height class="lightbox white--text">
        <v-spacer></v-spacer>
        <v-flex shrink>
          <v-avatar :tile="false" size="32px" color="primary">
            <img :src="getAvatar" :alt="getName">
          </v-avatar>
          <div class="subheading">{{user.name}}</div>
          <div class="body-1">{{user.email}}</div>
        </v-flex>
      </v-layout>
    </v-img>

    <v-list dense>
      <v-list-tile @click="goTo('/')">
        <v-list-tile-action>
          <v-icon>home</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Home</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>

      <v-list-tile @click="goTo('/profile')">
        <v-list-tile-action>
          <v-icon>person</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Profile</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>

      <v-divider></v-divider>

      <v-list-tile @click="doLogout">
        <v-list-tile-action>
          <v-icon>power_settings_new</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Logout</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
  </v-navigation-drawer>
</template>

<script lang="ts">
import { Component, Emit, Prop, Vue } from 'vue-property-decorator';
import { Action, Getter, State } from 'vuex-class';

import { IUserState } from '../store/modules/user';

@Component
export default class Sidebar extends Vue {
  @Prop(Boolean) public value!: boolean;

  @Action('logout', { namespace: 'user' }) private logout: any;
  @Getter('getName', { namespace: 'user' }) private getName!: string;
  @Getter('getAvatar', { namespace: 'user' }) private getAvatar!: string;
  @State('user') private user!: IUserState;

  private items = [
    { title: 'Home', icon: 'dashboard' },
    { title: 'About', icon: 'question_answer' },
  ];

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
