CREATE TABLE userlist
(
    id serial not null unique,
    username varchar(255) not null unique,
    password varchar(255) not null
);


CREATE TABLE sellinglist
(
    id serial not null unique,
    title varchar(150) not null,
    price int not null,
    date date not null,
    description varchar(1000) not null,
    url varchar(100) not null
);

CREATE TABLE usersellingtable
(
    id serial not null unique,
    user_login varchar(255) references userlist(username) on delete cascade not null,
    user_id int references userlist(id) on delete cascade not null,
    list_id int references sellinglist(id) on delete cascade not null
);