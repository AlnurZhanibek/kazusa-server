// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://kazusa.kz",
            "email": "aln.zh.621@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/course": {
            "get": {
                "description": "read courses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Read courses",
                "operationId": "course.read",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Course"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            },
            "post": {
                "description": "create module",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create module",
                "operationId": "module.create",
                "parameters": [
                    {
                        "description": "new module body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/NewModule"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login a user",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "login body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_handler.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_handler.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_handler.LoginResponse"
                        }
                    }
                }
            }
        },
        "/module": {
            "get": {
                "description": "read modules",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Read modules",
                "operationId": "module.read",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "course_id",
                        "name": "course_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Module"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Register a user",
                "operationId": "register",
                "parameters": [
                    {
                        "description": "register body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_handler.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_handler.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_handler.RegisterResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Course": {
            "type": "object",
            "required": [
                "createdAt",
                "description",
                "id",
                "price",
                "title"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "modules": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Module"
                    }
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "Module": {
            "type": "object",
            "required": [
                "content",
                "courseId",
                "createdAt",
                "durationMinutes",
                "id",
                "name"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "courseId": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "durationMinutes": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "NewCourse": {
            "type": "object",
            "required": [
                "description",
                "price",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "NewModule": {
            "type": "object",
            "required": [
                "content",
                "courseId",
                "durationMinutes",
                "name"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "courseId": {
                    "type": "string"
                },
                "durationMinutes": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "internal_handler.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "internal_handler.LoginResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "internal_handler.RegisterRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "passwordConfirmation": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "internal_handler.RegisterResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "kazusa.kz",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger KazUSA API",
	Description:      "This is the KazUSA server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
