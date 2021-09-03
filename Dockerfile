FROM debian:10-slim
COPY install-slu /
ENTRYPOINT [ "/install-slu" ]
