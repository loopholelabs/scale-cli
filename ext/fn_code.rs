
use signature::types;
use HttpFetch::types as fetch;
use HttpFetch::New;
use HttpFetch::HttpConnector;

pub fn scale(ctx: Option<&mut types::Context>) -> Result<Option<types::Context>, Box<dyn std::error::Error>> {

    let c = fetch::HttpConfig{ timeout: 10 };

    let fetcher = New(c);

    if fetcher.is_err() {
      // Return an error...
      let val = ctx.unwrap();
      val.my_string = "Error New err".to_string();
      return signature::next(Some(val));
    }

    let f = fetcher.unwrap();

    if f.is_none() {
      // Return an error...
      let val = ctx.unwrap();
      val.my_string = "Error New none".to_string();
      return signature::next(Some(val));
    }


    let res = f.unwrap().Fetch(fetch::ConnectionDetails{ url: "https://ifconfig.me".to_string() });

    if res.is_err() {
      // Return an error...
      let val = ctx.unwrap();
      val.my_string = "Error res err".to_string();
      return signature::next(Some(val));
    }

    let f1 = res.unwrap();

    if f1.is_none() {
      // Return an error...
      let val = ctx.unwrap();
      val.my_string = "Error res none".to_string();
      return signature::next(Some(val));
    }

    let r = f1.unwrap();

    let val = ctx.unwrap();

    let string = String::from_utf8(r.body);

    val.my_string = format!("Fetch extension StatusCode={} Body={}", r.status_code, string.unwrap()).to_string();

    return signature::next(Some(val));
}
