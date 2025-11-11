<script setup>
import { useRouter } from "vue-router";

const router = useRouter();

const goToMoments = () => {
  router.push("/moments");
};
</script>

<template>
  <div
    class="min-h-screen flex flex-col items-center justify-center relative overflow-hidden p-6 text-center"
  >
    <!-- Animated background -->
    <div
      class="absolute inset-0 bg-gradient-to-br from-purple-950 via-black to-purple-900"
    >
      <div
        class="absolute top-20 left-20 w-72 h-72 bg-purple-600/20 rounded-full blur-3xl animate-pulse"
      ></div>
      <div
        class="absolute bottom-20 right-20 w-96 h-96 bg-pink-600/10 rounded-full blur-3xl animate-pulse delay-1000"
      ></div>
    </div>

    <!-- Content -->
    <div class="relative z-10 max-w-2xl">
      <!-- Clock icon -->
      <div class="mb-8 inline-block">
        <n-icon size="80" color="#a78bfa" class="animate-spin-slow">
          <svg viewBox="0 0 24 24">
            <path
              fill="currentColor"
              d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
              :fill="none"
            />
          </svg>
        </n-icon>
      </div>

      <h1
        class="text-5xl md:text-6xl font-bold mb-6 bg-gradient-to-r from-purple-400 via-pink-400 to-purple-400 bg-clip-text text-transparent animate-gradient"
      >
        ยินดีต้อนรับเข้าสู่ MyWebMyBoo
      </h1>

      <n-space vertical :size="16" class="mb-10">
        <n-text class="text-lg md:text-xl text-purple-200 leading-relaxed">
          ขอบคุณที่อยู่เคียงข้างกันมาตลอด<br />
          ทุกช่วงเวลาของเรามีค่ามากจริง ๆ
        </n-text>
        <n-blockquote class="text-purple-300 text-sm italic">
          "เวลาที่ดีที่สุด คือเวลาที่เราอยู่ด้วยกัน"
        </n-blockquote>
      </n-space>

      <n-button
        type="primary"
        size="large"
        @click="goToMoments"
        class="custom-button"
        style="height: 56px; font-size: 1.125rem; padding: 0 2rem"
      >
        <template #icon>
          <svg
            class="w-5 h-5"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 5l7 7-7 7"
            />
          </svg>
        </template>
        ไปดูความทรงจำของเรา
      </n-button>
    </div>

    <!-- Floating hearts -->
    <div class="absolute inset-0 pointer-events-none overflow-hidden">
      <div class="heart-float" style="left: 10%; animation-delay: 0s">❤️</div>
      <div class="heart-float" style="left: 30%; animation-delay: 2s">💜</div>
      <div class="heart-float" style="left: 50%; animation-delay: 4s">🖤</div>
      <div class="heart-float" style="left: 70%; animation-delay: 1s">💗</div>
      <div class="heart-float" style="left: 90%; animation-delay: 3s">💕</div>
    </div>
  </div>
</template>

<style scoped>
.custom-button {
  background: linear-gradient(to right, #9333ea, #ec4899);
  border: none;
  transition: all 0.3s ease;
}

.custom-button:hover {
  transform: scale(1.05);
  box-shadow: 0 10px 25px -5px rgba(147, 51, 234, 0.5);
}

@keyframes spin-slow {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

.animate-spin-slow {
  animation: spin-slow 20s linear infinite;
}

@keyframes gradient {
  0%,
  100% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
}

.animate-gradient {
  background-size: 200% auto;
  animation: gradient 3s ease infinite;
}

@keyframes float-up {
  0% {
    transform: translateY(100vh) scale(0);
    opacity: 0;
  }
  10% {
    opacity: 1;
  }
  90% {
    opacity: 1;
  }
  100% {
    transform: translateY(-100vh) scale(1);
    opacity: 0;
  }
}

.heart-float {
  position: absolute;
  bottom: -50px;
  font-size: 2rem;
  animation: float-up 15s infinite ease-in;
}
</style>
