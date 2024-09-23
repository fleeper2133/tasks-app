CREATE TABLE users
(
    id serial not null unique,
    email varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE tasks
(
    id serial not null unique,
    title varchar(255) not null,
    description text,
    user_id int references users(id) on delete CASCADE not null,
    is_finish boolean default false,
    date_time_created timestamp default current_timestamp,
    date_time timestamp not null
);