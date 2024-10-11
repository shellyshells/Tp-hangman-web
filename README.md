# promo

This project is a simple web application demonstrating dynamic content handling, conditional logic, and user input validation using Golang.

## Prerequisites

- Git
- Go (version 1.16 or later)
- A text editor or IDE (e.g., Visual Studio Code, GoLang, or Sublime Text)

## Setup and Installation

1. Clone the repository:
   ```
   git clone https://github.com/shellyshells/promo.git
   cd promo
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

## Running the Application

1. Start the server:
   ```
   go run main.go
   ```

2. You should see the message: "Server is running on http://localhost:8080"

## Testing the Application

Open a web browser and visit the following URLs to test different routes:

1. Promo page: http://localhost:8080/promo
2. Change page: http://localhost:8080/change
3. User form: http://localhost:8080/user/form

For the user form, fill out the information and submit. You will be redirected to a page displaying the submitted information.

## Troubleshooting

- Ensure that port 8080 is not being used by another application.
- If you encounter any errors, make sure you have the latest version of Go installed.
