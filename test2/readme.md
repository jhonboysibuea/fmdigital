# Project Name: Soal Test Tahap 2

This project provides a wallet service REST API for user registration, login, top-up, payment, transfer, transaction history, and profile update functionality. It is built using **Java Spring Boot** and connects to a **MySQL** database.

## Prerequisites

Before running this project, ensure you have the following installed:
- **Java 11** or later
- **Maven 3.6+**
- **MySQL** (configured with database name, username, and password)
- **Postman** or **Curl** (for API testing)

## Setup Instructions

### 1. Clone the repository
```bash
git clone 
cd test2


## Database Configuration

# MySQL Database Configuration
spring.datasource.url=jdbc:mysql://localhost:3306/your_database_name?useSSL=false&serverTimezone=UTC
spring.datasource.username=your_mysql_username
spring.datasource.password=your_mysql_password

# Hibernate Configuration
spring.jpa.hibernate.ddl-auto=update
spring.jpa.show-sql=true


 ## Run Project
 mvn clean install
 mvn spring-boot:run




