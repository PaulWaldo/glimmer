# Glimmer

Glimmer is a Fyne application that provides a user interface for interacting with the Flickr API.

## Overview
Glimmer allows users to authenticate with Flickr and perform various actions. It utilizes the Fyne GUI toolkit to create a seamless user experience.

## Core Technologies
- Go: The programming language used to develop the application.
- Fyne: A cross-platform GUI toolkit for Go.
- Flickr API: The application interacts with the Flickr API to provide various functionalities.

## Architecture
The application is structured around a main window with a menu and a view stack. The view stack is responsible for managing the different views of the application.

## Getting Started
### Prerequisites
- Go 1.16 or later
- Fyne 2.0 or later

### Installation
```bash
go get -u github.com/PaulWaldo/glimmer
```
### Running the Application
To run the application, navigate to the project directory and execute the following command:
```bash
go run main.go
```
