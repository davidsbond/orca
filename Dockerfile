FROM gcr.io/distroless/static

COPY orca /usr/bin/orca

ENTRYPOINT ["/usr/bin/orca"]
