# Caching Proxy Server

A simple CLI tool that implements a caching proxy server. This server forwards requests to an origin server and caches the responses. If the same request is made again, it returns the cached response instead of forwarding the request to the origin server.

## Table of Contents
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Cache Management](#cache-management)
- [Contributing](#contributing)
- [License](#license)

## Features
- Forward requests to an origin server.
- Cache responses to improve performance.
- Return cached responses for repeated requests.
- Indicate whether the response was served from the cache or the origin server.
- Clear the cache using a command.

## Requirements
- Go 1.16 or higher

## Installation
1. Clone the repository:
```bash
   git clone https://github.com/Venukishore-R/caching-proxy.git
   cd caching-proxy
```
2. Make sure you have the necessary dependencies:
```bash
   go mod tidy
```
3. Install the application:
```bash
   make install
```
  It will the run the build and install go command


## Usage
To start the caching proxy server, run the following command:
```bash
  caching-proxy --port <number> --origin <url>
```
### Example
To start the server on port `3000` and forward requests to `http://dummyjson.com`, use:
```bash
   caching-proxy --port 3000 --origin http://dummyjson.com
```
Now, if you make a request to `http://localhost:3000/products`, the caching proxy server will:
- Forward the request to `http://dummyjson.com/products`.
- Return the response along with headers.
- Cache the response for future requests.

### Response Headers
- If the response is served from the cache:
  X-Cache: HIT

- If the response is served from the origin server:
  X-Cache: MISS

## Cache Management
To clear the cache, run the following command:
```bash
   curl <YOUR_URL>/clear-cache
```
### Example
To clear the cache, use:
```bash
   curl http://localhost:3000/clear-cache
```
This command will remove all cached responses.

## Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Conclusion
For further insights into building caching servers and to enhance your understanding of the topic, you can visit [this roadmap](https://roadmap.sh/projects/caching-server).

