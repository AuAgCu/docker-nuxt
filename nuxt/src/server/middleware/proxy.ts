import { createProxyMiddleware } from 'http-proxy-middleware'

export default createProxyMiddleware('/api/**', {
    target: 'http://go:3001/',
    changeOrigin: true,
    ws: true
});