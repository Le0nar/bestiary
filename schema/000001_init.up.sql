CREATE TABLE npc
(
    id serial not null unique,
    name varchar(255) not null,
    description varchar(255) not null,
    hp int not null
);

CREATE TABLE attack_type
(
    id serial not null unique,
    name varchar(255) not null,
    range int not null
);

CREATE TABLE enemy
(
    id serial not null unique,
    name varchar(255) not null,
    description varchar(255) not null,
    hp int not null,
    attack_type int references attack_type (id) on delete cascade not null,
    damage int not null,
    haste int not null
);
