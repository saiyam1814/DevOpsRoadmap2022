- When I started with kubernetes, in the starting months I have only learnt about pods and deployments mainly. 
- Today I tried to learn about jobs. 
- To set up jobs in kubernetes I tried creating a container image and using that image I tried to run a job in my kubernetes cluster. 

- I tried writing this simple go program that will run for 10 seconds. 
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting job at time 😐 : ", time.Now().Format("01-02-2006 Monday 15:04:05"))
	fmt.Println("Doing my Job...")
	time.Sleep(time.Second * 7)
	fmt.Println("Just few seconds more to finish 🙃")
	time.Sleep(time.Second * 3)
	fmt.Println("Finishing job at time 😎 : ", time.Now().Format("01-02-2006 Monday 15:04:05"))
}
```


- I created a Dockerfile to containerize the above program 
```Dockerfile
FROM golang:alpine
COPY . . 
CMD ["go", "run", "main.go"]
```

- Create the container image using `podman build -t kranurag7/simplejob .` 
- push the image to container registry using `podman push kranurag7/simplejob` 