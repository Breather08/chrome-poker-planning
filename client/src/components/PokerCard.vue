<script setup lang="ts">
interface IProps {
  content: string;
}

const props = defineProps<IProps>();
</script>

<template>
  <button class="poker-card">
    <div class="poker-card__content">
      {{ props.content }}
    </div>
  </button>
</template>

<style lang="scss">
.poker-card {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 70px;
  height: 100px;
  border-radius: 8px;
  font-size: 18px;
  box-shadow: 0 2px 1px -1px #0003, 0 1px 1px #00000024, 0 1px 3px #0000001f;
  color: black;
  background-color: #fff;
  transition: background-color 0.05s ease-out;
  outline: none;
  border: none;
  cursor: pointer;

  &__content {
    font-weight: bold;
  }

  &:active {
    box-shadow: none;
    background-color: #f0f0f0;
  }

  &:hover {
    background-color: #f0f0f0;
  }
}
</style>
