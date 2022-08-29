import type { IncomingMessage, ServerResponse } from "http";

export default async (req: IncomingMessage, res: ServerResponse) => {
    console.log("hogehoge")
    if (req.method != 'GET') {
      console.log(req.method)
      res.statusCode = 500
    }
    res.end(JSON.stringify([{firstName: "first", lastName: "lastName"}]))
}