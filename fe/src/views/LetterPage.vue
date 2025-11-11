<script setup>
  import { ref, onMounted } from "vue"
  import { useRouter } from "vue-router"

  const router = useRouter()
  const showLetter = ref(false)
  const letterVisible = ref(false)

  onMounted(() => {
    setTimeout(() => {
      showLetter.value = true
      setTimeout(() => {
        letterVisible.value = true
      }, 300)
    }, 400)
  })

  const backToStart = () => router.push("/")
</script>

<template>
  <div class="letter-page">
    <!-- Enhanced gradient background -->
    <div class="bg-base"></div>
    <div class="bg-gradient bg-gradient-1"></div>
    <div class="bg-gradient bg-gradient-2"></div>
    <div class="bg-gradient bg-gradient-3"></div>
    <div class="bg-noise"></div>

    <!-- Ambient particles -->
    <div class="particles-layer">
      <div
        v-for="i in 15"
        :key="i"
        class="particle"
        :style="{
          left: `${Math.random() * 100}%`,
          top: `${Math.random() * 100}%`,
          animationDelay: `${Math.random() * 8}s`,
          animationDuration: `${10 + Math.random() * 15}s`,
        }"
      ></div>
    </div>

    <!-- Floating hearts -->
    <div class="floating-elements">
      <div class="heart-float" style="left: 10%; animation-delay: 0s">💜</div>
      <div class="heart-float" style="left: 50%; animation-delay: 4s">💗</div>
      <div class="heart-float" style="left: 85%; animation-delay: 8s">💕</div>
      <div class="heart-float" style="left: 30%; animation-delay: 12s">✨</div>
    </div>

    <!-- Content -->
    <div class="content-wrapper">
      <div class="content-container">
        <!-- Envelope icon -->
        <div class="icon-section" v-if="showLetter">
          <div class="icon-container">
            <div class="icon-ring icon-ring-1"></div>
            <div class="icon-ring icon-ring-2"></div>
            <div class="icon-core">
              <svg viewBox="0 0 24 24" class="icon-svg">
                <rect
                  x="3"
                  y="5"
                  width="18"
                  height="14"
                  rx="2"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                >
                  <animate
                    attributeName="stroke-dasharray"
                    values="0 64; 64 0"
                    dur="1s"
                    fill="freeze"
                  />
                </rect>
                <path
                  d="M3 7l9 6 9-6"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2.5"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <animate
                    attributeName="stroke-dasharray"
                    values="0 27; 27 0"
                    dur="0.8s"
                    begin="0.3s"
                    fill="freeze"
                  />
                </path>
              </svg>
            </div>
          </div>

          <h1 class="page-title">
            <span class="title-word">จดหมาย</span>
            <span class="title-word">จากฉัน</span>
            <span class="title-word">ถึงเธอ</span>
          </h1>
        </div>

        <!-- Letter card -->
        <div class="letter-container" v-if="letterVisible">
          <div class="letter-card">
            <div class="card-shimmer"></div>

            <!-- Decorative corners -->
            <div class="corner-decor corner-tl"></div>
            <div class="corner-decor corner-tr"></div>
            <div class="corner-decor corner-bl"></div>
            <div class="corner-decor corner-br"></div>

            <!-- Letter content -->
            <div class="letter-content">
              <p class="letter-opening">
                <span class="first-letter">ข</span
                >อบคุณที่เข้ามาในชีวิตของเค้านะครับ ทุก ๆ
                วันมีความหมายขึ้นเพราะมีเธออยู่ข้าง ๆ เธอเก่งมาก ๆ เลย
                ยินดีกับความสำเร็จในครั้งนี้ด้วยนะครับ
                และขอให้ทุกความฝันที่เธอหวังไว้เป็นจริงนะ
              </p>

              <div class="letter-quote">
                <svg class="quote-icon" viewBox="0 0 24 24">
                  <path
                    d="M6 7h6v6H6V7zm8 0h6v6h-6V7z"
                    fill="currentColor"
                    opacity="0.3"
                  />
                </svg>
                <p>
                  ไม่ว่าอนาคตจะเป็นยังไง ขอให้รู้ไว้ว่าเค้าจะอยู่ตรงนี้เสมอ
                  เพราะเธอคือ "ของขวัญ" ที่ดีที่สุดที่เค้าเคยได้รับมาเลย
                </p>
              </div>

              <div class="letter-closing">
                <p class="closing-text">รักที่สุดในโลกเลยนะ</p>
                <div class="closing-signature">
                  <span class="signature-emoji">💗</span>
                  <span class="signature-text"
                    >จากนายอริน ผู้โชคดีที่มีเธออยู่ข้าง ๆ</span
                  >
                  <span class="signature-emoji">🌙</span>
                </div>
              </div>
            </div>

            <!-- Floating decorations -->
            <div class="card-decoration heart-1">💜</div>
            <div class="card-decoration heart-2">🖤</div>
            <div class="card-decoration sparkle-1">✨</div>
            <div class="card-decoration sparkle-2">✨</div>

            <div class="card-glow"></div>
          </div>
        </div>

        <!-- Back button -->
        <div class="button-section" v-if="letterVisible">
          <button class="back-button" @click="backToStart">
            <span class="button-glow"></span>
            <span class="button-content">
              <svg class="button-icon" viewBox="0 0 24 24">
                <path
                  d="M15 18l-6-6 6-6"
                  stroke="currentColor"
                  fill="none"
                  stroke-width="2.5"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
              </svg>
              <span>กลับไปหน้าแรก</span>
            </span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
  * {
    box-sizing: border-box;
  }

  .letter-page {
    min-height: 100vh;
    position: relative;
    overflow: hidden;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #ffffff;
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Sukhumvit Set",
      sans-serif;
  }

  /* === Enhanced Background === */
  .bg-base {
    position: fixed;
    inset: 0;
    background: radial-gradient(
        circle at 15% 15%,
        rgba(139, 92, 246, 0.18) 0%,
        transparent 50%
      ),
      radial-gradient(
        circle at 85% 85%,
        rgba(236, 72, 153, 0.15) 0%,
        transparent 50%
      ),
      linear-gradient(
        135deg,
        #0a0118 0%,
        #1a0b2e 30%,
        #2d1b4e 60%,
        #1a0b2e 85%,
        #0a0118 100%
      );
    z-index: 0;
  }

  .bg-gradient {
    position: fixed;
    border-radius: 50%;
    filter: blur(90px);
    opacity: 0.35;
    mix-blend-mode: screen;
    z-index: 1;
  }

  .bg-gradient-1 {
    width: 700px;
    height: 700px;
    top: -250px;
    left: -250px;
    background: radial-gradient(
      circle,
      rgba(167, 139, 250, 0.7),
      transparent 70%
    );
    animation: float-1 22s ease-in-out infinite;
  }

  .bg-gradient-2 {
    width: 600px;
    height: 600px;
    bottom: -200px;
    right: -200px;
    background: radial-gradient(
      circle,
      rgba(244, 114, 182, 0.6),
      transparent 70%
    );
    animation: float-2 28s ease-in-out infinite;
  }

  .bg-gradient-3 {
    width: 500px;
    height: 500px;
    top: 40%;
    left: 50%;
    transform: translate(-50%, -50%);
    background: radial-gradient(
      circle,
      rgba(59, 130, 246, 0.4),
      transparent 70%
    );
    animation: float-3 34s ease-in-out infinite;
  }

  .bg-noise {
    position: fixed;
    inset: 0;
    opacity: 0.025;
    pointer-events: none;
    z-index: 2;
    background-image: url("data:image/svg+xml,%3Csvg viewBox='0 0 200 200' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.9' numOctaves='4'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)'/%3E%3C/svg%3E");
  }

  /* === Particles === */
  .particles-layer {
    position: fixed;
    inset: 0;
    pointer-events: none;
    z-index: 3;
  }

  .particle {
    position: absolute;
    width: 3px;
    height: 3px;
    background: radial-gradient(circle, rgba(255, 255, 255, 0.9), transparent);
    border-radius: 50%;
    animation: particle-drift 18s linear infinite;
  }

  /* === Floating Hearts === */
  .floating-elements {
    position: fixed;
    inset: 0;
    pointer-events: none;
    z-index: 4;
    overflow: hidden;
  }

  .heart-float {
    position: absolute;
    bottom: -60px;
    font-size: 1.75rem;
    opacity: 0;
    animation: float-up 18s ease-in infinite;
  }

  /* === Content === */
  .content-wrapper {
    position: relative;
    z-index: 10;
    width: 100%;
    padding: 60px 24px;
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
  }

  .content-container {
    width: 100%;
    max-width: 900px;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 48px;
  }

  /* === Icon Section === */
  .icon-section {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 32px;
    animation: fade-in-up 1s ease-out;
  }

  .icon-container {
    position: relative;
    width: 140px;
    height: 140px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .icon-ring {
    position: absolute;
    border-radius: 50%;
    border: 2px solid;
    animation: ring-rotate 25s linear infinite;
  }

  .icon-ring-1 {
    inset: 0;
    border-color: rgba(167, 139, 250, 0.4);
    border-style: dashed;
    border-width: 2px;
  }

  .icon-ring-2 {
    inset: -16px;
    border-color: rgba(244, 114, 182, 0.25);
    border-style: dotted;
    border-width: 3px;
    animation-direction: reverse;
    animation-duration: 18s;
  }

  .icon-core {
    position: relative;
    width: 90px;
    height: 90px;
    display: flex;
    align-items: center;
    justify-content: center;
    background: radial-gradient(
      circle,
      rgba(139, 92, 246, 0.5),
      rgba(79, 70, 229, 0.3)
    );
    border-radius: 50%;
    box-shadow: 0 0 50px rgba(167, 139, 250, 0.7),
      0 0 100px rgba(236, 72, 153, 0.4), inset 0 0 30px rgba(167, 139, 250, 0.3);
    animation: pulse-glow 3.5s ease-in-out infinite;
  }

  .icon-svg {
    width: 52px;
    height: 52px;
    stroke: #fae8ff;
    fill: none;
    stroke-linecap: round;
    stroke-linejoin: round;
  }

  .page-title {
    font-size: clamp(2.5rem, 6vw, 4.5rem);
    font-weight: 900;
    letter-spacing: -0.02em;
    text-align: center;
    margin: 0;
    line-height: 1.1;
  }

  .title-word {
    display: inline-block;
    background: linear-gradient(
      135deg,
      #ffffff 0%,
      #fae8ff 30%,
      #f9a8d4 60%,
      #fbbf24 100%
    );
    background-size: 200% 200%;
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
    animation: gradient-shift 10s ease infinite;
    filter: drop-shadow(0 0 40px rgba(244, 114, 182, 0.6));
    margin: 0 10px;
  }

  /* === Letter Card === */
  .letter-container {
    width: 100%;
    animation: fade-in-scale 0.8s ease-out;
  }

  .letter-card {
    position: relative;
    background: linear-gradient(
        135deg,
        rgba(15, 23, 42, 0.85),
        rgba(30, 64, 175, 0.5)
      ),
      radial-gradient(
        circle at 10% 10%,
        rgba(244, 114, 182, 0.15),
        transparent 60%
      );
    border: 1px solid rgba(255, 255, 255, 0.15);
    border-radius: 32px;
    padding: 56px 48px;
    backdrop-filter: blur(20px);
    box-shadow: 0 25px 80px rgba(0, 0, 0, 0.6), 0 0 60px rgba(139, 92, 246, 0.3),
      inset 0 1px 0 rgba(255, 255, 255, 0.1);
    overflow: hidden;
    transition: transform 0.4s ease, box-shadow 0.4s ease;
  }

  .letter-card:hover {
    transform: translateY(-8px);
    box-shadow: 0 35px 100px rgba(0, 0, 0, 0.7),
      0 0 80px rgba(167, 139, 250, 0.5);
  }

  .card-shimmer {
    position: absolute;
    inset: -150%;
    background: linear-gradient(
      90deg,
      transparent,
      rgba(255, 255, 255, 0.08),
      transparent
    );
    transform: translateX(-100%) rotate(45deg);
    animation: shimmer-sweep 8s ease-in-out infinite;
  }

  /* === Corner Decorations === */
  .corner-decor {
    position: absolute;
    width: 40px;
    height: 40px;
    border: 2px solid rgba(244, 114, 182, 0.4);
  }

  .corner-tl {
    top: 16px;
    left: 16px;
    border-right: none;
    border-bottom: none;
    border-radius: 8px 0 0 0;
  }

  .corner-tr {
    top: 16px;
    right: 16px;
    border-left: none;
    border-bottom: none;
    border-radius: 0 8px 0 0;
  }

  .corner-bl {
    bottom: 16px;
    left: 16px;
    border-right: none;
    border-top: none;
    border-radius: 0 0 0 8px;
  }

  .corner-br {
    bottom: 16px;
    right: 16px;
    border-left: none;
    border-top: none;
    border-radius: 0 0 8px 0;
  }

  /* === Letter Content === */
  .letter-content {
    position: relative;
    z-index: 2;
  }

  .letter-opening {
    font-size: 1.125rem;
    line-height: 2;
    color: rgba(255, 255, 255, 0.9);
    margin: 0 0 32px;
    text-align: justify;
  }

  .first-letter {
    float: left;
    font-size: 4rem;
    font-weight: 900;
    line-height: 0.85;
    margin-right: 12px;
    margin-top: 8px;
    background: linear-gradient(135deg, #e9d5ff, #f9a8d4);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
    filter: drop-shadow(0 0 20px rgba(244, 114, 182, 0.6));
  }

  .letter-quote {
    position: relative;
    margin: 32px 0;
    padding: 24px 32px;
    padding-left: 56px;
    background: rgba(139, 92, 246, 0.1);
    border-left: 4px solid rgba(244, 114, 182, 0.8);
    border-radius: 0 16px 16px 0;
    backdrop-filter: blur(10px);
  }

  .quote-icon {
    position: absolute;
    top: 20px;
    left: 16px;
    width: 28px;
    height: 28px;
    color: rgba(244, 114, 182, 0.6);
  }

  .letter-quote p {
    font-size: 1.125rem;
    line-height: 1.9;
    color: #fae8ff;
    margin: 0;
    font-style: italic;
    text-shadow: 0 0 20px rgba(236, 72, 153, 0.3);
  }

  .letter-closing {
    margin-top: 48px;
    text-align: center;
  }

  .closing-text {
    font-size: 1.5rem;
    font-weight: 700;
    font-style: italic;
    background: linear-gradient(135deg, #fecaca, #fde68a);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
    margin: 0 0 16px;
    filter: drop-shadow(0 0 20px rgba(248, 113, 113, 0.5));
  }

  .closing-signature {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 12px;
    flex-wrap: wrap;
  }

  .signature-emoji {
    font-size: 1.5rem;
    animation: pulse-soft 2.5s ease-in-out infinite;
  }

  .signature-text {
    font-size: 1rem;
    color: rgba(255, 255, 255, 0.8);
  }

  /* === Card Decorations === */
  .card-decoration {
    position: absolute;
    font-size: 2rem;
    pointer-events: none;
    animation: float-gentle 4s ease-in-out infinite;
  }

  .heart-1 {
    top: 20px;
    right: 24px;
    animation-delay: 0s;
  }

  .heart-2 {
    bottom: 24px;
    left: 20px;
    animation-delay: 1s;
  }

  .sparkle-1 {
    top: 50%;
    right: 16px;
    font-size: 1.5rem;
    animation-delay: 0.5s;
  }

  .sparkle-2 {
    bottom: 40%;
    left: 16px;
    font-size: 1.5rem;
    animation-delay: 1.5s;
  }

  .card-glow {
    position: absolute;
    bottom: -60%;
    left: 50%;
    width: 80%;
    height: 80%;
    background: radial-gradient(
      circle,
      rgba(167, 139, 250, 0.4),
      transparent 70%
    );
    transform: translateX(-50%);
    opacity: 0;
    transition: opacity 0.6s;
    pointer-events: none;
  }

  .letter-card:hover .card-glow {
    opacity: 1;
  }

  /* === Button === */
  .button-section {
    animation: fade-in-up 1s ease-out 0.5s both;
  }

  .back-button {
    position: relative;
    padding: 18px 48px;
    background: linear-gradient(135deg, #8b5cf6, #ec4899, #f59e0b);
    background-size: 200% 200%;
    border: none;
    border-radius: 9999px;
    font-size: 1.125rem;
    font-weight: 700;
    color: #ffffff;
    cursor: pointer;
    overflow: hidden;
    transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
    box-shadow: 0 20px 40px rgba(139, 92, 246, 0.4),
      0 0 60px rgba(236, 72, 153, 0.3);
    animation: gradient-shift 6s ease infinite;
  }

  .back-button:hover {
    transform: translateY(-4px) scale(1.05);
    box-shadow: 0 30px 60px rgba(139, 92, 246, 0.6),
      0 0 80px rgba(236, 72, 153, 0.5);
  }

  .back-button:active {
    transform: translateY(-2px) scale(1.02);
  }

  .button-glow {
    position: absolute;
    inset: -20px;
    background: radial-gradient(
      circle,
      rgba(236, 72, 153, 0.6),
      transparent 70%
    );
    opacity: 0.5;
    filter: blur(25px);
    animation: pulse-glow 3s ease-in-out infinite;
  }

  .button-content {
    position: relative;
    display: flex;
    align-items: center;
    gap: 12px;
    z-index: 1;
  }

  .button-icon {
    width: 24px;
    height: 24px;
    animation: arrow-bounce 2s ease-in-out infinite;
  }

  /* === Animations === */
  @keyframes float-1 {
    0%,
    100% {
      transform: translate(0, 0) scale(1);
    }
    50% {
      transform: translate(120px, -120px) scale(1.15);
    }
  }

  @keyframes float-2 {
    0%,
    100% {
      transform: translate(0, 0) scale(1);
    }
    50% {
      transform: translate(-140px, 100px) scale(1.2);
    }
  }

  @keyframes float-3 {
    0%,
    100% {
      transform: translate(-50%, -50%) scale(1);
    }
    50% {
      transform: translate(-45%, -55%) scale(1.25);
    }
  }

  @keyframes particle-drift {
    0% {
      transform: translate(0, 0);
      opacity: 0;
    }
    10% {
      opacity: 1;
    }
    90% {
      opacity: 1;
    }
    100% {
      transform: translate(120px, -1000px);
      opacity: 0;
    }
  }

  @keyframes float-up {
    0% {
      transform: translateY(0) rotate(0deg) scale(0.8);
      opacity: 0;
    }
    10% {
      opacity: 1;
    }
    90% {
      opacity: 1;
    }
    100% {
      transform: translateY(-120vh) rotate(360deg) scale(1);
      opacity: 0;
    }
  }

  @keyframes fade-in-up {
    0% {
      opacity: 0;
      transform: translateY(40px);
    }
    100% {
      opacity: 1;
      transform: translateY(0);
    }
  }

  @keyframes fade-in-scale {
    0% {
      opacity: 0;
      transform: scale(0.95);
    }
    100% {
      opacity: 1;
      transform: scale(1);
    }
  }

  @keyframes ring-rotate {
    to {
      transform: rotate(360deg);
    }
  }

  @keyframes pulse-glow {
    0%,
    100% {
      opacity: 0.5;
      transform: scale(1);
    }
    50% {
      opacity: 1;
      transform: scale(1.15);
    }
  }

  @keyframes gradient-shift {
    0%,
    100% {
      background-position: 0% 50%;
    }
    50% {
      background-position: 100% 50%;
    }
  }

  @keyframes shimmer-sweep {
    0% {
      transform: translateX(-100%) rotate(45deg);
    }
    15% {
      transform: translateX(100%) rotate(45deg);
    }
    100% {
      transform: translateX(100%) rotate(45deg);
    }
  }

  @keyframes float-gentle {
    0%,
    100% {
      transform: translateY(0) scale(1);
      opacity: 0.7;
    }
    50% {
      transform: translateY(-10px) scale(1.1);
      opacity: 1;
    }
  }

  @keyframes pulse-soft {
    0%,
    100% {
      transform: scale(1);
      opacity: 0.9;
    }
    50% {
      transform: scale(1.15);
      opacity: 1;
    }
  }

  @keyframes arrow-bounce {
    0%,
    100% {
      transform: translateX(0);
    }
    50% {
      transform: translateX(-6px);
    }
  }

  /* === Responsive === */
  @media (max-width: 1024px) {
    .icon-container {
      width: 120px;
      height: 120px;
    }

    .icon-core {
      width: 80px;
      height: 80px;
    }

    .icon-svg {
      width: 46px;
      height: 46px;
    }
  }

  @media (max-width: 768px) {
    .content-wrapper {
      padding: 40px 20px;
    }

    .content-container {
      gap: 36px;
    }

    .icon-container {
      width: 100px;
      height: 100px;
    }

    .icon-core {
      width: 70px;
      height: 70px;
    }

    .icon-svg {
      width: 40px;
      height: 40px;
    }

    .page-title {
      font-size: 2.5rem;
    }

    .title-word {
      margin: 0 6px;
    }

    .letter-card {
      padding: 40px 28px;
      border-radius: 24px;
    }

    .letter-opening {
      font-size: 1rem;
      line-height: 1.9;
    }

    .first-letter {
      font-size: 3rem;
      margin-right: 8px;
    }

    .letter-quote {
      padding: 20px 24px;
      padding-left: 48px;
    }

    .letter-quote p {
      font-size: 1rem;
    }

    .closing-text {
      font-size: 1.25rem;
    }

    .back-button {
      padding: 16px 36px;
      font-size: 1rem;
    }

    .corner-decor {
      width: 30px;
      height: 30px;
    }

    .card-decoration {
      font-size: 1.5rem;
    }
  }

  @media (max-width: 480px) {
    .page-title {
      font-size: 2rem;
    }

    .letter-card {
      padding: 32px 20px;
    }

    .letter-opening {
      font-size: 0.9375rem;
    }

    .first-letter {
      font-size: 2.5rem;
      margin-top: 4px;
    }

    .letter-quote {
      padding: 16px 20px;
      padding-left: 44px;
    }

    .letter-quote p {
      font-size: 0.9375rem;
    }

    .closing-text {
      font-size: 1.125rem;
    }

    .signature-text {
      font-size: 0.875rem;
    }

    .back-button {
      padding: 14px 28px;
      font-size: 0.9375rem;
    }
  }
</style>
