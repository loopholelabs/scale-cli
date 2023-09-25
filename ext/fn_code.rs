
use signature::types;
use HttpFetch::types as fetch;

pub fn scale(ctx: Option<&mut types::Context>) -> Result<Option<types::Context>, Box<dyn std::error::Error>> {

    let c = fetch::HttpConfig{ timeout: 10 };

    let fetcher = fetch::New(c);

    if fetcher.is_none() {
      // Return an error...
      let val = ctx.unwrap();
      val.my_string = "Error".to_string();
      return signature::next(Some(val));
    }

    let res = fetcher.unwrap().Fetch(fetch::ConnectionDetails{ url: "https://ifconfig.me".to_string() });

    if res.is_none() {
      // Return an error...
      let val = ctx.unwrap();
      val.my_string = "Error".to_string();
      return signature::next(Some(val));
    }

    let val = ctx.unwrap();

    val.my_string = format!("Fetch extension StatusCode={} Body={}", res.StatusCode, res.Body).to_string();

    return signature::next(Some(val));
}
