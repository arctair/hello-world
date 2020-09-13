FROM debian
COPY bin/hello-world /bin/hello-world
ENTRYPOINT ["/bin/hello-world"]
