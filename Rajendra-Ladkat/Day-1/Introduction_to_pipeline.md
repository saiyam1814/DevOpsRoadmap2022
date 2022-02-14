


**Jenkinsfile structure**

For Declarative Pipeline, the first line must be pipeline to open a pipeline block in which all the rest of the code is included. Note the following elements:
agent — Specifies where the Pipeline or a specific stage executes
Declarative Pipeline requires a global declaration of agent as the first line in the pipeline block. It is possible to use different agents for different stages but the global agent statement is always required. For now, we are using agent any, meaning that this Pipeline can run on any available agent.
stage — A conceptually distinct subset of the Pipeline, such as MyBuild, MyTest, or MyDeploy Note the two stages that we have created, each of which contains steps.
Used to present Pipeline status/progress.
Stage labels should be significant for your application.
steps — A series of distinct tasks inside a stage. You see here that the Print Message steps we created have been converted to standard echo commands.
Real world Jenkinsfiles are, of course, much more complex — they can also include Sections, Directives, Options and so forth — but will always have this same basic structure.
For full information about Pipeline syntax, see jenkins.io/doc/book/pipeline/syntax
