<template>
  <div
    class="ds-icon"
    :class="[props.color]"
    :style="{ maskImage: `url(${src})` }"
    @click="$emit('click', $event)"
  ></div>
</template>

<script setup lang="ts">
import { computed } from "vue";

interface IProps {
  name: string;
  color?: string;
  size?: number;
}

const src = computed(() => "/vite.svg");
const props = defineProps<IProps>();
</script>

<style lang="scss">
.ds-icon {
  display: inline-block;
  -webkit-mask-repeat: no-repeat;
  mask-repeat: no-repeat;
  -webkit-mask-position: center;
  mask-position: center;
}
</style>
