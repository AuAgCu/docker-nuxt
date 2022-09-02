import { defineNuxtConfig } from 'nuxt'

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    // TODO: SSRをONにするとプロキシがうまく動かない
    ssr: false,
    vite: {
        logLevel: 'error',
        ssr: {

        },
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