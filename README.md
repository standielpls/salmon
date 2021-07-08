# salmon

Salmon is a CLI tool meant to help manage and automate Google form submissions.
Organizations often require a log of problems and solutions weekly and this project
is meant to help alleviate the mundane task of typing it all out.

This project uses Go.

## How to use

### Prerequisites

```
make init
```

config.yaml

```
user: {github user name}
full_name: {full name}
owner: {github owner}
repo: {github repo}
token: {github access token}
```

Build the binary

```
make build
```

Run the binary

```
bin/salmon
```
