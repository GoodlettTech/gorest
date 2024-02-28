FROM mcr.microsoft.com/devcontainers/go:1-1.22-bullseye

RUN curl -fsSL https://deb.nodesource.com/setup_21.x | bash -
RUN apt-get install -y nodejs
USER vscode
RUN go install github.com/cosmtrek/air@latest
USER root
