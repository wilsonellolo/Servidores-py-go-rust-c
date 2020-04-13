#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use] extern crate rocket;

use std::path::{Path, PathBuf};

use rocket::response::NamedFile;
//rocket.mount("/public", StaticFiles::from("/static"))
//#[get("/")]
//fn index() -> &'static str {
//    "Hello, world!"
//}

#[get("/")]
pub fn index() -> Option<NamedFile> {
    NamedFile::open("static/index.html").ok()
}

#[get("/mundo")]
fn mundo() -> &'static str {
    "Hello, MUNDO EN ESPAñOL"
}


fn main() {
    rocket::ignite().mount("/", routes![index,mundo]).launch();
}