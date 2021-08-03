package socket

var SwaggerJSON = `
{
  "swagger": "2.0",
  "info": {
    "title": "Socket Services",
    "version": "1.0",
    "contact": {
      "name": "Impervious AI",
      "url": "https://impervious.ai"
    }
  },
  "tags": [
    {
      "name": "Socket"
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
    "/v1/socket/sendRequest": {
      "post": {
        "operationId": "Socket_SendSocket",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/socketSendSocketResponse"
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
              "$ref": "#/definitions/socketSendSocketRequest"
            }
          }
        ],
        "tags": [
          "Socket"
        ]
      }
    },
    "/v1/socket/start": {
      "post": {
        "operationId": "Socket_StartSocket",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/socketStartSocketResponse"
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
              "$ref": "#/definitions/socketStartSocketRequest"
            }
          }
        ],
        "tags": [
          "Socket"
        ]
      }
    },
    "/v1/socket/stop": {
      "post": {
        "operationId": "Socket_StopSocket",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/socketStopSocketResponse"
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
              "$ref": "#/definitions/socketStopSocketRequest"
            }
          }
        ],
        "tags": [
          "Socket"
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
    "socketSendSocketRequest": {
      "type": "object",
      "properties": {
        "pubkey": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a request to send socket connection information to a far end node"
    },
    "socketSendSocketResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a response back from the request to establish a socket"
    },
    "socketStartSocketRequest": {
      "type": "object",
      "title": "*\nRepresents a request to start a socket on an owned IMP node"
    },
    "socketStartSocketResponse": {
      "type": "object",
      "properties": {
        "protocol": {
          "type": "string"
        },
        "ip": {
          "type": "string"
        },
        "port1": {
          "type": "string"
        },
        "port2": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a response containing socket information from a started socket"
    },
    "socketStopSocketRequest": {
      "type": "object",
      "title": "*\nRepresents a request to stop a socket on an owned IMP node"
    },
    "socketStopSocketResponse": {
      "type": "object",
      "title": "*\nRepresents a response back from a stopped socket"
    }
  },
  "externalDocs": {
    "description": "Documentation on IMP",
    "url": "https://github.com/imperviousai/imp-releases"
  }
}
`
