CREATE TABLE items (
	timestamp TIMESTAMP NOT NULL,
	uid STRING(MAX) NOT NULL,
	number INT64,
	part STRING(MAX),
	text STRING(MAX) NOT NULL,
) PRIMARY KEY (uid, timestamp)
