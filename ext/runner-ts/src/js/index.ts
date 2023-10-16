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

let error_new = false;
let error_fetch = false;

// An HttpConnector
class HttpConnectorImpl implements HttpConnector {
  public Fetch(d: ConnectionDetails): HttpResponse {
    if (error_fetch) {
      throw new Error("Error from fetch");
    }
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
    if (error_new) {
      throw new Error("Error from new");
    }
    console.log("#EXT# New()", c);
    return new HttpConnectorImpl();
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

    let ex = FetchNew(impl);

  // First try go
  error_new = false;
  error_fetch = false;
  console.log("\nRunning go tests...\n");
  await runTest([sfn_go], ex);
  error_new = true;
  console.log("\nRunning go tests [ error from new ]...\n");
  await runTest([sfn_go], ex);
  error_new = false;
  error_fetch = true;
  console.log("\nRunning go tests [ error from fetch ]...\n");
  await runTest([sfn_go], ex);
  
  error_new = false;
  error_fetch = false;
  console.log("\nRunning rust tests...\n");
  await runTest([sfn_rs], ex);
  error_new = true;
  console.log("\nRunning rust tests [ error from new ]...\n");
  await runTest([sfn_rs], ex);
  error_new = false;
  error_fetch = true;
  console.log("\nRunning rust tests [ error from fetch ]...\n");
  await runTest([sfn_rs], ex);

  } catch(e) {
    console.log("Exception setting up scale functions", e);
  }
}

async function runTest(functions: ScaleFunc[], ex: Extension) {
  try {
    // First try go
    let config = new Config(SigNew);
    config.WithFunctions(functions);
    config.WithExtension(ex);
    scale_runtime = await New(config);

    const ctx = SigNew();
    ctx.context.myString = "Hello world";

    const i = await scale_runtime.Instance(undefined);
    await i.Run(ctx);

    console.log("Output was (" + ctx.context.myString + ")");
  } catch(e) {
      // There was an error! Need to write an error to context etc
      console.log("Error from scale", e);   
  }
}

init();
