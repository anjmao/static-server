# About
Static files server useful for serving local static files like single page apps. Static server also includes gzip compression which helps to estimate production bundle sizes.

## Installation

```
go get github.com/anjmao/static-server
```

## Usage


Serve from current directory
```
static-server
```

Serve from another directory
```
static-server -dir ./dist
```
