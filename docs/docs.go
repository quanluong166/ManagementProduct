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
        "/products/export": {
            "get": {
                "description": "Export list of products by filter to PDF",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Export list of products by filter to PDF",
                "parameters": [
                    {
                        "description": "Export product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.ExportProductRequest"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "export error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products/filter": {
            "post": {
                "description": "Get list of products by filter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get product by filter",
                "parameters": [
                    {
                        "description": "Get product by filter",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.GetProductsByFilterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Get product successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ExportProductRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "api.GetProductsByFilterRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "city_sort": {
                    "description": "ASC, DESC",
                    "type": "string"
                },
                "paging": {
                    "$ref": "#/definitions/models.Paging"
                },
                "price_sort": {
                    "description": "ASC, DESC",
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.Paging": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
