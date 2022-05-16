# Question1

```zsh
junjie@ubuntu:~$ redis-benchmark -t get,set -d 10 -n 1000000
====== SET ======
  1000000 requests completed in 6.31 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

99.91% <= 1 milliseconds
99.94% <= 2 milliseconds
99.95% <= 15 milliseconds
99.95% <= 16 milliseconds
99.95% <= 17 milliseconds
99.96% <= 18 milliseconds
99.97% <= 19 milliseconds
99.99% <= 22 milliseconds
99.99% <= 25 milliseconds
100.00% <= 36 milliseconds
100.00% <= 36 milliseconds
158403.30 requests per second

====== GET ======
  1000000 requests completed in 6.01 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

99.98% <= 1 milliseconds
100.00% <= 1 milliseconds
166389.34 requests per second


junjie@ubuntu:~$ redis-benchmark -t get,set -d 20 -n 1000000
====== SET ======
  1000000 requests completed in 6.15 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

99.96% <= 1 milliseconds
99.97% <= 2 milliseconds
99.97% <= 3 milliseconds
99.98% <= 15 milliseconds
99.98% <= 16 milliseconds
99.98% <= 18 milliseconds
99.99% <= 19 milliseconds
99.99% <= 25 milliseconds
99.99% <= 26 milliseconds
100.00% <= 38 milliseconds
100.00% <= 39 milliseconds
100.00% <= 39 milliseconds
162654.53 requests per second

====== GET ======
  1000000 requests completed in 6.01 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

99.99% <= 1 milliseconds
100.00% <= 1 milliseconds
166334.00 requests per second


junjie@ubuntu:~$ redis-benchmark -t get,set -d 50 -n 1000000
====== SET ======
  1000000 requests completed in 6.04 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

99.98% <= 1 milliseconds
100.00% <= 1 milliseconds
165453.34 requests per second

====== GET ======
  1000000 requests completed in 6.02 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

99.99% <= 1 milliseconds
100.00% <= 1 milliseconds
166030.22 requests per second


junjie@ubuntu:~$ redis-benchmark -t get,set -d 100 -n 1000000
====== SET ======
  1000000 requests completed in 6.19 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

99.94% <= 1 milliseconds
99.97% <= 17 milliseconds
99.98% <= 18 milliseconds
99.99% <= 35 milliseconds
100.00% <= 39 milliseconds
100.00% <= 40 milliseconds
100.00% <= 40 milliseconds
161498.70 requests per second

====== GET ======
  1000000 requests completed in 6.05 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

99.98% <= 1 milliseconds
100.00% <= 2 milliseconds
100.00% <= 2 milliseconds
165289.25 requests per second


junjie@ubuntu:~$ redis-benchmark -t get,set -d 200 -n 1000000
====== SET ======
  1000000 requests completed in 6.00 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

99.96% <= 1 milliseconds
100.00% <= 1 milliseconds
166638.89 requests per second

====== GET ======
  1000000 requests completed in 5.96 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

99.98% <= 1 milliseconds
100.00% <= 1 milliseconds
167785.23 requests per second


junjie@ubuntu:~$ redis-benchmark -t get,set -d 1000 -n 1000000
====== SET ======
  1000000 requests completed in 6.21 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1

99.99% <= 1 milliseconds
100.00% <= 1 milliseconds
161108.44 requests per second

====== GET ======
  1000000 requests completed in 6.07 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1

99.98% <= 1 milliseconds
100.00% <= 1 milliseconds
164744.64 requests per second


junjie@ubuntu:~$ redis-benchmark -t get,set -d 5000 -n 1000000
====== SET ======
  1000000 requests completed in 6.91 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1

99.98% <= 1 milliseconds
100.00% <= 2 milliseconds
144613.16 requests per second

====== GET ======
  1000000 requests completed in 6.65 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1

99.99% <= 1 milliseconds
100.00% <= 1 milliseconds
150285.55 requests per second

```

---

# Question2

```zsh
junjie@ubuntu:~/tmp$ redis-cli
127.0.0.1:6379> memory usage 10_bytes_0
(integer) 66
127.0.0.1:6379> memory usage 20_bytes_0
(integer) 76
127.0.0.1:6379> memory usage 50_bytes_0
(integer) 108
127.0.0.1:6379> memory usage 100_bytes_0
(integer) 159
127.0.0.1:6379> memory usage 200_bytes_0
(integer) 259
127.0.0.1:6379> memory usage 1000_bytes_0
(integer) 1062
127.0.0.1:6379> memory usage 5000_bytes_0
(integer) 5062
127.0.0.1:6379> 
```