FROM gcr.io/distroless/static-debian11:nonroot
ENTRYPOINT ["/baton-network-security"]
COPY baton-network-security /