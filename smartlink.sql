create database smartlink;
use smartlink;

create  table users (
	id int(11) auto_increment primary key,
	user_id varchar(15),
	nama varchar(50),
	username varchar(15),
	password varchar(50),
	telepon varchar(15)
);

create  table layanans (
	id int(11) auto_increment primary key,
	layanan_id varchar(15),
	nama varchar(50),
	unit varchar (3),
	harga double(10,2),
	user_id int(11),
	user_id_s varchar (15),
	foreign key (user_id) references users(id)
);

