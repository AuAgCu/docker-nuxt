import { defineNuxtConfig } from 'nuxt'

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    ssr: true,
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
    // modules: ['@nuxtjs-alt/proxy',],
    // proxy: {
    //     '/api': {
    //         target: 'http://go:3001/',
    //     },
    // }
});