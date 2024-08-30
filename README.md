# URL Shortener With Go

[![Go Version](https://img.shields.io/badge/Go-1.21-blue.svg)](https://golang.org/doc/go1.18)
[![MongoDB](https://img.shields.io/badge/MongoDB-v5.0-green)](https://www.mongodb.com/try/download/community)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A modern URL Shortener application built using Go, with a clean architecture and MongoDB integration. This project demonstrates the implementation of a three-tier architecture using Docker for containerization. The CI/CD pipeline is managed with Jenkins, and security analysis is conducted using Trivy. The application is built and pushed to AWS ECR (Elastic Container Registry) and deployed on an AWS EKS (Elastic Kubernetes Service) cluster.

## Table of Contents

- [Project Overview](#project-overview)
- [Features](#features)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Project Overview

The URL Shortener With Go is a web application designed to shorten long URLs into concise links. The application is structured using a three-tier architecture, leveraging Go for the backend, MongoDB for data storage, and a static frontend for user interaction. 

This project is an ideal demonstration for those looking to understand how to create scalable web applications with Go and MongoDB while employing containerization for deployment.

## Features

- **URL Shortening:** Convert long URLs into short, manageable links.
- **Analytics Dashboard:** Track click counts and URL usage statistics.
- **User Authentication:** Secure login and user management.
- **RESTful API:** Comprehensive API for integration with other services.
- **Dockerized Environment:** Seamless setup and deployment with Docker.
- **Scalable Design:** Three-tier architecture for robust and scalable solutions.

## Architecture

The project follows a three-tier architecture:

1. **Presentation Layer (Frontend):** 
   - A static landing page built with HTML, CSS, and JavaScript. 
   - Provides an intuitive interface for users to shorten URLs and view analytics.

2. **Application Layer (Backend):** 
   - Built with Go, handling all business logic.
   - Exposes RESTful API endpoints for URL operations.

3. **Data Layer (Database):**
   - MongoDB used for storing URL mappings and user data.
   - Ensures high availability and easy scalability.

```plaintext
   +-------------------+
   |  Presentation     |
   |    (Frontend)     |
   +-------------------+
            |
            |
   +-------------------+
   |   Application     |
   |   (Backend/Go)    |
   +-------------------+
            |
            |
   +-------------------+
   |       Data        |
   |   (MongoDB)       |
   +-------------------+
```

## Tech Stack

- **Go (Golang):** v1.21
- **MongoDB:** v5.0
- **Docker:** v20.10
- **HTML/CSS/JavaScript:** Static files for frontend

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.21 or higher)
- [Docker](https://docs.docker.com/get-docker/)
- [MongoDB](https://www.mongodb.com/try/download/community)

### Clone the Repository

```bash
git clone https://github.com/ARUP-G/URL-Shortener-With-GO.git
cd URL-Shortener-With-GO
```

### Build and Run

#### Without Docker:

1. **Start MongoDB:**
   
   Ensure MongoDB is running on your machine. You can start it with:

   ```bash
   mongod --config /usr/local/etc/mongod.conf --fork
   ```

2. **Install Dependencies:**

   Navigate to the backend folder and install dependencies:

   ```bash
   cd backend
   go mod tidy
   ```

3. **Run the Application:**

   ```bash
   go run main.go
   ```

4. **Open the Frontend:**

   Open the `index.html` file in your browser to start using the application.

#### With Docker:

1. **Build and Start Containers:**

   ```bash
   docker-compose up --build
   ```

2. **Access the Application:**

   Visit `http://localhost:8181` in your browser.

## Usage

- **Shorten URL:** Enter a long URL on the landing page to receive a shortened link.
- **Track URL:** Log in to access analytics and view click statistics for your shortened URLs.
- **API Access:** Use the provided API to integrate URL shortening functionality into your applications.

### API Endpoints

| Method | Endpoint              | Description                       |
|--------|-----------------------|-----------------------------------|
| POST   | `/api/v1/shorten`     | Shortens a provided URL           |
| GET    | `/api/v1/{shortURL}`  | Redirects to the original URL     |


### Continuous Integration

The project includes a GitHub Actions workflow to automate testing and ensure code quality.

## Contributing

Contributions are welcome! Please read the [CONTRIBUTING.md](CONTRIBUTING.md) file for guidelines on how to contribute to this project.

1. Fork the repository.
2. Create a new feature branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add your feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries or feedback, feel free to reach out:

- **GitHub:** [ARUP-G](https://github.com/ARUP-G)
- **Email:** [arupdascontact@gmail.com](mailto:arupdascontact@gmail.com)

---

### Acknowledgments

- Special thanks to the open-source community for providing valuable resources and inspiration.
- Inspired by various URL shortener projects available online.

---

Feel free to customize the sections and content according to your project's specific details and requirements. If you have any additional features or specific instructions, you can add them to the appropriate sections.

