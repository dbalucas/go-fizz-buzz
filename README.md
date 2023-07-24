# docker-gs-ping

A simple Go server/microservice example for [Docker's Go Language Guide](https://docs.docker.com/language/golang/).

Notable features:

* Includes a [multi-stage `Dockerfile`](https://github.com/olliefr/docker-gs-ping/blob/main/Dockerfile.multistage).
* Has a CI pipeline using GitHub Actions to run tests.
(* Has a CD pipeline using GitHub Actions to publish to Docker Hub.)

## Contributing

This was written for an _introduction_ section of the fizz-buzz and as such it favours clarity over robustness. 

Thus, feedback is welcome, but please no nits or pedantry. Ain't nobody got time for that 🙃

## License

[Apache-2.0 License](LICENSE)
