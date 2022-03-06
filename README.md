This projet is a backend service that generates random hellos from different languages.
The port is set to 9000.
# Endpoints

## /randomhellos
returns a random hello from a random language.
The response is a json object with the following fields:
* `hello`: the random hello
* `language`: the language of the hello
It will return it as a json object, as such:
``` json
{
  "language": "French",
  "hello": "bonjour"
}
```

# Docker
Run the following command to build the docker image:
`docker run -p 9000:9000 valentinrudloff/randomhellos`