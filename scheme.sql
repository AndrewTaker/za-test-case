DROP TABLE IF EXISTS NEWS, NEWS_CATEGORIES, USERS;

CREATE TABLE NEWS (
  ID BIGSERIAL PRIMARY KEY,
  TITLE TEXT NOT NULL,
  CONTENT TEXT NOT NULL
);

CREATE TABLE NEWS_CATEGORIES (
  ID SERIAL NOT NULL,
  NEWS_ID BIGINT NOT NULL,
  PRIMARY KEY (NEWS_ID, ID),
  FOREIGN KEY (NEWS_ID) REFERENCES NEWS(ID) ON DELETE CASCADE
);

CREATE TABLE USERS (
  ID SERIAL PRIMARY KEY,
  API_KEY TEXT NULL UNIQUE
);

INSERT INTO USERS (API_KEY) VALUES
('vyISUXX08FPi1gnTdnaV0UeVBWpx9yN3'),
('vAZfft1mPx8QUVpRFUBKntb5rOBqIjry'),
('8gmq1PcOCubrPFVoLgBxjqgk8ofNLP3p'),
('sgNxqEGdj9Uws7sQj9b22ArZiVvSTnyL'),
('sxdZfdyYn7o2V5xr0GYV1CzeO8vdHM1A'),
('19z8oAoLUSRgCPURoCQeL5q2PhJ7EPiu'),
('5EiFJtQpXarwy0hK7SdlTAkUbO59wmDg'),
('VB90ZyMRgkqAfYmYardU9ldehZRrqDnP'),
('5mWMZSsc5ugaIpPk2xpU3giYdL8PEaFq'),
('4ysDI1hL33GpgW1WvGgLII2tDfyi234W');

INSERT INTO NEWS (TITLE, CONTENT) VALUES 
('Title 1', 'Content for news 1'),
('Title 2', 'Content for news 2'),
('Title 3', 'Content for news 3'),
('Title 4', 'Content for news 4'),
('Title 5', 'Content for news 5'),
('Title 6', 'Content for news 6'),
('Title 7', 'Content for news 7'),
('Title 8', 'Content for news 8'),
('Title 9', 'Content for news 9'),
('Title 10', 'Content for news 10'),
('Title 11', 'Content for news 11'),
('Title 12', 'Content for news 12'),
('Title 13', 'Content for news 13'),
('Title 14', 'Content for news 14'),
('Title 15', 'Content for news 15'),
('Title 16', 'Content for news 16'),
('Title 17', 'Content for news 17'),
('Title 18', 'Content for news 18'),
('Title 19', 'Content for news 19'),
('Title 20', 'Content for news 20');

INSERT INTO NEWS_CATEGORIES (NEWS_ID, ID) VALUES 
(1, 1), (1, 2), (2, 1), (2, 3), (3, 1), (3, 4),
(4, 1), (4, 5), (5, 1), (5, 6), (6, 2), (6, 7),
(7, 2), (7, 8), (8, 2), (8, 9), (9, 2), (9, 10),
(10, 3), (10, 4), (11, 3), (11, 5), (12, 3), (12, 6),
(13, 3), (13, 7), (14, 4), (14, 8), (15, 4), (15, 9),
(16, 4), (16, 10), (17, 5), (17, 6), (18, 5), (18, 7),
(19, 5), (19, 8), (20, 5), (20, 9);