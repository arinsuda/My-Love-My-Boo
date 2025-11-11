<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
const moments = ref([]);
const loading = ref(true);

const loadMoments = async () => {
  try {
    const res = await fetch("/api/moments");
    if (!res.ok) throw new Error("load fail");
    moments.value = await res.json();
  } catch {
    moments.value = [
      {
        id: 1,
        date: "2019-08-10",
        title: "วันแรกที่เราเจอกัน",
        description: "ตอนนั้นยังไม่รู้เลยว่าจะกลายเป็นคนสำคัญขนาดนี้",
        images: ["/images/first-day-1.jpg"],
        tag: "เริ่มต้น",
      },
      {
        id: 2,
        date: "2020-02-14",
        title: "วาเลนไทน์ครั้งแรก",
        description: "ช็อกโกแลตวันนั้น กับรอยยิ้มของคุณ ยังจำได้อยู่เลย",
        images: ["/images/valentine-2020.jpg"],
        tag: "เดต",
      },
    ];
  } finally {
    loading.value = false;
  }
};

onMounted(loadMoments);

const goToLetter = () => {
  router.push("/letter");
};
</script>

<template>
  <div class="min-h-screen relative overflow-hidden">
    <!-- Background -->
    <div
      class="absolute inset-0 bg-gradient-to-b from-purple-950 via-black to-purple-900"
    >
      <div
        class="absolute top-0 left-1/4 w-96 h-96 bg-purple-600/10 rounded-full blur-3xl"
      ></div>
      <div
        class="absolute bottom-0 right-1/4 w-96 h-96 bg-pink-600/10 rounded-full blur-3xl"
      ></div>
    </div>

    <!-- Content -->
    <div class="relative z-10 px-6 py-12">
      <!-- Header -->
      <div class="text-center mb-12">
        <div class="inline-block mb-4">
          <n-icon size="64" color="#a78bfa">
            <svg viewBox="0 0 24 24">
              <path
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
          </n-icon>
        </div>
        <h2
          class="text-4xl md:text-5xl font-bold bg-gradient-to-r from-purple-400 to-pink-400 bg-clip-text text-transparent mb-4"
        >
          ความทรงจำของเรา
        </h2>
        <n-text class="text-purple-300 text-lg">
          ทุกช่วงเวลาที่เราสร้างขึ้นมาด้วยกัน
        </n-text>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="flex justify-center items-center py-20">
        <n-spin size="large" />
      </div>

      <!-- Moments Grid -->
      <div v-else class="max-w-6xl mx-auto mb-12">
        <n-grid :cols="1" :md-cols="2" :x-gap="24" :y-gap="24">
          <n-gi v-for="m in moments" :key="m.id">
            <n-card
              hoverable
              class="moment-card group"
              content-style="padding: 0;"
            >
              <!-- Image -->
              <div
                v-if="m.images && m.images.length"
                class="relative h-64 overflow-hidden"
              >
                <img
                  :src="m.images[0]"
                  alt="moment image"
                  class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-500"
                />
                <div
                  class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent"
                ></div>

                <!-- Tag -->
                <div class="absolute top-4 right-4">
                  <n-tag type="primary" round>
                    {{ m.tag }}
                  </n-tag>
                </div>
              </div>

              <!-- Content -->
              <div class="p-6 space-y-3">
                <n-space align="center" :size="8">
                  <n-icon size="16" color="#d8b4fe">
                    <svg viewBox="0 0 24 24">
                      <path
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
                      />
                    </svg>
                  </n-icon>
                  <n-text depth="3">{{ m.date }}</n-text>
                </n-space>

                <h3
                  class="text-2xl font-bold text-white group-hover:text-purple-300 transition-colors"
                >
                  {{ m.title }}
                </h3>

                <n-text class="text-purple-200 leading-relaxed">
                  {{ m.description }}
                </n-text>
              </div>
            </n-card>
          </n-gi>
        </n-grid>
      </div>

      <!-- Navigation Button -->
      <div class="flex justify-center">
        <n-button
          type="primary"
          size="large"
          @click="goToLetter"
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
                d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
              />
            </svg>
          </template>
          อ่านจดหมายจากใจ
        </n-button>
      </div>
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

:deep(.moment-card) {
  background: linear-gradient(
    to bottom right,
    rgba(88, 28, 135, 0.4),
    rgba(0, 0, 0, 0.4)
  );
  border: 1px solid rgba(139, 92, 246, 0.2);
  backdrop-filter: blur(8px);
  transition: all 0.3s ease;
  overflow: hidden;
}

:deep(.moment-card:hover) {
  border-color: rgba(139, 92, 246, 0.5);
  box-shadow: 0 20px 40px -10px rgba(139, 92, 246, 0.2);
  transform: scale(1.02);
}

:deep(.n-tag) {
  background-color: rgba(147, 51, 234, 0.8);
  backdrop-filter: blur(4px);
}
</style>
