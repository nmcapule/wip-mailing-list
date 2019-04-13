# WIP + Mailing List Temp Site

This repo is created to quickly whip up temporary sites for websites that are in development.

Note that the current website is more geared towards the AppEngine platform.

## Requirements

1. Make sure you install the Google Cloud SDK and Go tools.
2. Setup your default project to the AppEngine project that you wnt to deploy this app into.
3. Edit `app.yaml` and change all instances of the string INSTANCE_CONNECTION_NAME w/ what's in your DB instance connection name field.

## Deploying locally

Run the ff. on the command line:

```sh
$ dev_appserver.py app.yaml
```

## Deploying to AppEngine

1. Go configure `app.yaml` and uncomment the line for local instance connection string.

2. Run the ff. on the command line:

```sh
$ gcloud app deploy --version="<myversion>"
```