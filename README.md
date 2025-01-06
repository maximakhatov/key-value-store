Simple Key-Value storage, made to learn Go and try some libraries.

Based on [Build Redis from scratch](https://www.build-redis-from-scratch.dev). List of my improvements:
- Project is structured
- Few minor optimizations
- Env configuration support (with Viper module)
- Multiple connections support
- Client for get/set operations and simple protocol test

TODO: 
- Logging (with metadata, like level and timestamp); configurable logging level