#[macro_use] extern crate rocket;

#[get("/")]
fn hello() -> &'static str{
    "Hewwo"
}

#[launch]
fn rocket() -> _ {
    rocket::build()
    .mount("/v1", routes![hello])
}