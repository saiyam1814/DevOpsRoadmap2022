
```
The key difference between Declarative pipeline and Scripted pipeline would be with respect to their syntaxes and their flexibility. 

Declarative pipeline is a relatively new feature that supports the pipeline as code concept. It makes the pipeline code easier to read and write. This code is written in a Jenkinsfile which can be checked into a source control management system such as Git.

Whereas, the scripted pipeline is a traditional way of writing the code. In this pipeline, the Jenkinsfile is written on the Jenkins UI instance. 

Though both these pipelines are based on the groovy DSL, the scripted pipeline uses stricter groovy based syntaxes because it was the first pipeline to be built on the groovy foundation. Since this Groovy script was not typically desirable to all the users, the declarative pipeline was introduced to offer a simpler and more optioned Groovy syntax.

The declarative pipeline is defined within a block labelled ‘pipeline’ whereas the scripted pipeline is defined within a ‘node’. 

Structure and syntax of the Declarative pipeline:

The Agent is where the whole pipeline runs. Example, Docker. The Agent has following parameters:

any – Which mean the whole pipeline will run on any available agent.

none – Which mean all the stages under the block will have to declared with agent separately.

label –  this is just a label for the Jenkins environment

docker –  this is to run the pipeline in Docker environment.

## **The Declarative pipeline code will looks like this**:
````
```
pipeline {
  agent { label 'node-1' }
  stages {
    stage('Source') {
      steps {
        git 'https://github.com/digitalvarys/jenkins-tutorials.git''
      }
    }
    stage('Compile') {
      tools {
        gradle 'gradle4'
      }
      steps {
        sh 'gradle clean compileJava test'
      }
    }
  }
}
```
````
## **Scripted Pipeline**:

The scripted pipeline is a traditional way of writing the Jenkins pipeline as code. Ideally, Scripted pipeline is written in Jenkins file on web UI of Jenkins.

 Unlike Declarative pipeline, the scripted pipeline strictly uses groovy based syntax. Since this, The scripted pipeline provides huge control over the script and can manipulate the flow of script extensively. 

This helps developers to develop advance and complex pipeline as code.

Structure and syntax of the scripted pipeline:

**Node Block**:

Node is the part of the Jenkins architecture where Node or agent node will run the part of the workload of the jobs and master node will handle the configuration of the job. So this will be defined in the first place as

node {
}
**Stage Block**:

Stage block can be a single stage or multiple as the task goes. And it may have common stages like

Cloning the code from SCM

Building the project

Running the Unit Test cases

Deploying the code

Other functional and performance tests.

So the stages can be written as mentioned below:

stage {
}
So, Together, the scripted pipeline can be written as mentioned below.
````
```
node ('node-1') {
  stage('Source') {
    git 'https://github.com/digitalvarys/jenkins-tutorials.git''
  }
  stage('Compile') {
    def gradle_home = tool 'gradle4'
    sh "'${gradle_home}/bin/gradle' clean compileJava test"
  }
}
```
````
Declarative Pipeline encourages a declarative programming model, whereas Scripted Pipelines follow a more imperative programming model. 

Declarative type imposes limitations to the user with a more strict and pre-defined structure, which would be ideal  for simpler continuous delivery pipelines. 

Scripted type has very few limitations that to with respect to structure and syntax that tend to be defined by Groovy,thus making it ideal for users with more complex requirements. 
```
