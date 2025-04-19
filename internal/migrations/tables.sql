CREATE TABLE students (
    id uuid PRIMARY KEY,
    telegram_id BIGINT,
    first_name VARCHAR(50),
    middle_name VARCHAR(50),
    last_name VARCHAR(50),
    github VARCHAR(100),
    job VARCHAR(100),
    idea TEXT,
    about TEXT
);

CREATE TABLE teachers (
    id uuid PRIMARY KEY,
    telegram_id BIGINT,
    first_name VARCHAR(50),
    middle_name VARCHAR(50),
    last_name VARCHAR(50),
    degree VARCHAR(50),
    position VARCHAR(50),
    department VARCHAR(50),
    is_free BOOLEAN,
    idea TEXT,
    about TEXT
);


INSERT INTO teachers 
(id, telegram_id, first_name, middle_name, last_name, degree, position, department, is_free, idea, about)
VALUES
  (gen_random_uuid(), 123456789, 'Иван', 'Иванович', 'Иванов', 'PhD', 'Профессор', 'Кафедра информатики', TRUE, 'Разработка квантовых алгоритмов', 'Специалист в области квантовых вычислений с 10-летним опытом'),
  ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 987654321, 'Мария', 'Петровна', 'Сидорова', 'Кандидат наук', 'Доцент', 'Кафедра математики', FALSE, 'Применение ИИ в математическом моделировании', 'Преподает высшую математику и занимается исследованиями в области ИИ'),
  (gen_random_uuid(), 555555555, 'Алексей', NULL, 'Петров', 'Магистр', 'Ассистент', 'Кафедра физики', TRUE, 'Исследование новых материалов', 'Молодой ученый, увлеченный нанотехнологиями'),
  ('b3d7a2e1-0c4f-4a8d-b6e2-f8c7d3e4a5b6', 111223344, 'Елена', 'Владимировна', 'Кузнецова', 'Доктор наук', 'Заведующая кафедрой', 'Кафедра биологии', FALSE, 'Генетические модификации растений', 'Автор более 50 научных работ по генетике'),
  (gen_random_uuid(), 999888777, 'Дмитрий', 'Сергеевич', 'Смирнов', NULL, 'Инженер', 'Научно-исследовательский отдел', TRUE, 'Разработка робототехнических систем', 'Инженер-робототехник с практическим опытом в промышленной автоматизации');