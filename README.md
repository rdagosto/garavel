# Garavel - A Lightweight Go Framework Inspired by Laravel

**Garavel** is a lightweight yet powerful Go framework inspired by Laravel, designed to make web development simpler, faster, and more enjoyable. The goal of this project is to provide a minimalistic structure while still offering many of the features developers love about Laravel, such as routing, middleware, database integration, and easy configuration management.

**Status**: _In Development_  
**Current Version**: 0.1.0

## Features (Planned)

- **Routing**: Simple and intuitive routing with RESTful support.
- **Middleware**: Custom middleware for request and response processing.
- **Database ORM**: Integration with GORM for database interactions.
- **Request Validation**: Easy-to-use validation for incoming requests.
- **Environment Management**: Use of `.env` files for configuration.
- **Dependency Injection**: Built-in support for managing application dependencies.
- **Controllers & Views**: Structured approach for handling HTTP requests and responses.

## Installation

Clone this repository and install dependencies:

```bash
git clone git@github.com:rdagosto/garavel.git
cd garavel
mv .env.example .env
go run main.go
```

## Getting Started

1. Start the server:

```bash
go run main.go
```

Now, your server should be running on `http://localhost:8080/`. Visit `/hello` to see the response.

## Project Structure

Garavel follows a structured approach similar to Laravel:

```
/garavel
    /cmd
    /internal
        /configs
        /controllers
        /libs
        /middlewares
        /models
        /routes
        /validators
        /views
    /tests
    main.go
```

- **controllers/**: Handles HTTP request logic.
- **middleware/**: Stores middleware functions for request handling.
- **models/**: Defines database models using GORM.
- **views/**: Contains HTML templates (planned feature).
- **routes/**: Manages application routes.

## Planned Features

- **Eloquent-Like ORM**: A simple and intuitive ORM inspired by Laravel's Eloquent.
- **Testing Utilities**: Built-in support for testing controllers and routes.
- **Session Management**: Support for sessions, cookies, and authentication.
- **Queueing System**: A lightweight system for background tasks.
- **Caching**: Built-in caching support for performance optimization.

## Contribution

We welcome contributions! If you'd like to help develop Garavel, feel free to open an issue or create a pull request. Contributions can include bug fixes, new features, or documentation improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- Inspired by Laravel, a powerful PHP framework.
- Also influenced by Go frameworks like Gin and Echo.

## TODO

- Repository
- Factory
- Base tests (transaction)
- Migrations
- Gate
- Service
- Command
- Collections

---

Stay tuned for updates as Garavel evolves into a powerful yet lightweight Go framework!
