<b>Snippetbox</b> 🚀

Веб-приложение для управления заметками на Go и MySQL</b>

<b>Описание проекта</b>

Snippetbox — это простое, но мощное веб-приложение, созданное для изучения и демонстрации возможностей бэкенд-разработки на языке <strong>Go (Golang)</strong> с использованием <strong>MySQL</strong> в качестве базы данных.
Приложение позволяет пользователям:
<ul>
  <li>📝 <strong>Создавать</strong> заметки с указанием заголовка, содержимого и срока действия.</li>
  <li>🔍 <strong>Просматривать</strong> список последних заметок.</li>
  <li>📄 <strong>Изучать</strong> детали конкретной заметки.</li>
  <li>🗑️ <strong>Удалять</strong> заметки по их ID.</li>
</ul>

<strong>Основные функции</strong>

🏠 Просмотр заметок

На главной странице отображаются последние 10 заметок. Каждая заметка содержит:
<ul>
  <li>Заголовок</li>
  <li>Содержимое (краткое описание)</li>
  <li>Дату создания</li>
  <li>Срок действия</li>
</ul>

✏️ Создание заметок

Пользователь может легко создать новую заметку, заполнив простую форму. После успешного создания заметки, пользователь автоматически перенаправляется на страницу с её деталями.

🔎 Просмотр деталей заметки

Каждая заметка имеет свою уникальную страницу, где можно увидеть полное содержимое.

🗑️ Удаление заметок
Пользователь может удалить заметку, указав её ID. После удаления пользователь перенаправляется на главную страницу.

<b>Технологии</b> 🛠️

Язык программирования: Go (Golang)

База данных: MySQL

Шаблонизация: Встроенный пакет html/template

Маршрутизация: Стандартный пакет net/http

Логирование: Встроенный пакет log

<b>Как запустить проект</b> 🚀
<ol>
  <li>Установите Go и MySQL
Убедитесь, что на вашем компьютере установлены:

Go

MySQL
</li>
<li>Создайте базу данных
Подключитесь к MySQL и выполните следующие команды:

CREATE DATABASE snippetbox;

USE snippetbox;

CREATE TABLE snippets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);</li>
<li>Настройте подключение к базе данных
В файле main.go укажите параметры подключения к MySQL:

dsn := "web:pass@/snippetbox?parseTime=true"</li>
<li>Запустите приложение
Перейдите в директорию проекта и выполните команду:

go run cmd/web/main.go</li>
<li>Откройте в браузере
Перейдите по адресу: http://localhost:4000.</li>
</ol>








<b>Зависимости</b> 📦

MySQL Driver: github.com/go-sql-driver/mysql

Примеры использования 🖥️

<b>Главная страница</b>

Перейдите по адресу http://localhost:4000, чтобы увидеть список последних заметок.

<b>Создание заметки</b>

Перейдите по адресу http://localhost:4000/snippet/create, чтобы создать новую заметку.

<b>Просмотр заметки</b>

Перейдите по адресу http://localhost:4000/snippet?id=1, чтобы просмотреть заметку с ID = 1.

<b>Удаление заметки</b>

Перейдите по адресу http://localhost:4000/snippet/delete, чтобы удалить заметку по её ID.


👤 Автор
TREFLX

[![GitHub](https://img.shields.io/badge/GitHub-TREFLX-blue?style=flat&logo=github)](https://github.com/TREFLX)
[![Telegram](https://img.shields.io/badge/Telegram-Treflx-blue?style=flat&logo=telegram)](https://t.me/Treflx)

📷 Скриншоты
![image](https://github.com/user-attachments/assets/96306f50-48bf-4f9a-915a-4ea4a31699c2)
![image](https://github.com/user-attachments/assets/254ff4c0-cc0e-4a73-9e29-e3235810dfe2)


Ссылки 🔗

- [GitHub репозиторий](https://github.com/TREFLX/Snippetbox)  
  Перейдите в репозиторий проекта на GitHub.
- [Документация Go](https://golang.org/doc/)  
  Официальная документация по языку программирования Go.
- [Документация MySQL](https://dev.mysql.com/doc/)  
  Официальная документация по MySQL.
