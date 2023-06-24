# url-shortener-openapi

`url-shortener-openapi` is a backend implementation for a url shortener service written in golang with REST openapi 3.0 server specification and advantages Hexagonal architecture (ports & adaptors) and DDD (domain-driven design).

[![Tests](https://github.com/aria3ppp/url-shortener-openapi/actions/workflows/tests.yml/badge.svg)](https://github.com/aria3ppp/url-shortener-openapi/actions/workflows/tests.yml)
[![Coverage Status](https://coveralls.io/repos/github/aria3ppp/url-shortener-openapi/badge.svg?branch=master)](https://coveralls.io/github/aria3ppp/url-shortener-openapi?branch=master)

### To run the server:

```bash
cp .env.example .env && make server-testdeploy-up
```

Now server is up and running at port `8080` on your `localhost`.
