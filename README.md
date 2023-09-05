# wb-L0

Учебное задание для Wildberries. <Br>
Сервис генерирует данные, которые в последствии сохраняются в кэш и базу. При запуске сервера происходит выгрузка данных в кэш. <br>
Доступный хендлер: /{id} - запросить информацию о заказе по id.<br>

# Результаты нагрузочного тестирования
Для тестирования использовалась Vegeta. Файлы с результатами доступны в директории /test-reports.

## rps=5000; duration=15s
| Название параметра  | Значения |
| ------------- | ------------- |
| Requests      [total, rate, throughput]   | 75000, 5000.07, 4999.97  |
| Duration      [total, attack, wait]  | 15s, 15s, 285.682µs |
| Latencies     [min, mean, 50, 90, 95, 99, max] | 87.937µs, 283.046µs, 226.186µs, 386.038µs, 484.738µs, 904.841µs, 41.682ms  |
| Bytes In      [total, mean]   | 130950000, 1746.00  |
| Bytes Out     [total, mean]  | 0, 0.00  |
| Success       [ratio]   | 100.00%  |
| Status Codes  [code:count]  | 200:75000  |

## rps=5000; duration=60s
| Название параметра  | Значения |
| ------------- | ------------- |
| Requests      [total, rate, throughput]   | 300000, 5000.02, 5000.00  |
| Duration      [total, attack, wait]  | 1m0s, 1m0s, 259.037µs |
| Latencies     [min, mean, 50, 90, 95, 99, max] | 80.483µs, 255.58µs, 225.697µs, 361.976µs, 428.831µs, 680.511µs, 7.408ms  |
| Bytes In      [total, mean]   | 523800000, 1746.00  |
| Bytes Out     [total, mean]  | 0, 0.00  |
| Success       [ratio]   | 100.00%  |
| Status Codes  [code:count]  | 200:300000  |


## rps=10000; duration=60s
| Название параметра  | Значения |
| ------------- | ------------- |
| Requests      [total, rate, throughput]   | 600000, 10000.04, 10000.01  |
| Duration      [total, attack, wait]  | 1m0s, 1m0s, 160.989µs |
| Latencies     [min, mean, 50, 90, 95, 99, max] | 73.133µs, 148.138µs, 108.496µs, 261.058µs, 321.066µs, 592.255µs, 8.095ms  |
| Bytes In      [total, mean]   | 1047600000, 1746.00  |
| Bytes Out     [total, mean]  | 0, 0.00  |
| Success       [ratio]   | 100.00%  |
| Status Codes  [code:count]  | 200:600000  |

## rps=20000; duration=60s
| Название параметра  | Значения |
| ------------- | ------------- |
| Requests      [total, rate, throughput]   | 1200000, 20000.13, 20000.09  |
| Duration      [total, attack, wait]  | 1m0s, 1m0s, 116.137µs |
| Latencies     [min, mean, 50, 90, 95, 99, max] | 65.639µs, 193.971µs, 112.643µs, 244.805µs, 466.087µs, 1.463ms, 64.505ms  |
| Bytes In      [total, mean]   | 2095200000, 1746.00  |
| Bytes Out     [total, mean]  | 0, 0.00  |
| Success       [ratio]   | 100.00%  |
| Status Codes  [code:count]  | 200:1200000    |

## rps=20000; duration=120s
| Название параметра  | Значения |
| ------------- | ------------- |
| Requests      [total, rate, throughput]   | 2400000, 20000.05, 20000.00  |
| Duration      [total, attack, wait]  | 2m0s, 2m0s, 296.933µs |
| Latencies     [min, mean, 50, 90, 95, 99, max] | 67.581µs, 252.213µs, 113.934µs, 283.631µs, 593.228µs, 2.456ms, 97.701ms  |
| Bytes In      [total, mean]   | 4190400000, 1746.00  |
| Bytes Out     [total, mean]  | 0, 0.00  |
| Success       [ratio]   | 100.00%  |
| Status Codes  [code:count]  | 200:2400000    |

## rps=30000; duration=120s
| Название параметра  | Значения |
| ------------- | ------------- |
| Requests      [total, rate, throughput]   | 3600036, 30000.35, 30000.24  |
| Duration      [total, attack, wait]  | 2m0s, 2m0s, 447.893µs |
| Latencies     [min, mean, 50, 90, 95, 99, max] | 61.155µs, 707.802µs, 155.128µs, 1.87ms, 3.068ms, 8.019ms, 83.352ms  |
| Bytes In      [total, mean]   | 6285662856, 1746.00  |
| Bytes Out     [total, mean]  | 0, 0.00  |
| Success       [ratio]   | 100.00%  |
| Status Codes  [code:count]  | 200:3600036    |


## rps=50000; duration=120s
| Название параметра  | Значения |
| ------------- | ------------- |
| Requests      [total, rate, throughput]   | 1794855, 27153.04, 27109.02  |
| Duration      [total, attack, wait]  | 1m6s, 1m6s, 8.386ms |
| Latencies     [min, mean, 50, 90, 95, 99, max] | 82.903µs, 547.374ms, 386.478ms, 1.204s, 1.577s, 3.093s, 5.491s  |
| Bytes In      [total, mean]   | 3129134058, 1743.39  |
| Bytes Out     [total, mean]  | 0, 0.00  |
| Success       [ratio]   | 99.85%  |
| Status Codes  [code:count]  | 0:2682  200:1792173      |
