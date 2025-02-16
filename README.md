# Agoravote App Backend

Agoravote is a backend application built with Go that serves as a platform for managing groups, posts, votes, and user authentication. This README provides an overview of the project, setup instructions, and usage guidelines.

## Features

- **Group Management**: Create and manage groups.
- **Post Management**: Create and retrieve posts within groups.
- **Voting System**: Cast and retrieve votes on posts.
- **User Authentication**: Login and manage user sessions.

## Project Structure

```
agoravote-app-backend
├── src
│   ├── controllers
│   ├── models
│   ├── routes
│   ├── services
│   └── main.go
├── go.mod
└── README.md
```

## Setup Instructions

1. **Clone the Repository**:
   ```
   git clone https://github.com/yourusername/agoravote-app-backend.git
   cd agoravote-app-backend
   ```

2. **Install Dependencies**:
   Ensure you have Go installed. Run the following command to download the necessary dependencies:
   ```
   go mod tidy
   ```

3. **Run the Application**:
   Start the server by executing:
   ```
   go run src/main.go
   ```

4. **Access the API**:
   The server will start on `http://localhost:8080`. You can use tools like Postman or curl to interact with the API endpoints.

## Usage Guidelines

- **Groups**:
  - Create a group: `POST /groups`
  - Get all groups: `GET /groups`

- **Posts**:
  - Create a post: `POST /posts`
  - Get all posts: `GET /posts`

- **Votes**:
  - Cast a vote: `POST /votes`
  - Get all votes: `GET /votes`

- **User Authentication**:
  - User login: `POST /users/login`
  - Get user details: `GET /users/{id}`

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.