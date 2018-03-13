# httpdperfd

Simple httpd for performance testing different load aspects



### Prerequisites

goland and possibly docker



## Getting Started
Binds to port 8000. 

```
go build httpdperfd.go -o httpdperfd
```
or
```
docker build -t <somename> -f Docker.build .
```



## Running the tests
Query Params, both optional:
* weight (default 1000), num shasum iterations to simulate cpu intensive load
* response_body_bytes, size of the response object. Defaults to the timing of the http query in total.

```
curl http://localhost:8000/?weight=1&response_body_bytes=10000
```

We use [SemVer](http://semver.org/) for versioning. For the versions available,
see the [tags on this repository](https://github.com/mattiasberge/httpdperfd/tags). 

## Authors

* **Mattias Berge** - *Initial work*


## License

This project is licensed under the MIT License - see the
[LICENSE.md](LICENSE.md) file for details

## Acknowledgments

...

