create table if not exists user (
    id bigint not null auto_increment,
    user_id bigint not null,
    username varchar(64) collate utf8mb4_general_ci not null,
    password varchar(64) collate utf8mb4_general_ci not null,
    email varchar(64),
    gender tinyint not null default '0',
    create_time timestamp null default current_timestamp,
    update_time timestamp null default current_timestamp on update current_timestamp,
    primary key(id),
    unique key idx_username (username) using btree,
    unique key idx_user_id (user_id) using btree
);

drop table if exists `community`;
create table `community` (
    id int not null auto_increment,
    community_id int not null,
    community_name varchar(128) not null,
    introduction varchar(256) not null,
    create_time timestamp not null default current_timestamp,
    update_time timestamp not null default current_timestamp on update current_timestamp,
    primary key(id),
    unique key `idx_community_id` (community_id),
    unique key `idx_community_name` (community_name)
);

insert into community (community_id,community_name,introduction) 
values (1,"go", "go社区"),(2,"python", "python社区"),(3,"rust","rust社区");

drop table if exists post;
create table post(
    id bigint not null auto_increment,
    post_id bigint not null comment '帖子id',
    title varchar(128) not null comment '标题',
    content varchar(8192) not null comment '内容',
    author_id bigint not null comment '作者的用户id',
    community_id bigint not null comment '所属社区',
    status tinyint not null default '1' comment '帖子状态',
    create_time timestamp null default current_timestamp comment '创建时间',
    update_time timestamp null default current_timestamp on update current_timestamp comment '更新时间',
    primary key(id),
    unique key idx_post_id (post_id),
    key idx_author_id (author_id),
    key idx_community_id (community_id)
);