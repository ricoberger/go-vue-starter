<template>
  <v-text-field v-model="email" :rules="emailRules" prepend-icon="email" name="email" label="Email address" type="email" required></v-text-field>
</template>

<script lang="ts">
import { Component, Emit, Prop, Vue } from 'vue-property-decorator';

@Component
export default class EmailTextField extends Vue {
  @Prop(String) public value!: string;

  private emailRules = [
    (v: string) => !!v || 'Email address is required',
    (v: string) => /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(v) || 'Email address must be valid',
  ];

  private get email(): string {
    return this.value;
  }

  private set email(newVal: string) {
    this.changeValue(newVal);
  }

  @Emit('input')
  public changeValue(val: string) { /* */ }
}
</script>
