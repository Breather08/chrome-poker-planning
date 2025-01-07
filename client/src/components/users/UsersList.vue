<template>
  <div class="users-list">
    <article
      v-for="user in props.users"
      :key="user.id"
      :class="[
        'users-list__item',
        {
          'users-list__item--self': props.selfId === user.id,
        },
      ]"
    >
      <div class="users-list__item-name">
        {{ user.name }}
        <template v-if="props.selfId === user.id">(Me)</template>
      </div>
      <div
        :class="[
          'users-list__item-result',
          {
            'users-list__item-result--revealed': revealed,
          },
        ]"
      >
        <div v-if="revealed">
          {{ user.vote }}
        </div>
        <div v-else-if="user.voted">Voted</div>
        <div v-else>Not Voted</div>
      </div>
    </article>
  </div>
</template>

<script setup lang="ts">
import type { User } from "./types";

const props = defineProps<{
  users: User[];
  selfId: string;
  revealed: boolean;
}>();
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
    border-radius: 8px;
    padding: 8px 16px;
    box-shadow: 0px 2px 1px -1px rgba(0, 0, 0, 0.2),
      0px 1px 1px 0px rgba(0, 0, 0, 0.14), 0px 1px 3px 0px rgba(0, 0, 0, 0.12);

    &--self {
    }
  }
}
</style>
