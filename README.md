# Mini Social Media API

A backend RESTful API for a mini social media application using Go and Gin.

## Features

1. Create, update, and retrieve posts.
2. Like posts.
3. Add comments to posts.
4. Retrieve post details, including likes and comments.

## Setup Instructions

1. Clone the repository.
2. Install dependencies:


## API Endpoints

- **POST** `/posts/`: Create a new post.
- **PUT** `/posts/:id`: Update a post by ID.
- **GET** `/posts/`: Retrieve all posts.
- **POST** `/posts/:id/like`: Like a specific post by ID.
- **POST** `/posts/:id/comment`: Add a comment to a post by ID.
- **GET** `/posts/:id`: Get post details by ID.

## Assumptions

- Data is stored in memory (no database).
- Concurrency is managed using `sync.RWMutex`.

## Suggestions for Improvement

1. Add persistent storage (e.g., PostgreSQL).
2. Implement user authentication.
3. Add pagination for large datasets.

## Postman Collection

https://documenter.getpostman.com/view/9463833/2sAYBUDY8v
