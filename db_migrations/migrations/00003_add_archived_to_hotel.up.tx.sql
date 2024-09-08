ALTER TABLE hotels ADD archived BOOLEAN DEFAULT false;

ALTER TABLE hotels ADD archived_at DATETIME DEFAULT current_timestamp;