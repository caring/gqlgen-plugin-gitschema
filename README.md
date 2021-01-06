# gqlgen-plugin-gitschema

This plugin works with [github.com/99designs/gqlgen](https://github.com/99designs/gqlgen) GraphQL code generation to allow reading schemas from remote git repositories.

See [the official gqlgen plugin documentation](https://gqlgen.com/reference/plugins/) to understand how plugins work.

This plugin requires two arguments when initializing:
```
gitschema.New("github.com/caring/addressgeo/api/graphql", "addressgeo.graphqls")
```

The first is the full repo path to the directory in git holding the schema file, the second is the filename itself. Please note that the git path excluded host http path like `tree`, `ref`, etc.

Additionally, this plugin uses OAuth authentication to access the GIT repository holding the schema file. It uses an environment variable `GIT_OAUTH_TOKEN`, and assumes all git repos being accessed will use the same authentication.
