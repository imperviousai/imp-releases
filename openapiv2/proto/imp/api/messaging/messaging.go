package messaging

var SwaggerJSON = `
{
  "swagger": "2.0",
  "info": {
    "title": "Messaging Services",
    "version": "1.0",
    "contact": {
      "name": "Impervious AI",
      "url": "https://impervious.ai"
    }
  },
  "tags": [
    {
      "name": "Messaging"
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
    "/v1/message/send": {
      "post": {
        "summary": "*\nSendMessage sends a text message to another node.",
        "operationId": "Messaging_SendMessage",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/messagingSendMessageResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/messagingSendMessageRequest"
            }
          }
        ],
        "tags": [
          "Messaging"
        ]
      }
    }
  },
  "definitions": {
    "messagingSendMessageRequest": {
      "type": "object",
      "properties": {
        "msg": {
          "type": "string",
          "format": "byte"
        },
        "pubkey": {
          "type": "string"
        },
        "amount": {
          "type": "string",
          "format": "int64"
        },
        "replyToId": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a message send to a far end node"
    },
    "messagingSendMessageResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a response back from a sent message"
    },
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
    }
  },
  "externalDocs": {
    "description": "Documentation on IMP",
    "url": "https://docs.impervious.ai"
  }
}
`
