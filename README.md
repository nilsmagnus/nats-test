# nats-test


start nats-streaming in docker:

    $ docker run -p 4223:4223 -p 8223:8223 nats-streaming -p 4223 -m 8223


run the client :

    $ go run main.go