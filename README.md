# Comies :hamburger:

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

## 📖 Overview

Comies is a platform that helps food services and stores organize themselves better. We allow users to sell products, manage stocks, keep track of profit. 

## 🧩 Project

### Architecture

As it is intended to make this service increase their capabilities, we used [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) to structure our project. This way, the service functions can grow and change with less possibilities of breaking other parts

```
.
├── app
│      ├── core
│      │        ├── entities # structs representing real world objects that satifies the reason of this app to exist 
│      │        ├── types # general purpose types to be used throughout the application
│      │        └── workflows # usecases (https://martinfowler.com/bliki/UseCase.html)
├── cmd # main files for API and workers
└── docs # swagger and domain documentation
```

### More
Below there are links referencing specific documentation for the entities envolving this service
* [product](docs/product/README.md)
* [category](docs/category/README.md)
* [stock](docs/stock/README.md)
