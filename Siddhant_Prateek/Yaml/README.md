# YAML 

Yaml (Yaml ain't markup language), earlier it was known as Yet another markup language. It is data format, used to exchange data. It is similar to XML and JSON(Javascript Object Notation).

- It is used to store information about some configuratio, logs, caches.
- In Yaml, you can store only data not commands.

### Data Serialisation

It is process of converting data objects into a complex data structure or stream of bytes.

examples of data-serialisation languages are: JSON, XML , YAML

![](https://i.imgur.com/PhKs0mc.png)

### Data De-Serialisation

It is process of converting stream of bytes or data back to data object.


### It's Benefits

- It's simple and easy to read
- It has strict syntax.
- Easily convertable to JSON, XML
- Widely  used in various programming language.
- More powerful when representing complex data.
- Parsing( Reading the data ) is Easy.

Example of `.yaml` file
```yaml=

"apple" : " its a red fruit"
1 : "gaustovo"
---
# lists

- apple
- mango
- banana
---
# block stype

cities:
 - new york
 - madrid
 - kolkata
--- 
cities: [new york, madrid, kolkata]
...
# ended
```

example for storing in json from format.
```json=
{
    "cities": [
      "new york",
      "madrid",
      "kolkata"
    ]
  }
```

```yaml=
# to Write a single line in multiple 

message: >
    hey,
    this statement
    will be read in 
    single line.
```


### Data types in YAML
```yaml 

"apple" : " its a red fruit"
1 : "gaustovo"
---
# lists

- apple
- mango
- banana
---
# block stype

cities:
 - new york
 - madrid
 - kolkata
--- 
cities: [new york, madrid, kolkata]

---
bio : |
 My name
 is Siddhant Prateek
 Mahanayak

message: >
    hey,
    this statement
    will be read in 
    single line.
---
# data types
booleanValue: No
boolean2: False
boolean3: false
boolean4: FALSE

boolean_Value: Yes
boolean_Value2: True
boolean_Value3: !!bool TRUE
boolean_Value4: yes
---
# specific type of data

zero: !!int 0
positiveNum: !!int 45
negativeNum: !!int -45
binaryNum: !!int 0xb11001
octalNum: !!int 06574
hexa: !!int 0x45
exponential number: 4.77E36
---
# floating point numbers
marks: !!float 80.4
infinite: !!float .inf
not a num: .nan
---
# null
name: !!null Null # NULL null ~
~: Null key
...
# ended
```

### Advance Data types

```yaml 
student: !!seq 
 - marks
 - name
 - roll 

# similarly

cities: [new york, madrid, kolkata]

# Some of the keys of the sequence will be empty
# sparse seq.
sparse seq:
 - hey
 - what
 - 
 - null
 - cool 
---
# nested sequence
- 
 - mango
 - apple 
 - grapes
- 
 - name
 - roll
 - date
---
#  key: value pairs are called maps !!map

name: Siddhant Prateek Mahanayak
role: 
 age: 12
 status: active

---
# similarly
name: Siddhant Prateek Mahanayak
role: {age: 12, statue: active}

# pairs: keys may have duplicate values !!pairs

pair example: !!pairs
 - job: student 
 - class: 2
# this will be an array of hashtable

# similarly
pair example2: !!pairs [job: student, class: 2]

# '!!set' will allow you to have unique values

names: !!set
 ? andrew
 ? rodri
 ? jack

#  dictionary !!omap

people: !!omap 
 - dex: 
   - name: Siddhant
   - age: 21
 - lex: 
   - name: rex 

# reusing some properties we use anchors

likes: &like
 - fav fruit: grapes
 - role: student

person1: 
 name: Siddhant
 <<: **like

person2: 
 name: Prateek
 <<: **like
 
# updated
person3: 
 name: Max
 <<: **like
 role: developer
```