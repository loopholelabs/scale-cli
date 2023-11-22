
use signature::types;
use testext::types as fetch;
use testext::New;
use testext::HttpConnector;

pub fn scale(ctx: Option<types::Context>) -> Result<Option<types::Context>, Box<dyn std::error::Error>> {

    let c = fetch::HttpConfig{ timeout: 10 };

    let fetcher = New(c);

    if let Err(e) = fetcher {
      let mut val = ctx.unwrap();
      val.my_string = format!("Error New err {e}");
      return signature::next(Some(val));
    }

    let f = fetcher.unwrap();

    if f.is_none() {
      // Return an error...
      let mut val = ctx.unwrap();
      val.my_string = "Error New none".to_string();
      return signature::next(Some(val));
    }

    let res = f.unwrap().Fetch(fetch::ConnectionDetails{ url: "https://ifconfig.me".to_string() });

    if let Err(e) = res {
      let mut val = ctx.unwrap();
      val.my_string = format!("Error Fetch err {e}");
      return signature::next(Some(val));
    }
  
    let f1 = res.unwrap();

    if f1.is_none() {
      // Return an error...
      let mut val = ctx.unwrap();
      val.my_string = "Error Fetch none".to_string();
      return signature::next(Some(val));
    }

    let r = f1.unwrap();

    let mut val = ctx.unwrap();

    let string = String::from_utf8(r.body);

    val.my_string = format!("Fetch extension StatusCode={} Body={}", r.status_code, string.unwrap()).to_string();

    return signature::next(Some(val));
}
