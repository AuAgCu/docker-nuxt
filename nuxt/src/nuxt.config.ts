import { defineNuxtConfig } from 'nuxt'

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    vite: {
        server: {
            hmr: {
                protocol: 'ws',
                host: 'localhost'
            }
        }
    }
})
