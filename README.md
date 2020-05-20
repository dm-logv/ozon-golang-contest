# Ozon GoLang School

Qualifying contest for the Ozon GoLang school


## A. Unique numbers
### Description 

```
A. Уникальное число

Ограничение времени	1 секунда
Ограничение памяти	64Mb
Ввод	стандартный ввод или input-201.txt
Вывод	стандартный вывод или input-201.a.txt

На вход программе подается большое количество целых чисел. Все числа, кроме одного, имеют пару, причем может быть несколько одинаковых пар. Найдите число без пары.


Формат ввода

stdin десятеричные числа по одному на каждой строке
Формат вывода
stdout десятеричное число

Пример

Ввод	Вывод
1       3
2
2
1
2
3
2
```

### Solution

```bash
unique_number$ cat input.txt | python3 unique_number.py
9
```


## B. Tags
### Description

```
B. Теги

Ограничение времени	1 секунда
Ограничение памяти	64Mb
Ввод	стандартный ввод или input.txt
Вывод	стандартный вывод или output.txt

В базе данных имеется таблица с товарами goods (id INTEGER, name TEXT), таблица с тегами tags (id INTEGER, name TEXT) и таблица связки товаров и тегов tags_goods (tag_id INTEGER, goods_id INTEGER, UNIQUE (tag_id, goods_id)).

Выведите id и названия всех товаров, которые имеют все возможные теги в этой базе.

БД - SQLite3. В качестве языка решения выберите make2.

Формат ввода
SQL-запрос.
```

### Solution

```bash
tags$ rm -f tags.db && sqlite3 -init tags_db.sql tags.db '.read tags.sql'
-- Loading resources from tags_db.sql
2|Good 2
8|Good 8
```

