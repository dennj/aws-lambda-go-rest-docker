# Build:
`docker build --platform linux/amd64 -t docker-image:test .`


# Run container:
`docker run -d -p 9000:9000 --name lambda-tolki docker-image:test`

# Stop container:
- `docker ps`
- `docker stop lambda-tolki`
- `docker rm lambda-tolki`

# Test:
- `curl -X GET "http://localhost:9000/v1/getinfo/12345"`
- `curl -X POST "http://localhost:9000/v1/set/12345/of/67890"`