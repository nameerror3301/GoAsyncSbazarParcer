# GoAsyncSbazarParcer

## Описание 

Микросервис для сбора данных с сайта **Sbazar.cz**

Микросервис будет собирать данные формировать их в JSON и отправлять другому микросервису который эти данные будет сохранять.

## Структура проекта

```.
├── app
│   └── cmd
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   └── run.go
│   ├── config
│   │   ├── config.go
│   │   └── config.yaml
│   ├── logs
│   │   └── info.txt
│   └── models
│       └── sbazar.go
└── Readme.md
```

## Использование

На данном этапе разработки, пока сервис не работает в контейнере, вам потребуется иметь на своем устройстве **Golang версии 1.19**.

Требуется перейти в папку **/app/cmd/** и запустить файл **main.go**. После чего в консоли вы увидите результат.