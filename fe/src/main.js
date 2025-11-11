import { createApp } from "vue";
import { MotionPlugin } from "@vueuse/motion";

import App from "./App.vue";
import router from "./router";

import "./style.css";

import naive from "naive-ui";

import "vfonts/Lato.css";
import "vfonts/FiraCode.css";

const app = createApp(App);

app.use(MotionPlugin);
app.use(router);
app.use(naive);

app.mount("#app");
