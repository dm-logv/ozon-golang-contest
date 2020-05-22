# Ozon GoLang School

Qualifying contest for the Ozon GoLang school


## A. Unique numbers
### Description 

<details>
<summary>A. Уникальное число</summary>

```
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

</details>

### Solution

```bash
unique_number$ cat input.txt | python3 unique_number.py
9
```


## B. Tags
### Description

<details>
<summary>B. Теги</summary>

```
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

</details>


### Solution

```bash
tags$ rm -f tags.db && sqlite3 -init tags_db.sql tags.db '.read tags.sql'
-- Loading resources from tags_db.sql
2|Good 2
8|Good 8
```


## D. Sum of two numbers
### Description

<details>
<summary>D. Сложение чисел</summary>

```
Ограничение времени	1 секунда
Ограничение памяти	64Mb
Ввод	стандартный ввод
Вывод	стандартный вывод

Даны два числа неотрицательных числа A и B(числа могут содержать до 1000 цифр). Вам нужно вычислить их сумму.


Формат ввода

Первая строка ввода содержит числа A и B, разделенные пробелом


Формат вывода

Результат сложения двух чисел нужно вывести на отдельной строке.


Пример 1

Ввод	Вывод
1       2
3


Пример 2

Ввод	Вывод
199     1
200
```

</details>


### Solution

```bash
number_sum$ c++ number_sum.cpp && echo "9999999993 99999" | ./a.out
10000099992%
```

Tests:

```bash
number_sum$ python3 test_number_sum.py
```


## F. Find a sum
### Description

<details>
<summary>F. Сумма двух</summary>

```
Ограничение времени	1 секунда
Ограничение памяти	64Mb
Ввод	input.txt
Вывод	output.txt

Дано целое положительное число "target". Также дана последовательность из целых положительных чисел. Необходимо записать в выходной файл "1", если в последовательности есть два числа сумма, которых равна значению "target" или "0" если таких нет.


Формат ввода

5
1 7 3 4 7 9

Формат вывода

1


Примечания

Все числа используемы в задаче находятся в диапазоне 0 < N < 999999999

Название входной файл: input.txt
Название выходной файл: output.txt
```

</details>


### Solution

```bash
find_sum$ go run find_sum.go && cat output.txt
```

Some tests:

```bash
find_sum$ python3 test_find_sum.py
```
