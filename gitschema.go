package gitschema

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/plugin"
	"github.com/posener/gitfs"
	"github.com/vektah/gqlparser/v2/ast"
	"golang.org/x/oauth2"
)

type gitschema struct {
	Config *Config
}

type Config struct {
	GitPath string
	Schema  string
}

func New(gitpath string, schema string) plugin.Plugin {
	c := &Config{
		gitpath,
		schema,
	}
	return &gitschema{c}
}

// Name returns the plugin name
func (g *gitschema) Name() string {
	return `gitschema-` + g.Config.Schema
}

func (g *gitschema) InjectSourceEarly() *ast.Source {

	// open the git repository
	client := GitAuth()
	fs, err := gitfs.New(context.Background(), g.Config.GitPath, gitfs.OptClient(client))
	if err != nil {
		log.Fatalln("Failed to open git path")
	}

	// access the file
	f, err := fs.Open(g.Config.Schema)
	if err != nil {
		log.Fatalln("Failed to open schema file: " + err.Error())
	}

	// read the file
	schemaRaw, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("Failed to read the schema file")
	}

	// return the source
	return &ast.Source{
		Name:    g.Config.GitPath,
		Input:   string(schemaRaw),
		BuiltIn: true,
	}
}

func GitAuth() *http.Client {
	token := os.Getenv("GIT_OAUTH_TOKEN")
	if token == "" {
		log.Fatalln("Missing env var: GIT_OAUTH_TOKEN")
	}
	return oauth2.NewClient(
		context.Background(),
		oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))
}
