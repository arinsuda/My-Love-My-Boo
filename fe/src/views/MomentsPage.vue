<script setup>
  import { ref, onMounted, computed } from "vue"
  import { useRouter } from "vue-router"

  const router = useRouter()
  const moments = ref([])
  const loading = ref(true)
  const viewMode = ref("grid")
  const hoveredMoment = ref(null)

  // 👇 สำหรับ media viewer
  const activeMoment = ref(null) // moment ที่กำลังดู
  const activeIndex = ref(0) // index ของรูป/วิดีโอใน moment นั้น

  const API_BASE = import.meta.env.VITE_API_BASE_URL

  const loadMoments = async () => {
    loading.value = true
    try {
      const res = await fetch(`${API_BASE}/api/moments`)
      if (!res.ok) throw new Error("load fail")

      const data = await res.json()
      moments.value = Array.isArray(data) ? data : []
    } catch (err) {
      console.error("โหลด moments ไม่สำเร็จ:", err)
      moments.value = []
    } finally {
      loading.value = false
    }
  }

  onMounted(loadMoments)

  const goToLetter = () => {
    router.push("/letter")
  }

  const goToWelcome = () => {
    router.push("/welcome")
  }

  const setViewMode = mode => {
    viewMode.value = mode
  }

  const formatDate = dateStr => {
    const [year, month, day] = dateStr.split("-")
    const months = [
      "ม.ค.",
      "ก.พ.",
      "มี.ค.",
      "เม.ย.",
      "พ.ค.",
      "มิ.ย.",
      "ก.ค.",
      "ส.ค.",
      "ก.ย.",
      "ต.ค.",
      "พ.ย.",
      "ธ.ค.",
    ]
    return `${parseInt(day)} ${months[parseInt(month) - 1]} ${
      parseInt(year) + 543
    }`
  }

  // ---------- media viewer helpers ----------
  const currentMediaUrl = computed(() => {
    if (!activeMoment.value) return ""
    return activeMoment.value.images?.[activeIndex.value] || ""
  })

  const isVideo = url => !!url && /\.(mp4|webm|ogg)$/i.test(url) // ถ้าไฟล์ลงท้ายแบบนี้ให้ถือว่าเป็นวิดีโอ

  const openViewer = (moment, index = 0) => {
    if (!moment?.images?.length) return
    activeMoment.value = moment
    activeIndex.value = index
  }

  const closeViewer = () => {
    activeMoment.value = null
    activeIndex.value = 0
  }

  const nextMedia = () => {
    if (!activeMoment.value) return
    const total = activeMoment.value.images?.length || 0
    if (!total) return
    activeIndex.value = (activeIndex.value + 1) % total
  }

  const prevMedia = () => {
    if (!activeMoment.value) return
    const total = activeMoment.value.images?.length || 0
    if (!total) return
    activeIndex.value = (activeIndex.value - 1 + total) % total
  }
</script>

