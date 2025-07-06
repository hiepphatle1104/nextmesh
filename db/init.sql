create table accounts (
  id varchar(36) primary key,
  name varchar(255) not null,
  email varchar(255) not null unique,
  password varchar(255) not null,
  created_at timestamp default current_timestamp
);

create table sessions (
  id varchar(255) primary key,
  account_id varchar(36) not null,
  expires_at timestamp not null,
  created_at timestamp default current_timestamp,
  foreign key (account_id) references accounts(id)
);
