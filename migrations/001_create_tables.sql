-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- users table
CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  username VARCHAR UNIQUE NOT NULL,
  email VARCHAR UNIQUE NOT NULL,
  password_hash VARCHAR NOT NULL,
  role VARCHAR NOT NULL,
  first_name VARCHAR NOT NULL,
  last_name VARCHAR NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- books table
CREATE TABLE books (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  google_book_id VARCHAR UNIQUE NOT NULL,
  isbn_10 VARCHAR,
  isbn_13 VARCHAR,
  title VARCHAR NOT NULL,
  publisher VARCHAR,
  published_date VARCHAR,
  description TEXT,
  page_count INT,
  language VARCHAR,
  thumbnail_url VARCHAR,
  total_copies INT NOT NULL,
  available_copies INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- authors table
CREATE TABLE authors (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR UNIQUE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- categories table
CREATE TABLE categories (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR UNIQUE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- book_authors junction table
CREATE TABLE book_authors (
  book_id UUID NOT NULL,
  author_id UUID NOT NULL,
  PRIMARY KEY (book_id, author_id),
  FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
  FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE
);

-- book_categories junction table
CREATE TABLE book_categories (
  book_id UUID NOT NULL,
  category_id UUID NOT NULL,
  PRIMARY KEY (book_id, category_id),
  FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
  FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

-- loans table
CREATE TABLE loans (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL,
  book_id UUID NOT NULL,
  borrowed_date DATE NOT NULL,
  due_date DATE NOT NULL,
  returned_date DATE,
  status VARCHAR NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);

-- book_reviews table
CREATE TABLE book_reviews (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  book_id UUID NOT NULL,
  user_id UUID NOT NULL,
  rating INT NOT NULL,
  review_text TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);

-- Add check constraints
ALTER TABLE book_reviews ADD CONSTRAINT rating_range CHECK (rating >= 1 AND rating <= 5);
ALTER TABLE users ADD CONSTRAINT valid_role CHECK (role IN ('admin', 'member'));
ALTER TABLE loans ADD CONSTRAINT valid_status CHECK (status IN ('active', 'returned', 'overdue'));

-- Create indexes for performance
CREATE INDEX idx_books_title ON books(title);
CREATE INDEX idx_loans_status ON loans(status);
CREATE INDEX idx_loans_user_id ON loans(user_id);
CREATE INDEX idx_loans_book_id ON loans(book_id);
CREATE INDEX idx_book_reviews_book_id ON book_reviews(book_id);
CREATE INDEX idx_book_reviews_user_id ON book_reviews(user_id);

-- +goose Down
DROP TABLE IF EXISTS book_reviews;
DROP TABLE IF EXISTS loans;
DROP TABLE IF EXISTS book_categories;
DROP TABLE IF EXISTS book_authors;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS authors;
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS users;
