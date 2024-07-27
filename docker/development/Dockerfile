FROM golang:bookworm as development

RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y \
    # to work inside the container
    locales \
    build-essential \
    cron \
    curl \
    wget \
    vim \
    sudo \
    iputils-ping \
    watch \
    tree \
    nano \
    less \
    software-properties-common \
    bash-completion \
    default-mysql-client \
    make \
    # VSCode container requirements
    net-tools \
    # shells
    bash \
    zsh \
    # other tools
    file \
    && rm -rf /var/lib/apt/lists/*

RUN sed -i -e 's/# en_US.UTF-8 UTF-8/en_US.UTF-8 UTF-8/' /etc/locale.gen \
    && dpkg-reconfigure --frontend=noninteractive locales

# Create new user with home directory, improve docker compatibility with UID/GID 1000,
# add user to sudo group, allow passwordless sudo, switch to that user
# and change directory to user home directory
RUN groupadd -g 1000 goalist \
    && useradd --no-log-init -r -m -u 1000 -g 1000 -G sudo goalist \
    && echo "goalist ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers

USER goalist
WORKDIR /home/goalist

# Uses "robbyrussell" theme (original Oh My Zsh theme), with no plugins
RUN sh -c "$(curl -L https://github.com/deluan/zsh-in-docker/releases/download/v1.1.5/zsh-in-docker.sh)" -- \
    -t robbyrussell \
    -p git

# Install Node via nvm
ENV NODE_VERSION_16=16.20.2
ENV NODE_VERSION=18.18.2
ENV NVM_DIR /home/goalist/.nvm
ENV PATH ${NVM_DIR}/versions/node/v${NODE_VERSION}/bin/:${PATH}

RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.5/install.sh | bash \
    && . ${NVM_DIR}/nvm.sh \
    && nvm install ${NODE_VERSION_16} \
    && nvm use v${NODE_VERSION_16} \
    && npm install -g yarn \
    && nvm install ${NODE_VERSION} \
    && nvm use v${NODE_VERSION} \
    && npm install -g yarn \
    && nvm alias default v${NODE_VERSION} \
    && rm -rf ${NVM_DIR}/.cache \
    && echo 'export NVM_DIR="/home/goalist/.nvm"' >>~/.bashrc \
    && echo '[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm' >> ~/.bashrc \
    && echo '[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion' >> ~/.bashrc \
    && echo 'export NVM_DIR="/home/goalist/.nvm"' >>~/.zshrc \
    && echo '[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm' >> ~/.zshrc 

# Install project dependencies
RUN npm i -g @redocly/cli
RUN go install -v github.com/ogen-go/ogen/cmd/ogen@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/google/wire/cmd/wire@latest
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to ~/bin
RUN echo 'export PATH="$PATH:$HOME/bin"' >>~/.zshrc \
    && echo 'export PATH="$PATH:$HOME/bin"' >>~/.bashrc

EXPOSE 8000-8005 9000-9005
