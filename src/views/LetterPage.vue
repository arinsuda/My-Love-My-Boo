<script setup>
import { useRouter } from "vue-router";

const router = useRouter();
const backToStart = () => router.push("/");
</script>

<template>
  <div
    class="min-h-screen flex flex-col justify-center items-center relative overflow-hidden px-6 py-12"
  >
    <!-- Animated Background -->
    <div
      class="absolute inset-0 bg-gradient-to-br from-black via-purple-950 to-pink-950"
    >
      <div
        class="absolute top-20 left-20 w-96 h-96 bg-purple-600/20 rounded-full blur-3xl animate-pulse"
      ></div>
      <div
        class="absolute bottom-20 right-20 w-96 h-96 bg-pink-600/20 rounded-full blur-3xl animate-pulse"
        style="animation-delay: 1s"
      ></div>
      <div
        class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-full h-full bg-[radial-gradient(ellipse_at_center,_var(--tw-gradient-stops))] from-purple-900/20 via-transparent to-transparent"
      ></div>
    </div>

    <!-- Floating particles -->
    <div class="absolute inset-0 overflow-hidden pointer-events-none">
      <div class="particle" style="left: 10%; animation-delay: 0s"></div>
      <div class="particle" style="left: 30%; animation-delay: 2s"></div>
      <div class="particle" style="left: 50%; animation-delay: 4s"></div>
      <div class="particle" style="left: 70%; animation-delay: 1s"></div>
      <div class="particle" style="left: 90%; animation-delay: 3s"></div>
    </div>

    <!-- Content -->
    <div class="relative z-10 max-w-3xl w-full text-center">
      <!-- Envelope Icon -->
      <div class="mb-8 inline-block animate-bounce-slow">
        <n-icon size="80" color="#a78bfa">
          <svg viewBox="0 0 24 24">
            <path
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
            />
          </svg>
        </n-icon>
      </div>

      <!-- Title -->
      <h1
        class="text-4xl md:text-5xl font-bold mb-8 bg-gradient-to-r from-purple-400 via-pink-400 to-purple-400 bg-clip-text text-transparent animate-gradient"
      >
        จดหมายจากฉันถึงเธอ
      </h1>

      <!-- Letter Card -->
      <n-card class="letter-card relative group">
        <!-- Decorative corners -->
        <div
          class="absolute top-0 left-0 w-16 h-16 border-t-2 border-l-2 border-purple-400/50 rounded-tl-3xl"
        ></div>
        <div
          class="absolute top-0 right-0 w-16 h-16 border-t-2 border-r-2 border-purple-400/50 rounded-tr-3xl"
        ></div>
        <div
          class="absolute bottom-0 left-0 w-16 h-16 border-b-2 border-l-2 border-purple-400/50 rounded-bl-3xl"
        ></div>
        <div
          class="absolute bottom-0 right-0 w-16 h-16 border-b-2 border-r-2 border-purple-400/50 rounded-br-3xl"
        ></div>

        <!-- Letter content -->
        <div
          class="space-y-6 text-left text-purple-100 text-lg leading-relaxed"
        >
          <n-p
            class="first-letter:text-5xl first-letter:font-bold first-letter:text-purple-400 first-letter:mr-1 first-letter:float-left"
          >
            ขอบคุณที่เข้ามาในชีวิตของฉันนะ ทุก ๆ
            วันมีความหมายขึ้นเพราะมีเธออยู่ข้าง ๆ ตั้งแต่วันแรกจนถึงวันนี้
            ฉันก็ยังรู้สึกเหมือนเดิม และมากขึ้นทุกวันเลย
          </n-p>

          <n-blockquote class="border-l-4 border-purple-500/50 pl-4">
            ไม่ว่าอนาคตจะเป็นยังไง ขอให้รู้ไว้ว่าฉันจะอยู่ตรงนี้เสมอ เพราะเธอคือ
            "ของขวัญ" ที่ดีที่สุดที่ฉันเคยได้รับ
          </n-blockquote>

          <div class="pt-6 text-center">
            <n-text
              class="text-pink-300 text-xl font-semibold italic block mb-2"
            >
              รักที่สุดในโลกเลยนะ
            </n-text>
            <n-space justify="center" align="center" :size="8">
              <span class="text-2xl">💗</span>
              <n-text class="text-purple-300">
                จากฉัน ผู้โชคดีที่มีเธออยู่ข้าง ๆ
              </n-text>
              <span class="text-2xl">🌙</span>
            </n-space>
          </div>
        </div>

        <!-- Heart decoration -->
        <div class="absolute -top-4 -right-4 text-5xl animate-pulse">💜</div>
        <div
          class="absolute -bottom-4 -left-4 text-4xl animate-pulse"
          style="animation-delay: 0.5s"
        >
          🖤
        </div>
      </n-card>

      <!-- Back Button -->
      <n-button
        type="primary"
        size="large"
        @click="backToStart"
        class="custom-button mt-12"
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
              d="M15 19l-7-7 7-7"
            />
          </svg>
        </template>
        กลับไปหน้าแรก
      </n-button>
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

:deep(.letter-card) {
  background: linear-gradient(
    to bottom right,
    rgba(88, 28, 135, 0.6),
    rgba(0, 0, 0, 0.6)
  );
  border: 1px solid rgba(139, 92, 246, 0.3);
  backdrop-filter: blur(16px);
  box-shadow: 0 20px 60px -15px rgba(0, 0, 0, 0.8);
}

:deep(.letter-card:hover) {
  border-color: rgba(139, 92, 246, 0.4);
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

@keyframes bounce-slow {
  0%,
  100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
}

.animate-bounce-slow {
  animation: bounce-slow 2s ease-in-out infinite;
}

@keyframes float {
  0%,
  100% {
    transform: translateY(0) translateX(0);
    opacity: 0;
  }
  10% {
    opacity: 1;
  }
  90% {
    opacity: 1;
  }
  100% {
    transform: translateY(-100vh) translateX(50px);
    opacity: 0;
  }
}

.particle {
  position: absolute;
  bottom: -50px;
  width: 4px;
  height: 4px;
  background: linear-gradient(45deg, #a78bfa, #ec4899);
  border-radius: 50%;
  animation: float 15s infinite ease-in;
}
</style>
