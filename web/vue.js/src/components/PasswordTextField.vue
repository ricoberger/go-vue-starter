<template>
  <v-text-field v-model="password" :rules="passwordRules" prepend-icon="lock" name="password" label="Password" type="password" required></v-text-field>
</template>

<script lang="ts">
import { Component, Emit, Prop, Vue } from 'vue-property-decorator';

@Component
export default class PasswordTextField extends Vue {
  @Prop(String) public value!: string;

  private passwordRules = [
    (v: string) => !!v || 'Password is required',
    (v: string) => /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/.test(v) || 'Password must be eight characters long, with at least one letter and one number',
  ];

  private get password(): string {
    return this.value;
  }

  private set password(newVal: string) {
    this.changeValue(newVal);
  }

  @Emit('input')
  public changeValue(val: string) { /* */ }
}
</script>
