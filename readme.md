# Программа для вывода дерева каталогов и файлов

## Описание работы

Программа при запуске принимает на вход два аргумента: путь к файлу и флаг, определяющий отображение файлов при выводе дерева.

*Например:*  

При указании флага -f будет осуществлен вывод дерева с каталогами и файлами.
```
go run main.go testdata -f 
```
*Вывод:*
```
├───project
│       ├───file.txt (19b)
│       └───gopher.png (70372b)
├───static
│       ├───a_lorem
│       │       ├───dolor.txt (empty)
│       │       ├───gopher.png (70372b)
│       │       └───ipsum
│       │               └───gopher.png (70372b)
│       ├───css
│       │       └───body.css (28b)
│       ├───empty.txt (empty)
│       ├───html
│       │       └───index.html (57b)
│       ├───js
│       │       └───site.js (10b)
│       └───z_lorem
│               ├───dolor.txt (empty)
│               ├───gopher.png (70372b)
│               └───ipsum
│                       └───gopher.png (70372b)
├───zline
│       ├───empty.txt (empty)
│       └───lorem
│               ├───dolor.txt (empty)
│               ├───gopher.png (70372b)
│               └───ipsum
│                       └───gopher.png (70372b)
└───zzfile.txt (empty)
```

Если флаг не указывать, то дерево будет состоять только из каталогов.
```
go run main.go testdata
```
*Вывод в этом случае:*
```
├───project
├───static
│       ├───a_lorem
│       │       └───ipsum
│       ├───css
│       ├───html
│       ├───js
│       └───z_lorem
│               └───ipsum
└───zline
        └───lorem
                └───ipsum
```

## Запуск скрипта

```

```