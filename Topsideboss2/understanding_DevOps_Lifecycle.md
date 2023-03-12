# Introduction to DevOps Lifecycle

## Development

Taking a brand new example of an application, to start with we have nothing created, maybe as a developer, you have to discuss with your client or end user the requirements and come up with some sort of plan or requirements for your application. We then need to create from the requirements our brand new application.

In regards to tools, there is no real requirement here other than choosing your IDE and the programming language you wish to use to write your application.

As a DevOps engineer, remember you are probably not the one creating this plan or coding the application for the end user, this will be a skilled developer. 

But it also wouldn't hurt for you to be able to read some of the code so that you can make the best infrastructure decisions moving forward for your application.

This application will be maintained using a version control system because the application can be written in any language.
It is also likely that it will not be one developer working on this project although this could be the case even so best practices would require a code repository to store and collaborate on the code, this could be private or public and could be hosted or privately deployed generally speaking you would hear the likes of **Github** or **Gitlab** being used as a code repository. 

## Testing

At this stage, we have our requirements and we have our application being developed. But we need to make sure we are testing our code in all the different environments that we have available to us specifically maybe to the programming language chosen. This phase enables Quaity Assurance to test for bugs, more frequently we see containers being used for simulating the test for environments which overall can improve on cost overheads of physical or cloud infrastructure.

This phase is also going to be automated as part of the next area which is Continuous Integration. The ability to automate this testing vs 10s, 100s or even 1000s of quality assurance engineers having to do this manually speaks for itself, these engineers can focus on something else within the stack to ensure you are moving faster and developing more functionality vs testing bugs and software which tends to be the hold up on most traditional software releases that use a waterfall methodology.

## Integration

Quite importantly Integration is at the middle of the DevOps Lifecycle. It is the practice in which developers require to commit changes to the source code more frequently. This could be on a daily or weekly basis. 
With every commit, your application can go through the automated testing phases and this allows for early detection of issues or bugs before the next phase.

## Deployment

Ok so we have our application built and tested against the requirements of our user and we now need to go ahead and deploy this application into application for our end users to consume.

This is the stage where the code is deployed to the production servers, now this is where things get extremely interesting. Because different applications require different possibly hardware or configurations. This is where **Application Configuration Management** and **Infrastructure as Code** could play a key part in your DevOps lifecycle. It might be that your application is **Cointainerised** but also available to run on a virtual machine. This then also leads us onto platforms like **Kubernetes** which would be orchestrating these containers and making sure you have the desired state available to your end users.

## Monitoring

Things are moving fast here and we have our application that we are continuously updating with new features and functionality and we have our testing making sure no gremlins are being found. We have the application running in our environment that can be continually keeping the required configuration and performance.
But now we need to be sure that our end users are getting the experience they expire. Here we need to make sure that our Application Performance is continuously being monitored, this phase is going to allow your developers to make better decisions about enhancements to the application in future releases to better serve the end users.
Monitoring is also where we are going to capture that feedbacl wheel about the features that have been implemented and how the end users would like to make these better for them.
Reliability is a key factor here as well, at the end of the day we want our Application to be available all the time it is required. This then leads to other observability, seecurity and data management areas that should be continuously monitored and feedback can always be used to better enhance, update and release the application continuously.

In the words of @MichaelCade

> I think it is also a good time to bring up the "DevOps Engineer" mentioned above, albeit there are many DevOps Engineer positions in the wild that people hold, this is not the ideal way of positioning the process of DevOps. What I mean is from speaking to others in the community the title of DevOps Engineer should not be the goal for anyone because really any position should be adopting DevOps processes and culture explained here. DevOps should be used in many different positions such as Cloud-Native engineer/architect, virtualization admin, cloud architect/engineer, and infrastructure admin. This is to name a few but the reason for using DevOps Engineer above was really to highlight the scope of the process used by any of the above positions and more.





