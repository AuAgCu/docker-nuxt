import { defineNuxtConfig } from 'nuxt'

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    // TODO: SSRをONにするとプロキシがうまく動かない
    ssr: true,
    vite: {
        logLevel: 'error',
        ssr: {

        },
        server: {
            proxy: {
                "/api/": {
                    target: process.env.BASE_URL,
                    secure: false
                }
            }
        }
    },
    publicRuntimeConfig: {
        BASE_URL: process.env.BASE_URL,
    },
    // modules: ['@nuxtjs-alt/proxy',],
    // proxy: {
    //     '/api': {
    //         target: 'http://go:3001/',
    //     },
    // }
});