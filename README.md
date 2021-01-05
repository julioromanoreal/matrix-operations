## Matrix Handling Web Server

Program to handle a matrix from a CSV file and perform a few operations on it.

-----

<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#future-improvements">Future improvements</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

---

## About The Project

This is a simple project to handle a matrix from a CSV file and perform a few operations on it.

Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

The following operations should be available:

1. Echo
   - Return the matrix as a string in matrix format.
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. Invert
   - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. Flatten
   - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 
4. Sum
   - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. Multiply
   - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. `resources/matrix.csv` is an example of a valid input.

### Built With

* Go 1.14

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

Make sure Go is installed in your machine. Run the following command to confirm:

```sh-session
$ go version
go version go1.14 windows/amd64
```

### Installation

1. Clone the repo
```sh
$ git clone https://github.com/julioromanoreal/matrix-operations.git
```
2. Start the application
```sh
$ go run cmd/matrix-operations/main.go
```

## Usage

The following endpoints will be available:

1. Echo: `http://localhost:8080/echo`:
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```
2. Flatten: `http://localhost:8080/flatten`:
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"
```
3. Invert: `http://localhost:8080/invert`:
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/invert"
```
4. Sum: `http://localhost:8080/sum`:
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"
```
5. Multiply: `http://localhost:8080/multiply`:
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"
```

## Future improvements
* Consider protecting the URLs using JWT
* Consider using [gorilla/mux](https://github.com/gorilla/mux) in the web server

## Contributing

Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Julio Romano - [@julioromano_](https://twitter.com/julioromano_) - julio.romano@gmail.com

Project Link: [https://github.com/julioromanoreal/matrix-operations](https://github.com/julioromanoreal/matrix-operations)
