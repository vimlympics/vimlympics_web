CREATE TABLE events (
    event_id INTEGER PRIMARY KEY,
    event_type INTEGER NOT NULL,
    event_level INTEGER NOT NULL
);

CREATE TABLE scores (
    score_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    event_id INTEGER NOT NULL,
    timems INTEGER NOT NULL,
    date_entered DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (event_id) REFERENCES events(event_id)
);

CREATE TABLE users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(150) NOT NULL,
    api_key VARCHAR(150) NOT NULL,
    country VARCHAR(2)
);
