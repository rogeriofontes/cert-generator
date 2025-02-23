// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://meusistema.com/termos/",
        "contact": {
            "name": "Suporte da API",
            "url": "http://meusistema.com/suporte",
            "email": "suporte@meusistema.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/communities": {
            "post": {
                "description": "ption Criar um novo Comunidadeo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comunidade"
                ],
                "summary": "Criar um novo Comunidadeo",
                "parameters": [
                    {
                        "description": "Comunidade a ser criada",
                        "name": "community",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Community"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Comunidade criada com sucesso!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Dados inválidos",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro ao criar Comunidadeo",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/community": {
            "get": {
                "description": "Buscar todos os Comunidadeos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comunidade"
                ],
                "summary": "Buscar todos os Comunidadeos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Community"
                            }
                        }
                    }
                }
            }
        },
        "/events": {
            "get": {
                "description": "Buscar todos os eventos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Eventos"
                ],
                "summary": "Buscar todos os eventos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Event"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Cadastrar um novo evento",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Eventos"
                ],
                "summary": "Cadastrar um novo evento",
                "parameters": [
                    {
                        "description": "Evento a ser cadastrado",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Event"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Evento cadastrado com sucesso",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Dados inválidos",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Erro ao cadastrar evento",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/events/{eventID}/participants": {
            "get": {
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do evento",
                        "name": "eventID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Participantes encontrados",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Participant"
                            }
                        }
                    },
                    "400": {
                        "description": "ID do evento inválido",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Erro ao buscar participantes",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/participants": {
            "get": {
                "description": "Buscar todos os participantes cadastrados",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Participantes"
                ],
                "summary": "Buscar todos os participantes",
                "responses": {
                    "200": {
                        "description": "Participantes encontrados",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Participant"
                            }
                        }
                    },
                    "500": {
                        "description": "Erro ao buscar participantes",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Cadastrar um novo participante em um evento",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Participantes"
                ],
                "summary": "Cadastrar um novo participante",
                "parameters": [
                    {
                        "description": "Participante a ser cadastrado",
                        "name": "participant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Participant"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Participante cadastrado com sucesso",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Dados inválidos",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Erro ao cadastrar participante",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/participants/validate": {
            "get": {
                "description": "Valida um certificado pelo código UUID gerado",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Certificados"
                ],
                "summary": "Valida um certificado",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Código do certificado",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Certificado válido",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Código ausente",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Certificado não encontrado",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/participants/{id}": {
            "get": {
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do participante",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Participante encontrado",
                        "schema": {
                            "$ref": "#/definitions/domain.Participant"
                        }
                    },
                    "400": {
                        "description": "ID inválido",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Participante não encontrado",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Community": {
            "type": "object"
        },
        "domain.Event": {
            "type": "object"
        },
        "domain.Participant": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9393",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "API de Certificados",
	Description:      "API para geração e validação de certificados.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
