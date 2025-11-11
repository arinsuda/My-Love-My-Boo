<script setup>
import { darkTheme } from "naive-ui";
import { computed } from "vue";

const customTheme = computed(() => ({
  common: {
    primaryColor: "#b794f6",
    primaryColorHover: "#c4b5fd",
    primaryColorPressed: "#a78bfa",
    primaryColorSuppl: "#ec4899",
    errorColor: "#f87171",
    successColor: "#34d399",
    warningColor: "#fbbf24",
    infoColor: "#93c5fd",
    borderRadius: "14px",
  },
  Button: {
    borderRadiusMedium: "14px",
    borderRadiusLarge: "18px",
  },
  Input: {
    borderRadius: "14px",
  },
}));
</script>

<template>
  <n-config-provider :theme="darkTheme" :theme-overrides="customTheme">
    <n-message-provider>
      <div
        class="min-h-screen bg-kurumi-dark text-white font-sans flex flex-col"
      >
        <router-view v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </div>
    </n-message-provider>
  </n-config-provider>
</template>

<style>
:root {
  --kurumi-dark: #0a0412;
  --kurumi-purple-dark: #1b0e2b;
  --kurumi-purple: #2b1743;
  --kurumi-purple-light: #50316b;
  --kurumi-accent: #b794f6;
  --kurumi-accent-light: #c4b5fd;
  --kurumi-pink: #f472b6;
  --kurumi-pink-light: #f9a8d4;
  --kurumi-red: #fb7185;
  --kurumi-gradient: linear-gradient(
    135deg,
    #1a0b28 0%,
    #2b1743 40%,
    #4c2886 100%
  );
}

/* ⚙️ Reset */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  margin: 0;
  background: var(--kurumi-gradient);
  font-family: "Inter", "Noto Sans Thai", system-ui, sans-serif;
  color: #f5f3ff;
  overflow-x: hidden;
}

/* 🪄 Background Helper Class */
.bg-kurumi-dark {
  background: var(--kurumi-gradient);
}

/* ✨ Smooth Fade Transition */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* 🌸 Custom Scrollbar */
::-webkit-scrollbar {
  width: 8px;
}
::-webkit-scrollbar-track {
  background: var(--kurumi-purple-dark);
}
::-webkit-scrollbar-thumb {
  background: var(--kurumi-accent);
  border-radius: 4px;
}
::-webkit-scrollbar-thumb:hover {
  background: var(--kurumi-accent-light);
}
</style>
