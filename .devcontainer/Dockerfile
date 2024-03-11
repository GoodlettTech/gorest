FROM mcr.microsoft.com/devcontainers/go:1-1.22-bullseye

EXPOSE 3000
EXPOSE 3001

RUN curl -fsSL https://deb.nodesource.com/setup_21.x | bash -
RUN apt-get install -y nodejs
USER vscode
RUN go install github.com/cosmtrek/air@latest
USER root
