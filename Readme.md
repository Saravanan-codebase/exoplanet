# Exoplanet 

Exoplanet API is for managing a collection of exoplanets.

## Features

- Add new exoplanets.
- List all exoplanets with optional filtering and sorting.
- Retrieve an exoplanet by its ID.
- Update an existing exoplanet.
- Delete an exoplanet.
- Calculate fuel estimation for traveling to an exoplanet.

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/Saravanan-codebase/exoplanet.git
    cd exoplanet
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

### Running the API LOCAL

1. Run the application:

    ```bash
    go run main.go
    ```

2. The API will be available at `http://localhost:8080`.

### Running the API DOCKER

1. Build the Docker Image:

docker build -t exoplanet

2. Run the Docker Container:

docker run -p 8080:8080 exoplanet



