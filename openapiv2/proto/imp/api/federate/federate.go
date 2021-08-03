package federate

var SwaggerJSON = `
{
  "swagger": "2.0",
  "info": {
    "title": "Federate Services",
    "version": "1.0",
    "contact": {
      "name": "Impervious AI",
      "url": "https://impervious.ai"
    }
  },
  "tags": [
    {
      "name": "Federate"
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
    "/v1/federate/leave": {
      "post": {
        "summary": "*\nLeaveFederation performs the removal of a federated peer (upon message receipt).",
        "operationId": "Federate_LeaveFederation",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/federateLeaveFederationResponse"
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
              "$ref": "#/definitions/federateLeaveFederationRequest"
            }
          }
        ],
        "tags": [
          "Federate"
        ]
      }
    },
    "/v1/federate/request": {
      "post": {
        "summary": "*\nRequestFederation performs the federation request to a specific peer.",
        "operationId": "Federate_RequestFederate",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/federateRequestFederateResponse"
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
              "$ref": "#/definitions/federateRequestFederateRequest"
            }
          }
        ],
        "tags": [
          "Federate"
        ]
      }
    }
  },
  "definitions": {
    "federateLeaveFederationRequest": {
      "type": "object",
      "properties": {
        "pubkey": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a request to leave a federation from a far end node"
    },
    "federateLeaveFederationResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a response back from a Leave Federation Request"
    },
    "federateRequestFederateRequest": {
      "type": "object",
      "properties": {
        "pubkey": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a request to federate with a far end node"
    },
    "federateRequestFederateResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a response back from a Federation Request"
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
    "url": "https://github.com/imperviousai/imp-releases"
  }
}
`
