CREATE TABLE IF NOT EXISTS customers (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  cognito_id TEXT NOT NULL,
  created_at DATETIME DEFAULT current_timestamp,
  first_name TEXT NOT NULL,
  middle_name TEXT,
  last_name TEXT NOT NULL,
  date_of_birth DATE NOT NULL,
  email_address TEXT NOT NULL,
  phone_number TEXT
);

CREATE TABLE IF NOT EXISTS addresses (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  customer_id REFERENCES customers(id),
  line_1 TEXT NOT NULL,
  line_2 TEXT,
  city TEXT NOT NULL,
  country TEXT NOT NULL,
  postcode TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS organisations (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  address_id REFERENCES addresses(id)
);

CREATE TABLE IF NOT EXISTS hotels (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  organisation_id REFERENCES organisations(id),
  address_id REFERENCES addresses(id),
  star_rating SMALLINT NOT NULL
);

CREATE TABLE IF NOT EXISTS rooms (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  room_number TEXT NOT NULL,
  hotel_id REFERENCES hotels(id),
  room_features CLOB NOT NULL,
  capacity SMALLINT NOT NULL,
  base_nightly_price DECIMAL(10,2) NOT NULL
);

CREATE TABLE IF NOT EXISTS payment_methods (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  customer_id REFERENCES customers(id),
  stripe_payment_id TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS daily_rate_increases (
  hotel_id REFERENCES hotels(id) NOT NULL,
  day_of_adjustment DATE NOT NULL,
  rate_multiplier DECIMAL(10,2) NOT NULL,
  PRIMARY KEY(hotel_id, day_of_adjustment)
);

CREATE TABLE IF NOT EXISTS payments (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_at DATETIME NOT NULL,
  description TEXT NOT NULL,
  stripe_payment_id TEXT NOT NULL,
  payment_method_id REFERENCES payment_methods(id) NOT NULL
);

CREATE TABLE IF NOT EXISTS bookings (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  check_in_date DATE NOT NULL,
  check_out_date DATE NOT NULL,
  room_id REFERENCES rooms(id) NOT NULL,
  deposit_id REFERENCES payments(id) NOT NULL,
  payment_id REFERENCES payments(id) NOT NULL,
  total_price DECIMAL(10, 2) NOT NULL,
  customer_id REFERENCES customers(id) NOT NULL
);

CREATE TABLE IF NOT EXISTS reviews (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  booking_id REFERENCES bookings(id) NOT NULL,
  rating INTEGER NOT NULL,
  title TEXT,
  body TEXT,
  created_at DATETIME DEFAULT current_timestamp
);