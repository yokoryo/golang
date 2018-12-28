use go_apps;
create table tasks(id int auto_increment not null primary key, username varchar(255), task text, end_flag varchar(5), created_at timestamp not null default current_timestamp, updated_at timestamp not null default current_timestamp on update current_timestamp, index(id));
