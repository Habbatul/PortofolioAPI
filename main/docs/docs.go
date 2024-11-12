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
        "/projects": {
            "post": {
                "description": "Create a new project with associated tags and images using form-data",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Create a new project with images and tags",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project Name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Project Description",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Project Note",
                        "name": "note",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Project URL",
                        "name": "url_project",
                        "in": "formData"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "Tag item",
                        "name": "tags[]",
                        "in": "formData"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "file"
                        },
                        "collectionFormat": "multi",
                        "description": "Images",
                        "name": "images",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.Project"
                        }
                    },
                    "400": {
                        "description": "Invalid form data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to create project",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/projects/{id}": {
            "get": {
                "description": "Retrieve project details by project ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Get project by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ProjectResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid project ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Project not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Image": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "project_id": {
                    "type": "integer"
                },
                "url_img": {
                    "type": "string"
                }
            }
        },
        "entity.Project": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Image"
                    }
                },
                "name": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Tag"
                    }
                },
                "url_project": {
                    "type": "string"
                }
            }
        },
        "entity.Tag": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name_tag": {
                    "type": "string"
                },
                "projects": {
                    "description": "ProjectIDs akan dikelola oleh GORM sebagai bagian dari relasi many-to-many",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Project"
                    }
                }
            }
        },
        "model.ProjectResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "url_project": {
                    "type": "string"
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