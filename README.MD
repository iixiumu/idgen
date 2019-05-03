#   memory atomic
    ab -n 100000 -c 100 -k 'http://127.0.0.1:8052/getid?type=1'

    Benchmarking 127.0.0.1 (be patient)
    Completed 10000 requests
    Completed 20000 requests
    Completed 30000 requests
    Completed 40000 requests
    Completed 50000 requests
    Completed 60000 requests
    Completed 70000 requests
    Completed 80000 requests
    Completed 90000 requests
    Completed 100000 requests
    Finished 100000 requests


    Server Software:
    Server Hostname:        127.0.0.1
    Server Port:            8052

    Document Path:          /getid?type=1
    Document Length:        14 bytes

    Concurrency Level:      100
    Time taken for tests:   9.093 seconds
    Complete requests:      100000
    Failed requests:        0
    Keep-Alive requests:    100000
    Total transferred:      16100000 bytes
    HTML transferred:       1400000 bytes
    Requests per second:    10997.03 [#/sec] (mean)
    Time per request:       9.093 [ms] (mean)
    Time per request:       0.091 [ms] (mean, across all concurrent requests)
    Transfer rate:          1729.02 [Kbytes/sec] received

    Connection Times (ms)
                min  mean[+/-sd] median   max
    Connect:        0    0   0.9      0      42
    Processing:     0    9   9.1      8     152
    Waiting:        0    9   9.1      8     152
    Total:          0    9   9.2      8     152

    Percentage of the requests served within a certain time (ms)
    50%      8
    66%      9
    75%     10
    80%     11
    90%     16
    95%     20
    98%     28
    99%     40
    100%    152 (longest request)
#   memory mutex
    ab -n 100000 -c 100 -k 'http://127.0.0.1:8052/getid?type=2'

    Benchmarking 127.0.0.1 (be patient)
    Completed 10000 requests
    Completed 20000 requests
    Completed 30000 requests
    Completed 40000 requests
    Completed 50000 requests
    Completed 60000 requests
    Completed 70000 requests
    Completed 80000 requests
    Completed 90000 requests
    Completed 100000 requests
    Finished 100000 requests


    Server Software:
    Server Hostname:        127.0.0.1
    Server Port:            8052

    Document Path:          /getid?type=2
    Document Length:        14 bytes

    Concurrency Level:      100
    Time taken for tests:   9.623 seconds
    Complete requests:      100000
    Failed requests:        0
    Keep-Alive requests:    100000
    Total transferred:      16100000 bytes
    HTML transferred:       1400000 bytes
    Requests per second:    10391.31 [#/sec] (mean)
    Time per request:       9.623 [ms] (mean)
    Time per request:       0.096 [ms] (mean, across all concurrent requests)
    Transfer rate:          1633.79 [Kbytes/sec] received

    Connection Times (ms)
                min  mean[+/-sd] median   max
    Connect:        0    0   0.2      0      10
    Processing:     0   10   7.9      9     134
    Waiting:        0   10   7.9      9     134
    Total:          0   10   7.9      9     134

    Percentage of the requests served within a certain time (ms)
    50%      9
    66%     10
    75%     11
    80%     14
    90%     18
    95%     23
    98%     30
    99%     36
    100%    134 (longest request)
#   redis incr
    ab -n 100000 -c 100 -k 'http://127.0.0.1:8052/getid?type=3'

    Benchmarking 127.0.0.1 (be patient)
    Completed 10000 requests
    Completed 20000 requests
    Completed 30000 requests
    Completed 40000 requests
    Completed 50000 requests
    Completed 60000 requests
    Completed 70000 requests
    Completed 80000 requests
    Completed 90000 requests
    Completed 100000 requests
    Finished 100000 requests


    Server Software:
    Server Hostname:        127.0.0.1
    Server Port:            8052

    Document Path:          /getid?type=3
    Document Length:        8 bytes

    Concurrency Level:      100
    Time taken for tests:   12.896 seconds
    Complete requests:      100000
    Failed requests:        99991
    (Connect: 0, Receive: 0, Length: 99991, Exceptions: 0)
    Keep-Alive requests:    100000
    Total transferred:      15888796 bytes
    HTML transferred:       1188895 bytes
    Requests per second:    7754.36 [#/sec] (mean)
    Time per request:       12.896 [ms] (mean)
    Time per request:       0.129 [ms] (mean, across all concurrent requests)
    Transfer rate:          1203.20 [Kbytes/sec] received

    Connection Times (ms)
                min  mean[+/-sd] median   max
    Connect:        0    0   0.3      0      12
    Processing:     0   13   8.3     11     146
    Waiting:        0   13   8.3     11     146
    Total:          0   13   8.4     11     146

    Percentage of the requests served within a certain time (ms)
    50%     11
    66%     14
    75%     16
    80%     17
    90%     22
    95%     27
    98%     34
    99%     42
    100%    146 (longest request)
#   mysql lock
    ab -n 100000 -c 100 -k 'http://127.0.0.1:8052/getid?type=4'
#   mysql cas
    ab -n 100000 -c 100 -k 'http://127.0.0.1:8052/getid?type=5'

    Completed 10000 requests
    Completed 20000 requests
    Completed 30000 requests
    Completed 40000 requests
    Completed 50000 requests
    Completed 60000 requests
    Completed 70000 requests
    Completed 80000 requests
    Completed 90000 requests
    Completed 100000 requests
    Finished 100000 requests


    Server Software:
    Server Hostname:        127.0.0.1
    Server Port:            8052

    Document Path:          /getid?type=5
    Document Length:        8 bytes

    Concurrency Level:      100
    Time taken for tests:   199.457 seconds
    Complete requests:      100000
    Failed requests:        99991
    (Connect: 0, Receive: 0, Length: 99991, Exceptions: 0)
    Keep-Alive requests:    100000
    Total transferred:      15888796 bytes
    HTML transferred:       1188895 bytes
    Requests per second:    501.36 [#/sec] (mean)
    Time per request:       199.457 [ms] (mean)
    Time per request:       1.995 [ms] (mean, across all concurrent requests)
    Transfer rate:          77.79 [Kbytes/sec] received

    Connection Times (ms)
                min  mean[+/-sd] median   max
    Connect:        0    0   0.3      0      13
    Processing:    25  199  52.0    178     535
    Waiting:       11  199  52.0    178     535
    Total:         25  199  52.0    178     535

    Percentage of the requests served within a certain time (ms)
    50%    178
    66%    195
    75%    214
    80%    228
    90%    270
    95%    311
    98%    361
    99%    397
    100%    535 (longest request)   <!-- step=100 -->
    Server Software:
    Server Hostname:        127.0.0.1
    Server Port:            8052

    Document Path:          /getid?type=6
    Document Length:        8 bytes

    Concurrency Level:      100
    Time taken for tests:   9.742 seconds
    Complete requests:      100000
    Failed requests:        99991
    (Connect: 0, Receive: 0, Length: 99991, Exceptions: 0)
    Keep-Alive requests:    100000
    Total transferred:      15888796 bytes
    HTML transferred:       1188895 bytes
    Requests per second:    10264.46 [#/sec] (mean)
    Time per request:       9.742 [ms] (mean)
    Time per request:       0.097 [ms] (mean, across all concurrent requests)
    Transfer rate:          1592.67 [Kbytes/sec] received

    Connection Times (ms)
                min  mean[+/-sd] median   max
    Connect:        0    0   0.3      0      12
    Processing:     0   10   6.2      9     159
    Waiting:        0   10   6.2      9     159
    Total:          0   10   6.2      9     159

    Percentage of the requests served within a certain time (ms)
    50%      9
    66%     10
    75%     11
    80%     11
    90%     16
    95%     19
    98%     26
    99%     30
    100%    159 (longest request)
#   mysql pessimism step
    ab -n 100000 -c 100 -k 'http://127.0.0.1:8052/getid?type=6'

    Benchmarking 127.0.0.1 (be patient)
    Completed 10000 requests
    Completed 20000 requests
    Completed 30000 requests
    Completed 40000 requests
    Completed 50000 requests
    Completed 60000 requests
    Completed 70000 requests
    Completed 80000 requests
    Completed 90000 requests
    Completed 100000 requests
    Finished 100000 requests

    <!-- step=10 -->
    Server Software:
    Server Hostname:        127.0.0.1
    Server Port:            8052

    Document Path:          /getid?type=6
    Document Length:        8 bytes

    Concurrency Level:      100
    Time taken for tests:   29.523 seconds
    Complete requests:      100000
    Failed requests:        99991
    (Connect: 0, Receive: 0, Length: 99991, Exceptions: 0)
    Keep-Alive requests:    100000
    Total transferred:      15888796 bytes
    HTML transferred:       1188895 bytes
    Requests per second:    3387.15 [#/sec] (mean)
    Time per request:       29.523 [ms] (mean)
    Time per request:       0.295 [ms] (mean, across all concurrent requests)
    Transfer rate:          525.56 [Kbytes/sec] received

    Connection Times (ms)
                min  mean[+/-sd] median   max
    Connect:        0    0   0.2      0       9
    Processing:     0   29  12.5     26     215
    Waiting:        0   29  12.5     26     215
    Total:          0   29  12.5     26     215

    Percentage of the requests served within a certain time (ms)
    50%     26
    66%     28
    75%     29
    80%     31
    90%     36
    95%     52
    98%     69
    99%     83
    100%    215 (longest request)

    <!-- step=100 -->
    Server Software:
    Server Hostname:        127.0.0.1
    Server Port:            8052

    Document Path:          /getid?type=6
    Document Length:        8 bytes

    Concurrency Level:      100
    Time taken for tests:   9.742 seconds
    Complete requests:      100000
    Failed requests:        99991
    (Connect: 0, Receive: 0, Length: 99991, Exceptions: 0)
    Keep-Alive requests:    100000
    Total transferred:      15888796 bytes
    HTML transferred:       1188895 bytes
    Requests per second:    10264.46 [#/sec] (mean)
    Time per request:       9.742 [ms] (mean)
    Time per request:       0.097 [ms] (mean, across all concurrent requests)
    Transfer rate:          1592.67 [Kbytes/sec] received

    Connection Times (ms)
                min  mean[+/-sd] median   max
    Connect:        0    0   0.3      0      12
    Processing:     0   10   6.2      9     159
    Waiting:        0   10   6.2      9     159
    Total:          0   10   6.2      9     159

    Percentage of the requests served within a certain time (ms)
    50%      9
    66%     10
    75%     11
    80%     11
    90%     16
    95%     19
    98%     26
    99%     30
    100%    159 (longest request)
#   issue
1. ab -k Use HTTP KeepAlive feature
2. redis pool MaxIdle > 100
