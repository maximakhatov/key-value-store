Simple Key-Value storage, made to try Go and some of its libraries.

Project is based on tutorial [Build Redis from scratch](https://www.build-redis-from-scratch.dev). List of my improvements:
- Better project structure
- Few minor optimizations
- Env configuration support (with Viper module)
- Logging (with Zerolog module)
- Multiple connections support
- Client for get/set operations
- A couple of tests to try `go test`