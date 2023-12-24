FROM golang:1.21.5-bullseye

# Install Protoc
RUN apt update \
  && apt install -y protobuf-compiler

# Install Go Tools
RUN go install -v golang.org/x/tools/gopls@latest

# Install Node.js (required for Salesforce CLI)
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
  apt-get install -y nodejs

# Install Salesforce CLI
RUN npm install -g sfdx-cli

# Install Yarn
RUN npm install -g yarn