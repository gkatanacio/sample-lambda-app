# go-serverless-template

Template for serverless golang projects.

Since this tries to follow the [3 Musketeers](https://3musketeers.io) methodology, the following are heavily used:
- Docker
- Docker Compose
- Make

## Usage

#### configure
```bash
$ make .env
```
* see generated `.env` file for configuration

**Note:** For deployment via CI/CD, `CICD_MODE` environment variable should be set to `true` in the build server. All configuration should also be set in the build server's environment variables. They will automatically be used if `.env` is generated from `.env.cicd`.

#### run tests
```bash
$ make test
```

#### tidy dependencies
```bash
$ make deps
```

#### build serverless functions
```bash
$ make build
```

#### deploy serverless application
```bash
$ make deploy
```

### Helpers during development:

#### format all .go files in project (using go fmt)
```bash
$ make fmt
```

#### generate test mocks (to be used with [stretchr/testify](https://github.com/stretchr/testify)) for all interfaces in project
```bash
$ make genMocks
```
