# Zeet Golang test

  

## Requirement

- Fix broken repo so it can be deployed on Zeet

- Add a feature to the now working project

## Specifications

1.  Fork the Github repo, and deploy it as a project on the Zeet team.  
    
2.  Fix anything that needs to be fixed, so that by the end you have an endpoint that you can open in your browser that responds with `{"page": "home"}`
    
3.  Make it so that if I go to `{url}/hello/johnny`, it responds with `{"page": "hello", "hello": "johnny"}`
    
4.  for bonus points, add some UI so it's not just a JSON response (not required)

## Solution
To fix this project, I had to change a couple of things. 
1. `main.go` was only listening on `127.0.0.1` so when I created the Docker image, I'd get an Empty Reply. Had to fix this by listening to all IPs (`0.0.0.0`)
2. The Dockerfile was giving me many issues. This could have to do with the Golang version and/or the the Dockerfile itself. Fixed this by rewriting the Dockerfile. 
3. To implement the "hello functionality", we simply add a parameter to the route and send it to the same handler. gin-gonic defines [two types of parameters](https://chenyitian.gitbooks.io/gin-web-framework/content/docs/8.html) in path. With using  __*param__ to extract the value, we limit the possibility of further route branching as it consumes all forward slashes after {URL}/hello/. 
