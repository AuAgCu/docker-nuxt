// server/middleware/proxy.ts
import { createProxyMiddleware } from 'http-proxy-middleware'

export default defineEventHandler((event) => {
    console.log('New request: ' + event.req.url);
})