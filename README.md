# Goodlett.Tech

## Description

Briefly describe your project, highlighting its main features and functionalities.

## Table of Contents

- [Goodlett.Tech](#goodletttech)
	- [Description](#description)
	- [Table of Contents](#table-of-contents)
	- [Features](#features)
	- [Devcontainer](#devcontainer)
		- [Installation](#installation)
		- [Configuration](#configuration)
	- [Metrics](#metrics)
	- [Logging](#logging)
	- [Visualization](#visualization)
	- [Documentation](#documentation)


## Features
The application currently only really has a log in and user creation page. I had been trying to decide if I wanted to move over to an htmx frontend so I decided to spend some more time building out services around my application so if I decided to change it would require minimal changes.

## Devcontainer

The project comes with a devcontainer config with the following containers:

- application container
- postgres database
- pgadmin
- prometheus
- loki
- grafana

The bulk of work is done inside of the application container which is configured for golang and javascript. This allows for developing a golang rest api in the backend and a solidjs based frontend to consume it. Because these containers are configured with everything needed to develop, the only requirements are docker, vscode, and the dev container plugin for vscode.

### Installation

1. Install Docker
2. Install VS Code
3. install VS Code Dev Containers plugin

### Configuration

The configuration of my application is done through a .env file that is passed to the dev container when launching it. This means that the server has access to these variables through the environment variables, but updating them requires restarting the container.

## Metrics

The backend is configured with a prometheus exporter to enable the prometheus container to scrape metrics from it. The metrics include data such as the success and error rates for each api endpoint based on the HTTP Method. This allows for very deep insights into how the backend is being utilized.

## Logging

Logs are exported to the loki container via the server's promtail operator. Logs are currently kind of limited in the project but once I learn how to tie the logs and metrics together using labels they will provide a rich understanding of what my application is doing.

## Visualization

I am using Grafana to display my log data and metrics data. Soon I will be working on creating a nice dashboard so the data is a bit easier to view and understand. I would also eventually like to add alerting but it isn't really necessary at the moment.

## Documentation

The project's routes are documented using Swagger docs. Swagger is a powerful tool for documenting APIs, providing a clear and interactive documentation for developers to understand the available endpoints, request/response formats, and any additional information.

To access the Swagger documentation when hosting the project, follow these steps:

1. Start the project and ensure it is running.
2. Open a web browser and navigate to the following URL: `http://localhost:PORT/swagger/index.html`, where `PORT` is the port number on which your project is running.
3. The Swagger UI will be displayed, showing a list of all the available routes and their details.
4. Explore the different endpoints, request/response schemas, and any additional information provided in the Swagger documentation.

Having well-documented routes with Swagger not only helps developers understand the API, but also makes it easier to collaborate with others and integrate the project with other systems.