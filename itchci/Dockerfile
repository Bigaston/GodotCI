FROM debian:bookworm

# Install required libraries
RUN apt update
RUN apt install -y curl zip

# Install butler
RUN curl -L -o butler.zip https://broth.itch.ovh/butler/linux-amd64/LATEST/archive/default
RUN unzip butler.zip
RUN mv butler /bin/butler
RUN chmod +x /bin/butler

# Clean
RUN rm butler.zip
RUN rm -rf /var/lib/apt/lists/*

# Check version
RUN butler -V