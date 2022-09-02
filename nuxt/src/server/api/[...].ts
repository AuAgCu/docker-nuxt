import type { IncomingMessage, ServerResponse } from 'http'
import httpProxy from 'http-proxy'

const proxy = httpProxy.createProxyServer({
    target: `http://go:3001/`,
    changeOrigin: true,
})

// export default defineEventHandler((event) => {
//     const req = event.req;
//     const res = event.res;
//     const prefix = '/api'
//     if (req.url.startsWith(prefix)) {
//         const extraOptions = {

//         }
//         proxy.web(req, res, extraOptions);
//         return [{ firstName: "hoge", lastName: "gaga" }]
//     } else {
//         console.log("hugehuge")
//         // next()
//     }
// })

export default function (req: IncomingMessage, res: ServerResponse) {
    console.log("catch")
    console.log(req.url);
    // const prefix = '/api'


    // console.log(req);
    // if (req.url.startsWith(prefix)) {
    //     const extraOptions = {

    //     }
    //     proxy.web(req, res, extraOptions);
    //     console.log(typeof next);

    //     // return [{firstName: "firstName", lastName: "hoge"}];
    // }
}