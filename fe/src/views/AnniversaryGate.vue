<script setup>
  import { useRouter } from "vue-router"
  import { useMotion } from "@vueuse/motion"
  import AnniversaryUnlockForm from "@/components/AnniversaryUnlockForm.vue"

  const router = useRouter()

  const handleUnlocked = () => {
    router.push("/welcome")
  }

  // motion config ต่าง ๆ
  const cardMotion = {
    initial: { opacity: 0, scale: 0.9, y: 40 },
    enter: {
      opacity: 1,
      scale: 1,
      y: 0,
      transition: { duration: 0.7, easing: "ease-out" },
    },
  }

  const innerCardMotion = {
    initial: { opacity: 0, y: 24 },
    enter: {
      opacity: 1,
      y: 0,
      transition: { delay: 0.15, duration: 0.6 },
    },
  }

  const headerMotion = {
    initial: { opacity: 0, y: -20 },
    enter: {
      opacity: 1,
      y: 0,
      transition: { delay: 0.25, duration: 0.6 },
    },
  }

  const formMotion = {
    initial: { opacity: 0, y: 30 },
    enter: {
      opacity: 1,
      y: 0,
      transition: { delay: 0.45, duration: 0.7 },
    },
  }
</script>

<template>
  <div class="page-container">
    <!-- background overlay -->
    <div class="page-overlay"></div>

    <!-- floating orbs / stars -->
    <div class="bg-orb orb-1"></div>
    <div class="bg-orb orb-2"></div>
    <div class="bg-orb orb-3"></div>

    <!-- Main Card -->
    <div class="card-wrapper" v-motion="cardMotion">
      <div class="card-content" v-motion="innerCardMotion">
        <!-- Header -->
        <div class="header-section" v-motion="headerMotion">
          <h1 class="main-title">WELCOME TO OUR<br />WORLD</h1>

          <p class="subtitle">
            ใส่ <span class="highlight">วันที่เราเริ่มคบกัน</span>
            <br class="mobile-break" />
            เพื่อปลดล็อกโลกของเราสองคน 💫
          </p>
        </div>

        <!-- Form -->
        <div class="form-section" v-motion="formMotion">
          <AnniversaryUnlockForm @unlocked="handleUnlocked" />
        </div>
      </div>
    </div>
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
    z-index: 1;
  }

  /* orbs ด้านหลัง ลอย ๆ */
  .bg-orb {
    position: absolute;
    border-radius: 999px;
    filter: blur(26px);
    opacity: 0.7;
    mix-blend-mode: screen;
    will-change: transform, opacity;
    z-index: 0;
  }

  .orb-1 {
    width: 260px;
    height: 260px;
    background: radial-gradient(
      circle,
      rgba(244, 114, 182, 0.9),
      rgba(168, 85, 247, 0)
    );
    top: -60px;
    left: -40px;
    animation: orb-float-1 18s ease-in-out infinite alternate;
  }

  .orb-2 {
    width: 220px;
    height: 220px;
    background: radial-gradient(
      circle,
      rgba(129, 140, 248, 0.9),
      rgba(30, 64, 175, 0)
    );
    bottom: -40px;
    right: -40px;
    animation: orb-float-2 20s ease-in-out infinite alternate;
  }

  .orb-3 {
    width: 140px;
    height: 140px;
    background: radial-gradient(
      circle,
      rgba(52, 211, 153, 0.7),
      rgba(15, 118, 110, 0)
    );
    bottom: 20%;
    left: 10%;
    animation: orb-float-3 22s ease-in-out infinite alternate;
  }

  .card-wrapper {
    width: 100%;
    max-width: 480px;
    position: relative;
    z-index: 2;
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
    will-change: transform, box-shadow;
  }

  /* hover ให้เด้งนิด ๆ + เงาสีม่วง */
  .card-content:hover {
    transform: translateY(-6px) scale(1.01);
    box-shadow: 0 25px 60px rgba(168, 85, 247, 0.4);
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

  /* glow ของตัวหนังสือแบบพุ่ง ๆ */
  .main-title::after {
    content: "✨";
    display: inline-block;
    margin-left: 0.35rem;
    font-size: 1.4rem;
    animation: title-sparkle 2.3s ease-in-out infinite;
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

  /* keyframes */

  @keyframes orb-float-1 {
    0% {
      transform: translate3d(0, 0, 0) scale(1);
      opacity: 0.75;
    }
    50% {
      transform: translate3d(40px, 20px, 0) scale(1.05);
      opacity: 1;
    }
    100% {
      transform: translate3d(10px, 60px, 0) scale(0.98);
      opacity: 0.7;
    }
  }

  @keyframes orb-float-2 {
    0% {
      transform: translate3d(0, 0, 0) scale(1.05);
    }
    50% {
      transform: translate3d(-30px, -30px, 0) scale(1);
    }
    100% {
      transform: translate3d(-10px, 10px, 0) scale(1.08);
    }
  }

  @keyframes orb-float-3 {
    0% {
      transform: translate3d(0, 0, 0) scale(1);
    }
    100% {
      transform: translate3d(40px, -30px, 0) scale(1.1);
    }
  }

  @keyframes title-sparkle {
    0%,
    100% {
      transform: translateY(0) scale(1);
      opacity: 0.8;
      text-shadow: 0 0 8px rgba(250, 250, 210, 0.7);
    }
    30% {
      transform: translateY(-2px) scale(1.25) rotate(-5deg);
      opacity: 1;
      text-shadow: 0 0 18px rgba(250, 250, 210, 1);
    }
    60% {
      transform: translateY(1px) scale(0.95) rotate(4deg);
      opacity: 0.9;
    }
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
