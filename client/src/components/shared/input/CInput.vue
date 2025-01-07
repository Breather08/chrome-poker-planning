<template>
  <label>
    <span>{{ label }}</span>
    <input
      v-model="model"
      v-bind="$attrs"
      :type="props.type"
      :placeholder="props.placeholder"
    />
  </label>
</template>

<script setup lang="ts">
interface IProps {
  label?: string;
  placeholder?: string;
  type?: string;
}

const props = withDefaults(defineProps<IProps>(), {
  label: "",
  placeholder: "",
  type: "text",
});

const model = defineModel();
</script>
