CREATE TABLE widgets (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL DEFAULT '',
  description TEXT NOT NULL DEFAULT '',
  inventory_level INTEGER NOT NULL,
  price INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  image VARCHAR(255) NOT NULL DEFAULT ''
);

CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  widget_id INTEGER NOT NULL,
  transaction_id INTEGER NOT NULL,
  status_id INTEGER NOT NULL,
  quantity INTEGER NOT NULL,
  amount INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT orders_widgets_id_fk FOREIGN KEY (widget_id) REFERENCES widgets (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT orders_transactions_id_fk FOREIGN KEY (transaction_id) REFERENCES transactions (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT orders_statuses_id_fk FOREIGN KEY (status_id) REFERENCES statuses (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(60) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  amount INTEGER NOT NULL,
  currency VARCHAR(255) NOT NULL,
  last_four VARCHAR(255) NOT NULL,
  bank_return_code VARCHAR(255) NOT NULL,
  transaction_status_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT transactions_transaction_statuses_id_fk FOREIGN KEY (transaction_status_id) REFERENCES transaction_statuses (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE transaction_statuses (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE statuses (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);