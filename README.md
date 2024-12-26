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
    CREATE DATABASE inventory OWNER postgres_user;
    \q


## 🚀 How to Set Up the Project

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/inventory-backend.git
   cd inventory-backend


2. Install all the dependecies

    go mod tidy

3. Execute the server 
    go run main.go

4. Set up the database:
Create a local PostgreSQL database named inventory_db.
Update the connection settings in the config/database.go file:
    const (
    Host     = "localhost"
    Port     = 5432
    User     = "your_username"
    Password = "your_password"
    DBName   = "inventory_db"
)

# inventory-backend-project
