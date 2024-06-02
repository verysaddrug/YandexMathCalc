[Russian Version](README.md)

# Technical Documentation for the "Distributed Calculator" Project

## Table of Contents

- [Overview](#overview)
- [Application Launch](#application-launch)
- [Project Structure](#project-structure)
- [Running with Docker](#running-with-docker)
- [How It Works](#how-it-works)
- [Additional Notes](#additional-notes)
- [Conclusion](#conclusion)
- [Contributing](#contributing)
- [License](#license)

## Overview

The **"Distributed Calculator"** project is a distributed computation system that allows for the execution of mathematical expressions by distributing tasks across multiple computational nodes.

## Application Launch

### Prerequisites

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)

### Steps

1. **Clone the repository:**

   ```bash
   git clone https://github.com/verysaddrug/YandexMathCalc
   cd YandexMathCalc
   ```

2. **Start the calculator:**
   ```bash
   go run cmd/calculator/main.go
   ```

3. **Start the orchestrator:**
   ```bash
   go run cmd/orchestrator/main.go
   ```

4. **Open your browser and go to [http://localhost:8000/](http://localhost:8000/).**


## Project Structure

- **cmd/calculator/main.go**: Main file to start the calculator service.
- **cmd/orchestrator/main.go**: Main file to start the orchestrator.
- **handler/**: Package for handling HTTP requests.
  - **add-expression.go**: Handler for adding a new expression.
  - **index.go**: Handler for the main page.
  - **register.go**: Handler for user registration.
  - **auth.go**: Handler for user authentication.
  - **resource.go**: Handler for managing computational resources.
  - **settings.go**: Handler for setting parameters.
  - **login.go**: Handler for user login.
- **calculator/**: Package containing the logic for executing mathematical expressions.
  - **calculator.go**: Main file with the calculator logic.
- **proto/**: Package containing the gRPC protocol description for interaction between services.
  - **calculator.proto**: Description of the gRPC calculator service.
  - **calculator_grpc.pb.go**: Generated gRPC code for the calculator.
  - **calculator.pb.go**: Generated protocol code for the calculator.
- **db/**: Package for database interactions.
  - **user.go**: Logic for handling user data.
  - **setting.go**: Logic for handling settings.
  - **expression.go**: Logic for handling expressions.
  - **init.go**: Database initialization.
- **model/**: Package containing data models.
  - **user.go**: User model.
  - **setting.go**: Setting model.
  - **computing-resource.go**: Computational resource model.
  - **expression.go**: Expression model.
- **templates/**: HTML templates for the web interface.
  - **index.html**: Main page template.
  - **register.html**: Registration page template.
  - **login.html**: Login page template.
  - **settings.html**: Settings page template.
- **orchestrator/**: Package with the logic for orchestrating computational tasks.
  - **orchestrator.go**: Main file with the orchestrator logic.

## Running with Docker

The project includes a `docker-compose.yml` file to facilitate deployment using Docker. To start the application in Docker containers, execute the following command in the project's root directory:

```bash
docker-compose up --build
```

This will automatically create and start the necessary containers for the calculator and orchestrator.

## How It Works

The system consists of two main components: the calculator and the orchestrator. The calculator is responsible for executing mathematical expressions, while the orchestrator distributes tasks among multiple calculators and manages their states. Users submit requests to execute expressions through a web interface, which are then processed by the orchestrator and sent to the calculators for execution.

## Additional Notes

- **Automatic Status Refresh:** Automatic status refresh for expressions is implemented. The page will refresh every 10 seconds to show the most up-to-date information.

## Conclusion

The **"Distributed Calculator"** project provides a powerful system for distributed computations, allowing for efficient load distribution across multiple nodes. By following this documentation, you can easily deploy and run the application and understand the main components and their interactions.

For more detailed information and usage examples, please refer to the source code files and the README included in the project.

## Contributing

We welcome contributions to the project! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add new feature'`).
5. Push to the remote repository (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
