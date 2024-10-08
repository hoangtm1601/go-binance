// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/merchants": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "CreateMerchant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "merchants"
                ],
                "summary": "CreateMerchant",
                "parameters": [
                    {
                        "description": "merchant",
                        "name": "merchant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_VuThanhThien_golang-gorm-postgres_merchant_internal_models_dto.CreateMerchantDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/merchants/merchant-id/{merchantID}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "GetMerchantByMerchantID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "merchants"
                ],
                "summary": "GetMerchantByMerchantID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "merchantID",
                        "name": "merchantID",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/merchants/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "GetMerchant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "merchants"
                ],
                "summary": "GetMerchant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_VuThanhThien_golang-gorm-postgres_merchant_internal_models_dto.CreateMerchantDTO": {
            "type": "object",
            "required": [
                "merchant_id",
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "merchant_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        },
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8001",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Swagger Merchant API",
	Description:      "This is merchant golang server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
