# Project Setup

This project is a Go-based application that uses Firebase for authentication, Firestore as a NoSQL database, and IPFS for decentralized storage. It also uses a variety of design patterns such as the Visitor pattern.

## Prerequisites

- Go 1.16 or later
- Firebase account
- IPFS installed and running locally

## Steps to Setup

1. Clone the repository to your local machine.
2. Install the required Go packages by running `go mod tidy`.
3. Set up Firebase:
    - Create a new Firebase project.
    - Enable Email/Password sign-in under the Authentication sign-in method.
    - Generate a new private key file for your service account.
    - Save the JSON file and note down the path.
4. Set up environment variables:
    - Create a `.env` file in the root directory.
    - Add the following lines to the file:
      ```
      FIREBASE_CREDENTIALS_PATH=<path_to_your_firebase_credentials_json_file>
      APP_PORT=8080
      IPFS_MESSAGE="Hello, IPFS!"
      ```
5. Run the application by executing `go run main.go`.

# Documentation

## `main.go`

This is the entry point of the application. It loads environment variables, initializes Firebase, sets up the HTTP server, and listens for incoming requests.

## `server.go`

This file sets up the HTTP router and defines the routes for the application.

## `controllers` package

This package contains the HTTP handlers for the application. Each file corresponds to a different part of the application:

- `commentsController.go`: Handles HTTP requests related to comments.
- `entitiesController.go`: Handles HTTP requests related to entities.
- `payementsController.go`: Handles HTTP requests related to payments.
- `ratingController.go`: Handles HTTP requests related to ratings.

## `models` package

This package contains the data structures used in the application:

- `comment.go`: Defines the `Comment` struct and the `Visitor` interface for the Visitor pattern.
- `entities.go`: Defines the `Entity` struct.
- `rating.go`: Defines the `Rating` struct.
- `user.go`: Defines the `User` struct.
- `visitorComment.go`: Defines the `BadWordDetector` struct which implements the `Visitor` interface.

## `facade` package

This package contains the business logic of the application. It interacts with the repositories to fetch and store data.

## `repositories` package

This package contains the data access layer of the application. It interacts with Firestore and IPFS to fetch and store data.

## `routes` package

This package sets up the routes for the application. Each file corresponds to a different part of the application:

- `commentsRoutes.go`: Sets up the routes for comments.
- `entitiesRoutes.go`: Sets up the routes for entities.
- `payementRoutes.go`: Sets up the routes for payments.
- `ratingRoutes.go`: Sets up the routes for ratings.

# How It Works

When a request comes in, it is first handled by the router which routes the request to the appropriate handler in the `controllers` package. The handler then calls the appropriate function in the `facade` package to perform the business logic. The `facade` package interacts with the `repositories` package to fetch and store data. The `repositories` package interacts with Firestore and IPFS. The response is then sent back to the client.
