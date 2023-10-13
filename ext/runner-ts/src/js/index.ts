/*
	Copyright 2022 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

import { New, Config, Scale } from "@loopholelabs/scale";
import { ScaleFunc } from "@loopholelabs/scale/scalefunc/scalefunc"
import { Extension, InstallableFunc, ModuleMemory, Resizer } from "@loopholelabs/scale-extension-interfaces";

import { New as SigNew, Signature } from "@loopholelabs/signature";
import { Decoder, Encoder, Kind } from "@loopholelabs/polyglot";

import { HttpConfig, HttpResponse, ConnectionDetails } from "@loopholelabs/HttpFetch";

import { HttpFetchIfc, HttpConnector, New as FetchNew } from "@loopholelabs/HttpFetch";

let scale_runtime: Scale<Signature>;

// An HttpConnector
class HttpConnectorImpl implements HttpConnector {
  public Fetch(d: ConnectionDetails): HttpResponse {
    console.log("#EXT# HttpConnector.Fetch()", d);
    // TODO: Implement actual fetch logic here...
    const r = new HttpResponse();

    const enc = new TextEncoder();
    r.statusCode = 200;
    r.body = enc.encode("Hello back");
    return r;
  }
}

// Make the extension implementation
class Fetcher implements HttpFetchIfc {
  public New(c: HttpConfig): HttpConnector {
    console.log("#EXT# New()", c);
    return new HttpConnectorImpl();
  }
}

// Make something that implements the bits from Extension
class Ext implements Extension {
  implementation: HttpFetchIfc;
  instances_map: HttpConnectorImpl[];

  constructor(i: HttpFetchIfc) {
    this.implementation = i;
    this.instances_map = [];
  }

  public Init(): Map<string, InstallableFunc> {
    let fns = new Map<string, InstallableFunc>();

// Add function for
// ext_5c7d22390f9101d459292d76c11b5e9f66c327b1766aae34b9cc75f9f40e8206_HttpConnector_Fetch

    const fetchfn = function(t) {
      return (mem: ModuleMemory, resize: Resizer, params: number[]) => {
        const id = params[0];
        const ii = t.instances_map[id];
        const d = mem.Read(params[1], params[2]);
        const c = ConnectionDetails.decode(new Decoder(d));
        if (c==undefined) {
          throw new Error("ConnectionDetails is undefined");
        }

        const r = ii.Fetch(c);
        const e = new Encoder();
        r.encode(e);
        const data = e.bytes;

        const ptr = resize("ext_5c7d22390f9101d459292d76c11b5e9f66c327b1766aae34b9cc75f9f40e8206_Resize", data.length);
  
        mem.Write(ptr, data);

        params[0] = 0;
      }
    }(this)

    fns.set("ext_5c7d22390f9101d459292d76c11b5e9f66c327b1766aae34b9cc75f9f40e8206_HttpConnector_Fetch", fetchfn);

    let newfn = function(t) {
      return (mem: ModuleMemory, resize: Resizer, params: number[]) => {
        const d = mem.Read(params[1], params[2]);
        const c = HttpConfig.decode(new Decoder(d));

        const con = t.implementation.New(c);
        // Store it in instances map...

        const id = t.instances_map.length;

        t.instances_map.push(con);

        // throw new Error("Testing error...");
        // Return an instance
        params[0] = id;
      }
    }(this);

    fns.set("ext_5c7d22390f9101d459292d76c11b5e9f66c327b1766aae34b9cc75f9f40e8206_New", newfn);


    console.log("Adding some functions into the mix...", fns);

    return fns;
  }

  public Reset() {

  }
}

async function init() {
  try {
    console.log("Setting up scale app...");

    const mod_go = await fetch("local-testfngo-latest.scale");
    const data_go = Buffer.from(await mod_go.arrayBuffer());

    console.log("Scale function go=" + data_go.length)

    console.log("ScaleFunc is ", ScaleFunc);

    const sfn_go = ScaleFunc.Decode(data_go);

    const mod_rs = await fetch("local-testfnrs-latest.scale");
    const data_rs = Buffer.from(await mod_rs.arrayBuffer());

    console.log("Scale function rs=" + data_rs.length)

    const sfn_rs = ScaleFunc.Decode(data_rs);

    let impl = new Fetcher();
    let ex = new Ext(impl);

    let ex2 = FetchNew(impl);

    console.log("Old fetch", ex);
    console.log("New fetch", ex2);

    let config = new Config(SigNew);
    config.WithFunctions([sfn_go]);
    config.WithExtension(ex2);

    scale_runtime = await New(config);
  
  } catch(e) {
    console.log("Exception setting up scale functions", e);
  }  

  const ctx = SigNew();
  ctx.context.myString = "Hello world";

  try {
    const i = await scale_runtime.Instance(undefined);
    await i.Run(ctx);

    console.log("Output was (" + ctx.context.myString + ")");

  } catch(e) {
      // There was an error! Need to write an error to context etc
      console.log("Error from scale", e);   
  }
}

init();