<template>
  <div class="moments-page">
    <button class="top-back-button" @click="goToWelcome">
      <span class="top-back-glow"></span>
      <span class="top-back-content">
        <svg class="top-back-icon" viewBox="0 0 24 24">
          <path d="M19 12H5" />
          <path d="M12 19l-7-7 7-7" />
        </svg>
        <span class="top-back-text">กลับ</span>
      </span>
    </button>

    <!-- Enhanced gradient mesh background -->
    <div class="bg-base"></div>
    <div class="bg-gradient bg-gradient-1"></div>
    <div class="bg-gradient bg-gradient-2"></div>
    <div class="bg-gradient bg-gradient-3"></div>
    <div class="bg-noise"></div>

    <!-- Ambient particles -->
    <div class="particles-layer">
      <div
        v-for="i in 20"
        :key="i"
        class="particle"
        :style="{
          left: `${Math.random() * 100}%`,
          top: `${Math.random() * 100}%`,
          animationDelay: `${Math.random() * 8}s`,
          animationDuration: `${8 + Math.random() * 12}s`,
        }"
      ></div>
    </div>

    <!-- Floating elements -->
    <div class="floating-elements">
      <div class="heart-float" style="left: 8%; animation-delay: 0s">💜</div>
      <div class="heart-float" style="left: 50%; animation-delay: 3s">💗</div>
      <div class="heart-float" style="left: 85%; animation-delay: 6s">✨</div>
      <div class="heart-float" style="left: 30%; animation-delay: 9s">💕</div>
    </div>

    <!-- Content -->
    <div class="content-wrapper">
      <div class="content-container">
        <!-- Header section with enhanced design -->
        <div class="header-section">
          <div class="header-content">
            <!-- Icon with enhanced effects -->
            <div class="icon-container">
              <div class="icon-ring icon-ring-1"></div>
              <div class="icon-ring icon-ring-2"></div>
              <div class="icon-core">
                <svg viewBox="0 0 24 24" class="icon-svg">
                  <circle
                    cx="12"
                    cy="12"
                    r="9"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                  >
                    <animate
                      attributeName="stroke-dasharray"
                      values="0 56.5; 56.5 0"
                      dur="1.2s"
                      fill="freeze"
                    />
                  </circle>
                  <path
                    d="M12 7v5l3 3"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2.5"
                    stroke-linecap="round"
                  />
                </svg>
              </div>
            </div>

            <h1 class="main-title">
              <span class="title-word">อัลบั้ม</span>
              <span class="title-word">ความทรงจำ</span>
              <span class="title-word">ของเรา</span>
            </h1>

            <p class="subtitle">
              เก็บทุกช่วงเวลาพิเศษไว้เหมือนอัลบั้มรูปส่วนตัวของเราสองคน
            </p>

            <div v-if="!loading" class="meta-info">
              <span class="meta-item">
                <svg
                  class="meta-icon"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                >
                  <rect
                    x="3"
                    y="4"
                    width="18"
                    height="18"
                    rx="2"
                    ry="2"
                    stroke-width="2"
                  />
                  <line
                    x1="16"
                    y1="2"
                    x2="16"
                    y2="6"
                    stroke-width="2"
                    stroke-linecap="round"
                  />
                  <line
                    x1="8"
                    y1="2"
                    x2="8"
                    y2="6"
                    stroke-width="2"
                    stroke-linecap="round"
                  />
                  <line x1="3" y1="10" x2="21" y2="10" stroke-width="2" />
                </svg>
                {{ moments.length }} ช่วงเวลา
              </span>
              <span class="meta-divider">·</span>
              <span class="meta-item">
                <svg
                  class="meta-icon"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                >
                  <path
                    d="M12 20l9-11H3z"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  />
                </svg>
                {{ viewMode === "grid" ? "Grid View" : "List View" }}
              </span>
            </div>
          </div>

          <!-- View toggle with refined design -->
          <div class="view-controls">
            <div class="toggle-group">
              <button
                class="toggle-btn"
                :class="{ active: viewMode === 'grid' }"
                @click="setViewMode('grid')"
              >
                <svg viewBox="0 0 24 24">
                  <rect x="3" y="3" width="7" height="7" rx="1.5" />
                  <rect x="14" y="3" width="7" height="7" rx="1.5" />
                  <rect x="3" y="14" width="7" height="7" rx="1.5" />
                  <rect x="14" y="14" width="7" height="7" rx="1.5" />
                </svg>
                <span>Grid</span>
              </button>
              <button
                class="toggle-btn"
                :class="{ active: viewMode === 'list' }"
                @click="setViewMode('list')"
              >
                <svg viewBox="0 0 24 24">
                  <line x1="4" y1="6" x2="20" y2="6" />
                  <line x1="4" y1="12" x2="20" y2="12" />
                  <line x1="4" y1="18" x2="20" y2="18" />
                </svg>
                <span>List</span>
              </button>
            </div>
          </div>
        </div>

        <!-- Loading state -->
        <div v-if="loading" class="loading-container">
          <div class="loading-spinner">
            <div class="spinner-ring"></div>
            <div class="spinner-ring"></div>
            <div class="spinner-ring"></div>
          </div>
          <p class="loading-text">กำลังเปิดอัลบั้มรูปของเรา...</p>
        </div>

        <!-- Grid View -->
        <div v-else-if="viewMode === 'grid'" class="grid-view">
          <div
            v-for="(m, idx) in moments"
            :key="m.id"
            class="grid-card"
            :style="{ animationDelay: `${idx * 0.1}s` }"
            @mouseenter="hoveredMoment = m.id"
            @mouseleave="hoveredMoment = null"
            @click="openViewer(m, 0)"
          >
            <div class="card-shimmer"></div>
            <div class="card-image-wrapper">
              <!-- ถ้ามีรูปเดียว -->
              <template v-if="(m.images?.length || 0) <= 1">
                <img
                  :src="m.images?.[0]"
                  :alt="m.title"
                  class="card-image single"
                />
                <div class="image-overlay"></div>
              </template>

              <!-- ถ้ามีมากกว่า 1 รูป -->
              <template v-else>
                <div class="dual-images">
                  <!-- รูปซ้าย (รูปแรก) -->
                  <div class="dual-left">
                    <img
                      :src="m.images?.[0]"
                      :alt="m.title"
                      class="card-image"
                    />
                  </div>

                  <!-- รูปขวา (รูปที่ 2 + overlay + count) -->
                  <div class="dual-right">
                    <img
                      :src="m.images?.[1]"
                      :alt="m.title"
                      class="card-image"
                    />
                    <div class="dual-overlay"></div>

                    <!-- แสดงจำนวนรูปที่เหลือ (total - 2) -->
                    <div v-if="(m.images?.length || 0) > 2" class="dual-count">
                      +{{ (m.images?.length || 0) - 2 }}
                    </div>
                  </div>
                </div>

                <!-- overlay ด้านล่างเพื่อให้อ่าน title ได้ -->
                <div class="image-overlay"></div>
              </template>
            </div>

            <div class="card-content">
              <div class="card-tag" v-if="m.tag">{{ m.tag }}</div>
              <h3 class="card-title">{{ m.title }}</h3>
              <div class="card-date">
                <svg viewBox="0 0 24 24" class="date-icon">
                  <circle
                    cx="12"
                    cy="12"
                    r="10"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                  />
                  <path
                    d="M12 6v6l4 2"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                  />
                </svg>
                {{ formatDate(m.date) }}
              </div>
            </div>
            <div class="card-glow"></div>
          </div>
        </div>

        <!-- List View -->
        <div v-else class="list-view">
          <div
            v-for="(m, idx) in moments"
            :key="m.id"
            class="list-card"
            :style="{ animationDelay: `${idx * 0.1}s` }"
            @mouseenter="hoveredMoment = m.id"
            @mouseleave="hoveredMoment = null"
            @click="openViewer(m, 0)"
          >
            <div class="list-shimmer"></div>
            <div class="list-image-wrapper">
              <img :src="m.images?.[0]" :alt="m.title" class="list-image" />
              <div class="list-image-overlay"></div>
            </div>
            <div class="list-content">
              <div class="list-header">
                <div class="list-tag" v-if="m.tag">{{ m.tag }}</div>
                <div class="list-date">
                  <svg viewBox="0 0 24 24" class="date-icon">
                    <rect
                      x="3"
                      y="4"
                      width="18"
                      height="18"
                      rx="2"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                    />
                    <line
                      x1="16"
                      y1="2"
                      x2="16"
                      y2="6"
                      stroke="currentColor"
                      stroke-width="2"
                    />
                    <line
                      x1="8"
                      y1="2"
                      x2="8"
                      y2="6"
                      stroke="currentColor"
                      stroke-width="2"
                    />
                  </svg>
                  {{ formatDate(m.date) }}
                </div>
              </div>
              <h3 class="list-title">{{ m.title }}</h3>
              <p class="list-description">{{ m.description }}</p>
            </div>
            <div class="list-glow"></div>
          </div>
        </div>

        <!-- CTA Button -->
        <div class="cta-section">
          <button class="cta-button" @click="goToLetter">
            <span class="cta-glow"></span>
            <span class="cta-content">
              <svg class="cta-icon" viewBox="0 0 24 24">
                <path
                  d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                />
              </svg>
              <span>อ่านจดหมาย</span>
            </span>
          </button>
        </div>

        <!-- Media Viewer Modal -->
        <div
          v-if="activeMoment"
          class="media-viewer-overlay"
          @click.self="closeViewer"
        >
          <div class="media-viewer-content">
            <button class="media-viewer-close" @click="closeViewer">✕</button>

            <div class="media-viewer-body">
              <button
                class="nav-btn nav-btn-left"
                @click.stop="prevMedia"
                v-if="activeMoment.images?.length > 1"
              >
                ‹
              </button>

              <div class="media-wrapper">
                <!-- แสดงรูปหรือวิดีโอ ตามชนิด -->
                <img
                  v-if="!isVideo(currentMediaUrl)"
                  :src="currentMediaUrl"
                  :alt="activeMoment.title"
                  class="media-image"
                />
                <video v-else class="media-video" controls autoplay>
                  <source :src="currentMediaUrl" />
                  Your browser does not support the video tag.
                </video>

                <div class="media-caption">
                  <div class="media-title">{{ activeMoment.title }}</div>
                  <div class="media-date">
                    {{ formatDate(activeMoment.date) }}
                  </div>
                  <div class="media-desc" v-if="activeMoment.description">
                    {{ activeMoment.description }}
                  </div>
                  <div
                    class="media-counter"
                    v-if="activeMoment.images?.length > 1"
                  >
                    {{ activeIndex + 1 }} / {{ activeMoment.images.length }}
                  </div>
                </div>
              </div>

              <button
                class="nav-btn nav-btn-right"
                @click.stop="nextMedia"
                v-if="activeMoment.images?.length > 1"
              >
                ›
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
  * {
    box-sizing: border-box;
  }

  .moments-page {
    min-height: 100vh;
    position: relative;
    overflow: hidden;
    color: #ffffff;
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Sukhumvit Set",
      sans-serif;
  }

  /* === Enhanced Background === */
  .bg-base {
    position: fixed;
    inset: 0;
    background: radial-gradient(
        circle at 20% 20%,
        rgba(139, 92, 246, 0.15) 0%,
        transparent 50%
      ),
      radial-gradient(
        circle at 80% 80%,
        rgba(236, 72, 153, 0.12) 0%,
        transparent 50%
      ),
      linear-gradient(
        135deg,
        #0a0118 0%,
        #1a0b2e 25%,
        #2d1b4e 50%,
        #1a0b2e 75%,
        #0a0118 100%
      );
    z-index: 0;
  }

  .bg-gradient {
    position: fixed;
    border-radius: 50%;
    filter: blur(80px);
    opacity: 0.4;
    mix-blend-mode: screen;
    z-index: 1;
  }

  .bg-gradient-1 {
    width: 600px;
    height: 600px;
    top: -200px;
    left: -200px;
    background: radial-gradient(
      circle,
      rgba(167, 139, 250, 0.6),
      transparent 70%
    );
    animation: float-1 20s ease-in-out infinite;
  }

  .bg-gradient-2 {
    width: 500px;
    height: 500px;
    bottom: -150px;
    right: -150px;
    background: radial-gradient(
      circle,
      rgba(244, 114, 182, 0.5),
      transparent 70%
    );
    animation: float-2 25s ease-in-out infinite;
  }

  .bg-gradient-3 {
    width: 400px;
    height: 400px;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background: radial-gradient(
      circle,
      rgba(59, 130, 246, 0.3),
      transparent 70%
    );
    animation: float-3 30s ease-in-out infinite;
  }

  .bg-noise {
    position: fixed;
    inset: 0;
    opacity: 0.03;
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
    width: 2px;
    height: 2px;
    background: radial-gradient(circle, rgba(255, 255, 255, 0.8), transparent);
    border-radius: 50%;
    animation: particle-drift 15s linear infinite;
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
    bottom: -50px;
    font-size: 1.5rem;
    opacity: 0;
    animation: float-up 15s ease-in infinite;
  }

  /* === Content === */
  .content-wrapper {
    position: relative;
    z-index: 10;
    min-height: 100vh;
    display: flex;
    justify-content: center;
    padding: 60px 24px;
  }

  .content-container {
    width: 100%;
    max-width: 1200px;
  }

  .header-section {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 32px;
    margin-bottom: 64px;
    animation: fade-in-up 0.8s ease-out;
  }

  .header-content {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .icon-container {
    position: relative;
    width: 120px;
    height: 120px;
    margin: 0 auto 32px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .icon-ring {
    position: absolute;
    border-radius: 50%;
    border: 2px solid;
    animation: ring-rotate 20s linear infinite;
  }

  .icon-ring-1 {
    inset: 0;
    border-color: rgba(167, 139, 250, 0.3);
    border-style: dashed;
  }

  .icon-ring-2 {
    inset: -12px;
    border-color: rgba(244, 114, 182, 0.2);
    border-style: dotted;
    animation-direction: reverse;
    animation-duration: 15s;
  }

  .icon-core {
    position: relative;
    width: 80px;
    height: 80px;
    display: flex;
    align-items: center;
    justify-content: center;
    background: radial-gradient(
      circle,
      rgba(139, 92, 246, 0.4),
      rgba(79, 70, 229, 0.2)
    );
    border-radius: 50%;
    box-shadow: 0 0 40px rgba(167, 139, 250, 0.6),
      inset 0 0 20px rgba(167, 139, 250, 0.3);
    animation: pulse-glow 3s ease-in-out infinite;
  }

  .icon-svg {
    width: 48px;
    height: 48px;
    stroke: #e9d5ff;
    stroke-width: 2;
    fill: none;
    stroke-linecap: round;
    stroke-linejoin: round;
  }

  .main-title {
    font-size: clamp(2.5rem, 5vw, 4rem);
    font-weight: 900;
    letter-spacing: -0.02em;
    text-align: center;
    margin: 0 0 16px;
    line-height: 1.1;
  }

  .title-word {
    display: inline-block;
    background: linear-gradient(135deg, #ffffff 0%, #e9d5ff 50%, #f9a8d4 100%);
    background-size: 200% 200%;
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
    animation: gradient-shift 8s ease infinite;
    text-shadow: 0 0 80px rgba(244, 114, 182, 0.5);
    margin: 0 8px;
  }

  .subtitle {
    font-size: 1.125rem;
    color: rgba(255, 255, 255, 0.75);
    text-align: center;
    margin: 0 0 24px;
    line-height: 1.6;
  }

  .meta-info {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 12px;
    font-size: 0.875rem;
    color: rgba(255, 255, 255, 0.5);
  }

  .meta-item {
    display: flex;
    align-items: center;
    gap: 6px;
  }

  .meta-icon {
    width: 16px;
    height: 16px;
    stroke-width: 2;
  }

  .meta-divider {
    opacity: 0.5;
  }

  /* === View Controls === */
  .view-controls {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    gap: 12px;
  }

  .toggle-group {
    display: flex;
    gap: 8px;
    padding: 4px;
    background: rgba(0, 0, 0, 0.3);
    border-radius: 12px;
    border: 1px solid rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
  }

  .toggle-btn {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 10px 16px;
    background: transparent;
    border: none;
    border-radius: 8px;
    color: rgba(255, 255, 255, 0.6);
    font-size: 0.875rem;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .toggle-btn svg {
    width: 16px;
    height: 16px;
    stroke: currentColor;
    fill: none;
    stroke-width: 2;
    stroke-linecap: round;
    stroke-linejoin: round;
  }

  .toggle-btn.active {
    background: linear-gradient(
      135deg,
      rgba(139, 92, 246, 0.4),
      rgba(236, 72, 153, 0.4)
    );
    color: #ffffff;
    box-shadow: 0 4px 12px rgba(139, 92, 246, 0.3);
  }

  .toggle-btn:hover:not(.active) {
    background: rgba(255, 255, 255, 0.05);
    color: rgba(255, 255, 255, 0.8);
  }

  /* === Loading === */
  .loading-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 120px 0;
  }

  .loading-spinner {
    position: relative;
    width: 80px;
    height: 80px;
  }

  .spinner-ring {
    position: absolute;
    inset: 0;
    border: 3px solid transparent;
    border-top-color: #a78bfa;
    border-radius: 50%;
    animation: spin 1.5s cubic-bezier(0.5, 0, 0.5, 1) infinite;
  }

  .spinner-ring:nth-child(2) {
    border-top-color: #ec4899;
    animation-delay: -0.3s;
  }

  .spinner-ring:nth-child(3) {
    border-top-color: #3b82f6;
    animation-delay: -0.6s;
  }

  .loading-text {
    margin-top: 24px;
    font-size: 1rem;
    color: rgba(255, 255, 255, 0.6);
    animation: pulse-text 2s ease-in-out infinite;
  }

  /* === Grid View === */
  .grid-view {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 24px;
    margin-bottom: 64px;
  }

  .grid-card {
    position: relative;
    aspect-ratio: 1;
    border-radius: 24px;
    overflow: hidden;
    background: rgba(0, 0, 0, 0.4);
    border: 1px solid rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    cursor: pointer;
    opacity: 0;
    transform: translateY(30px);
    animation: fade-in-up 0.6s ease-out forwards;
    transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .grid-card:hover {
    transform: translateY(-8px) scale(1.02);
    border-color: rgba(167, 139, 250, 0.4);
    box-shadow: 0 20px 60px rgba(0, 0, 0, 0.5),
      0 0 40px rgba(167, 139, 250, 0.3);
  }

  .card-shimmer {
    position: absolute;
    inset: -100%;
    background: linear-gradient(
      90deg,
      transparent,
      rgba(255, 255, 255, 0.1),
      transparent
    );
    transform: translateX(-100%);
    transition: transform 0.6s;
  }

  .grid-card:hover .card-shimmer {
    transform: translateX(100%);
  }

  .card-image-wrapper {
    position: relative;
    width: 100%;
    height: 100%;
    overflow: hidden;
  }

  .card-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform 0.6s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .grid-card:hover .card-image {
    transform: scale(1.1);
  }

  .image-overlay {
    position: absolute;
    inset: 0;
    background: linear-gradient(
      to top,
      rgba(0, 0, 0, 0.9) 0%,
      rgba(0, 0, 0, 0.5) 40%,
      transparent 100%
    );
  }

  .card-content {
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    padding: 24px;
    z-index: 2;
  }

  .card-tag {
    display: inline-block;
    padding: 4px 12px;
    background: rgba(139, 92, 246, 0.6);
    border-radius: 12px;
    font-size: 0.75rem;
    font-weight: 600;
    margin-bottom: 8px;
    backdrop-filter: blur(10px);
  }

  .card-title {
    font-size: 1.25rem;
    font-weight: 700;
    margin: 0 0 8px;
    line-height: 1.3;
    color: #ffffff;
  }

  .card-date {
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 0.875rem;
    color: rgba(255, 255, 255, 0.7);
  }

  .date-icon {
    width: 14px;
    height: 14px;
    stroke-width: 2;
  }

  .card-glow {
    position: absolute;
    bottom: -50%;
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
    transition: opacity 0.4s;
    pointer-events: none;
  }

  .grid-card:hover .card-glow {
    opacity: 1;
  }

  /* === List View === */
  .list-view {
    display: flex;
    flex-direction: column;
    gap: 24px;
    margin-bottom: 64px;
  }

  .list-card {
    position: relative;
    display: flex;
    gap: 24px;
    padding: 24px;
    background: rgba(0, 0, 0, 0.4);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 24px;
    backdrop-filter: blur(10px);
    cursor: pointer;
    opacity: 0;
    transform: translateY(30px);
    animation: fade-in-up 0.6s ease-out forwards;
    transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
    overflow: hidden;
  }

  .list-card:hover {
    transform: translateY(-4px);
    border-color: rgba(244, 114, 182, 0.4);
    box-shadow: 0 20px 60px rgba(0, 0, 0, 0.5),
      0 0 40px rgba(244, 114, 182, 0.3);
  }

  .list-shimmer {
    position: absolute;
    inset: -100%;
    background: linear-gradient(
      90deg,
      transparent,
      rgba(255, 255, 255, 0.08),
      transparent
    );
    transform: translateX(-100%);
    transition: transform 0.6s;
  }

  .list-card:hover .list-shimmer {
    transform: translateX(100%);
  }

  .list-image-wrapper {
    position: relative;
    width: 200px;
    min-width: 200px;
    aspect-ratio: 4/3;
    border-radius: 16px;
    overflow: hidden;
    background: rgba(0, 0, 0, 0.3);
  }

  .list-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform 0.6s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .list-card:hover .list-image {
    transform: scale(1.08);
  }

  .list-image-overlay {
    position: absolute;
    inset: 0;
    background: linear-gradient(
      to right,
      rgba(0, 0, 0, 0.6) 0%,
      transparent 100%
    );
  }

  .list-content {
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 12px;
  }

  .list-header {
    display: flex;
    align-items: center;
    gap: 12px;
    flex-wrap: wrap;
  }

  .list-tag {
    padding: 6px 14px;
    background: rgba(236, 72, 153, 0.6);
    border-radius: 20px;
    font-size: 0.75rem;
    font-weight: 600;
    backdrop-filter: blur(10px);
  }

  .list-date {
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 0.875rem;
    color: rgba(255, 255, 255, 0.6);
  }

  .list-title {
    font-size: 1.5rem;
    font-weight: 700;
    margin: 0;
    line-height: 1.3;
    color: #ffffff;
    letter-spacing: -0.01em;
  }

  .list-description {
    font-size: 1rem;
    line-height: 1.7;
    color: rgba(255, 255, 255, 0.75);
    margin: 0;
  }

  .list-glow {
    position: absolute;
    bottom: -50%;
    left: 50%;
    width: 60%;
    height: 60%;
    background: radial-gradient(
      circle,
      rgba(244, 114, 182, 0.4),
      transparent 70%
    );
    transform: translateX(-50%);
    opacity: 0;
    transition: opacity 0.4s;
    pointer-events: none;
  }

  .list-card:hover .list-glow {
    opacity: 1;
  }

  /* === CTA Button === */
  .cta-section {
    display: flex;
    justify-content: center;
    padding: 40px 0;
  }

  .cta-button {
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

  .cta-button:hover {
    transform: translateY(-4px) scale(1.05);
    box-shadow: 0 30px 60px rgba(139, 92, 246, 0.6),
      0 0 80px rgba(236, 72, 153, 0.5);
  }

  .cta-button:active {
    transform: translateY(-2px) scale(1.02);
  }

  .cta-glow {
    position: absolute;
    inset: -20px;
    background: radial-gradient(
      circle,
      rgba(236, 72, 153, 0.5),
      transparent 70%
    );
    opacity: 0.6;
    filter: blur(20px);
    animation: pulse-glow 3s ease-in-out infinite;
  }

  .cta-content {
    position: relative;
    display: flex;
    align-items: center;
    gap: 12px;
    z-index: 1;
  }

  .cta-icon {
    width: 24px;
    height: 24px;
    stroke: currentColor;
    fill: none;
    stroke-width: 2;
    stroke-linecap: round;
    stroke-linejoin: round;
    animation: icon-bounce 2s ease-in-out infinite;
  }

  /* === Animations === */
  @keyframes float-1 {
    0%,
    100% {
      transform: translate(0, 0) scale(1);
    }
    50% {
      transform: translate(100px, -100px) scale(1.1);
    }
  }

  @keyframes float-2 {
    0%,
    100% {
      transform: translate(0, 0) scale(1);
    }
    50% {
      transform: translate(-120px, 80px) scale(1.15);
    }
  }

  @keyframes float-3 {
    0%,
    100% {
      transform: translate(-50%, -50%) scale(1);
    }
    50% {
      transform: translate(-40%, -60%) scale(1.2);
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
      transform: translate(100px, -800px);
      opacity: 0;
    }
  }

  @keyframes float-up {
    0% {
      transform: translateY(0) rotate(0deg);
      opacity: 0;
    }
    10% {
      opacity: 1;
    }
    90% {
      opacity: 1;
    }
    100% {
      transform: translateY(-100vh) rotate(360deg);
      opacity: 0;
    }
  }

  @keyframes fade-in-up {
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  @keyframes ring-rotate {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }

  @keyframes pulse-glow {
    0%,
    100% {
      opacity: 0.6;
      transform: scale(1);
    }
    50% {
      opacity: 1;
      transform: scale(1.1);
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

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }

  @keyframes pulse-text {
    0%,
    100% {
      opacity: 0.6;
    }
    50% {
      opacity: 1;
    }
  }

  @keyframes icon-bounce {
    0%,
    100% {
      transform: translateX(0);
    }
    50% {
      transform: translateX(6px);
    }
  }

  /* === Responsive === */
  @media (max-width: 1024px) {
    .header-section {
      flex-direction: column;
      align-items: center;
      text-align: center;
    }

    .view-controls {
      align-items: center;
    }

    .grid-view {
      grid-template-columns: repeat(auto-fill, minmax(240px, 1fr));
      gap: 20px;
    }
  }

  @media (max-width: 768px) {
    .content-wrapper {
      padding: 40px 16px;
    }

    .main-title {
      font-size: 2.5rem;
    }

    .title-word {
      margin: 0 4px;
    }

    .subtitle {
      font-size: 1rem;
    }

    .grid-view {
      grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
      gap: 16px;
    }

    .list-card {
      flex-direction: column;
      gap: 16px;
    }

    .list-image-wrapper {
      width: 100%;
      min-width: 0;
    }

    .list-title {
      font-size: 1.25rem;
    }

    .list-description {
      font-size: 0.9375rem;
    }

    .cta-button {
      padding: 16px 36px;
      font-size: 1rem;
    }

    .icon-container {
      width: 100px;
      height: 100px;
      margin-bottom: 24px;
    }

    .icon-core {
      width: 70px;
      height: 70px;
    }

    .icon-svg {
      width: 40px;
      height: 40px;
    }
  }

  @media (max-width: 480px) {
    .grid-view {
      grid-template-columns: 1fr;
    }

    .main-title {
      font-size: 2rem;
    }

    .header-section {
      margin-bottom: 48px;
    }

    .toggle-group {
      width: 100%;
    }

    .toggle-btn {
      flex: 1;
      justify-content: center;
    }

    .card-content {
      padding: 16px;
    }

    .card-title {
      font-size: 1.125rem;
    }
  }

  /* === Media Viewer Modal === */
  .media-viewer-overlay {
    position: fixed;
    inset: 0;
    z-index: 50;
    background: rgba(3, 7, 18, 0.85);
    backdrop-filter: blur(10px);
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .media-viewer-content {
    position: relative;
    max-width: 960px;
    width: 100%;
    padding: 24px;
  }

  .media-viewer-close {
    position: absolute;
    top: 8px;
    right: 16px;
    border: none;
    background: transparent;
    color: #e5e7eb;
    font-size: 1.5rem;
    cursor: pointer;
    padding: 4px;
    line-height: 1;
    opacity: 0.8;
    transition: opacity 0.2s;
  }

  .media-viewer-close:hover {
    opacity: 1;
  }

  .media-viewer-body {
    display: flex;
    align-items: center;
    gap: 16px;
  }

  .media-wrapper {
    flex: 1;
    background: rgba(15, 23, 42, 0.9);
    border-radius: 24px;
    padding: 16px;
    box-shadow: 0 25px 50px rgba(0, 0, 0, 0.8);
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  .media-image,
  .media-video {
    max-height: 70vh;
    width: 100%;
    border-radius: 16px;
    object-fit: contain;
    background: #020617;
  }

  .media-caption {
    display: flex;
    flex-direction: column;
    gap: 4px;
    font-size: 0.9rem;
    color: #e5e7eb;
  }

  .media-title {
    font-weight: 700;
    font-size: 1.1rem;
  }

  .media-date {
    font-size: 0.8rem;
    opacity: 0.7;
  }

  .media-desc {
    margin-top: 4px;
    line-height: 1.5;
  }

  .media-counter {
    margin-top: 4px;
    font-size: 0.8rem;
    opacity: 0.7;
    align-self: flex-end;
  }

  .nav-btn {
    border: none;
    background: rgba(15, 23, 42, 0.8);
    color: #e5e7eb;
    width: 36px;
    height: 36px;
    border-radius: 999px;
    font-size: 1.5rem;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
    transition: background 0.2s, transform 0.2s;
  }

  .nav-btn:hover {
    background: rgba(37, 99, 235, 0.8);
    transform: translateY(-1px);
  }

  .nav-btn-left {
    margin-right: 4px;
  }

  .nav-btn-right {
    margin-left: 4px;
  }

  @media (max-width: 768px) {
    .media-viewer-content {
      padding: 16px;
    }

    .media-viewer-body {
      gap: 8px;
    }

    .media-wrapper {
      padding: 12px;
    }

    .media-image,
    .media-video {
      max-height: 60vh;
    }

    .nav-btn {
      display: none; /* มือถือเลื่อนด้วยนิ้วได้อยู่แล้ว */
    }
  }

  .top-back-button {
    position: fixed;
    top: 20px;
    left: 20px;
    z-index: 40;
    border: none;
    cursor: pointer;
    padding: 10px 14px;
    border-radius: 9999px;
    color: #fff;
    background: rgba(15, 23, 42, 0.65);
    backdrop-filter: blur(10px);
    border: 1px solid rgba(255, 255, 255, 0.12);
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.45),
      0 0 30px rgba(167, 139, 250, 0.25);
    transition: all 0.3s ease;
    overflow: hidden;
  }

  .top-back-button:hover {
    transform: translateY(-2px) scale(1.05);
    border-color: rgba(167, 139, 250, 0.5);
    box-shadow: 0 14px 40px rgba(0, 0, 0, 0.6),
      0 0 40px rgba(244, 114, 182, 0.35);
  }

  .top-back-button:active {
    transform: translateY(0) scale(0.98);
  }

  .top-back-glow {
    position: absolute;
    inset: -12px;
    background: radial-gradient(
      circle,
      rgba(139, 92, 246, 0.55),
      transparent 70%
    );
    opacity: 0.6;
    filter: blur(14px);
    animation: pulse-glow 3s ease-in-out infinite;
    pointer-events: none;
  }

  .top-back-content {
    position: relative;
    display: flex;
    align-items: center;
    gap: 8px;
    z-index: 1;
  }

  .top-back-icon {
    width: 18px;
    height: 18px;
    stroke: currentColor;
    fill: none;
    stroke-width: 2.5;
    stroke-linecap: round;
    stroke-linejoin: round;
    animation: arrow-bounce-left 2s ease-in-out infinite;
  }

  .top-back-text {
    font-size: 0.95rem;
    font-weight: 700;
    letter-spacing: 0.2px;
    opacity: 0.95;
  }

  /* เล็กลงบนมือถือ */
  @media (max-width: 768px) {
    .top-back-button {
      top: 14px;
      left: 14px;
      padding: 8px 12px;
    }
    .top-back-text {
      font-size: 0.85rem;
    }
    .top-back-icon {
      width: 16px;
      height: 16px;
    }
  }

  @keyframes arrow-bounce-left {
    0%,
    100% {
      transform: translateX(0);
    }
    50% {
      transform: translateX(-4px);
    }
  }

  .dual-images {
    display: grid;
    grid-template-columns: 1.35fr 1fr; /* ซ้ายใหญ่ ขวาเล็ก */
    width: 100%;
    height: 100%;
    gap: 6px;
  }

  .dual-left,
  .dual-right {
    position: relative;
    width: 100%;
    height: 100%;
    overflow: hidden;
    border-radius: 14px;
  }

  .dual-right {
    border-radius: 14px;
  }

  /* ให้รูป fit เต็มช่อง */
  .dual-left img,
  .dual-right img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  /* overlay ดำจางบนรูปขวา */
  .dual-overlay {
    position: absolute;
    inset: 0;
    background: rgba(0, 0, 0, 0.55);
    backdrop-filter: blur(1px);
  }

  /* ตัวเลข +N เหมือนตัวอย่าง */
  .dual-count {
    position: absolute;
    inset: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: clamp(1.4rem, 2.2vw, 2.3rem);
    font-weight: 900;
    color: #fff;
    text-shadow: 0 6px 18px rgba(0, 0, 0, 0.7);
    letter-spacing: 0.5px;
  }

  /* เวลา hover ให้รูปซ้ายขยายเบา ๆ */
  .grid-card:hover .dual-left img,
  .list-card:hover .dual-left img {
    transform: scale(1.08);
  }

  .grid-card:hover .dual-right img,
  .list-card:hover .dual-right img {
    transform: scale(1.05);
  }
</style>
