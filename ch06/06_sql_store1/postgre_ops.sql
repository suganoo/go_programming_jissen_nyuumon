create role gwp with createdb login password='gwp';

select * from pg_stat_user_tables;

select * from pg_database;

insert into posts (content, author) values ('tgggggg', 'hogdfase');

alter sequence posts_id_seq restart with 1;
