FROM rustlang/rust:nightly-slim

WORKDIR /rust

COPY /rust /rust

RUN cargo build
CMD ["cargo","run"]
