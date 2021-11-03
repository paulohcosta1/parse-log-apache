CREATE TABLE IF NOT EXISTS logs(
   id serial PRIMARY KEY,
   ip VARCHAR (50),
   date DATE,
   method VARCHAR (50),
   resource VARCHAR (255),
   version VARCHAR (50),
   status VARCHAR (50),
   size VARCHAR (255),
   referer VARCHAR (255),
   user_agent VARCHAR (255)
);

