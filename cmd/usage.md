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

### Commands
#### Visualise

The output from the visualise subcommand is [Graphviz dot notation](https://www.graphviz.org/doc/info/lang.html)

The output from a call to 'abstrakt visualise' can be piped into Graphviz to generate a graphical output. See the example in the Examples section. 

Alternatively, copy the output and paste into a Graphviz rendering tool to see the graph produced. Some sites listed below (rendering option in the utility to be developed).  

[Graphviz online](https://dreampuf.github.io/GraphvizOnline/)  
[Webgraphviz](http://www.webgraphviz.com/)  


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

Run visualise on a file  
	
	abstrakt visualise -f basic_azure_event_hubs.yaml
	digraph Azure_Event_Hubs_Sample {
	        Event_Generator->Azure_Event_Hub;
	        Azure_Event_Hub->Event_Logger;
	        Azure_Event_Hub;
	        Event_Generator;
	        Event_Logger;
	
	}
	
Pipe visualise output to Graphviz producing a file called result.png (assumes Graphviz is installed and can be called from the location abstrakt is being run)

	abstrakt visualise -f ./sample/constellation/sample_consteallation.yaml | dot -Tpng > result.png