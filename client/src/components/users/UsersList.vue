<template>
  <div class="users-list">
    <article
      v-for="user in props.users"
      :key="user.id"
      class="users-list__item"
    >
      <div class="users-list__item-name">{{ user.name }}</div>
      <div
        :class="[
          'users-list__item-result',
          {
            'users-list__item-result--voted': user.voted,
          },
        ]"
      >
        {{ user.vote }}
      </div>
    </article>
  </div>
</template>

<script setup lang="ts">
import type { User } from "./types";

const props = defineProps<{ users: User[] }>();
</script>

<style lang="scss">
.users-list {
  display: flex;
  flex-direction: column;
  gap: 4px;

  &__item {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
}
</style>
