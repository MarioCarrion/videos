# System Design in Go

[Video](https://youtu.be/ErA92edMta8)

## Building Web APIs using the Design First Approach and OpenAPI 3.0

This is an example of using the _Design First Approach_ to build a Web API, first generating the corresponding API Description Format in OpenAPI 3.0 and then implementing the corresponding code.

## Hypothetical Example

Create an API that includes options for:

* Listing Users and
* Creating Users

Where an `User` has a:

* `name` string (minimum length 5)
* `birth_year` int (minimum value 1900, maximum value 2022)

## Installing `goa`

Notice this project is using `direnv`, for more details you can refer to a post [I wrote](https://mariocarrion.com/2020/11/20/golang-go-tool-direnv.html).

To install run the following instruction:

```
go install goa.design/goa/v3/cmd/goa@v3.5.3
```

Then `go get` the go packages:

```
go get goa.design/goa/v3@v3.5.3
```

## Recommended Literature

* [The Design of Web APIs](https://amzn.to/2PNtQvi) (affiliate link)
* [Irresistible APIs: Designing web APIs that developers will love](https://amzn.to/38n5ChW) (affiliate link)
