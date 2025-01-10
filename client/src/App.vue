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
    <div class="users-header">
      <h3>Users</h3>
      <CBtn variant="ghost" @click="resetVotes">Reset</CBtn>
    </div>
    <UsersList :users="users" :self-id="id" :revealed="isRevealed" />
    <CBtn @click="revealAnswers" class="reveal-button" full
      >Reveal answers</CBtn
    >
  </div>
  <template v-else>
    <CInput v-model="username" />
    <CBtn @click="setUser">Join Room</CBtn>
  </template>
</template>

<script setup lang="ts">
import { ref } from "vue";
import PokerCard from "./components/PokerCard.vue";
import UsersList from "./components/users/UsersList.vue";
import type { User } from "./components/users/types";
import CBtn from "./components/shared/btn/CBtn.vue";
import CInput from "./components/shared/input/CInput.vue";

enum SendCommand {
  Join = "join",
  Pick = "pick",
  Show = "show",
  Reset = "reset",
  GetState = "get_state",
}

enum ReceiveCommand {
  PlayerEnter = "player_enter",
  PickCard = "pick_card",
  GameState = "game_state",
}

interface Player {
  id: string;
  name: string;
  voted: boolean;
  vote: string;
}

interface GameState {
  revealed: boolean;
  players: Player[];
}

interface ResponsePayload {
  type: ReceiveCommand;
  body: any;
}

const WS_API_PATH = "ws://192.168.1.154:8080/ws";
const roomId = window.location.href.replace(/.+\/((\w|-)+)(\?.+)?$/g, "$1");
const username = ref(localStorage.getItem("username") || "");
const id = ref(localStorage.getItem("user_id") || "");
const isPlayable = ref(false);
const isRevealed = ref(false);
const users = ref<User[]>([]);

const socket = new WebSocket(`${WS_API_PATH}?room_id=${roomId}`);

socket.addEventListener("error", (e) => {
  console.error("WebSocket error:", e);
});

socket.addEventListener("message", (e) => {
  try {
    const payload = JSON.parse(e.data) as ResponsePayload;
    console.log("Message Received:", payload);
    switch (payload.type) {
      case "player_enter":
        handlePlayerEnter(payload.body);
        break;
      case "pick_card":
        handlePickCard(payload.body);
        break;
      case "game_state":
        handleGameState(payload.body);
        break;
      default:
        console.error("Unhandled message type", payload.type);
    }
  } catch (e) {
    console.error("Parse error", e);
  }
});

socket.addEventListener("open", () => {
  console.log("Connection open", id.value);
  if (username.value && id.value) {
    requestJoin();
    requestGameState();
  }
});

function resetVotes() {
  sendMessage(SendCommand.Reset);
}

function revealAnswers() {
  sendMessage(SendCommand.Show);
}

function requestGameState() {
  sendMessage(SendCommand.GetState);
}

function requestJoin() {
  sendMessage(SendCommand.Join, {
    id: id.value,
    username: username.value,
  });
}

function sendMessage<T = any>(type: SendCommand, payload?: T) {
  socket.send(
    JSON.stringify({
      type,
      payload,
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
  console.log("Card Picked", payload);
}

function handleGameState(payload: GameState) {
  isRevealed.value = payload.revealed;
  users.value = payload.players;
}

function setUser() {
  localStorage.setItem("username", username.value);
  requestJoin();
  requestGameState();
}

function selectCard(num: number) {
  sendMessage(SendCommand.Pick, {
    cardId: num.toString(),
  });
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

.users-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.reveal-button {
  margin-top: 24px;
}
</style>
