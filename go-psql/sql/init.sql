
CREATE DATABASE jigdra;

CREATE TABLE IF NOT EXISTS workspace
(
    id uuid,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS workspace_user
(
    id uuid,
    nickname varchar,
    user_id uuid NOT NULL,
    workspace_id uuid NOT NULL,
    FOREIGN KEY (workspace_id) REFERENCES workspace(id) ON DELETE CASCADE,
    PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS task
(
    id uuid,
    workspace_user_id uuid NOT NULL,
    workspace_id uuid NOT NULL,
    title varchar NOT NULL,
    FOREIGN KEY (workspace_user_id) REFERENCES workspace_user(id),
    FOREIGN KEY (workspace_id) REFERENCES workspace(id),
    PRIMARY KEY (id)
);
