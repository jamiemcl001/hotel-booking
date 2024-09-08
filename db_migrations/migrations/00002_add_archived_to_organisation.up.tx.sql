ALTER TABLE organisations ADD archived BOOLEAN DEFAULT false;

ALTER TABLE organisations ADD archived_at DATETIME DEFAULT current_timestamp;