import type { IncomingMessage, ServerResponse } from 'http'

export default async (req: IncomingMessage, _res: ServerResponse) => {
  console.log(`${new Date().toLocaleString()}:`, req.method)
  console.log(`${new Date().toLocaleString()}:`, req.url)

  // jwtTokenをここでつける予定、req.urlにapiがついているときにやるでもいいかも
  req.headers.authorization = "Bearer eyJhbGciOiJIUz";
}