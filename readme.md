<!-- [link](https://archive.ph/th5cL) -->
# Mouse Counter

A simple web application built with Go that counts the number of times your mouse enters a specific area on the page. Built using:

- [Echo](https://echo.labstack.com/) - Web framework
- [HTMX](https://htmx.org/) - Dynamic HTML updates
- [Templ](https://templ.guide/) - Go templating engine

## Features

- Real-time counter updates without page refresh
- Simple and lightweight interface
- Server-side state management

## Getting Started

### Prerequisites

- Go 1.23.4 or later
- Git

### Installation

1. Clone the repository:
```bash
git clone [repository-url]
cd mouse-counter
go run .
```
## How It Works
The application displays a counter that increments each time your mouse enters a designated area. It uses:

- HTMX for seamless client-side updates
- Echo framework for handling HTTP requests
- Templ for HTML templating

Server-side state management for counting
Project Structure

- main.go - Main application entry point and server setup
- views - Template files and view components
- assets - Static assets (if any)