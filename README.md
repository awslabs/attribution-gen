# attribution-gen CLI tool

`attribution-gen` is a tools that helps you generate license attributions
files. Currently it only works with Go projects using modules.

We hope you find this tool useful; please verify the accuracy of the detected
third party code and licenses.

## Installing

- Install to $GOBIN
    - `go install github.com/awslabs/attribution-gen/...@latest`


## Using attribution-gen

The easiest way to use `attribution-gen` is to run it inside the directory 
of your Go project.

```bash
attribution-gen --debug # default generated file name is ATTRIBUTIONS.md
```

By default the max depth allowed while exploring the dependency graph is 2,
you can override this value by using the `--depth` flag.

```bash
attribution-gen --depth 5 --debug
```

You can also set the output/input and the templates used to generate the
attributions file.

```bash
attribution-gen --output ATTRIBUTIONS.md --modfile go.mod\
    --attr-header-template $(HEADER_TMP)\
    --attr-block-template $(BLOCK_TMP)
```

You can also print the dependency graph using the `--show-graph` flag

```bash
attribution-gen --show-graph --depth 2

# OUTPUT
INFO[0004] attribution-gen
├── github.com/google/go-cmdtest@v0.4.0 Apache-2.0
│   ├── github.com/google/go-cmp@v0.5.9 BSD-3-Clause
│   └── github.com/google/renameio@v1.0.1 Apache-2.0
├── github.com/google/licenseclassifier/v2@v2.0.0 Apache-2.0
│   ├── github.com/davecgh/go-spew@v1.1.1 ISC
│   ├── github.com/google/go-cmp@v0.5.9 BSD-3-Clause
│   └── github.com/sergi/go-diff@v1.2.0 Copyright
├── github.com/sirupsen/logrus@v1.8.1 MIT
│   ├── github.com/davecgh/go-spew@v1.1.1 ISC
│   ├── github.com/pmezard/go-difflib@v1.0.0 BSD-3-Clause
│   ├── github.com/stretchr/testify@v1.2.2 MIT
│   └── golang.org/x/sys@v0.1.0 BSD-3-Clause
├── github.com/spf13/cobra@v1.4.0 Apache-2.0
│   ├── github.com/cpuguy83/go-md2man/v2@v2.0.1 MIT
│   ├── github.com/inconshreveable/mousetrap@v1.0.0 Apache-2.0
│   ├── github.com/spf13/pflag@v1.0.5 BSD-3-Clause
│   └── gopkg.in/yaml.v2@v2.4.0 Apache-2.0
├── github.com/xlab/treeprint@v1.1.0 MIT
│   └── github.com/stretchr/testify@v1.7.0 MIT
├── golang.org/x/exp@v0.0.0-20230522175609-2e198f4a06a1 BSD-3-Clause
│   ├── github.com/google/go-cmp@v0.5.9 BSD-3-Clause
│   ├── golang.org/x/mod@v0.6.0 BSD-3-Clause
│   ├── golang.org/x/tools@v0.2.0 BSD-3-Clause
│   └── golang.org/x/sys@v0.1.0 BSD-3-Clause
└── golang.org/x/mod@v0.6.0 BSD-3-Clause 
```

## Testing

`make cli-test`

## Security

See [CONTRIBUTING](CONTRIBUTING.md#security-issue-notifications) for more information.

## License

This project is licensed under the Apache-2.0 License.

