#[macro_use] extern crate rocket;

extern crate mongodb;
extern crate serde;
extern crate rocket_contrib;
use std::panic::catch_unwind;

use rocket::
    {Request, Response, data::FromData, form::Form, http::{ContentType, Status}, request::FromRequest, response::{self, Responder}, serde::json::serde_json::json};
use mongodb::{Client, bson::{doc, oid::ObjectId}, options::ClientOptions};
use serde::{Serialize,Deserialize};
use rocket_contrib::json::Json;
#[derive(Debug)]
struct ApiResponse{
    code: Status,
    content_type: ContentType,
    text: String
}

impl<'r, 'o> Responder<'r,'r> for ApiResponse{
    fn respond_to(self, request: &'r Request<'_>) -> response::Result<'r> {
        Response::build_from(self.text.respond_to(request)?)
        .status(self.code)
        .header(self.content_type).ok()
    }
}
// enum ApiError{
//     Unauthorized(String),
//     NotFound(rocket::http::hyper::Error)
// }

// impl<'r> Responder<'r,'r> for ApiError{
//     fn respond_to(self, request: &'r Request<'_>) -> std::result::Result<Response<'r>, Status> {
//         match self {
//             ApiError::Unauthorized(_) => 
//             Response::build()
//                 .status(Status::Unauthorized)
//                 .header(ContentType::new("application","json")).ok(),
//             ApiError::NotFound(_) => 
//             Response::build()
//             .status(Status::NotFound)
//             .header(ContentType::new("application","json")).ok(),
//         }
//     }
// }

// impl From<Error> for ApiError {
//     fn from(_: Error) -> Self {
//         ApiError::NotFound(Error)
//     }
// }

#[derive(Serialize, Deserialize)]
struct User{
    id: ObjectId,
    username: String,
    email: String,
    pw_hash: String
}

#[derive(Serialize,Deserialize, FromForm)]
struct FormUser{
    username: String,
    password: String
}

// #[rocket::async_trait]
// impl<'r> FromData<'r> for FormUser{
//     type Error = ();

    
//     async fn from_data(req: &'r Request<'_>, data: rocket::Data<'r>) -> rocket::data::Outcome<'r, Self> {
//         todo!()
//     }
// }

#[catch(404)]
fn not_found(req: &Request) -> ApiResponse{
    ApiResponse{
        code: Status::NotFound,
        content_type: ContentType::new("application", "json"),
        text: json!({
            "code": 404,
            "msg": "Page not found"
        }).to_string()
    }
}

#[catch(500)]
fn server_error(req: &Request) -> ApiResponse{
    ApiResponse{
        code: Status::InternalServerError,
        content_type: ContentType::new("application", "json"),
        text: json!({
            "code": 500,
            "msg": "Internal server error"
        }).to_string()
    }
}


#[get("/")]
async fn hello() -> Option<ApiResponse>{
   let client = establish_connection().await.ok()?;
   let code = check_connection(&client).await.ok()?;

   Some(ApiResponse{
     code: Status::new(code as u16),
     content_type: ContentType::new("application", "json"),
     text: json!({
         "code": code,
         "msg": "success"
     }).to_string()
    })
}


#[rocket::main]
async fn main() -> Result<(),rocket::Error> {
    rocket::build()
    .register("/", catchers![server_error,not_found])
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

