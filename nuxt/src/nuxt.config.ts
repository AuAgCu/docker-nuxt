import { defineNuxtConfig } from 'nuxt'

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    // TODO: SSRをONにするとプロキシがうまく動かない
    ssr: true,
    modules: ['nuxt-proxy'],
    // See options here https://github.com/chimurai/http-proxy-middleware#options
    proxy: {
        options: {
            target: process.env.BASE_URL,
            changeOrigin: true,
            pathFilter: [
                '/api'
            ],
        }
    },
    publicRuntimeConfig: {
        // BASE_URL: process.env.BASE_URL,
        BASE_URL: "http://localhost:3000",
    },
});