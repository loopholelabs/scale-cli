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

import {ScaleFunc, V1Alpha, Go, Rust} from "@loopholelabs/scalefile/scalefunc";

import { New, Config, Scale } from "@loopholelabs/scale";

import { New as SigNew } from "@loopholelabs/signature";
/*
import { New as ExtNew } from "@loopholelabs/HttpFetch";
*/
import * as fs from 'fs';

let scale_runtime: Scale<SigNew>;

async function init() {
  try {
    console.log("Setting up scale app...");

    let data_go = fs.readFileSync("../local-testfngo-latest.scale");

    const sfn_go = ScaleFunc.Decode(data_go);

    let data_rs = fs.readFileSync("../local-testfnrs-latest.scale");
    const sfn_rs = ScaleFunc.Decode(data_rs);

    let config = new Config(SigNew);

    scale_runtime = await New(config);
  
  } catch(e) {
    console.log("Exception setting up scale functions", e);
  }  

  const ctx = SigNew();
  ctx.MyString = "Hello world";

  try {
    const i = await scale_runtime.Instance(undefined);
    i.Run(ctx);

    const td = new TextDecoder();

    console.log("Output was (" + td.decode(ctx.MyString) + ")");

  } catch(e) {
      // There was an error! Need to write an error to context etc
      console.log("Error from scale", e);   
  }
}

init();
