#![allow(unused_mut)]

#[path = "signature/signature.rs"]
mod signature

use signature::Context;

pub fn scale (mut context: Context) -> Context {
    return context
}