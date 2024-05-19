# MySQL Installtion Steps
```brew install mysql
brew services start mysql 
mysql -u root -p admin123
create database bird_encyclopedia;
show databases;
use bird_encyclopedia;
CREATE TABLE birds (   id INT AUTO_INCREMENT PRIMARY KEY,   bird VARCHAR(256),   description VARCHAR(1024) );
show tables;

INSERT INTO birds (bird, description) VALUES 
('pigeon', 'common in cities'),
('eagle', 'bird of prey');

select * from birds;
```

# Refernce 
// Ref : https://www.sohamkamani.com/golang/sql-database/