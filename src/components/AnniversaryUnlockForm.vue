<script setup>
  import { ref, computed, nextTick } from "vue"
  import { useMessage } from "naive-ui"

  const emit = defineEmits(["unlocked"])
  const message = useMessage()

  const day = ref("")
  const month = ref("")
  const year = ref("")
  const isLoading = ref(false)

  // effect state
  const shouldShake = ref(false)
  const successPulse = ref(false)

  const dayInputRef = ref(null)
  const monthInputRef = ref(null)
  const yearInputRef = ref(null)

  const TARGET_DAY = 8
  const TARGET_MONTH = 12
  const TARGET_YEAR = 2567

  const isDisabled = computed(
    () => !day.value || !month.value || !year.value || isLoading.value
  )

  // motion config สำหรับฟอร์ม (ใช้ @vueuse/motion)
  const formMotion = {
    initial: { opacity: 0, y: 30, scale: 0.94 },
    enter: {
      opacity: 1,
      y: 0,
      scale: 1,
      transition: {
        duration: 0.8,
        easing: "ease-out",
      },
    },
  }

  const handleInput = (value, field, maxLength) => {
    let cleanValue = value.replace(/[^0-9]/g, "")
    if (cleanValue.length > maxLength) {
      cleanValue = cleanValue.slice(0, maxLength)
    }

    if (field === "day") day.value = cleanValue
    if (field === "month") month.value = cleanValue
    if (field === "year") year.value = cleanValue

    if (cleanValue.length === maxLength) {
      const nextMap = {
        day: monthInputRef,
        month: yearInputRef,
      }
      const nextRef = nextMap[field]
      if (nextRef) {
        nextTick(() => {
          nextRef.value?.focus?.()
        })
      }
    }
  }

  const triggerShake = () => {
    shouldShake.value = false
    // force reflow เผื่อเคย shake ไปแล้ว
    requestAnimationFrame(() => {
      shouldShake.value = true
      setTimeout(() => {
        shouldShake.value = false
      }, 350)
    })
  }

  const triggerSuccessPulse = () => {
    successPulse.value = false
    requestAnimationFrame(() => {
      successPulse.value = true
      setTimeout(() => {
        successPulse.value = false
      }, 650)
    })
  }

  const handleSubmit = async () => {
    if (!day.value || !month.value || !year.value) {
      message.error("กรุณากรอกข้อมูลให้ครบทุกช่องนะคะ", { duration: 3000 })
      triggerShake()
      return
    }

    const dayNum = Number(day.value)
    const monthNum = Number(month.value)
    const yearNum = Number(year.value)

    if (dayNum < 1 || dayNum > 31) {
      message.error("วันที่ไม่ถูกต้อง (1-31)", { duration: 3000 })
      triggerShake()
      return
    }

    if (monthNum < 1 || monthNum > 12) {
      message.error("เดือนไม่ถูกต้อง (1-12)", { duration: 3000 })
      triggerShake()
      return
    }

    if (yearNum < 2500 || yearNum > 2600) {
      message.error("ปีไม่ถูกต้อง (พ.ศ. 2500-2600)", { duration: 3000 })
      triggerShake()
      return
    }

    isLoading.value = true

    await new Promise(resolve => setTimeout(resolve, 200))

    const isCorrect =
      dayNum === TARGET_DAY &&
      monthNum === TARGET_MONTH &&
      yearNum === TARGET_YEAR

    if (isCorrect) {
      message.success("ใช่เลย! 💖 ยินดีต้อนรับค่ะ", { duration: 1500 })
      isLoading.value = false
      triggerSuccessPulse()

      setTimeout(() => {
        emit("unlocked")
      }, 450)
    } else {
      message.warning("เหมือนจะไม่ใช่วันนั้นนะ ลองทบทวนดูอีกที 💭", {
        duration: 3000,
      })
      isLoading.value = false
      triggerShake()
    }
  }
</script>

