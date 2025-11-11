<script setup>
  import { useRouter } from "vue-router"

  const router = useRouter()

  const goToMoments = () => {
    router.push("/moments")
  }
</script>

<template>
  <div
    class="welcome-page min-h-screen flex flex-col items-center justify-center relative overflow-hidden p-6 text-center"
  >
    <!-- Animated gradient + aurora background -->
    <div class="bg-layer"></div>
    <div class="bg-aurora bg-aurora-1"></div>
    <div class="bg-aurora bg-aurora-2"></div>
    <div class="bg-noise"></div>

    <!-- Floating hearts & particles -->
    <div class="floating-layer pointer-events-none">
      <div class="heart-float heart-1">❤️</div>
      <div class="heart-float heart-2">💜</div>
      <div class="heart-float heart-3">🖤</div>
      <div class="heart-float heart-4">💗</div>
      <div class="heart-float heart-5">💕</div>

      <div class="star star-1"></div>
      <div class="star star-2"></div>
      <div class="star star-3"></div>
      <div class="star star-4"></div>
      <div class="star star-5"></div>
    </div>

    <!-- Content -->
    <div class="content-wrapper relative z-10 max-w-3xl">
      <!-- Glass card -->
      <div class="hero-card">
        <!-- Clock icon -->
        <div class="icon-wrapper">
          <div class="icon-orbit"></div>
          <n-icon size="80" color="#c4b5fd" class="clock-icon">
            <svg viewBox="0 0 24 24" class="w-16 h-16">
              <!-- outer circle stroke draw -->
              <circle
                cx="12"
                cy="12"
                r="9"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-dasharray="60"
                stroke-dashoffset="60"
              >
                <animate
                  attributeName="stroke-dashoffset"
                  values="60;0"
                  dur="0.9s"
                  fill="freeze"
                />
              </circle>
              <!-- clock hands -->
              <path
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 7v5l3 3"
              >
                <animate
                  attributeName="stroke-dashoffset"
                  from="16"
                  to="0"
                  dur="0.6s"
                  begin="0.3s"
                  fill="freeze"
                  calcMode="linear"
                />
              </path>
            </svg>
          </n-icon>
        </div>

        <!-- Heading -->
        <h1 class="hero-title">
          ของขวัญชิ้นนี้ <br />
          อาจจะไม่ได้พิเศษอะไร แต่ก็ตั้งใจทำให้เธอ  <br />
          <span class="hero-title-accent">My Babe</span>
        </h1>

        <!-- Subcopy -->
        <n-space vertical :size="16" class="mb-8">
          <n-text class="hero-text">
            ขอบคุณที่อยู่เคียงข้างกันมาตลอดเลย 💫<br />
            วันนี้เธอประสบความสำเร็จอีกก้าวหนึ่งในชีวิตแล้วนะ<br />
            เค้าภูมิใจในตัวเธอมาก ๆ เลย
          </n-text>

          <n-blockquote class="hero-quote">
            “เวลาที่ดีที่สุด คือเวลาที่เราอยู่ด้วยกัน”
          </n-blockquote>
        </n-space>

        <!-- Button -->
        <div class="btn-wrapper">
          <div class="btn-glow"></div>
          <n-button
            type="primary"
            size="large"
            @click="goToMoments"
            class="hero-button"
          >
            <template #icon>
              <svg
                class="w-5 h-5 arrow-icon"
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
            ต่อไป
          </n-button>
        </div>

        <div class="hint-text">
          ไม่ค่อยมีเวลาทำของให้เลย ทำเว็บให้ล่ะกันนะคะ💖
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
  /* === Layout & Background === */
  .welcome-page {
    color: #f9fafb;
  }

  .bg-layer {
    position: absolute;
    inset: 0;
    background: radial-gradient(
        circle at 10% 0%,
        rgba(196, 181, 253, 0.25),
        transparent 55%
      ),
      radial-gradient(
        circle at 90% 100%,
        rgba(248, 113, 113, 0.22),
        transparent 55%
      ),
      radial-gradient(
        circle at 50% 50%,
        rgba(56, 189, 248, 0.12),
        transparent 60%
      ),
      linear-gradient(
        135deg,
        #020617 0%,
        #1e1b4b 40%,
        #4c1d95 80%,
        #020617 100%
      );
    z-index: 0;
  }

  /* aurora ribbons */
  .bg-aurora {
    position: absolute;
    inset: -40%;
    opacity: 0.55;
    mix-blend-mode: screen;
    filter: blur(32px);
    background: conic-gradient(
      from 120deg,
      rgba(244, 114, 182, 0.6),
      rgba(56, 189, 248, 0.2),
      rgba(167, 139, 250, 0.7),
      rgba(248, 113, 113, 0.35),
      rgba(56, 189, 248, 0.2)
    );
    animation: aurora-move 22s ease-in-out infinite alternate;
    z-index: 0;
  }

  .bg-aurora-2 {
    animation-duration: 28s;
    animation-delay: -6s;
    opacity: 0.4;
  }

  .bg-noise {
    position: absolute;
    inset: -50px;
    opacity: 0.16;
    pointer-events: none;
    z-index: 1;
    background-image: url("data:image/svg+xml,%3Csvg viewBox='0 0 160 160' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='.8' numOctaves='3' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)' opacity='0.9'/%3E%3C/svg%3E");
    mix-blend-mode: soft-light;
  }

  /* === Floating emojis & stars === */
  .floating-layer {
    position: absolute;
    inset: 0;
    overflow: hidden;
    z-index: 2;
  }

  .heart-float {
    position: absolute;
    bottom: -60px;
    font-size: 2rem;
    animation: float-up 14s linear infinite;
    opacity: 0.9;
  }

  .heart-1 {
    left: 8%;
    animation-duration: 15s;
  }
  .heart-2 {
    left: 25%;
    animation-duration: 17s;
    animation-delay: 2s;
  }
  .heart-3 {
    left: 50%;
    animation-duration: 19s;
    animation-delay: 4s;
  }
  .heart-4 {
    left: 72%;
    animation-duration: 16s;
    animation-delay: 1s;
  }
  .heart-5 {
    left: 90%;
    animation-duration: 18s;
    animation-delay: 3s;
  }

  /* tiny glowing stars */
  .star {
    position: absolute;
    width: 6px;
    height: 6px;
    border-radius: 999px;
    background: radial-gradient(circle, #e5e7eb, transparent);
    box-shadow: 0 0 12px rgba(248, 250, 252, 0.9);
    animation: star-twinkle 4s ease-in-out infinite;
  }

  .star-1 {
    top: 15%;
    left: 18%;
  }
  .star-2 {
    top: 30%;
    right: 20%;
    animation-delay: 1s;
  }
  .star-3 {
    top: 65%;
    left: 28%;
    animation-delay: 2s;
  }
  .star-4 {
    top: 75%;
    right: 12%;
    animation-delay: 1.5s;
  }
  .star-5 {
    top: 45%;
    left: 50%;
    animation-delay: 0.5s;
  }

  /* === Hero card === */
  .content-wrapper {
    animation: fade-up 0.9s ease-out forwards;
    transform-origin: center;
  }

  .hero-card {
    position: relative;
    padding: 2.5rem 2rem 2.75rem;
    border-radius: 28px;
    backdrop-filter: blur(22px);
    background: linear-gradient(
        135deg,
        rgba(15, 23, 42, 0.65),
        rgba(30, 64, 175, 0.35)
      ),
      radial-gradient(
        circle at 0% 0%,
        rgba(244, 114, 182, 0.18),
        transparent 55%
      ),
      radial-gradient(
        circle at 100% 100%,
        rgba(129, 140, 248, 0.16),
        transparent 55%
      );
    border: 1px solid rgba(251, 207, 232, 0.18);
    box-shadow: 0 22px 60px rgba(15, 23, 42, 0.9),
      0 0 45px rgba(147, 51, 234, 0.35);
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 1.75rem;
    overflow: hidden;
  }

  .hero-card::before {
    content: "";
    position: absolute;
    inset: -40%;
    background: conic-gradient(
      from 180deg,
      transparent,
      rgba(248, 250, 252, 0.18),
      transparent 30%
    );
    opacity: 0;
    transform: rotate(0deg);
    animation: card-sheen 6s ease-in-out infinite;
    pointer-events: none;
  }

  /* === Icon area === */
  .icon-wrapper {
    position: relative;
    margin-bottom: 0.75rem;
  }

  .icon-orbit {
    position: absolute;
    inset: -18px;
    border-radius: 999px;
    border: 1px dashed rgba(196, 181, 253, 0.4);
    box-shadow: 0 0 18px rgba(196, 181, 253, 0.5);
    opacity: 0.9;
    animation: orbit-spin 18s linear infinite;
  }

  .n-icon {
    display: flex !important;
    align-items: center;
    justify-content: center;
    line-height: 1 !important;
  }

  .clock-icon {
    position: relative;
    z-index: 1;
    padding: 0.6rem;
    border-radius: 999px;
    background: radial-gradient(
      circle,
      rgba(79, 70, 229, 0.8),
      rgba(15, 23, 42, 0.8)
    );
    box-shadow: 0 0 25px rgba(129, 140, 248, 0.9),
      0 0 45px rgba(59, 130, 246, 0.55);
    animation: icon-bounce 3s ease-in-out infinite;
  }

  /* === Typography === */
  .hero-title {
    font-size: 2.4rem;
    line-height: 1.2;
    font-weight: 800;
    letter-spacing: 0.03em;
    margin-bottom: 0.25rem;
    background: linear-gradient(
      90deg,
      #e5e7eb,
      #e9d5ff,
      #f9a8d4,
      #e5e7eb,
      #e9d5ff
    );
    background-size: 200% auto;
    -webkit-background-clip: text;
    color: transparent;
    animation: title-gradient 5s ease-in-out infinite alternate;
  }

  .hero-title-accent {
    display: block;
    font-size: 2.6rem;
    letter-spacing: 0.08em;
    text-transform: uppercase;
    margin-top: 0.25rem;
    text-shadow: 0 0 20px rgba(244, 114, 182, 0.7),
      0 0 40px rgba(129, 140, 248, 0.75);
  }

  .hero-text {
    font-size: 1.05rem;
    color: rgba(224, 231, 255, 0.95);
    line-height: 1.75;
  }

  .hero-quote {
    color: #e9d5ff;
    border-left-color: rgba(244, 114, 182, 0.7);
    font-size: 0.95rem;
    text-shadow: 0 0 14px rgba(236, 72, 153, 0.3);
  }

  /* === Button === */
  .btn-wrapper {
    position: relative;
    display: inline-flex;
    margin-bottom: 0.75rem;
  }

  .btn-glow {
    position: absolute;
    inset: 6px;
    border-radius: 999px;
    background: radial-gradient(
      circle,
      rgba(248, 113, 113, 0.45),
      rgba(79, 70, 229, 0)
    );
    opacity: 0.8;
    filter: blur(10px);
    z-index: -1;
    animation: btn-glow 3.4s ease-in-out infinite;
  }

  .hero-button {
    position: relative;
    height: 3.6rem;
    padding: 0 2.4rem;
    font-size: 1.05rem;
    font-weight: 700;
    border-radius: 999px;
    background-image: linear-gradient(120deg, #a855f7, #ec4899, #f97316);
    background-size: 200% auto;
    border: none;
    box-shadow: 0 18px 35px rgba(236, 72, 153, 0.6),
      0 0 25px rgba(129, 140, 248, 0.8);
    transition: transform 0.18s ease, box-shadow 0.18s ease,
      background-position 0.35s ease;
    display: inline-flex;
    align-items: center;
    gap: 0.55rem;
  }

  .hero-button:hover {
    transform: translateY(-2px) scale(1.02);
    background-position: 100% 0;
    box-shadow: 0 22px 40px rgba(236, 72, 153, 0.8),
      0 0 32px rgba(129, 140, 248, 1);
  }

  .hero-button:active {
    transform: translateY(0) scale(0.99);
    box-shadow: 0 10px 24px rgba(236, 72, 153, 0.6),
      0 0 18px rgba(129, 140, 248, 0.9);
  }

  .arrow-icon {
    transform-origin: center;
    animation: arrow-pulse 1.4s ease-in-out infinite;
  }

  .hint-text {
    font-size: 0.85rem;
    color: rgba(209, 213, 219, 0.85);
    margin-top: 0.25rem;
  }

  /* === Animations === */
  @keyframes aurora-move {
    0% {
      transform: translate3d(-10%, 10%, 0) scale(1.05) rotate(0deg);
    }
    50% {
      transform: translate3d(10%, -5%, 0) scale(1.08) rotate(8deg);
    }
    100% {
      transform: translate3d(-8%, -12%, 0) scale(1.02) rotate(-6deg);
    }
  }

  @keyframes float-up {
    0% {
      transform: translate3d(0, 100vh, 0) scale(0.6) rotate(0deg);
      opacity: 0;
    }
    10% {
      opacity: 1;
    }
    50% {
      transform: translate3d(-15px, 0, 0) scale(1) rotate(-10deg);
    }
    80% {
      opacity: 1;
    }
    100% {
      transform: translate3d(15px, -110vh, 0) scale(1.1) rotate(8deg);
      opacity: 0;
    }
  }

  @keyframes star-twinkle {
    0%,
    100% {
      opacity: 0.2;
      transform: scale(0.8);
    }
    50% {
      opacity: 1;
      transform: scale(1.3);
    }
  }

  @keyframes fade-up {
    0% {
      opacity: 0;
      transform: translate3d(0, 28px, 0) scale(0.98);
    }
    100% {
      opacity: 1;
      transform: translate3d(0, 0, 0) scale(1);
    }
  }

  @keyframes card-sheen {
    0% {
      opacity: 0;
      transform: translateX(-40%) rotate(0deg);
    }
    8% {
      opacity: 1;
    }
    18% {
      transform: translateX(40%) rotate(18deg);
      opacity: 0;
    }
    100% {
      opacity: 0;
      transform: translateX(40%) rotate(18deg);
    }
  }

  @keyframes orbit-spin {
    0% {
      transform: rotate(0deg) scale(1);
    }
    100% {
      transform: rotate(360deg) scale(1);
    }
  }

  @keyframes icon-bounce {
    0%,
    100% {
      transform: translateY(0) scale(1);
    }
    40% {
      transform: translateY(-4px) scale(1.03);
    }
    60% {
      transform: translateY(1px) scale(0.99);
    }
  }

  @keyframes title-gradient {
    0% {
      background-position: 0% 50%;
    }
    100% {
      background-position: 100% 50%;
    }
  }

  @keyframes btn-glow {
    0%,
    100% {
      opacity: 0.65;
      transform: scale(0.95);
    }
    50% {
      opacity: 1;
      transform: scale(1.05);
    }
  }

  @keyframes arrow-pulse {
    0%,
    100% {
      transform: translateX(0);
    }
    50% {
      transform: translateX(4px);
    }
  }

  /* === Responsive tweaks === */
  @media (max-width: 768px) {
    .hero-card {
      padding: 2.1rem 1.4rem 2.25rem;
      border-radius: 24px;
    }

    .hero-title {
      font-size: 1.9rem;
    }

    .hero-title-accent {
      font-size: 2.1rem;
    }

    .hero-text {
      font-size: 0.98rem;
    }

    .hero-button {
      height: 3.25rem;
      padding: 0 1.9rem;
      font-size: 0.98rem;
    }
  }
</style>
