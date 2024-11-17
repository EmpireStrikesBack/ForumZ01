CREATE TABLE IF NOT EXISTS Clients (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    last_name TEXT,
    first_name TEXT,
    user_name TEXT UNIQUE,
    email TEXT UNIQUE,
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

CREATE TABLE IF NOT EXISTS Posts (
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
    FOREIGN KEY(author_id) REFERENCES Clients(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Comments (
    comment_id INT AUTO_INCREMENT PRIMARY KEY, 
    post_id INT NOT NULL,
    user_id INT NOT NULL,
    content TEXT,
    creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES Posts(post_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES Clients(user_id) ON DELETE CASCADE
);