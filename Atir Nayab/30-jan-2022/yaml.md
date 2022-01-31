# Yaml

## why yaml?
- configuration files all written in YAML.
- widely used format
- for different DevOps tools and applications

## what is YAML?
- YAML is a data serialization language
- what is data serialization 
  - standard format to transfer data
- YAML Ain't Markup Language
- file extension: .yaml, .yml

## YAML Format compared to others
- human readable and intuitive

## Diff between YAML, XML, JSON
```yaml
calling-birds:
   - huey
   - dewey
   - louie
   - fred
```
```xml
<calling-birds>huey</calling-birds>
<calling-birds>dewey</calling-birds>
<calling-birds>louie</calling-birds>
<calling-birds>fred</calling-birds>
```
```json
{
   "calling-birds": [
      "huey",
      "dewey",
      "louie",
      "fred"
   ]
}
```
*YAML is superset of JSON any valid JSON file is also a valid YAML file*

## syntax

- key-value pair
```yml
app: user-authentication
port: 9000
version: 1.7
```
- comment
```yaml
# this is a comment
```
- objects
```yml
microservice: 
  app: test
  port: 9000
  version: 1.7
```
*for safety use yaml validator*

- list
```yml
microservice: 
  - app: test
    port: 7000
```

- boolean
```yml
delployed: true
delployed: false
delployed: yes
delployed: no
delployed: on
delployed: off
```

- more list
```yaml
microservices:
  - app: test
    port: 9000
  - app: test2
    port: 10000
app-name: 
  - shopping
  - user-auth
versions: [1.9, "2.2", 1.2]
```

- multiline string
```yml
# all string are considered single line
multilinestring: 
  this is a multiline string
  and this is the next line
  again next line
# multiline string
multilinestring: |
  this is a multiline string
  and this is the next line
  again next line
# all string are considered single line
multilinestring: >
  this is a multiline string
  and this is the next line
  again next line
```

- environmental variables
```yml
app: $APP_NAME
```

-placeholders
```yml
# value inside brackets is replaced by brackets
app: {{ testing }}
```

- multiple yaml documents 
```yml
app: app one cofig

---
app: app two config
```