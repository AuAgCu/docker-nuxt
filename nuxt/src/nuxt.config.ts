import { defineNuxtConfig } from 'nuxt'

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    vite: {
        server: {
            proxy: {
                "/api/": {
                    target: process.env.BACKEND_HOST,
                    secure: false
                }
            }
        }
    },
})
