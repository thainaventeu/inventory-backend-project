# Inventory Backend - Project setup 

This is a backend project developed with **Go**, using **Gin Framework** and **PostgreSQL** to manage an inventory system.

## 📋 Prerequisites

Make sure you have the following installed on your system:
- [Go](https://go.dev/doc/install) (version 1.18 or later)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/downloads)

    ```bash   
    install postgresql@14 
    psql postgres
    CREATE ROLE postgres_user WITH LOGIN PASSWORD 'your_password'
    ALTER ROLE postgres CREATEDB;

    \q


## 🚀 How to Set Up the Project

1. **Clone the repository:**
   ```bash
   git clone https://github.com/thainaventeu/inventory-backend-project.git
   cd inventory-backend


2. **Install all the dependecies**
   ```bash
   go mod tidy

3. **Execute the server in background** 
    ```bash
    go run main.go &

4. **Set up the database:**
Create a local PostgreSQL database named inventory.
Update the connection settings in the config/database.go file:

    ```bash
    const (
    Host     = "localhost"
    Port     = 5432
    User     = "your_username"
    Password = "your_password"
    DBName   = "inventory"
)

## ✅ How To Validate This Project

1. **Create an item:**
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"name":"Monitor","stock":3,"price":300}' http://localhost:8080/inventory

2. **Update an item:**
    ```bash
    curl -X PUT -H "Content-Type: application/json" -d '{"name":"Updated Laptop","stock":10,"price":2000}' http://localhost:8080/inventory/1

3. **Delete an item:**
    ```bash
    curl -X DELETE http://localhost:8080/inventory/1

4. **Find a specific item:**
    ```bash
    curl -X GET http://localhost:8080/inventory/1

