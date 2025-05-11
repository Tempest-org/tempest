CREATE TABLE IF NOT EXISTS accounts (
  id varchar(255) NOT NULL,

  email varchar(255) NOT NULL,
  username varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  phone varchar(255) DEFAULT NULL,

  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp DEFAULT NULL,
  PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS accounts_email_idx ON accounts (email);
CREATE UNIQUE INDEX IF NOT EXISTS accounts_username_idx ON accounts (username);
CREATE UNIQUE INDEX IF NOT EXISTS accounts_phone_idx ON accounts (phone) WHERE phone IS NOT NULL;
