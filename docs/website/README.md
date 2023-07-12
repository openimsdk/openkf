# OpenKF's docs

OpenKF's docs are built with [hugo](https://gohugo.io/), and can be browsed live at https://kind.sigs.k8s.io/

To browse them locally, install hugo and run `make serve` from this directory.

Read [tutorial](https://getdoks.org/docs/overview/introduction/), learning to use OpenKF document construction


## File Directory Description

| Directory  | Description                                                  |
| ---------- | ------------------------------------------------------------ |
| archetypes | Directory for creating new content files with Hugo command.  |
| assets     | Directory for storing files that need to be processed by Hugo Pipes. |
| config     | Directory for storing configuration directives in JSON, YAML, or TOML format. |
| content    | Directory for storing all the content of your website. Each top-level folder is considered a content section. |
| data       | Directory for storing configuration files that can be used by Hugo during website generation. |
| layouts    | Directory for storing template files specifying how to render content views in the static website. |
| static     | Directory for storing all static content such as images, CSS, JavaScript, etc. |
| resources  | Directory for caching files to speed up the generation process. |


## Command

> **Note**💡: 
> You can change the commands in the scripts section of `./package.json`.

### create

Create new content for your site:

```bash
npm run create [path] [flags]
```

See also the Hugo docs: [hugo new](https://gohugo.io/commands/hugo_new/).

### Docs based tree

Create a docs based tree — with a single command:

```bash
npm run create -- --kind docs [section]
```

For example, create a docs based tree named guides:

```bash
npm run create -- --kind docs guides
```

## lint

Check scripts, styles, and markdown for errors:

```bash
npm run lint
```

### scripts

Check scripts for errors:

```bash
npm run lint:scripts [-- --fix]
```

### styles

Check styles for errors:

```bash
npm run lint:styles [-- --fix]
```

### markdown

Check markdown for errors:

```bash
npm run lint:markdown [-- --fix]
```

## clean

Delete temporary directories:

```bash
npm run clean
```

## start

Start local development server:

```bash
npm run start
```

## build

Build production website:

```bash
npm run build
```

### functions

Build Lambda functions:

```bash
npm run build:functions
```

### preview

Build production website including draft and future content:

```bash
npm run build:preview
```