<template>
  <!-- ใส่ v-motion + shake class -->
  <div
    class="form-container"
    v-motion="formMotion"
    :class="{ 'shake-error': shouldShake }"
  >
    <n-form @submit.prevent="handleSubmit">
      <n-space vertical :size="24">
        <n-grid :cols="3" :x-gap="12">
          <!-- Day -->
          <n-gi>
            <n-space vertical :size="8">
              <label
                class="block text-xs font-medium text-purple-200 text-center tracking-wide"
              >
                วันที่
              </label>
              <n-input
                ref="dayInputRef"
                :value="day"
                @update:value="val => handleInput(val, 'day', 2)"
                placeholder="00"
                maxlength="2"
                size="large"
                :disabled="isLoading"
                :input-props="{ inputmode: 'numeric' }"
                class="date-input"
              >
                <template #prefix>
                  <svg
                    class="w-4 h-4 text-purple-300"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
                    />
                  </svg>
                </template>
              </n-input>
            </n-space>
          </n-gi>

          <!-- Month -->
          <n-gi>
            <n-space vertical :size="8">
              <label
                class="block text-xs font-medium text-purple-200 text-center tracking-wide"
              >
                เดือน
              </label>
              <n-input
                ref="monthInputRef"
                :value="month"
                @update:value="val => handleInput(val, 'month', 2)"
                placeholder="00"
                maxlength="2"
                size="large"
                :disabled="isLoading"
                :input-props="{ inputmode: 'numeric' }"
                class="date-input"
              />
            </n-space>
          </n-gi>

          <!-- Year -->
          <n-gi>
            <n-space vertical :size="8">
              <label
                class="block text-xs font-medium text-purple-200 text-center tracking-wide"
              >
                ปี พ.ศ.
              </label>
              <n-input
                ref="yearInputRef"
                :value="year"
                @update:value="val => handleInput(val, 'year', 4)"
                placeholder="0000"
                maxlength="4"
                size="large"
                :disabled="isLoading"
                :input-props="{ inputmode: 'numeric' }"
                class="date-input"
              />
            </n-space>
          </n-gi>
        </n-grid>

        <div
          class="button-container"
          :class="{ 'success-pulse': successPulse }"
        >
          <n-button
            type="primary"
            size="large"
            :loading="isLoading"
            :disabled="isDisabled"
            block
            strong
            class="submit-button"
            @click="handleSubmit"
            attr-type="submit"
          >
            <template #icon>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                class="love-icon w-5 h-5 md:w-6 md:h-6"
              >
                <g
                  fill="none"
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                >
                  <path
                    stroke-dasharray="36"
                    stroke-dashoffset="36"
                    d="M13 4h7c.55 0 1 .45 1 1v14c0 .55-.45 1-1 1h-7"
                  >
                    <animate
                      fill="freeze"
                      attributeName="stroke-dashoffset"
                      dur="0.5s"
                      values="36;0"
                    />
                  </path>
                  <path
                    stroke-dasharray="14"
                    stroke-dashoffset="14"
                    d="M3 12h11.5"
                  >
                    <animate
                      fill="freeze"
                      attributeName="stroke-dashoffset"
                      begin="0.6s"
                      dur="0.2s"
                      values="14;0"
                    />
                  </path>
                  <path
                    stroke-dasharray="6"
                    stroke-dashoffset="6"
                    d="M14.5 12l-3.5-3.5M14.5 12l-3.5 3.5"
                  >
                    <animate
                      fill="freeze"
                      attributeName="stroke-dashoffset"
                      begin="0.8s"
                      dur="0.2s"
                      values="6;0"
                    />
                  </path>
                </g>
              </svg>
            </template>

            {{ isLoading ? "กำลังตรวจสอบ..." : "Confirm" }}
          </n-button>
        </div>
      </n-space>
    </n-form>
  </div>
</template>

