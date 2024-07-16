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
                        "description": "course read request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/kazusa-server_internal_entity.CourseReadRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/kazusa-server_internal_entity.Course"
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
                            "$ref": "#/definitions/kazusa-server_internal_entity.NewModule"
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
                        "description": "module read request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/kazusa-server_internal_entity.ModuleReadRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/kazusa-server_internal_entity.Module"
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
        },
        "kazusa-server_internal_entity.Course": {
            "type": "object",
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
        "kazusa-server_internal_entity.CourseFilters": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "kazusa-server_internal_entity.CourseReadRequest": {
            "type": "object",
            "properties": {
                "filters": {
                    "$ref": "#/definitions/kazusa-server_internal_entity.CourseFilters"
                },
                "pagination": {
                    "$ref": "#/definitions/kazusa-server_internal_entity.Pagination"
                }
            }
        },
        "kazusa-server_internal_entity.Module": {
            "type": "object",
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
        "kazusa-server_internal_entity.ModuleFilters": {
            "type": "object",
            "properties": {
                "courseId": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "kazusa-server_internal_entity.ModuleReadRequest": {
            "type": "object",
            "properties": {
                "filters": {
                    "$ref": "#/definitions/kazusa-server_internal_entity.ModuleFilters"
                },
                "pagination": {
                    "$ref": "#/definitions/kazusa-server_internal_entity.Pagination"
                }
            }
        },
        "kazusa-server_internal_entity.NewCourse": {
            "type": "object",
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
        "kazusa-server_internal_entity.NewModule": {
            "type": "object",
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
        "kazusa-server_internal_entity.Pagination": {
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
