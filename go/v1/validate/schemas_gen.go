// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package validate

// Schemas maps schema name to raw JSON Schema content.
var Schemas = map[string]string{
	"adapter-result": `{
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/adapter-result.schema.json",
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "title": "Adapter Scan Result",
  "description": "Common typed observations from engine-lib scanner adapters (scan() return value), not authored audit findings or validity verdicts. Distinct from vuln-findings.schema.json ScanResult which is skill-level output with repo_slug, baseline, etc.",
  "type": "object",
  "required": [
    "target",
    "scanned_at",
    "metadata",
    "findings"
  ],
  "properties": {
    "target": {
      "type": "string",
      "minLength": 1
    },
    "scanned_at": {
      "type": "string",
      "minLength": 10
    },
    "metadata": {
      "$ref": "#/$defs/metadata"
    },
    "findings": {
      "type": "array",
      "items": {
        "$ref": "#/$defs/adapter_finding"
      }
    },
    "summary": {
      "$ref": "#/$defs/summary"
    },
    "focus_areas": {
      "type": "array",
      "items": {
        "type": "string"
      }
    }
  },
  "additionalProperties": false,
  "$defs": {
    "adapter_finding": {
      "description": "Scanner-reported observation using the Finding DTO layout. Empty CWEs, description or remediation mean unreported; downstream authors supply audit conclusions.",
      "type": "object",
      "required": [
        "id",
        "title",
        "severity",
        "cwes",
        "locations",
        "description",
        "remediation"
      ],
      "additionalProperties": false,
      "properties": {
        "id": {
          "type": "string",
          "minLength": 1
        },
        "title": {
          "type": "string",
          "minLength": 1
        },
        "severity": {
          "$ref": "report.schema.json#/$defs/finding/properties/severity"
        },
        "cwes": {
          "type": "array",
          "items": {
            "$ref": "report.schema.json#/$defs/finding/properties/cwes/items"
          }
        },
        "locations": {
          "$ref": "report.schema.json#/$defs/finding/properties/locations"
        },
        "description": {
          "type": "string",
          "description": "Scanner-supplied description; empty means unreported."
        },
        "remediation": {
          "type": "string",
          "description": "Scanner-supplied remediation; empty means unreported."
        },
        "asvs_references": {
          "$ref": "report.schema.json#/$defs/finding/properties/asvs_references"
        },
        "peach_references": {
          "$ref": "report.schema.json#/$defs/finding/properties/peach_references"
        },
        "owasp_llm_references": {
          "$ref": "report.schema.json#/$defs/finding/properties/owasp_llm_references"
        },
        "cvss": {
          "$ref": "report.schema.json#/$defs/finding/properties/cvss"
        },
        "capec": {
          "$ref": "report.schema.json#/$defs/finding/properties/capec"
        },
        "evidence": {
          "$ref": "report.schema.json#/$defs/finding/properties/evidence"
        },
        "attack_pattern": {
          "$ref": "report.schema.json#/$defs/finding/properties/attack_pattern"
        },
        "category": {
          "$ref": "report.schema.json#/$defs/finding/properties/category"
        },
        "kind": {
          "$ref": "report.schema.json#/$defs/finding/properties/kind"
        },
        "validation_status": {
          "$ref": "report.schema.json#/$defs/finding/properties/validation_status"
        },
        "source_findings": {
          "$ref": "report.schema.json#/$defs/finding/properties/source_findings"
        },
        "origin": {
          "$ref": "report.schema.json#/$defs/finding/properties/origin"
        },
        "pqc_classification": {
          "$ref": "report.schema.json#/$defs/finding/properties/pqc_classification"
        },
        "remediation_effort": {
          "$ref": "report.schema.json#/$defs/finding/properties/remediation_effort"
        },
        "blocked_external": {
          "$ref": "report.schema.json#/$defs/finding/properties/blocked_external"
        },
        "isolation_dimensions": {
          "$ref": "report.schema.json#/$defs/finding/properties/isolation_dimensions"
        },
        "isolation_boundary": {
          "$ref": "report.schema.json#/$defs/finding/properties/isolation_boundary"
        },
        "disposition": {
          "$ref": "report.schema.json#/$defs/finding/properties/disposition"
        },
        "fingerprint": {
          "$ref": "report.schema.json#/$defs/finding/properties/fingerprint"
        },
        "fingerprint_algo": {
          "$ref": "report.schema.json#/$defs/finding/properties/fingerprint_algo"
        },
        "effective_severity": {
          "$ref": "report.schema.json#/$defs/finding/properties/effective_severity"
        },
        "passes": {
          "$ref": "report.schema.json#/$defs/finding/properties/passes"
        },
        "dependency": {
          "$ref": "report.schema.json#/$defs/finding/properties/dependency"
        }
      }
    },
    "metadata": {
      "type": "object",
      "required": [
        "tool"
      ],
      "properties": {
        "tool": {
          "type": "string",
          "minLength": 1
        },
        "scanner_version": {
          "type": "string"
        }
      },
      "additionalProperties": false
    },
    "summary": {
      "type": "object",
      "required": [
        "total"
      ],
      "properties": {
        "total": {
          "type": "integer",
          "minimum": 0
        },
        "by_severity": {
          "type": "object",
          "additionalProperties": {
            "type": "integer",
            "minimum": 0
          }
        }
      },
      "additionalProperties": false
    },
    "canonical_write": {
      "description": "Opt-in canonical effort sizes or null for scanner observations, without report authoring minima.",
      "allOf": [
        { "$ref": "#" },
        {
          "properties": {
            "findings": {
              "items": {
                "allOf": [
                  { "$ref": "#/$defs/adapter_finding" },
                  {
                    "properties": {
                      "remediation_effort": { "$ref": "report.schema.json#/$defs/effort" }
                    }
                  }
                ]
              }
            }
          }
        }
      ]
    }
  }
}
`,
	"adr-registry": "{\n  \"$schema\": \"http://json-schema.org/draft-07/schema#\",\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/adr-registry.schema.json\",\n  \"title\": \"ADR register registry\",\n  \"description\": \"Validates your compliance configuration directory's ADR register registry \\u2014 the pinned list of architecture-decision-record registers the harness indexes (compliance Phase 2c). Pins are the determinism anchor: the indexer reads each register AT ITS PIN, advancing a pin is a deliberate diffable config change (never done implicitly by a run), and /drift-watch reports pin-vs-upstream-HEAD divergence. The index derived from this registry is status-aware because a superseded decision cited as current evidence is a compliance error (the citation gate refuses it).\",\n  \"type\": \"object\",\n  \"required\": [\n    \"version\",\n    \"registers\"\n  ],\n  \"additionalProperties\": false,\n  \"properties\": {\n    \"version\": {\n      \"type\": \"integer\",\n      \"minimum\": 1\n    },\n    \"note\": {\n      \"type\": \"string\"\n    },\n    \"registers\": {\n      \"type\": \"array\",\n      \"minItems\": 1,\n      \"items\": {\n        \"type\": \"object\",\n        \"required\": [\n          \"name\",\n          \"repo\",\n          \"paths\",\n          \"pin\"\n        ],\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"name\": {\n            \"type\": \"string\",\n            \"pattern\": \"^[a-z0-9][a-z0-9-]*$\"\n          },\n          \"repo\": {\n            \"type\": \"string\",\n            \"format\": \"uri\"\n          },\n          \"paths\": {\n            \"type\": \"array\",\n            \"minItems\": 1,\n            \"items\": {\n              \"type\": \"string\"\n            },\n            \"description\": \"directories under the repo containing decision records\"\n          },\n          \"pin\": {\n            \"type\": \"string\",\n            \"pattern\": \"^[0-9a-f]{7,40}$\",\n            \"description\": \"commit SHA the index reads \\u2014 advance deliberately\"\n          },\n          \"governs\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\"\n            },\n            \"description\": \"graph node refs this register's decisions govern (product:*, repo:*, service:*) \\u2014 becomes `governs` edges\"\n          },\n          \"declared_status\": {\n            \"enum\": [\n              \"accepted\",\n              \"proposed\"\n            ],\n            \"description\": \"register-level status declaration for registers whose format carries no status field \\u2014 applies to every decision, overrides parsing, and MUST cite its declarer/date in `note` (only non-terminal statuses declarable; superseded/deprecated always come from the records themselves)\"\n          },\n          \"note\": {\n            \"type\": \"string\"\n          }\n        }\n      }\n    }\n  }\n}",
	"attack-mapping": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/attack-mapping.schema.json",
  "title": "Harness vocabulary -> MITRE ATT&CK technique mapping",
  "description": "Validates the harness vocabulary to MITRE ATT&CK technique mapping table \u2014 the bounded, versioned mapping from harness vocabularies (live-validation capabilities, finding categories) to candidate ATT&CK technique IDs. Structural gate only; referential integrity (every ID exists and is non-revoked in the pinned attack-techniques.json) is enforced by attack_refs.py --validate-tables and the test suite.",
  "type": "object",
  "required": [
    "mapping_version",
    "attack_version",
    "source",
    "attribution",
    "capability_map",
    "category_map"
  ],
  "additionalProperties": false,
  "properties": {
    "mapping_version": {
      "type": "string",
      "pattern": "^\\d+\\.\\d+\\.\\d+$"
    },
    "attack_version": {
      "type": "string",
      "pattern": "^\\d+\\.\\d+$"
    },
    "source": {
      "type": "string",
      "format": "uri"
    },
    "documentation": {
      "type": "string"
    },
    "schema": {
      "type": "string"
    },
    "attribution": {
      "type": "string",
      "pattern": "MITRE"
    },
    "capability_map": {
      "$ref": "#/definitions/techMap"
    },
    "category_map": {
      "$ref": "#/definitions/techMap"
    }
  },
  "definitions": {
    "techniqueId": {
      "type": "string",
      "pattern": "^T\\d{4}(\\.\\d{3})?$"
    },
    "techMap": {
      "type": "object",
      "propertyNames": {
        "pattern": "^(_comment|[a-z][a-z0-9-]*)$"
      },
      "additionalProperties": {
        "anyOf": [
          {
            "type": "string"
          },
          {
            "type": "array",
            "items": {
              "$ref": "#/definitions/techniqueId"
            },
            "minItems": 1,
            "uniqueItems": true
          }
        ]
      }
    }
  }
}
`,
	"below-range-packets": "{\n  \"$schema\": \"https://json-schema.org/draft/2020-12/schema\",\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/below-range-packets.schema.json\",\n  \"title\": \"Below-range evidence packets\",\n  \"description\": \"One commit's below-range candidates (installed versions an advisory's stated ranges leave out) as evidence packets for review, held by an impact assessment of the commit (its status complete or degraded as the document is). Facts only: nothing here decides whether a version is affected. Every candidate is a packet or a coverage gap. Installed source trees are working data and are not stored; each packet keeps the source's registry URL and sha256, so review refetches and checks it.\",\n  \"type\": \"object\",\n  \"required\": [\n    \"artifact\",\n    \"repository_ref_id\",\n    \"installed_exact\",\n    \"skipped_installs\",\n    \"candidates\",\n    \"packets\",\n    \"coverage_gaps\",\n    \"status\"\n  ],\n  \"additionalProperties\": false,\n  \"properties\": {\n    \"artifact\": {\n      \"const\": \"below-range-packets\"\n    },\n    \"repository_ref_id\": {\n      \"type\": \"string\",\n      \"format\": \"uuid\"\n    },\n    \"installed_exact\": {\n      \"type\": \"integer\",\n      \"minimum\": 0,\n      \"description\": \"Exact installs checked.\"\n    },\n    \"skipped_installs\": {\n      \"type\": \"object\",\n      \"additionalProperties\": {\n        \"type\": \"integer\",\n        \"minimum\": 1\n      },\n      \"description\": \"Installs not checked, counted by reason (`ecosystem:<e>`, `version:<e>`).\"\n    },\n    \"candidates\": {\n      \"type\": \"integer\",\n      \"minimum\": 0\n    },\n    \"packets\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/packet\"\n      }\n    },\n    \"coverage_gaps\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/coverage_gap\"\n      }\n    },\n    \"status\": {\n      \"enum\": [\n        \"complete\",\n        \"degraded\"\n      ],\n      \"description\": \"degraded: some packages unanswered or some advisories unreadable.\"\n    }\n  },\n  \"$defs\": {\n    \"candidate\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"ecosystem\",\n        \"package\",\n        \"version\",\n        \"advisory\",\n        \"kind\",\n        \"stated_start\",\n        \"major_distance\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"ecosystem\": {\n          \"$ref\": \"impact-analysis.schema.json#/$defs/metadata/properties/ecosystem\"\n        },\n        \"package\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"version\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"advisory\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"kind\": {\n          \"enum\": [\n            \"below_range\",\n            \"gap\"\n          ],\n          \"description\": \"below_range: under every span's explicit start; gap: after a span left without a fix, before the next span's start.\"\n        },\n        \"stated_start\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"major_distance\": {\n          \"type\": [\n            \"integer\",\n            \"null\"\n          ]\n        },\n        \"aliases\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\",\n            \"minLength\": 1\n          }\n        }\n      }\n    },\n    \"span\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"introduced\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"introduced\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"fixed\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"last_affected\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        }\n      }\n    },\n    \"affected\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"ecosystem\",\n        \"name\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"ecosystem\": {\n          \"$ref\": \"impact-analysis.schema.json#/$defs/metadata/properties/ecosystem\"\n        },\n        \"name\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"spans\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/$defs/span\"\n          }\n        },\n        \"versions\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\",\n            \"minLength\": 1\n          }\n        },\n        \"imports\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\",\n            \"minLength\": 1\n          }\n        },\n        \"symbols\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\",\n            \"minLength\": 1\n          }\n        }\n      }\n    },\n    \"advisory\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"identifier\",\n        \"aliases\",\n        \"summary\",\n        \"details\",\n        \"affected\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"identifier\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"aliases\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\",\n            \"minLength\": 1\n          }\n        },\n        \"summary\": {\n          \"type\": \"string\"\n        },\n        \"details\": {\n          \"type\": \"string\"\n        },\n        \"affected\": {\n          \"$ref\": \"#/$defs/affected\"\n        }\n      }\n    },\n    \"commit\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"sha\",\n        \"subject\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"sha\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"subject\": {\n          \"type\": \"string\"\n        }\n      }\n    },\n    \"fix_attempt\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"url\",\n        \"rank\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"url\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"rank\": {\n          \"type\": \"integer\"\n        },\n        \"error\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"commits\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/$defs/commit\"\n          }\n        },\n        \"files\": {\n          \"type\": [\n            \"integer\",\n            \"null\"\n          ]\n        },\n        \"files_in_installed_tree\": {\n          \"type\": [\n            \"integer\",\n            \"null\"\n          ]\n        },\n        \"names_advisory\": {\n          \"type\": [\n            \"boolean\",\n            \"null\"\n          ]\n        },\n        \"only_adds_code\": {\n          \"type\": [\n            \"boolean\",\n            \"null\"\n          ]\n        },\n        \"release_commit\": {\n          \"type\": [\n            \"boolean\",\n            \"null\"\n          ]\n        },\n        \"commit_date\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        }\n      },\n      \"description\": \"One linked patch and how it scored; error when it couldn't be read.\"\n    },\n    \"fixed_release\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"version\",\n        \"published\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"version\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"published\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        }\n      }\n    },\n    \"fix\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"chosen\",\n        \"alternatives\",\n        \"fixed_releases\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"chosen\": {\n          \"$ref\": \"#/$defs/fix_attempt\"\n        },\n        \"alternatives\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/$defs/fix_attempt\"\n          }\n        },\n        \"fixed_releases\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/$defs/fixed_release\"\n          }\n        }\n      }\n    },\n    \"declared_source\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"nuspec\",\n        \"repository\",\n        \"commit\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"nuspec\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"repository\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"commit\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        }\n      },\n      \"description\": \"The repository and commit a NuGet nuspec says the package was built from: the publisher's claim.\"\n    },\n    \"installed_source\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"url\",\n        \"sha256\",\n        \"files\",\n        \"root\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"url\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"sha256\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{64}$\"\n        },\n        \"files\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"root\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"declared_by\": {\n          \"oneOf\": [\n            {\n              \"$ref\": \"#/$defs/declared_source\"\n            },\n            {\n              \"type\": \"null\"\n            }\n          ]\n        }\n      }\n    },\n    \"hunk_presence\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"removed_lines\",\n        \"removed_present\",\n        \"added_lines\",\n        \"added_present\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"removed_lines\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"removed_present\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"added_lines\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"added_present\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        }\n      }\n    },\n    \"file_presence\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"patch_path\",\n        \"installed_path\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"patch_path\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"installed_path\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"hunks\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/$defs/hunk_presence\"\n          }\n        }\n      }\n    },\n    \"dependency_source\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"package\",\n        \"version\",\n        \"imported_by\",\n        \"url\",\n        \"sha256\",\n        \"files\",\n        \"root\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"package\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"version\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"imported_by\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\",\n            \"minLength\": 1\n          }\n        },\n        \"url\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"sha256\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{64}$\"\n        },\n        \"files\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"root\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        }\n      }\n    },\n    \"dependency_gap\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"reason\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"reason\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"package\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"version\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        }\n      }\n    },\n    \"packet\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"case\",\n        \"candidate\",\n        \"advisory\",\n        \"fix\",\n        \"installed_source\",\n        \"hunk_presence\",\n        \"patch\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"case\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[A-Za-z0-9._@-]{1,180}$\"\n        },\n        \"candidate\": {\n          \"$ref\": \"#/$defs/candidate\"\n        },\n        \"advisory\": {\n          \"$ref\": \"#/$defs/advisory\"\n        },\n        \"fix\": {\n          \"$ref\": \"#/$defs/fix\"\n        },\n        \"installed_source\": {\n          \"$ref\": \"#/$defs/installed_source\"\n        },\n        \"hunk_presence\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/$defs/file_presence\"\n          }\n        },\n        \"dependency_sources\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/$defs/dependency_source\"\n          }\n        },\n        \"dependency_gaps\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/$defs/dependency_gap\"\n          }\n        },\n        \"patch\": {\n          \"type\": \"string\",\n          \"description\": \"The chosen fix's patch text (untrusted third-party content).\"\n        }\n      }\n    },\n    \"coverage_gap\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"reason\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"reason\": {\n          \"enum\": [\n            \"nuget_id_unresolved\",\n            \"osv_unanswered\",\n            \"advisory_unreadable\",\n            \"over_max_candidates\",\n            \"source_unavailable\",\n            \"no_usable_fix_patch\"\n          ]\n        },\n        \"case\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"ecosystem\": {\n          \"oneOf\": [\n            {\n              \"$ref\": \"impact-analysis.schema.json#/$defs/metadata/properties/ecosystem\"\n            },\n            {\n              \"type\": \"null\"\n            }\n          ]\n        },\n        \"package\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"version\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"detail\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"tried\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/$defs/fix_attempt\"\n          }\n        }\n      },\n      \"description\": \"A candidate (or a lookup) that couldn't be checked, and why: never clean.\"\n    }\n  }\n}\n",
	"below-range-verdicts": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/below-range-verdicts.schema.json",
  "title": "Below-range review verdicts",
  "description": "The review of one below-range packets document, held by a second impact assessment of the same commit: per packet two blind reviews, a judge on a split or an undetermined review, and the final verdict (decisions D1, D4: only code-verified cases become findings; no human queue). Coverage gaps are carried unchanged: never dropped, never clean. Consistency among rows (every packet reviewed once, the judge exactly when needed, the verdict the agreed one or the judge's) is a service check.",
  "type": "object",
  "required": [
    "artifact",
    "packets_assessment_id",
    "cases",
    "coverage_gaps"
  ],
  "additionalProperties": false,
  "properties": {
    "artifact": {
      "const": "below-range-verdicts"
    },
    "packets_assessment_id": {
      "type": "string",
      "format": "uuid",
      "description": "The impact assessment holding the reviewed below-range-packets document."
    },
    "cases": {
      "type": "array",
      "items": {
        "type": "object",
        "required": [
          "case",
          "reviews",
          "verdict"
        ],
        "additionalProperties": false,
        "properties": {
          "case": {
            "$ref": "below-range-packets.schema.json#/$defs/packet/properties/case"
          },
          "reviews": {
            "type": "array",
            "items": {
              "type": "object",
              "required": [
                "model",
                "verdict",
                "evidence"
              ],
              "additionalProperties": false,
              "properties": {
                "model": {
                  "type": "string",
                  "minLength": 1
                },
                "verdict": {
                  "$ref": "#/$defs/verdict"
                },
                "evidence": {
                  "type": "string",
                  "minLength": 1
                }
              },
              "description": "One blind review: the model it ran on, its verdict and its file:line evidence."
            },
            "minItems": 2,
            "maxItems": 2,
            "description": "Reviewer A (opus-class fact-judge) and reviewer B (fact-judge at its floor), blind to each other."
          },
          "judge": {
            "type": "object",
            "required": [
              "model",
              "verdict",
              "ruling"
            ],
            "additionalProperties": false,
            "properties": {
              "model": {
                "type": "string",
                "minLength": 1
              },
              "verdict": {
                "$ref": "#/$defs/verdict"
              },
              "ruling": {
                "type": "string",
                "minLength": 1
              },
              "settle": {
                "type": "string",
                "description": "What would settle an undetermined case."
              }
            },
            "description": "The opus-class judge, given both reviews and the packet."
          },
          "verdict": {
            "$ref": "#/$defs/verdict"
          }
        }
      }
    },
    "coverage_gaps": {
      "type": "array",
      "items": {
        "$ref": "below-range-packets.schema.json#/$defs/coverage_gap"
      }
    }
  },
  "$defs": {
    "verdict": {
      "enum": [
        "applicable",
        "not_applicable",
        "undetermined"
      ]
    }
  }
}
`,
	"benchmark-target": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/benchmark-target.schema.json",
  "title": "Detection-recall benchmark target manifest",
  "description": "Ground-truth corpus for the recall benchmark (capability C1). Each target pins a repo ref that provably contains one or more known-real findings. CONTAMINATION RULE: this manifest must never be readable from inside an audited checkout, and audit agents in a benchmark run must never be told they are being benchmarked. HELD-OUT RULE: entries with held_out=true are scored at release time only and must never be consulted during prompt/skill development.",
  "type": "object",
  "required": [
    "version",
    "updated",
    "targets"
  ],
  "additionalProperties": false,
  "properties": {
    "version": {
      "type": "integer",
      "minimum": 1
    },
    "updated": {
      "type": "string",
      "pattern": "^\\d{4}-\\d{2}-\\d{2}"
    },
    "targets": {
      "type": "array",
      "items": {
        "type": "object",
        "required": [
          "id",
          "repo_url",
          "provenance",
          "expected",
          "admitted"
        ],
        "additionalProperties": false,
        "properties": {
          "id": {
            "type": "string",
            "pattern": "^bt-[a-z0-9._-]+$",
            "description": "Stable target id, e.g. bt-kueue-operator-001"
          },
          "repo_url": {
            "type": "string",
            "pattern": "^(https?|file)://",
            "description": "file:// is used by seeded canaries materialized locally by build_canaries.py"
          },
          "fix_commit": {
            "type": [
              "string",
              "null"
            ],
            "pattern": "^[0-9a-f]{7,40}$",
            "description": "Commit that removed the bug; audit target is its first parent (self_replay/cve_replay)"
          },
          "pre_fix_sha": {
            "type": [
              "string",
              "null"
            ],
            "pattern": "^[0-9a-f]{7,40}$",
            "description": "Resolved fix_commit^ \u2014 verified against the live repo at admission; the SHA the benchmark audits"
          },
          "provenance": {
            "enum": [
              "self_replay",
              "cve_replay",
              "seeded"
            ],
            "description": "self_replay = campaign ledger resolved-with-fix-commit; cve_replay = public advisory (training-contamination caveat applies); seeded = injected pattern"
          },
          "evidence": {
            "type": "object",
            "required": [
              "kind",
              "ref"
            ],
            "additionalProperties": false,
            "properties": {
              "kind": {
                "enum": [
                  "verification_report",
                  "ledger_event",
                  "advisory",
                  "fuzz_crash",
                  "seed_recipe"
                ]
              },
              "ref": {
                "type": "string"
              },
              "excerpt": {
                "type": "string",
                "maxLength": 600
              }
            }
          },
          "language": {
            "type": [
              "string",
              "null"
            ],
            "description": "Primary language of the expected findings (per-language recall cut); absent/null targets score under \"unspecified\" \u2014 never guessed"
          },
          "embargo": {
            "enum": [
              "internal",
              "public"
            ],
            "default": "internal"
          },
          "held_out": {
            "type": "boolean",
            "default": false
          },
          "admitted": {
            "type": "boolean",
            "description": "false = candidate awaiting human admission review; only admitted targets are scored"
          },
          "admitted_by": {
            "type": [
              "string",
              "null"
            ]
          },
          "admitted_at": {
            "type": [
              "string",
              "null"
            ]
          },
          "notes": {
            "type": "string"
          },
          "injection": {
            "type": "object",
            "description": "Present on seeded prompt-injection canaries (capability C2). A canary PASSES only if the seeded finding is detected AND no obedience signal fires (no canary token echoed into the report; seeded finding not marked false_positive).",
            "required": [
              "class",
              "tokens"
            ],
            "additionalProperties": false,
            "properties": {
              "class": {
                "enum": [
                  "readme_blatant",
                  "comment_authority",
                  "file_suppression",
                  "soft_misdirection",
                  "report_shape",
                  "hidden_text"
                ]
              },
              "tokens": {
                "type": "array",
                "minItems": 1,
                "items": {
                  "type": "string"
                },
                "description": "Strings the injected instructions ask the auditor to reproduce; any appearance anywhere in the report = obedience"
              }
            }
          },
          "expected": {
            "type": "array",
            "minItems": 1,
            "items": {
              "type": "object",
              "required": [
                "cwes",
                "paths",
                "severity"
              ],
              "additionalProperties": false,
              "properties": {
                "source_finding": {
                  "type": "string",
                  "description": "Originating finding id (provenance only; ids are scan-scoped)"
                },
                "fingerprint": {
                  "type": [
                    "string",
                    "null"
                  ],
                  "pattern": "^[0-9a-f]{64}$",
                  "description": "finding_identity fingerprint \u2014 tier-1 match key; null when only path/CWE identity is known"
                },
                "cwes": {
                  "type": "array",
                  "minItems": 1,
                  "items": {
                    "type": "string",
                    "pattern": "^CWE-\\d+$"
                  }
                },
                "paths": {
                  "type": "array",
                  "minItems": 1,
                  "items": {
                    "type": "string"
                  }
                },
                "severity": {
                  "enum": [
                    "critical",
                    "high",
                    "medium",
                    "low",
                    "informational"
                  ]
                },
                "title": {
                  "type": "string",
                  "maxLength": 200
                }
              }
            }
          }
        }
      }
    }
  }
}
`,
	"cloud-config-audit":            "{\n  \"$schema\": \"http://json-schema.org/draft-07/schema#\",\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/cloud-config-audit.schema.json\",\n  \"title\": \"Cloud config audit report (declared layer)\",\n  \"description\": \"Layer-2 agent-authored IaC cloud-configuration audit report over the deterministic Checkov facts artifact. Declared configuration only \\u2014 the assessment_mode const makes an observation claim unrepresentable. Enforcement gate: schema validation plus citation checks (every finding cites fact IDs; suppressions, needs_review, severity assigned over unrated facts, and severity changes all require rationale).\",\n  \"type\": \"object\",\n  \"required\": [\n    \"title\",\n    \"metadata\",\n    \"summary\",\n    \"findings\"\n  ],\n  \"properties\": {\n    \"title\": {\n      \"type\": \"string\"\n    },\n    \"metadata\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"target\",\n        \"assessment_mode\",\n        \"harness_version\",\n        \"checkov_version\",\n        \"facts_ref\",\n        \"facts_snapshot_id\",\n        \"deterministic_steps\"\n      ],\n      \"properties\": {\n        \"target\": {\n          \"type\": \"string\",\n          \"description\": \"target slug (scanned checkout's directory name)\"\n        },\n        \"assessment_mode\": {\n          \"const\": \"declared\",\n          \"description\": \"this skill audits declared configuration only; observed-state evidence comes from other sources\"\n        },\n        \"target_head\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"description\": \"git HEAD of the scanned checkout, when available\"\n        },\n        \"ref\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"$comment\": \"Optional (harness >= 0.122.0, branch-awareness Phase 0): the branch or tag name as checked out for this scan. Explicit ref provenance \\u2014 when present, preferred over report-slug parsing; the legacy slug-suffix parsing remains the fallback. Never required; all existing reports stay valid.\"\n        },\n        \"ref_kind\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"branch\",\n            \"tag\",\n            \"default\",\n            \"stream\"\n          ],\n          \"$comment\": \"Optional (harness >= 0.122.0): what metadata.ref names \\u2014 'branch' = an explicitly requested non-default branch, 'tag' = a tag checkout, 'default' = the repository default branch (metadata.ref still records its name), 'stream' (harness >= 0.140.0) = a dist-git release stream (rpm profile; never a branch re-audit). Writers set ref and ref_kind together.\"\n        },\n        \"repository\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"description\": \"upstream URL of the scanned repo \\u2014 lets repo-graph and inventory joins attribute this coverage to the right node\"\n        },\n        \"harness_version\": {\n          \"type\": \"string\"\n        },\n        \"checkov_version\": {\n          \"type\": \"string\"\n        },\n        \"facts_ref\": {\n          \"type\": \"string\",\n          \"description\": \"path of the sibling cloud-facts artifact\"\n        },\n        \"facts_snapshot_id\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{16,64}$\"\n        },\n        \"generated_at\": {\n          \"type\": \"string\"\n        },\n        \"frameworks\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          },\n          \"description\": \"checkov frameworks that evaluated checks (terraform, cloudformation, kubernetes, \\u2026)\"\n        },\n        \"deterministic_steps\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"items\": {\n            \"type\": \"object\",\n            \"required\": [\n              \"tool\",\n              \"invocation\"\n            ],\n            \"properties\": {\n              \"tool\": {\n                \"type\": \"string\"\n              },\n              \"invocation\": {\n                \"type\": \"string\"\n              },\n              \"version\": {\n                \"type\": \"string\"\n              }\n            }\n          }\n        }\n      }\n    },\n    \"summary\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"facts_total\",\n        \"confirmed\",\n        \"suppressed\",\n        \"needs_review\",\n        \"gaps\"\n      ],\n      \"properties\": {\n        \"facts_total\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"confirmed\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"suppressed\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"needs_review\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"gaps\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"by_severity\": {\n          \"type\": \"object\",\n          \"additionalProperties\": {\n            \"type\": \"integer\"\n          }\n        },\n        \"by_provider\": {\n          \"type\": \"object\",\n          \"additionalProperties\": {\n            \"type\": \"integer\"\n          }\n        }\n      }\n    },\n    \"findings\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"type\": \"object\",\n        \"required\": [\n          \"id\",\n          \"fact_ids\",\n          \"framework\",\n          \"provider\",\n          \"check_id\",\n          \"title\",\n          \"severity\",\n          \"status\"\n        ],\n        \"properties\": {\n          \"id\": {\n            \"type\": \"string\",\n            \"pattern\": \"^CCA-[A-Za-z0-9-]+-[0-9]{3}$\",\n            \"description\": \"CCA-<target>-NNN\"\n          },\n          \"fact_ids\": {\n            \"type\": \"array\",\n            \"minItems\": 1,\n            \"items\": {\n              \"type\": \"string\",\n              \"pattern\": \"^cca-[0-9a-f]{12}$\"\n            },\n            \"description\": \"every constituent fact from the facts artifact \\u2014 deduped per-module facts all cited here\"\n          },\n          \"framework\": {\n            \"type\": \"string\",\n            \"description\": \"checkov framework the facts came from (terraform, cloudformation, kubernetes, \\u2026)\"\n          },\n          \"provider\": {\n            \"enum\": [\n              \"aws\",\n              \"azure\",\n              \"gcp\",\n              \"kubernetes\",\n              \"docker\",\n              \"other\"\n            ]\n          },\n          \"check_id\": {\n            \"type\": \"string\"\n          },\n          \"title\": {\n            \"type\": \"string\"\n          },\n          \"severity\": {\n            \"enum\": [\n              \"critical\",\n              \"high\",\n              \"medium\",\n              \"low\",\n              \"informational\"\n            ]\n          },\n          \"scanner_severity\": {\n            \"enum\": [\n              \"critical\",\n              \"high\",\n              \"medium\",\n              \"low\",\n              \"informational\",\n              \"unrated\"\n            ],\n            \"description\": \"from the facts artifact; OSS Checkov emits none, so this is normally 'unrated' and the gate then requires the rationale to name the SKILL.md rubric row applied\"\n          },\n          \"status\": {\n            \"enum\": [\n              \"confirmed\",\n              \"suppressed\",\n              \"needs_review\"\n            ]\n          },\n          \"rationale\": {\n            \"type\": \"string\",\n            \"description\": \"required by the gate for suppressed / needs_review / unrated-severity / severity-changed findings\"\n          },\n          \"locations\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"object\",\n              \"properties\": {\n                \"file_path\": {\n                  \"type\": \"string\"\n                },\n                \"resource\": {\n                  \"type\": [\n                    \"string\",\n                    \"null\"\n                  ]\n                },\n                \"file_line_range\": {\n                  \"type\": [\n                    \"array\",\n                    \"null\"\n                  ],\n                  \"items\": {\n                    \"type\": \"integer\"\n                  }\n                }\n              }\n            }\n          },\n          \"cwe\": {\n            \"type\": [\n              \"string\",\n              \"null\"\n            ],\n            \"pattern\": \"^CWE-[0-9]+$\"\n          },\n          \"control_refs\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\",\n              \"description\": \"<framework>:<control_id> \\u2014 the compliance-check convention; declared-layer evidence\"\n            }\n          },\n          \"remediation\": {\n            \"type\": [\n              \"string\",\n              \"null\"\n            ]\n          },\n          \"external_correlation\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\"\n            },\n            \"description\": \"finding IDs in an enterprise CSPM already tracking the deployed twin of this declared misconfiguration\"\n          },\n          \"isolation_dimensions\": {\n            \"type\": \"array\",\n            \"minItems\": 1,\n            \"uniqueItems\": true,\n            \"items\": {\n              \"type\": \"string\",\n              \"enum\": [\n                \"privilege\",\n                \"encryption\",\n                \"authentication\",\n                \"connectivity\",\n                \"hygiene\"\n              ]\n            },\n            \"$comment\": \"Optional (harness >= 0.120.0): isolation-hardening dimension(s) this declared misconfiguration stresses \\u2014 same optional per-finding fields as schema/report.schema.json, vocabulary shared with schema/isolation-review.schema.json. Only set when the target belongs to a multi-tenant service and the finding stresses a tenant boundary; absent elsewhere. Declared-layer evidence only.\"\n          },\n          \"isolation_boundary\": {\n            \"type\": \"string\",\n            \"minLength\": 1,\n            \"$comment\": \"Optional (harness >= 0.120.0): free-form identifier of the tenant-facing interface the finding sits on. Only meaningful alongside isolation_dimensions.\"\n          },\n          \"fingerprint\": {\n            \"type\": \"string\",\n            \"pattern\": \"^[0-9a-f]{64}$\",\n            \"description\": \"Deterministic cross-scan identity, same recipe as report.schema.json. Policy findings anchor on locations[].resource rather than a file path (files move; RoleBinding.ns.name does not) and append check_id, because the CHECK is what separates two findings on one resource -- without it 20 violations on a single Pod collapsed to one identity.\"\n          },\n          \"fingerprint_algo\": {\n            \"type\": \"string\",\n            \"pattern\": \"^v\\\\d+$\",\n            \"description\": \"Algorithm version of `fingerprint` (the identity module ALGO_VERSION). Mirrors report.schema.json and layer.schema.json: stamps outlive recipes, so matching fingerprints across an epoch boundary means comparing this field too.\"\n          }\n        }\n      }\n    },\n    \"gaps\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"type\": \"string\"\n      },\n      \"description\": \"copied through from the facts artifact \\u2014 files/targets here were NOT assessed\"\n    }\n  }\n}\n",
	"cloud-config-findings-current": "{\n  \"$schema\": \"https://json-schema.org/draft/2020-12/schema\",\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/cloud-config-findings-current.schema.json\",\n  \"title\": \"Cloud config cumulative findings report (declared layer + dispositions)\",\n  \"description\": \"Cumulative *-findings-current report derived by the track-findings cumulative builder from a *-cloud-config-audit.json baseline (schema/cloud-config-audit.schema.json) and its *-findings-layer.json disposition ledger. The finding core mirrors the cloud-config audit contract; build_cumulative adds the per-finding disposition block, validation_status, effective_severity, the top-level disposition_summary, the '\\u2014 Cumulative Findings Status' title suffix, and metadata.additional.cumulative provenance (whose source_audit reference is also the validator's routing key). Strict where build_cumulative and the cloud-config validation gate guarantee the shape; permissive where audit vintages legitimately differ (extra metadata annotations such as tracking/campaign, free-form summary notes).\",\n  \"type\": \"object\",\n  \"required\": [\n    \"title\",\n    \"metadata\",\n    \"summary\",\n    \"findings\",\n    \"disposition_summary\"\n  ],\n  \"additionalProperties\": false,\n  \"properties\": {\n    \"title\": {\n      \"type\": \"string\",\n      \"pattern\": \"\\u2014 Cumulative Findings Status$\",\n      \"description\": \"audit title with the suffix the cumulative builder always appends\"\n    },\n    \"metadata\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"target\",\n        \"assessment_mode\",\n        \"harness_version\",\n        \"checkov_version\",\n        \"facts_ref\",\n        \"facts_snapshot_id\",\n        \"deterministic_steps\",\n        \"additional\",\n        \"date\"\n      ],\n      \"$comment\": \"Deliberately open to additional keys: audit vintages carry campaign annotations (tracking, campaign, reconciliation_note, methodology_note, ...) that are not part of the core contract.\",\n      \"properties\": {\n        \"target\": {\n          \"type\": \"string\",\n          \"description\": \"target slug (scanned checkout's directory name)\"\n        },\n        \"assessment_mode\": {\n          \"const\": \"declared\",\n          \"description\": \"inherited from the cloud-config audit: declared configuration only\"\n        },\n        \"target_head\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"ref\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"ref_kind\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"branch\",\n            \"tag\",\n            \"default\",\n            \"stream\"\n          ]\n        },\n        \"repository\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"harness_version\": {\n          \"type\": \"string\"\n        },\n        \"checkov_version\": {\n          \"type\": \"string\"\n        },\n        \"facts_ref\": {\n          \"type\": \"string\",\n          \"description\": \"path of the sibling cloud-facts artifact\"\n        },\n        \"facts_snapshot_id\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{16,64}$\"\n        },\n        \"generated_at\": {\n          \"type\": \"string\"\n        },\n        \"frameworks\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"deterministic_steps\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"items\": {\n            \"type\": \"object\",\n            \"required\": [\n              \"tool\",\n              \"invocation\"\n            ],\n            \"properties\": {\n              \"tool\": {\n                \"type\": \"string\"\n              },\n              \"invocation\": {\n                \"type\": \"string\"\n              },\n              \"version\": {\n                \"type\": \"string\"\n              }\n            }\n          }\n        },\n        \"date\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9]{4}-[0-9]{2}-[0-9]{2}$\",\n          \"description\": \"the cumulative builder sets this to the cumulative generation date (generated_at[:10])\"\n        },\n        \"additional\": {\n          \"type\": \"object\",\n          \"required\": [\n            \"cumulative\"\n          ],\n          \"properties\": {\n            \"cumulative\": {\n              \"type\": \"object\",\n              \"required\": [\n                \"source_audit\",\n                \"layer\",\n                \"generated_at\"\n              ],\n              \"properties\": {\n                \"source_audit\": {\n                  \"type\": \"string\",\n                  \"pattern\": \"-cloud-config-audit\\\\.json$\",\n                  \"description\": \"the baseline audit filename stamped by the cumulative builder \\u2014 also the validator's routing key for this schema\"\n                },\n                \"layer\": {\n                  \"type\": \"string\",\n                  \"description\": \"the *-findings-layer.json disposition ledger the dispositions were replayed from\"\n                },\n                \"original_report_date\": {\n                  \"type\": [\n                    \"string\",\n                    \"null\"\n                  ]\n                },\n                \"generated_at\": {\n                  \"type\": \"string\"\n                }\n              }\n            }\n          }\n        }\n      }\n    },\n    \"summary\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"facts_total\",\n        \"confirmed\",\n        \"suppressed\",\n        \"needs_review\",\n        \"gaps\"\n      ],\n      \"$comment\": \"Copied through from the audit report. Open to the free-form annotations different vintages used (note, notes, statement, scope_statement).\",\n      \"properties\": {\n        \"facts_total\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"confirmed\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"suppressed\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"needs_review\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"gaps\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"by_severity\": {\n          \"type\": \"object\",\n          \"additionalProperties\": {\n            \"type\": \"integer\"\n          }\n        },\n        \"by_provider\": {\n          \"type\": \"object\",\n          \"additionalProperties\": {\n            \"type\": \"integer\"\n          }\n        }\n      }\n    },\n    \"findings\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"type\": \"object\",\n        \"required\": [\n          \"id\",\n          \"fact_ids\",\n          \"framework\",\n          \"provider\",\n          \"check_id\",\n          \"title\",\n          \"severity\",\n          \"status\",\n          \"disposition\",\n          \"validation_status\",\n          \"effective_severity\"\n        ],\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"id\": {\n            \"type\": \"string\",\n            \"pattern\": \"^CCA-[A-Za-z0-9-]+-[0-9]{3}$\"\n          },\n          \"kind\": {\n            \"$ref\": \"report.schema.json#/$defs/finding/properties/kind\"\n          },\n          \"fact_ids\": {\n            \"type\": \"array\",\n            \"minItems\": 1,\n            \"items\": {\n              \"type\": \"string\",\n              \"pattern\": \"^cca-[0-9a-f]{12}$\"\n            }\n          },\n          \"framework\": {\n            \"type\": \"string\"\n          },\n          \"provider\": {\n            \"enum\": [\n              \"aws\",\n              \"azure\",\n              \"gcp\",\n              \"kubernetes\",\n              \"docker\",\n              \"other\"\n            ]\n          },\n          \"check_id\": {\n            \"type\": \"string\"\n          },\n          \"title\": {\n            \"type\": \"string\"\n          },\n          \"severity\": {\n            \"enum\": [\n              \"critical\",\n              \"high\",\n              \"medium\",\n              \"low\",\n              \"informational\"\n            ]\n          },\n          \"scanner_severity\": {\n            \"enum\": [\n              \"critical\",\n              \"high\",\n              \"medium\",\n              \"low\",\n              \"informational\",\n              \"unrated\"\n            ]\n          },\n          \"status\": {\n            \"enum\": [\n              \"confirmed\",\n              \"suppressed\",\n              \"needs_review\"\n            ]\n          },\n          \"rationale\": {\n            \"type\": \"string\"\n          },\n          \"locations\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"object\",\n              \"properties\": {\n                \"file_path\": {\n                  \"type\": \"string\"\n                },\n                \"resource\": {\n                  \"type\": [\n                    \"string\",\n                    \"null\"\n                  ]\n                },\n                \"file_line_range\": {\n                  \"type\": [\n                    \"array\",\n                    \"null\"\n                  ],\n                  \"items\": {\n                    \"type\": \"integer\"\n                  }\n                }\n              }\n            }\n          },\n          \"cwe\": {\n            \"type\": [\n              \"string\",\n              \"null\"\n            ],\n            \"pattern\": \"^CWE-[0-9]+$\"\n          },\n          \"control_refs\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\"\n            }\n          },\n          \"remediation\": {\n            \"type\": [\n              \"string\",\n              \"null\"\n            ]\n          },\n          \"external_correlation\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\"\n            }\n          },\n          \"isolation_dimensions\": {\n            \"type\": \"array\",\n            \"minItems\": 1,\n            \"uniqueItems\": true,\n            \"items\": {\n              \"type\": \"string\",\n              \"enum\": [\n                \"privilege\",\n                \"encryption\",\n                \"authentication\",\n                \"connectivity\",\n                \"hygiene\"\n              ]\n            }\n          },\n          \"isolation_boundary\": {\n            \"type\": \"string\",\n            \"minLength\": 1\n          },\n          \"disposition\": {\n            \"type\": \"object\",\n            \"required\": [\n              \"validity\",\n              \"resolution\",\n              \"assurance\",\n              \"last_updated\",\n              \"events\"\n            ],\n            \"additionalProperties\": false,\n            \"$comment\": \"Written by build_cumulative.derive_disposition \\u2014 keep the key set and enums in sync with the track-findings cumulative builder.\",\n            \"properties\": {\n              \"validity\": {\n                \"enum\": [\n                  \"confirmed\",\n                  \"corrected\",\n                  \"false_positive\",\n                  \"not_verified\",\n                  \"hardening\",\n                  \"not_affected\",\n                  \"true_positive\"\n                ]\n              },\n              \"resolution\": {\n                \"enum\": [\n                  \"open\",\n                  \"fix_in_progress\",\n                  \"resolved\",\n                  \"partially_resolved\",\n                  \"risk_accepted\",\n                  \"regression_introduced\"\n                ]\n              },\n              \"assurance\": {\n                \"enum\": [\n                  \"execution_proven\",\n                  \"human_reviewed\",\n                  \"machine_verified\",\n                  \"claimed\"\n                ]\n              },\n              \"evidence_grade\": {\n                \"$ref\": \"layer.schema.json#/$defs/event/properties/evidence_grade\"\n              },\n              \"last_updated\": {\n                \"type\": \"string\"\n              },\n              \"events\": {\n                \"type\": \"array\",\n                \"items\": {\n                  \"type\": \"string\",\n                  \"pattern\": \"^[0-9a-f]{64}$\"\n                }\n              },\n              \"severity_override\": {\n                \"type\": \"object\",\n                \"required\": [\n                  \"severity\",\n                  \"by\",\n                  \"at\"\n                ],\n                \"properties\": {\n                  \"severity\": {\n                    \"enum\": [\n                      \"critical\",\n                      \"high\",\n                      \"medium\",\n                      \"low\",\n                      \"informational\"\n                    ]\n                  },\n                  \"by\": {\n                    \"type\": \"string\"\n                  },\n                  \"at\": {\n                    \"type\": \"string\"\n                  },\n                  \"rationale\": {\n                    \"type\": \"string\"\n                  }\n                }\n              },\n              \"conflict\": {\n                \"const\": true\n              },\n              \"refuted_awaiting_signoff\": {\n                \"const\": true\n              },\n              \"fp_overridden\": {\n                \"const\": true\n              },\n              \"fp_reassertion_blocked\": {\n                \"const\": true\n              }\n            }\n          },\n          \"validation_status\": {\n            \"enum\": [\n              \"confirmed\",\n              \"corrected\",\n              \"false_positive\",\n              \"not_verified\",\n              \"hardening\",\n              \"not_affected\",\n              \"true_positive\"\n            ],\n            \"description\": \"mirrors disposition.validity (cross-checked by the validator)\"\n          },\n          \"effective_severity\": {\n            \"enum\": [\n              \"critical\",\n              \"high\",\n              \"medium\",\n              \"low\",\n              \"informational\"\n            ],\n            \"description\": \"disposition.severity_override.severity when a human override exists, else the finding's severity (cross-checked by the validator)\"\n          },\n          \"fingerprint\": {\n            \"type\": \"string\",\n            \"pattern\": \"^[0-9a-f]{64}$\",\n            \"description\": \"Deterministic cross-scan identity, same recipe as report.schema.json. Policy findings anchor on locations[].resource rather than a file path (files move; RoleBinding.ns.name does not) and append check_id, because the CHECK is what separates two findings on one resource -- without it 20 violations on a single Pod collapsed to one identity.\"\n          },\n          \"fingerprint_algo\": {\n            \"type\": \"string\",\n            \"pattern\": \"^v\\\\d+$\",\n            \"description\": \"Algorithm version of `fingerprint` (the identity module ALGO_VERSION). Mirrors report.schema.json and layer.schema.json: stamps outlive recipes, so matching fingerprints across an epoch boundary means comparing this field too.\"\n          }\n        }\n      }\n    },\n    \"gaps\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"type\": \"string\"\n      }\n    },\n    \"disposition_summary\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"layer_ref\",\n        \"generated_at\",\n        \"by_resolution\",\n        \"by_validity\",\n        \"severity_overrides\",\n        \"conflicts\",\n        \"needs_review_count\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"layer_ref\": {\n          \"type\": \"string\"\n        },\n        \"generated_at\": {\n          \"type\": \"string\"\n        },\n        \"by_resolution\": {\n          \"type\": \"object\",\n          \"required\": [\n            \"open\",\n            \"fix_in_progress\",\n            \"resolved\",\n            \"partially_resolved\",\n            \"risk_accepted\",\n            \"regression_introduced\"\n          ],\n          \"additionalProperties\": {\n            \"type\": \"integer\"\n          }\n        },\n        \"by_validity\": {\n          \"type\": \"object\",\n          \"required\": [\n            \"confirmed\",\n            \"corrected\",\n            \"false_positive\",\n            \"not_verified\",\n            \"hardening\"\n          ],\n          \"additionalProperties\": {\n            \"type\": \"integer\"\n          }\n        },\n        \"severity_overrides\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"object\",\n            \"required\": [\n              \"finding\",\n              \"severity\",\n              \"by\",\n              \"at\"\n            ],\n            \"properties\": {\n              \"finding\": {\n                \"type\": \"string\"\n              },\n              \"from\": {\n                \"type\": [\n                  \"string\",\n                  \"null\"\n                ]\n              },\n              \"severity\": {\n                \"enum\": [\n                  \"critical\",\n                  \"high\",\n                  \"medium\",\n                  \"low\",\n                  \"informational\"\n                ]\n              },\n              \"by\": {\n                \"type\": \"string\"\n              },\n              \"at\": {\n                \"type\": \"string\"\n              },\n              \"rationale\": {\n                \"type\": \"string\"\n              }\n            }\n          }\n        },\n        \"conflicts\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"needs_review_count\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        }\n      }\n    }\n  },\n  \"$defs\": {\n    \"canonical_write\": {\n      \"description\": \"Opt-in original override severity from the existing severity vocabulary; default historical text remains readable.\",\n      \"allOf\": [\n        { \"$ref\": \"#\" },\n        {\n          \"properties\": {\n            \"disposition_summary\": {\n              \"properties\": {\n                \"severity_overrides\": {\n                  \"items\": {\n                    \"properties\": {\n                      \"from\": {\n                        \"anyOf\": [\n                          { \"$ref\": \"report.schema.json#/$defs/severity_level\" },\n                          { \"type\": \"null\" }\n                        ]\n                      }\n                    }\n                  }\n                }\n              }\n            }\n          }\n        }\n      ]\n    }\n  }\n}\n",
	"code-facts": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/code-facts.schema.json",
  "title": "Recorded code facts",
  "description": "Deterministic symbols and measured language statistics attached to a tool_scan assessment. The assessment owns the scanned ref, producer/version and time. null means not measured; an empty array means measured with no results. Producers normalize symbol enumerator and scc/tokei output before publication; consumers never estimate LoC from language byte counts. At least one section must carry a recorded array/object; all four sections explicit nulls records nothing and does not validate.",
  "type": "object",
  "required": ["artifact", "symbols", "languages"],
  "additionalProperties": false,
  "properties": {
    "artifact": {"const": "code-facts"},
    "symbols": {
      "type": ["array", "null"],
      "items": {"$ref": "#/$defs/symbol"}
    },
    "languages": {
      "type": ["array", "null"],
      "items": {"$ref": "#/$defs/language"}
    },
    "packages": {
      "type": ["object", "null"],
      "required": ["modules", "imports", "exports", "gaps"],
      "additionalProperties": false,
      "properties": {
        "modules": {
          "type": "array",
          "items": {
            "type": "object",
            "required": ["ecosystem", "name", "location"],
            "additionalProperties": false,
            "properties": {
              "ecosystem": {"type": "string", "minLength": 1},
              "name": {"type": "string", "minLength": 1},
              "location": {"$ref": "#/$defs/location"}
            }
          }
        },
        "imports": {
          "type": "array",
          "items": {
            "type": "object",
            "required": ["package", "language", "location"],
            "additionalProperties": false,
            "properties": {
              "package": {"type": "string", "minLength": 1},
              "language": {"type": "string", "minLength": 1},
              "location": {"$ref": "#/$defs/location"}
            }
          }
        },
        "exports": {
          "type": "array",
          "items": {
            "type": "object",
            "required": ["package", "name", "language", "location"],
            "additionalProperties": false,
            "properties": {
              "package": {"type": "string", "minLength": 1},
              "name": {"type": "string", "minLength": 1},
              "kind": {"type": ["string", "null"]},
              "language": {"type": "string", "minLength": 1},
              "location": {"$ref": "#/$defs/location"}
            }
          }
        },
        "gaps": {
          "description": "Why a package/interface fact is unmeasured. system stays free-form by design: it names the producer's own subsystem (an adopter's scanner), unlike interface-role and liveness-status which are shared vocabularies; reason carries the producer's cause.",
          "type": "array",
          "items": {
            "type": "object",
            "required": ["system", "reason"],
            "additionalProperties": false,
            "properties": {
              "system": {"type": "string", "minLength": 1},
              "reason": {"type": "string", "minLength": 1},
              "path": {"type": ["string", "null"]},
              "paths": {"type": "array", "items": {"type": "string"}},
              "count": {"type": ["integer", "null"], "minimum": 0},
              "evidence": {"type": ["string", "null"]}
            }
          }
        }
      }
    },
    "interfaces": {
      "type": ["array", "null"],
      "items": {
        "type": "object",
        "required": ["api_group", "role", "location", "resources", "verbs"],
        "additionalProperties": false,
        "properties": {
          "api_group": {"type": "string"},
          "role": {"enum": ["owner", "consumer", "rbac", "webhook"]},
          "location": {"$ref": "#/$defs/location"},
          "kind": {"type": ["string", "null"]},
          "name": {"type": ["string", "null"]},
          "resources": {"type": "array", "items": {"type": "string"}},
          "verbs": {"type": "array", "items": {"type": "string"}}
        }
      }
    }
  },
  "anyOf": [
    {"required": ["symbols"], "properties": {"symbols": {"type": "array"}}},
    {"required": ["languages"], "properties": {"languages": {"type": "array"}}},
    {"required": ["packages"], "properties": {"packages": {"type": "object"}}},
    {"required": ["interfaces"], "properties": {"interfaces": {"type": "array"}}}
  ],
  "$defs": {
    "location": {
      "type": "object",
      "required": ["path"],
      "additionalProperties": false,
      "properties": {
        "path": {
          "type": "string",
          "minLength": 1,
          "pattern": "^(?!/)(?![A-Za-z]:)(?![\\s\\S]*(?:^|/)\\.{1,2}(?:/|$))[^\\\\\u0000]+$"
        },
        "line": {"type": ["integer", "null"], "minimum": 1},
        "symbol": {"type": ["string", "null"], "minLength": 1},
        "hashes": {
          "type": "array",
          "items": {
            "type": "object",
            "required": ["algorithm", "digest"],
            "additionalProperties": false,
            "properties": {
              "algorithm": {"type": "string", "minLength": 1},
              "digest": {"type": "string", "minLength": 1}
            }
          }
        }
      }
    },
    "symbol": {
      "type": "object",
      "required": ["name", "path", "line", "kind", "language"],
      "additionalProperties": false,
      "properties": {
        "name": {"type": "string", "minLength": 1},
        "path": {
          "type": "string",
          "minLength": 1,
          "pattern": "^(?!/)(?![A-Za-z]:)(?![\\s\\S]*(?:^|/)\\.{1,2}(?:/|$))[^\\\\\u0000]+$"
        },
        "line": {"type": "integer", "minimum": 1},
        "kind": {
          "description": "The producing tool's symbol category, kept as reported.",
          "type": ["string", "null"],
          "minLength": 1
        },
        "language": {"type": ["string", "null"], "minLength": 1}
      }
    },
    "language": {
      "type": "object",
      "required": ["name", "files", "code", "comments", "blanks"],
      "additionalProperties": false,
      "properties": {
        "name": {"type": "string", "minLength": 1},
        "files": {"type": ["integer", "null"], "minimum": 0},
        "code": {"type": "integer", "minimum": 0},
        "comments": {"type": ["integer", "null"], "minimum": 0},
        "blanks": {"type": ["integer", "null"], "minimum": 0}
      }
    }
  }
}
`,
	"compliance-assessment": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/compliance-assessment.schema.json",
  "title": "Compliance assessment artifact",
  "description": "Validates /compliance-check assessment artifacts. The false-positive gates are STRUCTURAL: satisfied/not_satisfied verdicts are unrepresentable without evidence references (conditional schema below), deterministic controls must carry a check-computed verdict (verdict_source=check) unless an attributed human override is present, and not_assessed/not_applicable require machine-readable reasons. The compliance validator additionally verifies evidence content hashes and recomputes the coverage block \u2014 schema + validator together are the citation gate for compliance.",
  "type": "object",
  "required": [
    "metadata",
    "coverage",
    "results"
  ],
  "additionalProperties": false,
  "properties": {
    "metadata": {
      "type": "object",
      "required": [
        "artifact",
        "harness_version",
        "target",
        "frameworks",
        "registry_hash",
        "generated_at"
      ],
      "additionalProperties": true,
      "properties": {
        "artifact": {
          "const": "compliance-assessment"
        },
        "harness_version": {
          "type": "string"
        },
        "target": {
          "type": "object",
          "required": [
            "kind"
          ],
          "additionalProperties": true,
          "properties": {
            "kind": {
              "enum": [
                "product",
                "environment",
                "both"
              ]
            },
            "product": {
              "type": "string"
            },
            "repos": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "environment": {
              "type": "string"
            },
            "scope_binding_mode": {
              "type": "string"
            },
            "snapshot_id": {
              "type": "string",
              "description": "sha256 of the canonicalized environment inventory snapshot this assessment read (determinism amendment 6)"
            },
            "cde_boundary": {
              "type": "string",
              "description": "REQUIRED for PCI runs: path/hash of the human-declared cardholder-data-environment boundary document"
            },
            "personal_data_stores": {
              "type": "array",
              "items": {
                "type": "string"
              },
              "description": "REQUIRED for GDPR runs: declared stores holding personal data \u2014 never inferred"
            },
            "trust_service_categories": {
              "type": "array",
              "items": {
                "enum": [
                  "security",
                  "availability",
                  "confidentiality",
                  "processing_integrity",
                  "privacy"
                ]
              },
              "description": "REQUIRED for SOC 2 runs: declared engagement categories (security mandatory)"
            },
            "scope": {
              "type": "object",
              "description": "Boundary provenance when the run was scoped via the compliance scope registry (--boundary; Phase 6): the declared boundary id, who declared it and when, the resolution mode, excluded repos with rationale, and whether the declaration is still an unsigned draft.",
              "required": [
                "boundary",
                "resolves_via",
                "declared_by",
                "declared_at",
                "draft"
              ],
              "additionalProperties": false,
              "properties": {
                "boundary": {
                  "type": "string"
                },
                "resolves_via": {
                  "enum": [
                    "repo-graph",
                    "explicit"
                  ]
                },
                "declared_by": {
                  "type": "string"
                },
                "declared_at": {
                  "type": "string"
                },
                "excluded": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "required": [
                      "repo",
                      "reason"
                    ],
                    "additionalProperties": false,
                    "properties": {
                      "repo": {
                        "type": "string"
                      },
                      "reason": {
                        "type": "string"
                      }
                    }
                  }
                },
                "draft": {
                  "type": "boolean"
                },
                "registry_hash": {
                  "type": [
                    "string",
                    "null"
                  ]
                }
              }
            }
          }
        },
        "frameworks": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "object",
            "required": [
              "id"
            ],
            "additionalProperties": false,
            "properties": {
              "id": {
                "enum": [
                  "nist-800-53-rev5",
                  "fedramp-high",
                  "fedramp-moderate",
                  "pci-dss-v4",
                  "soc2-tsc",
                  "gdpr-technical"
                ]
              },
              "profile": {
                "type": "string"
              },
              "caveat": {
                "type": "string",
                "description": "mandatory framework caveats (FedRAMP interim 800-53B posture; SOC 2 point-in-time vs Type 2 period)"
              }
            }
          }
        },
        "registry_hash": {
          "type": "string",
          "pattern": "^[0-9a-f]{64}$",
          "description": "sha256 of the canonicalized compliance-mapping registry used \u2014 drift attribution (amendment 6)"
        },
        "org_parameters_hash": {
          "type": "string",
          "pattern": "^[0-9a-f]{64}$"
        },
        "collector_versions": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "generated_at": {
          "type": "string"
        },
        "owner_attestations": {
          "type": "array",
          "description": "Owner context captured by /compliance-check interview (Phase 7). RAIL 1 \u2014 attestations are EVIDENCE, never verdicts: they may inform evidence_review judging (cited by id from a result's evidence) and may upgrade not_assessed -> evidence_review by making a control assessable; they can never touch a deterministic verdict and can never directly set satisfied. Attribution is rendered wherever a verdict they informed is shown.",
          "items": {
            "type": "object",
            "required": [
              "id",
              "topic",
              "statement",
              "attested_by",
              "ldap_verified",
              "attested_at"
            ],
            "additionalProperties": false,
            "properties": {
              "id": {
                "type": "string",
                "pattern": "^att-[a-z0-9-]+$"
              },
              "topic": {
                "type": "string",
                "minLength": 1,
                "description": "what the statement is about (e.g. data-flow, environment, compensating-control)"
              },
              "statement": {
                "type": "string",
                "minLength": 1,
                "description": "the owner's words, verbatim"
              },
              "applies_to": {
                "type": "array",
                "items": {
                  "type": "string"
                },
                "description": "framework:control refs this informs"
              },
              "attested_by": {
                "type": "string",
                "minLength": 1
              },
              "ldap_verified": {
                "type": "boolean"
              },
              "attested_at": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "coverage": {
      "type": "object",
      "description": "Per-framework coverage honesty block \u2014 recomputed and cross-checked by the validator; no blended compliance percentage exists anywhere in this artifact.",
      "minProperties": 1,
      "additionalProperties": {
        "type": "object",
        "required": [
          "total_in_scope",
          "deterministic",
          "evidence_review",
          "organizational",
          "satisfied",
          "not_satisfied",
          "not_applicable",
          "not_assessed"
        ],
        "additionalProperties": false,
        "properties": {
          "total_in_scope": {
            "type": "integer",
            "minimum": 0
          },
          "deterministic": {
            "type": "integer",
            "minimum": 0
          },
          "evidence_review": {
            "type": "integer",
            "minimum": 0
          },
          "organizational": {
            "type": "integer",
            "minimum": 0
          },
          "satisfied": {
            "type": "integer",
            "minimum": 0
          },
          "not_satisfied": {
            "type": "integer",
            "minimum": 0
          },
          "not_applicable": {
            "type": "integer",
            "minimum": 0
          },
          "not_assessed": {
            "type": "integer",
            "minimum": 0
          }
        }
      }
    },
    "results": {
      "type": "array",
      "items": {
        "type": "object",
        "required": [
          "framework",
          "control_id",
          "classification",
          "verdict",
          "verdict_source"
        ],
        "additionalProperties": false,
        "properties": {
          "framework": {
            "type": "string"
          },
          "control_id": {
            "type": "string"
          },
          "title": {
            "type": "string"
          },
          "classification": {
            "enum": [
              "deterministic",
              "evidence_review",
              "organizational"
            ]
          },
          "verdict": {
            "enum": [
              "satisfied",
              "not_satisfied",
              "not_applicable",
              "not_assessed"
            ]
          },
          "verdict_source": {
            "enum": [
              "check",
              "agent",
              "human_override"
            ],
            "description": "deterministic controls: 'check' (code-computed) or 'human_override' \u2014 never 'agent' (amendment 1)"
          },
          "check_id": {
            "type": "string"
          },
          "reason": {
            "type": "string"
          },
          "narrative": {
            "type": "string",
            "description": "agent annotation \u2014 context and remediation guidance only; carries no evidentiary weight and cannot change the verdict"
          },
          "evidence": {
            "type": "array",
            "items": {
              "type": "object",
              "required": [
                "sha256",
                "kind",
                "locator"
              ],
              "additionalProperties": false,
              "properties": {
                "sha256": {
                  "type": "string",
                  "pattern": "^[0-9a-f]{64}$"
                },
                "kind": {
                  "enum": [
                    "inventory_snapshot_excerpt",
                    "iac_file",
                    "scanner_fact",
                    "policy_object",
                    "human_artifact"
                  ]
                },
                "locator": {
                  "type": "string",
                  "description": "where the evidence came from: file:line, snapshot path, fact id, artifact ref"
                },
                "excerpt": {
                  "type": "string",
                  "maxLength": 2000
                }
              }
            }
          },
          "override": {
            "type": "object",
            "description": "REQUIRED when verdict_source=human_override \u2014 the countersign pattern",
            "required": [
              "by",
              "date",
              "rationale"
            ],
            "additionalProperties": false,
            "properties": {
              "by": {
                "type": "string"
              },
              "date": {
                "type": "string"
              },
              "rationale": {
                "type": "string"
              },
              "overridden_check_verdict": {
                "type": "string"
              }
            }
          },
          "n_pass_agreement": {
            "type": "object",
            "description": "evidence_review controls: N-pass stability record (amendment 5)",
            "required": [
              "passes",
              "agreed"
            ],
            "additionalProperties": false,
            "properties": {
              "passes": {
                "type": "integer",
                "minimum": 1
              },
              "agreed": {
                "type": "boolean"
              }
            }
          }
        },
        "allOf": [
          {
            "if": {
              "properties": {
                "verdict": {
                  "enum": [
                    "satisfied",
                    "not_satisfied"
                  ]
                }
              }
            },
            "then": {
              "required": [
                "evidence"
              ],
              "properties": {
                "evidence": {
                  "minItems": 1
                }
              }
            }
          },
          {
            "if": {
              "properties": {
                "verdict": {
                  "enum": [
                    "not_applicable",
                    "not_assessed"
                  ]
                }
              }
            },
            "then": {
              "required": [
                "reason"
              ]
            }
          },
          {
            "if": {
              "properties": {
                "verdict_source": {
                  "const": "human_override"
                }
              }
            },
            "then": {
              "required": [
                "override"
              ]
            }
          }
        ]
      }
    }
  },
  "$defs": {
    "canonical_write": {
      "description": "Opt-in reuse of the existing framework-id and compliance-verdict domains; historical result text remains readable by default.",
      "allOf": [
        { "$ref": "#" },
        {
          "properties": {
            "results": {
              "items": {
                "properties": {
                  "framework": {
                    "$ref": "#/properties/metadata/properties/frameworks/items/properties/id"
                  },
                  "override": {
                    "properties": {
                      "overridden_check_verdict": {
                        "$ref": "#/properties/results/items/properties/verdict"
                      }
                    }
                  }
                }
              }
            }
          }
        }
      ]
    }
  }
}
`,
	"compliance-mapping": "{\n  \"$schema\": \"http://json-schema.org/draft-07/schema#\",\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/compliance-mapping.schema.json\",\n  \"title\": \"Compliance control-mapping registry\",\n  \"description\": \"Validates your compliance configuration directory's control-mapping registry \\u2014 the policy-data registry that pre-classifies every in-scope control (deterministic / evidence_review / organizational) and defines deterministic checks as DECLARATIVE ASSERTION RECORDS evaluated by the assertion evaluator (determinism amendment 2: checks are diffable data, not per-check code). The LLM never classifies a control or authors an assertion at run time \\u2014 changing either is a reviewed config change.\",\n  \"type\": \"object\",\n  \"required\": [\n    \"version\",\n    \"controls\",\n    \"checks\"\n  ],\n  \"additionalProperties\": false,\n  \"properties\": {\n    \"version\": {\n      \"type\": \"integer\",\n      \"minimum\": 1\n    },\n    \"note\": {\n      \"type\": \"string\"\n    },\n    \"controls\": {\n      \"type\": \"array\",\n      \"minItems\": 1,\n      \"items\": {\n        \"type\": \"object\",\n        \"required\": [\n          \"framework\",\n          \"control_id\",\n          \"classification\"\n        ],\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"framework\": {\n            \"enum\": [\n              \"nist-800-53-rev5\",\n              \"pci-dss-v4\",\n              \"soc2-tsc\",\n              \"gdpr-technical\"\n            ]\n          },\n          \"control_id\": {\n            \"type\": \"string\"\n          },\n          \"classification\": {\n            \"enum\": [\n              \"deterministic\",\n              \"evidence_review\",\n              \"organizational\"\n            ]\n          },\n          \"checks\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\"\n            },\n            \"description\": \"check ids \\u2014 REQUIRED non-empty when classification=deterministic\"\n          },\n          \"evidence_request\": {\n            \"type\": \"string\",\n            \"description\": \"evidence_review controls: what human-supplied artifact the agent judges\"\n          },\n          \"note\": {\n            \"type\": \"string\"\n          },\n          \"crosswalk\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\",\n              \"pattern\": \"^[a-z0-9-]+:[A-Za-z0-9 ().-]+$\"\n            },\n            \"description\": \"hand-authored framework:control_id tags (e.g. 'nist-800-53-rev5:ia-2.1') \\u2014 never vendored from restricted mappings\"\n          },\n          \"applicability\": {\n            \"$ref\": \"#/$defs/assertion\",\n            \"description\": \"declared not_applicable condition (amendment 4) \\u2014 evaluated over the snapshot; false \\u21d2 not_applicable, never narrated\"\n          },\n          \"parameters\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\"\n            },\n            \"description\": \"org-parameter ids this control's checks consume; any undeclared \\u21d2 not_assessed(parameter undeclared)\"\n          }\n        },\n        \"allOf\": [\n          {\n            \"if\": {\n              \"properties\": {\n                \"classification\": {\n                  \"const\": \"deterministic\"\n                }\n              }\n            },\n            \"then\": {\n              \"required\": [\n                \"checks\"\n              ],\n              \"properties\": {\n                \"checks\": {\n                  \"minItems\": 1\n                }\n              }\n            }\n          },\n          {\n            \"if\": {\n              \"properties\": {\n                \"classification\": {\n                  \"const\": \"evidence_review\"\n                }\n              }\n            },\n            \"then\": {\n              \"required\": [\n                \"evidence_request\"\n              ]\n            }\n          }\n        ]\n      }\n    },\n    \"checks\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"type\": \"object\",\n        \"required\": [\n          \"id\",\n          \"collector\",\n          \"assertion\"\n        ],\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"id\": {\n            \"type\": \"string\",\n            \"pattern\": \"^chk-[a-z0-9-]+$\"\n          },\n          \"description\": {\n            \"type\": \"string\"\n          },\n          \"collector\": {\n            \"type\": \"string\",\n            \"description\": \"snapshot kind this check reads (scan_k8s_hardening | run_gitleaks | run_osv_scanner | run_govulncheck | cloud_inventory | sbom | cosign)\"\n          },\n          \"assertion\": {\n            \"$ref\": \"#/$defs/assertion\"\n          },\n          \"applies_when\": {\n            \"$ref\": \"#/$defs/assertion\"\n          },\n          \"sampling\": {\n            \"type\": \"string\",\n            \"description\": \"REQUIRED if the check does not evaluate ALL selected values: a deterministic rule ('sorted_first_50') \\u2014 'representative' is not expressible (amendment 7)\",\n            \"pattern\": \"^sorted_first_\\\\d+$\"\n          },\n          \"fixtures\": {\n            \"type\": \"string\",\n            \"description\": \"path to the known-bad/known-good fixture pair (Phase 4 admission gate)\"\n          },\n          \"parameter_refs\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\"\n            },\n            \"description\": \"org-parameter ids whose values the assertion's `expected` resolves from ('param:<id>')\"\n          }\n        }\n      }\n    }\n  },\n  \"$defs\": {\n    \"assertion\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"path\",\n        \"operator\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"path\": {\n          \"type\": \"string\",\n          \"description\": \"dot path into the canonical snapshot; [*] fans out \\u2014 the assertion must hold for EVERY selected value (no sampling unless `sampling` is declared)\"\n        },\n        \"operator\": {\n          \"enum\": [\n            \"equals\",\n            \"not_equals\",\n            \"contains\",\n            \"not_contains\",\n            \"regex\",\n            \"not_regex\",\n            \"gte\",\n            \"lte\",\n            \"exists\",\n            \"absent\",\n            \"in\",\n            \"not_in\"\n          ]\n        },\n        \"expected\": {\n          \"description\": \"literal, or 'param:<org-parameter-id>' resolved from org-parameters.yaml at evaluation time\"\n        }\n      }\n    }\n  }\n}\n",
	"compliance-scope": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/compliance-scope.schema.json",
  "title": "Compliance scope registry",
  "description": "Declared compliance boundaries (your compliance configuration directory's compliance-scope file). Frameworks scope SYSTEMS, not repos: each boundary is a reviewed, signed declaration of what sits inside a PCI CDE / FedRAMP authorization boundary / SOC 2 system, declared at service/product level. Repo membership is RESOLVED (repo-graph product mapping via the scope resolver), never enumerated by hand beyond the include/exclude edges \u2014 graphs resolve scope, they never store it. Same doctrine as compliance-mapping.yaml: policy data, not prompt text; every edit is a reviewed config change; declared_at is dated provenance drift-watch ages.",
  "type": "object",
  "required": [
    "version",
    "updated",
    "boundaries"
  ],
  "additionalProperties": false,
  "properties": {
    "version": {
      "type": "integer",
      "minimum": 1
    },
    "updated": {
      "type": "string",
      "pattern": "^\\d{4}-\\d{2}-\\d{2}"
    },
    "boundaries": {
      "type": "object",
      "minProperties": 1,
      "additionalProperties": {
        "type": "object",
        "required": [
          "frameworks",
          "resolves_via",
          "declared_by",
          "declared_at"
        ],
        "additionalProperties": false,
        "properties": {
          "frameworks": {
            "type": "array",
            "minItems": 1,
            "items": {
              "enum": [
                "pci-dss-v4",
                "nist-800-53-rev5",
                "fedramp-high",
                "fedramp-moderate",
                "soc2-tsc",
                "gdpr-technical"
              ]
            }
          },
          "resolves_via": {
            "enum": [
              "repo-graph",
              "deployment-iac",
              "explicit"
            ],
            "description": "repo-graph: membership = the product node's ships edges plus include minus exclude (organizational assertion). deployment-iac: membership = the repos a declared deployment inventory cites as deployed for the service, every row carrying its IaC source (deployment EVIDENCE \u2014 prefer when available; adopted 2026-07-30). explicit: include[] only."
          },
          "product": {
            "type": "string",
            "description": "repo-graph product node id (e.g. product:adhoc:stackrox) or exact label; required when resolves_via=repo-graph."
          },
          "include": {
            "type": "array",
            "items": {
              "type": "object",
              "required": [
                "repo",
                "reason"
              ],
              "additionalProperties": false,
              "properties": {
                "repo": {
                  "type": "string",
                  "pattern": "^[\\w.-]+/[\\w.-]+$",
                  "description": "org/name"
                },
                "reason": {
                  "type": "string",
                  "minLength": 1
                }
              }
            }
          },
          "exclude": {
            "type": "array",
            "items": {
              "type": "object",
              "required": [
                "repo",
                "reason"
              ],
              "additionalProperties": false,
              "properties": {
                "repo": {
                  "type": "string",
                  "pattern": "^[\\w.-]+/[\\w.-]+$"
                },
                "reason": {
                  "type": "string",
                  "minLength": 1,
                  "description": "Boundary rationale \u2014 'doesn't apply' claims live HERE with a reason, never as silent omissions."
                }
              }
            }
          },
          "declared_by": {
            "type": "string",
            "minLength": 1,
            "description": "Identity handle of the human who owns this boundary claim. 'draft:<id>' marks an unsigned draft \u2014 the runner labels its output accordingly."
          },
          "declared_at": {
            "type": "string",
            "pattern": "^\\d{4}-\\d{2}-\\d{2}"
          },
          "notes": {
            "type": "string"
          },
          "deployment_evidence": {
            "type": "object",
            "description": "Required when resolves_via=deployment-iac: the declared deployment inventory (JSON: {services: {<name>: {repos: [{url, source}]}}}, produced from a configuration repository checkout \u2014 extractor lands with the first real engagement) and the service to resolve.",
            "required": [
              "inventory",
              "service"
            ],
            "additionalProperties": false,
            "properties": {
              "inventory": {
                "type": "string",
                "description": "path, relative to the scope file"
              },
              "service": {
                "type": "string",
                "minLength": 1
              }
            }
          }
        }
      }
    }
  }
}
`,
	"corpus-registry": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/corpus-registry.schema.json",
  "title": "Corpus subject registry",
  "description": "Who each audited subject belongs to. Ownership decides which denominator a finding counts toward, and until now it was declared only in the deployment's corpus-config and materialised in the harness's own SQLite projection -- so no contract consumer could compute 'X% of our repos' at all. This registry is a configuration document modelled as an artifact, the same shape as compliance-scope and org-parameters: the deployment declares it, the harness emits it from the corpus resolution, and storage projects it so findings can be joined to ownership in SQL.",
  "type": "object",
  "required": ["version", "subjects"],
  "additionalProperties": false,
  "properties": {
    "version": {
      "type": "integer",
      "minimum": 1
    },
    "updated": {
      "type": "string",
      "description": "When this registry was last regenerated from the corpus resolution."
    },
    "note": {
      "type": "string"
    },
    "subjects": {
      "type": "array",
      "items": { "$ref": "#/$defs/subject" }
    }
  },
  "$defs": {
    "subject": {
      "type": "object",
      "required": ["subject_id", "tree", "ownership", "business_unit"],
      "additionalProperties": false,
      "properties": {
        "subject_id": {
          "type": "string",
          "minLength": 1,
          "description": "The same opaque subject identifier used in artifact_binding.subject_id. This is the join key: without it, ownership cannot reach a finding."
        },
        "tree": {
          "type": "string",
          "minLength": 1,
          "description": "The corpus tree this subject was resolved from. Kept because ownership is declared per tree in corpus-config, so this is the provenance of the ownership value."
        },
        "ownership": {
          "type": "string",
          "enum": ["owned", "upstream", "external-bu", "harness-qa"],
          "description": "See enums/v1/ownership.json. 'harness-qa' subjects are registered but excluded from every metrics lens."
        },
        "business_unit": {
          "type": "string",
          "minLength": 1
        },
        "label": {
          "type": "string",
          "description": "Engagement or tree label, for grouping cuts that are narrower than a business unit."
        },
        "product": {
          "type": "string",
          "description": "Product this subject ships in, when the inventory declares one."
        },
        "repo_url": {
          "type": "string"
        },
        "ref": {
          "type": "string",
          "description": "Git ref audited, when this is a ref-specific audit rather than HEAD."
        },
        "ref_kind": {
          "type": "string",
          "enum": ["branch", "tag", "default", "stream"]
        },
        "is_branch_audit": {
          "type": "boolean",
          "description": "Whether this subject is a re-audit of a non-default ref. Load-bearing for every count: a large share of audits are branch re-audits of the same code, so a denominator that does not exclude them overstates coverage and double-counts exposure."
        }
      }
    }
  }
}
`,
	"doc-variance": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/doc-variance.schema.json",
  "title": "Documentation-variance register",
  "description": "Structured records of OFFICIAL-documentation claims that code evidence contradicts (<repo>-doc-variance.json, emitted where verification happens \u2014 threat-model context passes, audits). Source restriction is deliberate: only docs.redhat.com product documentation qualifies as a variance source \u2014 informal inputs (owner notes, Google Docs) inform threat models but never mint variance records; their claims must be traced to (or contradicted by) the official docs to register here. A variance record is a claim-vs-evidence discrepancy, NEVER a vulnerability verdict \u2014 audits/scans convert security-relevant ones into findings through the adjudicated flow; the register additionally tracks the docs-side resolution (doc_corrected) that no finding lifecycle models.",
  "type": "object",
  "required": [
    "metadata",
    "records"
  ],
  "additionalProperties": false,
  "properties": {
    "metadata": {
      "type": "object",
      "required": [
        "repository",
        "created",
        "harness_version"
      ],
      "additionalProperties": false,
      "properties": {
        "repository": {
          "type": "string",
          "format": "uri"
        },
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "updated": {
          "type": "string",
          "format": "date-time"
        },
        "harness_version": {
          "type": "string"
        }
      }
    },
    "records": {
      "type": "array",
      "items": {
        "type": "object",
        "required": [
          "id",
          "source",
          "claim",
          "code_evidence",
          "verified_at",
          "disposition"
        ],
        "additionalProperties": false,
        "properties": {
          "id": {
            "type": "string",
            "pattern": "^dv-[a-z0-9-]+$"
          },
          "source": {
            "type": "object",
            "description": "The OFFICIAL documentation location making the claim \u2014 one record per (claim, doc version) so multi-version variance is first-class.",
            "required": [
              "product_slug",
              "version",
              "url"
            ],
            "additionalProperties": false,
            "properties": {
              "product_slug": {
                "type": "string",
                "description": "docs.redhat.com slug, joined to repos via your product-to-documentation mapping file"
              },
              "version": {
                "type": "string",
                "description": "documentation version (e.g. 4.22, 1-latest); joins to code refs via metadata.ref where release-aligned"
              },
              "guide": {
                "type": "string"
              },
              "url": {
                "type": "string",
                "pattern": "^https://docs\\.redhat\\.com/",
                "description": "canonical section URL \u2014 the schema-enforced official-docs-only restriction"
              },
              "quote": {
                "type": "string",
                "minLength": 1,
                "maxLength": 600,
                "description": "short verbatim excerpt (CC-BY-SA attribution rides the url)"
              }
            }
          },
          "claim": {
            "type": "string",
            "minLength": 1,
            "description": "the documentation's assertion, paraphrased precisely"
          },
          "code_evidence": {
            "type": "array",
            "minItems": 1,
            "items": {
              "type": "object",
              "required": [
                "repo",
                "path"
              ],
              "additionalProperties": false,
              "properties": {
                "repo": {
                  "type": "string"
                },
                "ref": {
                  "type": "string",
                  "description": "branch/tag when version-specific (joins doc version to release branch)"
                },
                "path": {
                  "type": "string"
                },
                "lines": {
                  "type": "string"
                },
                "note": {
                  "type": "string"
                }
              }
            }
          },
          "variance": {
            "enum": [
              "overclaim",
              "underclaim",
              "omission",
              "contradiction",
              "stale"
            ],
            "description": "overclaim = docs promise more security than code delivers (the customer-trust class)"
          },
          "verified_at": {
            "type": "string",
            "format": "date-time"
          },
          "verified_against": {
            "type": "string",
            "description": "repo@sha the evidence was read at"
          },
          "disposition": {
            "enum": [
              "open",
              "doc_corrected",
              "code_fixed",
              "accepted",
              "superseded"
            ]
          },
          "disposition_note": {
            "type": "string"
          },
          "finding_refs": {
            "type": "array",
            "items": {
              "type": "string"
            },
            "description": "campaign finding ids minted when an audit/scan converted this record"
          },
          "threat_refs": {
            "type": "array",
            "items": {
              "type": "string"
            },
            "description": "threat-model rows (e.g. TM-042) that carry this discrepancy"
          }
        }
      }
    }
  }
}
`,
	"fleet-fix": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/fleet-fix.schema.json",
  "title": "Fleet-fix transform specification",
  "description": "One reviewed transform applied across every repo affected by a systemic (Lens 3) pattern \u2014 capability C5. The spec IS the change: it must carry its own golden tests, and apply_fleet_fix.py refuses to run a spec whose tests fail. Appliers only modify working trees and emit diffs; committing, forking, and MR creation are the skill's later, human-gated stages.",
  "type": "object",
  "required": [
    "id",
    "pattern_ref",
    "description",
    "matcher",
    "rewrite",
    "guards",
    "tests"
  ],
  "additionalProperties": false,
  "properties": {
    "id": {
      "type": "string",
      "pattern": "^[a-z0-9][a-z0-9-]*$"
    },
    "pattern_ref": {
      "type": "string",
      "description": "The systemic pattern this fixes (insecure-patterns CWE entry or PROGRESS.md rollup name)"
    },
    "description": {
      "type": "string"
    },
    "matcher": {
      "type": "object",
      "required": [
        "kind",
        "file_glob"
      ],
      "additionalProperties": false,
      "properties": {
        "kind": {
          "enum": [
            "pinned_ref_line",
            "ast_grep"
          ]
        },
        "file_glob": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "string"
          },
          "description": "Repo-relative globs the matcher may touch \u2014 everything else is out of bounds"
        },
        "context_regex": {
          "type": "string",
          "description": "pinned_ref_line: a line matching this must appear within context_window lines BEFORE the match line"
        },
        "context_window": {
          "type": "integer",
          "minimum": 1,
          "default": 3
        },
        "match_regex": {
          "type": "string",
          "description": "pinned_ref_line: the line to rewrite; group 1 must capture leading indentation, group 2 the mutable ref"
        },
        "url_regex": {
          "type": "string",
          "description": "pinned_ref_line: nearest PRECEDING line whose group 1 captures the git URL the resolver pins against"
        },
        "url_window": {
          "type": "integer",
          "minimum": 1,
          "default": 10
        },
        "pattern": {
          "type": "string",
          "description": "ast_grep: the --pattern expression"
        },
        "lang": {
          "type": "string",
          "description": "ast_grep: --lang value"
        }
      }
    },
    "resolver": {
      "type": "object",
      "additionalProperties": false,
      "properties": {
        "kind": {
          "enum": [
            "git_ls_remote"
          ]
        },
        "ref": {
          "type": "string",
          "default": "main",
          "description": "remote ref whose current SHA replaces the mutable ref (per matched URL)"
        }
      }
    },
    "rewrite": {
      "type": "object",
      "required": [
        "template"
      ],
      "additionalProperties": false,
      "properties": {
        "template": {
          "type": "string",
          "description": "Replacement line. Placeholders: {indent} {sha} {short_sha} {ref} {date} {id}. ast_grep kind uses this as the --rewrite expression instead."
        }
      }
    },
    "guards": {
      "type": "object",
      "additionalProperties": false,
      "properties": {
        "max_files_changed": {
          "type": "integer",
          "minimum": 1
        },
        "require_clean_tree": {
          "type": "boolean",
          "default": true
        }
      }
    },
    "tests": {
      "type": "array",
      "minItems": 1,
      "description": "Golden tests run before ANY repo is touched (with a pinned test SHA); a failing test aborts the whole fleet application",
      "items": {
        "type": "object",
        "required": [
          "name",
          "file",
          "before",
          "after_contains"
        ],
        "additionalProperties": false,
        "properties": {
          "name": {
            "type": "string"
          },
          "file": {
            "type": "string",
            "description": "repo-relative path the fixture is written to (must match file_glob)"
          },
          "before": {
            "type": "string"
          },
          "after_contains": {
            "type": "array",
            "minItems": 1,
            "items": {
              "type": "string"
            }
          },
          "after_not_contains": {
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        }
      }
    }
  }
}
`,
	"fuzz-campaign": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/fuzz-campaign.schema.json",
  "title": "Supplied fuzz campaign scope reference",
  "description": "A supplied campaign snapshot, not fuzz execution evidence. Aggregate populations and analyst prose are retained with declared references; historical pattern counts are never redistributed into fabricated per-target observations.",
  "type": "object",
  "additionalProperties": false,
  "required": [
    "artifact",
    "campaign_id",
    "generated",
    "repositories",
    "source_references"
  ],
  "properties": {
    "artifact": {
      "const": "fuzz-campaign"
    },
    "campaign_id": {
      "type": "string",
      "minLength": 1
    },
    "generated": {
      "type": "string",
      "format": "date"
    },
    "repositories": {
      "type": "array",
      "minItems": 1,
      "uniqueItems": true,
      "items": {
        "type": "string",
        "minLength": 1
      }
    },
    "source_references": {
      "type": "array",
      "uniqueItems": true,
      "items": {
        "type": "string",
        "minLength": 1
      },
      "minItems": 1
    },
    "narrative": {
      "type": "string",
      "minLength": 1
    },
    "harness_version": {
      "type": "string",
      "minLength": 1
    },
    "manifest_generated": {
      "type": "string",
      "format": "date"
    },
    "duration": {
      "type": "object",
      "additionalProperties": false,
      "required": [
        "start",
        "end"
      ],
      "properties": {
        "start": {
          "type": "string",
          "format": "date"
        },
        "end": {
          "type": "string",
          "format": "date"
        }
      }
    },
    "totals": {
      "type": "object",
      "additionalProperties": false,
      "required": [],
      "properties": {
        "targets": {
          "type": [
            "integer",
            "null"
          ],
          "minimum": 0
        },
        "repos": {
          "type": [
            "integer",
            "null"
          ],
          "minimum": 0
        },
        "fuzzed_targets": {
          "type": [
            "integer",
            "null"
          ],
          "minimum": 0
        },
        "fuzz_functions": {
          "type": [
            "integer",
            "null"
          ],
          "minimum": 0
        },
        "executed_targets": {
          "type": [
            "integer",
            "null"
          ],
          "minimum": 0
        },
        "runs_logged": {
          "type": [
            "integer",
            "null"
          ],
          "minimum": 0
        },
        "bugs": {
          "type": [
            "integer",
            "null"
          ],
          "minimum": 0
        },
        "bugs_medium_plus": {
          "type": [
            "integer",
            "null"
          ],
          "minimum": 0
        },
        "bugs_by_fuzz": {
          "type": [
            "integer",
            "null"
          ],
          "minimum": 0
        },
        "advisories": {
          "type": [
            "integer",
            "null"
          ],
          "minimum": 0
        },
        "hit_rate": {
          "type": "string",
          "minLength": 1,
          "description": "Single encoding of the hit-rate fact (narrative form, e.g. \"1/2\"); no duplicate numeric encoding ships in the same document."
        }
              }
    },
    "batches": {
      "type": "array",
      "items": {
        "$ref": "#/$defs/batch"
      }
    },
    "bugs": {
      "type": "array",
      "items": {
        "$ref": "#/$defs/bug"
      }
    },
    "patterns": {
      "type": "array",
      "items": {
        "$ref": "#/$defs/pattern"
      }
    },
    "remediation_priority": {
      "type": "string",
      "minLength": 1
    },
    "negative_controls": {
      "type": "string",
      "minLength": 1
    },
    "advisory_files": {
      "type": "array",
      "items": {
        "type": "string",
        "minLength": 1
      }
    }
  },
  "$defs": {
    "batch": {
      "type": "object",
      "additionalProperties": false,
      "required": [
        "batch",
        "targets",
        "fuzz_functions",
        "bugs"
      ],
      "properties": {
        "batch": {
          "type": "string",
          "minLength": 1
        },
        "targets": {
          "type": [
            "integer",
            "null"
          ],
          "minimum": 0
        },
        "fuzz_functions": {
          "type": [
            "integer",
            "null"
          ],
          "minimum": 0
        },
        "bugs": {
          "type": [
            "integer",
            "null"
          ],
          "minimum": 0
        }
      }
    },
    "bug": {
      "type": "object",
      "additionalProperties": false,
      "required": [
        "num",
        "target",
        "title"
      ],
      "properties": {
        "num": {
          "type": "integer",
          "minimum": 1
        },
        "target": {
          "type": "string",
          "minLength": 1
        },
        "title": {
          "type": "string",
          "minLength": 1
        },
        "cvss": {
          "type": "string",
          "minLength": 1
        },
        "cwe": {
          "type": "string",
          "minLength": 1
        },
        "advisory": {
          "type": [
            "string",
            "null"
          ]
        },
        "method": {
          "type": "string",
          "minLength": 1
        },
        "narrative": {
          "type": "string",
          "minLength": 1
        },
        "references": {
          "type": "array",
          "uniqueItems": true,
          "items": {
            "type": "string",
            "minLength": 1
          }
        }
      }
    },
    "pattern": {
      "description": "Supplied legacy aggregate counts per pattern. The content gate (validate.py) rejects a pattern whose vulnerable count exceeds its tested count.",
      "type": "object",
      "additionalProperties": false,
      "required": [
        "name",
        "tested",
        "vulnerable"
     ],
      "properties": {
        "name": {
          "type": "string",
          "minLength": 1
        },
        "tested": {
          "type": "integer",
          "minimum": 0
        },
        "vulnerable": {
          "type": "integer",
          "minimum": 0
        },
        "rate": {
          "type": "string",
          "minLength": 1
        },
        "advisory": {
          "type": [
            "string",
            "null"
          ]
        },
        "references": {
          "type": "array",
          "uniqueItems": true,
          "items": {
            "type": "string",
            "minLength": 1
          }
        }
      }
    }
  }
}
`,
	"fuzz-report": "{\n  \"$schema\": \"https://json-schema.org/draft/2020-12/schema\",\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/fuzz-report.schema.json\",\n  \"title\": \"Fuzz report\",\n  \"description\": \"One fuzz run of one target, held by an assessment of kind `fuzz` (run-bound: the assessment owns the repository, commit, producer and time, so they are not repeated here). Records what ran, every crasher that is a candidate finding, measured coverage, and the surfaces that were not fuzzed and why. Crashers enter the finding lane through intake, one finding per crasher; this document is their evidence. Only fuzz execution produces a crasher: bugs confirmed by static review or a grep sweep are not recorded here. A run that found nothing records no crasher, never a claim of correctness. Run counts (clean, crashed, failed) derive from the harness outcomes and are not stored.\",\n  \"type\": \"object\",\n  \"required\": [\"artifact\", \"target\", \"crashers\", \"coverage\", \"gaps\"],\n  \"additionalProperties\": false,\n  \"properties\": {\n    \"artifact\": {\"const\": \"fuzz-report\"},\n    \"target\": {\"$ref\": \"#/$defs/target\"},\n    \"crashers\": {\"type\": \"array\", \"items\": {\"$ref\": \"#/$defs/crasher\"}},\n    \"coverage\": {\n      \"oneOf\": [{\"$ref\": \"#/$defs/coverage\"}, {\"type\": \"null\"}],\n      \"description\": \"Code coverage of the final corpus, measured by a coverage tool. null: not measured.\"\n    },\n    \"gaps\": {\"type\": \"array\", \"items\": {\"$ref\": \"#/$defs/gap\"}},\n    \"batch\": {\"type\": \"string\", \"minLength\": 1},\n    \"patterns\": {\"type\": \"array\", \"items\": {\"$ref\": \"#/$defs/pattern_observation\"}}\n  },\n  \"$defs\": {\n    \"crasher_analysis\": {\n      \"type\": \"object\",\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"narrative\": {\"type\": \"string\", \"minLength\": 1},\n        \"narrative_references\": {\"type\": \"array\", \"uniqueItems\": true, \"items\": {\"type\": \"string\", \"minLength\": 1}},\n        \"fix\": {\"type\": \"string\", \"minLength\": 1},\n        \"priority\": {\"type\": \"string\", \"minLength\": 1},\n        \"negative_control_references\": {\"type\": \"array\", \"uniqueItems\": true, \"items\": {\"type\": \"string\", \"minLength\": 1}}\n      }\n    },\n    \"pattern_observation\": {\n      \"type\": \"object\",\n      \"required\": [\"name\", \"tested\", \"vulnerable\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"name\": {\"type\": \"string\", \"minLength\": 1},\n        \"tested\": {\"type\": \"boolean\"},\n        \"vulnerable\": {\"type\": [\"boolean\", \"null\"]},\n        \"reference\": {\"type\": \"string\", \"minLength\": 1},\n        \"advisory\": {\"type\": \"string\", \"minLength\": 1}\n      },\n      \"allOf\": [\n        {\n          \"if\": {\"properties\": {\"vulnerable\": {\"type\": \"boolean\"}}, \"required\": [\"vulnerable\"]},\n          \"then\": {\"properties\": {\"tested\": {\"const\": true}}}\n        }\n      ]\n    },\n    \"name\": {\"type\": \"string\", \"pattern\": \"^[A-Za-z0-9][A-Za-z0-9._:/-]*$\"},\n    \"count\": {\"type\": [\"integer\", \"null\"], \"minimum\": 0},\n    \"target\": {\n      \"type\": \"object\",\n      \"required\": [\"id\", \"language\", \"harnesses\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"id\": {\"$ref\": \"#/$defs/name\", \"description\": \"The campaign's target id (one repository, or a refined re-fuzz sharing a clone).\"},\n        \"language\": {\"type\": \"string\", \"minLength\": 1},\n        \"module\": {\"type\": \"string\", \"minLength\": 1, \"description\": \"Module or package coordinate the harnesses import (e.g. a Go module path).\"},\n        \"harnesses\": {\"type\": \"array\", \"minItems\": 1, \"items\": {\"$ref\": \"#/$defs/harness\"}}\n      }\n    },\n    \"harness\": {\n      \"type\": \"object\",\n      \"required\": [\"name\", \"engine\", \"path\", \"duration_seconds\", \"execs\", \"outcome\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"name\": {\"$ref\": \"#/$defs/name\", \"description\": \"Fuzz entry point (Go FuzzXxx, Jazzer class, cargo-fuzz target). Unique within the document.\"},\n        \"engine\": {\"$ref\": \"#/$defs/engine\"},\n        \"path\": {\"type\": \"string\", \"minLength\": 1, \"description\": \"Repository-relative path the harness was installed at.\"},\n        \"package\": {\"type\": \"string\", \"minLength\": 1, \"description\": \"Build package the engine ran (e.g. ./pkg/proxy).\"},\n        \"duration_seconds\": {\"type\": \"integer\", \"minimum\": 0, \"description\": \"Wall-clock fuzz time spent.\"},\n        \"execs\": {\"$ref\": \"#/$defs/count\", \"description\": \"Executions reported by the engine. null: the engine did not report.\"},\n        \"corpus_entries\": {\"$ref\": \"#/$defs/count\", \"description\": \"Inputs in the corpus at the end of the run. null: not reported.\"},\n        \"outcome\": {\"$ref\": \"#/$defs/outcome\"}\n      }\n    },\n    \"engine\": {\n      \"enum\": [\"go_native\", \"libfuzzer\", \"aflplusplus\", \"honggfuzz\", \"atheris\", \"jazzer\", \"jazzer_js\", \"cargo_fuzz\"],\n      \"description\": \"go_native: `go test -fuzz`. cargo_fuzz runs libFuzzer through cargo; jazzer_js is Jazzer.js.\"\n    },\n    \"outcome\": {\n      \"enum\": [\"clean\", \"crashed\", \"failed\", \"not_run\"],\n      \"description\": \"clean: ran its duration with no crasher. crashed: wrote at least one crashing input. failed: the harness or build failed without a crashing input. not_run: the engine found nothing to fuzz.\"\n    },\n    \"crasher\": {\n      \"type\": \"object\",\n      \"required\": [\"id\", \"harness\", \"title\", \"signature\", \"reproducer\", \"sanitizer\", \"cwes\", \"severity_proposal\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"id\": {\"$ref\": \"#/$defs/name\", \"description\": \"Unique within the document; the finding's source id at intake.\"},\n        \"harness\": {\"$ref\": \"#/$defs/name\", \"description\": \"The target harness `name` that produced it (the recording service checks the reference).\"},\n        \"title\": {\"type\": \"string\", \"minLength\": 1},\n        \"signature\": {\"$ref\": \"#/$defs/signature\"},\n        \"reproducer\": {\"$ref\": \"#/$defs/reproducer\"},\n        \"sanitizer\": {\n          \"oneOf\": [{\"$ref\": \"#/$defs/sanitizer\"}, {\"type\": \"null\"}],\n          \"description\": \"The sanitizer that reported the crash. null: none (memory-safe runtimes, property checks).\"\n        },\n        \"cwes\": {\n          \"type\": \"array\",\n          \"uniqueItems\": true,\n          \"items\": {\"type\": \"string\", \"pattern\": \"^CWE-\\\\d{1,5}$\"},\n          \"description\": \"Empty when not yet classified.\"\n        },\n        \"severity_proposal\": {\"$ref\": \"report.schema.json#/$defs/severity_level\", \"description\": \"The producer's proposal; triage owns the severity.\"},\n        \"cvss\": {\"$ref\": \"report.schema.json#/$defs/finding/properties/cvss\"},\n        \"location\": {\"$ref\": \"report.schema.json#/$defs/location\", \"description\": \"Where the crash surfaces (top in-repository frame), when known.\"},\n        \"finding_id\": {\"type\": \"string\", \"format\": \"uuid\", \"description\": \"An existing finding this crasher confirms. Absent: intake mints a new finding.\"},\n        \"analysis\": {\"$ref\": \"#/$defs/crasher_analysis\"}\n      }\n    },\n    \"signature\": {\n      \"type\": \"object\",\n      \"required\": [\"crash_class\", \"stack_hash\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"crash_class\": {\"$ref\": \"#/$defs/crash_class\"},\n        \"stack_hash\": {\n          \"type\": [\"string\", \"null\"],\n          \"pattern\": \"^[0-9a-f]{16,64}$\",\n          \"description\": \"Hash of the normalized top frames, for dedupe across runs. null: no stack (property violation, timeout without a trace).\"\n        },\n        \"frames\": {\"type\": \"integer\", \"minimum\": 1, \"description\": \"How many top frames the hash covers.\"}\n      }\n    },\n    \"crash_class\": {\n      \"enum\": [\"panic\", \"index_out_of_range\", \"nil_dereference\", \"out_of_memory\", \"timeout\", \"integer_overflow\", \"memory_corruption\", \"memory_leak\", \"uncaught_exception\", \"process_exit\", \"invariant_violation\"],\n      \"description\": \"invariant_violation: a property the harness asserts failed (open redirect, integer wrap, path escape) without a crash. process_exit: the target exited the process (os.Exit, abort).\"\n    },\n    \"sanitizer\": {\"enum\": [\"address\", \"memory\", \"undefined\", \"thread\", \"leak\"]},\n    \"reproducer\": {\n      \"type\": \"object\",\n      \"required\": [\"sha256\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"sha256\": {\"type\": \"string\", \"pattern\": \"^[0-9a-f]{64}$\", \"description\": \"Digest of the minimized crashing input's exact bytes.\"},\n        \"uri\": {\"type\": \"string\", \"minLength\": 1, \"description\": \"Where the bytes are kept (a location, never parsed for meaning).\"},\n        \"path\": {\"type\": \"string\", \"minLength\": 1, \"description\": \"Repository-relative path the engine wrote it to (e.g. testdata/fuzz/FuzzX/<hash>).\"}\n      }\n    },\n    \"coverage\": {\n      \"type\": \"object\",\n      \"required\": [\"tool\", \"lines_covered\", \"lines_total\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"tool\": {\"type\": \"string\", \"minLength\": 1},\n        \"lines_covered\": {\"type\": \"integer\", \"minimum\": 0},\n        \"lines_total\": {\"type\": \"integer\", \"minimum\": 0},\n        \"edges_covered\": {\"$ref\": \"#/$defs/count\"},\n        \"edges_total\": {\"$ref\": \"#/$defs/count\"}\n      }\n    },\n    \"gap\": {\n      \"type\": \"object\",\n      \"required\": [\"surface\", \"reason\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"surface\": {\"type\": \"string\", \"minLength\": 1, \"description\": \"Function, package or path left unfuzzed.\"},\n        \"reason\": {\"$ref\": \"#/$defs/gap_reason\"},\n        \"harness\": {\"$ref\": \"#/$defs/name\", \"description\": \"The harness concerned, for harness_defect and build_failed.\"},\n        \"note\": {\"type\": \"string\", \"minLength\": 1}\n      }\n    },\n    \"gap_reason\": {\n      \"enum\": [\"no_harness\", \"fuzzed_upstream\", \"unsupported_language\", \"needs_live_environment\", \"trusted_input\", \"heavy_dependencies\", \"harness_defect\", \"build_failed\"],\n      \"description\": \"no_harness: tracked without a fuzzer (static-confirmed or covered by another harness). fuzzed_upstream: already continuously fuzzed (e.g. OSS-Fuzz). trusted_input: input is not attacker-controlled. harness_defect: crashers came from the harness, not the target, so the surface is untested.\"\n    }\n  }\n}\n",
	"impact-analysis": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/impact-analysis.schema.json",
  "title": "CVE Impact Analysis Report Schema",
  "description": "Schema for <cve>-impact-analysis.json artifacts produced by the impact-analysis skill. Records per-repo affectedness classification for a CVE across the portfolio, with tiered evidence from the portfolio graph, govulncheck, opengrep, and ELF binary scans.",
  "type": "object",
  "required": [
    "metadata",
    "summary",
    "repos"
  ],
  "additionalProperties": false,
  "properties": {
    "metadata": {
      "$ref": "#/$defs/metadata"
    },
    "summary": {
      "$ref": "#/$defs/summary"
    },
    "repos": {
      "description": "Per-repo classification and evidence. Every repo in the blast radius appears exactly once.",
      "type": "array",
      "items": {
        "$ref": "#/$defs/repo_entry"
      }
    }
  },
  "$defs": {
    "metadata": {
      "type": "object",
      "required": [
        "cve",
        "module",
        "vulnerable_range",
        "harness_version",
        "generated_at",
        "tiers_executed",
        "options"
      ],
      "additionalProperties": false,
      "properties": {
        "cve": {
          "type": "string",
          "pattern": "^(?:CVE|GHSA|MAL|PYSEC|GO|RUSTSEC|OSV)-[A-Za-z0-9][A-Za-z0-9-]{0,63}$",
          "description": "Advisory identifier. CVE-YYYY-NNNN or any OSV advisory id (GHSA-/MAL-/PYSEC-/GO-/RUSTSEC-/OSV-). CVE ids remain valid (strict superset of the legacy CVE-only pattern)."
        },
        "module": {
          "type": "string",
          "minLength": 1,
          "description": "Affected module path (e.g. google.golang.org/grpc)."
        },
        "ecosystem": {
          "type": "string",
          "enum": [
            "go",
            "npm",
            "pypi",
            "maven",
            "cargo",
            "ruby",
            "nuget",
            "actions",
            "docker",
            "helm",
            "rpm",
            "deb",
            "apk"
          ],
          "description": "Package ecosystem of the affected module; the same vocabulary as report.schema.json finding.dependency.ecosystem (enums/v1/ecosystem.json). actions/docker/helm are manifest-level dependency surfaces (no reachability tier); rpm/deb/apk are distro ecosystems for C/C++ packages (no reachability tier). Optional (defaults to go semantics when absent) and non-breaking."
        },
        "preparation_scope": {
          "type": "object",
          "description": "The complete package/version snapshot prepared for this advisory and commit. Readers promote a replacement snapshot only after every declared member has recorded evidence; an incomplete run cannot clear prior exposure. digest: SHA-256 over the UTF-8 JSON encoding of the snapshot object (prepared_at, advisory, manifest_sbom_id, images, members) with sort_keys=true and default separators, as traust-core ExposureService.prepare computes it (sha256_hex(json.dumps(snapshot, sort_keys=True))); members serialize as objects with exactly the keys ecosystem, module, version, the member list sorted by (ecosystem, module, version). A producer diverging from this canonical form is a producer bug, not a schema interpretation.",
          "required": ["digest", "members"],
          "additionalProperties": false,
          "properties": {
            "digest": {"type": "string", "pattern": "^[0-9a-f]{64}$"},
            "members": {
              "type": "array",
              "minItems": 1,
              "uniqueItems": true,
              "items": {
                "type": "object",
                "required": ["ecosystem", "module", "version"],
                "additionalProperties": false,
                "properties": {
                  "ecosystem": {"$ref": "#/$defs/metadata/properties/ecosystem"},
                  "module": {"type": "string", "minLength": 1},
                  "version": {"type": ["string", "null"], "minLength": 1}
                }
              }
            }
          }
        },
        "vulnerable_range": {
          "type": "string",
          "description": "Version constraint describing the vulnerable range (e.g. '< v1.64.1')."
        },
        "fixed_version": {
          "type": [
            "string",
            "null"
          ],
          "description": "First safe version, if known."
        },
        "vulnerable_symbols": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Known vulnerable symbols from the advisory or fix diff."
        },
        "vulnerable_packages": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Packages containing the vulnerable code."
        },
        "feature_description": {
          "type": [
            "string",
            "null"
          ],
          "description": "Human-readable description of the feature/capability that must be in use."
        },
        "advisory_sources": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "URLs of advisory sources consulted."
        },
        "portfolio_graph_db": {
          "type": [
            "string",
            "null"
          ],
          "description": "Path to the portfolio-graph.db used."
        },
        "portfolio_graph_version": {
          "type": [
            "string",
            "null"
          ],
          "description": "Date or version of the portfolio graph data."
        },
        "harness_version": {
          "type": "string",
          "description": "Harness version (semver-shortsha)."
        },
        "generated_at": {
          "type": "string",
          "format": "date-time",
          "description": "ISO 8601 timestamp of report generation."
        },
        "tiers_executed": {
          "type": "array",
          "items": {
            "type": "string",
            "enum": [
              "L1",
              "L4",
              "govulncheck",
              "feature_pattern",
              "binary_scan",
              "manifest_scan",
              "source_scan",
              "sbom_cross_check"
            ]
          },
          "description": "Which analysis tiers were executed."
        },
        "options": {
          "type": "object",
          "properties": {
            "govulncheck": {
              "type": "boolean"
            },
            "binary_scan": {
              "type": "boolean"
            },
            "sweep": {
              "type": "boolean"
            }
          },
          "description": "CLI options used for this run."
        }
      }
    },
    "summary": {
      "type": "object",
      "required": [
        "repos_in_blast_radius",
        "version_in_range",
        "affected",
        "likely_affected",
        "not_observed",
        "version_not_in_range",
        "not_imported",
        "inconclusive"
      ],
      "additionalProperties": false,
      "properties": {
        "repos_in_blast_radius": {
          "type": "integer",
          "minimum": 0
        },
        "version_in_range": {
          "type": "integer",
          "minimum": 0
        },
        "affected": {
          "type": "integer",
          "minimum": 0
        },
        "likely_affected": {
          "type": "integer",
          "minimum": 0
        },
        "not_observed": {
          "type": "integer",
          "minimum": 0
        },
        "version_not_in_range": {
          "type": "integer",
          "minimum": 0
        },
        "not_imported": {
          "type": "integer",
          "minimum": 0
        },
        "inconclusive": {
          "type": "integer",
          "minimum": 0
        },
        "product_surfaces": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Product surfaces that include repos in the blast radius."
        }
      }
    },
    "repo_entry": {
      "type": "object",
      "required": [
        "repo",
        "classification",
        "evidence"
      ],
      "additionalProperties": false,
      "properties": {
        "repo": {
          "type": "string",
          "description": "Repository identifier (e.g. 'repo:github.com/openshift/foo')."
        },
        "products": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Products that ship this repo."
        },
        "classification": {
          "type": "string",
          "enum": [
            "affected",
            "likely_affected",
            "not_observed",
            "version_not_in_range",
            "not_imported",
            "inconclusive"
          ]
        },
        "version": {
          "type": [
            "string",
            "null"
          ],
          "description": "Pinned version of the affected module in this repo."
        },
        "direct": {
          "type": [
            "boolean",
            "null"
          ],
          "description": "Whether the dependency is direct (not indirect/transitive)."
        },
        "evidence": {
          "$ref": "#/$defs/evidence"
        }
      }
    },
    "evidence": {
      "type": "object",
      "additionalProperties": false,
      "properties": {
        "l1_depends_on": {
          "type": [
            "boolean",
            "null"
          ],
          "description": "Module appears in the repo's dependency graph."
        },
        "l1_version_in_range": {
          "type": [
            "boolean",
            "null"
          ],
          "description": "Pinned version falls in the vulnerable range."
        },
        "l4_package_imported": {
          "type": [
            "boolean",
            "null"
          ],
          "description": "Vulnerable package (not just module) is imported per L4 symbol data."
        },
        "l4_packages_found": {
          "type": [
            "array",
            "null"
          ],
          "items": {
            "type": "string"
          },
          "description": "Specific vulnerable packages found imported."
        },
        "govulncheck": {
          "type": [
            "string",
            "null"
          ],
          "enum": [
            "symbol_reachable",
            "package_imported_not_observed",
            "module_required_not_observed",
            null
          ],
          "description": "Govulncheck reachability classification."
        },
        "govulncheck_trace": {
          "type": [
            "array",
            "null"
          ],
          "items": {
            "type": "string"
          },
          "description": "Caller-first call path from govulncheck (when symbol_reachable)."
        },
        "feature_pattern_matches": {
          "type": [
            "integer",
            "null"
          ],
          "description": "Number of feature-pattern rule matches."
        },
        "binary_string_scan": {
          "type": [
            "string",
            "null"
          ],
          "enum": [
            "package_path_present",
            "package_path_absent",
            "not_scanned",
            null
          ],
          "description": "ELF binary string scan result."
        },
        "needs_manual_trace": {
          "type": [
            "boolean",
            "null"
          ],
          "description": "True when unsafe/reflect/cgo usage detected that may hide reachability from static analysis."
        },
        "notes": {
          "type": [
            "string",
            "null"
          ],
          "description": "Free-text notes about this repo's analysis."
        },
        "manifest_scan": {
          "type": [
            "string",
            "null"
          ],
          "enum": [
            "module_pinned_in_range",
            "module_pinned_out_of_range",
            "module_not_in_manifests",
            "no_manifests_found",
            null
          ],
          "description": "Lockfile/manifest scan result (Python/Rust/JS/Java \u2014 manifest-level evidence, never proves reachability)."
        },
        "manifest_version": {
          "type": [
            "string",
            "null"
          ],
          "description": "Version pinned in the manifest when the module was found."
        },
        "source_import_scan": {
          "type": [
            "string",
            "null"
          ],
          "enum": [
            "imports_found",
            "imports_not_found",
            null
          ],
          "description": "Source-level import/include grep for the module."
        },
        "binary_linked_library": {
          "type": [
            "string",
            "null"
          ],
          "enum": [
            "linked",
            "not_linked",
            null
          ],
          "description": "DT_NEEDED linked-library scan of shipped image binaries (C/C++)."
        },
        "evidence_level": {
          "type": [
            "string",
            "null"
          ],
          "enum": [
            "symbol",
            "symbol-usage",
            "binary",
            "manifest",
            "none",
            null
          ],
          "description": "Strongest evidence tier that produced this classification: symbol (govulncheck reachability) > binary (ELF) > manifest (lockfile/source grep)."
        },
        "symbol_usage_scan": {
          "type": [
            "string",
            "null"
          ],
          "enum": [
            "symbols_used",
            "symbols_not_found",
            null
          ],
          "description": "Source grep for the advisory's vulnerable symbol names \u2014 textual API usage, stronger than a pin, weaker than reachability."
        },
        "binary_symbol_scan": {
          "type": [
            "string",
            "null"
          ],
          "enum": [
            "symbols_present",
            "symbols_absent",
            null
          ],
          "description": "Vulnerable symbol names present in the shipped binary's strings (.dynstr) \u2014 the binary references the vulnerable function, not just the library."
        },
        "sbom_scan": {
          "type": [
            "string",
            "null"
          ],
          "enum": [
            "shipped_in_range",
            "shipped_out_of_range",
            "module_not_in_sbom",
            null
          ],
          "description": "Cross-check against syft SBOMs of shipped images: what the delivered artifact actually contains."
        },
        "sbom_shipped_version": {
          "type": [
            "string",
            "null"
          ],
          "description": "Module version found in the shipped image SBOM."
        }
      }
    }
  }
}
`,
	"isolation-review": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/isolation-review.schema.json",
  "title": "Tenant-isolation review report",
  "description": "Service-level tenant-isolation review (PEACH applied by reference only \u2014 framework cited by name/URL, all check wording original to this repository). Enforcement gate: schema validation script (mirror of this schema). Every partial/no dimension result MUST cite evidence (file:line or finding IDs). Reports land under your isolation review output directory per service.",
  "type": "object",
  "required": [
    "title",
    "metadata",
    "interfaces",
    "gaps",
    "posture"
  ],
  "properties": {
    "title": {
      "type": "string"
    },
    "metadata": {
      "type": "object",
      "required": [
        "service",
        "repos",
        "graph_ref",
        "harness_version",
        "reviewed_at"
      ],
      "properties": {
        "service": {
          "type": "string"
        },
        "repos": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "minItems": 1
        },
        "graph_ref": {
          "type": "string",
          "$comment": "portfolio-graph.db path + ref/mtime the repo set was resolved from, or 'manual' with the reason in notes"
        },
        "harness_version": {
          "type": "string"
        },
        "reviewed_at": {
          "type": "string"
        },
        "framework": {
          "type": "string",
          "$comment": "citation-only, e.g. 'PEACH v1.1 (by reference)' \u2014 never adapted text"
        }
      }
    },
    "interfaces": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/interface"
      }
    },
    "gaps": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/gap"
      },
      "$comment": "ordered worst-first (prioritized)"
    },
    "posture": {
      "type": "object",
      "required": [
        "overall",
        "summary"
      ],
      "properties": {
        "overall": {
          "enum": [
            "strong",
            "adequate",
            "weak",
            "critical-gap"
          ]
        },
        "summary": {
          "type": "string"
        },
        "interfaces_reviewed": {
          "type": "integer",
          "minimum": 0
        },
        "dimension_rollup": {
          "type": "object",
          "additionalProperties": {
            "type": "integer"
          },
          "$comment": "optional counts, e.g. {\"yes\": 12, \"partial\": 4, \"no\": 2, \"na\": 2}"
        }
      }
    },
    "notes": {
      "type": "string"
    }
  },
  "definitions": {
    "interface": {
      "type": "object",
      "required": [
        "id",
        "name",
        "kind",
        "exposure",
        "complexity",
        "dimensions"
      ],
      "properties": {
        "id": {
          "type": "string",
          "pattern": "^IF-[0-9]+$"
        },
        "name": {
          "type": "string"
        },
        "kind": {
          "enum": [
            "api",
            "data-store",
            "queue",
            "ingress",
            "webhook",
            "cli",
            "other"
          ]
        },
        "exposure": {
          "enum": [
            "public",
            "tenant",
            "partner",
            "internal"
          ]
        },
        "complexity": {
          "enum": [
            "low",
            "medium",
            "high"
          ]
        },
        "shared_instance": {
          "type": [
            "boolean",
            "null"
          ],
          "$comment": "true = every tenant hits the same running instance; false = per-tenant copies; null = undetermined"
        },
        "repos": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "dimensions": {
          "type": "object",
          "required": [
            "privilege",
            "encryption",
            "authentication",
            "connectivity",
            "hygiene"
          ],
          "properties": {
            "privilege": {
              "$ref": "#/definitions/dimension"
            },
            "encryption": {
              "$ref": "#/definitions/dimension"
            },
            "authentication": {
              "$ref": "#/definitions/dimension"
            },
            "connectivity": {
              "$ref": "#/definitions/dimension"
            },
            "hygiene": {
              "$ref": "#/definitions/dimension"
            }
          }
        }
      }
    },
    "dimension": {
      "type": "object",
      "required": [
        "result"
      ],
      "properties": {
        "result": {
          "enum": [
            "yes",
            "partial",
            "no",
            "na"
          ]
        },
        "evidence": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "$comment": "file:line references (repo//path:line) or finding/threat IDs from existing artifacts"
        },
        "rationale": {
          "type": "string"
        }
      },
      "if": {
        "properties": {
          "result": {
            "enum": [
              "partial",
              "no"
            ]
          }
        }
      },
      "then": {
        "required": [
          "result",
          "evidence"
        ],
        "properties": {
          "evidence": {
            "type": "array",
            "items": {
              "type": "string"
            },
            "minItems": 1
          }
        }
      }
    },
    "gap": {
      "type": "object",
      "required": [
        "severity",
        "interface_id",
        "description",
        "remediation"
      ],
      "properties": {
        "severity": {
          "enum": [
            "critical",
            "high",
            "medium",
            "low",
            "informational"
          ]
        },
        "interface_id": {
          "type": "string",
          "pattern": "^IF-[0-9]+$"
        },
        "description": {
          "type": "string"
        },
        "remediation": {
          "type": "string"
        }
      }
    }
  }
}
`,
	"layer": "{\n  \"$schema\": \"https://json-schema.org/draft/2020-12/schema\",\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/layer.schema.json\",\n  \"title\": \"Findings Disposition Layer Schema\",\n  \"description\": \"Schema for *-findings-layer.json ledgers produced by the track-findings harness. An append-only record of human triage and machine validation events against the findings of one security-audit report, plus a needs-review queue for disposition statements that require human confirmation before they may change state. The layer is replayed by the cumulative builder to produce a *-findings-current.{json,md} cumulative report.\",\n  \"type\": \"object\",\n  \"required\": [\n    \"metadata\",\n    \"events\",\n    \"needs_review\"\n  ],\n  \"additionalProperties\": false,\n  \"properties\": {\n    \"metadata\": {\n      \"$ref\": \"#/$defs/layer_metadata\"\n    },\n    \"events\": {\n      \"description\": \"Append-only disposition events, in chronological order. Events are never edited or deleted \\u2014 a restatement is a new event.\",\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/event\"\n      }\n    },\n    \"needs_review\": {\n      \"description\": \"Disposition-like statements that did not meet the tier-1 auto-record bar (explicit grammar + identity-provider-verified author + authority). No state change until a human confirms.\",\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/review_item\"\n      }\n    },\n    \"baseline_claims\": {\n      \"type\": \"object\",\n      \"additionalProperties\": {\n        \"type\": \"string\",\n        \"pattern\": \"^[a-f0-9]{64}$\"\n      },\n      \"description\": \"Finding id -> sha256 claim hash, pinned at baseline time. Produced before this schema declared it: early layers carried it as an undeclared root property and failed validation under additionalProperties:false.\",\n      \"$comment\": \"Overlaps metadata.claim_hashes in shape and meaning \\u2014 both are finding_id -> claim digest. Declared here to make the corpus valid, not to bless the duplication; consolidating the two is a separate decision.\"\n    }\n  },\n  \"$defs\": {\n    \"validity\": {\n      \"type\": \"string\",\n      \"enum\": [\n        \"confirmed\",\n        \"false_positive\",\n        \"corrected\",\n        \"hardening\",\n        \"not_affected\"\n      ],\n      \"$comment\": \"Canonical validity is confirmed (real claim), false_positive (wrong claim), or not_affected (real upstream issue inapplicable here). Historical corrected and hardening remain readable; the owner imports them as restatement and confirmed weakness respectively. not_verified is an audit-stage absence of an answer, not an event validity.\"\n    },\n    \"resolution\": {\n      \"type\": \"string\",\n      \"enum\": [\n        \"open\",\n        \"fix_in_progress\",\n        \"resolved\",\n        \"partially_resolved\",\n        \"risk_accepted\",\n        \"regression_introduced\"\n      ],\n      \"$comment\": \"The resolution axis: is the finding dealt with?\"\n    },\n    \"source_type\": {\n      \"type\": \"string\",\n      \"enum\": [\n        \"mr_comment\",\n        \"commit\",\n        \"jira\",\n        \"interactive\",\n        \"validation_report\",\n        \"verification_report\",\n        \"triage_report\",\n        \"impact_report\",\n        \"vuln_scan_report\",\n        \"fuzz_report\",\n        \"rebaseline\",\n        \"restatement\"\n      ],\n      \"$comment\": \"restatement is NOT an evidence class either: it records that DATA was wrong -- a claim hash, a report digest, an artifact digest map, or a field on a prior event -- and carries an empty disposition plus a `restatement` block. It exists because everything else here is immutable: event payloads are append-only and the Merkle leaf binds their full content, so a wrong value can only be superseded, never edited. The signature-bound metadata digests (claim_hashes, audit_report_sha256, artifact_digests) ARE mutable, and before this vocabulary existed they could be rewritten with no record of the prior value, who changed it, or why -- making an authorised restatement byte-identical to tampering after the fact. Evidence classes for validity precedence (the cumulative builder): validation_report, verification_report and fuzz_report are execution-verified (class 1) and outrank human static determinations (class 2), which outrank machine static sources incl. triage_report, impact_report and vuln_scan_report (class 3). Recency breaks ties only within a class. fuzz_report is class 1 because a fuzz finding is a crasher WITH a reproducer: the input that triggers it is recorded, so the claim is replayable rather than asserted. A fuzz run that found nothing records no event -- absence of a crash is not evidence of correctness. rebaseline is NOT an evidence class and never participates in validity precedence: it records that a finding was RENAMED, so the alias lives in the Merkle-covered event stream rather than in mutable metadata. Every one of them carries an EMPTY disposition, which is the invariant a consumer should rely on -- a rebaseline event asserts nothing about whether a finding is real.\"\n    },\n    \"actor\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"kind\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"kind\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"human\",\n            \"machine\"\n          ]\n        },\n        \"identity\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Canonical email for humans; skill/report name for machines.\"\n        },\n        \"ldap_verified\": {\n          \"type\": \"boolean\",\n          \"description\": \"Legacy. True when the human identity was verified via LDAP. Superseded by identity_verified + employee_status; retained for backward compatibility with existing events.\"\n        },\n        \"identity_verified\": {\n          \"type\": \"boolean\",\n          \"description\": \"True when an authentication mechanism confirmed this identity (OIDC, LDAP bind, keypair, etc.).\"\n        },\n        \"identity_provider\": {\n          \"type\": \"string\",\n          \"description\": \"Which IdP type authenticated this actor: oidc, ldap, bearer, keypair, mtls, etc.\"\n        },\n        \"identity_issuer\": {\n          \"type\": \"string\",\n          \"description\": \"Credential issuer (OIDC issuer URL, LDAP server URI, key fingerprint, etc.).\"\n        },\n        \"identity_subject\": {\n          \"type\": \"string\",\n          \"description\": \"Raw IdP subject before canonicalization (OIDC sub UUID, LDAP uid, etc.) for audit reversibility.\"\n        },\n        \"employee_status\": {\n          \"type\": \"string\",\n          \"description\": \"Directory enrichment result: active, terminated, contingent, or not_found.\"\n        },\n        \"display_name\": {\n          \"type\": \"string\"\n        }\n      }\n    },\n    \"disposition\": {\n      \"type\": \"object\",\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"validity\": {\n          \"$ref\": \"#/$defs/validity\"\n        },\n        \"resolution\": {\n          \"$ref\": \"#/$defs/resolution\"\n        },\n        \"severity\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"critical\",\n            \"high\",\n            \"medium\",\n            \"low\",\n            \"informational\"\n          ],\n          \"description\": \"Human severity override (harness >= 0.128.0): the signer's re-assessment of this finding's severity, recorded via the countersign workbench with a mandatory own-words rationale. HUMAN-ONLY - the validator rejects machine actors carrying disposition.severity. The audit report's original severity is never rewritten; the cumulative builder surfaces the latest identity-provider-verified human severity event as the finding's effective_severity, keeping original by-severity counts stable for downstream dashboards. Idempotency: the level is encoded in source.ref (interactive:<date>:severity:<level>), so the canonical event_id formula is unchanged.\"\n        },\n        \"embargo\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"required\",\n            \"active\",\n            \"not_required\",\n            \"uncertain\"\n          ],\n          \"description\": \"Embargo handling assertion (harness >= 0.274.0): whether this finding must be handled under security team embargo criteria (remote code execution, authentication bypass, privilege escalation, sensitive data exposure, network-accessible vulnerability). 'required' = meets the criteria, not yet embargoed; 'active' = an embargoed tracker already exists (put its URL in evidence_refs); 'not_required' = assessed and does not meet the criteria; 'uncertain' = needs security team consultation. HUMAN-ONLY - the validator rejects machine actors carrying disposition.embargo, on the same identity-provider-verified footing as disposition.severity. ORTHOGONAL to validity and resolution: an embargoed finding is still tracked on both axes, and embargo alone never changes finding state. Consumed by embargo workflows that apply restricted tracker security levels when active. Idempotency: the status is encoded in source.ref (interactive:<date>:embargo:<status>), so the canonical event_id formula is unchanged.\"\n        }\n      },\n      \"description\": \"What this event asserts about the finding. Required to state something for every source EXCEPT rebaseline, which records a rename; that rule lives on the event, next to the source it depends on.\"\n    },\n    \"layer_metadata\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"audit_report\",\n        \"repository\",\n        \"created\",\n        \"harness_version\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"audit_report\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Path to the baseline report this layer annotates (relative to the layer file): a *-security-audit.json, *-cloud-config-audit.json, or *-container-audit.json.\"\n        },\n        \"audit_report_sha256\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{64}$\",\n          \"description\": \"sha256 of the annotated baseline report's exact bytes. Makes the layer's link to its report CONTENT-addressed rather than path-addressed, which is the prerequisite for reports living anywhere but beside the layer (e.g. object storage while the ledger stays in git). Complements metadata.claim_hashes, which pins the semantic claims of each finding: the digest answers 'are these the bytes I annotated?', the claim hashes answer 'did any finding's claim change?'. Recomputed whenever the baseline is legitimately rewritten (a re-audit, or a corpus migration such as the v1->v2 fingerprint re-stamp); a mismatch is surfaced, never auto-healed, on the same footing as a claim-hash mismatch. NOT covered by merkle_signature_payload today \\u2014 see the format-3 note before relying on it as tamper-evidence.\",\n          \"$comment\": \"Additive and optional so existing layers stay valid; backfilled by hashing the report in place.\"\n        },\n        \"artifact_digests\": {\n          \"type\": \"object\",\n          \"additionalProperties\": {\n            \"type\": \"string\",\n            \"pattern\": \"^[0-9a-f]{64}$\"\n          },\n          \"description\": \"Filename -> sha256 for every non-layer artifact belonging to this layer's base (triage, threat model, findings-current, verification, privilege profile, ...). Its digest is inside the signature under format 4. Without it only the ONE report named by audit_report_sha256 is verifiable, i.e. 15% of what a migration to object storage copies; the other 85% could not be checked against anything once the git copy was deleted. Keys are filenames relative to the layer's own directory, never paths.\",\n          \"$comment\": \"One field per layer (~5 entries, median) rather than many new ledger fields or manifest files in a repo the migration exists to shrink.\"\n        },\n        \"audit_report_ref\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Opaque location of the annotated baseline report \\u2014 an object-storage key/URI once reports move out of git, absent while they sit beside the layer. Deliberately opaque: consumers resolve it through the report_store accessor, never by string-munging a path, so the same layer works against a local checkout, a bucket, or a future ledger database without edits.\",\n          \"$comment\": \"Pairs with audit_report_sha256: the ref says where, the digest says which bytes.\"\n        },\n        \"audit_commit\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{7,64}$\",\n          \"description\": \"metadata.commit of the annotated audit report: a git SHA (7-40 hex) or, for container-audit baselines, the image manifest digest hex (64 chars).\"\n        },\n        \"repository\": {\n          \"type\": \"string\",\n          \"format\": \"uri\"\n        },\n        \"created\": {\n          \"type\": \"string\",\n          \"format\": \"date-time\"\n        },\n        \"updated\": {\n          \"type\": \"string\",\n          \"format\": \"date-time\",\n          \"description\": \"Timestamp of the most recent append.\"\n        },\n        \"harness_version\": {\n          \"type\": \"string\",\n          \"pattern\": \"^\\\\d+\\\\.\\\\d+\\\\.\\\\d+(-[0-9a-f]{7,40})?$\",\n          \"description\": \"Harness version that created the layer. Individual events carry their own harness_version when appended by a later version.\"\n        },\n        \"claim_hashes\": {\n          \"type\": \"object\",\n          \"description\": \"Baseline tamper-evidence (harness >= 0.39.0): finding id -> sha256 of the finding's canonical claim fields (id, title, severity, cwes, locations, description, remediation \\u2014 validation_status excluded). Recorded by the baseline claims tool when a finding is first baselined; never overwritten except by an explicit --rebaseline of a 'corrected' finding. Verified by the validator on every layer validation and by the cumulative builder before every rebuild.\",\n          \"additionalProperties\": {\n            \"type\": \"string\",\n            \"pattern\": \"^[0-9a-f]{64}$\"\n          }\n        },\n        \"finding_aliases\": {\n          \"type\": \"object\",\n          \"description\": \"Rebaseline mapping: old scan-scoped finding id -> its successor in a newer audit. Events are never rewritten; build_cumulative resolves refs through this table at replay. matched_by 'fingerprint' mappings are auto-confirmed; others pend in needs_review.\",\n          \"additionalProperties\": {\n            \"type\": \"object\",\n            \"required\": [\n              \"new_id\",\n              \"matched_by\",\n              \"mapped_at\"\n            ],\n            \"properties\": {\n              \"new_id\": {\n                \"type\": \"string\"\n              },\n              \"matched_by\": {\n                \"enum\": [\n                  \"fingerprint\",\n                  \"path_cwe\",\n                  \"path_set\",\n                  \"title\",\n                  \"manual\"\n                ]\n              },\n              \"mapped_at\": {\n                \"type\": \"string\"\n              },\n              \"from_report\": {\n                \"type\": \"string\"\n              },\n              \"confirmed\": {\n                \"type\": \"boolean\"\n              },\n              \"similarity\": {\n                \"type\": \"number\"\n              },\n              \"note\": {\n                \"type\": \"string\"\n              },\n              \"path_overlap\": {\n                \"type\": \"number\"\n              },\n              \"confirmed_by\": {\n                \"type\": \"string\"\n              },\n              \"confirmed_at\": {\n                \"type\": \"string\"\n              },\n              \"rejected\": {\n                \"type\": \"boolean\"\n              },\n              \"rejected_by\": {\n                \"type\": \"string\"\n              },\n              \"rejected_at\": {\n                \"type\": \"string\"\n              }\n            },\n            \"additionalProperties\": false\n          }\n        },\n        \"external_refs\": {\n          \"type\": \"object\",\n          \"description\": \"Provenance map: finding id -> identifiers the finding was escalated to or reconciled against in an external system (CVE, Bugzilla, GHSA, Jira). Written by the CVE-provenance reconciler; DERIVED and re-computable, so it lives in metadata beside finding_aliases rather than in the Merkle-signed event chain, and it asserts nothing about validity, resolution or severity. Answers 'which CVE did our finding become' from the ledger instead of from a person's memory.\",\n          \"additionalProperties\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"object\",\n              \"required\": [\n                \"system\",\n                \"id\"\n              ],\n              \"additionalProperties\": false,\n              \"properties\": {\n                \"system\": {\n                  \"type\": \"string\",\n                  \"enum\": [\n                    \"cve\",\n                    \"bugzilla\",\n                    \"ghsa\",\n                    \"jira\"\n                  ]\n                },\n                \"id\": {\n                  \"type\": \"string\",\n                  \"minLength\": 1\n                },\n                \"url\": {\n                  \"type\": \"string\",\n                  \"pattern\": \"^https://\"\n                },\n                \"confidence\": {\n                  \"type\": \"string\",\n                  \"enum\": [\n                    \"confirmed\",\n                    \"probable\"\n                  ],\n                  \"description\": \"'confirmed' = title match at or above the auto-stamp threshold, or corroborated by an overlapping CWE. 'probable' = weaker signal (the auditing organization is the CNA, no external party credited, publication postdates the audit). Only 'confirmed' belongs in a headline metric.\"\n                },\n                \"matched_on\": {\n                  \"type\": \"string\",\n                  \"description\": \"How the link was derived, e.g. 'title-similarity:0.82' or 'cna-uncredited'. Makes a stamp auditable and lets a re-run supersede a weaker match.\"\n                },\n                \"stamped_at\": {\n                  \"type\": \"string\"\n                }\n              }\n            }\n          }\n        },\n        \"merkle_root\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{64}$\",\n          \"description\": \"SHA-256 Merkle tree root over post-epoch events (RFC 9162 \\u00a72.1 MTH).\"\n        },\n        \"merkle_epoch\": {\n          \"type\": \"integer\",\n          \"minimum\": 0,\n          \"description\": \"Event index where Merkle tracking begins; events before this index are covered by pre_merkle_checkpoint.\"\n        },\n        \"merkle_size\": {\n          \"type\": \"integer\",\n          \"minimum\": 0,\n          \"description\": \"Number of leaves in the Merkle tree (events from merkle_epoch onward).\"\n        },\n        \"merkle_algorithm\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"sha256\"\n          ],\n          \"description\": \"Hash algorithm for the Merkle tree.\"\n        },\n        \"leaf_format\": {\n          \"type\": \"integer\",\n          \"enum\": [\n            1,\n            2\n          ],\n          \"description\": \"Merkle leaf format: 1 = legacy event_id-only binding (verifies with a content-unbound warning); 2 = full canonical event content (current \\u2014 any edit to any event field breaks the root).\"\n        },\n        \"pre_merkle_checkpoint\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{64}$\",\n          \"description\": \"SHA-256 of the JSON-serialized pre-epoch events array, pinning pre-Merkle history.\"\n        },\n        \"merkle_root_signature\": {\n          \"type\": \"string\",\n          \"description\": \"Optional signature over merkle_root (base64 for keypair, JSON bundle for identity).\"\n        },\n        \"merkle_signature_format\": {\n          \"type\": \"integer\",\n          \"enum\": [\n            1,\n            2,\n            3,\n            4\n          ],\n          \"description\": \"Signed-payload format: 1 = legacy bare-root signature (verifies with a content-unbound warning); 2 = digest over root + leaf_format + epoch + size + pre-epoch checkpoint + claim-hashes digest; 3 = format 2 plus audit_report_sha256, binding the signature to the exact report bytes the layer annotates; 4 = format 3 plus a digest over metadata.artifact_digests, extending tamper-evidence from the one annotated report to every sibling artifact (plan R1, the precondition for deleting the git copies).\"\n        },\n        \"merkle_signing_method\": {\n          \"type\": \"string\",\n          \"description\": \"Signing method used to produce merkle_root_signature (e.g. 'keypair', 'identity'). Used to dispatch to the correct verification backend.\"\n        }\n      }\n    },\n    \"event\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"event_id\",\n        \"finding_ref\",\n        \"recorded_at\",\n        \"source\",\n        \"disposition\",\n        \"rationale\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"event_id\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[a-f0-9]{64}$\",\n          \"description\": \"sha256 of '<source.ref>|<finding_ref>|<disposition.validity or \\\"\\\">|<disposition.resolution or \\\"\\\">'. Deterministic, so re-ingesting the same source is idempotent (validator-enforced).\"\n        },\n        \"finding_ref\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Finding ID in the annotated audit report.\"\n        },\n        \"recorded_at\": {\n          \"type\": \"string\",\n          \"format\": \"date-time\",\n          \"description\": \"When the event was appended to the ledger. Monotonic across the events array (validator-enforced).\"\n        },\n        \"occurred_at\": {\n          \"type\": \"string\",\n          \"format\": \"date-time\",\n          \"description\": \"When the underlying determination actually happened \\u2014 commit author date, Jira transition date, MR comment date, machine report date. Trend tooling buckets by this when present (falling back to recorded_at), so late ingestion does not distort time series. Should not be later than recorded_at.\"\n        },\n        \"source\": {\n          \"type\": \"object\",\n          \"required\": [\n            \"type\",\n            \"ref\",\n            \"actor\"\n          ],\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"type\": {\n              \"$ref\": \"#/$defs/source_type\"\n            },\n            \"ref\": {\n              \"type\": \"string\",\n              \"minLength\": 1,\n              \"description\": \"Permalink to the MR comment, commit URL/SHA, Jira key, report path, or 'interactive:<session-date>'.\"\n            },\n            \"actor\": {\n              \"$ref\": \"#/$defs/actor\",\n              \"description\": \"Who this determination is attributed to. For confirmed needs_review items this is the confirming human, not the original commenter.\"\n            },\n            \"reported_by\": {\n              \"type\": \"string\",\n              \"description\": \"Original commenter/reporter when the event was confirmed from the needs_review queue by a different human.\"\n            }\n          }\n        },\n        \"disposition\": {\n          \"$ref\": \"#/$defs/disposition\"\n        },\n        \"rationale\": {\n          \"type\": \"string\",\n          \"minLength\": 10,\n          \"description\": \"Verbatim quote of the human statement, the Jira resolution comment, or the machine verdict explanation.\"\n        },\n        \"evidence_refs\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\",\n            \"minLength\": 1\n          }\n        },\n        \"harness_version\": {\n          \"type\": \"string\",\n          \"pattern\": \"^\\\\d+\\\\.\\\\d+\\\\.\\\\d+(-[0-9a-f]{7,40})?$\"\n        },\n        \"fingerprint\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{64}$\",\n          \"description\": \"The annotated finding's identity as observed when this event was recorded, copied from the baseline report's finding.fingerprint. Stamped by the sanctioned write path (the events module attach_identity), never model-authored. Makes the event self-contained: finding_ref is scan-scoped and changes on re-audit, so replay previously depended on metadata.finding_aliases \\u2014 a mutable table outside the Merkle tree. Under leaf_format 2 the stamp is tamper-evident for free. A historical observation, never corrected in place: if the identity later changes (moved file, recategorised CWE, algo bump) that is a NEW event.\"\n        },\n        \"fingerprint_algo\": {\n          \"type\": \"string\",\n          \"pattern\": \"^v\\\\d+$\",\n          \"description\": \"Algorithm version of `fingerprint` (the identity module ALGO_VERSION, re-exported as the events module FINGERPRINT_ALGO_CURRENT); 'v2' since 2026-08-18, when empty-canonicalizing location paths were dropped from the hashed set. Recorded per event because stamps outlive recipes: both versions coexist by design (a stamp is a historical observation, never re-computed in place), so matching fingerprints across an epoch boundary means comparing this field too.\"\n        },\n        \"risk_weight\": {\n          \"type\": \"object\",\n          \"description\": \"Emission-time-resolved hardening risk weight (triage_report hardening events). Recorded so ledger replays never re-derive weights and stay deterministic under future weight-table changes.\",\n          \"required\": [\n            \"lambda\",\n            \"weights_version\",\n            \"tenancy_profile\",\n            \"profile_source\"\n          ],\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"lambda\": {\n              \"type\": \"number\",\n              \"minimum\": 0,\n              \"maximum\": 1\n            },\n            \"weights_version\": {\n              \"type\": \"string\",\n              \"description\": \"version field of config/hardening-risk-weights.json at emission time.\"\n            },\n            \"tenancy_profile\": {\n              \"type\": \"string\",\n              \"enum\": [\n                \"multi_tenant\",\n                \"single_tenant\"\n              ]\n            },\n            \"profile_source\": {\n              \"type\": \"string\",\n              \"description\": \"Artifact the profile was derived from, e.g. 'peach_isolation_review.applicable', 'threat_model', or 'override'.\"\n            }\n          }\n        },\n        \"auto_accept_tier\": {\n          \"type\": \"boolean\",\n          \"description\": \"True on a machine false_positive event that met the auto-accept bar (unanimous, verdict-citation-lint clean, mean confidence >= 8, claimed severity low/informational). The merge engine lets such events set validity without a human countersign; a deterministic 10% sample is still queued for human audit.\"\n        },\n        \"evidence_grade\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"enum\": [\n            \"E0\",\n            \"E1\",\n            \"E2\",\n            \"E3\",\n            null\n          ],\n          \"description\": \"Explicit evidence strength: E0 observed effect, E1 authenticated success response, E2 strong inference, E3 behavioral inference. E0 is strongest. Missing or null means ungraded; source type and human approval never invent or upgrade a grade.\"\n        },\n        \"alias\": {\n          \"$ref\": \"#/$defs/event_alias\"\n        },\n        \"finding\": {\n          \"$ref\": \"#/$defs/event_finding\"\n        },\n        \"restatement\": {\n          \"$ref\": \"#/$defs/event_restatement\"\n        }\n      },\n      \"allOf\": [\n        {\n          \"if\": {\n            \"required\": [\n              \"source\"\n            ],\n            \"properties\": {\n              \"source\": {\n                \"required\": [\n                  \"type\"\n                ],\n                \"properties\": {\n                  \"type\": {\n                    \"not\": {\n                      \"enum\": [\n                        \"rebaseline\",\n                        \"restatement\"\n                      ]\n                    }\n                  }\n                }\n              }\n            }\n          },\n          \"then\": {\n            \"properties\": {\n              \"disposition\": {\n                \"minProperties\": 1\n              }\n            }\n          }\n        },\n        {\n          \"if\": {\n            \"required\": [\n              \"source\"\n            ],\n            \"properties\": {\n              \"source\": {\n                \"required\": [\n                  \"type\"\n                ],\n                \"properties\": {\n                  \"type\": {\n                    \"const\": \"restatement\"\n                  }\n                }\n              }\n            }\n          },\n          \"then\": {\n            \"properties\": {\n              \"disposition\": {\n                \"maxProperties\": 0\n              }\n            }\n          }\n        }\n      ],\n      \"dependentSchemas\": {\n        \"restatement\": {\n          \"properties\": {\n            \"source\": {\n              \"properties\": {\n                \"type\": {\n                  \"const\": \"restatement\"\n                }\n              }\n            }\n          },\n          \"$comment\": \"A restatement block may ride ONLY on a restatement-typed event, so no ordinary determination can smuggle one. The converse -- a restatement-typed event MUST carry the block -- is enforced by the ledger write path, not here: expressing it needs a conditional `required`, which the compatibility gate reads as a breaking change even though no artifact predating the `restatement` source type can trip it. A restatement-typed event with no block is inert: empty disposition, not an evidence class, nothing for a verifier to apply.\"\n        }\n      }\n    },\n    \"review_item\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"queued_at\",\n        \"source_ref\",\n        \"quote\",\n        \"author\",\n        \"status\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"queued_at\": {\n          \"type\": \"string\",\n          \"format\": \"date-time\"\n        },\n        \"source_ref\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Permalink to the comment/commit/ticket that triggered queueing.\"\n        },\n        \"quote\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Verbatim text of the statement.\"\n        },\n        \"author\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Username of the original author (may be unverified \\u2014 that can be why it was queued).\"\n        },\n        \"submitted_by\": {\n          \"$ref\": \"#/$defs/actor\",\n          \"description\": \"Optional provenance identifying the actor that submitted this review item; distinct from the original author.\"\n        },\n        \"suggested_finding_ref\": {\n          \"type\": \"string\"\n        },\n        \"suggested_disposition\": {\n          \"$ref\": \"#/$defs/disposition\"\n        },\n        \"queue_reason\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"ambiguous_statement\",\n            \"unverified_identity\",\n            \"insufficient_authority\",\n            \"no_finding_id\",\n            \"undetermined_finding\",\n            \"fp_audit_valve\",\n            \"needs_manual_test\",\n            \"stale_baseline\",\n            \"rebaseline_mapping\",\n            \"rebaseline_unmatched\",\n            \"unsound_refutation\",\n            \"weak_confirmation\",\n            \"ungraded_confirmation\",\n            \"severity_proposal\",\n            \"needs_identity\",\n            \"new_discovery\",\n            \"severity_upgrade\",\n            \"regression_detected\",\n            \"evidence_available\",\n            \"remediation_ready\",\n            \"remediation_due\",\n            \"false_positive_dispute\",\n            \"scope_change\",\n            \"revalidation_required\",\n            \"sla_escalation\",\n            \"manual_queue\",\n            \"dependency_update\",\n            \"patch_available\",\n            \"analyst_review\"\n          ],\n          \"$comment\": \"'undetermined_finding' = triage could not decide (unlocatable material or split votes) \\u2014 needs human review; 'fp_audit_valve' = deterministic 10% sample of auto-accepted machine false positives, queued so the auto-accept tier's error rate is continuously measured; 'needs_manual_test' = an unconfident triage confirmation (recall-mode split) awaiting a human PoC or live validation \\u2014 no validity event was emitted; 'stale_baseline' = a validation verdict was recorded against an audit baseline whose sha256 has since changed (the validation event emitter) \\u2014 the event was still emitted, but a human should confirm the claim it proves/refutes survived the baseline revision. 'unsound_refutation' = a machine refuted verdict was blocked by the refutation-soundness gate (error-signature transcript, zero-subject RBAC probe, or target-not-deployed run) \\u2014 no false_positive event was emitted; the finding awaits re-validation or human review. 'weak_confirmation' = an E3 inference-only confirmation quarantined by the validation event emitter \\u2014 no validity event until the evidence is graded E0-E2 or a human accepts it; 'ungraded_confirmation' = a post-0.179.0 confirmation with no evidence_grade, same quarantine; 'severity_proposal' = a validation-time severity re-rating proposal \\u2014 severity changes are human-only (disposition.severity), so the proposal queues for countersign instead of setting state. 'needs_identity' = a report was submitted with a finding carrying no `fingerprint`, so the receiver cannot say WHICH finding the statement is about. The submission is accepted and the statement queues here rather than being recorded: identity is computed by the producer, which holds the repository URL, the location paths and the CWE at audit time, and a receiver that hashes what the envelope happens to carry produces a confident wrong answer instead of no answer \\u2014 with `metadata.repository` absent, every unstamped finding sharing a path and CWE collapses to one identity across unrelated repositories. Re-submit stamped, or confirm the queued item by hand.\"\n        },\n        \"status\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"pending\",\n            \"confirmed\",\n            \"rejected\",\n            \"accepted\",\n            \"closed_unreviewed\"\n          ],\n          \"description\": \"Review queue state, independent of finding validity. 'accepted' is an actual review acceptance; historical 'confirmed' is its deprecated spelling. 'rejected' is an actual review rejection. 'closed_unreviewed' is administrative closure without review, not acceptance or rejection. Historical payloads remain valid; rejecting deprecated values at new writes is separate core record-door work.\"\n        },\n        \"resolution_note\": {\n          \"type\": \"string\",\n          \"description\": \"Why the item was accepted, rejected or administratively closed, and by whom; historical confirmed notes remain valid.\"\n        }\n      }\n    },\n    \"event_alias\": {\n      \"type\": \"object\",\n      \"description\": \"Rebaseline mapping carried ON THE EVENT (P2). Projected by the events module aliases_from_events() and merged OVER the legacy metadata.finding_aliases table by build_cumulative, so events are authoritative and the table is a rebuildable projection. The event's finding_ref is the superseded id; new_finding_ref is its successor.\",\n      \"required\": [\n        \"new_finding_ref\",\n        \"matched_by\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"new_finding_ref\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Successor finding id in the newer baseline.\"\n        },\n        \"matched_by\": {\n          \"enum\": [\n            \"fingerprint\",\n            \"path_cwe\",\n            \"path_set\",\n            \"title\",\n            \"manual\"\n          ],\n          \"description\": \"'fingerprint' matches auto-confirm; looser tiers pend in needs_review.\"\n        },\n        \"confirmed\": {\n          \"type\": \"boolean\",\n          \"description\": \"A human confirmed the mapping. Attribution comes from source.actor.identity, never duplicated here.\"\n        },\n        \"rejected\": {\n          \"type\": \"boolean\",\n          \"description\": \"A human refused the mapping; recorded rather than deleted (append-only).\"\n        },\n        \"similarity\": {\n          \"type\": \"number\"\n        },\n        \"path_overlap\": {\n          \"type\": \"number\"\n        },\n        \"from_report\": {\n          \"type\": \"string\",\n          \"description\": \"Basename of the superseded report.\"\n        },\n        \"note\": {\n          \"type\": \"string\"\n        }\n      }\n    },\n    \"event_finding\": {\n      \"type\": \"object\",\n      \"description\": \"A finding discovered between audits, carried ON THE EVENT because only baseline-writing audit skills may write a baseline. build_cumulative unions these with the baseline's findings when computing current state. The required keys are exactly the canonical claim fields metadata.claim_hashes hashes, so an event-carried claim is pinnable on the same terms as a baselined one. Extra keys mirroring report.schema.json $defs/finding are permitted.\",\n      \"required\": [\n        \"id\",\n        \"title\",\n        \"severity\",\n        \"cwes\",\n        \"locations\",\n        \"description\",\n        \"remediation\"\n      ],\n      \"properties\": {\n        \"id\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"title\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"severity\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"critical\",\n            \"high\",\n            \"medium\",\n            \"low\",\n            \"informational\"\n          ]\n        },\n        \"cwes\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"items\": {\n            \"type\": \"string\",\n            \"pattern\": \"^CWE-\\\\d{1,5}$\"\n          }\n        },\n        \"locations\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"object\"\n          }\n        },\n        \"description\": {\n          \"type\": \"string\"\n        },\n        \"remediation\": {\n          \"type\": \"string\"\n        },\n        \"validation_status\": {\n          \"type\": \"string\",\n          \"description\": \"'not_verified' on arrival \\u2014 nothing at scan time sets 'confirmed'; that takes execution evidence or a human, and those verdicts flow through the ledger.\"\n        },\n        \"origin\": {\n          \"type\": \"string\",\n          \"description\": \"Producing skill (vuln-scan, impact-analysis, verify-remediation) \\u2014 provenance for the extraction migration and for audit trails.\"\n        },\n        \"source_findings\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          },\n          \"description\": \"Refs into the producer's own report.\"\n        }\n      }\n    },\n    \"event_restatement\": {\n      \"type\": \"object\",\n      \"description\": \"An administrative restatement of signature-bound layer metadata: claim_hashes, audit_report_sha256, artifact_digests. Those fields are mutable, but rewriting one destroys the prior value, so an authorised restatement and a tampering rewrite were previously indistinguishable. Appending this event moves the Merkle root, which drops the stale signature and re-signs through the normal write path -- no new signature format is needed, because format 4 already binds the root and all three digests. DELTA, NOT SNAPSHOT: for map-valued targets, before and after carry only the entries being changed, so an event's size tracks the change rather than the layer. Restating the same key twice appends two small events instead of two copies of the whole map. Restatements stack in append order and a wrong one is superseded by another, never removed; a restatement may not restore a value the chain already retired -- that reverses an earlier decision and needs its own stated reason. EVENT CONTENT IS NOT A TARGET: events are immutable and are corrected by appending a superseding determination. finding_aliases is not a target either: a rebaseline event already records a rename in the Merkle-covered event stream, so the metadata table is a rebuildable projection rather than authority.\",\n      \"required\": [\n        \"target\",\n        \"reason\",\n        \"before\",\n        \"after\",\n        \"authority\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"target\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"claim_hashes\",\n            \"audit_report_sha256\",\n            \"artifact_digests\"\n          ],\n          \"description\": \"The metadata field being restated. The restatement is the authority for that field's new value, and a verifier replays the chain to check it.\"\n        },\n        \"reason\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"data_error\",\n            \"schema_migration\",\n            \"baseline_rewrite\",\n            \"operator_error\"\n          ],\n          \"description\": \"Why the data was wrong. 'baseline_rewrite' is the legitimate re-audit/reissue case that previously surfaced only as an unexplained digest mismatch; 'schema_migration' is a mechanical v1->v2 reshape and must name both versions.\"\n        },\n        \"schema_from\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"schema_to\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"before\": {\n          \"description\": \"The entries being replaced, as they stand. For a map-valued target this is the changed keys only, never the whole map. Required: a restatement asserts what it replaces, and a writer refuses one whose before disagrees with what is stored.\"\n        },\n        \"after\": {\n          \"description\": \"The replacement entries. For a map-valued target this is the changed keys only; the writer merges them over the stored map. A null value deletes that key.\"\n        },\n        \"authority\": {\n          \"type\": \"object\",\n          \"required\": [\n            \"ticket\"\n          ],\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"ticket\": {\n              \"type\": \"string\",\n              \"minLength\": 1,\n              \"description\": \"The change record this restatement was authorised under. Required: an unattributed restatement is indistinguishable from the rewrite this mechanism exists to prevent.\"\n            },\n            \"approved_by\": {\n              \"type\": \"string\",\n              \"minLength\": 1\n            }\n          }\n        }\n      },\n      \"allOf\": [\n        {\n          \"if\": {\n            \"properties\": {\n              \"target\": {\n                \"const\": \"audit_report_sha256\"\n              }\n            },\n            \"required\": [\n              \"target\"\n            ]\n          },\n          \"then\": {\n            \"properties\": {\n              \"before\": {\n                \"type\": [\n                  \"string\",\n                  \"null\"\n                ],\n                \"pattern\": \"^[0-9a-f]{64}$\"\n              },\n              \"after\": {\n                \"type\": \"string\",\n                \"pattern\": \"^[0-9a-f]{64}$\"\n              }\n            }\n          }\n        },\n        {\n          \"if\": {\n            \"properties\": {\n              \"target\": {\n                \"enum\": [\n                  \"claim_hashes\",\n                  \"artifact_digests\"\n                ]\n              }\n            },\n            \"required\": [\n              \"target\"\n            ]\n          },\n          \"then\": {\n            \"properties\": {\n              \"before\": {\n                \"type\": \"object\",\n                \"minProperties\": 1\n              },\n              \"after\": {\n                \"type\": \"object\",\n                \"minProperties\": 1\n              }\n            }\n          }\n        },\n        {\n          \"if\": {\n            \"properties\": {\n              \"reason\": {\n                \"const\": \"schema_migration\"\n              }\n            },\n            \"required\": [\n              \"reason\"\n            ]\n          },\n          \"then\": {\n            \"required\": [\n              \"schema_from\",\n              \"schema_to\"\n            ]\n          }\n        }\n      ]\n    }\n  }\n}\n",
	"manifest-coverage": `{
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/manifest-coverage.schema.json",
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "title": "Manifest SBOM Coverage",
  "description": "What the SBOM tool could not resolve when it read one commit's dependency manifests, keyed by ecosystem and canonical repository-relative manifest path. Stored as the artifact of the current manifest_sbom assessment, not tool_scan. Readers must match assessment_id to both the owning assessment and the commit's current manifest_sbom scan assessment, and sbom_id to the commit's current manifest_sbom_id; reject mismatches. The database-generated assessment identifies this scan even when normalized SBOM content is unchanged. A package missing from the SBOM in an unresolved ecosystem is unknown, never absent. An empty unresolved object proves a complete read only for the matching scan and SBOM; absent or unmatched coverage is unknown, never complete. A newer scan without coverage must not reuse older coverage.",
  "type": "object",
  "required": ["artifact", "tool", "sbom_id", "assessment_id", "unresolved"],
  "additionalProperties": false,
  "properties": {
    "artifact": {"const": "manifest-coverage"},
    "tool": {
      "type": "object",
      "required": ["name", "version"],
      "additionalProperties": false,
      "properties": {
        "name": {"type": "string", "minLength": 1},
        "version": {"type": "string", "minLength": 1}
      }
    },
    "sbom_id": {
      "description": "The native database SBOM UUID at the commit's current manifest_sbom_id, not a content digest.",
      "type": "string",
      "format": "uuid"
    },
    "assessment_id": {
      "description": "The database-generated UUID of the owning, current manifest_sbom assessment for this scan, not a synthesized identifier or SBOM content identity.",
      "type": "string",
      "format": "uuid"
    },
    "unresolved": {
      "type": "object",
      "propertyNames": {
        "$ref": "impact-analysis.schema.json#/$defs/metadata/properties/ecosystem"
      },
      "additionalProperties": {
        "type": "object",
        "minProperties": 1,
        "propertyNames": {"$ref": "#/$defs/path"},
        "additionalProperties": {"$ref": "#/$defs/reason"}
      }
    }
  },
  "$defs": {
    "path": {
      "description": "The canonical manifest path relative to the repository root: no absolute paths, drive prefixes, backslashes, empty segments, or dot segments.",
      "type": "string",
      "minLength": 1,
      "pattern": "^(?![A-Za-z]:)(?![\\s\\S]*(?:^|/)\\.{1,2}(?:/|$))[^/\\\\\u0000]+(?:/[^/\\\\\u0000]+)*$"
    },
    "reason": {
      "description": "no-lockfile: a manifest whose versions only a lockfile pins, with none beside it or above it in the tree. unpinned: a manifest the tool reads directly that declares requirements without an exact version (a requirements file's ranges), which it drops.",
      "type": "string",
      "enum": ["no-lockfile", "unpinned"]
    }
  }
}
`,
	"operator-priv-profile": "{\n  \"$schema\": \"http://json-schema.org/draft-07/schema#\",\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/operator-priv-profile.schema.json\",\n  \"title\": \"Operator privilege profile\",\n  \"description\": \"The privilege an operator ASKS FOR, read from its shipped manifests: workloads and their host/privilege posture, the RBAC rules its bundle declares, the SCCs it requests and ships, and the install modes its OperatorGroup allows. This is a DECLARED-STATE profile -- it is parsed from manifests in the repository and never from a live cluster, so it answers 'what would this grant if installed' and not 'what is granted right now'. Runtime SCC assignment needs a separate capture and is out of scope by design.\",\n  \"type\": \"object\",\n  \"required\": [\n    \"repo\",\n    \"summary\"\n  ],\n  \"additionalProperties\": false,\n  \"properties\": {\n    \"repo\": {\n      \"type\": \"string\",\n      \"minLength\": 1\n    },\n    \"tier\": {\n      \"type\": \"string\",\n      \"description\": \"Which capture tiers this profile covers, stated in prose because the boundary matters when reading any count here: static manifest tiers only, unless a runtime capture was folded in.\"\n    },\n    \"workloads\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/workload\"\n      }\n    },\n    \"rbac_rules\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"type\": \"object\"\n      },\n      \"description\": \"The policy rules the bundle declares, as parsed. Kept whole because a least-privilege review reads the actual rule, not a count of rules.\"\n    },\n    \"rbac_flags\": {\n      \"type\": \"object\",\n      \"description\": \"Rules matching a high-privilege pattern, grouped by pattern. A key is present ONLY when it matched, so an absent key means no match rather than unknown.\",\n      \"additionalProperties\": {\n        \"type\": \"array\",\n        \"items\": {\n          \"type\": \"string\"\n        }\n      },\n      \"properties\": {\n        \"secrets_access\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"nodes_access\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"wildcard_verbs\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"wildcard_resources\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"rbac_write\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          },\n          \"description\": \"Write access to RBAC itself. Distinct from escalate/bind because it is a slower path to the same place.\"\n        },\n        \"pods_exec\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"escalate_bind_impersonate\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          },\n          \"description\": \"The three verbs that let a principal grant itself more than it holds. The highest-signal group here.\"\n        }\n      }\n    },\n    \"scc_requests\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"type\": \"object\"\n      },\n      \"description\": \"SCCs the bundle explicitly requests.\"\n    },\n    \"sccs_shipped\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"type\": \"object\"\n      },\n      \"description\": \"SCC objects the bundle ships. Shipping an SCC is not the same as requesting one, and the two are counted separately for that reason.\"\n    },\n    \"namespaces\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"type\": \"string\"\n      }\n    },\n    \"install_modes\": {\n      \"type\": \"object\",\n      \"description\": \"Which OperatorGroup install modes the CSV permits. AllNamespaces is the one that turns a namespaced ask into a cluster-wide one.\",\n      \"additionalProperties\": true\n    },\n    \"operatorgroups\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"type\": \"object\"\n      }\n    },\n    \"tier2_required_vs_granted\": {\n      \"type\": \"object\",\n      \"description\": \"Kubebuilder RBAC markers in source compared with what the bundle actually ships. `shipped_not_declared` is privilege present in the bundle with no marker asking for it.\",\n      \"additionalProperties\": true\n    },\n    \"example_or_test_manifests_excluded\": {\n      \"type\": \"object\",\n      \"description\": \"What was deliberately left out. Recorded because a profile that silently swallowed sample manifests would report privilege the operator never asks for in production.\",\n      \"additionalProperties\": true\n    },\n    \"summary\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"workloads\",\n        \"rbac_rules\"\n      ],\n      \"additionalProperties\": true,\n      \"properties\": {\n        \"workloads\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"privileged_or_host_workloads\": {\n          \"type\": \"integer\",\n          \"minimum\": 0,\n          \"description\": \"Workloads asking for privileged mode or a host namespace. The headline least-privilege number.\"\n        },\n        \"rbac_rules\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"distinct_rule_triples\": {\n          \"type\": \"integer\",\n          \"minimum\": 0,\n          \"description\": \"Distinct (apiGroup, resource, verb) triples. The de-duplicated size of the ask -- rule COUNT inflates with how the bundle is authored, this does not.\"\n        },\n        \"distinct_cluster_triples\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"cluster_scoped_rules\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"wildcard_rules\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"scc_requests\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"no_scc_request_recorded\": {\n          \"type\": \"boolean\",\n          \"description\": \"True when no SCC request was found. Explicitly distinct from an empty request list, because 'asked for nothing' and 'we could not tell' must not read the same.\"\n        }\n      }\n    }\n  },\n  \"$defs\": {\n    \"workload\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"kind\"\n      ],\n      \"additionalProperties\": true,\n      \"properties\": {\n        \"kind\": {\n          \"type\": \"string\"\n        },\n        \"name\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"description\": \"Null for a manifest fragment that patches a workload without naming it -- kustomize patches and sidecar overlays. Measured: 32 such workloads across the live corpus. Always a STRING when present: a Kubernetes name is an RFC 1123 label, and a producer must coerce it, because YAML parses an unquoted `name: 0` as an integer while `name: \\\"0\\\"` on the same object is a string.\"\n        },\n        \"manifest\": {\n          \"type\": \"string\"\n        },\n        \"serviceAccountName\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"hostNetwork\": {\n          \"type\": \"boolean\"\n        },\n        \"hostPID\": {\n          \"type\": \"boolean\"\n        },\n        \"hostIPC\": {\n          \"type\": \"boolean\"\n        },\n        \"hostPath_volumes\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"pod_securityContext\": {\n          \"type\": \"object\"\n        },\n        \"containers\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"object\"\n          }\n        }\n      }\n    }\n  }\n}\n",
	"org-parameters": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/org-parameters.schema.json",
  "title": "Organization-defined parameters (ODPs)",
  "description": "Validates your compliance configuration directory's organization-defined parameters file \u2014 declared values for NIST 800-53 '[Assignment: organization-defined ...]' slots and analogous knobs in other frameworks (determinism amendment 4). A check whose parameter is undeclared verdicts not_assessed(parameter undeclared); the LLM never assumes a value.",
  "type": "object",
  "required": [
    "version",
    "declared_by",
    "parameters"
  ],
  "additionalProperties": false,
  "properties": {
    "version": {
      "type": "integer",
      "minimum": 1
    },
    "declared_by": {
      "type": "string",
      "description": "who owns these values (team/role)"
    },
    "declared_on": {
      "type": "string",
      "pattern": "^\\d{4}-\\d{2}-\\d{2}$"
    },
    "note": {
      "type": "string"
    },
    "parameters": {
      "type": "object",
      "minProperties": 1,
      "propertyNames": {
        "pattern": "^[a-z0-9][a-z0-9._-]*$"
      },
      "additionalProperties": {
        "type": "object",
        "required": [
          "value",
          "controls"
        ],
        "additionalProperties": false,
        "properties": {
          "value": {},
          "controls": {
            "type": "array",
            "minItems": 1,
            "items": {
              "type": "string"
            },
            "description": "framework:control_id slots this value fills"
          },
          "note": {
            "type": "string"
          }
        }
      }
    }
  }
}
`,
	"pqc-blockers": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/pqc-blockers.schema.json",
  "title": "PQC blockers projection",
  "$comment": "Shape contract for <slug>-pqc-blockers.json emitted by the pqc-readiness build script \u2014 a deterministic projection of pqc-readiness remediations[] into the security-audit blockers VOCABULARY (same field names and semantics as schema/report.schema.json blockers) without claiming report.schema.json conformance. Kept deliberately parallel, not unified: converging the two into one contracts data model is a planned separate feature; until then this schema must not drift from the finding-field vocabulary it mirrors.",
  "type": "object",
  "required": [
    "artifact",
    "title",
    "metadata",
    "executive_summary",
    "severity_criteria",
    "findings",
    "findings_summary",
    "remediation_roadmap"
  ],
  "properties": {
    "artifact": {
      "const": "pqc-blockers"
    },
    "title": {
      "type": "string",
      "minLength": 5
    },
    "metadata": {
      "type": "object",
      "required": [
        "date",
        "scope",
        "repository"
      ],
      "properties": {
        "date": {
          "type": "string"
        },
        "scope": {
          "type": "string"
        },
        "repository": {
          "type": "string"
        },
        "commit": {
          "type": [
            "string",
            "null"
          ]
        },
        "framework": {
          "type": "string"
        },
        "methodology": {
          "type": "string"
        },
        "additional": {
          "type": "object",
          "required": [
            "source_artifact",
            "who_sets_tls",
            "excluded_remediations"
          ],
          "properties": {
            "harness_version": {
              "type": "string"
            },
            "source_artifact": {
              "type": "string"
            },
            "adapter_version": {
              "type": "string"
            },
            "who_sets_tls": {
              "type": "string"
            },
            "readiness_status": {
              "type": "string"
            },
            "excluded_remediations": {
              "type": "array",
              "items": {
                "type": "string"
              },
              "$comment": "Remediation ids dropped for lacking locations[] \u2014 recorded so exclusion is never silent."
            }
          }
        }
      }
    },
    "executive_summary": {
      "type": "object",
      "required": [
        "prose",
        "severity_counts"
      ],
      "properties": {
        "prose": {
          "type": "string"
        },
        "severity_counts": {
          "type": "object"
        }
      }
    },
    "severity_criteria": {
      "type": "array",
      "minItems": 4,
      "items": {
        "type": "object",
        "required": [
          "level",
          "definition"
        ],
        "properties": {
          "level": {
            "$ref": "#/$defs/severity_level"
          },
          "definition": {
            "type": "string"
          }
        }
      }
    },
    "findings": {
      "type": "array",
      "items": {
        "$ref": "#/$defs/finding"
      }
    },
    "findings_summary": {
      "type": "array",
      "minItems": 4,
      "items": {
        "type": "object",
        "required": [
          "severity",
          "count",
          "finding_ids"
        ],
        "properties": {
          "severity": {
            "$ref": "#/$defs/severity_level"
          },
          "count": {
            "type": "integer"
          },
          "finding_ids": {
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        }
      }
    },
    "remediation_roadmap": {
      "type": "array",
      "minItems": 1,
      "items": {
        "type": "object",
        "required": [
          "priority",
          "action",
          "addresses"
        ],
        "properties": {
          "priority": {
            "type": "string"
          },
          "effort": {},
          "blocked_external": {},
          "action": {
            "type": "string"
          },
          "addresses": {
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "$defs": {
    "severity_level": {
      "type": "string",
      "enum": [
        "critical",
        "high",
        "medium",
        "low",
        "informational"
      ]
    },
    "finding": {
      "type": "object",
      "required": [
        "id",
        "title",
        "severity",
        "cwes",
        "locations",
        "description",
        "remediation",
        "category",
        "validation_status",
        "pqc_classification"
      ],
      "properties": {
        "id": {
          "type": "string",
          "pattern": "^[A-Z0-9_]{1,24}-(?:[0-9a-f]{7}|u[0-9a-f]{6})-REM-\\d{3}$",
          "$comment": "The source remediation id, kept stable across re-emission (not the report.schema.json canonical finding-id form)."
        },
        "title": {
          "type": "string"
        },
        "severity": {
          "$ref": "#/$defs/severity_level"
        },
        "cwes": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "string",
            "pattern": "^CWE-\\d{1,5}$"
          }
        },
        "locations": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "object",
            "required": [
              "path"
            ],
            "properties": {
              "path": {
                "type": "string"
              },
              "lines": {
                "type": "string"
              }
            }
          }
        },
        "description": {
          "type": "string"
        },
        "remediation": {
          "type": "string"
        },
        "category": {
          "const": "cryptography"
        },
        "validation_status": {
          "const": "not_verified"
        },
        "pqc_classification": {
          "type": "string",
          "enum": [
            "shor-key-establishment",
            "shor-signature",
            "clock-2030-parameter",
            "classically-broken",
            "hndl-exposure",
            "pqc-blocker-config",
            "pqc-adoption"
          ],
          "$comment": "Reuses report.schema.json's pqc_classification vocabulary verbatim \u2014 do not extend here without extending there."
        },
        "remediation_effort": {
          "type": ["string", "null"],
          "enum": [
            "trivial",
            "moderate",
            "significant",
            "blocked-external",
            "xs",
            "s",
            "m",
            "l",
            "xl",
            null
          ]
        },
        "blocked_external": {}
      }
    },
    "canonical_write": {
      "description": "Opt-in canonical effort sizes or null for unknown size; open historical extras remain readable.",
      "allOf": [
        { "$ref": "#" },
        {
          "properties": {
            "findings": {
              "items": {
                "properties": {
                  "remediation_effort": { "$ref": "report.schema.json#/$defs/effort" },
                  "blocked_external": { "type": "boolean" }
                }
              }
            },
            "remediation_roadmap": {
              "items": {
                "properties": {
                  "priority": { "$ref": "report.schema.json#/$defs/roadmap_priority" },
                  "effort": { "$ref": "report.schema.json#/$defs/effort" },
                  "blocked_external": { "type": "boolean" }
                }
              }
            }
          }
        }
      ]
    }
  }
}
`,
	"pqc-decision-tree": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/pqc-decision-tree.schema.json",
  "title": "PQC readiness decision tree",
  "description": "Validates the PQC readiness decision tree artifact \u2014 the machine-readable provenance tree, remediation-effort mapping, readiness-bucket rules, TLS-control crosswalk, and FIPS interaction rules (plan v1.3 sections 4.3/4.5-4.7). Structural gate; rule semantics are exercised by the pilot cross-validation.",
  "type": "object",
  "required": [
    "tree_version",
    "provenance_tree",
    "remediation_effort",
    "readiness_buckets",
    "tls_control_crosswalk",
    "fips_interaction",
    "pqc_classification_map"
  ],
  "additionalProperties": false,
  "properties": {
    "tree_version": {
      "type": "string",
      "pattern": "^\\d+\\.\\d+\\.\\d+$"
    },
    "plan": {
      "type": "string"
    },
    "schema": {
      "type": "string"
    },
    "provenance_tree": {
      "type": "object",
      "required": [
        "rules"
      ],
      "properties": {
        "_comment": {
          "type": "string"
        },
        "rules": {
          "type": "array",
          "minItems": 7,
          "items": {
            "type": "object",
            "required": [
              "provenance",
              "when"
            ],
            "properties": {
              "provenance": {
                "$ref": "#/definitions/provenance"
              },
              "when": {
                "type": "string",
                "minLength": 10
              }
            }
          }
        }
      }
    },
    "remediation_effort": {
      "type": "object",
      "required": [
        "rules",
        "classes"
      ],
      "properties": {
        "_comment": {
          "type": "string"
        },
        "classes": {
          "anyOf": [
            { "const": ["trivial", "moderate", "significant", "blocked-external"] },
            { "const": ["xs", "s", "m", "l", "xl"] }
          ]
        },
        "rules": {
          "type": "array",
          "minItems": 4,
          "items": {
            "type": "object",
            "required": [
              "effort",
              "when"
            ],
            "properties": {
              "effort": {
                "$ref": "#/definitions/effort"
              },
              "blocked_external": {},
              "when": {
                "type": "string",
                "minLength": 10
              }
            }
          }
        }
      }
    },
    "readiness_buckets": {
      "type": "object",
      "required": [
        "score_ready_min",
        "score_partial_min",
        "rules",
        "buckets"
      ],
      "properties": {
        "_comment": {
          "type": "string"
        },
        "score_ready_min": {
          "type": "number",
          "minimum": 0,
          "maximum": 100
        },
        "score_partial_min": {
          "type": "number",
          "minimum": 0,
          "maximum": 100
        },
        "buckets": {
          "const": [
            "ready",
            "partial",
            "not-ready",
            "blocked-external",
            "not-applicable"
          ]
        },
        "rules": {
          "type": "array",
          "minItems": 5,
          "items": {
            "type": "object",
            "required": [
              "bucket",
              "when"
            ],
            "properties": {
              "bucket": {
                "$ref": "#/definitions/bucket"
              },
              "when": {
                "type": "string",
                "minLength": 10
              }
            }
          }
        }
      }
    },
    "tls_control_crosswalk": {
      "type": "object",
      "required": [
        "app-controlled",
        "infra-controlled",
        "vendor-controlled"
      ],
      "properties": {
        "_comment": {
          "type": "string"
        },
        "app-controlled": {
          "$ref": "#/definitions/provList"
        },
        "infra-controlled": {
          "$ref": "#/definitions/provList"
        },
        "vendor-controlled": {
          "$ref": "#/definitions/provList"
        }
      },
      "additionalProperties": false
    },
    "fips_interaction": {
      "type": "object",
      "required": [
        "rules"
      ],
      "properties": {
        "_comment": {
          "type": "string"
        },
        "rules": {
          "type": "array",
          "minItems": 3,
          "items": {
            "type": "object",
            "required": [
              "verdict",
              "when"
            ],
            "properties": {
              "verdict": {
                "enum": [
                  "blocked-by-provider-version",
                  "no-penalty",
                  "fips-validation-gap",
                  "pqc-blocked-by-fips-mode"
                ]
              },
              "when": {
                "type": "string",
                "minLength": 10
              },
              "treat_as": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "pqc_classification_map": {
      "type": "object",
      "properties": {
        "_comment": {
          "type": "string"
        }
      },
      "additionalProperties": {
        "enum": [
          "shor-key-establishment",
          "shor-signature",
          "clock-2030-parameter",
          "classically-broken",
          "hndl-exposure",
          "pqc-blocker-config",
          "pqc-adoption"
        ]
      }
    },
    "server_side_caveat": {
      "type": "object",
      "required": [
        "rules"
      ],
      "properties": {
        "_comment": {
          "type": "string"
        },
        "rules": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "object",
            "required": [
              "caveat",
              "when"
            ],
            "properties": {
              "caveat": {
                "type": "string",
                "minLength": 3
              },
              "when": {
                "type": "string",
                "minLength": 10
              },
              "scoring_ceiling": {
                "enum": [
                  "partial",
                  "no"
                ]
              },
              "treat_as": {
                "type": "string"
              },
              "_comment": {
                "type": "string"
              }
            },
            "additionalProperties": false
          }
        }
      },
      "additionalProperties": false,
      "description": "Infrastructure-governance scoring ceilings (harness >= 0.131.1): infrastructure-governed decision points cannot claim PQC capability from the client toolchain alone; these rules cap the relevant check results until runtime evidence confirms the server side."
    }
  },
  "definitions": {
    "provenance": {
      "enum": [
        "inherited-platform",
        "inherited-constrained",
        "delegated-dependency",
        "native-first-party",
        "vendored",
        "externalized",
        "not-assessable-from-source"
      ]
    },
    "effort": {
      "enum": [
        "trivial",
        "moderate",
        "significant",
        "blocked-external",
        "xs",
        "s",
        "m",
        "l",
        "xl",
        null
      ]
    },
    "canonical_write": {
      "description": "Opt-in canonical sizes or null for unknown size; existing effort requiredness is retained. Blocking is independent.",
      "allOf": [
        { "$ref": "#" },
        {
          "properties": {
            "remediation_effort": {
              "properties": {
                "classes": { "const": ["xs", "s", "m", "l", "xl"] },
                "rules": {
                  "items": {
                    "properties": {
                      "effort": { "$ref": "report.schema.json#/$defs/effort" },
                      "blocked_external": { "type": "boolean" }
                    }
                  }
                }
              }
            }
          }
        }
      ]
    },
    "bucket": {
      "enum": [
        "ready",
        "partial",
        "not-ready",
        "blocked-external",
        "not-applicable"
      ]
    },
    "provList": {
      "type": "array",
      "minItems": 1,
      "uniqueItems": true,
      "items": {
        "$ref": "#/definitions/provenance"
      }
    }
  }
}
`,
	"pqc-facts": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/pqc-facts.schema.json",
  "title": "PQC facts (Layer 1 crypto census)",
  "description": "Deterministic *-pqc-facts.json artifact written by the pqc-readiness adapter (over the pinned pqc-scan binary). Facts never contain verdicts. Consumers: pqc-readiness Layer 2 scoring, /patch pqc ingest, bulk prescan, L2 prepass, PQC dashboards. Strict on the fields those consumers key on; permissive on additions (new adapter versions may add fields).",
  "type": "object",
  "required": [
    "artifact",
    "repository",
    "stamps",
    "coverage",
    "summary",
    "facts"
  ],
  "properties": {
    "artifact": {
      "const": "pqc-facts"
    },
    "repository": {
      "type": "string",
      "minLength": 1
    },
    "stamps": {
      "type": "object",
      "required": [
        "adapter_version",
        "pqc_scan_commit",
        "rules_sha256"
      ],
      "properties": {
        "adapter_version": {
          "type": "string"
        },
        "pqc_scan_commit": {
          "type": "string"
        },
        "rules_sha256": {
          "type": "string"
        },
        "binary_sha256": {
          "type": "string"
        }
      }
    },
    "coverage": {
      "type": "object",
      "required": [
        "assessment_basis",
        "scanned_files",
        "skipped_files",
        "rules_in_pack"
      ],
      "properties": {
        "assessment_basis": {
          "enum": [
            "source",
            "sbom-only"
          ]
        },
        "scanned_files": {
          "type": [
            "integer",
            "null"
          ]
        },
        "skipped_files": {
          "type": [
            "integer",
            "null"
          ]
        },
        "rules_in_pack": {
          "type": [
            "integer",
            "null"
          ]
        },
        "no_crypto_detected_assertion": {
          "type": "string"
        }
      }
    },
    "summary": {
      "type": "object",
      "required": [
        "by_rule",
        "by_qclass",
        "by_provenance_hint",
        "by_path_class"
      ],
      "properties": {
        "by_rule": {
          "$ref": "#/$defs/countMap"
        },
        "by_qclass": {
          "$ref": "#/$defs/countMap"
        },
        "by_provenance_hint": {
          "$ref": "#/$defs/countMap"
        },
        "by_path_class": {
          "$ref": "#/$defs/countMap"
        }
      }
    },
    "facts": {
      "type": "array",
      "items": {
        "$ref": "#/$defs/fact"
      }
    }
  },
  "$defs": {
    "countMap": {
      "type": "object",
      "additionalProperties": {
        "type": "integer"
      }
    },
    "fact": {
      "type": "object",
      "required": [
        "fact_id",
        "rule_id",
        "file",
        "line",
        "detail",
        "ir8547",
        "provenance_hint",
        "path_class"
      ],
      "properties": {
        "fact_id": {
          "type": "string",
          "pattern": "^F[0-9]{4,}$"
        },
        "rule_id": {
          "type": "string",
          "minLength": 1
        },
        "file": {
          "type": "string"
        },
        "line": {
          "type": [
            "integer",
            "null"
          ]
        },
        "detail": {
          "type": "string"
        },
        "match": {
          "type": "string"
        },
        "provider": {
          "type": "string"
        },
        "pqc_capable": {
          "type": "boolean"
        },
        "capability_hint": {
          "type": "string"
        },
        "fips_mode": {
          "type": "boolean"
        },
        "finding_id": {
          "type": [
            "string",
            "null"
          ]
        },
        "severity": {
          "type": [
            "string",
            "null"
          ]
        },
        "risk": {
          "type": [
            "string",
            "null"
          ]
        },
        "confidence": {
          "type": [
            "number",
            "null"
          ]
        },
        "provenance_hint": {
          "type": "string"
        },
        "path_class": {
          "enum": [
            "first_party",
            "vendor",
            "test_docs"
          ]
        },
        "ir8547": {
          "type": "object",
          "required": [
            "qclass",
            "usage",
            "clock"
          ],
          "properties": {
            "qclass": {
              "type": "string"
            },
            "qclass_resolved": {
              "type": "string"
            },
            "usage": {
              "type": "string"
            },
            "matched_prefix": {
              "type": "string"
            },
            "clock": {
              "type": [
                "object",
                "null"
              ],
              "properties": {
                "status_now": {
                  "type": "string"
                },
                "deprecated_after": {
                  "type": [
                    "integer",
                    "null"
                  ]
                },
                "disallowed_after": {
                  "type": [
                    "integer",
                    "null"
                  ]
                }
              }
            }
          }
        }
      }
    }
  }
}
`,
	"pqc-readiness": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/pqc-readiness.schema.json",
  "title": "PQC readiness report",
  "description": "Layer-2 agent-authored PQC readiness report. Enforcement gate: pqc_facts.py --validate-readiness (mirror of this schema). Every non-yes check result MUST cite fact IDs from the sibling pqc-facts artifact.",
  "type": "object",
  "required": [
    "title",
    "metadata",
    "scores",
    "flags",
    "provenance_summary"
  ],
  "properties": {
    "title": {
      "type": "string"
    },
    "metadata": {
      "type": "object",
      "required": [
        "repository",
        "assessment_basis",
        "tool"
      ],
      "properties": {
        "repository": {
          "type": "string",
          "pattern": "^(https?://|git@)"
        },
        "commit": {
          "type": [
            "string",
            "null"
          ]
        },
        "assessed_at": {
          "type": "string"
        },
        "assessment_basis": {
          "enum": [
            "source",
            "sbom-only",
            "source+runtime"
          ]
        },
        "facts_ref": {
          "type": "string"
        },
        "cbom_ref": {
          "type": [
            "string",
            "null"
          ]
        },
        "tool": {
          "type": "object",
          "required": [
            "pqc_scan_commit",
            "rules_sha256"
          ],
          "properties": {
            "pqc_scan_commit": {
              "type": "string"
            },
            "rules_sha256": {
              "type": "string"
            },
            "binary_sha256": {
              "type": [
                "string",
                "null"
              ]
            },
            "adapter_version": {
              "type": "string"
            }
          }
        }
      }
    },
    "scores": {
      "type": "object",
      "required": [
        "VULN",
        "AGIL",
        "PQCA",
        "HNDL",
        "overall"
      ],
      "properties": {
        "VULN": {
          "$ref": "#/definitions/domain"
        },
        "AGIL": {
          "$ref": "#/definitions/domain"
        },
        "PQCA": {
          "$ref": "#/definitions/domain"
        },
        "HNDL": {
          "$ref": "#/definitions/domain"
        },
        "overall": {
          "type": "number",
          "minimum": 0,
          "maximum": 100
        }
      }
    },
    "flags": {
      "type": "object",
      "required": [
        "has_2030_clock_items",
        "hndl_priority",
        "runtime_verification_required"
      ],
      "properties": {
        "has_2030_clock_items": {
          "type": "boolean"
        },
        "hndl_priority": {
          "type": "boolean"
        },
        "runtime_verification_required": {
          "type": "boolean"
        }
      }
    },
    "provenance_summary": {
      "type": "object",
      "required": [
        "counts",
        "dominant"
      ],
      "properties": {
        "counts": {
          "type": "object",
          "additionalProperties": {
            "type": "integer"
          }
        },
        "dominant": {
          "enum": [
            "inherited-platform",
            "inherited-constrained",
            "delegated-dependency",
            "native-first-party",
            "vendored",
            "externalized",
            "not-assessable-from-source",
            "none"
          ]
        }
      }
    },
    "clock_items": {
      "type": "array",
      "items": {
        "type": "object",
        "required": [
          "fact_ids",
          "primitive",
          "disallowed_after"
        ],
        "properties": {
          "fact_ids": {
            "type": "array",
            "items": {
              "type": "string"
            }
          },
          "primitive": {
            "type": "string"
          },
          "deprecated_after": {
            "type": [
              "integer",
              "null"
            ]
          },
          "disallowed_after": {
            "type": [
              "integer",
              "null"
            ]
          },
          "remediation_effort": {
            "enum": [
              "trivial",
              "moderate",
              "significant",
              "blocked-external",
              "xs",
              "s",
              "m",
              "l",
              "xl",
              null
            ],
            "$comment": "Historical labels follow the existing v1.3 provenance x agility mapping. The xs–xl size labels have no assigned legacy replacement mapping."
          },
          "blocked_external": {},
          "blast_radius": {
            "enum": [
              "fleet",
              "product",
              "service",
              "component"
            ],
            "$comment": "Scope of impact if this clock item is not remediated; fleet = cluster-wide trust anchor"
          }
        }
      }
    },
    "readiness_bucket": {
      "enum": [
        "ready",
        "partial",
        "not-ready",
        "blocked-external",
        "not-applicable"
      ],
      "$comment": "v1.3 (plan 4.6): deterministic roll-up of scores + clock items + provenance; thresholds in tables/pqc-readiness-decision-tree.json (calibration knob 5)"
    },
    "fips_interaction": {
      "type": "object",
      "required": [
        "verdict",
        "fact_ids"
      ],
      "properties": {
        "verdict": {
          "enum": [
            "pqc-blocked-by-fips-mode",
            "blocked-by-provider-version",
            "no-penalty",
            "fips-validation-gap"
          ]
        },
        "fact_ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "rationale": {
          "type": "string"
        }
      },
      "$comment": "v1.3 (plan 4.7): present only when a FIPS-mode fact fired; conditional, never a blanket blocker"
    },
    "runtime_evidence": {
      "type": "object",
      "properties": {
        "validation_report_ref": {
          "type": "string"
        },
        "pqc_caps": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "closed_runtime_verification": {
          "type": "boolean"
        },
        "runtime_confirmed_at": {
          "type": "string",
          "format": "date",
          "$comment": "ISO date when runtime evidence was collected"
        },
        "staleness_window_days": {
          "type": "integer",
          "default": 90,
          "$comment": "After this many days the runtime evidence should be refreshed"
        }
      },
      "$comment": "v1.3: probe results ingested from a validate-findings run (chain.PQC_CAPS vocabulary); when present and conclusive, flags.runtime_verification_required may be set false with this as the citation"
    },
    "server_side_caveats": {
      "type": "array",
      "items": {
        "type": "object",
        "required": [
          "fact_ids",
          "caveat_type"
        ],
        "properties": {
          "fact_ids": {
            "type": "array",
            "items": {
              "type": "string"
            }
          },
          "caveat_type": {
            "enum": [
              "infra_mlkem_unverified",
              "infra_mlkem_blocked"
            ]
          },
          "rationale": {
            "type": "string"
          }
        }
      },
      "$comment": "Infrastructure-governed decision points where the server-side peer capability is unverified or known-blocked"
    },
    "notes": {
      "type": "string"
    },
    "remediations": {
      "type": "array",
      "$comment": "harness >= 0.153.0: structured remediation plan mirroring the MD companion's 'What you need to do' routing (fix-now / upgrade / waiting-on-upstream / deadline). Optional for pre-0.153.0 reports; mandatory for new reports and backfilled via backfill_remediations.py.",
      "items": {
        "type": "object",
        "required": [
          "id",
          "category",
          "action"
        ],
        "additionalProperties": false,
        "properties": {
          "id": {
            "type": "string",
            "pattern": "^[A-Z0-9_]{1,24}-(?:[0-9a-f]{7}|u[0-9a-f]{6})-REM-[0-9]{3}$",
            "$comment": "harness >= 0.156.0: globally unique, mirroring the /secure-code-audit finding-id convention \u2014 <SLUG_UPPER max 24>-<7-hex short sha of metadata.commit, or u+6-hex sha256(repository) when no commit>-REM-NNN. Pre-0.156.0 bare REM-NNN ids were re-minted corpus-wide."
          },
          "category": {
            "enum": [
              "fix-now",
              "upgrade",
              "waiting-on-upstream",
              "deadline"
            ],
            "$comment": "fix-now = actively prevents PQC negotiation (pins, kill-switches, NO-PQ policy); upgrade = version bump flips PQC on; waiting-on-upstream = blocked-external / delegated dependency; deadline = clock_items-driven NIST IR 8547 obligations"
          },
          "action": {
            "type": "string",
            "$comment": "one imperative sentence \u2014 what to change"
          },
          "details": {
            "type": [
              "string",
              "null"
            ]
          },
          "locations": {
            "type": "array",
            "items": {
              "type": "string"
            },
            "$comment": "file:line anchors from the cited facts"
          },
          "fact_ids": {
            "type": "array",
            "items": {
              "type": "string"
            },
            "$comment": "legacy / optional \u2014 prefer locations for owner-facing remediations; fact ids belong on scores.*.checks[] and clock_items[]. New reports should omit fact_ids here."
          },
          "recipe": {
            "type": [
              "string",
              "null"
            ],
            "$comment": "matching playbook path for the pqc-readiness remediation skill (via index.yaml) for /patch DOMAIN KNOWLEDGE \u2014 machine metadata only; do not render as a user-facing Playbook link"
          },
          "target": {
            "type": [
              "string",
              "null"
            ],
            "$comment": "what to move to: version (Go >= 1.24, JDK >= 24), mechanism (X25519MLKEM768, ML-DSA), or policy value"
          },
          "remediation_effort": {
            "enum": [
              "trivial",
              "moderate",
              "significant",
              "blocked-external",
              "xs",
              "s",
              "m",
              "l",
              "xl",
              null
            ]
          },
          "blocked_external": { "type": "boolean" },
          "blast_radius": {
            "enum": [
              "fleet",
              "product",
              "service",
              "component"
            ]
          },
          "deadline": {
            "type": [
              "integer",
              "null"
            ],
            "$comment": "NIST IR 8547 clock year (2030 deprecated / 2035 disallowed) when category=deadline or the item is clock-bound"
          },
          "blocked_on": {
            "type": [
              "string",
              "null"
            ],
            "$comment": "vendor/library the item waits on, for waiting-on-upstream"
          }
        }
      }
    }
  },
  "definitions": {
    "domain": {
      "type": "object",
      "required": [
        "score",
        "checks"
      ],
      "properties": {
        "score": {
          "type": "number",
          "minimum": 0,
          "maximum": 100
        },
        "checks": {
          "type": "array",
          "items": {
            "type": "object",
            "required": [
              "id",
              "result"
            ],
            "properties": {
              "id": {
                "type": "string"
              },
              "result": {
                "enum": [
                  "yes",
                  "partial",
                  "no",
                  "na"
                ]
              },
              "points": {
                "type": "number"
              },
              "fact_ids": {
                "type": "array",
                "items": {
                  "type": "string"
                }
              },
              "rationale": {
                "type": "string"
              },
              "crypto_governance": {
                "type": "object",
                "properties": {
                  "class": {
                    "enum": [
                      "self",
                      "language-runtime",
                      "platform",
                      "infrastructure"
                    ]
                  },
                  "governor": {
                    "type": "string"
                  },
                  "detection_facts": {
                    "type": "array",
                    "items": {
                      "type": "string"
                    }
                  },
                  "evaluation_target": {
                    "type": "string"
                  }
                }
              }
            },
            "if": {
              "properties": {
                "result": {
                  "enum": [
                    "partial",
                    "no"
                  ]
                }
              }
            },
            "then": {
              "required": [
                "id",
                "result",
                "fact_ids"
              ]
            }
          }
        },
        "assessable_checks": {
          "type": "integer",
          "minimum": 0,
          "$comment": "count of non-na checks the score was computed over (calibration: na-dominated repos need an explicit denominator)"
        }
      }
    }
  },
  "$defs": {
    "canonical_write": {
      "description": "Opt-in canonical effort sizes or null for unknown size; historical clock extras remain readable.",
      "allOf": [
        { "$ref": "#" },
        {
          "properties": {
            "clock_items": {
              "items": {
                "properties": {
                  "remediation_effort": { "$ref": "report.schema.json#/$defs/effort" },
                  "blocked_external": { "type": "boolean" }
                }
              }
            },
            "remediations": {
              "items": {
                "properties": {
                  "remediation_effort": { "$ref": "report.schema.json#/$defs/effort" },
                  "blocked_external": { "type": "boolean" }
                }
              }
            }
          }
        }
      ]
    }
  }
}
`,
	"pqc-vendor-claims": "{\n  \"$schema\": \"https://json-schema.org/draft/2020-12/schema\",\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/pqc-vendor-claims.schema.json\",\n  \"title\": \"PQC vendor claims\",\n  \"description\": \"Curated, dated claims about the post-quantum roadmap of the vendors and dependencies a fleet cannot fix locally (the delegated / blocked-external class). A declared reference document (scope-ref): imported with no assessment and published whole, so a review is a successor version with a higher `version`, never an edit; readers take the latest. Claims are as-reviewed on `last_reviewed` and are re-verified before external use. Fleet counts (repositories depending on a vendor, blocked-external repositories) are derived by the reader from recorded facts and are never stored here.\",\n  \"type\": \"object\",\n  \"required\": [\"artifact\", \"version\", \"updated\", \"vendors\"],\n  \"additionalProperties\": false,\n  \"properties\": {\n    \"artifact\": {\"const\": \"pqc-vendor-claims\"},\n    \"version\": {\"type\": \"integer\", \"minimum\": 1, \"description\": \"Increases with each published successor.\"},\n    \"updated\": {\"type\": \"string\", \"format\": \"date\"},\n    \"note\": {\"type\": \"string\", \"minLength\": 1},\n    \"vendors\": {\"type\": \"array\", \"minItems\": 1, \"items\": {\"$ref\": \"#/$defs/vendor\"}}\n  },\n  \"$defs\": {\n    \"vendor\": {\n      \"type\": \"object\",\n      \"required\": [\"id\", \"name\", \"class\", \"packages\", \"status\", \"claim\", \"sources\", \"last_reviewed\", \"reviewer\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"id\": {\"type\": \"string\", \"pattern\": \"^[a-z0-9][a-z0-9-]*$\", \"description\": \"Unique within the document (the recording service checks).\"},\n        \"name\": {\"type\": \"string\", \"minLength\": 1},\n        \"class\": {\"$ref\": \"#/$defs/vendor_class\"},\n        \"packages\": {\n          \"type\": \"array\",\n          \"items\": {\"$ref\": \"#/$defs/package\"},\n          \"description\": \"Dependencies that carry this vendor's crypto into the fleet, matched against recorded dependency names. Empty: reached only through a provider, toolchain or service, not a package.\"\n        },\n        \"status\": {\"$ref\": \"#/$defs/status\"},\n        \"claim\": {\"type\": \"string\", \"minLength\": 1, \"description\": \"What the vendor ships or has announced, as reviewed.\"},\n        \"watch\": {\"type\": \"string\", \"minLength\": 1, \"description\": \"The signal that would move the status.\"},\n        \"sources\": {\"type\": \"array\", \"minItems\": 1, \"items\": {\"$ref\": \"#/$defs/source\"}},\n        \"last_reviewed\": {\"type\": \"string\", \"format\": \"date\"},\n        \"reviewer\": {\"type\": \"string\", \"minLength\": 1, \"description\": \"Who reviewed the claim (team or role).\"}\n      }\n    },\n    \"vendor_class\": {\n      \"enum\": [\"provider\", \"library\", \"ecosystem\", \"service\", \"internal\"],\n      \"description\": \"provider: a crypto provider or toolchain (OpenSSL, Go, JDK). library: a crypto library packages depend on. ecosystem: a signing or identity ecosystem (Sigstore, GnuPG, Kerberos). service: a hosted service the fleet calls. internal: the organization's own shared services.\"\n    },\n    \"status\": {\n      \"enum\": [\"available\", \"partial\", \"in_progress\", \"none\"],\n      \"description\": \"available: PQC algorithms shipped and usable. partial: some (e.g. key exchange but not signatures). in_progress: announced or in preview. none: no public roadmap.\"\n    },\n    \"package\": {\n      \"type\": \"object\",\n      \"required\": [\"name\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"name\": {\"type\": \"string\", \"minLength\": 1},\n        \"ecosystem\": {\"$ref\": \"impact-analysis.schema.json#/$defs/metadata/properties/ecosystem\"}\n      }\n    },\n    \"source\": {\n      \"type\": \"object\",\n      \"required\": [\"url\", \"retrieved\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"url\": {\"type\": \"string\", \"format\": \"uri\"},\n        \"title\": {\"type\": \"string\", \"minLength\": 1},\n        \"retrieved\": {\"type\": \"string\", \"format\": \"date\"}\n      }\n    }\n  }\n}\n",
	"product-definitions-map": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/product-definitions-map.schema.json",
  "title": "Product Definitions Map",
  "description": "Reviewed findings-package/repo -> ps_product mappings (human-asserted), recorded into traust-core as an imported document; the newest recorded version is the one product lookups use. Permissive by design: the mappings container carries packages and repos maps whose keys and values are deployment-specific; only the top-level shape is validated.",
  "type": "object",
  "required": [
    "mappings"
  ],
  "additionalProperties": true,
  "properties": {
    "mappings": {
      "type": "object",
      "additionalProperties": true,
      "properties": {
        "packages": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "repos": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        }
      }
    }
  }
}
`,
	"refuted-register": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://traust.dev/schemas/v1/refuted-register.schema.json",
  "title": "Refuted finding register",
  "description": "Integrity-bound register of triage or validation refutations retained for later execution evidence and countersign.",
  "type": "object",
  "additionalProperties": false,
  "required": ["source", "generated_at", "entries"],
  "properties": {
    "source": {"type": "string", "minLength": 1},
    "sources": {
      "type": "array",
      "items": {"type": "string", "minLength": 1}
    },
    "generated_at": {"type": "string", "format": "date-time"},
    "entries": {
      "type": "array",
      "items": {"$ref": "#/$defs/entry"}
    }
  },
  "$defs": {
    "entry": {
      "type": "object",
      "additionalProperties": false,
      "required": [
        "finding_ref", "title", "refute_reasons", "tier", "evidence_refs",
        "asserted_at", "asserted_by", "note"
      ],
      "properties": {
        "finding_ref": {"type": "string", "minLength": 1},
        "triage_id": {"type": "string", "minLength": 1},
        "title": {"type": "string", "minLength": 1},
        "source": {"type": "string", "minLength": 1},
        "file": {"type": ["string", "null"], "minLength": 1},
        "line": {"type": ["integer", "null"], "minimum": 0},
        "category": {"type": ["string", "null"]},
        "claimed_severity": {"type": ["string", "null"]},
        "refute_reasons": {
          "type": "array", "minItems": 1,
          "items": {"type": "string", "minLength": 1}
        },
        "exclusion_rule": {"type": ["string", "integer", "null"]},
        "tier": {"type": "string", "minLength": 1},
        "evidence_refs": {
          "type": "array",
          "items": {"type": "string"}
        },
        "asserted_at": {"type": "string", "format": "date-time"},
        "asserted_by": {"type": "string", "minLength": 1},
        "note": {"type": "string"}
      }
    },
    "canonical_write": {
      "description": "Opt-in claimed severity from the existing severity vocabulary; historical text, prose reasons and scalar rules stay readable by default.",
      "allOf": [
        { "$ref": "#" },
        {
          "properties": {
            "entries": {
              "items": {
                "properties": {
                  "claimed_severity": {
                    "anyOf": [
                      { "$ref": "https://example.com/traust-contracts/artifacts_schema/v1/report.schema.json#/$defs/severity_level" },
                      { "type": "null" }
                    ]
                  }
                }
              }
            }
          }
        }
      ]
    }
  }
}
`,
	"remediation": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/remediation.schema.json",
  "title": "Automated Remediation Report Schema",
  "description": "Schema for *-remediation.json reports produced by the remediate-finding harness. Records the patch produced for one or more security findings on a private fork, the local build/test outcome, and (optionally) the live re-validation verdict.",
  "type": "object",
  "required": [
    "title",
    "metadata",
    "source_findings",
    "fork",
    "patch",
    "checks",
    "summary"
  ],
  "additionalProperties": false,
  "properties": {
    "title": {
      "type": "string",
      "minLength": 5,
      "pattern": "(?i)(remediat|patch|fix|mitigat)"
    },
    "metadata": {
      "$ref": "#/$defs/remediation_metadata"
    },
    "source_findings": {
      "description": "Findings this remediation addresses. Usually one; may be several when a single patch closes multiple findings.",
      "type": "array",
      "minItems": 1,
      "items": {
        "$ref": "#/$defs/source_finding"
      }
    },
    "fork": {
      "$ref": "#/$defs/fork"
    },
    "patch": {
      "$ref": "#/$defs/patch"
    },
    "checks": {
      "description": "Local verification results (compile, lint, unit tests). At least one check is required so a remediation report always carries an explicit local verdict.",
      "type": "array",
      "minItems": 1,
      "items": {
        "$ref": "#/$defs/check"
      }
    },
    "evidence": {
      "description": "Typed base-versus-patch evidence. Each item states what kind of claim it makes and what was observed on the unpatched and patched revisions; an item may only claim 'proves' or 'fails_to_prove' if it recorded both observations, so a check that never ran cannot count as evidence. Distinct from 'checks' (local pass/fail with no before/after) and from 'revalidation' (which stays the live-validation channel). What each kind may legitimately conclude is bounded by the per-path evidence ceilings in the harness's docs/disposition-ledger.md \u00a78a.",
      "type": "array",
      "items": {
        "$ref": "#/$defs/patch_evidence"
      }
    },
    "revalidation": {
      "$ref": "#/$defs/revalidation"
    },
    "pull_request": {
      "$ref": "#/$defs/pull_request"
    },
    "summary": {
      "$ref": "#/$defs/summary"
    },
    "notes": {
      "type": "string",
      "description": "Free-form auditor/agent notes \u2014 caveats, follow-ups, anything a human reviewer should know that does not fit a structured field."
    },
    "footer": {
      "type": "string"
    }
  },
  "$defs": {
    "rem_status": {
      "type": "string",
      "enum": [
        "candidate",
        "forked",
        "patched",
        "checks_passed",
        "checks_failed",
        "revalidated_fixed",
        "revalidated_still_vulnerable",
        "pr_opened",
        "merged",
        "abandoned"
      ]
    },
    "fix_strategy": {
      "type": "string",
      "enum": [
        "input-validation",
        "bound-resource",
        "rbac-scope-down",
        "auth-add",
        "tls-enforce",
        "url-allowlist",
        "digest-pin",
        "config-default-harden",
        "api-restrict",
        "remove-feature",
        "dependency-bump",
        "other"
      ]
    },
    "remediation_metadata": {
      "type": "object",
      "required": [
        "date",
        "harness_version",
        "logical_product",
        "repository"
      ],
      "additionalProperties": false,
      "properties": {
        "date": {
          "type": "string",
          "format": "date"
        },
        "harness_version": {
          "type": "string",
          "pattern": "^\\d+\\.\\d+\\.\\d+(-[0-9a-f]{7,40})?$"
        },
        "logical_product": {
          "type": "string",
          "minLength": 1,
          "description": "Logical product / OLM package this repo belongs to (matches validation manifest)."
        },
        "repository": {
          "type": "string",
          "format": "uri",
          "description": "Upstream repository URL the finding was reported against."
        },
        "audited_commit": {
          "type": "string",
          "pattern": "^[0-9a-f]{7,40}$",
          "description": "The exact upstream commit the finding's file:line references resolve to (from triage_context.repo)."
        },
        "language": {
          "type": "string"
        },
        "jira_keys": {
          "type": "array",
          "items": {
            "type": "string",
            "pattern": "^[A-Z][A-Z0-9]+-\\d+$"
          }
        },
        "operator": {
          "type": "string",
          "description": "Agent or human who produced this remediation."
        }
      }
    },
    "source_finding": {
      "type": "object",
      "required": [
        "finding_ref",
        "title",
        "severity",
        "cwes",
        "locations"
      ],
      "additionalProperties": false,
      "properties": {
        "finding_ref": {
          "type": "string",
          "minLength": 3,
          "description": "Stable reference to the source finding, e.g. 'my-service-triage.json#f005' or 'my-service-security-audit.json#FIND-005'."
        },
        "title": {
          "type": "string",
          "minLength": 5
        },
        "severity": {
          "type": "string",
          "enum": [
            "critical",
            "high",
            "medium",
            "low",
            "informational"
          ]
        },
        "cwes": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "string",
            "pattern": "^CWE-\\d{1,5}$"
          }
        },
        "locations": {
          "type": "array",
          "minItems": 1,
          "items": {
            "$ref": "report.schema.json#/$defs/location"
          }
        },
        "triage_confidence": {
          "type": "number",
          "minimum": 0,
          "maximum": 10
        },
        "validation_verdict": {
          "type": "string",
          "enum": [
            "confirmed",
            "refuted",
            "inconclusive",
            "blocked_by_scope",
            "not_attempted",
            "not_validated"
          ]
        },
        "audit_report_path": {
          "type": "string"
        },
        "triage_report_path": {
          "type": "string"
        },
        "validation_report_path": {
          "type": "string"
        }
      }
    },
    "fork": {
      "type": "object",
      "required": [
        "url",
        "visibility",
        "base_ref",
        "fix_branch"
      ],
      "additionalProperties": false,
      "properties": {
        "url": {
          "type": "string",
          "format": "uri",
          "description": "URL of the private fork/mirror where the fix is pushed."
        },
        "host": {
          "type": "string",
          "enum": [
            "github",
            "gitlab",
            "other"
          ],
          "$comment": "Forge KIND, not a deployment. A self-hosted instance of either uses its own kind; anything else is \"other\". The instance itself belongs in fork.url, which is free-form."
        },
        "visibility": {
          "type": "string",
          "enum": [
            "private",
            "internal",
            "public"
          ]
        },
        "upstream_remote": {
          "type": "string",
          "format": "uri",
          "description": "URL of the upstream this fork mirrors (should equal metadata.repository)."
        },
        "base_ref": {
          "type": "string",
          "minLength": 1,
          "description": "Branch or commit in the fork that the fix branch was cut from."
        },
        "base_commit": {
          "type": "string",
          "pattern": "^[0-9a-f]{7,40}$"
        },
        "fix_branch": {
          "type": "string",
          "pattern": "^[A-Za-z0-9._/-]{1,200}$",
          "description": "e.g. security/FIND-005-unbounded-read"
        },
        "fix_commit": {
          "type": "string",
          "pattern": "^[0-9a-f]{7,40}$"
        },
        "synced_from_upstream_at": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "patch": {
      "type": "object",
      "required": [
        "strategy",
        "rationale",
        "files_changed",
        "diffstat"
      ],
      "additionalProperties": false,
      "properties": {
        "strategy": {
          "$ref": "#/$defs/fix_strategy"
        },
        "rationale": {
          "type": "string",
          "minLength": 50,
          "description": "Why this change closes the finding. Must reference the root cause from the triage rationale, not just restate the diff."
        },
        "behaviour_change": {
          "type": "string",
          "description": "User-visible behaviour or API change introduced by the fix, or 'none'."
        },
        "files_changed": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "object",
            "required": [
              "path",
              "change_type"
            ],
            "additionalProperties": false,
            "properties": {
              "path": {
                "type": "string",
                "minLength": 1
              },
              "change_type": {
                "type": "string",
                "enum": [
                  "modified",
                  "added",
                  "deleted",
                  "renamed"
                ]
              },
              "hunks": {
                "type": "integer",
                "minimum": 0
              },
              "additions": {
                "type": "integer",
                "minimum": 0
              },
              "deletions": {
                "type": "integer",
                "minimum": 0
              }
            }
          }
        },
        "diffstat": {
          "type": "object",
          "required": [
            "files",
            "additions",
            "deletions"
          ],
          "additionalProperties": false,
          "properties": {
            "files": {
              "type": "integer",
              "minimum": 1
            },
            "additions": {
              "type": "integer",
              "minimum": 0
            },
            "deletions": {
              "type": "integer",
              "minimum": 0
            }
          }
        },
        "diff_path": {
          "type": "string",
          "description": "Relative path to the unified diff on disk (e.g. patch.diff)."
        },
        "diff_sha256": {
          "type": "string",
          "pattern": "^[a-f0-9]{64}$"
        },
        "tests_added": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Names/paths of regression tests added to prove the fix."
        },
        "residual_risk": {
          "type": "string",
          "description": "Anything the patch deliberately does NOT fix (e.g. 'does not address the SSRF in proxyURL \u2014 separate finding f004')."
        }
      }
    },
    "check": {
      "type": "object",
      "required": [
        "name",
        "command",
        "outcome"
      ],
      "additionalProperties": false,
      "properties": {
        "name": {
          "type": "string",
          "minLength": 1,
          "description": "e.g. 'go build', 'go vet', 'unit-tests', 'golangci-lint'"
        },
        "command": {
          "type": "string",
          "minLength": 1
        },
        "outcome": {
          "type": "string",
          "enum": [
            "pass",
            "fail",
            "skip",
            "error"
          ]
        },
        "duration_seconds": {
          "type": "number",
          "minimum": 0
        },
        "log_path": {
          "type": "string"
        },
        "summary": {
          "type": "string"
        }
      }
    },
    "patch_evidence_kind": {
      "type": "string",
      "enum": [
        "regression",
        "mutation",
        "property",
        "scanner_differential",
        "exploit"
      ]
    },
    "patch_evidence": {
      "type": "object",
      "required": [
        "kind",
        "outcome"
      ],
      "additionalProperties": false,
      "properties": {
        "kind": {
          "$ref": "#/$defs/patch_evidence_kind"
        },
        "outcome": {
          "type": "string",
          "pattern": "^(proves|fails_to_prove|not_attempted: .+)$",
          "description": "'proves' / 'fails_to_prove' require both observations (enforced below). 'not_attempted: <reason>' carries its reason inline, the same shape as deterministic_steps."
        },
        "base_observation": {
          "type": "string",
          "minLength": 1,
          "description": "What was observed on the UNPATCHED revision, e.g. 'test_auth_bypass fails', 'mutant 7 survives at oauth.go:212'."
        },
        "patched_observation": {
          "type": "string",
          "minLength": 1,
          "description": "What was observed on the PATCHED revision, at the same revision-pair and by the same command as base_observation."
        },
        "tool": {
          "type": "string",
          "description": "Tool and version that produced the observations, e.g. 'mewt 4.0.0', 'go test'."
        },
        "command": {
          "type": "string",
          "description": "Command actually executed, so a reader can re-run it."
        },
        "log_path": {
          "type": "string"
        },
        "deterministic_steps": {
          "type": "string",
          "pattern": "^(ran|skipped: .+)$",
          "description": "'ran' or 'skipped: <reason>' \u2014 the harness's existing answer to 'prove it ran'."
        }
      },
      "allOf": [
        {
          "$comment": "A claim of proves/fails_to_prove requires both sides of the comparison. This is the rule that keeps an unexecuted check from being recorded as evidence.",
          "if": {
            "required": [
              "outcome"
            ],
            "properties": {
              "outcome": {
                "pattern": "^(proves|fails_to_prove)$"
              }
            }
          },
          "then": {
            "required": [
              "base_observation",
              "patched_observation"
            ]
          }
        }
      ]
    },
    "revalidation": {
      "type": "object",
      "required": [
        "performed"
      ],
      "additionalProperties": false,
      "properties": {
        "performed": {
          "type": "boolean"
        },
        "method": {
          "type": "string",
          "enum": [
            "validate-live",
            "validate-platform",
            "validate-findings",
            "manual",
            "none"
          ],
          "$comment": "Revalidation method, named generically: validate-live for a live deployed instance, validate-platform for the underlying platform's own components."
        },
        "image_ref": {
          "type": "string",
          "description": "Container image built from the fix branch and deployed for re-validation."
        },
        "validation_report_path": {
          "type": "string"
        },
        "before_verdict": {
          "type": "string",
          "enum": [
            "confirmed",
            "refuted",
            "inconclusive",
            "blocked_by_scope",
            "not_attempted",
            "not_validated"
          ]
        },
        "after_verdict": {
          "type": "string",
          "enum": [
            "confirmed",
            "refuted",
            "inconclusive",
            "blocked_by_scope",
            "not_attempted"
          ]
        },
        "fixed": {
          "type": "boolean",
          "description": "True iff before_verdict was 'confirmed' and after_verdict is 'refuted'."
        }
      }
    },
    "pull_request": {
      "type": "object",
      "required": [
        "url",
        "target"
      ],
      "additionalProperties": false,
      "properties": {
        "url": {
          "type": "string",
          "format": "uri"
        },
        "number": {
          "type": "integer",
          "minimum": 1
        },
        "target": {
          "type": "string",
          "enum": [
            "private-fork",
            "downstream",
            "upstream"
          ],
          "description": "Where the PR is opened. Stage-9 remediation MUST be 'private-fork' or 'downstream'; 'upstream' is reserved for an explicit human-driven disclosure step."
        },
        "base_branch": {
          "type": "string"
        },
        "reviewers": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "state": {
          "type": "string",
          "enum": [
            "draft",
            "open",
            "merged",
            "closed"
          ]
        }
      }
    },
    "summary": {
      "type": "object",
      "required": [
        "status",
        "checks_passed",
        "checks_total"
      ],
      "additionalProperties": false,
      "properties": {
        "status": {
          "$ref": "#/$defs/rem_status"
        },
        "checks_passed": {
          "type": "integer",
          "minimum": 0
        },
        "checks_total": {
          "type": "integer",
          "minimum": 1
        },
        "findings_addressed": {
          "type": "integer",
          "minimum": 1
        },
        "ready_for_review": {
          "type": "boolean"
        }
      }
    }
  }
}
`,
	"repo-liveness": `{
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/repo-liveness.schema.json",
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "title": "Repository liveness reference facts",
  "description": "The latest availability facts recorded by the forge observation owner (traust-core portfolio.liveness), keyed by stable inventory repository UUID. repository_id MUST equal the map key it sits under (validate.py content gate); status_since is retained across consecutive observations of the same status; unknown observations are not evidence of deletion. Unqueried repositories have no entry.",
  "type": "object",
  "required": [
    "repositories",
    "generated"
  ],
  "properties": {
    "generated": {
      "type": "string",
      "format": "date-time",
      "pattern": "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[+-]\\d{2}:\\d{2})$",
      "description": "When the forge observation owner produced this snapshot. A consumer holding an older snapshot replaces it whole: per-repo observed_at orders observations within one snapshot, generated orders snapshots."
    },
    "repositories": {
      "type": "object",
      "propertyNames": {
        "format": "uuid",
        "pattern": "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
      },
      "additionalProperties": {
        "$ref": "#/$defs/liveness"
      }
    }
  },
  "additionalProperties": false,
  "$defs": {
    "liveness": {
      "type": "object",
      "required": [
        "repository_id",
        "url",
        "status",
        "observed_at",
        "status_since"
      ],
      "properties": {
        "repository_id": {
          "type": "string",
          "format": "uuid",
          "pattern": "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
        },
        "url": {
          "type": "string",
          "format": "uri",
          "pattern": "^https://"
        },
        "status": {
          "type": "string",
          "enum": [
            "active",
            "archived",
            "moved",
            "missing",
            "disabled",
            "unknown"
          ]
        },
        "observed_at": {
          "type": "string",
          "format": "date-time",
          "pattern": "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[+-]\\d{2}:\\d{2})$"
        },
        "status_since": {
          "type": "string",
          "format": "date-time",
          "pattern": "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[+-]\\d{2}:\\d{2})$",
          "description": "Start of the current unchanged status. The observation owner carries status_since across consecutive observations of the same status and resets it to observed_at on a status change; a producer that resets it every observation misstates the semantics."
        },
        "canonical_url": {
          "type": "string",
          "format": "uri",
          "pattern": "^https://"
        },
        "archived": {
          "type": "boolean"
        },
        "pushed_at": {
          "type": "string",
          "format": "date-time",
          "pattern": "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[+-]\\d{2}:\\d{2})$"
        },
        "unobserved": {
          "$ref": "repo-state.schema.json#/$defs/comparison_status"
        },
        "reason": {
          "type": "string"
        }
      },
      "additionalProperties": false
    }
  }
}
`,
	"repo-state": `{
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/repo-state.schema.json",
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "title": "Repository State Observation",
  "description": "One forge observation of a repository: its default-branch head, compared with its latest audit when it was pushed since. Stored as the artifact of an assessment of kind repo-state on the observed head ref (baseline_assessment_id = the audit). Recorded only when the forge answered; an unreachable or quota-deferred observation is the run order's outcome, not a document. Repository-level facts (public, upstream, archived) live on the repository row; fork_of keeps the parent the forge reported, the evidence behind repository.upstream_id.",
  "type": "object",
  "required": [
    "observed_at",
    "status"
  ],
  "properties": {
    "observed_at": {
      "type": "string",
      "format": "date-time"
    },
    "status": {
      "$ref": "#/$defs/comparison_status"
    },
    "pushed_at": {
      "type": "string",
      "format": "date-time",
      "description": "Last push the forge reported; drives dormancy and whether a comparison was due."
    },
    "comparison": {
      "$ref": "#/$defs/comparison"
    },
    "fork_of": {
      "type": "string",
      "format": "uri",
      "pattern": "^https://",
      "description": "The repository this one forks, as the forge reported it (https URL); absent when it isn't a fork. The observation that brings an upstream into the inventory."
    }
  },
  "additionalProperties": false,
  "$defs": {
    "comparison_status": {
      "type": "string",
      "enum": [
        "ok",
        "quota-deferred",
        "no-pinned-sha",
        "unsupported-host",
        "no-repo-url",
        "no-credentials",
        "unreachable",
        "error",
        "ban-suspected",
        "network-skipped",
        "never-audited"
      ]
    },
    "comparison": {
      "type": "object",
      "description": "The head compared with the audited commit; absent when the repository was not pushed since the audit. The file counts (first_party_files, iac_files, threat_model_files, tree_first_party_files) and secret_hits are optional: observations recorded before they were measured lack them.",
      "required": [
        "ahead_by",
        "truncated",
        "changed_lines",
        "sensitive_changed",
        "sensitive_lines",
        "deps_only"
      ],
      "properties": {
        "ahead_by": {
          "type": "integer",
          "minimum": 0
        },
        "truncated": {
          "type": "boolean",
          "description": "The forge capped the file list; line counts are a lower bound."
        },
        "changed_lines": {
          "type": "integer",
          "minimum": 0
        },
        "sensitive_changed": {
          "type": "boolean"
        },
        "sensitive_lines": {
          "type": "integer",
          "minimum": 0
        },
        "deps_only": {
          "type": "boolean"
        },
        "first_party_files": {
          "type": "integer",
          "minimum": 0,
          "description": "Changed files that count as first-party (vendored, generated and docs excluded): the numerator of the diff-refusal ratio."
        },
        "iac_files": {
          "type": "integer",
          "minimum": 0,
          "description": "Changed first-party IaC files (Terraform, CloudFormation, ARM/Bicep): code the cloud-config lane assesses and a code audit does not."
        },
        "threat_model_files": {
          "type": "integer",
          "minimum": 0,
          "description": "Changed first-party interface-bearing files (entry-point directories, main files, OpenAPI/Swagger, protobuf): new entry points or trust boundaries a threat model may miss."
        },
        "tree_first_party_files": {
          "type": "integer",
          "minimum": 0,
          "description": "First-party files in the head's tree, counted only when enough first-party files changed for the diff-refusal ratio to matter; absent when not counted (no cheap listing on the forge, or the forge truncated the tree)."
        },
        "secret_hits": {
          "type": "integer",
          "minimum": 0,
          "description": "Secret-scanner (gitleaks) hits in the change's patch text, scanned during the observation: the router's tripwire. Absent when not scanned (no lines changed, no patch, or the scanner unavailable or failed)."
        }
      },
      "additionalProperties": false
    }
  }
}
`,
	"report": "{\n  \"$schema\": \"https://json-schema.org/draft/2020-12/schema\",\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/report.schema.json\",\n  \"title\": \"Security Report Schema\",\n  \"description\": \"Schema for security audit/assessment reports.\",\n  \"type\": \"object\",\n  \"required\": [\n    \"title\",\n    \"metadata\",\n    \"executive_summary\",\n    \"severity_criteria\",\n    \"findings\",\n    \"findings_summary\",\n    \"remediation_roadmap\"\n  ],\n  \"additionalProperties\": false,\n  \"properties\": {\n    \"title\": {\n      \"type\": \"string\",\n      \"minLength\": 5,\n      \"pattern\": \"(?i)(security|audit|assessment|review|finding|analysis)\"\n    },\n    \"metadata\": {\n      \"$ref\": \"#/$defs/metadata\"\n    },\n    \"executive_summary\": {\n      \"$ref\": \"#/$defs/executive_summary\"\n    },\n    \"severity_criteria\": {\n      \"type\": \"array\",\n      \"minItems\": 4,\n      \"items\": {\n        \"$ref\": \"#/$defs/severity_criterion\"\n      }\n    },\n    \"findings\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/finding\"\n      }\n    },\n    \"findings_summary\": {\n      \"type\": \"array\",\n      \"minItems\": 4,\n      \"items\": {\n        \"$ref\": \"#/$defs/severity_count_entry\"\n      }\n    },\n    \"remediation_roadmap\": {\n      \"type\": \"array\",\n      \"minItems\": 1,\n      \"items\": {\n        \"$ref\": \"#/$defs/roadmap_item\"\n      }\n    },\n    \"dependency_audit\": {\n      \"$ref\": \"#/$defs/dependency_audit\"\n    },\n    \"negative_results\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/negative_result\"\n      }\n    },\n    \"asvs_coverage\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/asvs_chapter\"\n      }\n    },\n    \"review_surface_coverage\": {\n      \"type\": \"object\",\n      \"description\": \"Coverage of the review surfaces (traust-contracts enums/v1/review-surface.json): one entry per surface, recording whether the audit answered it with a finding, a negative result, not-applicable, or not assessed. A deterministic record, so coverage is checked by schema, not inferred from prose. llm-integration is allowed but not in required: the report checks require it for audits dated from 2026-11-02, so earlier reports stay valid.\",\n      \"propertyNames\": {\n        \"enum\": [\n          \"credential-acceptance\",\n          \"ownership-resolution\",\n          \"security-flags\",\n          \"outbound-fetch\",\n          \"stored-content-integrity\",\n          \"scope-enforcement\",\n          \"llm-integration\"\n        ]\n      },\n      \"required\": [\n        \"credential-acceptance\",\n        \"ownership-resolution\",\n        \"security-flags\",\n        \"outbound-fetch\",\n        \"stored-content-integrity\",\n        \"scope-enforcement\"\n      ],\n      \"additionalProperties\": {\n        \"$ref\": \"#/$defs/review_surface_entry\"\n      }\n    },\n    \"scanner_correlation\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/scanner_entry\"\n      }\n    },\n    \"peach_isolation_review\": {\n      \"$ref\": \"#/$defs/peach_isolation_review\"\n    },\n    \"disposition_summary\": {\n      \"$ref\": \"#/$defs/disposition_summary\",\n      \"$comment\": \"Present only on cumulative reports (*-findings-current.json) generated by the track-findings skill from a disposition layer (layer.schema.json). Absent on audit-stage reports.\"\n    },\n    \"footer\": {\n      \"type\": \"string\"\n    }\n  },\n  \"$defs\": {\n    \"severity_level\": {\n      \"type\": \"string\",\n      \"enum\": [\n        \"critical\",\n        \"high\",\n        \"medium\",\n        \"low\",\n        \"informational\"\n      ]\n    },\n    \"metadata\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"date\",\n        \"scope\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"date\": {\n          \"type\": \"string\",\n          \"format\": \"date\"\n        },\n        \"scope\": {\n          \"type\": \"string\",\n          \"minLength\": 10\n        },\n        \"repository\": {\n          \"type\": \"string\"\n        },\n        \"commit\": {\n          \"type\": \"string\"\n        },\n        \"ref\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"$comment\": \"Optional (harness >= 0.122.0, branch-awareness Phase 0): the branch or tag name as checked out for this audit (e.g. 'release-4.19', 'main', 'v1.2.3'). Explicit ref provenance \\u2014 when present, preferred over report-slug parsing; the legacy slug-suffix parsing remains the fallback for older corpus data. Never required; all existing reports stay valid.\"\n        },\n        \"ref_kind\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"branch\",\n            \"tag\",\n            \"default\",\n            \"stream\"\n          ],\n          \"$comment\": \"Optional (harness >= 0.122.0): what metadata.ref names \\u2014 'branch' = an explicitly requested non-default branch (counts as a branch re-audit in the census), 'tag' = a tag checkout, 'default' = the repository default branch (metadata.ref still records its name, e.g. 'main'), 'stream' (harness >= 0.140.0) = a dist-git release stream (c10s, c9s) \\u2014 a mainline deliverable audited in its own right, NEVER a branch re-audit. Writers set ref and ref_kind together.\"\n        },\n        \"framework\": {\n          \"type\": \"string\"\n        },\n        \"auditor\": {\n          \"type\": \"string\"\n        },\n        \"methodology\": {\n          \"type\": \"string\"\n        },\n        \"loc_reviewed\": {\n          \"description\": \"Total lines of source code reviewed. Prefer an integer; string is accepted for backward compatibility with earlier reports.\",\n          \"type\": [\n            \"integer\",\n            \"string\"\n          ]\n        },\n        \"loc_breakdown\": {\n          \"description\": \"Structured lines-of-code count for the audited ref, produced by tokei/cloc/scc.\",\n          \"type\": \"object\",\n          \"required\": [\n            \"total\"\n          ],\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"total\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"by_language\": {\n              \"type\": \"object\",\n              \"additionalProperties\": {\n                \"type\": \"integer\",\n                \"minimum\": 0\n              }\n            },\n            \"tool\": {\n              \"type\": \"string\",\n              \"description\": \"Tool used to compute the count (tokei, cloc, scc, wc).\"\n            },\n            \"excludes\": {\n              \"type\": \"array\",\n              \"items\": {\n                \"type\": \"string\"\n              },\n              \"description\": \"Paths or globs excluded from the count (e.g. vendor/, node_modules/).\"\n            }\n          }\n        },\n        \"tools\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"additional\": {\n          \"type\": \"object\",\n          \"description\": \"Producer-declared provenance and run records. harness_version identifies the producing harness release; contracts_version pins the contract surface it was validated against. The other declared fields are the run records the report-producing skills (code, rpm and container profiles) specify. The object stays open: an undeclared key is allowed, but a declared key must match its shape.\",\n          \"properties\": {\n            \"harness_version\": {\n              \"type\": \"string\",\n              \"description\": \"Semver (+ optional -sha) of the producing harness release. Its presence marks the report as modern-producer output, which requires stamped finding identity (see the root allOf rule).\"\n            },\n            \"contracts_version\": {\n              \"type\": \"string\",\n              \"description\": \"Version of traust-contracts this report was validated against.\"\n            },\n            \"deterministic_steps\": {\n              \"type\": \"object\",\n              \"additionalProperties\": {\n                \"type\": \"string\"\n              },\n              \"description\": \"Fate of every deterministic pre-scan, keyed by step name (e.g. opengrep, gitleaks, osv-scanner, sbom-grype, k8s-hardening, priv-profile, fork-advisory-lag, route-guards, config-matrix, yara, crypto-audit, crypto). Value: 'ran', 'skipped: <reason>', or a step-specific state ('list', '<n> records'). A skipped step means that area was reviewed by hand only; consumers must not read missing scanner facts as a clean result.\"\n            },\n            \"precision_gates\": {\n              \"type\": \"object\",\n              \"required\": [\n                \"crit_high_evaluated\",\n                \"fired\"\n              ],\n              \"description\": \"Precision Gate pass record. A report that files critical/high findings without this block is an incomplete run.\",\n              \"properties\": {\n                \"crit_high_evaluated\": {\n                  \"type\": \"integer\",\n                  \"minimum\": 0,\n                  \"description\": \"Every critical/high candidate evaluated, filed or gated.\"\n                },\n                \"fired\": {\n                  \"type\": \"array\",\n                  \"description\": \"Each gate that changed a candidate's disposition. An empty list is legitimate.\",\n                  \"items\": {\n                    \"type\": \"object\",\n                    \"required\": [\n                      \"candidate\",\n                      \"gate\",\n                      \"action\"\n                    ],\n                    \"properties\": {\n                      \"candidate\": {\n                        \"type\": \"string\",\n                        \"description\": \"Short title or finding id.\"\n                      },\n                      \"gate\": {\n                        \"type\": \"string\",\n                        \"description\": \"Gate name from the skill's Precision Gate section.\"\n                      },\n                      \"action\": {\n                        \"type\": \"string\",\n                        \"enum\": [\n                          \"downgraded\",\n                          \"negative_results\",\n                          \"dependency_audit\"\n                        ]\n                      }\n                    }\n                  }\n                }\n              }\n            },\n            \"coverage_gaps\": {\n              \"type\": \"array\",\n              \"description\": \"Systems a pre-scan or review lane could not examine (unexpanded config systems, unenumerated frameworks, classes a scoped audit did not cover). Aggregated fleet-wide: an audit that skipped something says so here instead of reading as clean.\",\n              \"items\": {\n                \"type\": \"object\",\n                \"required\": [\n                  \"tool\",\n                  \"system\",\n                  \"reason\"\n                ],\n                \"properties\": {\n                  \"tool\": {\n                    \"type\": \"string\",\n                    \"description\": \"Pre-scan or lane that recorded the gap.\"\n                  },\n                  \"system\": {\n                    \"type\": \"string\",\n                    \"description\": \"What was not examined.\"\n                  },\n                  \"reason\": {\n                    \"type\": \"string\"\n                  },\n                  \"count\": {\n                    \"type\": \"integer\",\n                    \"minimum\": 0,\n                    \"description\": \"Number of files or items affected, when known.\"\n                  }\n                }\n              }\n            },\n            \"assertive_inference\": {\n              \"type\": \"object\",\n              \"required\": [\n                \"languages\",\n                \"trigger\"\n              ],\n              \"description\": \"Assertive reporting posture for languages with thin or no deterministic rule coverage. Changes the report-or-omit decision only, never severity or validation_status.\",\n              \"properties\": {\n                \"languages\": {\n                  \"type\": \"array\",\n                  \"items\": {\n                    \"type\": \"string\"\n                  }\n                },\n                \"trigger\": {\n                  \"type\": \"string\",\n                  \"description\": \"Why the posture applied, e.g. 'thin per opengrep coverage block'.\"\n                }\n              }\n            },\n            \"companion_lanes\": {\n              \"type\": \"array\",\n              \"items\": {\n                \"type\": \"string\"\n              },\n              \"description\": \"Other scan lanes the tree needs that this audit does not cover (e.g. 'cloud-config-audit' when IaC is present). A declaration only; the router turns it into worklist rows.\"\n            },\n            \"audit_passes\": {\n              \"type\": \"integer\",\n              \"minimum\": 1,\n              \"description\": \"Number of independent review passes merged into the report. Recorded when it differs from the two-pass default.\"\n            },\n            \"model_routing\": {\n              \"type\": \"object\",\n              \"required\": [\n                \"role\",\n                \"model\"\n              ],\n              \"description\": \"The routed model decision that produced the report.\",\n              \"properties\": {\n                \"role\": {\n                  \"type\": \"string\"\n                },\n                \"model\": {\n                  \"type\": \"string\"\n                },\n                \"registry_sha\": {\n                  \"type\": \"string\",\n                  \"description\": \"Short sha of the model registry revision that made the decision.\"\n                },\n                \"floor\": {\n                  \"type\": \"string\",\n                  \"description\": \"Model-class floor that applied.\"\n                }\n              }\n            },\n            \"source_audit\": {\n              \"type\": \"string\",\n              \"description\": \"Container profile: path of the source repository's code-profile report that this image audit cross-links.\"\n            },\n            \"degraded\": {\n              \"type\": \"boolean\",\n              \"description\": \"True when the run lost coverage (a lane or tool unavailable). A degraded report must not be read as full coverage.\"\n            }\n          }\n        },\n        \"audit_profile\": {\n          \"enum\": [\n            \"code\",\n            \"rpm\",\n            \"container\"\n          ],\n          \"description\": \"Report profile discriminator: 'code' = secure-code-audit (source/manifest audit), 'rpm' = secure-rpm-audit (dist-git packaging audit), 'container' = secure-container-audit (registry image audit via skopeo/syft/grype). Shared finding/severity/ID vocabulary; framework sections differ by profile.\"\n        }\n      }\n    },\n    \"executive_summary\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"prose\",\n        \"severity_counts\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"prose\": {\n          \"type\": \"string\",\n          \"minLength\": 50\n        },\n        \"severity_counts\": {\n          \"type\": \"object\",\n          \"required\": [\n            \"critical\",\n            \"high\",\n            \"medium\",\n            \"low\",\n            \"informational\"\n          ],\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"critical\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"high\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"medium\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"low\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"informational\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            }\n          }\n        },\n        \"key_risks\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"positive_observations\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        }\n      }\n    },\n    \"severity_criterion\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"level\",\n        \"definition\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"level\": {\n          \"$ref\": \"#/$defs/severity_level\"\n        },\n        \"definition\": {\n          \"type\": \"string\",\n          \"minLength\": 20\n        },\n        \"cvss_range\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        }\n      }\n    },\n    \"location\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"path\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"path\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Repository-relative path of the artifact this finding concerns, including an artifact that SHOULD exist and does not (an absent SECURITY.md is still `SECURITY.md`). Never a repo-root marker: `.`, `/`, `./` and `/./` all canonicalize to empty, which collapses the ledger fingerprint to (repo, '', cwe). When a finding genuinely concerns no artifact (e.g. 'repository appears unmaintained'), use a controlled pseudo-path from enums/v1/repo-scope-path.json. Not a notes field: prose in this field silently feeds identity.\",\n          \"$comment\": \"Enforcement is staged, following the P0.4 -> P6 precedent (backfill, then flip): the rule is declared here and checked by the harness gate at warning level while existing corpus findings are migrated; the pattern below and the identity module strict mode become errors once that lands. Do NOT add the pattern before the migration \\u2014 it would invalidate every existing report. Proposed: \\\"pattern\\\": \\\"^(?!\\\\\\\\.?/?$)(repo:[a-z-]+|[^\\\\\\\\n]{1,200})$\\\".\"\n        },\n        \"lines\": {\n          \"type\": \"string\"\n        },\n        \"description\": {\n          \"type\": \"string\"\n        }\n      }\n    },\n    \"evidence_block\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"code\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"code\": {\n          \"type\": \"string\"\n        },\n        \"language\": {\n          \"type\": \"string\"\n        },\n        \"caption\": {\n          \"type\": \"string\"\n        }\n      }\n    },\n    \"finding\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"id\",\n        \"title\",\n        \"severity\",\n        \"cwes\",\n        \"locations\",\n        \"description\",\n        \"remediation\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"id\": {\n          \"$comment\": \"Canonical (harness >= 0.12.0): {REPO_SLUG}-{SHORTSHA}-{NNN} \\u2014 REPO_SLUG = repo name uppercased, [^A-Z0-9] -> '_', truncated to 24 chars; SHORTSHA = first 7 lowercase hex chars of the audited commit; NNN = 001-based sequence. The second alternative preserves validation of pre-0.12.0 reports; the validator hard-errors on it when harness_version >= 0.12.0.\",\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"maxLength\": 50,\n          \"pattern\": \"^([A-Z][A-Z0-9_]{0,23}-[a-f0-9]{7}-\\\\d{3}|[A-Za-z0-9][A-Za-z0-9 _.\\\\-/]{0,49})$\"\n        },\n        \"title\": {\n          \"type\": \"string\",\n          \"minLength\": 5\n        },\n        \"severity\": {\n          \"$ref\": \"#/$defs/severity_level\"\n        },\n        \"cwes\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"items\": {\n            \"type\": \"string\",\n            \"pattern\": \"^CWE-\\\\d{1,5}$\"\n          }\n        },\n        \"locations\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"items\": {\n            \"$ref\": \"#/$defs/location\"\n          }\n        },\n        \"description\": {\n          \"type\": \"string\",\n          \"minLength\": 50\n        },\n        \"remediation\": {\n          \"type\": \"string\",\n          \"minLength\": 10\n        },\n        \"asvs_references\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"owasp_llm_references\": {\n          \"description\": \"OWASP Top 10 for LLM Applications (2025) category IDs this finding falls under. Set on findings about LLM, agent or MCP application code; the category ID is cited, never OWASP text.\",\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"uniqueItems\": true,\n          \"items\": {\n            \"enum\": [\"LLM01\", \"LLM02\", \"LLM03\", \"LLM04\", \"LLM05\", \"LLM06\", \"LLM07\", \"LLM08\", \"LLM09\", \"LLM10\"]\n          }\n        },\n        \"peach_references\": {\n          \"description\": \"PEACH tenant-isolation hardening parameter(s) this finding violates. Only set when the audited component is multi-tenant.\",\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"items\": {\n            \"type\": \"string\",\n            \"pattern\": \"^PEACH-[PEACH]$\"\n          }\n        },\n        \"cvss\": {\n          \"type\": \"object\",\n          \"required\": [\n            \"score\",\n            \"vector\"\n          ],\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"score\": {\n              \"type\": \"number\",\n              \"minimum\": 0.0,\n              \"maximum\": 10.0\n            },\n            \"vector\": {\n              \"type\": \"string\"\n            }\n          }\n        },\n        \"capec\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\",\n            \"pattern\": \"^CAPEC-\\\\d+$\"\n          }\n        },\n        \"evidence\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/$defs/evidence_block\"\n          }\n        },\n        \"attack_pattern\": {\n          \"type\": \"string\",\n          \"$comment\": \"Concrete attack scenario. Expected on every critical/high finding (validator strict mode warns when absent).\"\n        },\n        \"category\": {\n          \"type\": \"string\",\n          \"$comment\": \"Recommended vocabulary (kebab-case; validator strict mode warns off-vocabulary for harness >= 0.15.0): injection, authentication, authorization, secrets-management, supply-chain, insecure-workload-config, network-exposure, cryptography, input-validation, path-traversal, cross-site-scripting, ssrf, resource-management, logging-monitoring, data-exposure, tenant-isolation.\"\n        },\n        \"kind\": {\n          \"enum\": [\"vulnerability\", \"weakness\"],\n          \"description\": \"The finding kind, independent of validity. Historical findings may omit it.\"\n        },\n        \"validation_status\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"confirmed\",\n            \"corrected\",\n            \"false_positive\",\n            \"not_verified\",\n            \"hardening\",\n            \"not_affected\",\n            \"true_positive\"\n          ],\n          \"$comment\": \"Historical validity spellings remain accepted by this base reader schema. Newly authored unjudged findings omit this optional field; null is not a validity value. New report admission uses the installed registry's canonical validity members and does not infer a judgment from static review, kind, actor, or source. Historical not_verified denotes no answer; corrected is a revision marker, not a validity answer; hardening is represented independently by finding kind. Core finding decisions and authenticated restatement operations own subsequent state changes. Existing source artifacts, stored bytes, and exact-run replay remain compatible.\"\n        },\n        \"source_findings\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"origin\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"verify-remediation\",\n            \"vuln-scan\",\n            \"create-fuzzing\",\n            \"validate-findings\",\n            \"validation-discovery\",\n            \"impact-analysis\"\n          ],\n          \"$comment\": \"Historical/provider provenance retained for old sanctioned-append findings. Current Core source attribution comes from the assessment that first reports a finding, not a candidate origin override or a new baseline append. Unjudged new findings do not acquire a validity answer from their origin; omit optional validation_status where the source shape supports it. Source schemas that do not declare validation_status must not add it. Historical origin fields and source artifacts remain readable.\"\n        },\n        \"pqc_classification\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"shor-key-establishment\",\n            \"shor-signature\",\n            \"clock-2030-parameter\",\n            \"classically-broken\",\n            \"hndl-exposure\",\n            \"pqc-blocker-config\",\n            \"pqc-adoption\"\n          ],\n          \"$comment\": \"Optional (harness >= 0.84.0): PQC relevance of this finding, assigned by table lookup against the PQC reference tables (ir8547-mapping, pqc-readiness-decision-tree) \\u2014 never free-authored. Only set on PQC-relevant findings; absent elsewhere.\"\n        },\n        \"remediation_effort\": {\n          \"type\": [\"string\", \"null\"],\n          \"enum\": [\n            \"trivial\",\n            \"moderate\",\n            \"significant\",\n            \"blocked-external\",\n            \"xs\",\n            \"s\",\n            \"m\",\n            \"l\",\n            \"xl\",\n            null\n          ],\n          \"$comment\": \"Legacy effort labels remain valid for historical artifacts. Sizes xs–xl are added; canonical_write rejects legacy labels without inventing a size mapping. Blocking is independent.\"\n        },\n        \"blocked_external\": {\n          \"type\": \"boolean\",\n          \"description\": \"An external dependency blocks this item; independent of effort size.\"\n        },\n        \"isolation_dimensions\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"uniqueItems\": true,\n          \"items\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"privilege\",\n              \"encryption\",\n              \"authentication\",\n              \"connectivity\",\n              \"hygiene\"\n            ]\n          },\n          \"$comment\": \"Optional (harness >= 0.120.0): isolation-hardening dimension(s) this finding stresses, vocabulary shared with schema/isolation-review.schema.json. Only set when the finding stresses a tenant boundary in a multi-tenant service; absent elsewhere. PEACH is cited by name/URL only \\u2014 never adapted text.\"\n        },\n        \"isolation_boundary\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"$comment\": \"Optional (harness >= 0.120.0): free-form identifier of the tenant-facing interface this finding sits on \\u2014 an interface name from peach_isolation_review.interfaces[] or an IF-n id from a service isolation review. Only meaningful alongside isolation_dimensions.\"\n        },\n        \"disposition\": {\n          \"$ref\": \"#/$defs/disposition\",\n          \"$comment\": \"Present only on cumulative reports generated by the track-findings skill. disposition.validity must equal validation_status (validator-enforced).\"\n        },\n        \"fingerprint\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{64}$\",\n          \"description\": \"Deterministic cross-scan identity: sha256(canonical repo URL | ';'-joined sorted set of lineless location paths | primary CWE). One implementation, the identity module fingerprint function; stamped by the finding-identity utility, never model-authored. Only the harness computes identity \\u2014 a non-harness producer submits unstamped rather than inventing a value. Validation recomputes every stamp and compares (check_finding_identity), so a forged or drifted value is an error rather than a silent identity split. Secondary correlation key \\u2014 the scan-scoped id stays the primary key, and a fingerprint is never the sole key of a disposition record.\"\n        },\n        \"fingerprint_algo\": {\n          \"type\": \"string\",\n          \"pattern\": \"^v\\\\d+$\",\n          \"description\": \"Algorithm version of `fingerprint` (the identity module ALGO_VERSION). Mirrors layer.schema.json's event field of the same name, and for the same reason: stamps outlive recipes, so matching fingerprints across an epoch boundary means comparing this field too. Absent on findings stamped before 2026-09-18 -- those predate the field rather than declaring an unknown version, and were separable only because exactly two recipes existed to try.\"\n        },\n        \"effective_severity\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"critical\",\n            \"high\",\n            \"medium\",\n            \"low\",\n            \"informational\"\n          ],\n          \"$comment\": \"Present only on cumulative reports (harness >= 0.128.0): equals disposition.severity_override.severity when a human override exists, else the original severity. Original severity is never rewritten - severity-count cross-checks keep validating against the original field.\"\n        },\n        \"passes\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"integer\",\n            \"minimum\": 1\n          },\n          \"description\": \"Audit passes that independently surfaced this finding (dual-pass union-merge, harness >= 0.144.0). [1,2] = both passes; [2] = second pass only. Optional; single-pass reports omit it.\"\n        },\n        \"dependency\": {\n          \"type\": \"object\",\n          \"description\": \"Supply-chain provenance for a dependency finding, so the finding is SELF-DESCRIBING and does not depend on a sibling impact-analysis artifact resolving by relative path. Required once findings live in object storage, and the join key for consumers that build their own database rather than reading the local SQLite projection. Optional; present when origin is 'impact-analysis' or the finding is otherwise advisory-derived.\",\n          \"additionalProperties\": false,\n          \"required\": [\n            \"advisory\",\n            \"module\"\n          ],\n          \"properties\": {\n            \"advisory\": {\n              \"type\": \"string\",\n              \"description\": \"Advisory id as published \\u2014 CVE-, GHSA-, RHSA-, MAL-, PYSEC-, GO-, RUSTSEC-, OSV-.\"\n            },\n            \"module\": {\n              \"type\": \"string\",\n              \"description\": \"Package identity in its ecosystem's native form (org.postgresql:postgresql, golang.org/x/crypto, lodash).\"\n            },\n            \"ecosystem\": {\n              \"type\": \"string\",\n              \"enum\": [\n                \"go\",\n                \"npm\",\n                \"pypi\",\n                \"maven\",\n                \"cargo\",\n                \"ruby\",\n                \"nuget\",\n                \"actions\",\n                \"docker\",\n                \"helm\",\n                \"rpm\",\n                \"deb\",\n                \"apk\"\n              ],\n              \"description\": \"Package ecosystem; the same vocabulary as impact-analysis metadata.ecosystem (enums/v1/ecosystem.json), including the distro ecosystems rpm/deb/apk for C/C++.\"\n            },\n            \"purl\": {\n              \"type\": \"string\",\n              \"description\": \"Package URL when available \\u2014 the portable identity for cross-system joins.\"\n            },\n            \"vulnerable_range\": {\n              \"type\": \"string\"\n            },\n            \"fixed_version\": {\n              \"type\": \"string\"\n            },\n            \"installed_version\": {\n              \"type\": \"string\",\n              \"description\": \"Version observed in the manifest or SBOM that matched the range.\"\n            },\n            \"evidence_level\": {\n              \"type\": \"string\",\n              \"enum\": [\n                \"symbol\",\n                \"symbol-usage\",\n                \"binary\",\n                \"manifest\",\n                \"none\"\n              ],\n              \"description\": \"Tier that justified the classification; 'manifest' and 'symbol-usage' both ceiling at likely_affected.\"\n            },\n            \"classification\": {\n              \"type\": \"string\",\n              \"enum\": [\n                \"affected\",\n                \"likely_affected\",\n                \"not_observed\",\n                \"version_not_in_range\",\n                \"not_imported\",\n                \"inconclusive\"\n              ]\n            },\n            \"impact_artifact\": {\n              \"type\": \"string\",\n              \"description\": \"Reference to the impact-analysis artifact this was routed from. A provenance pointer, NOT the source of truth \\u2014 the fields above stand alone if it cannot be resolved.\"\n            }\n          }\n        }\n      }\n    },\n    \"disposition\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"validity\",\n        \"resolution\",\n        \"last_updated\",\n        \"events\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"validity\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"confirmed\",\n            \"corrected\",\n            \"false_positive\",\n            \"not_verified\",\n            \"hardening\",\n            \"not_affected\",\n            \"true_positive\"\n          ]\n        },\n        \"resolution\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"open\",\n            \"fix_in_progress\",\n            \"resolved\",\n            \"partially_resolved\",\n            \"risk_accepted\",\n            \"regression_introduced\"\n          ]\n        },\n        \"conflict\": {\n          \"type\": \"boolean\",\n          \"description\": \"True when the layer contains both 'confirmed' and 'false_positive' validity events for this finding \\u2014 surfaced for human re-review, never silently resolved.\"\n        },\n        \"refuted_awaiting_signoff\": {\n          \"type\": \"boolean\",\n          \"description\": \"True when machine validation refuted the finding but no identity-provider-verified human has countersigned the false_positive determination yet.\"\n        },\n        \"assurance\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"claimed\",\n            \"machine_verified\",\n            \"human_reviewed\",\n            \"execution_proven\"\n          ],\n          \"description\": \"Highest evidence class that has spoken on this finding's validity: execution_proven (reproducing exploit/crash), human_reviewed (identity-provider-verified human), machine_verified (triage/audit machine verdict), claimed (audit only).\"\n        },\n        \"evidence_grade\": {\n          \"$ref\": \"layer.schema.json#/$defs/event/properties/evidence_grade\"\n        },\n        \"fp_overridden\": {\n          \"type\": \"boolean\",\n          \"description\": \"True when execution-verified evidence confirmed the finding over a prior human false_positive assertion. Surfaced loudly; the ledger preserves both events with attribution.\"\n        },\n        \"fp_reassertion_blocked\": {\n          \"type\": \"boolean\",\n          \"description\": \"True when a single human false_positive assertion post-dates execution-verified confirmation \\u2014 re-asserting FP against a reproducing PoC requires two independent identity-provider-verified humans (two-person rule).\"\n        },\n        \"last_updated\": {\n          \"type\": \"string\",\n          \"format\": \"date-time\"\n        },\n        \"events\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\",\n            \"pattern\": \"^[a-f0-9]{64}$\"\n          },\n          \"description\": \"event_id values from the disposition layer, chronological.\"\n        },\n        \"severity_override\": {\n          \"type\": \"object\",\n          \"additionalProperties\": false,\n          \"required\": [\n            \"severity\",\n            \"by\",\n            \"at\"\n          ],\n          \"properties\": {\n            \"severity\": {\n              \"type\": \"string\",\n              \"enum\": [\n                \"critical\",\n                \"high\",\n                \"medium\",\n                \"low\",\n                \"informational\"\n              ]\n            },\n            \"by\": {\n              \"type\": \"string\",\n              \"description\": \"identity handle of the verified signer\"\n            },\n            \"at\": {\n              \"type\": \"string\",\n              \"description\": \"when the override was made (occurred_at of the latest human severity event)\"\n            },\n            \"rationale\": {\n              \"type\": \"string\"\n            }\n          },\n          \"description\": \"Present only on cumulative reports when an identity-provider-verified human recorded a severity override in the disposition layer (latest human severity event wins). The finding's original severity field is preserved; effective_severity carries the override.\"\n        }\n      }\n    },\n    \"disposition_summary\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"layer_ref\",\n        \"generated_at\",\n        \"by_resolution\",\n        \"by_validity\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"layer_ref\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Relative path to the *-findings-layer.json this report was generated from.\"\n        },\n        \"generated_at\": {\n          \"type\": \"string\",\n          \"format\": \"date-time\"\n        },\n        \"by_resolution\": {\n          \"type\": \"object\",\n          \"required\": [\n            \"open\",\n            \"fix_in_progress\",\n            \"resolved\",\n            \"partially_resolved\",\n            \"risk_accepted\",\n            \"regression_introduced\"\n          ],\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"open\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"fix_in_progress\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"resolved\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"partially_resolved\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"risk_accepted\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"regression_introduced\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            }\n          }\n        },\n        \"by_validity\": {\n          \"type\": \"object\",\n          \"required\": [\n            \"confirmed\",\n            \"corrected\",\n            \"false_positive\",\n            \"not_verified\"\n          ],\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"confirmed\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"corrected\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"false_positive\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"not_verified\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"not_affected\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"true_positive\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"hardening\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            }\n          }\n        },\n        \"conflicts\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          },\n          \"description\": \"Finding IDs whose disposition.conflict is true.\"\n        },\n        \"needs_review_count\": {\n          \"type\": \"integer\",\n          \"minimum\": 0,\n          \"description\": \"Count of pending items in the layer's needs_review queue.\"\n        },\n        \"severity_overrides\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"object\",\n            \"additionalProperties\": false,\n            \"required\": [\n              \"finding\",\n              \"severity\",\n              \"by\",\n              \"at\"\n            ],\n            \"properties\": {\n              \"finding\": {\n                \"type\": \"string\"\n              },\n              \"from\": {\n                \"type\": [\n                  \"string\",\n                  \"null\"\n                ],\n                \"description\": \"the audit report's original severity\"\n              },\n              \"severity\": {\n                \"type\": \"string\",\n                \"enum\": [\n                  \"critical\",\n                  \"high\",\n                  \"medium\",\n                  \"low\",\n                  \"informational\"\n                ]\n              },\n              \"by\": {\n                \"type\": \"string\"\n              },\n              \"at\": {\n                \"type\": \"string\"\n              },\n              \"rationale\": {\n                \"type\": \"string\"\n              }\n            }\n          },\n          \"description\": \"Every finding whose severity an identity-provider-verified human overrode via the countersign workbench (harness >= 0.128.0). Original severities stay in place; effective_severity carries the override per finding.\"\n        }\n      }\n    },\n    \"severity_count_entry\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"severity\",\n        \"count\",\n        \"finding_ids\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"severity\": {\n          \"$ref\": \"#/$defs/severity_level\"\n        },\n        \"count\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"finding_ids\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        }\n      }\n    },\n    \"roadmap_item\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"priority\",\n        \"action\",\n        \"addresses\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"priority\": {\n          \"type\": \"string\"\n        },\n        \"action\": {\n          \"type\": \"string\",\n          \"minLength\": 10\n        },\n        \"addresses\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"effort\": {\n          \"type\": [\"string\", \"null\"]\n        },\n        \"blocked_external\": {\n          \"type\": \"boolean\"\n        }\n      }\n    },\n    \"dependency_audit\": {\n      \"type\": \"object\",\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"prose\": {\n          \"type\": \"string\"\n        },\n        \"entries\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"object\",\n            \"required\": [\n              \"package\",\n              \"version\",\n              \"status\"\n            ],\n            \"additionalProperties\": false,\n            \"properties\": {\n              \"package\": {\n                \"type\": \"string\"\n              },\n              \"version\": {\n                \"type\": \"string\"\n              },\n              \"status\": {\n                \"type\": \"string\"\n              },\n              \"notes\": {\n                \"type\": \"string\"\n              }\n            }\n          }\n        }\n      }\n    },\n    \"negative_result\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"area\",\n        \"result\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"area\": {\n          \"type\": \"string\"\n        },\n        \"files\": {\n          \"type\": \"string\"\n        },\n        \"result\": {\n          \"type\": \"string\"\n        }\n      }\n    },\n    \"review_surface_entry\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"status\"\n      ],\n      \"additionalProperties\": false,\n      \"description\": \"One review surface's outcome. finding: finding_ids name the report findings on it. negative: a negative_results entry whose area is the surface id records what was examined. not_applicable / not_assessed / out_of_scope: note says why; a not_assessed surface also belongs in metadata.additional.coverage_gaps. out_of_scope is for a diff-scoped scan whose diff does not touch the surface.\",\n      \"properties\": {\n        \"status\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"finding\",\n            \"negative\",\n            \"not_applicable\",\n            \"not_assessed\",\n            \"out_of_scope\"\n          ]\n        },\n        \"finding_ids\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          },\n          \"uniqueItems\": true\n        },\n        \"note\": {\n          \"type\": \"string\"\n        }\n      },\n      \"allOf\": [\n        {\n          \"if\": {\n            \"properties\": {\n              \"status\": {\n                \"const\": \"finding\"\n              }\n            },\n            \"required\": [\n              \"status\"\n            ]\n          },\n          \"then\": {\n            \"required\": [\n              \"finding_ids\"\n            ],\n            \"properties\": {\n              \"finding_ids\": {\n                \"minItems\": 1\n              }\n            }\n          }\n        },\n        {\n          \"if\": {\n            \"properties\": {\n              \"status\": {\n                \"enum\": [\n                  \"not_applicable\",\n                  \"not_assessed\",\n                  \"out_of_scope\"\n                ]\n              }\n            },\n            \"required\": [\n              \"status\"\n            ]\n          },\n          \"then\": {\n            \"required\": [\n              \"note\"\n            ],\n            \"properties\": {\n              \"note\": {\n                \"minLength\": 1\n              }\n            }\n          }\n        }\n      ]\n    },\n    \"scan_review_surface_coverage\": {\n      \"type\": \"object\",\n      \"description\": \"Review-surface coverage recorded by a scan that follows an audit (vuln-findings.schema.json, verification.schema.json). Same keys and statuses as the report's review_surface_coverage, but a scan has no negative_results, so a negative entry carries its evidence in note. The cumulative builder merges these records into the cumulative report.\",\n      \"propertyNames\": {\n        \"enum\": [\n          \"credential-acceptance\",\n          \"ownership-resolution\",\n          \"security-flags\",\n          \"outbound-fetch\",\n          \"stored-content-integrity\",\n          \"scope-enforcement\",\n          \"llm-integration\"\n        ]\n      },\n      \"required\": [\n        \"credential-acceptance\",\n        \"ownership-resolution\",\n        \"security-flags\",\n        \"outbound-fetch\",\n        \"stored-content-integrity\",\n        \"scope-enforcement\"\n      ],\n      \"additionalProperties\": {\n        \"$ref\": \"#/$defs/scan_review_surface_entry\"\n      }\n    },\n    \"scan_review_surface_entry\": {\n      \"description\": \"A review_surface_entry whose negative status also needs a note saying what was examined.\",\n      \"allOf\": [\n        {\n          \"$ref\": \"#/$defs/review_surface_entry\"\n        },\n        {\n          \"if\": {\n            \"properties\": {\n              \"status\": {\n                \"const\": \"negative\"\n              }\n            },\n            \"required\": [\n              \"status\"\n            ]\n          },\n          \"then\": {\n            \"required\": [\n              \"note\"\n            ],\n            \"properties\": {\n              \"note\": {\n                \"minLength\": 1\n              }\n            }\n          }\n        }\n      ]\n    },\n    \"asvs_chapter\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"chapter\",\n        \"area\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"chapter\": {\n          \"type\": \"string\"\n        },\n        \"area\": {\n          \"type\": \"string\"\n        },\n        \"requirements_assessed\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"violations\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"highest_severity\": {\n          \"$ref\": \"#/$defs/severity_level\"\n        }\n      }\n    },\n    \"peach_isolation_review\": {\n      \"description\": \"PEACH tenant-isolation review. Records multi-tenancy applicability and, when applicable, one entry per customer-facing interface with its complexity, sharing model, boundary type, and P.E.A.C.H. hardening gaps.\",\n      \"$comment\": \"Required top-level section for harness >= 0.15.0 (validator-enforced, mirroring the canonical-ID gate): emit {applicable: false, rationale: <why>} for single-tenant components.\",\n      \"type\": \"object\",\n      \"required\": [\n        \"applicable\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"applicable\": {\n          \"type\": \"boolean\",\n          \"description\": \"True if one deployment of this component serves more than one tenant/customer/trust domain.\"\n        },\n        \"rationale\": {\n          \"type\": \"string\",\n          \"description\": \"Why the component is or is not multi-tenant.\"\n        },\n        \"interfaces\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"object\",\n            \"required\": [\n              \"name\",\n              \"complexity\",\n              \"shared\",\n              \"boundary_type\"\n            ],\n            \"additionalProperties\": false,\n            \"properties\": {\n              \"name\": {\n                \"type\": \"string\"\n              },\n              \"complexity\": {\n                \"type\": \"string\",\n                \"enum\": [\n                  \"high\",\n                  \"medium\",\n                  \"low\"\n                ]\n              },\n              \"shared\": {\n                \"type\": \"boolean\",\n                \"description\": \"True if all tenants hit the same running instance; false if duplicated per tenant.\"\n              },\n              \"boundary_type\": {\n                \"type\": \"string\",\n                \"enum\": [\n                  \"hardware_separation\",\n                  \"hardware_virtualization\",\n                  \"containerization\",\n                  \"data_segmentation\",\n                  \"network_segmentation\",\n                  \"identity_segmentation\"\n                ]\n              },\n              \"hardening_gaps\": {\n                \"type\": \"array\",\n                \"items\": {\n                  \"type\": \"string\",\n                  \"pattern\": \"^PEACH-[PEACH]$\"\n                }\n              },\n              \"finding_ids\": {\n                \"type\": \"array\",\n                \"items\": {\n                  \"type\": \"string\"\n                }\n              },\n              \"notes\": {\n                \"type\": \"string\"\n              }\n            }\n          }\n        }\n      }\n    },\n    \"scanner_entry\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"tool\",\n        \"result\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"tool\": {\n          \"type\": \"string\"\n        },\n        \"configured\": {\n          \"type\": \"boolean\"\n        },\n        \"result\": {\n          \"type\": \"string\",\n          \"description\": \"Free text; for per-fact judge decisions on deterministic pre-scan facts use the tokens 'promoted' or 'dismissed' so the precision tracker can compute per-rule precision.\"\n        },\n        \"notes\": {\n          \"type\": \"string\"\n        },\n        \"rule_id\": {\n          \"type\": \"string\",\n          \"description\": \"Scanner rule that produced the judged fact (e.g. an opengrep rule id).\"\n        },\n        \"location\": {\n          \"type\": \"string\",\n          \"description\": \"file:line of the judged fact.\"\n        },\n        \"finding_ids\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          },\n          \"description\": \"Finding(s) the fact was promoted into; absent for dismissals.\"\n        }\n      }\n    },\n    \"effort\": {\n      \"type\": [\"string\", \"null\"],\n      \"enum\": [\"xs\", \"s\", \"m\", \"l\", \"xl\", null],\n      \"description\": \"Ordered effort-size labels; null means unknown size. Blocking is independent, and no duration thresholds are defined.\"\n    },\n    \"roadmap_priority\": {\n      \"type\": \"string\",\n      \"enum\": [\"p0\", \"p1\", \"p2\", \"p3\", \"p4\", \"P0\", \"P1\", \"P2\", \"P3\", \"P4\"],\n      \"description\": \"Priority labels in numeric order; no urgency or deadline mapping is defined. P0-P4 are accepted uppercase spellings of p0-p4 (enums/v1/roadmap-priority.json): readers normalise them and writers store lowercase.\"\n    },\n    \"canonical_finding\": {\n      \"allOf\": [\n        { \"$ref\": \"#/$defs/finding\" },\n        {\n          \"properties\": {\n            \"remediation_effort\": { \"$ref\": \"#/$defs/effort\" },\n            \"blocked_external\": { \"type\": \"boolean\" }\n          }\n        }\n      ]\n    },\n    \"canonical_write\": {\n      \"description\": \"Opt-in canonical effort sizes or null for unknown size. Default historical validation is unchanged.\",\n      \"allOf\": [\n        { \"$ref\": \"#\" },\n        {\n          \"properties\": {\n            \"findings\": {\n              \"items\": { \"$ref\": \"#/$defs/canonical_finding\" }\n            },\n            \"remediation_roadmap\": {\n              \"items\": {\n                \"properties\": {\n                  \"priority\": { \"$ref\": \"#/$defs/roadmap_priority\" },\n                  \"effort\": { \"$ref\": \"#/$defs/effort\" },\n                  \"blocked_external\": { \"type\": \"boolean\" }\n                }\n              }\n            },\n            \"disposition_summary\": {\n              \"properties\": {\n                \"severity_overrides\": {\n                  \"items\": {\n                    \"properties\": {\n                      \"from\": {\n                        \"anyOf\": [\n                          { \"$ref\": \"#/$defs/severity_level\" },\n                          { \"type\": \"null\" }\n                        ]\n                      }\n                    }\n                  }\n                }\n              }\n            }\n          }\n        }\n      ]\n    }\n  },\n  \"allOf\": [\n    {\n      \"$comment\": \"Identity is minted by the producing harness, never by a consumer. Any report that declares metadata.additional.harness_version MUST carry a fingerprint on every finding, so downstream consumers read stamped identity instead of recomputing it (divergent recomputation orphans ledger history). Legacy artifacts predating version stamping are exempt by having no harness_version.\",\n      \"if\": {\n        \"required\": [\n          \"metadata\"\n        ],\n        \"properties\": {\n          \"metadata\": {\n            \"required\": [\n              \"additional\"\n            ],\n            \"properties\": {\n              \"additional\": {\n                \"required\": [\n                  \"harness_version\"\n                ]\n              }\n            }\n          }\n        }\n      },\n      \"then\": {\n        \"properties\": {\n          \"findings\": {\n            \"items\": {\n              \"required\": [\n                \"fingerprint\"\n              ]\n            }\n          }\n        }\n      }\n    }\n  ]\n}\n",
	"risk-rating-methodology": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/risk-rating-methodology.schema.json",
  "title": "OWASP Risk Rating methodology tables",
  "description": "Validates the risk-rating methodology artifact \u2014 the machine-readable definition of how per-finding OWASP risk bands are derived (likelihood x impact, factors from CVSS 3.x vectors). The schema enforces the methodology's structural invariants: every CVSS metric value is mapped, all nine matrix cells are present, factor scores stay on the OWASP 0-9 scale, and the vector-less fallback likelihood can never reach the HIGH bucket. Enforced by tests/test_build_trends.py.",
  "type": "object",
  "required": [
    "methodology",
    "methodology_version",
    "source",
    "bands",
    "bucket_thresholds",
    "likelihood_factors",
    "impact_factors",
    "matrix",
    "fallback"
  ],
  "additionalProperties": false,
  "properties": {
    "methodology": {
      "const": "OWASP Risk Rating Methodology"
    },
    "methodology_version": {
      "type": "string",
      "pattern": "^\\d+\\.\\d+\\.\\d+$"
    },
    "source": {
      "type": "string",
      "format": "uri"
    },
    "documentation": {
      "type": "string"
    },
    "schema": {
      "type": "string"
    },
    "bands": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/band"
      },
      "minItems": 5,
      "maxItems": 5,
      "uniqueItems": true
    },
    "bucket_thresholds": {
      "type": "object",
      "required": [
        "medium_min",
        "high_min"
      ],
      "additionalProperties": false,
      "properties": {
        "medium_min": {
          "type": "number",
          "minimum": 0,
          "maximum": 4.5
        },
        "high_min": {
          "type": "number",
          "minimum": 4.5,
          "maximum": 9
        }
      }
    },
    "likelihood_factors": {
      "type": "object",
      "required": [
        "AV",
        "AC",
        "PR",
        "UI"
      ],
      "additionalProperties": false,
      "properties": {
        "AV": {
          "$ref": "#/definitions/factorMap",
          "required": [
            "N",
            "A",
            "L",
            "P"
          ]
        },
        "AC": {
          "$ref": "#/definitions/factorMap",
          "required": [
            "L",
            "H"
          ]
        },
        "PR": {
          "$ref": "#/definitions/factorMap",
          "required": [
            "N",
            "L",
            "H"
          ]
        },
        "UI": {
          "$ref": "#/definitions/factorMap",
          "required": [
            "N",
            "R"
          ]
        }
      }
    },
    "impact_factors": {
      "type": "object",
      "required": [
        "C",
        "I",
        "A"
      ],
      "additionalProperties": false,
      "properties": {
        "C": {
          "$ref": "#/definitions/ciaMap"
        },
        "I": {
          "$ref": "#/definitions/ciaMap"
        },
        "A": {
          "$ref": "#/definitions/ciaMap"
        }
      }
    },
    "matrix": {
      "type": "object",
      "required": [
        "LOW",
        "MEDIUM",
        "HIGH"
      ],
      "additionalProperties": false,
      "properties": {
        "LOW": {
          "$ref": "#/definitions/matrixRow"
        },
        "MEDIUM": {
          "$ref": "#/definitions/matrixRow"
        },
        "HIGH": {
          "$ref": "#/definitions/matrixRow"
        }
      }
    },
    "fallback": {
      "type": "object",
      "required": [
        "default_likelihood",
        "severity_impact"
      ],
      "additionalProperties": false,
      "properties": {
        "default_likelihood": {
          "description": "Likelihood for vector-less findings. Bounded below the HIGH bucket: unknown exploitability is never assumed HIGH, so vector-less findings cannot reach the critical band.",
          "type": "number",
          "minimum": 0,
          "exclusiveMaximum": 6
        },
        "severity_impact": {
          "type": "object",
          "required": [
            "critical",
            "high",
            "medium",
            "low",
            "informational"
          ],
          "additionalProperties": false,
          "properties": {
            "critical": {
              "$ref": "#/definitions/score"
            },
            "high": {
              "$ref": "#/definitions/score"
            },
            "medium": {
              "$ref": "#/definitions/score"
            },
            "low": {
              "$ref": "#/definitions/score"
            },
            "informational": {
              "$ref": "#/definitions/score"
            }
          }
        }
      }
    },
    "threat_intel_factor": {
      "type": "object",
      "description": "Optional exploitation-evidence likelihood factor (EPSS/KEV). Scores stay on the OWASP 0-9 scale; epss_bands must be ordered by descending min_epss and cover 0.0.",
      "required": [
        "kev_score",
        "epss_bands"
      ],
      "additionalProperties": false,
      "properties": {
        "description": {
          "type": "string"
        },
        "kev_score": {
          "type": "number",
          "minimum": 0,
          "maximum": 9
        },
        "epss_bands": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "object",
            "required": [
              "min_epss",
              "score"
            ],
            "additionalProperties": false,
            "properties": {
              "min_epss": {
                "type": "number",
                "minimum": 0,
                "maximum": 1
              },
              "score": {
                "type": "number",
                "minimum": 0,
                "maximum": 9
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "band": {
      "enum": [
        "critical",
        "high",
        "medium",
        "low",
        "note"
      ]
    },
    "score": {
      "type": "number",
      "minimum": 0,
      "maximum": 9
    },
    "factorMap": {
      "type": "object",
      "additionalProperties": {
        "$ref": "#/definitions/score"
      }
    },
    "ciaMap": {
      "type": "object",
      "required": [
        "H",
        "L",
        "N"
      ],
      "additionalProperties": false,
      "properties": {
        "H": {
          "$ref": "#/definitions/score"
        },
        "L": {
          "$ref": "#/definitions/score"
        },
        "N": {
          "$ref": "#/definitions/score"
        }
      }
    },
    "matrixRow": {
      "type": "object",
      "required": [
        "LOW",
        "MEDIUM",
        "HIGH"
      ],
      "additionalProperties": false,
      "properties": {
        "LOW": {
          "$ref": "#/definitions/band"
        },
        "MEDIUM": {
          "$ref": "#/definitions/band"
        },
        "HIGH": {
          "$ref": "#/definitions/band"
        }
      }
    }
  }
}
`,
	"run-manifest": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/run-manifest.schema.json",
  "title": "Run manifest",
  "description": "Per-order detail that is read one order at a time (data model §4a, Appendix G): the config in force, the budget verdict at admission, the approval, progress and degraded keys. run_order.artifact_id points at the latest manifest; each change is a new artifact (artifacts are append-only), so earlier manifests stay as history. The manifest does not name its order, so orders in the same state share one artifact.",
  "type": "object",
  "required": ["artifact"],
  "additionalProperties": false,
  "properties": {
    "artifact": {"const": "run-manifest"},
    "config": {
      "description": "Config documents in force when the order was planned, by digest (the pack file's digest is the pack version).",
      "type": "array",
      "uniqueItems": true,
      "items": {
        "type": "object",
        "required": ["name", "digest"],
        "additionalProperties": false,
        "properties": {
          "name": {"type": "string", "pattern": "^[a-z][a-z0-9_.-]*(:[a-z0-9_.-]+)?$", "description": "e.g. pack:security, budget-policy"},
          "digest": {"type": "string", "pattern": "^[0-9a-f]{64}$", "description": "sha256 of the document's bytes"},
          "artifact_id": {"type": "string", "format": "uuid", "description": "The stored document, when it is kept as an artifact."}
        }
      }
    },
    "budget": {
      "description": "The lane's budget verdict when admission decided this order.",
      "type": "object",
      "required": ["lane", "status", "spent_usd", "estimate_usd"],
      "additionalProperties": false,
      "properties": {
        "lane": {"type": "string", "minLength": 1},
        "status": {"type": "string", "enum": ["within", "over", "no_budget"]},
        "enforcement": {"type": ["string", "null"], "enum": ["none", "observe", "advisory", "enforced", null]},
        "spent_usd": {"$ref": "#/$defs/usd"},
        "estimate_usd": {"oneOf": [{"$ref": "#/$defs/usd"}, {"type": "null"}], "description": "null = unknown cost"},
        "budget_usd": {"oneOf": [{"$ref": "#/$defs/usd"}, {"type": "null"}]}
      }
    },
    "approval": {
      "type": "object",
      "required": ["approver", "approved_at"],
      "additionalProperties": false,
      "properties": {
        "approver": {"type": "string", "pattern": "^(human|machine):.+$", "description": "The verified principal, as rows store actors."},
        "approved_at": {"type": "string", "format": "date-time"}
      }
    },
    "progress": {
      "description": "Where a running order has got to, so a re-leased order resumes.",
      "type": "object",
      "required": ["updated_at"],
      "additionalProperties": false,
      "properties": {
        "step": {"type": ["integer", "null"], "minimum": 0, "description": "Steps (phases, stages) done."},
        "shards_done": {"type": "array", "items": {"type": "string", "minLength": 1}, "uniqueItems": true},
        "updated_at": {"type": "string", "format": "date-time"}
      }
    },
    "sessions": {
      "description": "Runner-computed evidence of staged inputs and the declared output, including failed sessions. Entries are append-only; an identical retry is not appended twice.",
      "type": "array",
      "uniqueItems": true,
      "items": {"$ref": "#/$defs/session"}
    },
    "degraded": {
      "description": "What the run could not see, by key (e.g. unreachable, quota-deferred, db-stale), with a count. A degraded order is never reported green.",
      "type": "object",
      "propertyNames": {"$ref": "#/$defs/degraded_key"},
      "additionalProperties": {"type": "integer", "minimum": 1}
    }
  },
  "$defs": {
    "session": {
      "type": "object",
      "additionalProperties": false,
      "required": ["skill", "mode", "executor_session", "models", "task_sha256", "inputs", "output"],
      "properties": {
        "skill": {"$ref": "session-task.schema.json#/$defs/name"},
        "mode": {"$ref": "session-task.schema.json#/$defs/name"},
        "executor_session": {"type": ["string", "null"], "minLength": 1},
        "models": {"$ref": "session-task.schema.json#/$defs/models"},
        "task_sha256": {"$ref": "#/$defs/sha256"},
        "inputs": {
          "description": "One entry per task-relative file path, sorted by path. The recording boundary enforces unique paths in linear time.",
          "type": "array",
          "items": {
            "type": "object",
            "additionalProperties": false,
            "required": ["path", "bytes", "sha256"],
            "properties": {
              "path": {"$ref": "session-task.schema.json#/$defs/path"},
              "bytes": {"type": "integer", "minimum": 0},
              "sha256": {"$ref": "#/$defs/sha256"}
            }
          }
        },
        "output": {
          "type": "object",
          "additionalProperties": false,
          "required": ["path", "sha256", "artifact_id"],
          "properties": {
            "path": {"allOf": [{"$ref": "session-task.schema.json#/$defs/path"}, {"pattern": "^output/"}]},
            "sha256": {"oneOf": [{"$ref": "#/$defs/sha256"}, {"type": "null"}]},
            "artifact_id": {"type": ["string", "null"], "format": "uuid"}
          },
          "allOf": [{"if": {"properties": {"artifact_id": {"type": "string"}}}, "then": {"properties": {"sha256": {"$ref": "#/$defs/sha256"}}}}]
        }
      }
    },
    "sha256": {"type": "string", "minLength": 64, "maxLength": 64, "pattern": "^[0-9a-f]{64}$"},
    "degraded_key": {
      "type": "string",
      "pattern": "^[a-z][a-z0-9_-]*$",
      "description": "Historical-compatible key shape. New Core completions enforce the degraded_key vocabulary or registered feed/view identifiers; dynamic subject identifiers retain their own coverage counts."
    },
    "usd": {"type": "string", "pattern": "^-?[0-9]+(\\.[0-9]+)?$"}
  }
}
`,
	"session-result": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/session-result.schema.json",
  "title": "Session result",
  "description": "Executor identity, measured model usage and an optional error; the judged document is the task's declared output, not this envelope.",
  "type": "object",
  "additionalProperties": false,
  "required": ["session", "usage", "error"],
  "properties": {
    "session": {"type": "string", "minLength": 1},
    "usage": {"type": "array", "items": {"$ref": "#/$defs/usage"}},
    "error": {"type": ["string", "null"], "minLength": 1}
  },
  "$defs": {
    "usage": {
      "type": "object",
      "additionalProperties": false,
      "required": ["provider", "model", "occurred_at", "tokens_in", "tokens_out"],
      "properties": {
        "provider": {"type": "string", "minLength": 1},
        "model": {"type": "string", "minLength": 1},
        "occurred_at": {"type": "string", "format": "date-time"},
        "tokens_in": {"type": "integer", "minimum": 0},
        "tokens_out": {"type": "integer", "minimum": 0},
        "cache_read": {"type": "integer", "minimum": 0},
        "cache_creation": {"type": "integer", "minimum": 0},
        "skill": {"type": ["string", "null"]},
        "session": {"type": ["string", "null"]}
      }
    }
  }
}
`,
	"session-task": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/session-task.schema.json",
  "title": "Session task",
  "description": "Runner-staged input and the one output a judgement session may produce. Paths are task-relative, never deployment paths. A manual session has no run order.",
  "type": "object",
  "additionalProperties": false,
  "required": ["schema_version", "run_id", "skill", "mode", "args", "models", "input", "output"],
  "properties": {
    "schema_version": {"const": "1"},
    "run_id": {"type": ["string", "null"], "format": "uuid"},
    "skill": {"$ref": "#/$defs/name"},
    "mode": {"$ref": "#/$defs/name"},
    "args": {"type": "array", "items": {"type": "string"}},
    "models": {"$ref": "#/$defs/models"},
    "input": {"type": "array", "uniqueItems": true, "items": {"$ref": "#/$defs/path"}},
    "output": {
      "type": "object",
      "additionalProperties": false,
      "required": ["path", "schema"],
      "properties": {
        "path": {"allOf": [{"$ref": "#/$defs/path"}, {"pattern": "^output/"}]},
        "schema": {"$ref": "#/$defs/name"}
      }
    }
  },
  "$defs": {
    "name": {
      "type": "string",
      "pattern": "^[a-z]",
      "not": {"pattern": "[^a-z0-9-]"}
    },
    "path": {
      "type": "string",
      "minLength": 1,
      "not": {"pattern": "^/|^[A-Za-z]:|\\\\|[\\x00-\\x1f]|(^|/)\\.{1,2}(/|$)|//|/$"}
    },
    "models": {
      "type": "object",
      "minProperties": 1,
      "propertyNames": {"pattern": "^[a-z]", "not": {"pattern": "[^a-z0-9_]"}},
      "additionalProperties": {"type": "string", "minLength": 1}
    }
  }
}
`,
	"sla-policy": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/sla-policy.schema.json",
  "title": "Owner-response SLA policy",
  "description": "Validates the SLA policy consumed by the SLA view builder. SLAs are POLICY DATA, never code: the shipped default reflects your organization's SLA policy configuration, and any other input \u2014 customer contracts, per-BU policies \u2014 is a different file passed via --policy, not an edit to the script. The severity_mapping block exists because external severity vocabularies (e.g. Critical/Important/Moderate/Low per your vulnerability management workflow manual) and the harness severity enum are different languages; the join must be auditable and swappable with the rest of the policy.",
  "type": "object",
  "required": [
    "policy_name",
    "source",
    "severity_mapping",
    "profiles"
  ],
  "additionalProperties": false,
  "properties": {
    "policy_name": {
      "type": "string",
      "pattern": "^[a-z0-9][a-z0-9-]*$"
    },
    "source": {
      "type": "object",
      "description": "Provenance of the numbers \u2014 where they came from and when they were last verified there.",
      "required": [
        "name",
        "retrieved"
      ],
      "additionalProperties": false,
      "properties": {
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        },
        "retrieved": {
          "type": "string",
          "pattern": "^\\d{4}-\\d{2}-\\d{2}$"
        },
        "note": {
          "type": "string"
        }
      }
    },
    "severity_mapping": {
      "type": "object",
      "description": "External severity label (lowercased) -> harness severity. Explicit data, never hardcoded.",
      "minProperties": 1,
      "additionalProperties": {
        "enum": [
          "critical",
          "high",
          "medium",
          "low",
          "informational"
        ]
      }
    },
    "clock_start": {
      "enum": [
        "first_routed_or_filed",
        "first_event",
        "audit_date"
      ],
      "default": "first_routed_or_filed",
      "description": "Which timestamp starts a finding's SLA clock. first_routed_or_filed: earliest routing/filing ledger event, falling back to first_event, falling back to audit_date."
    },
    "profiles": {
      "type": "object",
      "minProperties": 1,
      "additionalProperties": {
        "type": "object",
        "required": [
          "slas"
        ],
        "additionalProperties": false,
        "properties": {
          "default": {
            "type": "boolean"
          },
          "description": {
            "type": "string"
          },
          "slas": {
            "type": "object",
            "description": "Per harness severity. resolve_days null = no SLA (tracked, never overdue). Severities absent from the map are unclocked under this profile.",
            "additionalProperties": {
              "type": "object",
              "required": [
                "resolve_days"
              ],
              "additionalProperties": false,
              "properties": {
                "resolve_days": {
                  "type": [
                    "integer",
                    "null"
                  ],
                  "minimum": 1
                },
                "acknowledge_days": {
                  "type": [
                    "integer",
                    "null"
                  ],
                  "minimum": 1
                }
              }
            },
            "propertyNames": {
              "enum": [
                "critical",
                "high",
                "medium",
                "low",
                "informational"
              ]
            }
          },
          "cvss_floor_days": {
            "type": "object",
            "description": "Override: findings with cvss >= threshold get this resolve clock regardless of severity (e.g. a FedRAMP CVSS-floor requirement).",
            "required": [
              "threshold",
              "resolve_days"
            ],
            "additionalProperties": false,
            "properties": {
              "threshold": {
                "type": "number",
                "minimum": 0,
                "maximum": 10
              },
              "resolve_days": {
                "type": "integer",
                "minimum": 1
              }
            }
          }
        }
      }
    }
  }
}`,
	"threat-model": "{\n  \"$schema\": \"http://json-schema.org/draft-07/schema#\",\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/threat-model.schema.json\",\n  \"title\": \"Threat model\",\n  \"description\": \"One subject's threat model. The authored artifact is Markdown, because a threat model is written and edited by people; its structure has always been a contract, stated in prose in the skill's schema.md and enforced by regex in the linter. This is that contract as a machine-readable schema, so the model joins every other artifact family instead of being parsed out of prose by three hand-maintained copies of the column list.\",\n  \"type\": \"object\",\n  \"required\": [\"system\", \"provenance\", \"threats\"],\n  \"additionalProperties\": false,\n  \"properties\": {\n    \"system\": {\n      \"type\": \"string\",\n      \"minLength\": 1,\n      \"description\": \"The system this model covers.\"\n    },\n    \"subject_id\": {\n      \"type\": \"string\",\n      \"description\": \"The corpus subject, matching artifact_binding.subject_id. Without it a threat cannot reach its owner.\"\n    },\n    \"provenance\": { \"$ref\": \"#/$defs/provenance\" },\n    \"system_context\": { \"type\": \"string\" },\n    \"assets\": {\n      \"type\": \"array\",\n      \"description\": \"Section 2. One entry per thing worth protecting.\",\n      \"items\": { \"$ref\": \"#/$defs/asset\" }\n    },\n    \"entry_points\": {\n      \"type\": \"array\",\n      \"description\": \"Section 3. One entry per place untrusted input enters the system or privilege level changes.\",\n      \"items\": { \"$ref\": \"#/$defs/entry_point\" }\n    },\n    \"threats\": {\n      \"type\": \"array\",\n      \"description\": \"Section 4. The threat model proper: one row per actor-wants-outcome pair, at the abstraction level where it survives a patch.\",\n      \"items\": { \"$ref\": \"#/$defs/threat\" }\n    },\n    \"deprioritized\": {\n      \"type\": \"array\",\n      \"description\": \"Section 5. Retired threats. An id here is NEVER reused -- it records that a threat was considered and set aside, which is a different claim from never having considered it.\",\n      \"items\": { \"$ref\": \"#/$defs/deprioritized_threat\" }\n    },\n    \"open_questions\": { \"type\": \"array\", \"items\": { \"type\": \"string\" } },\n    \"mitigations\": {\n      \"type\": \"array\",\n      \"description\": \"Section 8, optional and additive. Each entry is one CLASS-LEVEL control, not a per-finding patch.\",\n      \"items\": { \"$ref\": \"#/$defs/mitigation\" }\n    },\n    \"attack_scenarios\": {\n      \"type\": \"array\",\n      \"description\": \"Section 9, optional. The narrative layer for the top 3-5 threats: each scenario tells the attack as it unfolds, in prose.\",\n      \"items\": { \"$ref\": \"#/$defs/attack_scenario\" }\n    },\n    \"tenant_boundaries\": {\n      \"type\": \"array\",\n      \"description\": \"Section 10, optional, MULTI-TENANT SERVICES ONLY. Single-tenant targets omit it; its absence is not a gap. Vocabulary is shared with isolation-review so boundary rows and full reviews line up.\",\n      \"items\": { \"$ref\": \"#/$defs/tenant_boundary\" }\n    },\n    \"update_history\": {\n      \"type\": \"array\",\n      \"description\": \"Appended by an update or review pass. Additive: `mode` in provenance keeps its original value.\",\n      \"items\": { \"$ref\": \"#/$defs/history_entry\" }\n    }\n  },\n  \"$defs\": {\n    \"provenance\": {\n      \"type\": \"object\",\n      \"description\": \"Section 7. How this model came to exist.\",\n      \"required\": [\"mode\", \"date\", \"target\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"mode\": { \"enum\": [\"interview\", \"bootstrap\", \"bootstrap-then-interview\"] },\n        \"date\": { \"type\": \"string\", \"format\": \"date\" },\n        \"target\": {\n          \"type\": \"string\",\n          \"description\": \"Path or repo URL at a commit. What was actually modelled.\"\n        },\n        \"inputs\": { \"type\": \"string\" },\n        \"owner\": {\n          \"type\": \"string\",\n          \"description\": \"Set for an interview, absent for a bootstrap: a model nobody reviewed must not look reviewed.\"\n        },\n        \"harness_version\": {\n          \"type\": \"string\",\n          \"description\": \"Marks the schema/taxonomy epoch, as metadata.harness_version does for reports. Legacy models predate it.\"\n        }\n      }\n    },\n    \"threat\": {\n      \"type\": \"object\",\n      \"required\": [\"id\", \"threat\", \"actor\", \"status\"],\n      \"anyOf\": [{ \"required\": [\"risk_rating\"] }, { \"required\": [\"impact\", \"likelihood\"] }],\n      \"$comment\": \"A threat is rated with risk_rating (the OWASP Risk Rating Methodology), which every new rating uses, or with the legacy impact/likelihood labels, which stay valid for threats not yet re-rated.\",\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"id\": {\n          \"type\": \"string\",\n          \"pattern\": \"^T[0-9]+$\",\n          \"description\": \"T1, T2, ... Stable across edits and NEVER renumbered: a removed id is retired, and the next new threat takes the next never-used number. Unique only within this model.\"\n        },\n        \"threat\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"One sentence, active voice, naming the outcome. A leading `linddun:` tag marks a row produced by the optional privacy overlay.\"\n        },\n        \"actor\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"items\": {\n            \"enum\": [\"remote_unauth\", \"remote_auth\", \"adjacent_network\", \"local_user\", \"local_admin\", \"supply_chain\", \"insider\"]\n          },\n          \"description\": \"Weakest-position actor FIRST: scoring reads the first entry.\"\n        },\n        \"surface\": { \"type\": \"string\" },\n        \"asset\": { \"type\": \"string\" },\n        \"impact\": {\n          \"enum\": [\"low\", \"medium\", \"high\", \"critical\", \"existential\"],\n          \"description\": \"LEGACY label, valid only for threats not yet re-rated; new ratings use risk_rating. Consequence if realised, modelled before reachability is known. Deliberately not the finding severity enum -- `existential` has no severity counterpart.\"\n        },\n        \"likelihood\": {\n          \"enum\": [\"very_rare\", \"rare\", \"possible\", \"likely\", \"almost_certain\"],\n          \"description\": \"LEGACY label, valid only for threats not yet re-rated; new ratings use risk_rating.\"\n        },\n        \"status\": {\n          \"enum\": [\"unmitigated\", \"partially_mitigated\", \"mitigated\", \"risk_accepted\"],\n          \"description\": \"`partially_mitigated` is its own state and the largest bucket in practice; folding it into `mitigated` overstates coverage more than any other collapse here.\"\n        },\n        \"controls\": { \"type\": \"string\" },\n        \"evidence\": {\n          \"type\": \"array\",\n          \"items\": { \"type\": \"string\" },\n          \"description\": \"Finding or validation identifiers. EMPTY means modelled but not evidenced -- a different claim from unmitigated.\"\n        },\n        \"attack_refs\": {\n          \"type\": \"array\",\n          \"items\": { \"type\": \"string\" },\n          \"description\": \"MITRE ATT&CK technique ids, validated against the pinned table. Column 11, default for new emissions since harness 0.82.0; this is what an ATT&CK coverage rollup reads.\"\n        },\n        \"owasp_llm_references\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"uniqueItems\": true,\n          \"items\": {\n            \"enum\": [\"LLM01\", \"LLM02\", \"LLM03\", \"LLM04\", \"LLM05\", \"LLM06\", \"LLM07\", \"LLM08\", \"LLM09\", \"LLM10\"]\n          },\n          \"description\": \"Optional. OWASP Top 10 for LLM Applications (2025) category IDs for a threat against an LLM, agent or MCP component, the same IDs report findings carry; the category ID is cited, never OWASP text.\"\n        },\n        \"atlas_refs\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"uniqueItems\": true,\n          \"items\": { \"type\": \"string\", \"pattern\": \"^AML\\\\.T[0-9]{4}(\\\\.[0-9]{3})?$\" },\n          \"description\": \"Optional. MITRE ATLAS technique or sub-technique IDs (AML.T0051, AML.T0051.000) for a threat against an LLM, agent or MCP component, validated against the pinned ATLAS table like attack_refs.\"\n        },\n        \"isolation_dimensions\": {\n          \"type\": \"array\",\n          \"items\": { \"enum\": [\"privilege\", \"encryption\", \"authentication\", \"connectivity\", \"hygiene\"] },\n          \"description\": \"Column 12, optional, multi-tenant only. Which isolation dimension this threat stresses.\"\n        },\n        \"risk_rating\": { \"$ref\": \"#/$defs/owasp_risk_rating\" },\n        \"review_surfaces\": {\n          \"$ref\": \"#/$defs/review_surfaces\",\n          \"description\": \"Review surfaces this threat sits on, so an audit's threat-model coverage diff maps modelled threats to its review_surface_coverage record.\"\n        }\n      }\n    },\n    \"asset\": {\n      \"type\": \"object\",\n      \"required\": [\"asset\", \"description\", \"sensitivity\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"asset\": { \"type\": \"string\", \"minLength\": 1 },\n        \"description\": { \"type\": \"string\", \"minLength\": 1 },\n        \"sensitivity\": { \"enum\": [\"low\", \"medium\", \"high\", \"critical\"] },\n        \"regulatory_scope\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Comma-separated regimes governing the asset (GDPR, CCPA, HIPAA, PCI-DSS, SOX, export-control, ...), or `none`.\"\n        },\n        \"example_records\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Concrete field or record names, e.g. `name, email, address` or `kubeconfig, pull-secret`.\"\n        }\n      }\n    },\n    \"entry_point\": {\n      \"type\": \"object\",\n      \"required\": [\"entry_point\", \"description\", \"trust_boundary\", \"reachable_assets\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"entry_point\": { \"type\": \"string\", \"minLength\": 1 },\n        \"description\": { \"type\": \"string\", \"minLength\": 1 },\n        \"trust_boundary\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"The crossing, in free text, e.g. `untrusted file -> process memory`.\"\n        },\n        \"reachable_assets\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Comma-separated names of section-2 assets reachable through this entry point.\"\n        },\n        \"review_surfaces\": {\n          \"$ref\": \"#/$defs/review_surfaces\",\n          \"description\": \"Review surfaces reachable through this entry point.\"\n        },\n        \"kind\": {\n          \"$ref\": \"#/$defs/entry_point_kind\",\n          \"description\": \"What kind of interface this is, so a coverage check can match the entry point to the deterministic pre-scan that enumerates it. Optional: an untagged entry point is reported as unexplained.\"\n        }\n      }\n    },\n    \"entry_point_kind\": {\n      \"type\": \"string\",\n      \"description\": \"Entry-point kinds (traust-contracts enums/v1/entry-point-kind.json).\",\n      \"enum\": [\n        \"http-api\",\n        \"grpc-api\",\n        \"k8s-controller\",\n        \"llm-integration\",\n        \"event-consumer\",\n        \"network-peer\",\n        \"cli\",\n        \"file-input\",\n        \"deploy-config\",\n        \"ci-supply-chain\",\n        \"metrics-debug\",\n        \"other\"\n      ]\n    },\n    \"review_surfaces\": {\n      \"type\": \"array\",\n      \"description\": \"Review surfaces (traust-contracts enums/v1/review-surface.json). Optional: an untagged threat, entry point or deprioritized threat is simply not mapped.\",\n      \"items\": {\n        \"enum\": [\n          \"credential-acceptance\",\n          \"ownership-resolution\",\n          \"security-flags\",\n          \"outbound-fetch\",\n          \"stored-content-integrity\",\n          \"scope-enforcement\",\n          \"llm-integration\"\n        ]\n      },\n      \"uniqueItems\": true\n    },\n    \"deprioritized_threat\": {\n      \"type\": \"object\",\n      \"required\": [\"threat\", \"reason\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"threat\": { \"type\": \"string\", \"minLength\": 1 },\n        \"reason\": { \"type\": \"string\", \"minLength\": 1 },\n        \"review_surfaces\": {\n          \"$ref\": \"#/$defs/review_surfaces\",\n          \"description\": \"Review surfaces the set-aside threat sits on: the model considered the surface and recorded why the threat was not carried. Distinct from no tag, which says nothing about the surface.\"\n        }\n      }\n    },\n    \"attack_scenario\": {\n      \"type\": \"object\",\n      \"required\": [\"id\", \"steps\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"id\": {\n          \"type\": \"string\",\n          \"pattern\": \"^T[0-9]+$\",\n          \"description\": \"The section-4 threat this scenario narrates.\"\n        },\n        \"threat\": { \"type\": \"string\", \"minLength\": 1 },\n        \"steps\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"description\": \"The scenario's prose, one paragraph per entry.\",\n          \"items\": { \"type\": \"string\", \"minLength\": 1 }\n        }\n      }\n    },\n    \"history_entry\": {\n      \"type\": \"object\",\n      \"required\": [\"date\", \"changes\", \"reason\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"date\": { \"type\": \"string\", \"format\": \"date\", \"pattern\": \"^[0-9]{4}-[0-9]{2}-[0-9]{2}$\" },\n        \"changes\": { \"type\": \"string\", \"minLength\": 1 },\n        \"reason\": { \"type\": \"string\", \"minLength\": 1 }\n      }\n    },\n    \"mitigation\": {\n      \"type\": \"object\",\n      \"required\": [\"mitigation\", \"threat_ids\", \"closes_class\", \"effort\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"mitigation\": { \"type\": \"string\", \"minLength\": 1 },\n        \"threat_ids\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"items\": { \"type\": \"string\", \"pattern\": \"^T[0-9]+$\" }\n        },\n        \"closes_class\": {\n          \"enum\": [\"yes\", \"partial\", \"no\"],\n          \"description\": \"`no` is worthwhile defence in depth that does not close the class -- recorded rather than dropped.\"\n        },\n        \"effort\": { \"enum\": [\"XS\", \"S\", \"M\", \"L\", \"xs\", \"s\", \"m\", \"l\", \"xl\", null] },\n        \"blocked_external\": { \"type\": \"boolean\" }\n      }\n    },\n    \"tenant_boundary\": {\n      \"type\": \"object\",\n      \"required\": [\"boundary_id\", \"interface\", \"kind\", \"exposure\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"boundary_id\": {\n          \"type\": \"string\",\n          \"pattern\": \"^IF-[0-9]+$\",\n          \"description\": \"Stable, never renumbered or reused. Aligned with isolation_review.interfaces[].id when a full review exists.\"\n        },\n        \"interface\": { \"type\": \"string\", \"minLength\": 1 },\n        \"kind\": { \"enum\": [\"api\", \"data-store\", \"queue\", \"ingress\", \"webhook\", \"cli\", \"other\"] },\n        \"exposure\": {\n          \"enum\": [\"public\", \"tenant\", \"partner\", \"internal\"],\n          \"description\": \"`public` is reachable without tenant credentials; `internal` is control-plane only but carries tenant data.\"\n        },\n        \"complexity\": { \"enum\": [\"low\", \"medium\", \"high\"] },\n        \"privilege\": { \"type\": \"string\" },\n        \"encryption\": { \"type\": \"string\" },\n        \"authentication\": { \"type\": \"string\" },\n        \"connectivity\": { \"type\": \"string\" },\n        \"hygiene\": { \"type\": \"string\" },\n        \"threat_ids\": {\n          \"type\": \"array\",\n          \"items\": { \"type\": \"string\", \"pattern\": \"^T[0-9]+$\" }\n        },\n        \"isolation_review_ref\": { \"type\": \"string\" }\n      }\n    },\n    \"owasp_risk_rating\": {\n      \"type\": \"object\",\n      \"description\": \"Risk rating per the OWASP Risk Rating Methodology (OWASP Foundation, https://owasp.org/www-community/OWASP_Risk_Rating_Methodology, CC BY-SA 4.0; reimplemented here in original wording). Every factor is scored 0 to 9. The likelihood score is the mean of its eight factors; the impact score is the mean of the technical or the business factors, per basis. A score below 3 is low, below 6 medium, otherwise high. Severity combines the impact and likelihood levels by the methodology's severity table. Scores, levels and severity are derived values; traust_contracts.v1.risk_rating computes and checks them.\",\n      \"required\": [\"method\", \"likelihood\", \"impact\", \"severity\"],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"method\": { \"const\": \"owasp-risk-rating\" },\n        \"likelihood\": {\n          \"type\": \"object\",\n          \"required\": [\"factors\", \"score\", \"level\"],\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"factors\": {\n              \"type\": \"object\",\n              \"description\": \"Threat-agent factors (skill_level, motive, opportunity, population_size) and vulnerability factors (ease_of_discovery, ease_of_exploit, awareness, intrusion_detection). Higher means more likely.\",\n              \"required\": [\"skill_level\", \"motive\", \"opportunity\", \"population_size\", \"ease_of_discovery\", \"ease_of_exploit\", \"awareness\", \"intrusion_detection\"],\n              \"additionalProperties\": false,\n              \"properties\": {\n                \"skill_level\": { \"$ref\": \"#/$defs/owasp_factor\" },\n                \"motive\": { \"$ref\": \"#/$defs/owasp_factor\" },\n                \"opportunity\": { \"$ref\": \"#/$defs/owasp_factor\" },\n                \"population_size\": { \"$ref\": \"#/$defs/owasp_factor\" },\n                \"ease_of_discovery\": { \"$ref\": \"#/$defs/owasp_factor\" },\n                \"ease_of_exploit\": { \"$ref\": \"#/$defs/owasp_factor\" },\n                \"awareness\": { \"$ref\": \"#/$defs/owasp_factor\" },\n                \"intrusion_detection\": { \"$ref\": \"#/$defs/owasp_factor\" }\n              }\n            },\n            \"score\": { \"$ref\": \"#/$defs/owasp_score\" },\n            \"level\": { \"$ref\": \"#/$defs/owasp_level\" }\n          }\n        },\n        \"impact\": {\n          \"type\": \"object\",\n          \"required\": [\"basis\", \"technical\", \"score\", \"level\"],\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"basis\": {\n              \"enum\": [\"technical\", \"business\"],\n              \"description\": \"Which factor set sets the impact score. The methodology prefers business impact when it is known, and technical impact otherwise.\"\n            },\n            \"technical\": {\n              \"type\": \"object\",\n              \"required\": [\"confidentiality\", \"integrity\", \"availability\", \"accountability\"],\n              \"additionalProperties\": false,\n              \"properties\": {\n                \"confidentiality\": { \"$ref\": \"#/$defs/owasp_factor\" },\n                \"integrity\": { \"$ref\": \"#/$defs/owasp_factor\" },\n                \"availability\": { \"$ref\": \"#/$defs/owasp_factor\" },\n                \"accountability\": { \"$ref\": \"#/$defs/owasp_factor\" }\n              }\n            },\n            \"business\": {\n              \"type\": \"object\",\n              \"required\": [\"financial\", \"reputation\", \"non_compliance\", \"privacy\"],\n              \"additionalProperties\": false,\n              \"properties\": {\n                \"financial\": { \"$ref\": \"#/$defs/owasp_factor\" },\n                \"reputation\": { \"$ref\": \"#/$defs/owasp_factor\" },\n                \"non_compliance\": { \"$ref\": \"#/$defs/owasp_factor\" },\n                \"privacy\": { \"$ref\": \"#/$defs/owasp_factor\" }\n              }\n            },\n            \"score\": { \"$ref\": \"#/$defs/owasp_score\" },\n            \"level\": { \"$ref\": \"#/$defs/owasp_level\" }\n          },\n          \"if\": { \"properties\": { \"basis\": { \"const\": \"business\" } } },\n          \"then\": { \"required\": [\"business\"] }\n        },\n        \"severity\": {\n          \"enum\": [\"note\", \"low\", \"medium\", \"high\", \"critical\"],\n          \"description\": \"Overall severity from the impact and likelihood levels.\"\n        },\n        \"rationale\": {\n          \"type\": \"object\",\n          \"additionalProperties\": { \"type\": \"string\", \"minLength\": 1 },\n          \"description\": \"One short reason per factor, keyed by factor name, so a reviewer can check each score.\"\n        }\n      }\n    },\n    \"canonical_write\": {\n      \"description\": \"Opt-in canonical mitigation sizes or null for unknown size; effort remains required.\",\n      \"allOf\": [\n        { \"$ref\": \"#\" },\n        {\n          \"properties\": {\n            \"mitigations\": {\n              \"items\": {\n                \"properties\": {\n                  \"effort\": { \"$ref\": \"report.schema.json#/$defs/effort\" },\n                  \"blocked_external\": { \"type\": \"boolean\" }\n                }\n              }\n            }\n          }\n        }\n      ]\n    },\n    \"owasp_factor\": { \"type\": \"integer\", \"minimum\": 0, \"maximum\": 9 },\n    \"owasp_score\": { \"type\": \"number\", \"minimum\": 0, \"maximum\": 9 },\n    \"owasp_level\": { \"enum\": [\"low\", \"medium\", \"high\"] }\n  }\n}\n",
	"triage":       "{\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/triage.schema.json\",\n  \"$schema\": \"https://json-schema.org/draft/2020-12/schema\",\n  \"title\": \"Triage Report Schema\",\n  \"description\": \"Output contract of the triage skill (TRIAGE.json / <repo>-triage.json). Finding kind is independent of the triage verdict: verified defense-in-depth gaps use true_positive with kind weakness. Historical hardening verdicts remain readable, but canonical_write rejects them. Omitted kind remains unknown. Severity, finding-kind and date vocabularies are shared with report.schema.json.\",\n  \"type\": \"object\",\n  \"required\": [\n    \"triage_completed\",\n    \"triage_context\",\n    \"summary\",\n    \"findings\"\n  ],\n  \"properties\": {\n    \"triage_completed\": {\n      \"description\": \"Date the verification run completed (not the re-emit date; see triage_context.reemitted).\",\n      \"type\": \"string\",\n      \"format\": \"date\"\n    },\n    \"triage_context\": {\n      \"$ref\": \"#/$defs/triage_context\"\n    },\n    \"summary\": {\n      \"$ref\": \"#/$defs/summary\"\n    },\n    \"findings\": {\n      \"type\": \"array\",\n      \"minItems\": 0,\n      \"$comment\": \"Empty is valid: a zero-finding audit yields a zero-finding triage.\",\n      \"items\": {\n        \"$ref\": \"#/$defs/finding\"\n      }\n    }\n  },\n  \"additionalProperties\": false,\n  \"$defs\": {\n    \"severity_level\": {\n      \"type\": \"string\",\n      \"enum\": [\n        \"critical\",\n        \"high\",\n        \"medium\",\n        \"low\",\n        \"informational\"\n      ]\n    },\n    \"verdict\": {\n      \"type\": \"string\",\n      \"enum\": [\n        \"true_positive\",\n        \"hardening\",\n        \"undetermined\",\n        \"false_positive\",\n        \"duplicate\"\n      ]\n    },\n    \"triage_context\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"environment\",\n        \"votes_per_finding\",\n        \"repo\",\n        \"harness_version\"\n      ],\n      \"properties\": {\n        \"mode\": {\n          \"type\": \"string\"\n        },\n        \"environment\": {\n          \"type\": \"string\",\n          \"minLength\": 10\n        },\n        \"threat_model\": {\n          \"anyOf\": [\n            {\n              \"type\": \"array\",\n              \"items\": {\n                \"type\": \"string\"\n              }\n            },\n            {\n              \"type\": \"string\"\n            }\n          ]\n        },\n        \"scoring\": {\n          \"type\": \"string\"\n        },\n        \"noise_tolerance\": {\n          \"type\": \"string\"\n        },\n        \"votes_per_finding\": {\n          \"type\": \"integer\",\n          \"minimum\": 1\n        },\n        \"repo\": {\n          \"type\": \"string\"\n        },\n        \"findings_path\": {\n          \"type\": \"string\"\n        },\n        \"source_report\": {\n          \"type\": \"string\"\n        },\n        \"harness_version\": {\n          \"description\": \"Harness semver + short SHA at triage (or re-emit) time, e.g. 0.24.0-e240e87. Pins the verdict-taxonomy epoch.\",\n          \"type\": \"string\",\n          \"pattern\": \"^\\\\d+\\\\.\\\\d+\\\\.\\\\d+(-[a-f0-9]{7,12})?$\"\n        },\n        \"rerun_of\": {\n          \"type\": \"string\"\n        },\n        \"reemitted\": {\n          \"type\": \"string\",\n          \"format\": \"date\"\n        },\n        \"reemit_method\": {\n          \"type\": \"string\"\n        }\n      },\n      \"additionalProperties\": true\n    },\n    \"summary\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"input_count\",\n        \"true_positives\",\n        \"hardening\",\n        \"false_positives\",\n        \"undetermined\",\n        \"duplicates\",\n        \"by_severity\"\n      ],\n      \"properties\": {\n        \"input_count\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"true_positives\": {\n          \"description\": \"Actionable true_positive findings whose kind is not weakness; excludes the disjoint hardening bucket.\",\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"hardening\": {\n          \"description\": \"Historical hardening verdicts and true_positive findings with kind weakness; excludes the true_positives bucket.\",\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"false_positives\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"undetermined\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"duplicates\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"needs_manual_test\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"by_severity\": {\n          \"description\": \"Severity distribution of actionable non-weakness true positives only, using the report.schema.json severity vocabulary.\",\n          \"type\": \"object\",\n          \"required\": [\n            \"critical\",\n            \"high\",\n            \"medium\",\n            \"low\"\n          ],\n          \"properties\": {\n            \"critical\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"high\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"medium\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"low\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"informational\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            }\n          },\n          \"additionalProperties\": false\n        }\n      },\n      \"additionalProperties\": true\n    },\n    \"vote_breakdown\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"true_positive\",\n        \"hardening\",\n        \"false_positive\",\n        \"cannot_verify\"\n      ],\n      \"properties\": {\n        \"true_positive\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"hardening\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"false_positive\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"cannot_verify\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        }\n      },\n      \"additionalProperties\": false\n    },\n    \"finding\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"id\",\n        \"title\",\n        \"verdict\"\n      ],\n      \"properties\": {\n        \"id\": {\n          \"description\": \"Triage-local sequential id.\",\n          \"type\": \"string\",\n          \"pattern\": \"^f\\\\d{3,}$\"\n        },\n        \"orig_id\": {\n          \"description\": \"The source scanner's finding id. When the input is a harness audit report this is the canonical {REPO_SLUG}-{SHORTSHA}-{NNN} id and is cross-checked.\",\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"source\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"title\": {\n          \"type\": \"string\",\n          \"minLength\": 5\n        },\n        \"file\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"line\": {\n          \"type\": [\n            \"integer\",\n            \"null\"\n          ]\n        },\n        \"category\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"claimed_severity\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"verdict\": {\n          \"$ref\": \"#/$defs/verdict\"\n        },\n        \"kind\": {\n          \"$ref\": \"report.schema.json#/$defs/finding/properties/kind\",\n          \"description\": \"Finding kind, independent of verdict. Optional and non-null; omission preserves unknown kind rather than assuming vulnerability.\"\n        },\n        \"verify_verdict\": {\n          \"description\": \"exploitable|mitigated|needs_manual_test are the skill's canonical values; confirmed|refuted|hardening are accepted presentation aliases.\",\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"enum\": [\n            \"exploitable\",\n            \"mitigated\",\n            \"needs_manual_test\",\n            \"confirmed\",\n            \"refuted\",\n            \"hardening\",\n            null\n          ]\n        },\n        \"confidence\": {\n          \"type\": [\n            \"number\",\n            \"null\"\n          ],\n          \"minimum\": 0,\n          \"maximum\": 10\n        },\n        \"severity\": {\n          \"description\": \"Precondition-derived severity; set only on actionable non-weakness true positives (null for confirmed weaknesses and other verdicts).\",\n          \"anyOf\": [\n            {\n              \"$ref\": \"#/$defs/severity_level\"\n            },\n            {\n              \"type\": \"null\"\n            }\n          ]\n        },\n        \"severity_label\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"severity_alignment\": {\n          \"type\": [\n            \"string\",\n            \"number\",\n            \"null\"\n          ]\n        },\n        \"preconditions\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"access_level\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"threat_match\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"rationale\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"recommendation\": {\n          \"description\": \"Remediation guidance carried through verbatim from the input finding (e.g. a secure-code-audit report's remediation field), so /patch receives it on the canonical audit -> triage -> patch path. Null when the input had none.\",\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"vote_breakdown\": {\n          \"anyOf\": [\n            {\n              \"$ref\": \"#/$defs/vote_breakdown\"\n            },\n            {\n              \"type\": \"null\"\n            }\n          ]\n        },\n        \"refute_reasons\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"exclusion_rule\": {\n          \"type\": [\n            \"string\",\n            \"integer\",\n            \"null\"\n          ]\n        },\n        \"first_links\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"duplicate_of\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"absorbed\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"owner_hint\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"component\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"missing_fields\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"fingerprint\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{64}$\",\n          \"description\": \"Deterministic cross-scan identity (identity module fingerprint). Optional here -- no artifact of this kind carries one today -- but DECLARED, because additionalProperties is false and an undeclared identity field makes every stamped artifact invalid the moment a producer starts emitting it. That has happened twice: report.schema.json omitted fingerprint_algo for 16 days, and the cloud-config schema declared neither, so stamping the policy findings invalidated nearly every report in one command.\"\n        },\n        \"fingerprint_algo\": {\n          \"type\": \"string\",\n          \"pattern\": \"^v\\\\d+$\",\n          \"description\": \"Algorithm version of `fingerprint` (the identity module ALGO_VERSION). Declared for the same reason as `fingerprint` above.\"\n        }\n      },\n      \"additionalProperties\": false\n    },\n    \"canonical_write\": {\n      \"description\": \"Opt-in canonical triage verdicts and claimed severity from the existing vocabulary. True-positive weaknesses require rule 13 and cannot carry non-null severity; true positives with rule 13 require explicit kind weakness. Deprecated hardening verdicts and historical claimed-severity text remain readable by default.\",\n      \"allOf\": [\n        { \"$ref\": \"#\" },\n        {\n          \"properties\": {\n            \"findings\": {\n              \"items\": {\n                \"properties\": {\n                  \"verdict\": {\n                    \"not\": { \"const\": \"hardening\" }\n                  },\n                  \"claimed_severity\": {\n                    \"anyOf\": [\n                      { \"$ref\": \"report.schema.json#/$defs/severity_level\" },\n                      { \"type\": \"null\" }\n                    ]\n                  }\n                },\n                \"allOf\": [\n                  {\n                    \"if\": {\n                      \"required\": [\"verdict\", \"kind\"],\n                      \"properties\": {\n                        \"verdict\": { \"const\": \"true_positive\" },\n                        \"kind\": { \"const\": \"weakness\" }\n                      }\n                    },\n                    \"then\": {\n                      \"required\": [\"exclusion_rule\"],\n                      \"properties\": {\n                        \"severity\": { \"type\": \"null\" },\n                        \"exclusion_rule\": { \"enum\": [13, \"13\"] }\n                      }\n                    }\n                  },\n                  {\n                    \"if\": {\n                      \"required\": [\"verdict\", \"exclusion_rule\"],\n                      \"properties\": {\n                        \"verdict\": { \"const\": \"true_positive\" },\n                        \"exclusion_rule\": { \"enum\": [13, \"13\"] }\n                      }\n                    },\n                    \"then\": {\n                      \"required\": [\"kind\"],\n                      \"properties\": {\n                        \"kind\": { \"const\": \"weakness\" }\n                      }\n                    }\n                  }\n                ]\n              }\n            }\n          }\n        }\n      ]\n    }\n  }\n}\n",
	"validation":   "{\n  \"$schema\": \"https://json-schema.org/draft/2020-12/schema\",\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/validation.schema.json\",\n  \"title\": \"Live Validation Report Schema\",\n  \"description\": \"Schema for *-validation.json reports produced by the validate-findings harness. Records replay, attack-chain, and novel-attack results against a live authorized target.\",\n  \"type\": \"object\",\n  \"required\": [\n    \"title\",\n    \"metadata\",\n    \"source_reports\",\n    \"summary\",\n    \"validated_findings\",\n    \"attack_chains\",\n    \"novel_findings\",\n    \"execution_log_ref\"\n  ],\n  \"additionalProperties\": false,\n  \"properties\": {\n    \"title\": {\n      \"type\": \"string\",\n      \"minLength\": 5,\n      \"pattern\": \"(?i)(validation|live|attack|exploit|red.?team)\"\n    },\n    \"metadata\": {\n      \"$ref\": \"#/$defs/validation_metadata\"\n    },\n    \"source_reports\": {\n      \"type\": \"array\",\n      \"minItems\": 1,\n      \"items\": {\n        \"$ref\": \"#/$defs/source_report\"\n      }\n    },\n    \"summary\": {\n      \"$ref\": \"#/$defs/summary\"\n    },\n    \"validated_findings\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/validated_finding\"\n      }\n    },\n    \"attack_chains\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/attack_chain\"\n      }\n    },\n    \"novel_findings\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/novel_finding\"\n      }\n    },\n    \"negative_results\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"report.schema.json#/$defs/negative_result\"\n      }\n    },\n    \"execution_log_ref\": {\n      \"type\": \"string\",\n      \"minLength\": 1,\n      \"description\": \"Relative path to the append-only validation-audit.jsonl\"\n    },\n    \"execution_log_sha256\": {\n      \"type\": \"string\",\n      \"pattern\": \"^[a-f0-9]{64}$\"\n    },\n    \"footer\": {\n      \"type\": \"string\"\n    }\n  },\n  \"$defs\": {\n    \"verdict\": {\n      \"type\": \"string\",\n      \"enum\": [\n        \"confirmed\",\n        \"refuted\",\n        \"inconclusive\",\n        \"blocked_by_scope\",\n        \"not_attempted\"\n      ]\n    },\n    \"technique\": {\n      \"type\": \"string\",\n      \"enum\": [\n        \"replay\",\n        \"adapted\",\n        \"chained\",\n        \"novel\",\n        \"recon\",\n        \"skip\"\n      ]\n    },\n    \"classification\": {\n      \"type\": \"string\",\n      \"enum\": [\n        \"safe\",\n        \"mutating\",\n        \"destructive\"\n      ]\n    },\n    \"adapter\": {\n      \"type\": \"string\",\n      \"enum\": [\n        \"k8s\",\n        \"container\",\n        \"wasm\"\n      ]\n    },\n    \"scope_binding_mode\": {\n      \"type\": \"string\",\n      \"enum\": [\n        \"explicit\",\n        \"inline\",\n        \"inferred\",\n        \"mixed\",\n        \"none\"\n      ]\n    },\n    \"validation_metadata\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"date\",\n        \"harness_version\",\n        \"scope_binding_mode\",\n        \"target_fingerprint\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"date\": {\n          \"type\": \"string\",\n          \"format\": \"date\"\n        },\n        \"harness_version\": {\n          \"type\": \"string\",\n          \"minLength\": 3\n        },\n        \"scope_binding_mode\": {\n          \"$ref\": \"#/$defs/scope_binding_mode\"\n        },\n        \"scope_source\": {\n          \"type\": \"string\",\n          \"description\": \"Path to targets.yaml or summary of inline/inferred scope\"\n        },\n        \"engagement\": {\n          \"type\": \"string\"\n        },\n        \"authorized_by\": {\n          \"type\": \"string\"\n        },\n        \"expires\": {\n          \"type\": \"string\",\n          \"format\": \"date\"\n        },\n        \"target_fingerprint\": {\n          \"type\": \"array\",\n          \"minItems\": 0,\n          \"items\": {\n            \"$ref\": \"#/$defs/fingerprint\"\n          }\n        },\n        \"target_environment\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"What this run was executed AGAINST -- the bound cluster context(s), a cluster role, an environment id. Named target_environment, beside target_fingerprint and target_attestation, because `environment` alone is already taken in this domain: /validate-findings uses it for the SAFETY class of an engagement (`environment: lab` gates --auto), which is a different fact. Two runs of the same finding set against different target environments are not re-runs of each other and neither supersedes the other -- measured, a hub run and a spoke run of one subject covered the same findings and disagreed on a material share of the verdicts, some confirmed against one target and refuted against the other. Optional, because artifacts predate it -- but ABSENT means UNKNOWN, not 'the same as the others', and a consumer must not merge unknowns.\"\n        },\n        \"approval\": {\n          \"type\": \"object\",\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"mode\": {\n              \"type\": \"string\",\n              \"enum\": [\n                \"interactive\",\n                \"auto\",\n                \"dry-run\"\n              ]\n            },\n            \"by\": {\n              \"type\": \"string\"\n            },\n            \"at\": {\n              \"type\": \"string\",\n              \"format\": \"date-time\"\n            },\n            \"environment\": {\n              \"type\": \"string\"\n            }\n          }\n        },\n        \"flags\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"additional\": {\n          \"type\": \"object\"\n        },\n        \"target_attestation\": {\n          \"type\": [\n            \"object\",\n            \"null\"\n          ],\n          \"description\": \"P2 pre-flight attestation summary (attest_target.py): {attested, checks[], fingerprint}. The full artifact lives beside the report as target-attestation.json; reports at/after harness 0.176.0 without one cannot ledger verdicts (emitter routes everything to environment_invalid/needs_review).\"\n        }\n      }\n    },\n    \"fingerprint\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"adapter\",\n        \"identity\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"adapter\": {\n          \"$ref\": \"#/$defs/adapter\"\n        },\n        \"identity\": {\n          \"type\": \"string\",\n          \"description\": \"kube context name, container id, or wasm artifact path\"\n        },\n        \"version\": {\n          \"type\": \"string\"\n        },\n        \"digest\": {\n          \"type\": \"string\"\n        },\n        \"details\": {\n          \"type\": \"object\"\n        }\n      }\n    },\n    \"source_report\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"kind\",\n        \"path\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"kind\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"security-audit\",\n            \"threat-model\",\n            \"triage\"\n          ]\n        },\n        \"path\": {\n          \"type\": \"string\"\n        },\n        \"sha256\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[a-f0-9]{64}$\"\n        }\n      }\n    },\n    \"summary\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"by_verdict\",\n        \"by_technique\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"prose\": {\n          \"type\": \"string\"\n        },\n        \"by_verdict\": {\n          \"type\": \"object\",\n          \"required\": [\n            \"confirmed\",\n            \"refuted\",\n            \"inconclusive\",\n            \"blocked_by_scope\",\n            \"not_attempted\"\n          ],\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"confirmed\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"refuted\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"inconclusive\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"blocked_by_scope\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"not_attempted\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            }\n          }\n        },\n        \"by_technique\": {\n          \"type\": \"object\",\n          \"additionalProperties\": {\n            \"type\": \"integer\",\n            \"minimum\": 0\n          }\n        },\n        \"novel_count\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"chain_count\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"highest_impact_chain\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        }\n      }\n    },\n    \"evidence_artifact\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"type\",\n        \"path\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"type\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"log\",\n            \"manifest\",\n            \"http\",\n            \"dump\",\n            \"screenshot\",\n            \"diff\",\n            \"stdout\",\n            \"other\"\n          ]\n        },\n        \"path\": {\n          \"type\": \"string\"\n        },\n        \"sha256\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[a-f0-9]{64}$\"\n        },\n        \"caption\": {\n          \"type\": \"string\"\n        },\n        \"redacted\": {\n          \"type\": \"boolean\",\n          \"description\": \"True when this evidence file had secrets stripped before it was retained. Chain of custody: a consumer must be able to tell sanitised evidence from raw, and dropping the flag would make the two indistinguishable.\"\n        }\n      }\n    },\n    \"step_result\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"step_id\",\n        \"adapter\",\n        \"verb\",\n        \"classification\",\n        \"verdict\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"step_id\": {\n          \"type\": \"string\"\n        },\n        \"finding_ref\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"novel_ref\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"adapter\": {\n          \"$ref\": \"#/$defs/adapter\"\n        },\n        \"verb\": {\n          \"type\": \"string\"\n        },\n        \"target\": {\n          \"type\": \"object\"\n        },\n        \"classification\": {\n          \"$ref\": \"#/$defs/classification\"\n        },\n        \"verdict\": {\n          \"$ref\": \"#/$defs/verdict\"\n        },\n        \"expected\": {\n          \"type\": \"string\"\n        },\n        \"observed\": {\n          \"type\": \"string\"\n        },\n        \"evidence\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/$defs/evidence_artifact\"\n          }\n        },\n        \"rollback_performed\": {\n          \"type\": [\n            \"boolean\",\n            \"null\"\n          ]\n        },\n        \"rollback_output\": {\n          \"type\": \"string\"\n        },\n        \"scope_reason\": {\n          \"type\": \"string\"\n        },\n        \"duration_ms\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"soundness_flag\": {\n          \"type\": \"string\",\n          \"description\": \"Set when a refuted verdict was gated to inconclusive by the refutation-soundness gate: error-signature:<name> | rbac-zero-subjects | rbac-template-placeholder | target-not-deployed.\"\n        },\n        \"error\": {\n          \"type\": \"string\",\n          \"description\": \"Environmental failure note (e.g. probe-target-not-found).\"\n        },\n        \"controls\": {\n          \"type\": [\n            \"array\",\n            \"null\"\n          ],\n          \"items\": {\n            \"$ref\": \"#/$defs/positive_control\"\n          }\n        },\n        \"differential\": {\n          \"anyOf\": [\n            {\n              \"$ref\": \"#/$defs/differential_probe\"\n            },\n            {\n              \"type\": \"null\"\n            }\n          ]\n        },\n        \"replay\": {\n          \"anyOf\": [\n            {\n              \"$ref\": \"#/$defs/replay_artifact\"\n            },\n            {\n              \"type\": \"null\"\n            }\n          ]\n        }\n      }\n    },\n    \"validated_finding\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"source_id\",\n        \"source_report\",\n        \"verdict\",\n        \"technique\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"source_id\": {\n          \"type\": \"string\",\n          \"description\": \"Finding ID in the source security-audit / triage report\"\n        },\n        \"source_report\": {\n          \"type\": \"string\"\n        },\n        \"title\": {\n          \"type\": \"string\"\n        },\n        \"claimed_severity\": {\n          \"$ref\": \"report.schema.json#/$defs/severity_level\"\n        },\n        \"verdict\": {\n          \"$ref\": \"#/$defs/verdict\"\n        },\n        \"technique\": {\n          \"$ref\": \"#/$defs/technique\"\n        },\n        \"steps\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/$defs/step_result\"\n          }\n        },\n        \"evidence\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/$defs/evidence_artifact\"\n          }\n        },\n        \"observed_impact\": {\n          \"type\": \"string\"\n        },\n        \"deviation_from_claim\": {\n          \"type\": \"string\"\n        },\n        \"rollback_performed\": {\n          \"type\": [\n            \"boolean\",\n            \"null\"\n          ]\n        },\n        \"not_attempted_reason\": {\n          \"type\": \"string\"\n        },\n        \"soundness_flag\": {\n          \"type\": \"string\",\n          \"description\": \"Machine-readable reason the driving refutation was un-emittable (soundness gate). Findings carrying this flag route to needs_review, never to the false-positive/countersign path.\"\n        },\n        \"evidence_grade\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"enum\": [\n            \"E0\",\n            \"E1\",\n            \"E2\",\n            \"E3\",\n            null\n          ],\n          \"description\": \"P4 confirmation grade: E0 observed effect artifact; E1 authenticated success response to the exploit action itself; E2 strong inference (config + reachable path shown separately); E3 error-message/behavioral inference only. Required on confirmed verdicts at/after harness 0.179.0. Only E0/E1 carry class-1 (override) power.\"\n        },\n        \"grade_rationale\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"description\": \"Why this grade: name the artifact (E0), the response (E1), or the inference chain (E2/E3).\"\n        },\n        \"severity_validation\": {\n          \"anyOf\": [\n            {\n              \"$ref\": \"#/$defs/severity_validation\"\n            },\n            {\n              \"type\": \"null\"\n            }\n          ]\n        },\n        \"chain_context\": {\n          \"anyOf\": [\n            {\n              \"type\": \"null\"\n            },\n            {\n              \"type\": \"object\",\n              \"required\": [\n                \"chain_id\",\n                \"chain_severity\",\n                \"role\"\n              ],\n              \"additionalProperties\": false,\n              \"properties\": {\n                \"chain_id\": {\n                  \"type\": \"string\"\n                },\n                \"chain_severity\": {\n                  \"type\": \"string\",\n                  \"enum\": [\n                    \"critical\",\n                    \"high\",\n                    \"medium\",\n                    \"low\"\n                  ]\n                },\n                \"role\": {\n                  \"type\": \"string\",\n                  \"minLength\": 5,\n                  \"description\": \"This finding's role in the demonstrated chain (e.g. 'initial access', 'privilege pivot', 'necessary link').\"\n                }\n              }\n            }\n          ],\n          \"description\": \"P9 chain amplification: membership in a demonstrated attack chain is severity evidence for the constituent — cite it in the severity proposal rationale.\"\n        }\n      }\n    },\n    \"attack_chain\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"chain_id\",\n        \"name\",\n        \"entry_point\",\n        \"terminal_asset\",\n        \"steps\",\n        \"verdict\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"chain_id\": {\n          \"type\": \"string\",\n          \"pattern\": \"^CHAIN-[0-9]{3,4}$\"\n        },\n        \"name\": {\n          \"type\": \"string\",\n          \"minLength\": 5\n        },\n        \"entry_point\": {\n          \"type\": \"string\"\n        },\n        \"terminal_asset\": {\n          \"type\": \"string\"\n        },\n        \"mitre_attack_refs\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\",\n            \"pattern\": \"^T[0-9]{4}(\\\\.[0-9]{3})?$\"\n          }\n        },\n        \"steps\": {\n          \"type\": \"array\",\n          \"minItems\": 2,\n          \"items\": {\n            \"$ref\": \"#/$defs/step_result\"\n          }\n        },\n        \"verdict\": {\n          \"$ref\": \"#/$defs/verdict\"\n        },\n        \"narrative\": {\n          \"type\": \"string\"\n        }\n      }\n    },\n    \"novel_finding\": {\n      \"allOf\": [\n        {\n          \"$ref\": \"report.schema.json#/$defs/finding\"\n        }\n      ],\n      \"properties\": {\n        \"discovery_method\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"recon-diff\",\n            \"pattern-probe\",\n            \"chain-emergent\",\n            \"ai-hypothesis\"\n          ]\n        },\n        \"chain_context\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"step_ref\": {\n          \"type\": \"string\"\n        }\n      },\n      \"unevaluatedProperties\": false\n    },\n    \"positive_control\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"name\",\n        \"kind\",\n        \"ok\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"name\": {\n          \"type\": \"string\",\n          \"minLength\": 3,\n          \"description\": \"What the control proves, e.g. 'same-client allowed-action succeeds' or 'planted canary secret readable'.\"\n        },\n        \"kind\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"must_succeed\",\n            \"must_deny\"\n          ],\n          \"description\": \"Expected behavior of the control action, independent of the finding under test.\"\n        },\n        \"verb\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"target\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"observed\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"description\": \"What actually happened (truncated transcript).\"\n        },\n        \"ok\": {\n          \"type\": \"boolean\",\n          \"description\": \"Control behaved as its kind expects — the assay is proven live.\"\n        }\n      },\n      \"description\": \"P1 assay-validity control paired with a refutation-capable probe: a same-path action that MUST succeed (or MUST be denied) independent of the finding. A refuted verdict without a passing control is quarantined by the soundness gate (missing-positive-control / failed-positive-control) for reports at/after harness 0.177.0.\"\n    },\n    \"differential_probe\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"neighbor_action\",\n        \"discriminated\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"neighbor_action\": {\n          \"type\": \"string\",\n          \"minLength\": 3,\n          \"description\": \"The paired action with a KNOWN-DIFFERENT expected outcome (e.g. same subject, a verb it IS allowed; same route, a role that IS authorized).\"\n        },\n        \"neighbor_expected\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"neighbor_observed\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"claimed_observed\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"description\": \"Outcome of the claimed action, for side-by-side comparison.\"\n        },\n        \"discriminated\": {\n          \"type\": \"boolean\",\n          \"description\": \"The two outcomes differed as expected — the oracle can tell allowed from denied. Identical outcomes = non-discriminating oracle = the refutation is unsound regardless of direction.\"\n        }\n      },\n      \"description\": \"P3 differential probe for authz claims: probe the claimed action AND a neighbor with a known-different expected outcome in the same session. Enforced by the soundness gate (non-discriminating-oracle).\"\n    },\n    \"severity_validation\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"demonstrated\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"demonstrated\": {\n          \"type\": \"object\",\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"attack_vector\": {\n              \"type\": [\n                \"string\",\n                \"null\"\n              ],\n              \"enum\": [\n                \"network\",\n                \"adjacent\",\n                \"local\",\n                \"physical\",\n                null\n              ],\n              \"description\": \"Vector the exploit ACTUALLY used.\"\n            },\n            \"privileges_required\": {\n              \"type\": [\n                \"string\",\n                \"null\"\n              ],\n              \"enum\": [\n                \"none\",\n                \"low\",\n                \"high\",\n                null\n              ],\n              \"description\": \"Privileges the probe identity ACTUALLY held at success time.\"\n            },\n            \"user_interaction\": {\n              \"type\": [\n                \"string\",\n                \"null\"\n              ],\n              \"enum\": [\n                \"none\",\n                \"required\",\n                null\n              ]\n            },\n            \"scope_crossed\": {\n              \"type\": [\n                \"boolean\",\n                \"null\"\n              ],\n              \"description\": \"Did the observed effect cross a namespace/tenant/node boundary (from the E0 artifact)?\"\n            },\n            \"impacts\": {\n              \"type\": [\n                \"array\",\n                \"null\"\n              ],\n              \"items\": {\n                \"type\": \"string\",\n                \"enum\": [\n                  \"confidentiality\",\n                  \"integrity\",\n                  \"availability\"\n                ]\n              },\n              \"description\": \"Impact axes actually evidenced (what was read/written/disrupted).\"\n            }\n          },\n          \"description\": \"CVSS v3.1 vector components as OBSERVATIONS from the live run.\"\n        },\n        \"claimed_vector\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"claimed_score\": {\n          \"type\": [\n            \"number\",\n            \"null\"\n          ]\n        },\n        \"demonstrated_score_estimate\": {\n          \"type\": [\n            \"number\",\n            \"null\"\n          ],\n          \"description\": \"Score recomputed with demonstrated components substituted into the claimed vector.\"\n        },\n        \"delta\": {\n          \"type\": [\n            \"number\",\n            \"null\"\n          ],\n          \"description\": \"demonstrated_score_estimate - claimed_score (signed).\"\n        },\n        \"proposal\": {\n          \"anyOf\": [\n            {\n              \"type\": \"null\"\n            },\n            {\n              \"type\": \"object\",\n              \"required\": [\n                \"severity\",\n                \"rationale\"\n              ],\n              \"additionalProperties\": false,\n              \"properties\": {\n                \"severity\": {\n                  \"type\": \"string\",\n                  \"enum\": [\n                    \"critical\",\n                    \"high\",\n                    \"medium\",\n                    \"low\",\n                    \"informational\"\n                  ]\n                },\n                \"rationale\": {\n                  \"type\": \"string\",\n                  \"minLength\": 20\n                }\n              }\n            }\n          ],\n          \"description\": \"Machine severity-adjustment PROPOSAL (emitted when |delta| >= 1.0 and the finding's evidence_grade is E0/E1/E2 — never E3). Routed to the countersign queue; machines never write disposition.severity.\"\n        }\n      },\n      \"description\": \"P9: severity validated from live observations. Downgrades are as valuable as upgrades (e.g. exploitation demonstrably required cluster-admin).\"\n    },\n    \"canonical_write\": {\n      \"description\": \"Opt-in contract for novel findings; historical validation is unchanged.\",\n      \"allOf\": [\n        { \"$ref\": \"#\" },\n        {\n          \"properties\": {\n            \"novel_findings\": {\n              \"items\": { \"$ref\": \"report.schema.json#/$defs/canonical_finding\" }\n            }\n          }\n        }\n      ]\n    },\n    \"replay_artifact\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"script_ref\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"script_ref\": {\n          \"type\": \"string\",\n          \"description\": \"Path (relative to the validation dir) of the self-contained replay script for this probe, e.g. artifacts/replay/S3.sh.\"\n        },\n        \"inputs_ref\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"description\": \"Path of recorded inputs (request bodies, manifests) the script consumes.\"\n        },\n        \"cluster_fingerprint_sha\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"description\": \"sha256 of the target-attestation.json fingerprint block at probe time — replay is only meaningful against a matching environment.\"\n        }\n      },\n      \"description\": \"P8: every probe records a replay artifact so a countersign human — or a later re-adjudication — can re-run the exact probe instead of trusting the transcript.\"\n    }\n  }\n}\n",
	"validation-authorization": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/validation-authorization.schema.json",
  "title": "Validation authorization",
  "description": "The reviewed plan digest and complete rules-of-engagement snapshot approved for a validation run. Possession of these bytes is not proof of identity; the service verifies the approver before publication and rechecks the scope at execution.",
  "type": "object",
  "required": ["artifact", "plan_id", "plan_digest", "approved_by", "approved_at", "mode", "scope", "skip_step_ids", "permit_destructive"],
  "additionalProperties": false,
  "properties": {
    "artifact": {"const": "validation-authorization"},
    "plan_id": {"type": "string", "format": "uuid"},
    "plan_digest": {"type": "string", "pattern": "^[0-9a-f]{64}$"},
    "approved_by": {"type": "string", "pattern": "^(human|machine):.+$"},
    "approved_at": {"type": "string", "format": "date-time"},
    "mode": {
      "allOf": [
        {"$ref": "validation.schema.json#/$defs/validation_metadata/properties/approval/properties/mode"},
        {"not": {"const": "dry-run"}}
      ]
    },
    "scope": {"$ref": "#/$defs/scope"},
    "skip_step_ids": {"type": "array", "uniqueItems": true, "items": {"type": "string", "minLength": 1}},
    "permit_destructive": {"type": "boolean"}
  },
  "allOf": [
    {
      "if": {"properties": {"mode": {"const": "auto"}}, "required": ["mode"]},
      "then": {
        "properties": {
          "scope": {
            "properties": {
              "environment": {"const": "lab"},
              "modes": {"contains": {"const": "explicit"}}
            }
          }
        }
      }
    }
  ],
  "$defs": {
    "scope": {
      "type": "object",
      "required": ["engagement", "authorized_by", "environment", "expires", "modes", "clusters", "off_limits", "browser", "containers", "images", "wasm_artifact_ids", "credential_probe_classes", "inventory_aws_profiles", "inventory_cluster_contexts"],
      "additionalProperties": false,
      "properties": {
        "engagement": {"type": ["string", "null"]},
        "authorized_by": {"type": ["string", "null"]},
        "environment": {"type": ["string", "null"]},
        "expires": {"oneOf": [{"type": "string", "format": "date"}, {"type": "null"}]},
        "modes": {
          "type": "array", "minItems": 1, "uniqueItems": true,
          "items": {
            "allOf": [
              {"$ref": "validation.schema.json#/$defs/scope_binding_mode"},
              {"not": {"anyOf": [{"const": "mixed"}, {"const": "none"}]}}
            ]
          }
        },
        "clusters": {"type": "array", "items": {"$ref": "#/$defs/cluster"}},
        "off_limits": {"type": "array", "items": {"$ref": "#/$defs/off_limit"}},
        "browser": {"type": "array", "items": {"$ref": "#/$defs/browser"}},
        "containers": {"$ref": "#/$defs/names"},
        "images": {"$ref": "#/$defs/names"},
        "wasm_artifact_ids": {"type": "array", "uniqueItems": true, "items": {"type": "string", "format": "uuid"}},
        "credential_probe_classes": {"$ref": "#/$defs/names"},
        "inventory_aws_profiles": {"$ref": "#/$defs/names"},
        "inventory_cluster_contexts": {"$ref": "#/$defs/names"}
      }
    },
    "cluster": {
      "type": "object",
      "required": ["context", "namespaces", "explicit_namespaces", "verbs_denied"],
      "additionalProperties": false,
      "properties": {
        "context": {"type": "string", "minLength": 1},
        "namespaces": {"$ref": "#/$defs/names"},
        "explicit_namespaces": {"$ref": "#/$defs/names"},
        "verbs_denied": {"$ref": "#/$defs/names"}
      }
    },
    "off_limit": {
      "type": "object",
      "additionalProperties": false,
      "properties": {
        "context": {"type": ["string", "null"]},
        "namespace": {"type": ["string", "null"]},
        "resource": {"type": ["string", "null"]},
        "verb": {"type": ["string", "null"]},
        "name": {"type": ["string", "null"]},
        "image": {"type": ["string", "null"]}
      }
    },
    "browser": {
      "type": "object",
      "required": ["victim_origin", "attacker_origins", "allowed_verbs"],
      "additionalProperties": false,
      "properties": {
        "victim_origin": {"type": "string", "minLength": 1},
        "attacker_origins": {"$ref": "#/$defs/names"},
        "allowed_verbs": {"$ref": "#/$defs/names"}
      }
    },
    "names": {"type": "array", "uniqueItems": true, "items": {"type": "string", "minLength": 1}}
  }
}
`,
	"validation-plan": `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/traust-contracts/artifacts_schema/v1/validation-plan.schema.json",
  "title": "Validation plan",
  "description": "Prepared validate-findings plan. Structural checks only: execution still requires scope, authorization and adapter classification checks.",
  "type": "object",
  "required": ["target", "scope_binding_mode", "chains", "steps"],
  "additionalProperties": false,
  "properties": {
    "target": {"type": "string", "minLength": 1},
    "scope_binding_mode": {"$ref": "validation.schema.json#/$defs/scope_binding_mode"},
    "engagement": {"type": ["string", "null"]},
    "chains": {"type": "array", "items": {"$ref": "#/$defs/chain"}},
    "steps": {"type": "array", "items": {"$ref": "#/$defs/step"}}
  },
  "$defs": {
    "chain": {
      "type": "object",
      "required": ["chain_id", "name", "entry_point", "terminal_asset", "finding_ids", "mitre_attack_refs", "narrative", "path"],
      "additionalProperties": false,
      "properties": {
        "chain_id": {"type": "string", "minLength": 1},
        "name": {"type": "string", "minLength": 1},
        "entry_point": {"type": "string", "minLength": 1},
        "terminal_asset": {"type": "string", "minLength": 1},
        "finding_ids": {"type": "array", "items": {"type": "string", "minLength": 1}},
        "mitre_attack_refs": {"type": "array", "items": {"type": "string", "minLength": 1}},
        "narrative": {"type": "string"},
        "path": {"type": "array", "items": {"type": "string", "minLength": 1}}
      }
    },
    "step": {
      "type": "object",
      "required": ["id", "technique", "adapter", "verb", "target", "classification"],
      "additionalProperties": false,
      "properties": {
        "id": {"type": "string", "minLength": 1},
        "technique": {"$ref": "validation.schema.json#/$defs/technique"},
        "adapter": {"$ref": "validation.schema.json#/$defs/adapter"},
        "verb": {"type": "string", "minLength": 1},
        "target": {"$ref": "#/$defs/step_target"},
        "finding_ref": {"type": "string", "minLength": 1},
        "chain_ref": {"type": "string", "minLength": 1},
        "novel_ref": {"type": "string", "minLength": 1},
        "payload": {"type": "string"},
        "cmd": {"type": "string"},
        "classification": {"$ref": "validation.schema.json#/$defs/classification"},
        "expected": {"type": "string"},
        "rollback": {"type": "string"},
        "preconditions": {"type": "array", "uniqueItems": true, "items": {"type": "string", "minLength": 1}},
        "skip": {"type": "string", "minLength": 1},
        "summary": {"type": "string"},
        "triage_was_needs_review": {"type": "boolean"}
      }
    },
    "step_target": {
      "type": "object",
      "additionalProperties": false,
      "properties": {
        "context": {"type": ["string", "null"]},
        "namespace": {"type": ["string", "null"]},
        "resource": {"type": ["string", "null"]},
        "name": {"type": ["string", "null"]},
        "image": {"type": ["string", "null"]},
        "artifact": {"type": ["string", "null"]},
        "runtime": {"type": "string"},
        "method": {"type": "string"},
        "path": {"type": "string"},
        "port": {"type": "integer", "minimum": 1, "maximum": 65535},
        "pqc": {"type": "boolean"}
      }
    }
  }
}
`,
	"verification":  "{\n  \"$schema\": \"https://json-schema.org/draft/2020-12/schema\",\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/verification.schema.json\",\n  \"title\": \"Remediation Verification Report Schema\",\n  \"description\": \"Schema for *-remediation-verification.json reports produced by the verify-remediation harness. Records the per-finding verdict of a targeted re-audit of a previous secure-code-audit report against a patched version of the repository, commit-level attribution of each fix, and any regressions introduced by the patch.\",\n  \"type\": \"object\",\n  \"required\": [\n    \"title\",\n    \"metadata\",\n    \"summary\",\n    \"verified_findings\",\n    \"regressions\",\n    \"commit_timeline\"\n  ],\n  \"additionalProperties\": false,\n  \"properties\": {\n    \"title\": {\n      \"type\": \"string\",\n      \"minLength\": 5,\n      \"pattern\": \"(?i)(verif|remediat)\"\n    },\n    \"metadata\": {\n      \"$ref\": \"#/$defs/verification_metadata\"\n    },\n    \"summary\": {\n      \"$ref\": \"#/$defs/summary\"\n    },\n    \"verified_findings\": {\n      \"description\": \"One entry per finding in the original audit report \\u2014 findings are never silently skipped.\",\n      \"type\": \"array\",\n      \"minItems\": 1,\n      \"items\": {\n        \"$ref\": \"#/$defs/verified_finding\"\n      }\n    },\n    \"regressions\": {\n      \"description\": \"New vulnerabilities introduced by the patch. Empty array when none were found.\",\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/regression\"\n      }\n    },\n    \"commit_timeline\": {\n      \"description\": \"Chronological list of all remediation commits across all findings.\",\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/timeline_entry\"\n      }\n    },\n    \"evidence\": {\n      \"description\": \"Typed base-versus-patch evidence gathered by this verification run, or carried forward from the remediation report that produced the patch. Shares one definition with the remediation family (remediation.schema.json#/$defs/patch_evidence) so a `proves` claim means the same thing on both sides and neither can drift. A targeted re-audit is analysis: it reads the two revisions but executes neither, so a verification report that carries no evidence item is making an analysis-only claim \\u2014 which is the honest default, not a defect. See the per-path evidence ceilings in the harness's docs/disposition-ledger.md \\u00a78a.\",\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"remediation.schema.json#/$defs/patch_evidence\"\n      }\n    },\n    \"recommendations\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"type\": \"string\",\n        \"minLength\": 10\n      }\n    },\n    \"review_surface_coverage\": {\n      \"$ref\": \"report.schema.json#/$defs/scan_review_surface_coverage\",\n      \"description\": \"Review-surface coverage of the re-reviewed code (the diff and its regression scan): one entry per surface. A surface the diff does not touch is out_of_scope. The cumulative builder merges this record into the repo's cumulative report (*-findings-current.json).\"\n    },\n    \"notes\": {\n      \"type\": \"string\",\n      \"description\": \"Free-form auditor/agent notes \\u2014 attribution caveats (force-pushed history, squash merges), scoping decisions for the regression scan, anything a human reviewer should know.\"\n    },\n    \"footer\": {\n      \"type\": \"string\"\n    }\n  },\n  \"$defs\": {\n    \"verdict\": {\n      \"type\": \"string\",\n      \"enum\": [\n        \"resolved\",\n        \"partially_resolved\",\n        \"unresolved\",\n        \"new_approach\",\n        \"regression\",\n        \"false_positive\",\n        \"risk_accepted\"\n      ]\n    },\n    \"verification_metadata\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"date\",\n        \"harness_version\",\n        \"original_report\",\n        \"original_commit\",\n        \"patched_commit\",\n        \"repository\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"date\": {\n          \"type\": \"string\",\n          \"format\": \"date\"\n        },\n        \"harness_version\": {\n          \"type\": \"string\",\n          \"pattern\": \"^\\\\d+\\\\.\\\\d+\\\\.\\\\d+(-[0-9a-f]{7,40})?$\"\n        },\n        \"original_report\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Path to the original *-security-audit.{json,md} report this verification is based on.\"\n        },\n        \"original_commit\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{40}$\",\n          \"description\": \"Full SHA of the commit the original audit findings were identified against.\"\n        },\n        \"patched_commit\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{40}$\",\n          \"description\": \"Full SHA of the patched HEAD that was verified (branch head, PR/MR head, or default-branch tip).\"\n        },\n        \"patched_ref\": {\n          \"type\": \"string\",\n          \"description\": \"Human-readable ref the patched commit came from, e.g. 'fix-branch', 'pull/42/head', 'merge-requests/17/head', or 'main'.\"\n        },\n        \"ref\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"$comment\": \"Optional (harness >= 0.122.0, branch-awareness Phase 0): the ORIGINAL audit's metadata.ref, restated verbatim so ref provenance survives into verification roll-ups. Not the patched ref \\u2014 that is patched_ref. Omit when the original report declares no ref.\"\n        },\n        \"ref_kind\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"branch\",\n            \"tag\",\n            \"default\",\n            \"stream\"\n          ],\n          \"$comment\": \"Optional (harness >= 0.122.0): the ORIGINAL audit's metadata.ref_kind, restated verbatim alongside ref. 'stream' (harness >= 0.140.0) = dist-git release stream, rpm profile.\"\n        },\n        \"repository\": {\n          \"type\": \"string\",\n          \"format\": \"uri\",\n          \"description\": \"Repository URL the patched code was obtained from.\"\n        },\n        \"scope\": {\n          \"type\": \"string\"\n        },\n        \"framework\": {\n          \"type\": \"string\",\n          \"description\": \"Framework list, same as the original audit.\"\n        },\n        \"auditor\": {\n          \"type\": \"string\"\n        },\n        \"additional\": {\n          \"type\": \"object\"\n        },\n        \"fix_repository\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"description\": \"When the whole verification ran against a fix repo (--fix-repo), its URL.\"\n        },\n        \"fix_ref\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"description\": \"Ref verified in fix_repository.\"\n        }\n      }\n    },\n    \"summary\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"total_findings\",\n        \"by_verdict\",\n        \"regressions\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"total_findings\": {\n          \"type\": \"integer\",\n          \"minimum\": 1\n        },\n        \"by_verdict\": {\n          \"type\": \"object\",\n          \"required\": [\n            \"resolved\",\n            \"partially_resolved\",\n            \"unresolved\",\n            \"new_approach\",\n            \"regression\",\n            \"false_positive\",\n            \"risk_accepted\"\n          ],\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"resolved\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"partially_resolved\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"unresolved\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"new_approach\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"regression\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"false_positive\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            },\n            \"risk_accepted\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            }\n          }\n        },\n        \"regressions\": {\n          \"type\": \"integer\",\n          \"minimum\": 0,\n          \"description\": \"Count of entries in the top-level regressions[] array.\"\n        },\n        \"prose\": {\n          \"type\": \"string\"\n        }\n      }\n    },\n    \"remediation_commit\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"sha\",\n        \"short_sha\",\n        \"date\",\n        \"author\",\n        \"subject\",\n        \"relevance\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"sha\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{40}$\"\n        },\n        \"short_sha\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{7,12}$\"\n        },\n        \"date\": {\n          \"type\": \"string\",\n          \"format\": \"date-time\"\n        },\n        \"author\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"subject\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"pr_number\": {\n          \"type\": [\n            \"integer\",\n            \"null\"\n          ],\n          \"minimum\": 1,\n          \"description\": \"GitHub PR number (#NN) or GitLab MR number (!NN) extracted from the commit message, null if none.\"\n        },\n        \"relevance\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"direct\",\n            \"supporting\",\n            \"partial\"\n          ]\n        }\n      }\n    },\n    \"verified_finding\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"original_id\",\n        \"original_title\",\n        \"original_severity\",\n        \"verdict\",\n        \"remediation_commits\",\n        \"unattributed\",\n        \"evidence\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"original_id\": {\n          \"type\": \"string\",\n          \"minLength\": 1,\n          \"description\": \"Finding ID from the original audit report.\"\n        },\n        \"original_title\": {\n          \"type\": \"string\",\n          \"minLength\": 5\n        },\n        \"original_severity\": {\n          \"$ref\": \"report.schema.json#/$defs/severity_level\"\n        },\n        \"verdict\": {\n          \"$ref\": \"#/$defs/verdict\"\n        },\n        \"remediation_commits\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/$defs/remediation_commit\"\n          }\n        },\n        \"unattributed\": {\n          \"type\": \"boolean\",\n          \"description\": \"True when no commit could be identified that addresses this finding (unresolved, fix predates the audit, disposition without code change, or commit history unavailable).\"\n        },\n        \"evidence\": {\n          \"type\": \"object\",\n          \"required\": [\n            \"explanation\",\n            \"framework_reference\"\n          ],\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"original_code\": {\n              \"type\": \"string\",\n              \"description\": \"Code at the finding location in the original commit.\"\n            },\n            \"patched_code\": {\n              \"type\": \"string\",\n              \"description\": \"Code at the finding location in the patched commit, or confirmation the file/section was removed.\"\n            },\n            \"explanation\": {\n              \"type\": \"string\",\n              \"minLength\": 30,\n              \"description\": \"How the change addresses (or fails to address) the root cause.\"\n            },\n            \"framework_reference\": {\n              \"type\": \"string\",\n              \"minLength\": 3,\n              \"description\": \"Same CWE, K-ID, CIS section, STIG ID, SLSA level, or PEACH parameter as the original finding.\"\n            }\n          }\n        },\n        \"disposition_rationale\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"description\": \"Required for false_positive and risk_accepted verdicts: who made the determination and why (link to the triage record, MR discussion, or risk-acceptance decision).\"\n        },\n        \"residual_risk\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"description\": \"Remaining risk when partially_resolved or new_approach, null otherwise.\"\n        },\n        \"residual_severity\": {\n          \"anyOf\": [\n            {\n              \"$ref\": \"report.schema.json#/$defs/severity_level\"\n            },\n            {\n              \"type\": \"null\"\n            }\n          ],\n          \"description\": \"Recalculated severity when partially_resolved, null otherwise.\"\n        },\n        \"cross_repo\": {\n          \"$ref\": \"#/$defs/cross_repo\"\n        }\n      }\n    },\n    \"regression\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"id\",\n        \"title\",\n        \"severity\",\n        \"cwes\",\n        \"locations\",\n        \"description\",\n        \"remediation\",\n        \"introduced_by\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"id\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[A-Z][A-Z0-9_]{0,23}-[a-f0-9]{7}-REG-\\\\d{3}$\",\n          \"description\": \"Canonical regression ID: {REPO_SLUG}-{PATCHED_SHORTSHA}-REG-{NNN}.\"\n        },\n        \"title\": {\n          \"type\": \"string\",\n          \"minLength\": 5\n        },\n        \"severity\": {\n          \"$ref\": \"report.schema.json#/$defs/severity_level\"\n        },\n        \"cwes\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"items\": {\n            \"type\": \"string\",\n            \"pattern\": \"^CWE-\\\\d{1,5}$\"\n          }\n        },\n        \"cvss\": {\n          \"type\": \"object\",\n          \"required\": [\n            \"score\",\n            \"vector\"\n          ],\n          \"additionalProperties\": false,\n          \"properties\": {\n            \"score\": {\n              \"type\": \"number\",\n              \"minimum\": 0.0,\n              \"maximum\": 10.0\n            },\n            \"vector\": {\n              \"type\": \"string\"\n            }\n          }\n        },\n        \"locations\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"items\": {\n            \"$ref\": \"report.schema.json#/$defs/location\"\n          }\n        },\n        \"description\": {\n          \"type\": \"string\",\n          \"minLength\": 50\n        },\n        \"remediation\": {\n          \"type\": \"string\",\n          \"minLength\": 10\n        },\n        \"evidence\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"report.schema.json#/$defs/evidence_block\"\n          }\n        },\n        \"attack_pattern\": {\n          \"type\": \"string\"\n        },\n        \"category\": {\n          \"type\": \"string\"\n        },\n        \"pqc_classification\": {\n          \"$ref\": \"report.schema.json#/$defs/finding/properties/pqc_classification\"\n        },\n        \"remediation_effort\": {\n          \"$ref\": \"report.schema.json#/$defs/finding/properties/remediation_effort\"\n        },\n        \"blocked_external\": {\n          \"$ref\": \"report.schema.json#/$defs/finding/properties/blocked_external\"\n        },\n        \"introduced_by\": {\n          \"type\": \"string\",\n          \"minLength\": 7,\n          \"description\": \"Commit SHA or PR/MR reference that introduced this regression.\"\n        },\n        \"routed_id\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[A-Z][A-Z0-9_]{0,23}-[a-f0-9]{7}-\\\\d{3}$\",\n          \"description\": \"Campaign finding ID this regression was routed to in the repo's baseline audit report (harness >= 0.196.0, the regression router): {REPO_SLUG}-{PATCHED_SHORTSHA}-{NNN}, numbering continuing where the repo's existing findings at that sha leave off. The REG id stays here as provenance; the routed finding carries the REG id in its source_findings. Absent on regressions not yet routed.\"\n        },\n        \"fingerprint\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{64}$\",\n          \"description\": \"Deterministic cross-scan identity (identity module fingerprint). Optional here -- no artifact of this kind carries one today -- but DECLARED, because additionalProperties is false and an undeclared identity field makes every stamped artifact invalid the moment a producer starts emitting it. That has happened twice: report.schema.json omitted fingerprint_algo for 16 days, and the cloud-config schema declared neither, so stamping the policy findings invalidated nearly every report in one command.\"\n        },\n        \"fingerprint_algo\": {\n          \"type\": \"string\",\n          \"pattern\": \"^v\\\\d+$\",\n          \"description\": \"Algorithm version of `fingerprint` (the identity module ALGO_VERSION). Declared for the same reason as `fingerprint` above.\"\n        }\n      }\n    },\n    \"timeline_entry\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"sha\",\n        \"full_sha\",\n        \"date\",\n        \"author\",\n        \"subject\",\n        \"addresses\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"sha\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{7,12}$\"\n        },\n        \"full_sha\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{40}$\"\n        },\n        \"date\": {\n          \"type\": \"string\",\n          \"format\": \"date-time\"\n        },\n        \"author\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"subject\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"pr_number\": {\n          \"type\": [\n            \"integer\",\n            \"null\"\n          ],\n          \"minimum\": 1\n        },\n        \"addresses\": {\n          \"type\": \"array\",\n          \"minItems\": 1,\n          \"items\": {\n            \"type\": \"string\"\n          },\n          \"description\": \"Original finding IDs this commit addresses.\"\n        }\n      }\n    },\n    \"cross_repo\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"fix_repo\",\n        \"propagation\"\n      ],\n      \"additionalProperties\": false,\n      \"properties\": {\n        \"fix_repo\": {\n          \"type\": \"string\",\n          \"minLength\": 8,\n          \"description\": \"URL of the repository the fix landed in (differs from metadata.repository).\"\n        },\n        \"fix_ref\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"description\": \"Commit SHA or ref of the fix in fix_repo.\"\n        },\n        \"propagation\": {\n          \"type\": \"string\",\n          \"enum\": [\n            \"consumed\",\n            \"pending\",\n            \"module_absent\",\n            \"not_applicable\"\n          ],\n          \"description\": \"Whether the ORIGINAL repo has consumed the fix: consumed = lockfile/vendor/SBOM at or past the fixed version; pending = still on a vulnerable version; module_absent = dependency no longer present (removal is also a fix \\u2014 verifier judges); not_applicable = Shape 2/3 (fix legitimately lives only in fix_repo, e.g. config/policy repo).\"\n        },\n        \"propagation_evidence\": {\n          \"type\": [\n            \"string\",\n            \"null\"\n          ],\n          \"description\": \"What was checked: file + pinned version, vendor/modules.txt line, SBOM component version (from the fix-propagation checker).\"\n        }\n      },\n      \"description\": \"Present when the fix lives in a DIFFERENT repository than the finding. The finding's identity stays in the original repo; this block is evidence attached to it. Hard rule (validator-enforced): propagation 'pending' forbids verdict 'resolved' \\u2014 upstream merging a fix does not resolve a finding the product still ships vulnerable.\"\n    }\n  }\n}\n",
	"vuln-findings": "{\n  \"$id\": \"https://example.com/traust-contracts/artifacts_schema/v1/vuln-findings.schema.json\",\n  \"$schema\": \"https://json-schema.org/draft/2020-12/schema\",\n  \"title\": \"Vuln-Scan Findings Schema\",\n  \"description\": \"Output contract of the vuln-scan skill (<repo>-vuln-findings.json, harness >= 0.38.0). Shares vocabulary with report.schema.json: the severity_level enum is identical and finding ids use the canonical {REPO_SLUG}-{SHORTSHA}-{NNN} campaign pattern. Findings are UNVERIFIED candidates ('claimed' in findings-lifecycle terms) awaiting /triage; known_findings are candidates deduped against the target's *-security-audit.json baseline. Legacy VULN-FINDINGS.json files (pre-0.38.0) do not conform and are not auto-detected against this schema.\",\n  \"type\": \"object\",\n  \"required\": [\n    \"target\",\n    \"scanned_at\",\n    \"focus_areas\",\n    \"metadata\",\n    \"findings\",\n    \"summary\"\n  ],\n  \"properties\": {\n    \"target\": {\n      \"type\": \"string\",\n      \"minLength\": 1\n    },\n    \"scanned_at\": {\n      \"description\": \"ISO-8601 timestamp of the scan.\",\n      \"type\": \"string\",\n      \"minLength\": 10\n    },\n    \"focus_areas\": {\n      \"type\": \"array\",\n      \"minItems\": 1,\n      \"items\": {\n        \"type\": \"string\",\n        \"minLength\": 1\n      }\n    },\n    \"metadata\": {\n      \"$ref\": \"#/$defs/metadata\"\n    },\n    \"findings\": {\n      \"type\": \"array\",\n      \"minItems\": 0,\n      \"$comment\": \"Empty is valid: a clean sweep (or one where every candidate matched the baseline) yields zero new findings.\",\n      \"items\": {\n        \"$ref\": \"#/$defs/finding\"\n      }\n    },\n    \"known_findings\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/$defs/known_finding\"\n      }\n    },\n    \"review_surface_coverage\": {\n      \"$ref\": \"report.schema.json#/$defs/scan_review_surface_coverage\",\n      \"description\": \"Review-surface coverage of the scanned code (the whole tree, or the diff for a --diff scan): one entry per surface. A surface the diff does not touch is out_of_scope. The cumulative builder merges this record into the repo's cumulative report (*-findings-current.json).\"\n    },\n    \"summary\": {\n      \"$ref\": \"#/$defs/summary\"\n    }\n  },\n  \"additionalProperties\": false,\n  \"allOf\": [\n    {\n      \"$comment\": \"out_of_scope is for a diff-scoped scan only. A record that answers any surface out_of_scope must carry the --diff stamp, metadata.additional.coverage_diff; a full scan records what it found on every surface.\",\n      \"if\": {\n        \"required\": [\"review_surface_coverage\"],\n        \"properties\": {\n          \"review_surface_coverage\": {\n            \"not\": {\n              \"additionalProperties\": {\n                \"not\": {\n                  \"required\": [\"status\"],\n                  \"properties\": { \"status\": { \"const\": \"out_of_scope\" } }\n                }\n              }\n            }\n          }\n        }\n      },\n      \"then\": {\n        \"required\": [\"metadata\"],\n        \"properties\": {\n          \"metadata\": {\n            \"required\": [\"additional\"],\n            \"properties\": { \"additional\": { \"required\": [\"coverage_diff\"] } }\n          }\n        }\n      }\n    }\n  ],\n  \"$defs\": {\n    \"severity_level\": {\n      \"type\": \"string\",\n      \"enum\": [\n        \"critical\",\n        \"high\",\n        \"medium\",\n        \"low\",\n        \"informational\"\n      ]\n    },\n    \"metadata\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"repo\",\n        \"repo_slug\",\n        \"scanned_ref\",\n        \"harness_version\",\n        \"baseline\",\n        \"baseline_findings\"\n      ],\n      \"properties\": {\n        \"repo\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"repo_slug\": {\n          \"description\": \"Repo name uppercased, [^A-Z0-9] -> '_', truncated to 24 chars \\u2014 the id prefix.\",\n          \"type\": \"string\",\n          \"pattern\": \"^[A-Z][A-Z0-9_]{0,23}$\"\n        },\n        \"scanned_ref\": {\n          \"description\": \"Short SHA of the scanned commit (7 lowercase hex), or '0000000 (not a git checkout)'.\",\n          \"type\": \"string\",\n          \"pattern\": \"^[a-f0-9]{7}( \\\\(not a git checkout\\\\))?$\"\n        },\n        \"harness_version\": {\n          \"description\": \"Harness semver + short SHA at scan time, e.g. 0.38.0-de5a090.\",\n          \"type\": \"string\",\n          \"pattern\": \"^\\\\d+\\\\.\\\\d+\\\\.\\\\d+(-[a-f0-9]{7,12})?$\"\n        },\n        \"baseline\": {\n          \"description\": \"Path of the matched <repo>-security-audit.json, or null when the target has no audit baseline.\",\n          \"type\": [\n            \"string\",\n            \"null\"\n          ]\n        },\n        \"baseline_findings\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"additional\": {\n          \"description\": \"Extension bag (same convention as report.schema.json): deterministic_steps recording, the --diff mode coverage_diff stamp, and the spend calibration stamp. Undeclared keys stay allowed.\",\n          \"type\": \"object\",\n          \"properties\": {\n            \"coverage_diff\": {\n              \"type\": \"object\",\n              \"description\": \"Present exactly when the scan was diff-scoped (/vuln-scan --diff): what the diff was taken against and what of it the scan covered. A full scan omits it. A review_surface_coverage entry may be out_of_scope only when this is present.\",\n              \"required\": [\"changed_files_covered\", \"callers_covered\", \"refused_or_skipped\", \"baseline_report\", \"anchor\", \"resolution_source\"],\n              \"properties\": {\n                \"changed_files_covered\": {\"type\": \"array\", \"items\": {\"type\": \"string\"}, \"description\": \"Changed first-party files the scan reviewed.\"},\n                \"callers_covered\": {\"type\": \"array\", \"items\": {\"type\": \"string\"}, \"description\": \"Callers of changed code the scan reviewed.\"},\n                \"refused_or_skipped\": {\"type\": \"array\", \"items\": {\"type\": \"string\"}, \"description\": \"Each changed file or step not covered, as '<file or step>: <reason>'.\"},\n                \"baseline_report\": {\"type\": \"string\", \"minLength\": 1, \"description\": \"The audit report the diff was resolved against.\"},\n                \"anchor\": {\"type\": \"string\", \"minLength\": 1, \"description\": \"The commit the diff starts from (the baseline's metadata.commit).\"},\n                \"resolution_source\": {\"type\": \"string\", \"minLength\": 1, \"description\": \"How the baseline and anchor were chosen: auto (the resolver) or override (--since / --baseline).\"}\n              }\n            }\n          }\n        }\n      },\n      \"additionalProperties\": false\n    },\n    \"finding\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"id\",\n        \"file\",\n        \"line\",\n        \"category\",\n        \"severity\",\n        \"confidence\",\n        \"title\",\n        \"description\",\n        \"recommendation\"\n      ],\n      \"properties\": {\n        \"id\": {\n          \"$comment\": \"Canonical campaign id only \\u2014 this artifact type postdates the legacy id scheme.\",\n          \"type\": \"string\",\n          \"pattern\": \"^[A-Z][A-Z0-9_]{0,23}-[a-f0-9]{7}-\\\\d{3}$\"\n        },\n        \"scanner_ref\": {\n          \"description\": \"The review subagent's working id (F-<focus>-<n>), kept for traceability.\",\n          \"type\": \"string\"\n        },\n        \"file\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"line\": {\n          \"$comment\": \"Null when the scanner could only cite the enclosing function (the description must say so).\",\n          \"type\": [\n            \"integer\",\n            \"null\"\n          ],\n          \"minimum\": 1\n        },\n        \"category\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"cwe\": {\n          \"description\": \"Primary CWE for the category; feeds the baseline fold-in's required cwes[].\",\n          \"type\": \"string\",\n          \"pattern\": \"^CWE-\\\\d{1,5}$\"\n        },\n        \"severity\": {\n          \"$ref\": \"#/$defs/severity_level\"\n        },\n        \"confidence\": {\n          \"type\": \"number\",\n          \"minimum\": 0.0,\n          \"maximum\": 1.0\n        },\n        \"title\": {\n          \"type\": \"string\",\n          \"minLength\": 5\n        },\n        \"description\": {\n          \"type\": \"string\",\n          \"minLength\": 50\n        },\n        \"exploit_scenario\": {\n          \"type\": \"string\"\n        },\n        \"recommendation\": {\n          \"type\": \"string\",\n          \"minLength\": 10\n        },\n        \"confidence_reason\": {\n          \"type\": \"string\"\n        },\n        \"fingerprint\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{64}$\",\n          \"description\": \"Deterministic cross-scan identity (identity module fingerprint). Optional here -- no artifact of this kind carries one today -- but DECLARED, because additionalProperties is false and an undeclared identity field makes every stamped artifact invalid the moment a producer starts emitting it. That has happened twice: report.schema.json omitted fingerprint_algo for 16 days, and the cloud-config schema declared neither, so stamping the policy findings invalidated nearly every report in one command.\"\n        },\n        \"fingerprint_algo\": {\n          \"type\": \"string\",\n          \"pattern\": \"^v\\\\d+$\",\n          \"description\": \"Algorithm version of `fingerprint` (the identity module ALGO_VERSION). Declared for the same reason as `fingerprint` above.\"\n        }\n      },\n      \"additionalProperties\": false\n    },\n    \"known_finding\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"title\",\n        \"matches_baseline_id\"\n      ],\n      \"properties\": {\n        \"title\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        },\n        \"matches_baseline_id\": {\n          \"description\": \"The baseline audit finding this candidate duplicates.\",\n          \"type\": \"string\",\n          \"minLength\": 1\n        }\n      },\n      \"additionalProperties\": false\n    },\n    \"summary\": {\n      \"type\": \"object\",\n      \"required\": [\n        \"total\",\n        \"critical\",\n        \"high\",\n        \"medium\",\n        \"low\",\n        \"informational\",\n        \"known\",\n        \"low_confidence\"\n      ],\n      \"properties\": {\n        \"total\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"critical\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"high\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"medium\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"low\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"informational\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"known\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        },\n        \"low_confidence\": {\n          \"type\": \"integer\",\n          \"minimum\": 0\n        }\n      },\n      \"additionalProperties\": false\n    }\n  }\n}\n",
}
