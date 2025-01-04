Simple Key-Value storage, made to learn Go (and try Vim).

[Build Redis from scratch](https://www.build-redis-from-scratch.dev) was used as a reference. List of my improvements:
- Structured the project
- Optimized RESP memory usage by using `rune` instead of `string` for the `Type` field
- Env configuration support (with Viper module)

TODO: 
- Multiple connections support
- Client
- Logging