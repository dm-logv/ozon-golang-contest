create table goods (id INTEGER, name TEXT);
create table tags (id INTEGER, name TEXT);
create table tags_goods (tag_id INTEGER, goods_id INTEGER, UNIQUE (tag_id, goods_id));

insert into goods values (1, 'Good 1');
insert into goods values (2, 'Good 2');
insert into goods values (3, 'Good 3');
insert into goods values (4, 'Good 4');
insert into goods values (5, 'Good 5');
insert into goods values (6, 'Good 6');
insert into goods values (7, 'Good 7');
insert into goods values (8, 'Good 8');
insert into goods values (9, 'Good 9');

insert into tags values (1, 'Tag 1');
insert into tags values (2, 'Tag 2');
insert into tags values (3, 'Tag 3');
insert into tags values (4, 'Tag 4');
insert into tags values (5, 'Tag 5');

insert into tags_goods values (1, 2);
insert into tags_goods values (2, 2);
insert into tags_goods values (3, 2);
insert into tags_goods values (4, 2);
insert into tags_goods values (5, 2);
insert into tags_goods values (1, 1);
insert into tags_goods values (2, 1);
insert into tags_goods values (3, 1);
insert into tags_goods values (3, 4);
insert into tags_goods values (1, 8);
insert into tags_goods values (2, 8);
insert into tags_goods values (3, 8);
insert into tags_goods values (4, 8);
insert into tags_goods values (5, 8);
insert into tags_goods values (4, 7);
insert into tags_goods values (5, 7);
