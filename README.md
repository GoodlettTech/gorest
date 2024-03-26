# Goodlett.Tech

## Table of Contents

- [Goodlett.Tech](#goodletttech)
	- [Table of Contents](#table-of-contents)
	- [Features](#features)
	- [Devcontainer](#devcontainer)
		- [Configuration](#configuration)
	- [Logging](#logging)
	- [Visualization](#visualization)
	- [Documentation](#documentation)


## Features

The application offers the following features:

- User Creation: Users can create an account by providing their credentials.
- User Login: Users can log in to their account using their username and password.
- Error Handling: The backend includes general error handling to provide meaningful error messages to users.
- Middleware Configuration: The backend is configured with middleware to handle authentication and other common tasks.
- Loki Logging: The project utilizes Loki for logging, allowing for centralized log storage and analysis.
- Prometheus Metrics: The backend includes Prometheus exporters to collect and expose metrics, providing insights into the application's performance.
- Postgres Database: The application uses a Postgres database to store user data.
- PgAdmin: PgAdmin is included to provide a user-friendly interface for managing and viewing the database.
- Frontend with SolidJS: The frontend is built with SolidJS and includes a login page and a user creation page. These pages interact with the backend routes for user creation and login.

## Devcontainer

The project includes a devcontainer configuration with the following containers:

- Application container
- PostgreSQL database
- PgAdmin
- Prometheus
- Loki
- Grafana

The majority of the development is done inside the application container, which is configured for Go and JavaScript. This setup allows for the development of a Go-based REST API backend and a frontend using SolidJS to consume it. With these containers pre-configured with all the necessary tools, the only requirements for development are Docker, VS Code, and the VS Code Dev Container extension.

### Configuration

The configuration of my applicationtiondonea .env filee dev con that is passed tois means that the  when launching it. This meansthrough the envirhasbles, b to these requires rthrough thearting the container.
, but updatingm requires restarting the

The backend is configured with a prometheus exporter to enable the prometheus container to scrape metrics from it. The metrics include data such as the success and error rates for each api endpoint based on the HTTP Method. This allows for very deep insights into how the backend is being utilized.

## Logging

The project utilizes the promtail operator to export logs to the loki container. Although the current log implementation is limited, future enhancements will involve integrating logs and metrics using labels to provide a comprehensive understanding of the application's behavior and performance.

## Visualization

Utilizing Grafana, I have implemented a comprehensive log and metrics visualization solution. In the near future, I plan to develop a visually appealing dashboard to enhance data accessibility and comprehension. Additionally, I aim to incorporate alerting functionality to ensure timely response to critical events. While not currently essential, this feature will further enhance the monitoring capabilities of the system.

## Documentation

This project utilizes Swagger documentation to provide a unified and organized interface for viewing comprehensive information about its API endpoints. By integrating code and comments within the same file, the documentation becomes easier to maintain and facilitates seamless searching and interaction. This approach enhances the professionalism and efficiency of managing the project's documentation.

