# System Design in Go

## Building Web APIs using the Code First Approach and Swagger 2.0

This is an example of using the _Code First Approach_ to build a Web API and generating the corresponding documentation in `Swagger 2.0`, the _Code First Approach_ is:

* Traditional
* API is coded directly
* Documentation can be generated afterwards

This example includes two folders:

* [start/](start/): includes the initial example shown and
* [completed/](completed/): includes the final example shown

## Hypothetical Example

Create an API that includes options for:

* Listing Users and
* Creating Users

Where an `User` has a:

* `name` string (minimum length 5)
* `birth_year` int (minimum value 1900, maximum value 2022)

## Installing `go-swagger`

Notice this project is using `direnv`, for more details you can refer to a post [I wrote](https://mariocarrion.com/2020/11/20/golang-go-tool-direnv.html).

To install run the following instruction:

```
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
```

### Generating spec from comments

Run the following instruction:

```
swagger generate spec --scan-models --output=./swagger.yaml
```

### Generating server from spec

Run the following instruction:

```
swagger generate spec --scan-models --output=./swagger.yaml
```

## Recommended Literature

* [The Design of Web APIs](https://amzn.to/2PNtQvi) (affiliate link)
* [Irresistible APIs: Designing web APIs that developers will love](https://amzn.to/38n5ChW) (affiliate link)
