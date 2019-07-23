<template>
  <div>
    <LandingPage v-if="isAuthenticated === false" />
    <Dashboard v-if="isAuthenticated === true" />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Getter } from 'vuex-class';

import Dashboard from '../components/Dashboard.vue';
import LandingPage from '../components/LandingPage.vue';

@Component({
  components: {
    Dashboard,
    LandingPage,
  },
})
export default class Home extends Vue {
  @Getter('isAuthenticated', { namespace: 'user' }) private isAuthenticated!: boolean;
}
</script>
