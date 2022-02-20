# logging

- In case of containers, the containers collect all the logs and your can check that by using `docker logs <image-id>`
- Docker picks out all the logs from STDOUT and STDERR
- By default docker writes the log file in JSON format
- We can verify this by the command `docker system info`

## Why Do we need logs?

In the microservices architecture we have many units and all these communicate with each other and generate logs.

- Logging can be important because of several regions:
    - Compliance
    - Security
    - Debugging - To figure out what went wrong if something breaks
    

## How to applications log data?

- One common way is to write that into a file
- log directly into log database e.g. elastic

## FluentD

- Fluentd is the unified logging layer
- fluentd gets deployed into the cluster and it collects the log from all the applications. 
- Now each applications can have log of different format (JSON or any other format). 
- Fluentd collects data in the format in which the application writes and then it present to you in a unified format.
- You can enrich your log data with fluentd.
- In most cases our aim is to visualize the log data.
- Fluentd can send any data from any data source to any destination e.g. kafka, elastic search,

## How to configure fluentd?

- You must install fluentd as daemonSet. DaemonSet is a component that runs on each kubernetes node.
- Fluentd configuration file
- First you define the source from which fluentd will collect the logs
- Now you define how the data will be processed
- You now enrich the data with record transformers to structure the data.
- Now at last you define the output.
- FluentD have tags feature to filter the data easily.
- Fluentd saves data in a hard drive so if somehow fluentd goes down and restarts then it will continue from where it left.
- Fluentd also have the retry logic and if something goes wrong at the destination site then fluentd will keep on trying.
