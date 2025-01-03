<template>
  <div v-if="isPlayable">
    <div class="cards-wrapper">
      <PokerCard content="1" @click="selectCard(1)" />
      <PokerCard content="2" @click="selectCard(2)" />
      <PokerCard content="3" @click="selectCard(3)" />
      <PokerCard content="5" @click="selectCard(5)" />
      <PokerCard content="8" @click="selectCard(8)" />
      <PokerCard content="13" @click="selectCard(13)" />
    </div>
    <div>
      <h3>Users</h3>
      <UsersList :users="users" />
    </div>
  </div>
  <template v-else>
    <input v-model="username" />
    <button @click="setUser">Join Room</button>
  </template>
  <div></div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import PokerCard from "./components/PokerCard.vue";
import UsersList from "./components/users/UsersList.vue";
import type { User } from "./components/users/types";

const WS_API_PATH = "ws://localhost:8080/ws";
const roomId = window.location.href.replace(/.+\/((\w|-)+)(\?.+)?$/g, "$1");
const username = ref(localStorage.getItem("username") || "");
const id = ref(localStorage.getItem("user_id") || "");
const isPlayable = ref(false);
const users = ref<User[]>([]);

const socket = new WebSocket(`${WS_API_PATH}?room_id=${roomId}`);

socket.addEventListener("error", (e) => {
  console.error("WebSocket error:", e);
});

socket.addEventListener("message", (e) => {
  console.log("Message Received:", e.data);
  const payload = JSON.parse(e.data);
  switch (payload.type) {
    case "player_enter":
      handlePlayerEnter(payload.body);
      break;
    case "pick_card":
      handlePickCard(payload.body);
      break;
    case "players_list":
      handlePlayersList(payload.body);
      break;
    default:
      console.error("Unhandled message type");
  }
});

socket.addEventListener("open", () => {
  console.log("Connection open", id.value);

  if (username.value && id.value) {
    requestJoin();
    requestPlayersList();
  }
});

function requestPlayersList() {
  socket.send(
    JSON.stringify({
      type: "players",
    })
  );
}

function requestJoin() {
  socket.send(
    JSON.stringify({
      type: "join",
      payload: {
        id: id.value,
        username: username.value,
      },
    })
  );
}

function handlePlayerEnter(userId: string) {
  if (!userId || typeof userId !== "string") return;

  const localId = localStorage.getItem("user_id");
  if (localId) {
    id.value = localId;
  } else {
    id.value = userId;
    localStorage.setItem("user_id", userId);
  }

  isPlayable.value = true;
}

function handlePickCard(payload: Record<string, number>) {
  if (!(payload instanceof Object) || !(id.value in payload)) return;

  const picked = (payload as Record<string, number>)[id.value];
}

function handlePlayersList(payload: Record<string, User>) {
  users.value = Object.values(payload);
}

function setUser() {
  localStorage.setItem("username", username.value);
  requestJoin();
}

function selectCard(num: number) {
  socket.send(
    JSON.stringify({
      type: "pick",
      payload: {
        cardId: num.toString(),
      },
    })
  );
}
</script>

<style scoped lang="scss">
.cards-wrapper {
  width: 100%;
  flex: 1;
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 8px;
}
</style>
