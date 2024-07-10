# Gin CLI Ticket Booking App

# Overview

This is a Gin framework-based command-line interface (CLI) application for booking tickets to a conference. The application uses a MySQL database to store user data and ticket information.

# Features

- User registration and ticket booking
- Validation of user input (name, email, ticket number)
- Ticket booking with payment processing (simulated)
- Sending tickets to users via email (simulated)
- Display of remaining tickets and conference details

# Usage

1. Clone the repository and run go build to build the application.
2. Run go run main.go to start the application.
3. Use the command-line interface to book tickets.

# Commands

- book: Book tickets for the conference.
- welcome: Display a welcome message.
- static: Serve static files from the static directory.

# Configuration

- Database connection settings are stored in the mysql package.
- Conference details (name, comedian, tickets) are stored in the main package.

# Dependencies

- GO 1.22.2 or later
- Gin framework latest
- MySQL driver
- Go-SQL driver
