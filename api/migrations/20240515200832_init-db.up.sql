create table if not exists games (
	id text primary key,
	created_at date null,
	owner_id text not null,

	foreign key (owner_id) references users(id)
);

create table if not exists users (
	id text primary key,
	name text not null,
	created_at date null,
	game_id text null,

	foreign key (game_id) references games(id) on delete cascade
);