module github.com/microsoft/abstrakt

go 1.14

replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309

require (
	github.com/awalterschulze/gographviz v0.0.0-20190522210029-fa59802746ab
	github.com/deckarep/golang-set v1.7.1
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b
	github.com/mitchellh/go-homedir v1.1.0
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.6
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
	golang.org/x/crypto v0.0.0-20200311171314-f7b00557c8c4
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/dealancer/validate.v2 v2.1.0
	gopkg.in/yaml.v2 v2.2.8
	helm.sh/helm/v3 v3.1.2
	sigs.k8s.io/yaml v1.1.0
)