<style scoped>
  .form-container {
    width: 100%;
    position: relative;
    /* glow ใต้ฟอร์มแบบเบา ๆ */
  }

  .form-container::before,
  .form-container::after {
    content: "";
    position: absolute;
    inset: auto;
    left: 50%;
    transform: translateX(-50%);
    width: 220px;
    height: 220px;
    border-radius: 999px;
    pointer-events: none;
    z-index: -1;
    background: radial-gradient(
      circle,
      rgba(236, 72, 153, 0.32),
      rgba(76, 29, 149, 0)
    );
    opacity: 0.9;
    filter: blur(28px);
    animation: form-glow 9s ease-in-out infinite alternate;
    will-change: transform, opacity;
  }

  .form-container::after {
    width: 160px;
    height: 160px;
    background: radial-gradient(
      circle,
      rgba(129, 140, 248, 0.5),
      rgba(76, 29, 149, 0)
    );
    animation-duration: 11s;
    animation-delay: -3s;
    opacity: 0.8;
  }

  /* container ตอน error – shake */
  .shake-error {
    animation: shake-error 0.3s ease-in-out;
    will-change: transform;
  }

  .button-container {
    margin-top: 8px;
    padding-top: 4px;
  }

  /* pulse ตอน success */
  .success-pulse {
    animation: success-pulse 0.65s ease-out;
    transform-origin: center;
    will-change: transform, box-shadow;
  }

  .submit-button {
    height: 3.5rem;
    font-size: 1.0625rem;
    font-weight: 600;
    background: linear-gradient(135deg, #a855f7 0%, #ec4899 100%);
    border: none;
    transition: transform 0.18s ease, box-shadow 0.18s ease,
      background 0.18s ease;
    box-shadow: 0 8px 18px -4px rgba(168, 85, 247, 0.35);
    will-change: transform, box-shadow;
  }

  .submit-button:hover:not(:disabled) {
    transform: translateY(-1px) scale(1.01);
    box-shadow: 0 12px 26px -6px rgba(168, 85, 247, 0.6);
    background: linear-gradient(135deg, #9333ea 0%, #db2777 100%);
  }

  .submit-button:active:not(:disabled) {
    transform: translateY(0);
    box-shadow: 0 4px 12px -2px rgba(168, 85, 247, 0.4);
  }

  .submit-button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  /* icon หัวใจเต้น ๆ */
  .love-icon {
    animation: heartbeat 1.8s ease-in-out infinite;
    transform-origin: center;
    opacity: 0.9;
  }

  /* Date Input Styles */
  :deep(.date-input .n-input) {
    background: rgba(17, 24, 39, 0.5);
    border: 1.5px solid rgba(168, 85, 247, 0.3);
    border-radius: 12px;
    transition: border-color 0.18s ease, box-shadow 0.18s ease,
      background 0.18s ease, transform 0.18s ease;
    backdrop-filter: blur(8px);
    will-change: transform, box-shadow, border-color, background-color;
  }

  :deep(.date-input .n-input:hover) {
    border-color: rgba(168, 85, 247, 0.5);
    background: rgba(17, 24, 39, 0.6);
    box-shadow: 0 4px 10px rgba(168, 85, 247, 0.25);
    transform: translateY(-1px);
  }

  :deep(.date-input .n-input.n-input--focus) {
    border-color: #a855f7;
    box-shadow: 0 0 0 2px rgba(168, 85, 247, 0.2);
    background: rgba(17, 24, 39, 0.7);
    transform: translateY(-1px) scale(1.01);
  }

  :deep(.date-input .n-input input) {
    color: white;
    text-align: center;
    font-size: 1.2rem;
    font-weight: 600;
    letter-spacing: 0.05em;
  }

  :deep(.date-input .n-input input::placeholder) {
    color: rgba(196, 181, 253, 0.4);
    font-weight: 500;
  }

  /* keyframes */

  @keyframes shake-error {
    0% {
      transform: translateX(0);
    }
    20% {
      transform: translateX(-8px);
    }
    40% {
      transform: translateX(8px);
    }
    60% {
      transform: translateX(-6px);
    }
    80% {
      transform: translateX(4px);
    }
    100% {
      transform: translateX(0);
    }
  }

  @keyframes success-pulse {
    0% {
      transform: scale(1);
      box-shadow: 0 0 0 0 rgba(236, 72, 153, 0);
    }
    40% {
      transform: scale(1.04);
      box-shadow: 0 0 0 12px rgba(236, 72, 153, 0.3);
    }
    100% {
      transform: scale(1);
      box-shadow: 0 0 0 0 rgba(236, 72, 153, 0);
    }
  }

  @keyframes heartbeat {
    0%,
    100% {
      transform: scale(1);
    }
    15% {
      transform: scale(1.2);
    }
    30% {
      transform: scale(0.96);
    }
    45% {
      transform: scale(1.14);
    }
    60% {
      transform: scale(0.98);
    }
  }

  @keyframes form-glow {
    0% {
      opacity: 0.5;
      transform: translate3d(-6px, 10px, 0) scale(1);
    }
    50% {
      opacity: 0.9;
      transform: translate3d(4px, 0, 0) scale(1.08);
    }
    100% {
      opacity: 0.6;
      transform: translate3d(-2px, -8px, 0) scale(0.98);
    }
  }

  @media (max-width: 768px) {
    .submit-button {
      height: 3.25rem;
      font-size: 1rem;
    }

    :deep(.date-input .n-input input) {
      font-size: 1.05rem;
    }

    .form-container::before,
    .form-container::after {
      width: 160px;
      height: 160px;
    }
  }
</style>
