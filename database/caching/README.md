# Caching

- A cache is a high-speed data storage layer which stores a subset of data, typically transient ( _lasting only for a short time_ ) in nature, so that future requests for that data are served up faster than is possible by accessing the data’s primary storage location.
- Caching allows you to efficiently reuse previously retrieved or computed data.
- Caching is storing copies of files in a cache so that cached content can be accessed quickly.
- The cache is a smaller and faster memory that stores copies of the frequently used data.


---


#### Why Caching?

- Reduce Latency
- Increase capacity to handle requests
- Improve app availability


---


#### Types of Web Cache

- Site Cache - Ideal for static content.
- Browser Cache - A popular and effective client-side cache option.
- Server Cache - Best for high-traffic websites that need to reduce server strain.
- Micro Cache - A targeted option for highly dynamic sites.


---


#### Caching Strategies

- Cache-Aside ( _App hits the Cache & DB_ ) - The application is responsible for fetching data from the database and populating the cache.
    1. Cache hit - Read from Cache
    2. Cache miss - Read from Database
    3. Write data into Cache

- Read-Through Cache ( _App hits the cache & cache hits the db if cache miss and then set the cache_ ) - This logic is usually supported by the library or stand-alone cache provider.
    1. Cache hit - Reads from Cache ( _app -> cache_ )
    2. Cache miss - Read from database ( _cache -> database_ )
    3. Write into Cache ( _cache has responsibility if cache miss then fetch from db and write into cache_ )

- Write-Through Cache
  - Data is first written to the cache and then to the database.
  - The cache sits in-line with the database and writes always go through the cache to the main database.
  - It helps cache maintain consistency with the main database.

- Write-Back or Write-Behind
  - This is very similar to to Write-Through but there’s one crucial difference:
    - In Write-Through, the data written to the cache is _synchronously_ updated in the main database. 
    - In Write-Back, the data written to the cache is _asynchronously_ updated in the main database.


---

#### Use Cases

- CDN
- Web Caching
- APIs