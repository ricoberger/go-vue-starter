<template>
  <v-alert v-model="isVisible" :type="type" transition="scale-transition" dismissible outline>{{value}}</v-alert>
</template>

<script lang="ts">
import { Component, Emit, Prop, Vue } from 'vue-property-decorator';

@Component
export default class Alert extends Vue {
  @Prop(String) public value!: string;
  @Prop(String) public type!: string;

  private get isVisible() {
    if (this.value !== '') {
      return true;
    } else {
      return false;
    }
  }

  private set isVisible(newVal: boolean) {
    if (newVal === false) {
      this.changeValue('');
    }
  }

  @Emit('input')
  public changeValue(val: string) { /* */ }
}
</script>
