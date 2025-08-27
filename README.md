# package_main
Запустить в docker приложение
И написать к нему pipeline, который будет собирать image, пушить его в docker hub, деплоить его и потом проверять его работоспособность с помощью запроса

curl http://localhost:9200/metrics
Пример ответа

{
  "status": "success",
  "data": {
    "metrics": {
      "cpu_usage": 0.754,
      "memory_usage": 34.21,
      "thread_count": 287
    }
  }
}
