<template>
  <button
    :class="[
      'btn',
      `btn--${props.variant} btn--${props.size}`,
      {
        'btn--full': props.full,
      },
    ]"
  >
    <slot></slot>
  </button>
</template>

<script setup lang="ts">
import type { ButtonSize, ButtonVariant } from "./types";

interface IProps {
  variant?: ButtonVariant;
  size?: ButtonSize;
  full?: boolean;
}

const props = withDefaults(defineProps<IProps>(), {
  variant: "primary",
  size: "m",
  full: false,
});
</script>

<style lang="scss">
.btn {
  border: none;
  outline: none;
  color: white;
  border-radius: 12px;
  font-weight: 500;
  font-size: 18px;
  cursor: pointer;
  transition: background 0.1s ease-out;

  &--primary {
    background-color: #5347dd;

    &:hover {
      background-color: #4b04e2;
    }
  }
  &--secondary {
    background-color: #ff077e;

    &:hover {
      background-color: #b20558;
    }
  }
  &--outline {
    background-color: transparent;
    border: 1px solid #5347dd;
    color: #5347dd;

    &:hover {
      border: 1px solid #4b04e2;
      color: #4b04e2;
      background-color: rgba(83, 71, 221, 0.1);
    }
  }
  &--ghost {
    color: #5347dd;
    background-color: transparent;

    &:hover {
      color: #4b04e2;
      background-color: rgba(83, 71, 221, 0.1);
    }
  }

  &--s {
    padding: 8px 16px;
  }
  &--m {
    padding: 12px 16px;
  }
  &--l {
    padding: 16px 16px;
  }

  &--full {
    width: 100%;
  }
}
</style>
