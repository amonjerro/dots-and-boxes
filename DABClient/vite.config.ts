import tailwindcss from "@tailwindcss/vite";
import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";

export default defineConfig({
    server: {
        host: true,
        port: 5173,
        strictPort: true,
        origin: "http://192.168.40.225:5173",
    },
    plugins: [tailwindcss(), sveltekit()],
});
