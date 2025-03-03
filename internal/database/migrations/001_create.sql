CREATE TABLE streams (
    id INTEGER PRIMARY KEY,
    url TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    thumbnail TEXT
);

CREATE TABLE watch_parties (
    id INTEGER PRIMARY KEY,
    stream_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    password TEXT,
    FOREIGN KEY (stream_id) REFERENCES streams(id)
);

CREATE TABLE party_members (
    id INTEGER PRIMARY KEY,
    party_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (party_id) REFERENCES watch_parties(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE chat_messages (
    id INTEGER PRIMARY KEY,
    party_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    message TEXT NOT NULL,
    FOREIGN KEY (party_id) REFERENCES watch_parties(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);
