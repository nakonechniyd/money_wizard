// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/calc_years/{years}": {
            "get": {
                "description": "Calculate human wealth by specified parameters",
                "produces": [
                    "application/json"
                ],
                "summary": "Calculate human wealth",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Ammount of calculation years",
                        "name": "years",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Total percent of all taxes",
                        "name": "tax_percent",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Passive percent",
                        "name": "passive_percent",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Ammount of start capital",
                        "name": "capital",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Monthly replenishment",
                        "name": "monthly_replenishment",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Retire year",
                        "name": "retire_year",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "Life percent from net passive income",
                        "name": "life_percent",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/apiv1.YearCalc"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apiv1.MonthCalc": {
            "type": "object",
            "properties": {
                "capital_after": {
                    "type": "number"
                },
                "capital_before": {
                    "type": "number"
                },
                "month": {
                    "type": "integer"
                },
                "monthly_replenishment": {
                    "type": "number"
                },
                "net_life_passive_income": {
                    "type": "number"
                },
                "net_passive_income": {
                    "type": "number"
                },
                "passive_income": {
                    "type": "number"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "apiv1.YearCalc": {
            "type": "object",
            "properties": {
                "capital_after": {
                    "type": "number"
                },
                "capital_before": {
                    "type": "number"
                },
                "months": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apiv1.MonthCalc"
                    }
                },
                "total_monthly_replenishment": {
                    "type": "number"
                },
                "total_net_life_passive_income": {
                    "type": "number"
                },
                "total_net_passive_income": {
                    "type": "number"
                },
                "total_passive_income": {
                    "type": "number"
                },
                "year": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Money Wizard API",
	Description: "Simple calculation of human wealth accumulation",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
