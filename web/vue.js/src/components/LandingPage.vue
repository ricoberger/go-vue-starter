<template>
  <div>
    <v-parallax :src="require('../assets/vbanner.jpg')">
      <v-layout class="center-text" align-center column justify-center>
        <h1 class="display-2 font-weight-thin mb-3">Go &#9829; Vue Starter</h1>
        <v-btn to="/signup" outline color="white">Sign up</v-btn>
      </v-layout>
    </v-parallax>

    <v-layout class="padding-top-bottom" justify-center>
      <v-flex class="center-text" xs12 sm10>
        <h1>A starter project using Go, Vue, Vuetify and MongoDB</h1>
      </v-flex>
    </v-layout>

    <v-layout justify-center>
      <v-flex xs12 sm10>
        <v-divider></v-divider>
      </v-flex>
    </v-layout>

    <v-layout class="padding-top-bottom" justify-center>
      <v-flex xs12 sm10>
        <v-container fluid grid-list-md>
          <v-layout row wrap>
            <v-flex v-for="image in images" xs12 sm6 lg3 :key="image.title">
              <v-card :href="image.link" target="_blank">
                <v-img class="center-image" :src="image.src" aspect-ratio="1" height="200" width="200" />
              </v-card>
            </v-flex>
          </v-layout>
        </v-container>
      </v-flex>
    </v-layout>

    <v-footer dark height="auto">
      <v-card flat tile class="primary white--text text-xs-center">
        <v-card-text>
          <v-btn v-for="icon in icons" :key="icon.icon" class="mx-3 white--text" :href="icon.link" target="_blank" icon>
            <v-icon size="24px">{{icon.icon}}</v-icon>
          </v-btn>
        </v-card-text>
        <v-card-text class="white--text pt-0">
          <LoremIpsum />
        </v-card-text>
        <v-divider></v-divider>
        <v-card-text class="white--text">
          &copy;2019 — <strong>Rico Berger</strong>
        </v-card-text>
      </v-card>
    </v-footer>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import LoremIpsum from '../components/LoremIpsum.vue';

@Component({
  components: {
    LoremIpsum,
  },
})
export default class LandingPage extends Vue {
  private images = [
    { title: 'Go', src: require('../assets/go.png'), link: 'https://golang.org' },
    { title: 'Vue', src: require('../assets/vue.png'), link: 'https://vuejs.org' },
    { title: 'Vuetify', src: require('../assets/vuetify.png'), link: 'https://vuetifyjs.com/en/' },
    { title: 'MongoDB', src: require('../assets/mongodb.png'), link: 'https://www.mongodb.com' },
  ];

  private icons = [
    { icon: 'fab fa-linkedin-in', link: 'https://www.linkedin.com/in/ricoberger/' },
    { icon: 'fab fa-xing', link: 'https://www.xing.com/profile/Rico_Berger5' },
    { icon: 'fab fa-twitter', link: 'https://twitter.com/rico_berger' },
    { icon: 'fab fa-github', link: 'https://github.com/ricoberger' },
    { icon: 'fab fa-medium', link: 'https://medium.com/@ricoberger' },
  ];
}
</script>

<style scoped>
.center-text {
  text-align: center;
}

.center-image {
  display: block;
  margin-left: auto;
  margin-right: auto;
}

.padding-top-bottom {
  padding-top: 50px;
  padding-bottom: 50px;
}
</style>
