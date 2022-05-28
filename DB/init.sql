USE MedicalQuestionnaireService;
DROP TABLE IF EXISTS answers;
DROP TABLE IF EXISTS questions;
DROP TABLE IF EXISTS question_forms;
DROP TABLE IF EXISTS doctors;
DROP TABLE IF EXISTS clients;
DROP TABLE IF EXISTS master_questions;

CREATE TABLE doctors (
id INT UNSIGNED NOT NULL AUTO_INCREMENT,
first_name VARCHAR(100) NOT NULL,
last_name VARCHAR(100) NOT NULL,
mail VARCHAR(100) UNIQUE NOT NULL,
phone_number VARCHAR(100) NOT NULL,
password VARCHAR(100) NOT NULL,
sex VARCHAR(10) NOT NULL,
birthday DATETIME NOT NULL,
affiliation VARCHAR(100) NOT NULL,
speciality VARCHAR(100) NOT NULL,
refresh_token VARCHAR(100) NOT NULL,
access_token VARCHAR(100) NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
deleted_at DATETIME,
PRIMARY KEY (id)
);

ALTER TABLE `doctors` auto_increment = 1;

CREATE TABLE clients  (
id INT UNSIGNED NOT NULL AUTO_INCREMENT,
first_name VARCHAR(100) NOT NULL,
last_name VARCHAR(100) NOT NULL,
mail VARCHAR(100) UNIQUE NOT NULL,
phone_number VARCHAR(100) NOT NULL,
password VARCHAR(100) NOT NULL,
sex VARCHAR(100) NOT NULL,
birthday DATETIME NOT NULL,
refresh_token VARCHAR(100) NOT NULL,
access_token VARCHAR(100) NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
deleted_at DATETIME,
PRIMARY KEY (id)
);

ALTER TABLE `clients` auto_increment = 1;

CREATE TABLE questions (
id INT UNSIGNED NOT NULL AUTO_INCREMENT,
question_text VARCHAR(100) NOT NULL,
selection TINYINT UNSIGNED NOT NULL,
statement VARCHAR(100) NOT NULL,
doctor_id INT DEFAULT 0 NOT NULL,
multiple_choice VARCHAR(100) DEFAULT NULL,
single_choice VARCHAR(100) DEFAULT NULL,
paragraph VARCHAR(1000) DEFAULT NULL,
jump TINYINT UNSIGNED DEFAULT NULL,
ng TINYINT UNSIGNED DEFAULT NULL,
department VARCHAR(100) NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
deleted_at DATETIME,
PRIMARY KEY(id)
);

ALTER TABLE `questions` auto_increment = 1;

CREATE TABLE master_questions (
id INT UNSIGNED NOT NULL AUTO_INCREMENT,
master_question_text VARCHAR(100) NOT NULL,
selection TINYINT UNSIGNED NOT NULL,
statement VARCHAR(100) NOT NULL,
multiple_choice VARCHAR(100) DEFAULT NULL,
single_choice VARCHAR(100) DEFAULT NULL,
paragraph VARCHAR(1000) DEFAULT NULL,
jump TINYINT UNSIGNED DEFAULT NULL,
ng TINYINT UNSIGNED DEFAULT NULL,
department VARCHAR(100) NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
deleted_at DATETIME,
PRIMARY KEY(id)
);

ALTER TABLE `master_questions` auto_increment = 1;

CREATE TABLE question_forms  (
id INT UNSIGNED NOT NULL AUTO_INCREMENT,
doctor_id INT UNSIGNED NOT NULL,
question_id_list VARCHAR(100) NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
deleted_at DATETIME,
PRIMARY KEY(id),
FOREIGN KEY (doctor_id) REFERENCES doctors(id)
);

ALTER TABLE `question_forms` auto_increment = 1;

CREATE TABLE answers (
id INT UNSIGNED NOT NULL AUTO_INCREMENT,
doctor_id INT UNSIGNED NOT NULL,
client_id INT UNSIGNED NOT NULL,
paragraph VARCHAR(255),
statement VARCHAR(100),
selection  VARCHAR(100) DEFAULT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
deleted_at DATETIME,
PRIMARY KEY (id),
FOREIGN KEY (doctor_id) REFERENCES question_forms(doctor_id),
FOREIGN KEY (client_id) REFERENCES clients(id)
);

ALTER TABLE `answers` auto_increment = 1;
