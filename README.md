Simple Key-Value storage, made to learn Go.

Based on [Build Redis from scratch](https://www.build-redis-from-scratch.dev). List of my improvements:
- Structured the project
- Optimized RESP memory usage by using `rune` instead of `string` for the `Type` field
- Env configuration support (with Viper module)
- Multiple connections support
- Client for get/set operations and simple protocol test

TODO: 
- Logging (with metadata, like level and timestamp); configurable logging level
- Merge `RESP` with `Writer` into one single struct `Protocol` and create it once per connection.