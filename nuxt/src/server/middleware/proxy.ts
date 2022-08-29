import { createProxyMiddleware } from 'http-proxy-middleware'

export default createProxyMiddleware('/api', {
    target: 'http://example.com',
    changeOrigin: true,
    ws: true
})