# Camellia-ui

Алгоритм необходимо реализовать в оконном или Web-приложении.

(Де-)Шифрование должно производиться асинхронно (не допускается отвисание UI-потока) и многопоточно(для увеличения скорости, если это возможно).

Необходимо реализовать режимы шифрования: 
- электронной кодовой книги(ECB),
- сцепления блоков шифротекста(CBC),
- обратной связи по шифротексту(CFB),
- обратной связи по выходу(OFB).

Предусмотреть возможности ввода ключа шифрования и вектора инициализации для режимов шифрования из файла.
Приложение должно уметь шифровать любые файлы, также необходимо реализовать отображение прогресса операции шифрованияи(опционально)
отмену операции шифрования по запросу пользователя.


## Сборка
генерация .go файла из .ui
```shell script
 goqtuic -go-test-file ui/camellia.ui -go-ui-dir ui/
```

## Скриншоты

![Lab3](https://github.com/elizarpif/camelia/blob/develop/screenshots/6.png)

![Lab3](https://github.com/elizarpif/camelia/blob/develop/screenshots/7.png)