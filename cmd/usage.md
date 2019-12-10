# Abstrakt Usage

### Abstrakt
Scalable, config driven data pipelines for Kubernetes.

Abstrakt is a command line utility for processing constellation files. It has a number of subcommands shown below.

```bash
Usage:
  abstrakt [command]

Available Commands:
  compose     Compose a package into requested template type
  help        Help about any command
  validate    Validate a constellation file for correct schema and ensure correctness.
  version     The version of Abstrakt being used
  visualise   Format a constellation configuration as Graphviz dot notation

Flags:
  -h, --help      help for abstrakt
  -v, --verbose   Use verbose output logs

Use "abstrakt [command] --help" for more information about a command.
```

### abstrakt `compose`

```bash
Compose is for composing a package based on mapsFilePath and constellationFilePath and template (default value is helm).

Example: abstrakt [chart name] compose -t [templateType] -f [constellationFilePath] -m [mapsFilePath] -o [outputPath] -z

Usage:
  abstrakt compose [chart name] [flags]

Flags:
  -f, --constellationFilePath string   constellation file path
  -h, --help                           help for compose
  -m, --mapsFilePath string            maps file path
  -o, --outputPath string              destination directory
  -t, --templateType string            output template type (default "helm")
  -z, --zipChart                       zips the chart

Global Flags:
  -v, --verbose   Use verbose output logs
```

Can compose a Helm chart directory (default) or a __.tgz__ of the produced helm chart (with `-z` flag).

#### Examples

Create a Helm chart named `http-demo` to be generated under ./output.

```bash
./abstrakt compose http-demo -f ./sample/constellation/http_constellation.yaml -m ./sample/constellation/http_constellation_maps.yaml -o ./output/http-demo 
```

With __.tgz__
```bash
./abstrakt compose http-demo -f ./sample/constellation/http_constellation.yaml -m ./sample/constellation/http_constellation_maps.yaml -o ./output/http-demo -z
```

### abstrakt `validate`

```bash
Validate is used to ensure the correctness of a constellation file.

Example: abstrakt validate -f [constellationFilePath]

Usage:
  abstrakt validate [flags]

Flags:
  -f, --constellationFilePath string   constellation file path
  -h, --help                           help for validate

Global Flags:
  -v, --verbose   Use verbose output logs
```

### abstrakt `visualise`

```bash
Visualise is for producing Graphviz dot notation code of a constellation configuration

Example: abstrakt visualise -f [constellationFilePath]

Usage:
  abstrakt visualise [flags]

Flags:
  -f, --constellationFilePath string   constellation file path
  -h, --help                           help for visualise

Global Flags:
  -v, --verbose   Use verbose output logs
```

The output from the visualise subcommand is [Graphviz dot notation](https://www.graphviz.org/doc/info/lang.html)

The output from a call to 'abstrakt visualise' can be piped into Graphviz to generate a graphical output. See the example in the Examples section. 

Alternatively, copy the output and paste into a Graphviz rendering tool to see the graph produced. Some sites listed below (rendering option in the utility to be developed).  

[Graphviz online](https://dreampuf.github.io/GraphvizOnline/)  
[Webgraphviz](http://www.webgraphviz.com/)  


#### Examples

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