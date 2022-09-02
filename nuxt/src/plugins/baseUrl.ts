export default defineNuxtPlugin((nuxtApp) => {
    const client = "";
    const server = nuxtApp.$config.BASE_URL;
    const baseUrl = (url: string) => {
        const baseUrl = process.client ? client : server;
        return `${baseUrl}${url}`
    }
    nuxtApp.provide('baseUrl', baseUrl);
});