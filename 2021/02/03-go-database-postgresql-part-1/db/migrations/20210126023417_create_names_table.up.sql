CREATE TABLE names (
	nconst              varchar(255),
	primary_name        varchar(255),
	birth_year          varchar(4),
	death_year          varchar(4) DEFAULT '',
	primary_professions varchar[],
	known_for_titles    varchar[]
);
