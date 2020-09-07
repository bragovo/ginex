```
wrk -t12 -c400 -d30s http://localhost:8080/api/posts
```
```
Running 30s test @ http://localhost:8080/api/posts
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   970.70ms  347.44ms   2.00s    70.87%
    Req/Sec    24.87     17.96   110.00     76.69%
  7889 requests in 30.10s, 546.48MB read
  Socket errors: connect 0, read 316, write 0, timeout 2012
Requests/sec:    262.10
Transfer/sec:     18.16MB
```
