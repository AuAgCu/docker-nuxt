import { defineNuxtConfig } from 'nuxt'

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    modules: ['nuxt-vite'],
    vite: {
        server: {
            proxy: {
                "/api/": {
                    target: "http://go:3001/",
                    secure: false
                }
            }
        }
    },
});