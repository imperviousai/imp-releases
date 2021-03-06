package signing

var SwaggerJSON = `
{
  "swagger": "2.0",
  "info": {
    "title": "Signing Services",
    "version": "1.0",
    "contact": {
      "name": "Impervious AI",
      "url": "https://impervious.ai"
    }
  },
  "tags": [
    {
      "name": "Signing"
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
    "/v1/sign": {
      "post": {
        "summary": "*\nSignMessage signs a message with your node's private key.",
        "operationId": "Signing_SignMessage",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/signingSignResponse"
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
              "$ref": "#/definitions/signingSignRequest"
            }
          }
        ],
        "tags": [
          "Signing"
        ]
      }
    },
    "/v1/verify": {
      "post": {
        "summary": "*\nVerifymessage verifies a message was signed from another node.",
        "operationId": "Signing_VerifySignature",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/signingVerifyResponse"
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
              "$ref": "#/definitions/signingVerifyRequest"
            }
          }
        ],
        "tags": [
          "Signing"
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
    "signingSignRequest": {
      "type": "object",
      "properties": {
        "msg": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a request to sign a message"
    },
    "signingSignResponse": {
      "type": "object",
      "properties": {
        "signature": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a response from a signature request"
    },
    "signingVerifyRequest": {
      "type": "object",
      "properties": {
        "msg": {
          "type": "string"
        },
        "signature": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a request to verify a signature and message"
    },
    "signingVerifyResponse": {
      "type": "object",
      "properties": {
        "result": {
          "type": "boolean"
        }
      },
      "title": "*\nRepresents a response back from a verification request"
    }
  },
  "externalDocs": {
    "description": "Documentation on IMP",
    "url": "https://docs.impervious.ai"
  }
}
`
