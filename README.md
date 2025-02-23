# blog
Open-source blog system with Gin, GORM, SQLite, JWT middleware backend, and Tailwind CSS frontend. Features admin dashboard, comments, tags, categories. Simple, fast, self-hosted, customizable.

![Blog Screenshot](screenshot.png)

blog/
├── main.go # Main application entry
├── config/ # Configuration directory
│ └── config.go # Config structure and loading logic
├── router/ # Router layer for request handling
│ └── post.go # Blog post related routes
├── model/ # Data model layer
│ └── post.go # Blog post model
├── service/ # Business logic layer
│ └── post.go # Blog post service
├── pkg/ # Reusable packages
│ ├── cache/ # Cache management package
│ └── logger/ # Logging package
├── middleware/ # Custom middleware
│ └── jwt.go # JWT authentication middleware
├── template/ # HTML template files
│ └── default/ # Default theme
│ ├── layout/ # Layout templates
│ │ └── base.html # Base layout
│ ├── post/ # Post related pages
│ │ ├── list.html # Post list page
│ │ └── detail.html # Post detail page
│ └── static/ # Static resource files
│ ├── css/ # Style files
│ │ └── style.css # Compiled TailwindCSS file
│ ├── js/ # JavaScript files
│ └── img/ # Image resources
├── data/ # Data storage directory
│ ├── db/ # Database files (SQLite)
│ ├── upload/ # Upload files storage
│ ├── temp/ # Temporary files
│ └── log/ # Application log files
├── util/ # Utility functions
│ └── tools.go # Common tools and utilities
├── go.mod # Go module file
├── go.sum # Go dependency lockfile
├── package.json # Node.js package config (for TailwindCSS)
└── tailwind.config.js # TailwindCSS configuration


Directory Structure:
1. main.go: Application entry point, initialization and server startup
2. config: Configuration files and loading logic
3. router: Route definitions and request handling
4. model: Data model definitions
5. service: Business logic implementation
6. pkg: Reusable packages and components
   - cache: Redis cache management
7. middleware: Authentication and other middleware components
8. template: HTML template files
9. static: Static resource files
10. data: Data storage directory
11. util: Utility functions

Tech Stack:
- Gin: Web framework
- SQLite: Database
- Redis: External cache system + Built-in cache
- GORM: ORM framework for database operations
- HTML Template: Template engine
- TailwindCSS: CSS framework 
- JWT middleware for authentication
- Configuration management with INI file

