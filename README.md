# Abstrakt

[![build](https://github.com/microsoft/abstrakt/workflows/build/badge.svg?branch=master)](https://github.com/microsoft/abstrakt/actions?query=workflow%3Abuild) ![Azure DevOps coverage (branch)](https://abstrakt-shields.azurewebsites.net/azure-devops/coverage/aussiedevcrew/Wormhole/20/master) [![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

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

`/examples/constellation/`

Using these files you can test the app binary. 

#### Running the local Helm Microservice Sample

This sample uses the Http Microservice test constellation file located in `sample/constellation/http_constellation.yaml`.

The sample will install three services with relationships between them:

- A sender
- Another sender
- A receiver

The concept is that you send a message down the chain. 

- The first sender takes a message via http, adds value to the message and forwards it to the second sender. 
- The second sender repeats this, also adding a unique value before sending on to the receiver. 
- The receiver echos back the message
- Each of the sender returns the message back down the chain before the final message pops out the original http request. 

##### Running the sample

Run:

`make run-http-demo`

This will deploy the demo to the cluster. The templates are output to `output/http_sample/Output`. 

You can test it by:

```
kubectl wait pod -n default --for condition=ready --timeout=120s --all && kubectl exec -it sender-sender bash
curl sender:8080/api/CallReceiver?message=yourmessage
```

You should see the response "yourmessage--Chain-sender--Chain-sender1" which proves the message traversed the chain!

##### Remove the install

You can remove the installed pods and services by calling `make http-demo-delete`

##### Debugging Templates

You can debug the templates by outputting them by running `make http-demo-template-all`


