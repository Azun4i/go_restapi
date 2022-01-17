CREATE TABLE users(
    uuid      varchar not null,
    firstname varchar   not null,
    lastname  varchar   not null,
    email     varchar   not null unique,
    age       integer,
    created   timestamp without time zone

);