FROM gcr.io/distroless/static-debian11:nonroot
ENTRYPOINT ["/baton-panorama"]
COPY baton-panorama /