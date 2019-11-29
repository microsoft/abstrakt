# Abstrakt

[![Build Status](https://dev.azure.com/aussiedevcrew/Wormhole/_apis/build/status/microsoft.abstrakt?branchName=master)](https://dev.azure.com/aussiedevcrew/Wormhole/_build/latest?definitionId=20&branchName=master)

![Azure DevOps coverage (branch)](https://abstrakt-shields-lx.azurewebsites.net/azure-devops/coverage/aussiedevcrew/Wormhole/20/master)

## Quick Start

If you want to contribute to the project, we've made things easy for you to get everything up and running using [devcontainers](https://code.visualstudio.com/docs/remote/containers).

### Dependencies

To run the devcontainer, there's a couple things you need to install first on your machine.

- [Visual Studio Code](https://code.visualstudio.com/)
- [Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack)
- [Docker](https://docs.docker.com/install/)

A full devcontainer setup guide can be found [here](https://code.visualstudio.com/docs/remote/containers#_getting-started).

### Running

With everything installed and running, you can continue.

1. Open this project in a new instance of Visual Studio Code.
2. You may be prompted to `Reopen in Container`. Click on it.
    - If you do not see the prompt, bring up your command palette (`View` -> `Command Palette...`) and type `Remote Containers: Open Folder in Container...`
3. Wait for Visual Studio Code to reopen and build your container.
4. Once your container has finished building, you can now start developing!
5. Open a new terminal in Visual Studio Code (`Terminal` -> `New Terminal`).
6. Type `go run main.go`.

### Samples

You can find sample constellation files in the following location:

`/sample/constellation/`

Using these files you can test the app binary. 

