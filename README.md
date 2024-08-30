Installation
Follow these steps to set up and run the project on your local machine.

1. Clone the Repository
First, clone the repository to your local machine using Git:

bash
Copy code
git clone https://github.com/adimgozali13/01-xyz-finance.git

2. Add Environment Variables
Create a .env file in the root directory of the project. Add the following environment variables to configure your database connection:

makefile
Copy code
DB_USER=your_username
DB_PASSWORD=your_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=your_db
Replace your_username, your_password, and your_db with your actual database credentials.

3. Run the Application
To start the application, run the following command:

bash
Copy code
go run main.go
The application will start running on the default port specified in your code.