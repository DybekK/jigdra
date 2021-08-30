
CREATE DATABASE IF NOT EXISTS jigdra;

CREATE TABLE IF NOT EXISTS workspace_user (
	id			UUID PRIMARY KEY,
	user_id		TEXT,
	nickname	VARCHAR(30)
);

CREATE TABLE IF NOT EXISTS workspace (
	id					UUID PRIMARY KEY,
	workspace_user_id	UUID
);

CREATE TABLE IF NOT EXISTS task (
	id					UUID	PRIMARY KEY,
	workspace_id		UUID,
	workspace_user_id	UUID
);
