package horse

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/domains": {
            "get": {
                "tags": ["domain"],
                "summary": "List Domains",
                "description": "List all domains that the current user has a role.",
                "operationId": "v1_list_domains",
                "parameters": [
                    {
                        "required": false,
                        "schema": {
                            "title": "Roles",
                            "type": "array",
                            "items": { "type": "string" }
                        },
                        "name": "roles",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Groups",
                            "type": "array",
                            "items": { "type": "string" }
                        },
                        "name": "groups",
                        "in": "query"
                    },
                    {
                        "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                        "required": false,
                        "schema": {
                            "title": "Ordering",
                            "type": "string",
                            "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                            "default": ""
                        },
                        "name": "ordering",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Offset",
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 0
                        },
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Limit",
                            "maximum": 500.0,
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 100
                        },
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/DomainListResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "post": {
                "tags": ["domain"],
                "summary": "Create Domain",
                "operationId": "v1_create_domain",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/DomainCreate" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/DomainResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/groups": {
            "get": {
                "tags": ["domain"],
                "summary": "Search Domain Groups",
                "operationId": "v1_search_domain_groups",
                "parameters": [
                    {
                        "description": "search query",
                        "required": true,
                        "schema": {
                            "title": "Query",
                            "minLength": 2,
                            "type": "string",
                            "description": "search query"
                        },
                        "name": "query",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/DomainTagListResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}": {
            "get": {
                "tags": ["domain"],
                "summary": "Get Domain",
                "operationId": "v1_get_domain",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/DomainDetailResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "delete": {
                "tags": ["domain"],
                "summary": "Delete Domain",
                "description": "TODO: finish this part\n\ntc-imba: delete domain have many side effects, and is not urgent,\n                 marked it deprecated and implement it later",
                "operationId": "v1_delete_domain",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/EmptyResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "deprecated": true,
                "security": [{ "HTTPBearer": [] }]
            },
            "patch": {
                "tags": ["domain"],
                "summary": "Update Domain",
                "operationId": "v1_update_domain",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/DomainEdit" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/DomainResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/transfer": {
            "post": {
                "tags": ["domain"],
                "summary": "Transfer Domain",
                "operationId": "v1_transfer_domain",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/DomainTransfer" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/DomainResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/users": {
            "get": {
                "tags": ["domain"],
                "summary": "List Domain Users",
                "operationId": "v1_list_domain_users",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                        "required": false,
                        "schema": {
                            "title": "Ordering",
                            "type": "string",
                            "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                            "default": ""
                        },
                        "name": "ordering",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Offset",
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 0
                        },
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Limit",
                            "maximum": 500.0,
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 100
                        },
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/UserWithDomainRoleListResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "post": {
                "tags": ["domain"],
                "summary": "Add Domain User",
                "operationId": "v1_add_domain_user",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/DomainUserAdd" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/UserDetailWithDomainRoleResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/users/{user}": {
            "get": {
                "tags": ["domain"],
                "summary": "Get Domain User",
                "operationId": "v1_get_domain_user",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "user id or 'me' or empty",
                        "required": true,
                        "schema": {
                            "title": "User",
                            "type": "string",
                            "description": "user id or 'me' or empty"
                        },
                        "name": "user",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/UserWithDomainRoleResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "delete": {
                "tags": ["domain"],
                "summary": "Remove Domain User",
                "operationId": "v1_remove_domain_user",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "user id or 'me' or empty",
                        "required": true,
                        "schema": {
                            "title": "User",
                            "type": "string",
                            "description": "user id or 'me' or empty"
                        },
                        "name": "user",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/EmptyResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "patch": {
                "tags": ["domain"],
                "summary": "Update Domain User",
                "operationId": "v1_update_domain_user",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "user id or 'me' or empty",
                        "required": true,
                        "schema": {
                            "title": "User",
                            "type": "string",
                            "description": "user id or 'me' or empty"
                        },
                        "name": "user",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/DomainUserUpdate" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/UserDetailWithDomainRoleResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/users/{user}/permission": {
            "get": {
                "tags": ["domain"],
                "summary": "Get Domain User Permission",
                "operationId": "v1_get_domain_user_permission",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "user id or 'me' or empty",
                        "required": true,
                        "schema": {
                            "title": "User",
                            "type": "string",
                            "description": "user id or 'me' or empty"
                        },
                        "name": "user",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/DomainUserPermissionResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/candidates": {
            "get": {
                "tags": ["domain"],
                "summary": "Search Domain Candidates",
                "operationId": "v1_search_domain_candidates",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "search query",
                        "required": true,
                        "schema": {
                            "title": "Query",
                            "minLength": 2,
                            "type": "string",
                            "description": "search query"
                        },
                        "name": "query",
                        "in": "query"
                    },
                    {
                        "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: username",
                        "required": false,
                        "schema": {
                            "title": "Ordering",
                            "type": "string",
                            "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: username",
                            "default": ""
                        },
                        "name": "ordering",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/UserDetailWithDomainRoleListResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/roles": {
            "get": {
                "tags": ["domain"],
                "summary": "List Domain Roles",
                "operationId": "v1_list_domain_roles",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                        "required": false,
                        "schema": {
                            "title": "Ordering",
                            "type": "string",
                            "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                            "default": ""
                        },
                        "name": "ordering",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/DomainRoleListResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "post": {
                "tags": ["domain"],
                "summary": "Create Domain Role",
                "operationId": "v1_create_domain_role",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/DomainRoleCreate" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/DomainRoleResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/roles/{role}": {
            "get": {
                "tags": ["domain"],
                "summary": "Get Domain Role",
                "operationId": "v1_get_domain_role",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "name of the domain role",
                        "required": true,
                        "schema": {
                            "title": "Role",
                            "maxLength": 256,
                            "minLength": 1,
                            "type": "string",
                            "description": "name of the domain role"
                        },
                        "name": "role",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/DomainRoleDetailResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "delete": {
                "tags": ["domain"],
                "summary": "Delete Domain Role",
                "operationId": "v1_delete_domain_role",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "name of the domain role",
                        "required": true,
                        "schema": {
                            "title": "Role",
                            "maxLength": 256,
                            "minLength": 1,
                            "type": "string",
                            "description": "name of the domain role"
                        },
                        "name": "role",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/EmptyResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "patch": {
                "tags": ["domain"],
                "summary": "Update Domain Role",
                "operationId": "v1_update_domain_role",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "name of the domain role",
                        "required": true,
                        "schema": {
                            "title": "Role",
                            "maxLength": 256,
                            "minLength": 1,
                            "type": "string",
                            "description": "name of the domain role"
                        },
                        "name": "role",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/DomainRoleEdit" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/DomainRoleResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/invitations": {
            "get": {
                "tags": ["domain"],
                "summary": "List Domain Invitations",
                "operationId": "v1_list_domain_invitations",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                        "required": false,
                        "schema": {
                            "title": "Ordering",
                            "type": "string",
                            "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                            "default": ""
                        },
                        "name": "ordering",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Offset",
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 0
                        },
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Limit",
                            "maximum": 500.0,
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 100
                        },
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/DomainInvitationListResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "post": {
                "tags": ["domain"],
                "summary": "Create Domain Invitation",
                "operationId": "v1_create_domain_invitation",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "$ref": "#/components/schemas/DomainInvitationCreate"
                            }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/DomainInvitationResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/invitations/{invitation}": {
            "get": {
                "tags": ["domain"],
                "summary": "Get Domain Invitation",
                "operationId": "v1_get_domain_invitation",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the domain invitation",
                        "required": true,
                        "schema": {
                            "title": "Invitation",
                            "type": "string",
                            "description": "url or id of the domain invitation"
                        },
                        "name": "invitation",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/DomainInvitationResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "delete": {
                "tags": ["domain"],
                "summary": "Delete Domain Invitation",
                "operationId": "v1_delete_domain_invitation",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the domain invitation",
                        "required": true,
                        "schema": {
                            "title": "Invitation",
                            "type": "string",
                            "description": "url or id of the domain invitation"
                        },
                        "name": "invitation",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/EmptyResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "patch": {
                "tags": ["domain"],
                "summary": "Update Domain Invitation",
                "operationId": "v1_update_domain_invitation",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the domain invitation",
                        "required": true,
                        "schema": {
                            "title": "Invitation",
                            "type": "string",
                            "description": "url or id of the domain invitation"
                        },
                        "name": "invitation",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/DomainInvitationEdit" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/DomainInvitationResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/join": {
            "post": {
                "tags": ["domain"],
                "summary": "Join Domain By Invitation",
                "operationId": "v1_join_domain_by_invitation",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "required": true,
                        "schema": { "title": "Invitationcode", "type": "string" },
                        "name": "invitationCode",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/UserWithDomainRoleResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/problem_sets": {
            "get": {
                "tags": ["problem set"],
                "summary": "List Problem Sets",
                "operationId": "v1_list_problem_sets",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                        "required": false,
                        "schema": {
                            "title": "Ordering",
                            "type": "string",
                            "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                            "default": ""
                        },
                        "name": "ordering",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Offset",
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 0
                        },
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Limit",
                            "maximum": 500.0,
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 100
                        },
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/ProblemSetListResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "post": {
                "tags": ["problem set"],
                "summary": "Create Problem Set",
                "operationId": "v1_create_problem_set",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/ProblemSetCreate" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/ProblemSetResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/problem_sets/{problemSet}": {
            "get": {
                "tags": ["problem set"],
                "summary": "Get Problem Set",
                "operationId": "v1_get_problem_set",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem set",
                        "required": true,
                        "schema": {
                            "title": "Problemset",
                            "type": "string",
                            "description": "url or id of the problem set"
                        },
                        "name": "problemSet",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/ProblemSetDetailResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "delete": {
                "tags": ["problem set"],
                "summary": "Delete Problem Set",
                "operationId": "v1_delete_problem_set",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem set",
                        "required": true,
                        "schema": {
                            "title": "Problemset",
                            "type": "string",
                            "description": "url or id of the problem set"
                        },
                        "name": "problemSet",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/EmptyResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "deprecated": true,
                "security": [{ "HTTPBearer": [] }]
            },
            "patch": {
                "tags": ["problem set"],
                "summary": "Update Problem Set",
                "operationId": "v1_update_problem_set",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem set",
                        "required": true,
                        "schema": {
                            "title": "Problemset",
                            "type": "string",
                            "description": "url or id of the problem set"
                        },
                        "name": "problemSet",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/ProblemSetEdit" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/ProblemSetResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/problem_sets/{problemSet}/problems": {
            "get": {
                "tags": ["problem set"],
                "summary": "List Problems In Problem Set",
                "operationId": "v1_list_problems_in_problem_set",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem set",
                        "required": true,
                        "schema": {
                            "title": "Problemset",
                            "type": "string",
                            "description": "url or id of the problem set"
                        },
                        "name": "problemSet",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/ProblemWithLatestRecordListResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "post": {
                "tags": ["problem set"],
                "summary": "Add Problem In Problem Set",
                "operationId": "v1_add_problem_in_problem_set",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem set",
                        "required": true,
                        "schema": {
                            "title": "Problemset",
                            "type": "string",
                            "description": "url or id of the problem set"
                        },
                        "name": "problemSet",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/ProblemSetAddProblem" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/ProblemSetResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/problem_sets/{problemSet}/problems/{problem}": {
            "get": {
                "tags": ["problem set"],
                "summary": "Get Problem In Problem Set",
                "operationId": "v1_get_problem_in_problem_set",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem set",
                        "required": true,
                        "schema": {
                            "title": "Problemset",
                            "type": "string",
                            "description": "url or id of the problem set"
                        },
                        "name": "problemSet",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/ProblemDetailWithLatestRecordResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "delete": {
                "tags": ["problem set"],
                "summary": "Delete Problem In Problem Set",
                "operationId": "v1_delete_problem_in_problem_set",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem set",
                        "required": true,
                        "schema": {
                            "title": "Problemset",
                            "type": "string",
                            "description": "url or id of the problem set"
                        },
                        "name": "problemSet",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/ProblemSetResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "patch": {
                "tags": ["problem set"],
                "summary": "Update Problem In Problem Set",
                "operationId": "v1_update_problem_in_problem_set",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem set",
                        "required": true,
                        "schema": {
                            "title": "Problemset",
                            "type": "string",
                            "description": "url or id of the problem set"
                        },
                        "name": "problemSet",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "$ref": "#/components/schemas/ProblemSetUpdateProblem"
                            }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/ProblemSetResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/problem_sets/{problemSet}/problems/{problem}/submit": {
            "post": {
                "tags": ["problem set"],
                "summary": "Submit Solution To Problem Set",
                "operationId": "v1_submit_solution_to_problem_set",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem set",
                        "required": true,
                        "schema": {
                            "title": "Problemset",
                            "type": "string",
                            "description": "url or id of the problem set"
                        },
                        "name": "problemSet",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "multipart/form-data": {
                            "schema": { "$ref": "#/components/schemas/ProblemSolutionSubmit" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/RecordResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/problems": {
            "get": {
                "tags": ["problem"],
                "summary": "List Problems",
                "operationId": "v1_list_problems",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                        "required": false,
                        "schema": {
                            "title": "Ordering",
                            "type": "string",
                            "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                            "default": ""
                        },
                        "name": "ordering",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Offset",
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 0
                        },
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Limit",
                            "maximum": 500.0,
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 100
                        },
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/ProblemWithLatestRecordListResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "post": {
                "tags": ["problem"],
                "summary": "Create Problem",
                "operationId": "v1_create_problem",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/ProblemCreate" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/ProblemDetailResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/problems/{problem}": {
            "get": {
                "tags": ["problem"],
                "summary": "Get Problem",
                "operationId": "v1_get_problem",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/ProblemDetailWithLatestRecordResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "post": {
                "tags": ["problem"],
                "summary": "Submit Solution To Problem",
                "operationId": "v1_submit_solution_to_problem",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "multipart/form-data": {
                            "schema": { "$ref": "#/components/schemas/ProblemSolutionSubmit" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/RecordResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "delete": {
                "tags": ["problem"],
                "summary": "Delete Problem",
                "operationId": "v1_delete_problem",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/EmptyResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "patch": {
                "tags": ["problem"],
                "summary": "Update Problem",
                "operationId": "v1_update_problem",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/ProblemEdit" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/ProblemResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/problems/clone": {
            "post": {
                "tags": ["problem"],
                "summary": "Clone Problem",
                "operationId": "v1_clone_problem",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/ProblemClone" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/ProblemListResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/problems/{problem}/configs": {
            "get": {
                "tags": ["problem config"],
                "summary": "List Problem Config Commits",
                "operationId": "v1_list_problem_config_commits",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    },
                    {
                        "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                        "required": false,
                        "schema": {
                            "title": "Ordering",
                            "type": "string",
                            "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                            "default": ""
                        },
                        "name": "ordering",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Offset",
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 0
                        },
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Limit",
                            "maximum": 500.0,
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 100
                        },
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/ProblemConfigDetailListResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "post": {
                "tags": ["problem config"],
                "summary": "Update Problem Config By Archive",
                "description": "Completely replace the problem config with the archive. This will delete files not included in the archive.",
                "operationId": "v1_update_problem_config_by_archive",
                "parameters": [
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "required": false,
                        "schema": {
                            "allOf": [{ "$ref": "#/components/schemas/ConfigMissing" }],
                            "default": "raise_error"
                        },
                        "name": "configJsonOnMissing",
                        "in": "query"
                    }
                ],
                "requestBody": {
                    "content": {
                        "multipart/form-data": {
                            "schema": { "$ref": "#/components/schemas/FileUpload" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/ProblemConfigDetailResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/problems/{problem}/configs/json": {
            "post": {
                "tags": ["problem config"],
                "summary": "Update Problem Config Json",
                "operationId": "v1_update_problem_config_json",
                "parameters": [
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/ProblemConfigJson" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/ProblemConfigDetailResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/problems/{problem}/configs/{config}": {
            "get": {
                "tags": ["problem config"],
                "summary": "Download Problem Config Archive",
                "operationId": "v1_download_problem_config_archive",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    },
                    {
                        "description": "'latest' or id of the config",
                        "required": true,
                        "schema": {
                            "title": "Config",
                            "type": "string",
                            "description": "'latest' or id of the config"
                        },
                        "name": "config",
                        "in": "path"
                    },
                    {
                        "required": false,
                        "schema": {
                            "allOf": [{ "$ref": "#/components/schemas/ArchiveType" }],
                            "default": "zip"
                        },
                        "name": "archiveFormat",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "title": "Response Download Problem Config Archive Domains    Domain    Problems    Problem    Configs    Config    Get"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/problems/{problem}/configs/{config}/json": {
            "get": {
                "tags": ["problem config"],
                "summary": "Get Problem Config Json",
                "operationId": "v1_get_problem_config_json",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "'latest' or id of the config",
                        "required": true,
                        "schema": {
                            "title": "Config",
                            "type": "string",
                            "description": "'latest' or id of the config"
                        },
                        "name": "config",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/ProblemConfigDataDetailResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/problems/{problem}/configs/latest/ls": {
            "get": {
                "tags": ["problem config"],
                "summary": "List Latest Problem Config Objects Under A Given Prefix",
                "operationId": "v1_list_latest_problem_config_objects_under_a_given_prefix",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    },
                    {
                        "description": "return items after this value",
                        "required": false,
                        "schema": {
                            "title": "After",
                            "type": "string",
                            "description": "return items after this value",
                            "default": ""
                        },
                        "name": "after",
                        "in": "query"
                    },
                    {
                        "description": "how many items to return",
                        "required": false,
                        "schema": {
                            "title": "Amount",
                            "type": "integer",
                            "description": "how many items to return",
                            "default": 100
                        },
                        "name": "amount",
                        "in": "query"
                    },
                    {
                        "description": "delimiter used to group common prefixes by",
                        "required": false,
                        "schema": {
                            "title": "Delimiter",
                            "type": "string",
                            "description": "delimiter used to group common prefixes by",
                            "default": ""
                        },
                        "name": "delimiter",
                        "in": "query"
                    },
                    {
                        "description": "return items prefixed with this value",
                        "required": false,
                        "schema": {
                            "title": "Prefix",
                            "type": "string",
                            "description": "return items prefixed with this value",
                            "default": ""
                        },
                        "name": "prefix",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/ObjectStatsListResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/problems/{problem}/configs/latest/diff": {
            "get": {
                "tags": ["problem config"],
                "summary": "Diff Problem Config Default Branch",
                "operationId": "v1_diff_problem_config_default_branch",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the problem",
                        "required": true,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "url or id of the problem"
                        },
                        "name": "problem",
                        "in": "path"
                    },
                    {
                        "description": "return items after this value",
                        "required": false,
                        "schema": {
                            "title": "After",
                            "type": "string",
                            "description": "return items after this value",
                            "default": ""
                        },
                        "name": "after",
                        "in": "query"
                    },
                    {
                        "description": "how many items to return",
                        "required": false,
                        "schema": {
                            "title": "Amount",
                            "type": "integer",
                            "description": "how many items to return",
                            "default": 100
                        },
                        "name": "amount",
                        "in": "query"
                    },
                    {
                        "description": "delimiter used to group common prefixes by",
                        "required": false,
                        "schema": {
                            "title": "Delimiter",
                            "type": "string",
                            "description": "delimiter used to group common prefixes by",
                            "default": ""
                        },
                        "name": "delimiter",
                        "in": "query"
                    },
                    {
                        "description": "return items prefixed with this value",
                        "required": false,
                        "schema": {
                            "title": "Prefix",
                            "type": "string",
                            "description": "return items prefixed with this value",
                            "default": ""
                        },
                        "name": "prefix",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/DiffListResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/problem_groups": {
            "get": {
                "tags": ["problem group"],
                "summary": "List Problem Groups",
                "operationId": "v1_list_problem_groups",
                "parameters": [
                    {
                        "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                        "required": false,
                        "schema": {
                            "title": "Ordering",
                            "type": "string",
                            "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                            "default": ""
                        },
                        "name": "ordering",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Offset",
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 0
                        },
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Limit",
                            "maximum": 500.0,
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 100
                        },
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/ProblemGroupListResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/records": {
            "get": {
                "tags": ["record"],
                "summary": "List Records In Domain",
                "operationId": "v1_list_records_in_domain",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "description": "problem set id",
                        "required": false,
                        "schema": {
                            "title": "Problemset",
                            "type": "string",
                            "description": "problem set id",
                            "format": "uuid"
                        },
                        "name": "problemSet",
                        "in": "query"
                    },
                    {
                        "description": "problem id",
                        "required": false,
                        "schema": {
                            "title": "Problem",
                            "type": "string",
                            "description": "problem id",
                            "format": "uuid"
                        },
                        "name": "problem",
                        "in": "query"
                    },
                    {
                        "description": "submitter uid",
                        "required": false,
                        "schema": {
                            "title": "Submitterid",
                            "type": "string",
                            "description": "submitter uid",
                            "format": "uuid"
                        },
                        "name": "submitterId",
                        "in": "query"
                    },
                    {
                        "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                        "required": false,
                        "schema": {
                            "title": "Ordering",
                            "type": "string",
                            "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                            "default": ""
                        },
                        "name": "ordering",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Offset",
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 0
                        },
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Limit",
                            "maximum": 500.0,
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 100
                        },
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/RecordListDetailListResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/records/{record}": {
            "get": {
                "tags": ["record"],
                "summary": "Get Record",
                "operationId": "v1_get_record",
                "parameters": [
                    {
                        "required": true,
                        "schema": { "title": "Record", "type": "string", "format": "uuid" },
                        "name": "record",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/RecordDetailResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/users/me": {
            "get": {
                "tags": ["user"],
                "summary": "Get Current User",
                "operationId": "v1_get_current_user",
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/UserDetailResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "patch": {
                "tags": ["user"],
                "summary": "Update Current User",
                "operationId": "v1_update_current_user",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/UserEdit" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/UserDetailResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/users/me/password": {
            "patch": {
                "tags": ["user"],
                "summary": "Change Password",
                "operationId": "v1_change_password",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/UserResetPassword" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/UserDetailResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/users/{uid}": {
            "get": {
                "tags": ["user"],
                "summary": "Get User",
                "operationId": "v1_get_user",
                "parameters": [
                    {
                        "required": true,
                        "schema": { "title": "Uid", "type": "string" },
                        "name": "uid",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/UserPreviewResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/auth/login": {
            "post": {
                "tags": ["auth"],
                "summary": "Login",
                "operationId": "v1_login",
                "parameters": [
                    {
                        "description": "Add Set/Delete-Cookie on response header",
                        "required": false,
                        "schema": {
                            "title": "Cookie",
                            "type": "boolean",
                            "description": "Add Set/Delete-Cookie on response header",
                            "default": true
                        },
                        "name": "cookie",
                        "in": "query"
                    },
                    {
                        "required": true,
                        "schema": {
                            "title": "Responsetype",
                            "enum": ["redirect", "json"],
                            "type": "string"
                        },
                        "name": "responseType",
                        "in": "query"
                    },
                    {
                        "description": "The redirect url after the operation",
                        "required": false,
                        "schema": {
                            "title": "Redirecturl",
                            "type": "string",
                            "description": "The redirect url after the operation"
                        },
                        "name": "redirectUrl",
                        "in": "query"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/x-www-form-urlencoded": {
                            "schema": {
                                "$ref": "#/components/schemas/OAuth2PasswordRequestForm"
                            }
                        }
                    }
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/AuthTokensResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "tags": ["auth"],
                "summary": "Logout",
                "operationId": "v1_logout",
                "parameters": [
                    {
                        "description": "Add Set/Delete-Cookie on response header",
                        "required": false,
                        "schema": {
                            "title": "Cookie",
                            "type": "boolean",
                            "description": "Add Set/Delete-Cookie on response header",
                            "default": true
                        },
                        "name": "cookie",
                        "in": "query"
                    },
                    {
                        "required": true,
                        "schema": {
                            "title": "Responsetype",
                            "enum": ["redirect", "json"],
                            "type": "string"
                        },
                        "name": "responseType",
                        "in": "query"
                    },
                    {
                        "description": "The redirect url after the operation",
                        "required": false,
                        "schema": {
                            "title": "Redirecturl",
                            "type": "string",
                            "description": "The redirect url after the operation"
                        },
                        "name": "redirectUrl",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "title": "Response Logout Auth Logout Post" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/auth/register": {
            "post": {
                "tags": ["auth"],
                "summary": "Register",
                "operationId": "v1_register",
                "parameters": [
                    {
                        "description": "Add Set/Delete-Cookie on response header",
                        "required": false,
                        "schema": {
                            "title": "Cookie",
                            "type": "boolean",
                            "description": "Add Set/Delete-Cookie on response header",
                            "default": true
                        },
                        "name": "cookie",
                        "in": "query"
                    },
                    {
                        "required": true,
                        "schema": {
                            "title": "Responsetype",
                            "enum": ["redirect", "json"],
                            "type": "string"
                        },
                        "name": "responseType",
                        "in": "query"
                    },
                    {
                        "description": "The redirect url after the operation",
                        "required": false,
                        "schema": {
                            "title": "Redirecturl",
                            "type": "string",
                            "description": "The redirect url after the operation"
                        },
                        "name": "redirectUrl",
                        "in": "query"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/UserCreate" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/AuthTokensResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/auth/token": {
            "get": {
                "tags": ["auth"],
                "summary": "Get Token",
                "operationId": "v1_get_token",
                "parameters": [
                    {
                        "description": "Add Set/Delete-Cookie on response header",
                        "required": false,
                        "schema": {
                            "title": "Cookie",
                            "type": "boolean",
                            "description": "Add Set/Delete-Cookie on response header",
                            "default": true
                        },
                        "name": "cookie",
                        "in": "query"
                    },
                    {
                        "required": true,
                        "schema": {
                            "title": "Responsetype",
                            "enum": ["redirect", "json"],
                            "type": "string"
                        },
                        "name": "responseType",
                        "in": "query"
                    },
                    {
                        "description": "The redirect url after the operation",
                        "required": false,
                        "schema": {
                            "title": "Redirecturl",
                            "type": "string",
                            "description": "The redirect url after the operation"
                        },
                        "name": "redirectUrl",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/AuthTokensResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "tags": ["auth"],
                "summary": "Refresh",
                "operationId": "v1_refresh",
                "parameters": [
                    {
                        "description": "Add Set/Delete-Cookie on response header",
                        "required": false,
                        "schema": {
                            "title": "Cookie",
                            "type": "boolean",
                            "description": "Add Set/Delete-Cookie on response header",
                            "default": true
                        },
                        "name": "cookie",
                        "in": "query"
                    },
                    {
                        "required": true,
                        "schema": {
                            "title": "Responsetype",
                            "enum": ["redirect", "json"],
                            "type": "string"
                        },
                        "name": "responseType",
                        "in": "query"
                    },
                    {
                        "description": "The redirect url after the operation",
                        "required": false,
                        "schema": {
                            "title": "Redirecturl",
                            "type": "string",
                            "description": "The redirect url after the operation"
                        },
                        "name": "redirectUrl",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/AuthTokensResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                }
            }
        },
        "/auth/oauth2": {
            "get": {
                "tags": ["auth"],
                "summary": "List Oauth2",
                "operationId": "v1_list_oauth2",
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/OAuth2ClientListResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    }
                }
            }
        },
        "/auth/oauth2/{oauth2}/authorize": {
            "get": {
                "tags": ["auth"],
                "summary": "Oauth Authorize",
                "operationId": "v1_oauth_authorize",
                "parameters": [
                    {
                        "description": "OAuth client name",
                        "required": true,
                        "schema": {
                            "title": "Oauth2",
                            "type": "string",
                            "description": "OAuth client name"
                        },
                        "name": "oauth2",
                        "in": "path"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Scopes",
                            "type": "array",
                            "items": { "type": "string" }
                        },
                        "name": "scopes",
                        "in": "query"
                    },
                    {
                        "description": "Add Set/Delete-Cookie on response header",
                        "required": false,
                        "schema": {
                            "title": "Cookie",
                            "type": "boolean",
                            "description": "Add Set/Delete-Cookie on response header",
                            "default": true
                        },
                        "name": "cookie",
                        "in": "query"
                    },
                    {
                        "required": true,
                        "schema": {
                            "title": "Responsetype",
                            "enum": ["redirect", "json"],
                            "type": "string"
                        },
                        "name": "responseType",
                        "in": "query"
                    },
                    {
                        "description": "The redirect url after the operation",
                        "required": false,
                        "schema": {
                            "title": "Redirecturl",
                            "type": "string",
                            "description": "The redirect url after the operation"
                        },
                        "name": "redirectUrl",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/RedirectResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                }
            }
        },
        "/admin/users": {
            "get": {
                "tags": ["admin"],
                "summary": "Admin List Users",
                "operationId": "v1_admin_list_users",
                "parameters": [
                    {
                        "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                        "required": false,
                        "schema": {
                            "title": "Ordering",
                            "type": "string",
                            "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                            "default": ""
                        },
                        "name": "ordering",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Offset",
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 0
                        },
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Limit",
                            "maximum": 500.0,
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 100
                        },
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/UserListResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/admin/domain_roles": {
            "get": {
                "tags": ["admin"],
                "summary": "Admin List Domain Roles",
                "operationId": "v1_admin_list_domain_roles",
                "parameters": [
                    {
                        "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                        "required": false,
                        "schema": {
                            "title": "Ordering",
                            "type": "string",
                            "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                            "default": ""
                        },
                        "name": "ordering",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Offset",
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 0
                        },
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Limit",
                            "maximum": 500.0,
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 100
                        },
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/DomainRoleListResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/admin/judgers": {
            "get": {
                "tags": ["admin"],
                "summary": "Admin List Judgers",
                "operationId": "v1_admin_list_judgers",
                "parameters": [
                    {
                        "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                        "required": false,
                        "schema": {
                            "title": "Ordering",
                            "type": "string",
                            "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                            "default": ""
                        },
                        "name": "ordering",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Offset",
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 0
                        },
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Limit",
                            "maximum": 500.0,
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 100
                        },
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/JudgerDetailListResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            },
            "post": {
                "tags": ["admin"],
                "summary": "Admin Create Judger",
                "operationId": "v1_admin_create_judger",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/JudgerCreate" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/AuthTokensWithLakefsResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/admin/{uid}": {
            "get": {
                "tags": ["admin"],
                "summary": "Admin Get User",
                "operationId": "v1_admin_get_user",
                "parameters": [
                    {
                        "required": true,
                        "schema": { "title": "Uid", "type": "string" },
                        "name": "uid",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/UserDetailResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/admin/{uid}/domains": {
            "get": {
                "tags": ["admin"],
                "summary": "Admin List User Domains",
                "operationId": "v1_admin_list_user_domains",
                "parameters": [
                    {
                        "required": true,
                        "schema": { "title": "Uid", "type": "string" },
                        "name": "uid",
                        "in": "path"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Role",
                            "type": "array",
                            "items": { "type": "string" }
                        },
                        "name": "role",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Groups",
                            "type": "array",
                            "items": { "type": "string" }
                        },
                        "name": "groups",
                        "in": "query"
                    },
                    {
                        "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                        "required": false,
                        "schema": {
                            "title": "Ordering",
                            "type": "string",
                            "description": "Comma separated list of ordering the results.\nYou may specify reverse orderings by prefixing the field name with '-'.\n\nAvailable fields: created_at,updated_at",
                            "default": ""
                        },
                        "name": "ordering",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Offset",
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 0
                        },
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "required": false,
                        "schema": {
                            "title": "Limit",
                            "maximum": 500.0,
                            "minimum": 0.0,
                            "type": "integer",
                            "default": 100
                        },
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/DomainListResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/records/{record}/judge/claim": {
            "post": {
                "tags": ["judge"],
                "summary": "Claim Record By Judger",
                "operationId": "v1_claim_record_by_judger",
                "parameters": [
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    },
                    {
                        "required": true,
                        "schema": { "title": "Record", "type": "string" },
                        "name": "record",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/JudgerClaim" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "$ref": "#/components/schemas/JudgerCredentialsResp"
                                }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/records/{record}/judge": {
            "put": {
                "tags": ["judge"],
                "summary": "Submit Record By Judger",
                "operationId": "v1_submit_record_by_judger",
                "parameters": [
                    {
                        "required": true,
                        "schema": { "title": "Record", "type": "string" },
                        "name": "record",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/RecordSubmit" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/EmptyResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/domains/{domain}/records/{record}/cases/{index}/judge": {
            "put": {
                "tags": ["judge"],
                "summary": "Submit Case By Judger",
                "operationId": "v1_submit_case_by_judger",
                "parameters": [
                    {
                        "required": true,
                        "schema": { "title": "Index", "minimum": 0.0, "type": "integer" },
                        "name": "index",
                        "in": "path"
                    },
                    {
                        "required": true,
                        "schema": { "title": "Record", "type": "string" },
                        "name": "record",
                        "in": "path"
                    },
                    {
                        "description": "url or id of the domain",
                        "required": true,
                        "schema": {
                            "title": "Domain",
                            "type": "string",
                            "description": "url or id of the domain"
                        },
                        "name": "domain",
                        "in": "path"
                    }
                ],
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": { "$ref": "#/components/schemas/RecordCaseSubmit" }
                        }
                    },
                    "required": true
                },
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/EmptyResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/HTTPValidationError" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/version": {
            "get": {
                "tags": ["miscellaneous"],
                "summary": "Version",
                "operationId": "v1_version",
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Version" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    }
                }
            }
        },
        "/jwt_decoded": {
            "get": {
                "tags": ["miscellaneous"],
                "summary": "Jwt Decoded",
                "operationId": "v1_jwt_decoded",
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/JWTAccessTokenResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    }
                },
                "security": [{ "HTTPBearer": [] }]
            }
        },
        "/test/report": {
            "get": {
                "tags": ["miscellaneous"],
                "summary": "Test Error Report",
                "operationId": "v1_test_error_report",
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/EmptyResp" }
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "content": {
                            "application/json": {
                                "schema": { "$ref": "#/components/schemas/Detail" }
                            }
                        }
                    }
                }
            }
        }
    },
    "components": {
        "schemas": {
            "ArchiveType": {
                "title": "ArchiveType",
                "enum": ["zip", "tar", "rar", "unknown"],
                "type": "string",
                "description": "An enumeration."
            },
            "AuthTokens": {
                "title": "AuthTokens",
                "required": ["accessToken", "refreshToken", "tokenType"],
                "type": "object",
                "properties": {
                    "accessToken": { "title": "Accesstoken", "type": "string" },
                    "refreshToken": { "title": "Refreshtoken", "type": "string" },
                    "tokenType": { "title": "Tokentype", "type": "string" }
                }
            },
            "AuthTokensResp": {
                "title": "AuthTokensResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/AuthTokens" }
                }
            },
            "AuthTokensWithLakefs": {
                "title": "AuthTokensWithLakefs",
                "required": [
                    "accessToken",
                    "refreshToken",
                    "tokenType",
                    "accessKeyId",
                    "secretAccessKey"
                ],
                "type": "object",
                "properties": {
                    "accessToken": { "title": "Accesstoken", "type": "string" },
                    "refreshToken": { "title": "Refreshtoken", "type": "string" },
                    "tokenType": { "title": "Tokentype", "type": "string" },
                    "accessKeyId": { "title": "Accesskeyid", "type": "string" },
                    "secretAccessKey": { "title": "Secretaccesskey", "type": "string" }
                }
            },
            "AuthTokensWithLakefsResp": {
                "title": "AuthTokensWithLakefsResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/AuthTokensWithLakefs" }
                }
            },
            "Case": {
                "title": "Case",
                "type": "object",
                "properties": {
                    "category": {
                        "title": "Category",
                        "type": "string",
                        "default": "pretest"
                    },
                    "time": { "title": "Time", "type": "string", "default": "1s" },
                    "memory": { "title": "Memory", "type": "string", "default": "64m" },
                    "score": { "title": "Score", "type": "integer", "default": 10 },
                    "ignoreWhitespace": {
                        "title": "Ignorewhitespace",
                        "type": "boolean",
                        "default": true
                    },
                    "executeFiles": {
                        "title": "Executefiles",
                        "type": "array",
                        "items": { "type": "string" },
                        "default": ["./a.out"]
                    },
                    "executeArgs": {
                        "title": "Executeargs",
                        "type": "array",
                        "items": { "type": "string" },
                        "default": ["./a.out", "--test"]
                    },
                    "executeInputFile": {
                        "title": "Executeinputfile",
                        "type": "string",
                        "default": "case.in"
                    },
                    "executeOutputFile": {
                        "title": "Executeoutputfile",
                        "type": "string",
                        "default": "case.out"
                    }
                }
            },
            "ConfigMissing": {
                "title": "ConfigMissing",
                "enum": ["use_old", "use_default", "raise_error"],
                "type": "string",
                "description": "An enumeration."
            },
            "Detail": {
                "title": "Detail",
                "required": ["detail"],
                "type": "object",
                "properties": { "detail": { "title": "Detail", "type": "string" } }
            },
            "Diff": {
                "title": "Diff",
                "required": ["type", "path", "path_type"],
                "type": "object",
                "properties": {
                    "type": { "$ref": "#/components/schemas/DiffTypeEnum" },
                    "path": { "title": "Path", "type": "string" },
                    "path_type": { "$ref": "#/components/schemas/PathTypeEnum" },
                    "size_bytes": { "title": "Size Bytes", "type": "integer" }
                }
            },
            "DiffList": {
                "title": "DiffList",
                "required": ["pagination", "results"],
                "type": "object",
                "properties": {
                    "pagination": { "$ref": "#/components/schemas/Pagination" },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/Diff" }
                    }
                }
            },
            "DiffListResp": {
                "title": "DiffListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/DiffList" }
                }
            },
            "DiffTypeEnum": {
                "title": "DiffTypeEnum",
                "enum": ["added", "removed", "changed", "conflict"],
                "type": "string",
                "description": "An enumeration."
            },
            "Domain": {
                "title": "Domain",
                "required": ["id", "name"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "url": {
                        "title": "Url",
                        "type": "string",
                        "description": "(unique) url of the domain",
                        "default": ""
                    },
                    "name": {
                        "title": "Name",
                        "maxLength": 256,
                        "minLength": 1,
                        "type": "string",
                        "description": "displayed name of the domain"
                    },
                    "gravatar": {
                        "title": "Gravatar",
                        "maxLength": 256,
                        "type": "string",
                        "description": "gravatar url of the domain",
                        "default": ""
                    },
                    "bulletin": {
                        "title": "Bulletin",
                        "maxLength": 65536,
                        "type": "string",
                        "description": "bulletin of the domain",
                        "default": ""
                    },
                    "hidden": {
                        "title": "Hidden",
                        "type": "boolean",
                        "description": "is the domain hidden",
                        "default": true
                    },
                    "group": {
                        "title": "Group",
                        "maxLength": 256,
                        "type": "string",
                        "description": "group name of the domain",
                        "default": ""
                    },
                    "ownerId": { "title": "Ownerid", "type": "string", "format": "uuid" }
                }
            },
            "DomainCreate": {
                "title": "DomainCreate",
                "required": ["name"],
                "type": "object",
                "properties": {
                    "url": {
                        "title": "Url",
                        "type": "string",
                        "description": "(unique) url of the domain",
                        "default": ""
                    },
                    "name": {
                        "title": "Name",
                        "maxLength": 256,
                        "minLength": 1,
                        "type": "string",
                        "description": "displayed name of the domain"
                    },
                    "gravatar": {
                        "title": "Gravatar",
                        "maxLength": 256,
                        "type": "string",
                        "description": "gravatar url of the domain",
                        "default": ""
                    },
                    "bulletin": {
                        "title": "Bulletin",
                        "maxLength": 65536,
                        "type": "string",
                        "description": "bulletin of the domain",
                        "default": ""
                    },
                    "hidden": {
                        "title": "Hidden",
                        "type": "boolean",
                        "description": "is the domain hidden",
                        "default": true
                    },
                    "group": {
                        "title": "Group",
                        "maxLength": 256,
                        "type": "string",
                        "description": "group name of the domain",
                        "default": ""
                    }
                }
            },
            "DomainDetail": {
                "title": "DomainDetail",
                "required": ["id", "name"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "url": {
                        "title": "Url",
                        "type": "string",
                        "description": "(unique) url of the domain",
                        "default": ""
                    },
                    "name": {
                        "title": "Name",
                        "maxLength": 256,
                        "minLength": 1,
                        "type": "string",
                        "description": "displayed name of the domain"
                    },
                    "gravatar": {
                        "title": "Gravatar",
                        "maxLength": 256,
                        "type": "string",
                        "description": "gravatar url of the domain",
                        "default": ""
                    },
                    "bulletin": {
                        "title": "Bulletin",
                        "maxLength": 65536,
                        "type": "string",
                        "description": "bulletin of the domain",
                        "default": ""
                    },
                    "hidden": {
                        "title": "Hidden",
                        "type": "boolean",
                        "description": "is the domain hidden",
                        "default": true
                    },
                    "group": {
                        "title": "Group",
                        "maxLength": 256,
                        "type": "string",
                        "description": "group name of the domain",
                        "default": ""
                    },
                    "ownerId": { "title": "Ownerid", "type": "string", "format": "uuid" },
                    "createdAt": {
                        "title": "Createdat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "updatedAt": {
                        "title": "Updatedat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "tag": { "title": "Tag", "type": "string" }
                }
            },
            "DomainDetailResp": {
                "title": "DomainDetailResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/DomainDetail" }
                }
            },
            "DomainEdit": {
                "title": "DomainEdit",
                "type": "object",
                "properties": {
                    "url": { "title": "Url", "type": "string" },
                    "name": { "title": "Name", "maxLength": 256, "type": "string" },
                    "gravatar": {
                        "title": "Gravatar",
                        "maxLength": 256,
                        "type": "string"
                    },
                    "bulletin": {
                        "title": "Bulletin",
                        "maxLength": 65536,
                        "type": "string"
                    },
                    "hidden": { "title": "Hidden", "type": "boolean" },
                    "group": { "title": "Group", "maxLength": 256, "type": "string" }
                }
            },
            "DomainInvitation": {
                "title": "DomainInvitation",
                "required": ["id", "domainId", "code"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "url": {
                        "title": "Url",
                        "type": "string",
                        "description": "(unique) url of the domain",
                        "default": ""
                    },
                    "code": {
                        "title": "Code",
                        "maxLength": 256,
                        "type": "string",
                        "description": "invitation code"
                    },
                    "expireAt": {
                        "title": "Expireat",
                        "type": "string",
                        "description": "expire time of invitation",
                        "format": "date-time"
                    },
                    "role": {
                        "title": "Role",
                        "type": "string",
                        "description": "domain role after invitation",
                        "default": "user"
                    }
                }
            },
            "DomainInvitationCreate": {
                "title": "DomainInvitationCreate",
                "required": ["code"],
                "type": "object",
                "properties": {
                    "url": {
                        "title": "Url",
                        "type": "string",
                        "description": "(unique) url of the domain",
                        "default": ""
                    },
                    "code": {
                        "title": "Code",
                        "maxLength": 256,
                        "type": "string",
                        "description": "invitation code"
                    },
                    "expireAt": {
                        "title": "Expireat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "role": {
                        "title": "Role",
                        "type": "string",
                        "description": "domain role after invitation",
                        "default": "user"
                    }
                }
            },
            "DomainInvitationEdit": {
                "title": "DomainInvitationEdit",
                "type": "object",
                "properties": {
                    "code": {
                        "title": "Code",
                        "maxLength": 256,
                        "type": "string",
                        "description": "invitation code"
                    },
                    "expireAt": {
                        "title": "Expireat",
                        "type": "string",
                        "description": "expire time of invitation",
                        "format": "date-time"
                    },
                    "role": {
                        "title": "Role",
                        "type": "string",
                        "description": "domain role after invitation"
                    }
                }
            },
            "DomainInvitationList": {
                "title": "DomainInvitationList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/DomainInvitation" },
                        "default": []
                    }
                }
            },
            "DomainInvitationListResp": {
                "title": "DomainInvitationListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/DomainInvitationList" }
                }
            },
            "DomainInvitationResp": {
                "title": "DomainInvitationResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/DomainInvitation" }
                }
            },
            "DomainList": {
                "title": "DomainList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/Domain" },
                        "default": []
                    }
                }
            },
            "DomainListResp": {
                "title": "DomainListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/DomainList" }
                }
            },
            "DomainPermission": {
                "title": "DomainPermission",
                "type": "object",
                "properties": {
                    "general": {
                        "title": "General",
                        "allOf": [{ "$ref": "#/components/schemas/GeneralPermission" }],
                        "default": {
                            "view": true,
                            "edit_permission": false,
                            "view_mod_badge": true,
                            "edit": false,
                            "unlimited_quota": false
                        }
                    },
                    "problem": {
                        "title": "Problem",
                        "allOf": [{ "$ref": "#/components/schemas/ProblemPermission" }],
                        "default": {
                            "create": false,
                            "view": true,
                            "view_hidden": false,
                            "submit": true,
                            "edit": false,
                            "view_config": false
                        }
                    },
                    "problemSet": {
                        "title": "Problemset",
                        "allOf": [{ "$ref": "#/components/schemas/ProblemSetPermission" }],
                        "default": {
                            "create": false,
                            "view": true,
                            "view_hidden": false,
                            "claim": true,
                            "scoreboard": false,
                            "manage": false,
                            "edit": false,
                            "view_config": false
                        }
                    },
                    "record": {
                        "title": "Record",
                        "allOf": [{ "$ref": "#/components/schemas/RecordPermission" }],
                        "default": {
                            "view": false,
                            "detail": false,
                            "code": false,
                            "judge": false,
                            "rejudge": false
                        }
                    }
                },
                "description": "All permissions in a domain"
            },
            "DomainResp": {
                "title": "DomainResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/Domain" }
                }
            },
            "DomainRole": {
                "title": "DomainRole",
                "required": ["id", "domainId", "role", "permission"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "role": {
                        "title": "Role",
                        "maxLength": 256,
                        "minLength": 1,
                        "type": "string"
                    },
                    "permission": { "$ref": "#/components/schemas/DomainPermission" }
                }
            },
            "DomainRoleCreate": {
                "title": "DomainRoleCreate",
                "required": ["role", "permission"],
                "type": "object",
                "properties": {
                    "role": {
                        "title": "Role",
                        "maxLength": 256,
                        "minLength": 1,
                        "type": "string"
                    },
                    "permission": { "$ref": "#/components/schemas/DomainPermission" }
                }
            },
            "DomainRoleDetail": {
                "title": "DomainRoleDetail",
                "required": ["id", "domainId", "role", "permission"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "role": {
                        "title": "Role",
                        "maxLength": 256,
                        "minLength": 1,
                        "type": "string"
                    },
                    "permission": { "$ref": "#/components/schemas/DomainPermission" },
                    "createdAt": {
                        "title": "Createdat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "updatedAt": {
                        "title": "Updatedat",
                        "type": "string",
                        "format": "date-time"
                    }
                }
            },
            "DomainRoleDetailResp": {
                "title": "DomainRoleDetailResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/DomainRoleDetail" }
                }
            },
            "DomainRoleEdit": {
                "title": "DomainRoleEdit",
                "type": "object",
                "properties": {
                    "permission": {
                        "title": "Permission",
                        "type": "object",
                        "description": "The permission which needs to be updated"
                    }
                }
            },
            "DomainRoleList": {
                "title": "DomainRoleList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/DomainRole" },
                        "default": []
                    }
                }
            },
            "DomainRoleListResp": {
                "title": "DomainRoleListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/DomainRoleList" }
                }
            },
            "DomainRoleResp": {
                "title": "DomainRoleResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/DomainRole" }
                }
            },
            "DomainTag": { "title": "DomainTag", "maxLength": 256, "type": "string" },
            "DomainTagList": {
                "title": "DomainTagList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/DomainTag" },
                        "default": []
                    }
                }
            },
            "DomainTagListResp": {
                "title": "DomainTagListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/DomainTagList" }
                }
            },
            "DomainTransfer": {
                "title": "DomainTransfer",
                "required": ["targetUser"],
                "type": "object",
                "properties": {
                    "targetUser": {
                        "title": "Targetuser",
                        "type": "string",
                        "description": "'me' or id of the user"
                    }
                }
            },
            "DomainUserAdd": {
                "title": "DomainUserAdd",
                "required": ["user"],
                "type": "object",
                "properties": {
                    "role": { "title": "Role", "type": "string", "default": "user" },
                    "user": {
                        "title": "User",
                        "type": "string",
                        "description": "'me' or id of the user"
                    }
                }
            },
            "DomainUserPermission": {
                "title": "DomainUserPermission",
                "required": ["domainId", "userId", "role", "permission"],
                "type": "object",
                "properties": {
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "userId": { "title": "Userid", "type": "string", "format": "uuid" },
                    "role": { "title": "Role", "type": "string" },
                    "permission": { "$ref": "#/components/schemas/DomainPermission" }
                }
            },
            "DomainUserPermissionResp": {
                "title": "DomainUserPermissionResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/DomainUserPermission" }
                }
            },
            "DomainUserUpdate": {
                "title": "DomainUserUpdate",
                "type": "object",
                "properties": {
                    "role": { "title": "Role", "type": "string", "default": "user" }
                }
            },
            "Empty": { "title": "Empty", "type": "object", "properties": {} },
            "EmptyResp": {
                "title": "EmptyResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/Empty" }
                }
            },
            "ErrorCode": {
                "title": "ErrorCode",
                "enum": [
                    "Success",
                    "Error",
                    "InternalServerError",
                    "UnknownFieldError",
                    "IllegalFieldError",
                    "IntegrityError",
                    "LockError",
                    "APINotImplementedError",
                    "UserRegisterError",
                    "UsernamePasswordError",
                    "UserNotFoundError",
                    "DomainNotFoundError",
                    "InvalidUrlError",
                    "ProblemNotFoundError",
                    "ProblemSetNotFoundError",
                    "ProblemGroupNotFoundError",
                    "ProblemConfigNotFoundError",
                    "ProblemConfigJsonNotFoundError",
                    "FileValidationError",
                    "FileUpdateError",
                    "FileDownloadError",
                    "UnsupportedLanguageError",
                    "RecordNotFoundError",
                    "DeleteProblemBadRequestError",
                    "UserAlreadyInDomainBadRequestError",
                    "DomainInvitationBadRequestError",
                    "DomainInvitationExpiredError",
                    "ScoreboardHiddenBadRequestError",
                    "UserNotJudgerError",
                    "DomainNotOwnerError",
                    "DomainNotRootError",
                    "DomainRoleNotFoundError",
                    "DomainRoleNotUniqueError",
                    "DomainRoleReadOnlyError",
                    "DomainRoleUsedError",
                    "DomainUserNotFoundError",
                    "FileSystemError"
                ],
                "type": "string",
                "description": "An enumeration."
            },
            "FileUpload": {
                "title": "FileUpload",
                "required": ["file"],
                "type": "object",
                "properties": {
                    "file": { "title": "File", "type": "string", "format": "binary" }
                }
            },
            "GeneralPermission": {
                "title": "GeneralPermission",
                "type": "object",
                "properties": {
                    "view": { "title": "View", "type": "boolean", "default": true },
                    "editPermission": {
                        "title": "Editpermission",
                        "type": "boolean",
                        "default": false
                    },
                    "viewModBadge": {
                        "title": "Viewmodbadge",
                        "type": "boolean",
                        "default": true
                    },
                    "edit": { "title": "Edit", "type": "boolean", "default": false },
                    "unlimitedQuota": {
                        "title": "Unlimitedquota",
                        "type": "boolean",
                        "default": false
                    }
                }
            },
            "HTTPValidationError": {
                "title": "HTTPValidationError",
                "type": "object",
                "properties": {
                    "detail": {
                        "title": "Detail",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/ValidationError" }
                    }
                }
            },
            "JWTAccessToken": {
                "title": "JWTAccessToken",
                "required": [
                    "sub",
                    "iat",
                    "nbf",
                    "jti",
                    "exp",
                    "type",
                    "category",
                    "username",
                    "isActive"
                ],
                "type": "object",
                "properties": {
                    "sub": { "title": "Sub", "type": "string" },
                    "iat": { "title": "Iat", "type": "integer" },
                    "nbf": { "title": "Nbf", "type": "integer" },
                    "jti": { "title": "Jti", "type": "string" },
                    "exp": { "title": "Exp", "type": "integer" },
                    "type": { "title": "Type", "type": "string" },
                    "fresh": { "title": "Fresh", "type": "boolean", "default": false },
                    "csrf": { "title": "Csrf", "type": "string" },
                    "category": {
                        "title": "Category",
                        "enum": ["user", "oauth"],
                        "type": "string"
                    },
                    "username": { "title": "Username", "type": "string" },
                    "gravatar": { "title": "Gravatar", "type": "string", "default": "" },
                    "role": { "title": "Role", "type": "string" },
                    "isActive": { "title": "Isactive", "type": "boolean" },
                    "oauthName": { "title": "Oauthname", "type": "string" }
                }
            },
            "JWTAccessTokenResp": {
                "title": "JWTAccessTokenResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/JWTAccessToken" }
                }
            },
            "JudgerClaim": {
                "title": "JudgerClaim",
                "required": ["taskId"],
                "type": "object",
                "properties": { "taskId": { "title": "Taskid", "type": "string" } }
            },
            "JudgerCreate": {
                "title": "JudgerCreate",
                "type": "object",
                "properties": {
                    "username": { "title": "Username", "type": "string" },
                    "password": { "title": "Password", "type": "string" },
                    "email": { "title": "Email", "type": "string" }
                }
            },
            "JudgerCredentials": {
                "title": "JudgerCredentials",
                "required": [
                    "problemConfigRepoName",
                    "problemConfigCommitId",
                    "recordRepoName",
                    "recordCommitId"
                ],
                "type": "object",
                "properties": {
                    "problemConfigRepoName": {
                        "title": "Problemconfigreponame",
                        "type": "string"
                    },
                    "problemConfigCommitId": {
                        "title": "Problemconfigcommitid",
                        "type": "string"
                    },
                    "recordRepoName": { "title": "Recordreponame", "type": "string" },
                    "recordCommitId": { "title": "Recordcommitid", "type": "string" }
                }
            },
            "JudgerCredentialsResp": {
                "title": "JudgerCredentialsResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/JudgerCredentials" }
                }
            },
            "JudgerDetail": {
                "title": "JudgerDetail",
                "required": ["id", "username", "isAlive"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "username": { "title": "Username", "type": "string" },
                    "gravatar": { "title": "Gravatar", "type": "string", "default": "" },
                    "role": { "title": "Role", "type": "string", "default": "user" },
                    "isActive": {
                        "title": "Isactive",
                        "type": "boolean",
                        "default": false
                    },
                    "createdAt": {
                        "title": "Createdat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "updatedAt": {
                        "title": "Updatedat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "isAlive": { "title": "Isalive", "type": "boolean" }
                }
            },
            "JudgerDetailList": {
                "title": "JudgerDetailList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/JudgerDetail" },
                        "default": []
                    }
                }
            },
            "JudgerDetailListResp": {
                "title": "JudgerDetailListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/JudgerDetailList" }
                }
            },
            "Language": {
                "title": "Language",
                "type": "object",
                "properties": {
                    "compileFiles": {
                        "title": "Compilefiles",
                        "type": "array",
                        "items": { "type": "string" },
                        "default": ["a.cpp"]
                    },
                    "compileArgs": {
                        "title": "Compileargs",
                        "type": "array",
                        "items": { "type": "string" },
                        "default": ["gcc", "a.cpp"]
                    },
                    "caseDefault": { "$ref": "#/components/schemas/Case" },
                    "cases": {
                        "title": "Cases",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/Case" }
                    },
                    "name": { "title": "Name", "type": "string", "default": "c++" }
                }
            },
            "LanguageDefault": {
                "title": "LanguageDefault",
                "type": "object",
                "properties": {
                    "compileFiles": {
                        "title": "Compilefiles",
                        "type": "array",
                        "items": { "type": "string" },
                        "default": ["a.cpp"]
                    },
                    "compileArgs": {
                        "title": "Compileargs",
                        "type": "array",
                        "items": { "type": "string" },
                        "default": ["gcc", "a.cpp"]
                    },
                    "caseDefault": { "$ref": "#/components/schemas/Case" },
                    "cases": {
                        "title": "Cases",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/Case" }
                    }
                }
            },
            "OAuth2Client": {
                "title": "OAuth2Client",
                "required": ["oauthName", "displayName", "icon"],
                "type": "object",
                "properties": {
                    "oauthName": { "title": "Oauthname", "type": "string" },
                    "displayName": { "title": "Displayname", "type": "string" },
                    "icon": { "title": "Icon", "type": "string" }
                }
            },
            "OAuth2ClientList": {
                "title": "OAuth2ClientList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/OAuth2Client" },
                        "default": []
                    }
                }
            },
            "OAuth2ClientListResp": {
                "title": "OAuth2ClientListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/OAuth2ClientList" }
                }
            },
            "OAuth2PasswordRequestForm": {
                "title": "OAuth2PasswordRequestForm",
                "type": "object",
                "properties": {
                    "grantType": {
                        "title": "Granttype",
                        "pattern": "password",
                        "type": "string"
                    },
                    "username": { "title": "Username", "type": "string" },
                    "password": { "title": "Password", "type": "string" },
                    "scope": { "title": "Scope", "type": "string", "default": "" },
                    "clientId": { "title": "Clientid", "type": "string" },
                    "clientSecret": { "title": "Clientsecret", "type": "string" }
                }
            },
            "ObjectStats": {
                "title": "ObjectStats",
                "required": ["path", "path_type", "checksum", "mtime"],
                "type": "object",
                "properties": {
                    "path": { "title": "Path", "type": "string" },
                    "path_type": { "$ref": "#/components/schemas/PathTypeEnum" },
                    "checksum": { "title": "Checksum", "type": "string" },
                    "size_bytes": { "title": "Size Bytes", "type": "integer" },
                    "mtime": { "title": "Mtime", "type": "integer" },
                    "content_type": { "title": "Content Type", "type": "string" }
                }
            },
            "ObjectStatsList": {
                "title": "ObjectStatsList",
                "required": ["pagination", "results"],
                "type": "object",
                "properties": {
                    "pagination": { "$ref": "#/components/schemas/Pagination" },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/ObjectStats" }
                    }
                }
            },
            "ObjectStatsListResp": {
                "title": "ObjectStatsListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/ObjectStatsList" }
                }
            },
            "Pagination": {
                "title": "Pagination",
                "required": ["has_more", "next_offset", "results", "max_per_page"],
                "type": "object",
                "properties": {
                    "has_more": { "title": "Has More", "type": "boolean" },
                    "next_offset": { "title": "Next Offset", "type": "string" },
                    "results": { "title": "Results", "minimum": 0.0, "type": "integer" },
                    "max_per_page": {
                        "title": "Max Per Page",
                        "minimum": 0.0,
                        "type": "integer"
                    }
                }
            },
            "PathTypeEnum": {
                "title": "PathTypeEnum",
                "enum": ["common_prefix", "object"],
                "type": "string",
                "description": "An enumeration."
            },
            "Problem": {
                "title": "Problem",
                "required": ["id", "domainId", "title"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "url": {
                        "title": "Url",
                        "type": "string",
                        "description": "(unique) url of the domain",
                        "default": ""
                    },
                    "title": {
                        "title": "Title",
                        "minLength": 1,
                        "type": "string",
                        "description": "title of the problem"
                    },
                    "hidden": {
                        "title": "Hidden",
                        "type": "boolean",
                        "description": "is the problem hidden",
                        "default": false
                    },
                    "numSubmit": {
                        "title": "Numsubmit",
                        "type": "integer",
                        "default": 0
                    },
                    "numAccept": {
                        "title": "Numaccept",
                        "type": "integer",
                        "default": 0
                    },
                    "ownerId": { "title": "Ownerid", "type": "string", "format": "uuid" },
                    "problemGroupId": {
                        "title": "Problemgroupid",
                        "type": "string",
                        "format": "uuid"
                    }
                }
            },
            "ProblemClone": {
                "title": "ProblemClone",
                "required": ["fromDomain", "problems"],
                "type": "object",
                "properties": {
                    "fromDomain": {
                        "title": "Fromdomain",
                        "type": "string",
                        "description": "url or id of the domain to clone from"
                    },
                    "problems": {
                        "title": "Problems",
                        "type": "array",
                        "items": { "type": "string" }
                    },
                    "newGroup": {
                        "title": "Newgroup",
                        "type": "boolean",
                        "description": "whether to create new problem group",
                        "default": false
                    }
                }
            },
            "ProblemConfigDataDetail": {
                "title": "ProblemConfigDataDetail",
                "required": ["id", "data"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "commitMessage": {
                        "title": "Commitmessage",
                        "type": "string",
                        "default": ""
                    },
                    "dataVersion": {
                        "title": "Dataversion",
                        "type": "integer",
                        "default": 2
                    },
                    "commitId": { "title": "Commitid", "type": "string", "default": "" },
                    "committerId": {
                        "title": "Committerid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "createdAt": {
                        "title": "Createdat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "updatedAt": {
                        "title": "Updatedat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "data": { "$ref": "#/components/schemas/ProblemConfigJson" }
                }
            },
            "ProblemConfigDataDetailResp": {
                "title": "ProblemConfigDataDetailResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/ProblemConfigDataDetail" }
                }
            },
            "ProblemConfigDetail": {
                "title": "ProblemConfigDetail",
                "required": ["id"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "commitMessage": {
                        "title": "Commitmessage",
                        "type": "string",
                        "default": ""
                    },
                    "dataVersion": {
                        "title": "Dataversion",
                        "type": "integer",
                        "default": 2
                    },
                    "commitId": { "title": "Commitid", "type": "string", "default": "" },
                    "committerId": {
                        "title": "Committerid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "createdAt": {
                        "title": "Createdat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "updatedAt": {
                        "title": "Updatedat",
                        "type": "string",
                        "format": "date-time"
                    }
                }
            },
            "ProblemConfigDetailList": {
                "title": "ProblemConfigDetailList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/ProblemConfigDetail" },
                        "default": []
                    }
                }
            },
            "ProblemConfigDetailListResp": {
                "title": "ProblemConfigDetailListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/ProblemConfigDetailList" }
                }
            },
            "ProblemConfigDetailResp": {
                "title": "ProblemConfigDetailResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/ProblemConfigDetail" }
                }
            },
            "ProblemConfigJson": {
                "title": "ProblemConfigJson",
                "required": ["languages"],
                "type": "object",
                "properties": {
                    "languages": {
                        "title": "Languages",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/Language" }
                    },
                    "languageDefault": { "$ref": "#/components/schemas/LanguageDefault" }
                }
            },
            "ProblemCreate": {
                "title": "ProblemCreate",
                "required": ["title"],
                "type": "object",
                "properties": {
                    "url": {
                        "title": "Url",
                        "type": "string",
                        "description": "(unique) url of the domain",
                        "default": ""
                    },
                    "title": {
                        "title": "Title",
                        "minLength": 1,
                        "type": "string",
                        "description": "title of the problem"
                    },
                    "hidden": {
                        "title": "Hidden",
                        "type": "boolean",
                        "description": "is the problem hidden",
                        "default": false
                    },
                    "content": {
                        "title": "Content",
                        "maxLength": 65536,
                        "type": "string",
                        "description": "content of the problem",
                        "default": ""
                    }
                }
            },
            "ProblemDetail": {
                "title": "ProblemDetail",
                "required": ["id", "domainId", "title"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "url": {
                        "title": "Url",
                        "type": "string",
                        "description": "(unique) url of the domain",
                        "default": ""
                    },
                    "title": {
                        "title": "Title",
                        "minLength": 1,
                        "type": "string",
                        "description": "title of the problem"
                    },
                    "hidden": {
                        "title": "Hidden",
                        "type": "boolean",
                        "description": "is the problem hidden",
                        "default": false
                    },
                    "numSubmit": {
                        "title": "Numsubmit",
                        "type": "integer",
                        "default": 0
                    },
                    "numAccept": {
                        "title": "Numaccept",
                        "type": "integer",
                        "default": 0
                    },
                    "ownerId": { "title": "Ownerid", "type": "string", "format": "uuid" },
                    "problemGroupId": {
                        "title": "Problemgroupid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "content": {
                        "title": "Content",
                        "maxLength": 65536,
                        "type": "string",
                        "description": "content of the problem",
                        "default": ""
                    },
                    "createdAt": {
                        "title": "Createdat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "updatedAt": {
                        "title": "Updatedat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "languages": {
                        "title": "Languages",
                        "type": "array",
                        "items": { "type": "string" },
                        "default": []
                    }
                }
            },
            "ProblemDetailResp": {
                "title": "ProblemDetailResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/ProblemDetail" }
                }
            },
            "ProblemDetailWithLatestRecord": {
                "title": "ProblemDetailWithLatestRecord",
                "required": ["id", "domainId", "title"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "url": {
                        "title": "Url",
                        "type": "string",
                        "description": "(unique) url of the domain",
                        "default": ""
                    },
                    "title": {
                        "title": "Title",
                        "minLength": 1,
                        "type": "string",
                        "description": "title of the problem"
                    },
                    "hidden": {
                        "title": "Hidden",
                        "type": "boolean",
                        "description": "is the problem hidden",
                        "default": false
                    },
                    "numSubmit": {
                        "title": "Numsubmit",
                        "type": "integer",
                        "default": 0
                    },
                    "numAccept": {
                        "title": "Numaccept",
                        "type": "integer",
                        "default": 0
                    },
                    "ownerId": { "title": "Ownerid", "type": "string", "format": "uuid" },
                    "problemGroupId": {
                        "title": "Problemgroupid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "content": {
                        "title": "Content",
                        "maxLength": 65536,
                        "type": "string",
                        "description": "content of the problem",
                        "default": ""
                    },
                    "createdAt": {
                        "title": "Createdat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "updatedAt": {
                        "title": "Updatedat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "languages": {
                        "title": "Languages",
                        "type": "array",
                        "items": { "type": "string" },
                        "default": []
                    },
                    "latestRecord": { "$ref": "#/components/schemas/RecordPreview" }
                }
            },
            "ProblemDetailWithLatestRecordResp": {
                "title": "ProblemDetailWithLatestRecordResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": {
                        "$ref": "#/components/schemas/ProblemDetailWithLatestRecord"
                    }
                }
            },
            "ProblemEdit": {
                "title": "ProblemEdit",
                "type": "object",
                "properties": {
                    "url": { "title": "Url", "type": "string" },
                    "title": { "title": "Title", "minLength": 1, "type": "string" },
                    "content": {
                        "title": "Content",
                        "maxLength": 65536,
                        "type": "string"
                    },
                    "hidden": { "title": "Hidden", "type": "boolean" }
                }
            },
            "ProblemGroup": {
                "title": "ProblemGroup",
                "required": ["id"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" }
                }
            },
            "ProblemGroupList": {
                "title": "ProblemGroupList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/ProblemGroup" },
                        "default": []
                    }
                }
            },
            "ProblemGroupListResp": {
                "title": "ProblemGroupListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/ProblemGroupList" }
                }
            },
            "ProblemList": {
                "title": "ProblemList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/Problem" },
                        "default": []
                    }
                }
            },
            "ProblemListResp": {
                "title": "ProblemListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/ProblemList" }
                }
            },
            "ProblemPermission": {
                "title": "ProblemPermission",
                "type": "object",
                "properties": {
                    "create": { "title": "Create", "type": "boolean", "default": false },
                    "view": { "title": "View", "type": "boolean", "default": true },
                    "viewHidden": {
                        "title": "Viewhidden",
                        "type": "boolean",
                        "default": false
                    },
                    "submit": { "title": "Submit", "type": "boolean", "default": true },
                    "edit": { "title": "Edit", "type": "boolean", "default": false },
                    "viewConfig": {
                        "title": "Viewconfig",
                        "type": "boolean",
                        "default": false
                    }
                }
            },
            "ProblemPreviewWithLatestRecord": {
                "title": "ProblemPreviewWithLatestRecord",
                "required": ["id", "title"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "url": {
                        "title": "Url",
                        "type": "string",
                        "description": "(unique) url of the domain",
                        "default": ""
                    },
                    "title": {
                        "title": "Title",
                        "minLength": 1,
                        "type": "string",
                        "description": "title of the problem"
                    },
                    "hidden": {
                        "title": "Hidden",
                        "type": "boolean",
                        "description": "is the problem hidden",
                        "default": false
                    },
                    "ownerId": { "title": "Ownerid", "type": "string", "format": "uuid" },
                    "latestRecord": { "$ref": "#/components/schemas/RecordPreview" }
                }
            },
            "ProblemResp": {
                "title": "ProblemResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/Problem" }
                }
            },
            "ProblemSet": {
                "title": "ProblemSet",
                "required": ["id", "domainId", "title"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "url": {
                        "title": "Url",
                        "type": "string",
                        "description": "(unique) url of the domain",
                        "default": ""
                    },
                    "title": {
                        "title": "Title",
                        "maxLength": 256,
                        "minLength": 1,
                        "type": "string",
                        "description": "title of the problem set"
                    },
                    "content": {
                        "title": "Content",
                        "maxLength": 65536,
                        "type": "string",
                        "description": "content of the problem set",
                        "default": ""
                    },
                    "hidden": {
                        "title": "Hidden",
                        "type": "boolean",
                        "description": "whether the problem set is hidden",
                        "default": false
                    },
                    "scoreboardHidden": {
                        "title": "Scoreboardhidden",
                        "type": "boolean",
                        "description": "whether the scoreboard of the problem set is hidden",
                        "default": false
                    },
                    "dueAt": {
                        "title": "Dueat",
                        "type": "string",
                        "description": "the problem set is due at this date",
                        "format": "date-time"
                    },
                    "lockAt": {
                        "title": "Lockat",
                        "type": "string",
                        "description": "the problem set is locked after this date",
                        "format": "date-time"
                    },
                    "unlockAt": {
                        "title": "Unlockat",
                        "type": "string",
                        "description": "the problem set is unlocked after this date",
                        "format": "date-time"
                    },
                    "numSubmit": {
                        "title": "Numsubmit",
                        "type": "integer",
                        "default": 0
                    },
                    "numAccept": {
                        "title": "Numaccept",
                        "type": "integer",
                        "default": 0
                    },
                    "ownerId": { "title": "Ownerid", "type": "string", "format": "uuid" }
                }
            },
            "ProblemSetAddProblem": {
                "title": "ProblemSetAddProblem",
                "required": ["problem"],
                "type": "object",
                "properties": {
                    "position": {
                        "title": "Position",
                        "type": "integer",
                        "description": "the position of the problem in the problem set. if None, append to the end of the problems list."
                    },
                    "problem": {
                        "title": "Problem",
                        "type": "string",
                        "description": "url or id of the problem"
                    }
                }
            },
            "ProblemSetCreate": {
                "title": "ProblemSetCreate",
                "required": ["title"],
                "type": "object",
                "properties": {
                    "url": {
                        "title": "Url",
                        "type": "string",
                        "description": "(unique) url of the domain",
                        "default": ""
                    },
                    "title": {
                        "title": "Title",
                        "maxLength": 256,
                        "minLength": 1,
                        "type": "string",
                        "description": "title of the problem set"
                    },
                    "content": {
                        "title": "Content",
                        "maxLength": 65536,
                        "type": "string",
                        "description": "content of the problem set",
                        "default": ""
                    },
                    "hidden": {
                        "title": "Hidden",
                        "type": "boolean",
                        "description": "whether the problem set is hidden",
                        "default": false
                    },
                    "scoreboardHidden": {
                        "title": "Scoreboardhidden",
                        "type": "boolean",
                        "description": "whether the scoreboard of the problem set is hidden",
                        "default": false
                    },
                    "dueAt": {
                        "title": "Dueat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "lockAt": {
                        "title": "Lockat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "unlockAt": {
                        "title": "Unlockat",
                        "type": "string",
                        "format": "date-time"
                    }
                }
            },
            "ProblemSetDetail": {
                "title": "ProblemSetDetail",
                "required": ["id", "domainId", "title"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "url": {
                        "title": "Url",
                        "type": "string",
                        "description": "(unique) url of the domain",
                        "default": ""
                    },
                    "title": {
                        "title": "Title",
                        "maxLength": 256,
                        "minLength": 1,
                        "type": "string",
                        "description": "title of the problem set"
                    },
                    "content": {
                        "title": "Content",
                        "maxLength": 65536,
                        "type": "string",
                        "description": "content of the problem set",
                        "default": ""
                    },
                    "hidden": {
                        "title": "Hidden",
                        "type": "boolean",
                        "description": "whether the problem set is hidden",
                        "default": false
                    },
                    "scoreboardHidden": {
                        "title": "Scoreboardhidden",
                        "type": "boolean",
                        "description": "whether the scoreboard of the problem set is hidden",
                        "default": false
                    },
                    "dueAt": {
                        "title": "Dueat",
                        "type": "string",
                        "description": "the problem set is due at this date",
                        "format": "date-time"
                    },
                    "lockAt": {
                        "title": "Lockat",
                        "type": "string",
                        "description": "the problem set is locked after this date",
                        "format": "date-time"
                    },
                    "unlockAt": {
                        "title": "Unlockat",
                        "type": "string",
                        "description": "the problem set is unlocked after this date",
                        "format": "date-time"
                    },
                    "numSubmit": {
                        "title": "Numsubmit",
                        "type": "integer",
                        "default": 0
                    },
                    "numAccept": {
                        "title": "Numaccept",
                        "type": "integer",
                        "default": 0
                    },
                    "ownerId": { "title": "Ownerid", "type": "string", "format": "uuid" },
                    "createdAt": {
                        "title": "Createdat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "updatedAt": {
                        "title": "Updatedat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "problems": {
                        "title": "Problems",
                        "type": "array",
                        "items": {
                            "$ref": "#/components/schemas/ProblemPreviewWithLatestRecord"
                        },
                        "default": []
                    }
                }
            },
            "ProblemSetDetailResp": {
                "title": "ProblemSetDetailResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/ProblemSetDetail" }
                }
            },
            "ProblemSetEdit": {
                "title": "ProblemSetEdit",
                "type": "object",
                "properties": {
                    "url": { "title": "Url", "type": "string" },
                    "title": {
                        "title": "Title",
                        "maxLength": 256,
                        "minLength": 1,
                        "type": "string"
                    },
                    "content": {
                        "title": "Content",
                        "maxLength": 65536,
                        "type": "string"
                    },
                    "hidden": { "title": "Hidden", "type": "boolean" },
                    "scoreboardHidden": {
                        "title": "Scoreboardhidden",
                        "type": "boolean"
                    },
                    "dueAt": {
                        "title": "Dueat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "lockAt": {
                        "title": "Lockat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "unlockAt": {
                        "title": "Unlockat",
                        "type": "string",
                        "format": "date-time"
                    }
                }
            },
            "ProblemSetList": {
                "title": "ProblemSetList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/ProblemSet" },
                        "default": []
                    }
                }
            },
            "ProblemSetListResp": {
                "title": "ProblemSetListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/ProblemSetList" }
                }
            },
            "ProblemSetPermission": {
                "title": "ProblemSetPermission",
                "type": "object",
                "properties": {
                    "create": { "title": "Create", "type": "boolean", "default": false },
                    "view": { "title": "View", "type": "boolean", "default": true },
                    "viewHidden": {
                        "title": "Viewhidden",
                        "type": "boolean",
                        "default": false
                    },
                    "claim": { "title": "Claim", "type": "boolean", "default": true },
                    "scoreboard": {
                        "title": "Scoreboard",
                        "type": "boolean",
                        "default": false
                    },
                    "manage": { "title": "Manage", "type": "boolean", "default": false },
                    "edit": { "title": "Edit", "type": "boolean", "default": false },
                    "viewConfig": {
                        "title": "Viewconfig",
                        "type": "boolean",
                        "default": false
                    }
                }
            },
            "ProblemSetResp": {
                "title": "ProblemSetResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/ProblemSet" }
                }
            },
            "ProblemSetUpdateProblem": {
                "title": "ProblemSetUpdateProblem",
                "type": "object",
                "properties": {
                    "position": {
                        "title": "Position",
                        "type": "integer",
                        "description": "the position of the problem in the problem set. if None, append to the end of the problems list."
                    }
                }
            },
            "ProblemSolutionSubmit": {
                "title": "ProblemSolutionSubmit",
                "required": ["language", "files"],
                "type": "object",
                "properties": {
                    "language": {
                        "title": "Language",
                        "maxLength": 256,
                        "type": "string"
                    },
                    "files": {
                        "title": "Files",
                        "type": "array",
                        "items": { "type": "string", "format": "binary" }
                    }
                }
            },
            "ProblemWithLatestRecord": {
                "title": "ProblemWithLatestRecord",
                "required": ["id", "domainId", "title"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "url": {
                        "title": "Url",
                        "type": "string",
                        "description": "(unique) url of the domain",
                        "default": ""
                    },
                    "title": {
                        "title": "Title",
                        "minLength": 1,
                        "type": "string",
                        "description": "title of the problem"
                    },
                    "hidden": {
                        "title": "Hidden",
                        "type": "boolean",
                        "description": "is the problem hidden",
                        "default": false
                    },
                    "numSubmit": {
                        "title": "Numsubmit",
                        "type": "integer",
                        "default": 0
                    },
                    "numAccept": {
                        "title": "Numaccept",
                        "type": "integer",
                        "default": 0
                    },
                    "ownerId": { "title": "Ownerid", "type": "string", "format": "uuid" },
                    "problemGroupId": {
                        "title": "Problemgroupid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "latestRecord": { "$ref": "#/components/schemas/RecordPreview" }
                }
            },
            "ProblemWithLatestRecordList": {
                "title": "ProblemWithLatestRecordList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/ProblemWithLatestRecord" },
                        "default": []
                    }
                }
            },
            "ProblemWithLatestRecordListResp": {
                "title": "ProblemWithLatestRecordListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/ProblemWithLatestRecordList" }
                }
            },
            "Record": {
                "title": "Record",
                "required": ["id", "domainId", "language"],
                "type": "object",
                "properties": {
                    "createdAt": {
                        "title": "Createdat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "updatedAt": {
                        "title": "Updatedat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "state": {
                        "allOf": [{ "$ref": "#/components/schemas/RecordState" }],
                        "default": "processing"
                    },
                    "language": { "title": "Language", "type": "string" },
                    "score": { "title": "Score", "type": "integer", "default": 0 },
                    "timeMs": { "title": "Timems", "type": "integer", "default": 0 },
                    "memoryKb": { "title": "Memorykb", "type": "integer", "default": 0 },
                    "judgedAt": {
                        "title": "Judgedat",
                        "type": "string",
                        "format": "date-time"
                    }
                }
            },
            "RecordCase": {
                "title": "RecordCase",
                "type": "object",
                "properties": {
                    "state": {
                        "allOf": [{ "$ref": "#/components/schemas/RecordCaseResult" }],
                        "default": "etc"
                    },
                    "score": { "title": "Score", "type": "integer", "default": 0 },
                    "timeMs": { "title": "Timems", "type": "integer", "default": 0 },
                    "memoryKb": { "title": "Memorykb", "type": "integer", "default": 0 },
                    "returnCode": {
                        "title": "Returncode",
                        "type": "integer",
                        "default": 0
                    },
                    "stdout": { "title": "Stdout", "type": "string", "default": "" },
                    "stderr": { "title": "Stderr", "type": "string", "default": "" }
                }
            },
            "RecordCaseResult": {
                "title": "RecordCaseResult",
                "enum": [
                    "accepted",
                    "wrong_answer",
                    "time_limit_exceeded",
                    "memory_limit_exceeded",
                    "output_limit_exceeded",
                    "runtime_error",
                    "compile_error",
                    "system_error",
                    "canceled",
                    "etc"
                ],
                "type": "string",
                "description": "An enumeration."
            },
            "RecordCaseSubmit": {
                "title": "RecordCaseSubmit",
                "type": "object",
                "properties": {
                    "state": { "$ref": "#/components/schemas/RecordCaseResult" },
                    "score": { "title": "Score", "type": "integer" },
                    "timeMs": { "title": "Timems", "type": "integer" },
                    "memoryKb": { "title": "Memorykb", "type": "integer" },
                    "returnCode": { "title": "Returncode", "type": "integer" },
                    "stdout": { "title": "Stdout", "type": "string" },
                    "stderr": { "title": "Stderr", "type": "string" }
                }
            },
            "RecordDetail": {
                "title": "RecordDetail",
                "required": ["id", "domainId", "language"],
                "type": "object",
                "properties": {
                    "createdAt": {
                        "title": "Createdat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "updatedAt": {
                        "title": "Updatedat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "state": {
                        "allOf": [{ "$ref": "#/components/schemas/RecordState" }],
                        "default": "processing"
                    },
                    "language": { "title": "Language", "type": "string" },
                    "score": { "title": "Score", "type": "integer", "default": 0 },
                    "timeMs": { "title": "Timems", "type": "integer", "default": 0 },
                    "memoryKb": { "title": "Memorykb", "type": "integer", "default": 0 },
                    "judgedAt": {
                        "title": "Judgedat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "commitId": { "title": "Commitid", "type": "string" },
                    "taskId": { "title": "Taskid", "type": "string", "format": "uuid" },
                    "cases": {
                        "title": "Cases",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/RecordCase" },
                        "default": []
                    },
                    "problemSetId": {
                        "title": "Problemsetid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "problemId": {
                        "title": "Problemid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "problemConfigId": {
                        "title": "Problemconfigid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "committerId": {
                        "title": "Committerid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "judgerId": {
                        "title": "Judgerid",
                        "type": "string",
                        "format": "uuid"
                    }
                }
            },
            "RecordDetailResp": {
                "title": "RecordDetailResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/RecordDetail" }
                }
            },
            "RecordListDetail": {
                "title": "RecordListDetail",
                "required": ["id", "domainId", "language"],
                "type": "object",
                "properties": {
                    "createdAt": {
                        "title": "Createdat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "updatedAt": {
                        "title": "Updatedat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "state": {
                        "allOf": [{ "$ref": "#/components/schemas/RecordState" }],
                        "default": "processing"
                    },
                    "language": { "title": "Language", "type": "string" },
                    "score": { "title": "Score", "type": "integer", "default": 0 },
                    "timeMs": { "title": "Timems", "type": "integer", "default": 0 },
                    "memoryKb": { "title": "Memorykb", "type": "integer", "default": 0 },
                    "judgedAt": {
                        "title": "Judgedat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "problemSetId": {
                        "title": "Problemsetid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "problemSetTitle": { "title": "Problemsettitle", "type": "string" },
                    "problemId": {
                        "title": "Problemid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "problemTitle": { "title": "Problemtitle", "type": "string" },
                    "committerId": {
                        "title": "Committerid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "committerUsername": {
                        "title": "Committerusername",
                        "type": "string"
                    }
                }
            },
            "RecordListDetailList": {
                "title": "RecordListDetailList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/RecordListDetail" },
                        "default": []
                    }
                }
            },
            "RecordListDetailListResp": {
                "title": "RecordListDetailListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/RecordListDetailList" }
                }
            },
            "RecordPermission": {
                "title": "RecordPermission",
                "type": "object",
                "properties": {
                    "view": { "title": "View", "type": "boolean", "default": false },
                    "detail": { "title": "Detail", "type": "boolean", "default": false },
                    "code": { "title": "Code", "type": "boolean", "default": false },
                    "judge": { "title": "Judge", "type": "boolean", "default": false },
                    "rejudge": { "title": "Rejudge", "type": "boolean", "default": false }
                }
            },
            "RecordPreview": {
                "title": "RecordPreview",
                "required": ["id", "state", "createdAt"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "state": { "$ref": "#/components/schemas/RecordState" },
                    "createdAt": {
                        "title": "Createdat",
                        "type": "string",
                        "format": "date-time"
                    }
                }
            },
            "RecordResp": {
                "title": "RecordResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/Record" }
                }
            },
            "RecordState": {
                "title": "RecordState",
                "enum": [
                    "processing",
                    "queueing",
                    "retrying",
                    "fetched",
                    "compiling",
                    "running",
                    "judging",
                    "accepted",
                    "rejected",
                    "failed"
                ],
                "type": "string",
                "description": "An enumeration."
            },
            "RecordSubmit": {
                "title": "RecordSubmit",
                "type": "object",
                "properties": {
                    "state": { "$ref": "#/components/schemas/RecordState" },
                    "score": { "title": "Score", "type": "integer" },
                    "timeMs": { "title": "Timems", "type": "integer" },
                    "memoryKb": { "title": "Memorykb", "type": "integer" },
                    "judgedAt": {
                        "title": "Judgedat",
                        "type": "string",
                        "format": "date-time"
                    }
                }
            },
            "Redirect": {
                "title": "Redirect",
                "required": ["redirectUrl"],
                "type": "object",
                "properties": {
                    "redirectUrl": { "title": "Redirecturl", "type": "string" }
                }
            },
            "RedirectResp": {
                "title": "RedirectResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/Redirect" }
                }
            },
            "User": {
                "title": "User",
                "required": ["id", "username"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "username": { "title": "Username", "type": "string" },
                    "gravatar": { "title": "Gravatar", "type": "string", "default": "" },
                    "role": { "title": "Role", "type": "string", "default": "user" },
                    "isActive": {
                        "title": "Isactive",
                        "type": "boolean",
                        "default": false
                    }
                }
            },
            "UserCreate": {
                "title": "UserCreate",
                "type": "object",
                "properties": {
                    "username": { "title": "Username", "type": "string" },
                    "password": { "title": "Password", "type": "string" },
                    "email": { "title": "Email", "type": "string" },
                    "oauth_name": { "title": "Oauth Name", "type": "string" },
                    "oauth_account_id": { "title": "Oauth Account Id", "type": "string" }
                }
            },
            "UserDetail": {
                "title": "UserDetail",
                "required": [
                    "id",
                    "username",
                    "email",
                    "registerIp",
                    "loginAt",
                    "loginIp"
                ],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "username": { "title": "Username", "type": "string" },
                    "gravatar": { "title": "Gravatar", "type": "string", "default": "" },
                    "role": { "title": "Role", "type": "string", "default": "user" },
                    "isActive": {
                        "title": "Isactive",
                        "type": "boolean",
                        "default": false
                    },
                    "createdAt": {
                        "title": "Createdat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "updatedAt": {
                        "title": "Updatedat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "email": { "title": "Email", "type": "string", "format": "email" },
                    "studentId": {
                        "title": "Studentid",
                        "type": "string",
                        "default": ""
                    },
                    "realName": { "title": "Realname", "type": "string", "default": "" },
                    "registerIp": { "title": "Registerip", "type": "string" },
                    "loginAt": {
                        "title": "Loginat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "loginIp": { "title": "Loginip", "type": "string" }
                }
            },
            "UserDetailResp": {
                "title": "UserDetailResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/UserDetail" }
                }
            },
            "UserDetailWithDomainRole": {
                "title": "UserDetailWithDomainRole",
                "required": [
                    "id",
                    "username",
                    "email",
                    "registerIp",
                    "loginAt",
                    "loginIp"
                ],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "username": { "title": "Username", "type": "string" },
                    "gravatar": { "title": "Gravatar", "type": "string", "default": "" },
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "domainRole": { "title": "Domainrole", "type": "string" },
                    "role": { "title": "Role", "type": "string", "default": "user" },
                    "isActive": {
                        "title": "Isactive",
                        "type": "boolean",
                        "default": false
                    },
                    "createdAt": {
                        "title": "Createdat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "updatedAt": {
                        "title": "Updatedat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "email": { "title": "Email", "type": "string", "format": "email" },
                    "studentId": {
                        "title": "Studentid",
                        "type": "string",
                        "default": ""
                    },
                    "realName": { "title": "Realname", "type": "string", "default": "" },
                    "registerIp": { "title": "Registerip", "type": "string" },
                    "loginAt": {
                        "title": "Loginat",
                        "type": "string",
                        "format": "date-time"
                    },
                    "loginIp": { "title": "Loginip", "type": "string" }
                }
            },
            "UserDetailWithDomainRoleList": {
                "title": "UserDetailWithDomainRoleList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": {
                            "$ref": "#/components/schemas/UserDetailWithDomainRole"
                        },
                        "default": []
                    }
                }
            },
            "UserDetailWithDomainRoleListResp": {
                "title": "UserDetailWithDomainRoleListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": {
                        "$ref": "#/components/schemas/UserDetailWithDomainRoleList"
                    }
                }
            },
            "UserDetailWithDomainRoleResp": {
                "title": "UserDetailWithDomainRoleResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/UserDetailWithDomainRole" }
                }
            },
            "UserEdit": {
                "title": "UserEdit",
                "type": "object",
                "properties": { "gravatar": { "title": "Gravatar", "type": "string" } }
            },
            "UserList": {
                "title": "UserList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/User" },
                        "default": []
                    }
                }
            },
            "UserListResp": {
                "title": "UserListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/UserList" }
                }
            },
            "UserPreview": {
                "title": "UserPreview",
                "required": ["id", "username"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "username": { "title": "Username", "type": "string" },
                    "gravatar": { "title": "Gravatar", "type": "string", "default": "" }
                }
            },
            "UserPreviewResp": {
                "title": "UserPreviewResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/UserPreview" }
                }
            },
            "UserResetPassword": {
                "title": "UserResetPassword",
                "required": ["newPassword"],
                "type": "object",
                "properties": {
                    "currentPassword": {
                        "title": "Currentpassword",
                        "type": "string",
                        "default": ""
                    },
                    "newPassword": { "title": "Newpassword", "type": "string" }
                }
            },
            "UserWithDomainRole": {
                "title": "UserWithDomainRole",
                "required": ["id", "username"],
                "type": "object",
                "properties": {
                    "id": { "title": "Id", "type": "string", "format": "uuid" },
                    "username": { "title": "Username", "type": "string" },
                    "gravatar": { "title": "Gravatar", "type": "string", "default": "" },
                    "domainId": {
                        "title": "Domainid",
                        "type": "string",
                        "format": "uuid"
                    },
                    "domainRole": { "title": "Domainrole", "type": "string" }
                }
            },
            "UserWithDomainRoleList": {
                "title": "UserWithDomainRoleList",
                "type": "object",
                "properties": {
                    "count": { "title": "Count", "type": "integer", "default": 0 },
                    "results": {
                        "title": "Results",
                        "type": "array",
                        "items": { "$ref": "#/components/schemas/UserWithDomainRole" },
                        "default": []
                    }
                }
            },
            "UserWithDomainRoleListResp": {
                "title": "UserWithDomainRoleListResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/UserWithDomainRoleList" }
                }
            },
            "UserWithDomainRoleResp": {
                "title": "UserWithDomainRoleResp",
                "required": ["errorCode"],
                "type": "object",
                "properties": {
                    "errorCode": { "$ref": "#/components/schemas/ErrorCode" },
                    "errorMsg": { "title": "Errormsg", "type": "string" },
                    "data": { "$ref": "#/components/schemas/UserWithDomainRole" }
                }
            },
            "ValidationError": {
                "title": "ValidationError",
                "required": ["loc", "msg", "type"],
                "type": "object",
                "properties": {
                    "loc": {
                        "title": "Location",
                        "type": "array",
                        "items": { "anyOf": [{ "type": "string" }, { "type": "integer" }] }
                    },
                    "msg": { "title": "Message", "type": "string" },
                    "type": { "title": "Error Type", "type": "string" }
                }
            },
            "Version": {
                "title": "Version",
                "required": ["version", "git"],
                "type": "object",
                "properties": {
                    "version": { "title": "Version", "type": "string" },
                    "git": { "title": "Git", "type": "string" }
                }
            }
        },
        "securitySchemes": {
            "HTTPBearer": {
                "type": "http",
                "scheme": "bearer",
                "bearerFormat": "JWT"
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "API",
	Description:      "This is an auto-generated API Docs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
