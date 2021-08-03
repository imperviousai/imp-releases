package websocket

var SwaggerJSON = `
{
  "swagger": "2.0",
  "info": {
    "title": "Websocket Services",
    "version": "1.0",
    "contact": {
      "name": "Impervious AI",
      "url": "https://impervious.ai"
    }
  },
  "tags": [
    {
      "name": "Websocket"
    }
  ],
  "schemes": [
    "http",
    "https",
    "wss"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/subscribe": {
      "get": {
        "operationId": "Websocket_Subscribe",
        "responses": {
          "200": {
            "description": "A successful response.(streaming responses)",
            "schema": {
              "type": "object",
              "properties": {
                "result": {
                  "$ref": "#/definitions/websocketSubscribeResponse"
                },
                "error": {
                  "$ref": "#/definitions/rpcStatus"
                }
              },
              "title": "Stream result of websocketSubscribeResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "tags": [
          "Websocket"
        ]
      }
    }
  },
  "definitions": {
    "protobufAny": {
      "type": "object",
      "properties": {
        "typeUrl": {
          "type": "string"
        },
        "value": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "rpcStatus": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/protobufAny"
          }
        }
      }
    },
    "websocketSubscribeResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "replyToId": {
          "type": "string"
        },
        "fromPubkey": {
          "type": "string"
        },
        "data": {
          "type": "string"
        },
        "serviceType": {
          "type": "string"
        },
        "amount": {
          "type": "string",
          "format": "int64"
        }
      },
      "title": "*\nRepresents a response back from the websocket containing event information"
    }
  },
  "externalDocs": {
    "description": "Documentation on IMP",
    "url": "https://github.com/imperviousai/imp-releases"
  }
}
`
