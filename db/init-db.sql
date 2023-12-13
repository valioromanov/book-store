CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title TEXT,
    author TEXT,
    book_cover LONGBLOB,
    published_at TIMESTAMP
);

insert into books.books (id, title, author, book_cover, published_at)
values (1,'1984','George Orwel', '',STR_TO_DATE('2022-12-01', '%Y-%m-%d'));