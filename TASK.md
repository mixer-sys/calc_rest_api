1. Реализуйте HTTP-сервер на **`net/http`**, который принимает числа через POST-запрос и возвращает их сумму.
![alt text](image.png)
2. Перепишите сервер на [Echo](https://echo.labstack.com/guide/routing/).
3. Организуйте проект по go standart project layout.
4. Сохраняйте результаты вычислений в памяти (используйте самописную потокобезопасную мапу с `sync.Mutex`, дополнительно почитайте про `sync.Map` и когда она используется).
5. Расширьте функциональность калькулятора, добавив ручку для умножения.
6. Добавьте базовое разделение запросов от пользователей: пользователь отправляет вместе с числами токен (какую-либо строку, лучше UUID), результаты сохраняются по токену (в памяти приложения).
7. Задокументируйте API через Swagger (используйте [swaggo](https://github.com/swaggo/swag)).
8. Напишите Makefile для запуска генерации документации.