<script setup>
import { ref, computed, nextTick } from "vue";
import { useMessage } from "naive-ui";

const emit = defineEmits(["unlocked"]);
const message = useMessage();

const day = ref("");
const month = ref("");
const year = ref("");
const isLoading = ref(false);

const dayInputRef = ref(null);
const monthInputRef = ref(null);
const yearInputRef = ref(null);

const TARGET_DAY = 8;
const TARGET_MONTH = 12;
const TARGET_YEAR = 2567;

const isDisabled = computed(
  () => !day.value || !month.value || !year.value || isLoading.value
);

const handleInput = (value, field, maxLength) => {
  let cleanValue = value.replace(/[^0-9]/g, "");
  if (cleanValue.length > maxLength) {
    cleanValue = cleanValue.slice(0, maxLength);
  }

  if (field === "day") day.value = cleanValue;
  if (field === "month") month.value = cleanValue;
  if (field === "year") year.value = cleanValue;

  if (cleanValue.length === maxLength) {
    const nextMap = {
      day: monthInputRef,
      month: yearInputRef,
    };
    const nextRef = nextMap[field];
    if (nextRef) {
      nextTick(() => {
        nextRef.value?.focus?.();
      });
    }
  }
};

const handleSubmit = async () => {
  if (!day.value || !month.value || !year.value) {
    message.error("กรุณากรอกข้อมูลให้ครบทุกช่องนะคะ", { duration: 3000 });
    return;
  }

  const dayNum = Number(day.value);
  const monthNum = Number(month.value);
  const yearNum = Number(year.value);

  if (dayNum < 1 || dayNum > 31) {
    message.error("วันที่ไม่ถูกต้อง (1-31)", { duration: 3000 });
    return;
  }

  if (monthNum < 1 || monthNum > 12) {
    message.error("เดือนไม่ถูกต้อง (1-12)", { duration: 3000 });
    return;
  }

  if (yearNum < 2500 || yearNum > 2600) {
    message.error("ปีไม่ถูกต้อง (พ.ศ. 2500-2600)", { duration: 3000 });
    return;
  }

  isLoading.value = true;

  // ลด delay จาก 800ms → 200ms ให้รู้สึกตอบสนองไวขึ้น
  await new Promise((resolve) => setTimeout(resolve, 200));

  const isCorrect =
    dayNum === TARGET_DAY &&
    monthNum === TARGET_MONTH &&
    yearNum === TARGET_YEAR;

  if (isCorrect) {
    message.success("ใช่เลย! 💖 ยินดีต้อนรับค่ะ", { duration: 1500 });
    isLoading.value = false;

    // ไม่ต้องรอ 1.5 วิ เต็ม ๆ แล้วค่อยเปลี่ยนหน้า → ให้ route เร็วขึ้น
    setTimeout(() => {
      emit("unlocked");
    }, 400);
  } else {
    message.warning("เหมือนจะไม่ใช่วันนั้นนะ ลองทบทวนดูอีกที 💭", {
      duration: 3000,
    });
    isLoading.value = false;
  }
};
</script>

<template>
  <div class="form-container">
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
                @update:value="(val) => handleInput(val, 'day', 2)"
                placeholder="08"
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
                @update:value="(val) => handleInput(val, 'month', 2)"
                placeholder="12"
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
                @update:value="(val) => handleInput(val, 'year', 4)"
                placeholder="2567"
                maxlength="4"
                size="large"
                :disabled="isLoading"
                :input-props="{ inputmode: 'numeric' }"
                class="date-input"
              />
            </n-space>
          </n-gi>
        </n-grid>

        <div class="button-container">
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
              <img
                src="/love-svgrepo-com.svg"
                alt="love icon"
                class="w-5 h-5 md:w-6 md:h-6 opacity-90"
              />
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
}

.button-container {
  margin-top: 8px;
  padding-top: 4px;
}

/* ปุ่ม: ลด animation เหลือ hover ธรรมดา */
.submit-button {
  height: 3.5rem;
  font-size: 1.0625rem;
  font-weight: 600;
  background: linear-gradient(135deg, #a855f7 0%, #ec4899 100%);
  border: none;
  transition: transform 0.18s ease, box-shadow 0.18s ease, background 0.18s ease;
  box-shadow: 0 8px 18px -4px rgba(168, 85, 247, 0.35);
}

.submit-button:hover:not(:disabled) {
  transform: translateY(-1px);
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

/* Date Input Styles (ไม่ใส่ animation keyframes แล้ว) */
:deep(.date-input .n-input) {
  background: rgba(17, 24, 39, 0.5);
  border: 1.5px solid rgba(168, 85, 247, 0.3);
  border-radius: 12px;
  transition: border-color 0.18s ease, box-shadow 0.18s ease,
    background 0.18s ease, transform 0.18s ease;
  backdrop-filter: blur(8px);
}

:deep(.date-input .n-input:hover) {
  border-color: rgba(168, 85, 247, 0.5);
  background: rgba(17, 24, 39, 0.6);
  box-shadow: 0 4px 10px rgba(168, 85, 247, 0.25);
}

:deep(.date-input .n-input.n-input--focus) {
  border-color: #a855f7;
  box-shadow: 0 0 0 2px rgba(168, 85, 247, 0.2);
  background: rgba(17, 24, 39, 0.7);
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

@media (max-width: 768px) {
  .submit-button {
    height: 3.25rem;
    font-size: 1rem;
  }

  :deep(.date-input .n-input input) {
    font-size: 1.05rem;
  }
}
</style>
