<template>
  <v-app>
    <Sidebar v-if="isAuthenticated === true" v-model="drawer" />
    <Header v-model="drawer" />

    <v-content>

      <router-view/>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Getter } from 'vuex-class';

import Header from './components/Header.vue';
import Sidebar from './components/Sidebar.vue';

@Component({
  components: {
    Header,
    Sidebar,
  },
})
export default class App extends Vue {
  @Getter('isAuthenticated', { namespace: 'user' }) private isAuthenticated!: boolean;

  private drawer: boolean = false;
}
</script>
