import { createProxyMiddleware } from 'http-proxy-middleware'


export default createProxyMiddleware('/api', {
    target: 'http://go:3001',
    changeOrigin: true,
    ws: true,
    pathRewrite: { "^/api/": "/api/" },
    logLevel: 'debug',
    secure: false,
});

// export default hoge;


// export default function() {
//     if (!process.client) {
//         console.log("hogehoge")
//         const hoge = createProxyMiddleware('/api', {
//             target: 'http://go:3001',
//             // changeOrigin: true,
//             // ws: true,
//             pathRewrite: { "^/api/": "/api/" },
//             logLevel: 'debug',
//             // secure: false,
//         });

//         hoge
//     }
// }

