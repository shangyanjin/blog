#!/bin/bash

# Create main project directory
mkdir -p blog

# Create all subdirectories
cd blog

# Create application directories
mkdir -p config
mkdir -p router
mkdir -p model
mkdir -p service
mkdir -p util

# Create pkg directories
mkdir -p pkg/database
mkdir -p pkg/logger
mkdir -p pkg/middleware

# Create template directories
mkdir -p template/layout
mkdir -p template/post

# Create static directories
mkdir -p static/css
mkdir -p static/js
mkdir -p static/img

# Create data directories
mkdir -p data/db
mkdir -p data/upload
mkdir -p data/temp
mkdir -p data/log

# Create empty files to maintain directory structure
touch main.go
touch go.mod
touch readme.txt

# Create empty template files
touch template/layout/base.html
touch template/post/list.html
touch template/post/detail.html

# Create empty static files
touch static/css/style.css
touch static/js/app.js

# Create empty config file
touch config/config.go

# Create empty router file
touch router/router.go

# Set appropriate permissions
chmod -R 755 .
chmod -R 777 data

# Initialize go module
go mod init blog

# Install dependencies
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

echo "Project structure created successfully!" 