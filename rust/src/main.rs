#[macro_use] extern crate rocket;

extern crate mongodb;
use rocket::response::content;
use mongodb::{Client, bson::doc, options::ClientOptions};

// #[derive(Debug, Deserialize)]
// struct ApiResponse{
//     code: u16,
//     msg: String
// }

// impl<'r, 'o> Responder<'r, 'r> for Result<String, Error> {
//     fn respond_to(self, _request: &'r Request<'_>) -> response::Result<'r> {
//         Response::build().status(Status::new(self.code)).sized_body(self.msg.len(),std::io::Cursor::new(self.msg)).ok()
//     }
// }



#[get("/")]
async fn hello() -> Option<content::Json<String>>{
   let client = establish_connection().await.ok()?; 
   let code = check_connection(&client).await.ok()?;

   Some(content::Json(format!("{{
       'code': {} 
   }}", code.to_string())))
}


#[rocket::main]
async fn main() -> Result<(),rocket::Error> {
    rocket::build()
    .mount("/v1", routes![hello]).ignite().await?.launch().await
}

//establishes connection between mongodb database
async fn establish_connection() -> mongodb::error::Result<Client> {
    let mut client_options = ClientOptions::parse("mongodb://localhost:27017").await?;
    client_options.app_name = Some("unga bunga".to_string());
    Client::with_options(client_options)

}
//checks ping (doesn't verify whether a database exists or not so yea idc)
async fn check_connection(client: &Client) -> mongodb::error::Result<i32>{
    client.database("admin").run_command(doc! {"ping": 1}, None).await?;
    Ok(200)
}

