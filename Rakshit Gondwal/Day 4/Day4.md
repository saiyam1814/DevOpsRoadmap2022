Day 4-

Today I learned more about kubernetes. 

Here's a differnce between monolithic architecture and microservices-

# What is Monolithic architecture?
Let's suppose we have a MERN stack application. Now, this application is written in one single code base and is deployed as a single unit, it would be referred to as an application built on monolithic architecture. 

## Advantages- 
Though there are several advantages of monolithic architecture such as being the whole codebase in one file, it is easier to deploy these applications or easy debugging,

## Disadvantages- 
Now, let's say you want to update the frontend of the application then you will have to update the whole code base again and then deploy it as a whole again. This makes the updates restrictive and time-consuming. 


there can be many disadvantages of it too like Slower development speed, Scalability issues, A small change to the codebase would require in changing of the whole codebase, etc.

# What are Microservices?
Let's suppose, the same MERN stack application is deployed in small different parts and independently, like frontend is deployed in a different container, backend is deployed in a different container. This way of deploying applications into smaller independent services is known as Microservices.

## Advantages- 
There are a lot of advantages of microservices such as-
- Makes the applications easily scalable.
- Makes the applications highly maintainable and testable.
- More frequent deployments.
- High reliability. 

## Disadvantages- 
- Debugging challenges. 
- More complex.