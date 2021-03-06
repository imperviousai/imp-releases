package vpn

var SwaggerJSON = `
{
  "swagger": "2.0",
  "info": {
    "title": "VPN Services",
    "version": "1.0",
    "contact": {
      "name": "Impervious AI",
      "url": "https://impervious.ai"
    }
  },
  "tags": [
    {
      "name": "VPN"
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
    "/v1/vpn/contract": {
      "post": {
        "summary": "*\nAcceptContract accepts and pays for a specific contract.",
        "operationId": "VPN_AcceptContract",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/vpnAcceptContractResponse"
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
              "$ref": "#/definitions/vpnAcceptContractRequest"
            }
          }
        ],
        "tags": [
          "VPN"
        ]
      }
    },
    "/v1/vpn/quote": {
      "post": {
        "summary": "*\nRequestQuote requests a quote from another node.",
        "operationId": "VPN_RequestQuote",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/vpnRequestQuoteResponse"
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
              "$ref": "#/definitions/vpnRequestQuoteRequest"
            }
          }
        ],
        "tags": [
          "VPN"
        ]
      }
    },
    "/v1/vpn/refresh": {
      "post": {
        "summary": "*\nRefreshContract accepts and pays for a specific contract refresh on an existing VPN.",
        "operationId": "VPN_RefreshContract",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/vpnRefreshContractResponse"
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
              "$ref": "#/definitions/vpnRefreshContractRequest"
            }
          }
        ],
        "tags": [
          "VPN"
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
    "vpnAcceptContractRequest": {
      "type": "object",
      "properties": {
        "pubkey": {
          "type": "string"
        },
        "nonce": {
          "type": "string"
        },
        "price": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a request to Accept (Pay For) a VPN Quote"
    },
    "vpnAcceptContractResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a response back from an accepted VPN Quote"
    },
    "vpnRefreshContractRequest": {
      "type": "object",
      "properties": {
        "pubkey": {
          "type": "string"
        },
        "nonce": {
          "type": "string"
        },
        "price": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a request to extend/refresh an expiring VPN Connection (i.e. purchase more time)"
    },
    "vpnRefreshContractResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a reponse back from a refreshed VPN connection"
    },
    "vpnRequestQuoteRequest": {
      "type": "object",
      "properties": {
        "pubkey": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a request to receive a VPN quote from a far end node"
    },
    "vpnRequestQuoteResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      },
      "title": "*\nRepresents a response back from a VPN Quote Reqeust"
    }
  },
  "externalDocs": {
    "description": "Documentation on IMP",
    "url": "https://docs.impervious.ai"
  }
}
`
