
use signature::types;

pub fn scale(ctx: Option<&mut types::Context>) -> Result<Option<types::Context>, Box<dyn std::error::Error>> {
    return signature::next(ctx);
}
