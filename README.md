# Abstrakt

[![Build Status](https://dev.azure.com/aussiedevcrew/Wormhole/_apis/build/status/microsoft.abstrakt?branchName=master)](https://dev.azure.com/aussiedevcrew/Wormhole/_build/latest?definitionId=20&branchName=master)

# Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

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