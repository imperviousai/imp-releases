# Impervious Releases

This is the repository for impervious releases and supporting files and documentation.

## Binaries

The binaries are now released and the latest versions can be found in the release/tag section [https://github.com/imperviousai/imp-releases/releases](https://github.com/imperviousai/imp-releases/releases) in github. Find the version meant for your operating system and download. 

If you are on Mac, it may get flagged as unverified. You can shift+right click on the application and open it to approve it once. Close it and then run it normally from the command line, passing in your config.

## Docker

Instead of the binaries, you may use Docker with the binaries already loaded in. Just create your config file, set a few necessary lines specific to Docker, and you're ready to go just as if you were running the binary normally on your computer. Follow the instructions for that [here](https://hub.docker.com/r/impant/imp-releases). 

## Sample Config 

We have a base config file you may build off of to run with your Imp deamon. Copy `config.yml` to a location next to your deamon, fill in the blanks, and you can run with `./impervious --config=config.yml`.

More info on the config can be found here: https://docs.impervious.ai/#impervious-node-configuration 

## Docs

Markdown docs for the API's may be viewable in `/docs` for each service. There is a nice format of these docs imported into https://docs.impervious.ai as well.

You may also see the docs on your Imp deamon by accessing it through the http proxy. By default: `localhost:8882/docs`

## GRPC 

The Imp daemon API's can be accessible with both GRPC and HTTP. For accessing via GRPC, you may utilize the proto files located at `/proto/` of this repo. We have also generated a few client libraries so you can get started much faster by using them in your project. For now we support: `go`, `rust`, & `js` but if there's another language you'd like generated client libraries for, please let us know.

You can test out grpc from the command line to make sure it is working. Install [grpcurl](https://github.com/fullstorydev/grpcurl) and you can run an example command from this repo directory.

```
grpcurl -plaintext -import-path ./proto -proto ./proto/messaging.proto -d @ 127.0.0.1:8881 messaging.Messaging.SendMessage <<EOM
{
  "amount": 1234,
  "msg": "test123",
  "pubkey": "03ed306a57ae78941d8204bdda314560eb1651ff18d727404e93d639cc885eb03d"
}
EOM
```

Replace that pubkey with one of your other node's pubkeys (not the one you are accessing via grpc). Or you may leave it but you'll get a generic GRPC error, which is fine because at least you know grpc works. 

## OpenAPI

We also supply OpenAPIv2 files so that you may import them into your own tools, whether that's postman or something like https://editor.swagger.io

These files are split up by service and can be accessible at `/openapiv2/`

## Discord

Join our discord for help, support, and to communicate with the team. https://discord.gg/5YdUJvvNuZ 
