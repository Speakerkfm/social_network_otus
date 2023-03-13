# social_network_otus
### Установка

1. Склонировать проект
    ```sh
    git clone git@github.com:Speakerkfm/social_network_otus.git
    ```

2. Установить необходимые инструменты
    ```sh
    make install
    ```

3. Запустить бд в контейнере и накатить миграции
    ```sh
    make start-postgres
    make migration
    ```

4. Запустить API локально
    ```sh
    make run
    ```

### Документация
    Сваггер будет доступен по следующему URL'у
    http://127.0.0.1:84/docs

