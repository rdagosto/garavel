# Garavel - A Lightweight Go Framework Inspired by Laravel

**Garavel** is a lightweight yet powerful Go framework inspired by Laravel, designed to make web development simpler, faster, and more enjoyable. The goal of this project is to provide a minimalistic structure while still offering many of the features developers love about Laravel, such as routing, middleware, database integration, and easy configuration management.

**Status**: _In Development_  
**Current Version**: 0.1.0

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
    /internal
        /configs
        /controllers
        /databases
        /factories
        /gates
        /libs
        /middlewares
        /migrations
        /models
        /repositories
        /routes
        /services
        /validators
        /views
    /tests
    main.go
```

- **`internal/`** - Contains the core application logic and is structured into various subdirectories:
- **`configs/`** – Manages configuration settings (e.g., environment variables, database credentials).
- **`controllers/`** – Handles HTTP request logic, coordinating between models and views.
- **`databases/`** – Database connections and utilities for handling migrations.
- **`factories/`** – Provides model factories for generating test data.
- **`gates/`** – Defines authorization rules for user access control.
- **`libs/`** – External or internal utility libraries used throughout the project.
- **`middlewares/`** – Stores middleware functions for request handling (e.g., authentication, logging).
- **`migrations/`** – Database migration files for schema management.
- **`models/`** – Defines database models using GORM.
- **`repositories/`** – Implements data access layers, abstracting database queries.
- **`routes/`** – Manages application routes and defines API endpoints.
- **`services/`** – Business logic layer for separating concerns from controllers.
- **`validators/`** – Implements request validation logic (e.g., form validation, API input sanitization).
- **`views/`** – Contains HTML templates (planned feature).
- **`tests/`** - Holds unit and integration tests for the application.
- **`main.go`** - The entry point of the Garavel application. Initializes the app, loads configurations, and starts the server.

## Contribution

We welcome contributions! If you'd like to help develop Garavel, feel free to open an issue or create a pull request. Contributions can include bug fixes, new features, or documentation improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- Inspired by Laravel, a powerful PHP framework.
- Also influenced by Go frameworks like Gin and Echo.

## TODO

- Auth
- Base tests (transaction)
- Service
- Command
- Collections

---

Stay tuned for updates as Garavel evolves into a powerful yet lightweight Go framework!
