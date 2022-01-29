## Completed [Yaml Tutorial](https://www.youtube.com/watch?v=1uFVr15xDGg)

**YAML** is a data serialization language.

**What is a serialization language?**

Application written with different technologies languages etch which have different data structures can tranfer data to each other using standard format.
The most popular standard formats are YAML, JSON and XML.

**YAML** stands for YAML Ain't Markup Language


File extenstion: `.yaml` `.yml`


YAML Format compare to others is human readable and intuitive. <br/>

![image](https://user-images.githubusercontent.com/74575612/150760957-6270a9d9-5384-497d-ab46-217fa8eb987e.png)

![image](https://user-images.githubusercontent.com/74575612/150761029-7264c7de-542e-4d32-8b2b-5b51efceb1d3.png)


YAML is superset of JSON: any valid JSON file is also a valid YAML file.

YAML Use Case Files:
- Docker Compose Files
- Ansible
- Kubernetes

**Basic Syntax of YAML**
- Key-Value Pair <br/>
  Example:
  ```
  app: user-authentication
  port: 9000
  version: 1.7
  ```
- Objects <br/>
  Example:
  ```
  microservice:
    app: user-authentication
    port: 9000
    version: 1.7
  ```
 - Lists <br/>
   Example:
   ```
   microservice:
    - app: user-authentication
    port: 9000
    version: 1.7
   ```
 - Boolean <br/>
   Example
   ```
    microservice:
    - app: user-authentication
      port: 9000
      version: 1.7
      deployed: yes
    ```
  - More About Lists <br/>
    Example
    ```
    microservice:
    - app: user-authentication
      port: 9000
    - app: shopping-cart
      port: 9002
      versions: ["1.9", "2.0", "2.1"]
     ```
   - Multi Line Strings <br/>
     Example-1
     ```
     multilineString: |
       this is a multiline string,
       some other stuff
     ```
     Example-2
     ```
     multilineString: >
       this is a single line string,
       that should be on one line
       some other stuff
     ```
   - Environment Variables <br/>
     Example
     ```
     command:
       - /bin/sh
       - -ec
       - >-
         mysql -h 127.0.0.1 -u root -p$MYSQL_ROOT_PASSWORD -e 'SELECT 1'
     ```
   - Placeholders <br/>
     Example:
     ```
     apiVersion: v1
     kind: Service
     metadata:
       name: {{ .Values.service.name }}
     spec:
       selector: 
         app: {{ .Values.service.app }}
       ports:
         - protocol: TCP
           port: {{ .Values.service.port }}
           targetPort: {{ .Values.service.targetPort }}
     ```
   - Multiple YAML Docs <br/>
     Example:
     ```
     apiVersion: v1
     kind: ConfigMap
     metatda:
       name: mosquito-config-file
     data:
       mosquito.conf: |
         log_dest stdout
         log_type all
         log_timestamp true
         listener 9001
      ---
      apiVersion: v1
      kind: Secret
      metadata: 
        name: mosquito-secret-file
      type: Opaque
      ```
    
   
