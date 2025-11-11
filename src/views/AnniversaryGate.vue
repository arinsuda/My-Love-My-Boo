<script setup>
import { useRouter } from "vue-router";
import AnniversaryUnlockForm from "@/components/AnniversaryUnlockForm.vue";

const router = useRouter();

const handleUnlocked = () => {
  router.push("/welcome");
};
</script>

<template>
  <div class="page-container">
    <!-- background overlay -->
    <div class="page-overlay"></div>

    <!-- Main Card -->
    <motion-div
      v-motion
      :initial="{ opacity: 0, scale: 0.9, y: 40 }"
      :enter="{
        opacity: 1,
        scale: 1,
        y: 0,
        transition: { duration: 0.7, easing: 'ease-out' },
      }"
      class="card-wrapper"
    >
      <motion-div
        v-motion
        :initial="{ opacity: 0, y: 20 }"
        :enter="{ opacity: 1, y: 0, transition: { delay: 0.2, duration: 0.6 } }"
        class="card-content"
      >
        <!-- Header -->
        <motion-div
          v-motion
          :initial="{ opacity: 0, y: -20 }"
          :enter="{
            opacity: 1,
            y: 0,
            transition: { delay: 0.3, duration: 0.6 },
          }"
          class="header-section"
        >
          <h1 class="main-title">WELCOME TO OUR<br />WORLD</h1>

          <p class="subtitle">
            ใส่ <span class="highlight">วันที่เราเริ่มคบกัน</span>
            <br class="mobile-break" />
            เพื่อปลดล็อกโลกของเราสองคน 💫
          </p>
        </motion-div>

        <!-- Form -->
        <motion-div
          v-motion
          :initial="{ opacity: 0, y: 30 }"
          :enter="{
            opacity: 1,
            y: 0,
            transition: { delay: 0.5, duration: 0.7 },
          }"
          class="form-section"
        >
          <AnniversaryUnlockForm @unlocked="handleUnlocked" />
        </motion-div>
      </motion-div>
    </motion-div>
  </div>
</template>

<style scoped>
.page-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 1.5rem;
  background: radial-gradient(
      circle at 20% 30%,
      rgba(168, 85, 247, 0.2) 0%,
      transparent 60%
    ),
    radial-gradient(
      circle at 80% 70%,
      rgba(236, 72, 153, 0.2) 0%,
      transparent 60%
    ),
    linear-gradient(135deg, #1e1b4b 0%, #4c1d95 50%, #831843 100%);
  position: relative;
  overflow: hidden;
}

.page-overlay {
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.35);
  backdrop-filter: blur(6px);
}

.card-wrapper {
  width: 100%;
  max-width: 480px;
  position: relative;
  z-index: 10;
}

.card-content {
  background: rgba(255, 255, 255, 0.08);
  backdrop-filter: blur(16px);
  border: 1.5px solid rgba(255, 255, 255, 0.2);
  border-radius: 32px;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.45);
  padding: 2.5rem 2rem;
  display: flex;
  flex-direction: column;
  gap: 2rem;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.card-content:hover {
  transform: translateY(-5px) scale(1.01);
  box-shadow: 0 25px 60px rgba(168, 85, 247, 0.25);
}

.header-section {
  text-align: center;
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.main-title {
  font-size: 2rem;
  font-weight: 700;
  color: white;
  line-height: 1.2;
  letter-spacing: 0.03em;
  text-shadow: 0 2px 20px rgba(168, 85, 247, 0.5);
}

.subtitle {
  font-size: 0.9375rem;
  color: rgba(255, 255, 255, 0.9);
  line-height: 1.6;
  letter-spacing: 0.01em;
}

.highlight {
  font-weight: 600;
  color: #ffc0e9;
  text-shadow: 0 0 8px rgba(236, 72, 153, 0.3);
}

@media (min-width: 640px) {
  .main-title {
    font-size: 2.5rem;
  }
  .subtitle {
    font-size: 1rem;
  }
}
</style>
