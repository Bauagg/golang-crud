CREATE TABLE products(
  	id INT AUTO_INCREMENT PRIMARY KEY,
	name_product VARCHAR(255) NOT NULL,  
  	stock INT NOT NULL, 
  	price INT NOT NULL,
  	images VARCHAR(255),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- migrate -database "mysql://root:root@tcp(127.0.0.1:3306)/databases-golang?query" -path databases/migrations up

-- migrate create -ext sql -dir databases/migrations -seq create_users_table