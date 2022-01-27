**Learning Resource: [Yaml Tutorial | Learn YAML in 18 mins (TechWorld with Nana)](https://youtu.be/1uFVr15xDGg)**

# YAML
- Very popular for writing configuration files.
- Widely used format.
- Used for writing configuration files for many different DevOps tools and applications.

## What is YAML?
- YAML is a data serialization language just like XML or JSON

- #### What is a serialization language?
  - Applications written in different technologies and languages can transfer data using a common, agreed on and standard format
  - Most popular : YAML, JSON and XML

- YAML stands for YAML Ain't Markup Language
- File extensions: `.yaml` or `.yml`

## Why learn YAML? YAML compared to JSON and XML
- YAML is more human readable and intuitive, great fit for wiriting configuration files for all those recent DevOps tools like Docker, K8s etc.
- #### YAML:
  - No special characters to define data structures
  - Data structure is defined using line seperation and space indentation
```
microservices:
  - app: user-authentication
    port: 8000
    version: 1.0
```
- #### XML
  - Data structured defined using `<>` tags
```
<microservices>
  <microservices>
    <app>user-authentication</app>
    <port>8000</port>
    <version>1.0</version>
  </microservices>
</microservices>
```
- #### JSON
  - Data structures defined using `{}` curly braces
  - **NOTE: YAML is a superset of JSON: any valid JSON file is also a valid YAML file**

```
{
  microservices: [
    {
      app: "user-authentication"
      port: 8000
      version: "1.0"
    }
  ]
}
```
## YAML Syntax

- ### key-value pairs
  - Here, `app: user-authentication` is a String, we don't need to include it in `""`. However if we're using special characters, we need to `app: "user-authentication \n"`

```
app: user-authentication
port: 8000
version: 1.7
```

- ### comments
  - We can write comments in YAML using `#`
  - To make out file more readable and human understandable
```
# comment here
app: user-authentication
port: 8000
# comment here
version: 1.7
```

- ### objects
  - We can group key-value pairs inside of an object
  - Here, we have a object `microservice`
  - Take care of the spaces and identation (use a YAML validator if needed)
```
microservice:
  app: user-authentication
  port: 8000
  version: 1.7
```

- ### lists
  - If we have multiple "microservices", we can create a list using `-`
```
microservice:
  - app: user-authentication
    port: 8000
    version: 1.7
  - app: user-data
    port: 9000
    version:                                # list
      - 1.0
      - 2.0
      - 3.0
  - app: user-name
    port: 7000
    version: ["1.0", 2.0, 3.0, "latest"]    # list
    
microservices:                              # list
  - user-authentication
  - user-data
  - user-name
```
- ### booleans
  - We can use `true/false`, `yes/no` and `on/off` for boolean data.

```
microservice:
  app: user-authentication
  port: 8000
  version: 1.7
  deployed1: true # boolean true/false
  deployed2: yes  # boolean yes/no
  deployed3: on   # boolean on/off
```

- ### multi-line Strings
```
multilineString: |
  this is a multiline string
  and this is the next line
  next line
  etc
  
singlelineString: >
  this is a single line string,
  that should be all in same line.
  some other stuff
```

- ### environment variables
  - We can access environment variables in YAML using `$<name of env varibales>`
```
# readiness probe
command:
  - /bin/sh
  - -ec
  - >-
    mysql  -h 127.0.0.1 -u root -p$MYSQL_ROOT_PASSWORD -e 'SELECT 1'
```

- ### placeholders
  - We can use placeholders in YAML using `{{<value>}}` double curly braces
  - For eg: we need this in Helm and Ansible where placeholder value gets replaced using template generator

- ### multiple YAML documents
  - We can seprate different YAML using `---`
  - very handy in K8s where we have multiple components in one service and we want to group them in single yaml file
```
apiVersion: v1
kind: configMap
metadata
  name: config-file
  
---

apiVersion: v1
kind: configMap
metadata
  name: config-file
```

## Real K8s YAML Configuration Example

- ### Pod Configuration
  - key-value pairs
  - metadata = object
  - labels = object
  - spec = object
  - containers = list of objects
  - ports = list
  - volumeMounts = list of objects
  - 
```
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  container:
  - name: nginx-container
    image: nginx
    ports:
    - containerPort: 80
    volumeMounts:
    - name: nginx
      mountPath: /usr/nginx/html
  - name: sidecar-container
    image: curlimages/curl
    command: [" /bin/sh"]
    args: ["-c", "echo Hello from the sidecar container; sleep 300"]
```

## YAML and JSON in Kubernetes
- **We can write Kubernetes configuration files in both YAML and JSON format**
- YAML is more commonly used as its cleaner and more readable

