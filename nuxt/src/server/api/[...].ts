import type { IncomingMessage, ServerResponse } from "http";


export default async (req: IncomingMessage, res: ServerResponse) => {
    const proxyRequestHeaders = Object.assign({}, req.headers);

    const proxyUrl = 'http://go:3001/';
};