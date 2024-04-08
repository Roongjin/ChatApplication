CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
  username varchar(255) UNIQUE NOT NULL,
  is_online boolean NOT NULL DEFAULT FALSE
);


CREATE TABLE rooms (id UUID PRIMARY KEY DEFAULT gen_random_uuid ());


CREATE TABLE user_room_links (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
  user_id UUID REFERENCES users (id) ON UPDATE CASCADE ON DELETE SET NULL,
  room_id UUID REFERENCES rooms (id) ON UPDATE CASCADE ON DELETE SET NULL,
);


CREATE TABLE conversations (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
  text varchar(2000) NOT NULL,
  user_id UUID REFERENCES users (id) ON UPDATE CASCADE ON DELETE SET NULL,
  room_id UUID REFERENCES rooms (id) ON UPDATE CASCADE ON DELETE SET NULL,
);
