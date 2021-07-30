# SlackMessageSender

Для работы программы нужны файлы:
- hooks.json, где хранится информация в  
виде {"канал": "хук для канала", б}
- message.json, где хранится информация в  
виде {"channels": [{"channel": "имя канала", "text": "текст сообщения"}]}
Пример отправки представлен в cmd/SlackMessageSender/main.go
