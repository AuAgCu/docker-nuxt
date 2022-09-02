const client = "";
const server = "http://go:3001"

export default defineNuxtPlugin((nuxtApp) => {
    const baseUrl = (url: string) => {
        const baseUrl = process.client ? client : server;
        return `${baseUrl}${url}`
    }
    nuxtApp.provide('baseUrl', baseUrl);
});