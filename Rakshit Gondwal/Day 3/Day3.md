Day 3- 

Today I revised some of my concepts about docker. Here's what I learned - 


# Problems before docker.
Let's suppose you have made a web applications which is based on MERN stack.
React for frontend.
Node.js and express for backend.
MongoDB for the database.

You must be knowing that these applications are based on the versions of their tech stacks. 
Like let's say, 

- You built this website on node.js's latest version in the month of January and you've deployed it on the server and the code is public too.
- Time passed, and now I want to run that website in my own local environment by cloning the repo and all that stuff, but it's the month of august and I have the lastest version Node.js as of august.
- But the application is running on January's vresion of node.js. Now, I'll have to install the old version of node.js so that I would be able to run this application in my local environment.

Thus the main problems which arise are- 
- The local hosting is dependent on the OS.
- The local hosting is dependent on the versions of tech stacks.
- Everytime any service releases updates, you might need to recheck the compatibility. 

# How did docker solved this problems?
With all these problems in scene, docker came and solved all these problems. Docker basically provides us with a local isolated environment in which we can run the applications without having to worry about the dependencies, service updates, OS and all other stuff.

We can build new applications and upload them to docker from where anyone would be able to download our application on their local system(the dependencies gets installed too automatically) and can run them without having to worry about any other stuff.