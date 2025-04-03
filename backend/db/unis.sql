CREATE TABLE kazan_unis (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(256) NOT NULL DEFAULT ''
);

CREATE TABLE moscow_unis (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(256) NOT NULL DEFAULT ''
);

CREATE TABLE stp_unis (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(256) NOT NULL DEFAULT ''
);

INSERT INTO kazan_unis (name) 
VALUES 
('Казанский (Приволжский) федеральный университет'),
('Казанский национальный исследовательский технический университет им. А.Н. Туполева - КАИ'),
('Казанский государственный медицинский университет'),
('Казанский национальный исследовательский технологический университет'),
('Казанский государственный энергетический университет'),
('Казанский государственный архитектурно-строительный университет'),
('Казанский государственный аграрный университет'),
('Казанский инновационный университет имени В.Г. Тимирясова'),
('Казанский государственный институт культуры'),
('Казанский юридический институт МВД России');