# Installation instructions for Mac

## Prerequisites

### Install go
[Installation instructions](https://go.dev/doc/install)

### Install make
https://formulae.brew.sh/formula/make
```bash
$ brew install make
```

### Install NPM (optional, for hot-reloading with nodemon)
https://docs.npmjs.com/downloading-and-installing-node-js-and-npm

## Starting project
```bash
# EITHER NPM
$ npm install # only needed once
$ npm run dev # hot-reloading

# OR MAKE
$ make install # only needed once
$ make run
```

then navigate to http://localhost:3000/swagger to see the swagger docs
