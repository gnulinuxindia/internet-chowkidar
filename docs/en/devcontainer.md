# Devcontainers
Development containers, or dev containers, are Docker containers that are specifically configured to provide a fully featured development environment. 

## Prerequisites
In order to get started, you need to satisfy the following prerequisites:
- [Docker](https://docs.docker.com/get-docker/)
- [Visual Studio Code](https://code.visualstudio.com/)
- [Visual Studio Code Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

It is recommended you allocate at least 2GB of memory to Docker.
- [Instructions for Windows](https://docs.docker.com/docker-for-windows/#resources)
- [Instructions for macOS](https://docs.docker.com/docker-for-mac/#resources)

## Getting Started
Clone the repository and follow the steps below to start developing in a devcontainer.

### Use VSCode Remote Containers Extension
For most people getting started with development, the best solution is to use [VSCode Remote - Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers).

VSCode should automatically suggest installing the required extensions. They can also be installed manually as follows:
- Install Remote - Containers for VSCode
  - through command line code --install-extension ms-vscode-remote.remote-containers
  - clicking on the Install button in the Vistual Studio Marketplace: [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
  - View: Extensions command in VSCode (Windows: <kbd>Ctrl</kbd><kbd>Shift</kbd><kbd>X</kbd>; macOS: <kbd>Cmd (⌘)</kbd><kbd>Shift</kbd><kbd>X</kbd>) then search for extension `ms-vscode-remote.remote-containers`.

After the extension is installed, you can:
- Open the repository in VSCode
     - `code .`
     
Launch the following command from Command Palette (Windows: <kbd>Ctrl</kbd><kbd>Shift</kbd><kbd>P</kbd>, macOS: <kbd>Cmd (⌘)</kbd>+<kbd>Shift</kbd>+<kbd>X</kbd>) `Remote-Containers: Reopen in Container.`

You can also click the green button in the bottom left corner to access the remote container menu.

> [!NOTE] 
> The first time you open the repository in a devcontainer, it will take some time to build the container. Subsequent openings will be faster.

> [!TIP]
> The repository will be mounted and available in the container.

### Running without vscode
If you prefer not to use VSCode, you can still run the container manually to get a shell into the container. 
First, install the devcontainer CLI
```bash
npm install -g devcontainer-cli
```
Then, run the following command in the root of the repository.
```bash
devcontainer up --workspace-folder .
```

Finally, to get a shell in the container, run the following command.
```bash
devcontainer exec --workspace-folder . bash # (or zsh, depending on your shell preference)
```
