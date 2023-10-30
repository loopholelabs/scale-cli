import * as signature from "signature";
import { HttpConfig, ConnectionDetails, HttpResponse } from "HttpFetch/types";

import { New } from "HttpFetch";

export function scale(ctx?: signature.Context): signature.Context | undefined {

  try {
    let conf = new HttpConfig();
    conf.timeout = 999;

    let fetcher = New(conf);

    let details = new ConnectionDetails();
    details.url = "https://ifconfig.me";

    let resp = fetcher.Fetch(details);

    let tdec = new TextDecoder();
    ctx.myString = "Typescript says hi: " + tdec.decode(resp.body);
  } catch(e) {
    ctx.myString = "Error in ext: " + e;
  }
  return signature.Next(ctx);
}
