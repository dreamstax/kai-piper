# Kai-piper
A no frills pipeline orchestrator

## Description
Kai-piper enables clients to easily register and execute synchronous pipelines. You probably don't need a DAG. It is expected that steps can be signaled by way of some networking process (typically http) to perform the actual work. By separating the orchestration of the pipeline and the actual work being done in a pipeline we believe we can drastically simplify the overall configuration and management of today's pipelines. Also kai-piper is extremely lightweight and will probably scale really well but we haven't tested that yet ;)

## Kai controller
Although you could technically run kai-piper as a standalone service it's expected use-case is for it to be leveraged via the [Kai controller](https://github.com/dreamstax/kai). By using the controller you can automate the deployment and registration of your steps within a pipeline in a Kubernetes native way. 


