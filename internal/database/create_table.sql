CREATE TABLE IF NOT EXISTS Clients (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    last_name TEXT,
    first_name TEXT,
    user_name TEXT UNIQUE,
    email TEXT UNIQUE,
    oauth_provider TEXT,
    oautH_id TEXT,
    password TEXT,
    avatar TEXT,
    birth_date DATE,
    user_role TEXT DEFAULT 'user',
    creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deletion_date TIMESTAMP 
);

CREATE TABLE IF NOT EXISTS Sessions (
	session_id TEXT PRIMARY KEY,
	user_id INTEGER,
	user_role TEXT,
	user_name TEXT,
	is_logged_in BOOLEAN DEFAULT FALSE,
	expiration TIMESTAMP,
	creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deletion_date TIMESTAMP,
	is_deleted BOOLEAN DEFAULT FALSE,
	FOREIGN KEY(user_id) REFERENCES Clients(user_id)
);

CREATE TABLE IF NOT EXISTS posts (
	post_id INTEGER PRIMARY KEY AUTOINCREMENT,
	author_id INTEGER,
	title TEXT,
	category TEXT,
	tags TEXT,
	content TEXT,
	creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deletion_date TIMESTAMP,
	is_deleted BOOLEAN DEFAULT FALSE,
	likes INTEGER DEFAULT 0,
	dislikes INTEGER DEFAULT 0,
	FOREIGN KEY(author_id) REFERENCES Clients(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Comments (
	comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
	post_id INT NOT NULL,
	user_id INT NOT NULL,
	content TEXT,
	creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES Clients(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS likes_dislikes (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	post_id INT NOT NULL,
	user_id INT NOT NULL,
	liked BOOLEAN DEFAULT NULL,
	update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES Clients(user_id) ON DELETE CASCADE
	UNIQUE (post_id, user_id)
);
