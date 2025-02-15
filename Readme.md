# Разворачивание проекта

## Установка зависимостей

```
go get -u github.com/beego/beego/v2
go get -u github.com/beego/bee
go get -u github.com/lib/pq
```

### Cоздание нового проекта

```
bee api название проекта
```

### Генерация роутинга

```
bee generate routers
```

### запуск с генерацией swagger

```
bee run -gendoc=true -downdoc=true
```

url swagger после запуска
http://localhost:8080/swagger/

### Описание диррикторий

* conf 				дирриктория с конфигурационным файлом
* routers				содержит структуру HTTP путей
* controllers			контроллер, содержит код для автоматической генерации документации в формате swagger и описание точек входа
* swagger				дерриктория содержащая swagger
* tests               дирриктория с тестами, пока не трогаем
* models              дирриктория с моделами, описание структуры полей БД и обработка входящих запросов