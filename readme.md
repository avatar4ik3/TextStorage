# Запуск

- установить нужный коммит.
  - на коммит на котором все работет
    ```console
    git reset --hard a05f73f5a3667e19a0d3908f93e05a6037348e1f
    ```
  - либо коммит на понедельник 19 декабря
    ```console
    git reset --hard 991995ae133975d72edadbca4ba413e943b32376
    ```
- отредактировать .env файл под нужные порты
- выполнить

```console
docker-compose up
```
