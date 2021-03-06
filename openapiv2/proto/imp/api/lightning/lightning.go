package lightning

var SwaggerJSON = `
{
  "swagger": "2.0",
  "info": {
    "title": "Lightning Services",
    "version": "1.0",
    "contact": {
      "name": "Impervious AI",
      "url": "https://impervious.ai"
    }
  },
  "tags": [
    {
      "name": "Lightning"
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
    "/v1/lightning/checkinvoice": {
      "post": {
        "summary": "*\nCheckInvoice allows you to check a specific invoice to see if it was paid.",
        "operationId": "Lightning_CheckInvoice",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/lightningCheckInvoiceResponse"
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
              "$ref": "#/definitions/lightningCheckInvoiceRequest"
            }
          }
        ],
        "tags": [
          "Lightning"
        ]
      }
    },
    "/v1/lightning/generateinvoice": {
      "post": {
        "summary": "*\nGenerateInvoice allows you to generate an invoice for a specific payment amount from your lightning node.",
        "operationId": "Lightning_GenerateInvoice",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/lightningGenerateInvoiceResponse"
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
              "$ref": "#/definitions/lightningGenerateInvoiceRequest"
            }
          }
        ],
        "tags": [
          "Lightning"
        ]
      }
    },
    "/v1/lightning/payinvoice": {
      "post": {
        "summary": "*\nPayInvoice allows you to pay a specific invoice with your lightning node.",
        "operationId": "Lightning_PayInvoice",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/lightningPayInvoiceResponse"
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
              "$ref": "#/definitions/lightningPayInvoiceRequest"
            }
          }
        ],
        "tags": [
          "Lightning"
        ]
      }
    }
  },
  "definitions": {
    "lightningCheckInvoiceRequest": {
      "type": "object",
      "properties": {
        "invoice": {
          "type": "string"
        }
      },
      "description": "*\nRepresents an request to check an invoice."
    },
    "lightningCheckInvoiceResponse": {
      "type": "object",
      "properties": {
        "paid": {
          "type": "boolean"
        }
      },
      "description": "*\nRepresents a response back from the invoice check."
    },
    "lightningGenerateInvoiceRequest": {
      "type": "object",
      "properties": {
        "amount": {
          "type": "string",
          "format": "int64"
        },
        "memo": {
          "type": "string"
        }
      },
      "description": "*\nRepresents an invoice creation request from your lightning node."
    },
    "lightningGenerateInvoiceResponse": {
      "type": "object",
      "properties": {
        "invoice": {
          "type": "string"
        }
      },
      "description": "*\nRepresents a response back from an invoice generation request."
    },
    "lightningPayInvoiceRequest": {
      "type": "object",
      "properties": {
        "invoice": {
          "type": "string"
        }
      },
      "description": "*\nRepresents an invoice that will be paid by your lightning node."
    },
    "lightningPayInvoiceResponse": {
      "type": "object",
      "properties": {
        "preimage": {
          "type": "string"
        }
      },
      "description": "*\nRepresents a response back from the payment result."
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
