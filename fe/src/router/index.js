import { createRouter, createWebHistory } from "vue-router";

import AnniversaryGate from "@/views/AnniversaryGate.vue";
import WelcomePage from "@/views/WelcomePage.vue";
import MomentsPage from "@/views/MomentsPage.vue";
import LetterPage from "@/views/LetterPage.vue";

const routes = [
  { path: "/", name: "Gate", component: AnniversaryGate },
  { path: "/welcome", name: "Welcome", component: WelcomePage },
  { path: "/moments", name: "Moments", component: MomentsPage },
  { path: "/letter", name: "Letter", component: LetterPage },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
