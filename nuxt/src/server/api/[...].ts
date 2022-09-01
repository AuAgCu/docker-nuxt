import type { IncomingMessage, ServerResponse } from "http";
import request from 'request';


export default async (req: IncomingMessage, res: ServerResponse) => {
    const proxyRequestHeaders = Object.assign({}, req.headers);

    const proxyUrl = 'http://go:3001/';
    request({
        url: proxyUrl,
        method: req.method,
        headers: proxyRequestHeaders,
    }).pipe(res);
};