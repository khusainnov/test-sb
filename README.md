# Тестовое задание для поступления в GoCloudCamp

___

### Описание

* Реализовать сервис, который позволит динамически управлять конфигурацией приложений. Доступ к сервису должен
  осуществляться с помощью API. Конфигурация может храниться в любом источнике данных, будь то файл на диске, либо база
  данных. Для удобства интеграции с сервисом может быть реализована клиентская библиотека.

### Структура конфига

* #### entity path: *sbercloud/internal/entity/upload.go*
```
{
    Service string            `json:"service"`
    Data    map[string]string `json:"data"`   
}
```

* #### Конфиг файл
```
{
  "service": "anyCfgName",
  "data": {
    "key1": "value323",
    "key2": "value2",
    "key3": "value3"
  }
}
```

* #### Endpoints
  **POST:** `/v1/config` – *добавление нового конфига 
  или изменение имеющего файла*

  **GET:** `/v1/config?service={filename}` – *получение данных 
    из конфига*

  **DELETE:** `/v1/config?service={filename}` – *удаление конфигурации*

### Результат
    
* **Хранение данных происходит в redis**

  
* **Создание конифга**

`curl -d "@data.json" -H "Content-Type: application/json" -X POST "http://localhost:8080/v1/config"`

**Response:**```{"code":200,"message":"Uploaded"} ```
* **Получение конфига**

`curl -X GET "http://localhost:8080/v1/config?service=anyCfgName"`

**Response:**```{"body":{"key1":"value323","key2":"value2","key3":"value3"},"code":200}```

* **Удаление конфига**

`curl -X DELETE "http://localhost:8080/v1/config?service=anyCfgName"`

**Response:**```{"code":200,"message":"Deleted"}```