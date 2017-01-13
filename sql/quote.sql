DROP TABLE IF EXISTS quotes, quote_images;

-- TODO: Image, should figure out a uniform solution.
CREATE TABLE quote_images (
  ID SERIAL PRIMARY KEY,
  DATA BYTEA,
  URL VARCHAR(2083)
);

CREATE TABLE quotes(
  ID SERIAL PRIMARY KEY,
  -- User authentication from other platforms. Query index.
  USER_ID VARCHAR(255) NOT NULL,
  PLATFORM VARCHAR(32) NOT NULL,
  -- Data structure for quotes.
  CONTENT TEXT NOT NULL,
  AUTHOR TEXT NOT NULL,
  SOURCE TEXT,
  SECTION TEXT,
  IMAGE_ID INTEGER REFERENCES quote_images
);

-- Should support query by (user, platform).
CREATE INDEX quotes_idx ON quotes (USER_ID, PLATFORM);

-- Insert example data.
INSERT INTO quotes (USER_ID, PLATFORM, CONTENT, AUTHOR) VALUES (
	'edfward', 'console', 'Hello world.', 'Junjia He'
);
INSERT INTO quotes (USER_ID, PLATFORM, CONTENT, AUTHOR, SOURCE) VALUES (
	'edfward', 'console', 'Nice boots!', 'Jing Yu', 'SW 101'
);
