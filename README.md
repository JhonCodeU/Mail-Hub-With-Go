# Mail-Hub-With-Go
This project is based on Chapter 6 of the book "Mining the Social Web, 2nd Edition" by Matthew A. Russell. The objective is to create an interface to search for information in email databases. The project has several parts, including indexing the email database, profiling the indexer, creating a visualizer, and optional optimization and deployment.

### Technologies used:
* Go programming language
* ZincSearch
* Profiling tools in Go

This is a web application built with Golang and the Chi router. It provides APIs for managing users and emails.

## Features
* RESTful API built with the github.com/go-chi/chi/v5 router
* CORS enabled with github.com/go-chi/cors


| Method | Endpoint | Description |
| ------ | -------- | ----------- |
| GET | / | Home endpoint |
| GET | /users | Retrieve all users |
| POST | /user | Create a new user |
| PUT | /user | Update an existing user |
| DELETE | /user/{userID} | Delete a user with the specified ID |
| POST | /emails/search | Search for emails matching the given criteria |
| POST | /emails/search_all | Search for all emails |

## Usage
1. Build the application: `go build -o api`
2. Run the application: `air`

## Stats profiling

The application includes stats profiling, which tracks the following metrics:

- The number of requests (`request_count`)
- The latency of requests (`request_latency`)

These metrics are exposed via an endpoint at `/debug/vars`, which returns a JSON object with the metric values.
## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
