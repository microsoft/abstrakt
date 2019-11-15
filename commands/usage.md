# Abstrakt Usage

### Abstrakt
Scalable, config driven data pipelines for Kubernetes.

Abstrakt is a command line utility for processing constellation files. It has a number of subcommands shown below.


    Usage:
    abstrakt [command]

    Available Commands:
    compose   -  compose a package into requested template type 
    help - Help about any command
    version - The version of Abstrakt being used
    visualise  - format a constellation configuration as Graphviz dot notation

    Flags:
    -h, --help      help for abstrakt
    -v, --verbose   Use verbose output logs`


`Use "abstrakt [command] --help" for more information about a command.`


### Examples

Get help on a command 'visualise'  

    abstrakt visualise --help | abstrakt help visualise

    visualise is for producing Graphviz dot notation code of a constellation configuration
    abstrakt visualise -f [constellationFilePath]

    Usage:
    abstrakt visualise [flags]

    Flags:
    -f, --constellationFilePath string   constellation file path
    -h, --help                           help for visualise

    Global Flags:
    -v, --verbose   Use verbose output logs 
  



Show current application version  

    abstract version
    
    INFO[15-11-2019 05:01:16] abstrakt version 0.0.1 

    