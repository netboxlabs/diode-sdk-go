# Entity Examples Generation Specification

## Overview

This document specifies the automated generation of complete, runnable example code for all Diode entity types. The system parses protobuf definitions and generates language-specific examples (initially Go, extensible to Python) that demonstrate SDK usage.

## Goals

1. **Comprehensive Coverage**: Generate examples for all 96+ entity types automatically
2. **Programmatic**: No manual code writing - all examples generated from templates
3. **Maintainable**: Update proto → regenerate examples → all examples stay current
4. **Testable**: All examples must compile and optionally run against live Diode server
5. **Educational**: Examples demonstrate best practices and common patterns
6. **Extensible**: Support multiple languages (Go first, Python later)

## Architecture

### System Components

```
┌─────────────────────────────────────────────────────────────────┐
│                    Example Generator System                      │
├─────────────────────────────────────────────────────────────────┤
│                                                                   │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐      │
│  │ Proto Parser │───▶│  Analyzer    │───▶│  Generator   │      │
│  └──────────────┘    └──────────────┘    └──────────────┘      │
│         │                    │                    │              │
│         ▼                    ▼                    ▼              │
│  Extract entity        Determine field     Apply templates      │
│  definitions           requirements        Generate code        │
│  and metadata          and examples                             │
│                                                                   │
└─────────────────────────────────────────────────────────────────┘

Input: ingester.proto (4,223 lines)
Output: 96+ runnable example programs
```

### Data Flow

```
ingester.proto
    ↓
Proto Parser (Go AST or protoreflect)
    ↓
Entity Models (in-memory structures)
    ↓
Example Analyzer (determine required fields, relationships)
    ↓
Template Engine (text/template or similar)
    ↓
Generated Examples (main.go + go.mod per entity)
    ↓
Validation (go build to verify compilation)
```

## Proto Parsing Strategy

### Option 1: Custom Parser (Recommended for Phase 1)

Build on existing `protograph/parser.go` logic:

```go
package generator

import (
    "bufio"
    "regexp"
    "strings"
)

// EntityDefinition represents a parsed entity from the proto file
type EntityDefinition struct {
    Name           string
    ProtoName      string // e.g., "Device"
    OneofFieldName string // e.g., "device" in Entity.oneof
    Fields         []FieldDefinition
    Comments       []string
}

// FieldDefinition represents a field in an entity
type FieldDefinition struct {
    Name         string
    ProtoType    string // e.g., "string", "DeviceType", "repeated Tag"
    GoType       string // e.g., "string", "*diodepb.DeviceType", "[]*diodepb.Tag"
    Number       int
    Optional     bool
    Repeated     bool
    IsReference  bool   // true if field is another entity type
    IsEnum       bool
    EnumValues   []string
    Constraints  FieldConstraints
    Comments     []string
}

// FieldConstraints represents validation rules
type FieldConstraints struct {
    Required   bool
    MinValue   *int64
    MaxValue   *int64
    InValues   []string // enum values
    Pattern    string   // regex pattern
}

// ParseProtoFile extracts entity definitions from ingester.proto
func ParseProtoFile(path string) ([]EntityDefinition, error) {
    // Implementation:
    // 1. Find the Entity message oneof block
    // 2. Extract all entity type names from oneof
    // 3. For each entity type, find its message definition
    // 4. Parse fields, types, constraints, comments
    // 5. Build EntityDefinition structures
    return nil, nil
}
```

**Key Parsing Logic**:
1. Identify entity types from `Entity.oneof` block (lines 29-100 in proto)
2. For each entity type, locate its message definition
3. Extract field information:
   - Field name and number
   - Type (scalar, message, repeated)
   - Optional/required status (proto3 optional keyword)
   - Validation constraints (from validate.rules)
   - Comments (field-level documentation)

### Option 2: Use protoreflect (More Robust)

Use Go's `google.golang.org/protobuf/reflect/protoreflect` for proper parsing:

```go
import (
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/reflect/protodesc"
    "google.golang.org/protobuf/reflect/protoreflect"
    "google.golang.org/protobuf/types/descriptorpb"
)

// This approach uses compiled proto descriptors
// More robust but requires proto compilation first
func ParseProtoWithReflect(descriptorPath string) ([]EntityDefinition, error) {
    // Load proto descriptor
    // Use reflection to inspect messages and fields
    // Build EntityDefinition structures
    return nil, nil
}
```

**Trade-offs**:
- ✅ More accurate, handles all proto syntax correctly
- ✅ Gets validation rules automatically
- ❌ Requires proto compilation step
- ❌ More complex setup

**Recommendation**: Start with Option 1 (custom parser) for simplicity, can upgrade to Option 2 later.

## Example Analysis

### Determining Required Fields

Not all fields must be populated in examples. Analyze field requirements:

```go
// ExampleFieldRequirement determines if field should be in example
type ExampleFieldRequirement int

const (
    FieldOptional   ExampleFieldRequirement = iota // Can be omitted
    FieldRecommended                                // Should show in example
    FieldRequired                                   // Must be in example
)

// AnalyzeFieldRequirement determines field requirement level
func AnalyzeFieldRequirement(field FieldDefinition, entity EntityDefinition) ExampleFieldRequirement {
    // Rules:
    // 1. If field has validate.rules.required = true → Required
    // 2. If field name is "name" or "site" or "device" → Required (primary identifiers)
    // 3. If field is optional and not commonly used → Optional
    // 4. Everything else → Recommended

    if field.Constraints.Required {
        return FieldRequired
    }

    // Common primary fields
    primaryFields := []string{"name", "site", "device", "prefix", "address"}
    for _, pf := range primaryFields {
        if field.Name == pf {
            return FieldRequired
        }
    }

    // Optional fields that add value to examples
    educationalFields := []string{"description", "tags", "status"}
    for _, ef := range educationalFields {
        if field.Name == ef {
            return FieldRecommended
        }
    }

    return FieldOptional
}
```

### Example Complexity Levels

Generate multiple example variations per entity:

```go
type ExampleVariation string

const (
    MinimalExample     ExampleVariation = "minimal"     // Only required fields
    TypicalExample     ExampleVariation = "typical"     // Required + recommended
    ComprehensiveExample ExampleVariation = "comprehensive" // All fields with examples
)
```

**For documentation**: Use `TypicalExample` (shows common usage without overwhelming)

**For testing**: Generate all three variations to verify API coverage

## Go Template Design

### Directory Structure

```
tools/gen-entity-docs/
├── templates/
│   ├── go/
│   │   ├── main.go.tmpl           # Main template for runnable program
│   │   ├── go.mod.tmpl            # go.mod template
│   │   ├── entity.go.tmpl         # Entity construction
│   │   ├── field_scalar.go.tmpl   # Scalar field assignment
│   │   ├── field_ref.go.tmpl      # Reference field (nested entity)
│   │   ├── field_repeated.go.tmpl # Repeated field (slice)
│   │   └── field_enum.go.tmpl     # Enum field
│   └── python/
│       ├── main.py.tmpl           # (Future) Python example
│       └── entity.py.tmpl
├── data/
│   ├── example_values.yaml        # Sample data for examples
│   └── field_patterns.yaml        # Patterns for generating realistic data
└── generator.go
```

### Main Template (main.go.tmpl)

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/netboxlabs/diode-sdk-go/diode"
	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

func main() {
	// Configuration from environment
	target := os.Getenv("DIODE_TARGET")
	if target == "" {
		target = "grpc://localhost:8081"
	}
	apiKey := os.Getenv("DIODE_API_KEY")
	if apiKey == "" {
		log.Fatal("DIODE_API_KEY environment variable required")
	}

	// Create Diode client
	client, err := diode.NewClient(
		diode.WithTarget(target),
		diode.WithAPIKey(apiKey),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	{{template "entity_construction" .}}

	// Wrap in Entity message
	entity := &diodepb.Entity{
		Entity: &diodepb.Entity_{{.OneofFieldName}}{
			{{.ProtoName}}: {{.VariableName}},
		},
	}

	// Ingest to Diode
	ctx := context.Background()
	response, err := client.Ingest(ctx, []*diodepb.Entity{entity})
	if err != nil {
		log.Fatalf("Failed to ingest: %v", err)
	}

	log.Printf("Successfully ingested {{.ProtoName}}: %s", response.String())
}
```

### Entity Construction Template (entity.go.tmpl)

```go
{{define "entity_construction"}}
	// Create {{.ProtoName}} entity
	{{.VariableName}} := &diodepb.{{.ProtoName}}{
{{- range .Fields}}
{{- if .ShouldInclude}}
		{{template "field_assignment" .}}
{{- end}}
{{- end}}
	}
{{end}}
```

### Field Assignment Templates

**Scalar Field (field_scalar.go.tmpl)**:
```go
{{define "field_assignment_scalar"}}
		{{.GoFieldName}}: {{.ExampleValue}},
{{end}}
```

**Reference Field (field_ref.go.tmpl)**:
```go
{{define "field_assignment_reference"}}
		{{.GoFieldName}}: &diodepb.{{.ReferenceType}}{
{{- range .ReferenceFields}}
			{{.GoFieldName}}: {{.ExampleValue}},
{{- end}}
		},
{{end}}
```

**Repeated Field (field_repeated.go.tmpl)**:
```go
{{define "field_assignment_repeated"}}
		{{.GoFieldName}}: []*diodepb.{{.ElementType}}{
			{
{{- range .ElementFields}}
				{{.GoFieldName}}: {{.ExampleValue}},
{{- end}}
			},
		},
{{end}}
```

**Enum Field (field_enum.go.tmpl)**:
```go
{{define "field_assignment_enum"}}
		{{.GoFieldName}}: "{{.ExampleEnumValue}}",  // Options: {{.AllEnumValues}}
{{end}}
```

## Example Value Generation

### Realistic Sample Data

Define realistic example values for common fields:

```yaml
# data/example_values.yaml

# By field name
field_values:
  name:
    device: ["router01", "switch01", "fw01"]
    site: ["NYC-DC1", "LON-DC2", "SFO-HQ"]
    manufacturer: ["Cisco", "Juniper", "Arista"]

  address:
    ipv4: ["10.0.1.1", "192.168.1.10", "172.16.0.1"]
    ipv6: ["2001:db8::1", "fd00::1"]

  prefix:
    ipv4: ["10.0.0.0/24", "192.168.1.0/24"]
    ipv6: ["2001:db8::/64"]

  status:
    device: ["active", "planned", "offline"]

  description:
    generic: ["Production environment", "Management network", "Test infrastructure"]

# By entity type
entity_specific:
  Device:
    name: "router01"
    serial: "SN123456789"
    asset_tag: "ASSET-001"

  IPAddress:
    address: "10.0.1.1/24"
    dns_name: "router01.example.com"

  Interface:
    name: "GigabitEthernet0/0"
    type: "1000base-t"

  Site:
    name: "NYC-DC1"
    slug: "nyc-dc1"
    facility: "Equinix NY5"
```

### Value Generation Algorithm

```go
// ExampleValueGenerator generates realistic example values
type ExampleValueGenerator struct {
    values map[string][]string
    used   map[string]int // Track usage for variation
}

// GenerateValue creates example value for a field
func (g *ExampleValueGenerator) GenerateValue(field FieldDefinition, entityType string) string {
    // Priority:
    // 1. Entity-specific value from YAML
    // 2. Field name pattern match from YAML
    // 3. Type-based default value
    // 4. Constraint-based generation (if enum, use first value)

    // Check entity-specific values
    if val := g.getEntitySpecificValue(entityType, field.Name); val != "" {
        return val
    }

    // Check field name patterns
    if val := g.getFieldPatternValue(field.Name); val != "" {
        return val
    }

    // Type-based defaults
    switch field.GoType {
    case "string":
        return fmt.Sprintf("\"%s\"", g.generateStringValue(field))
    case "int32", "int64":
        return g.generateIntValue(field)
    case "bool":
        return "true"
    default:
        // Reference type - needs nested generation
        return g.generateReferenceValue(field)
    }
}
```

## Complete Generation Process

### Step-by-Step Algorithm

```go
// GenerateExamples is the main entry point
func GenerateExamples(protoPath, outputDir string, lang Language) error {
    // 1. Parse proto file
    entities, err := ParseProtoFile(protoPath)
    if err != nil {
        return fmt.Errorf("parse proto: %w", err)
    }

    // 2. Load example values
    values, err := LoadExampleValues("data/example_values.yaml")
    if err != nil {
        return fmt.Errorf("load values: %w", err)
    }

    // 3. For each entity, generate example
    for _, entity := range entities {
        // Analyze field requirements
        fields := analyzeFields(entity, values)

        // Build template data
        data := buildTemplateData(entity, fields, values)

        // Generate code from template
        code, err := generateFromTemplate(lang, data)
        if err != nil {
            return fmt.Errorf("generate %s: %w", entity.Name, err)
        }

        // Write to output directory
        entityDir := filepath.Join(outputDir, toSnakeCase(entity.Name))
        if err := writeExample(entityDir, code); err != nil {
            return fmt.Errorf("write %s: %w", entity.Name, err)
        }

        // Validate (compile check)
        if err := validateExample(entityDir, lang); err != nil {
            return fmt.Errorf("validate %s: %w", entity.Name, err)
        }
    }

    // 4. Generate index/README
    if err := generateExampleIndex(outputDir, entities); err != nil {
        return fmt.Errorf("generate index: %w", err)
    }

    return nil
}
```

### Template Data Structure

```go
// TemplateData contains all data needed for example generation
type TemplateData struct {
    // Entity information
    EntityName     string // "Device"
    VariableName   string // "device"
    OneofFieldName string // "device" in Entity.oneof
    ProtoName      string // "Device"

    // Fields to include in example
    Fields []FieldTemplateData

    // Metadata
    PackageName   string   // "main"
    ImportPath    string   // "github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
    Description   string   // From proto comments

    // Example configuration
    Variation     ExampleVariation
    IncludeComments bool
}

// FieldTemplateData represents a field in the template
type FieldTemplateData struct {
    Name           string // "device_type"
    GoFieldName    string // "DeviceType"
    GoType         string // "*diodepb.DeviceType"
    ExampleValue   string // Generated value

    // Field classification
    IsScalar       bool
    IsReference    bool
    IsRepeated     bool
    IsEnum         bool

    // For reference types
    ReferenceType  string
    ReferenceFields []FieldTemplateData

    // For enums
    EnumValues     []string
    ExampleEnumValue string

    // For repeated
    ElementType    string

    // Documentation
    Comment        string

    // Include in example?
    ShouldInclude  bool
}
```

## Comprehensive Coverage Strategy

### Ensuring All Entities Are Generated

```go
// VerifyCompleteness checks that all entities have examples
func VerifyCompleteness(protoPath, examplesDir string) error {
    // 1. Parse proto to get all entity types
    entities, err := ParseProtoFile(protoPath)
    if err != nil {
        return err
    }

    // 2. List generated example directories
    dirs, err := ioutil.ReadDir(examplesDir)
    if err != nil {
        return err
    }

    // 3. Build set of generated entities
    generated := make(map[string]bool)
    for _, dir := range dirs {
        if dir.IsDir() {
            generated[dir.Name()] = true
        }
    }

    // 4. Check for missing entities
    var missing []string
    for _, entity := range entities {
        dirName := toSnakeCase(entity.Name)
        if !generated[dirName] {
            missing = append(missing, entity.Name)
        }
    }

    if len(missing) > 0 {
        return fmt.Errorf("missing examples for: %v", missing)
    }

    return nil
}
```

### Entity Coverage Report

Generate a coverage report showing example completeness:

```go
// GenerateCoverageReport creates a report of example coverage
func GenerateCoverageReport(protoPath, examplesDir, outputPath string) error {
    entities, _ := ParseProtoFile(protoPath)

    report := &CoverageReport{
        TotalEntities: len(entities),
        Generated:     0,
        Missing:       []string{},
        Details:       []EntityCoverage{},
    }

    for _, entity := range entities {
        dirName := toSnakeCase(entity.Name)
        entityDir := filepath.Join(examplesDir, dirName)

        coverage := EntityCoverage{
            Name: entity.Name,
        }

        // Check if example exists
        if _, err := os.Stat(filepath.Join(entityDir, "main.go")); err == nil {
            coverage.HasExample = true
            report.Generated++

            // Check if it compiles
            if err := runGoBuild(entityDir); err == nil {
                coverage.Compiles = true
            } else {
                coverage.CompileError = err.Error()
            }
        } else {
            report.Missing = append(report.Missing, entity.Name)
        }

        // Calculate field coverage
        coverage.FieldCoverage = calculateFieldCoverage(entity, entityDir)

        report.Details = append(report.Details, coverage)
    }

    // Write report as markdown
    return writeMarkdownReport(report, outputPath)
}
```

## Advanced Features

### Field Pattern Recognition

Automatically detect common field patterns:

```go
// FieldPattern represents a recognized pattern
type FieldPattern int

const (
    PatternUnknown FieldPattern = iota
    PatternIdentifier           // name, slug, etc.
    PatternNetwork              // IP address, prefix, etc.
    PatternPhysical             // serial, asset_tag, etc.
    PatternRelationship         // device, site, etc.
    PatternMetadata             // description, comments, etc.
    PatternTags                 // tags, labels, etc.
)

// DetectFieldPattern identifies common patterns
func DetectFieldPattern(field FieldDefinition) FieldPattern {
    name := strings.ToLower(field.Name)

    identifiers := []string{"name", "slug", "id"}
    for _, id := range identifiers {
        if name == id {
            return PatternIdentifier
        }
    }

    network := []string{"address", "prefix", "ip", "subnet", "vlan"}
    for _, net := range network {
        if strings.Contains(name, net) {
            return PatternNetwork
        }
    }

    // ... more patterns

    return PatternUnknown
}
```

### Relationship Graph Generation

Build a graph of entity relationships:

```go
// EntityGraph represents relationships between entities
type EntityGraph struct {
    Entities map[string]*EntityNode
}

type EntityNode struct {
    Name         string
    References   []Reference  // Entities this refers to
    ReferencedBy []Reference  // Entities that refer to this
}

type Reference struct {
    EntityName string
    FieldName  string
    Required   bool
}

// BuildEntityGraph creates relationship graph from proto
func BuildEntityGraph(entities []EntityDefinition) *EntityGraph {
    graph := &EntityGraph{
        Entities: make(map[string]*EntityNode),
    }

    // First pass: create nodes
    for _, entity := range entities {
        graph.Entities[entity.Name] = &EntityNode{
            Name: entity.Name,
        }
    }

    // Second pass: build relationships
    for _, entity := range entities {
        node := graph.Entities[entity.Name]

        for _, field := range entity.Fields {
            if field.IsReference {
                refType := extractReferenceType(field.ProtoType)
                node.References = append(node.References, Reference{
                    EntityName: refType,
                    FieldName:  field.Name,
                    Required:   field.Constraints.Required,
                })

                // Add reverse reference
                if refNode, ok := graph.Entities[refType]; ok {
                    refNode.ReferencedBy = append(refNode.ReferencedBy, Reference{
                        EntityName: entity.Name,
                        FieldName:  field.Name,
                        Required:   field.Constraints.Required,
                    })
                }
            }
        }
    }

    return graph
}
```

### Dependency-Ordered Example Generation

Generate examples in dependency order (referenced entities first):

```go
// GenerateInDependencyOrder generates examples respecting dependencies
func GenerateInDependencyOrder(entities []EntityDefinition) ([]string, error) {
    graph := BuildEntityGraph(entities)

    // Topological sort to get dependency order
    order := []string{}
    visited := make(map[string]bool)

    var visit func(name string) error
    visit = func(name string) error {
        if visited[name] {
            return nil
        }

        node := graph.Entities[name]

        // Visit dependencies first (entities this references)
        for _, ref := range node.References {
            if err := visit(ref.EntityName); err != nil {
                return err
            }
        }

        visited[name] = true
        order = append(order, name)
        return nil
    }

    // Visit all entities
    for name := range graph.Entities {
        if err := visit(name); err != nil {
            return nil, err
        }
    }

    return order, nil
}
```

## Python Extension (Future)

### Python Template Structure

```python
# templates/python/main.py.tmpl
"""
Example: {{.EntityName}}

{{.Description}}
"""

from netboxlabs.diode.sdk import DiodeClient
from netboxlabs.diode.sdk.ingester import (
    {{.ProtoName}},
    Entity,
{{- range .ImportTypes}}
    {{.}},
{{- end}}
)

def main():
    # Configuration
    client = DiodeClient(
        target=os.getenv("DIODE_TARGET", "grpc://localhost:8081"),
        api_key=os.getenv("DIODE_API_KEY"),
    )

    # Create {{.ProtoName}} entity
    {{.VariableName}} = {{.ProtoName}}(
{{- range .Fields}}
{{- if .ShouldInclude}}
        {{.Name}}={{.PythonExampleValue}},
{{- end}}
{{- end}}
    )

    # Wrap in Entity
    entity = Entity({{.OneofFieldName}}={{.VariableName}})

    # Ingest to Diode
    response = client.ingest([entity])
    print(f"Successfully ingested {{.ProtoName}}: {response}")

if __name__ == "__main__":
    main()
```

### Cross-Language Consistency

Ensure Go and Python examples demonstrate the same patterns:

```go
// ExampleSpec defines language-agnostic example specification
type ExampleSpec struct {
    EntityName   string
    Fields       []FieldSpec
    Variation    ExampleVariation
}

// FieldSpec defines how a field should appear in examples
type FieldSpec struct {
    Name         string
    ExampleValue interface{} // Language-agnostic value
    Required     bool
}

// GenerateFromSpec creates language-specific code from spec
func GenerateFromSpec(spec ExampleSpec, lang Language) (string, error) {
    switch lang {
    case LanguageGo:
        return generateGoFromSpec(spec)
    case LanguagePython:
        return generatePythonFromSpec(spec)
    default:
        return "", fmt.Errorf("unsupported language: %v", lang)
    }
}
```

## Testing Strategy

### Example Validation

```go
// ValidateExample ensures example is correct
func ValidateExample(dir string, lang Language) error {
    switch lang {
    case LanguageGo:
        // 1. Check files exist
        mainGo := filepath.Join(dir, "main.go")
        goMod := filepath.Join(dir, "go.mod")
        if !fileExists(mainGo) || !fileExists(goMod) {
            return errors.New("missing required files")
        }

        // 2. Run go build
        cmd := exec.Command("go", "build", "-o", "/dev/null", ".")
        cmd.Dir = dir
        if output, err := cmd.CombinedOutput(); err != nil {
            return fmt.Errorf("build failed: %s", output)
        }

        // 3. Run go vet
        cmd = exec.Command("go", "vet", "./...")
        cmd.Dir = dir
        if output, err := cmd.CombinedOutput(); err != nil {
            return fmt.Errorf("vet failed: %s", output)
        }

        // 4. (Optional) Run golangci-lint
        cmd = exec.Command("golangci-lint", "run")
        cmd.Dir = dir
        if output, err := cmd.CombinedOutput(); err != nil {
            // Don't fail on lint errors, just warn
            log.Printf("lint warnings for %s: %s", dir, output)
        }

        return nil

    case LanguagePython:
        // Python validation
        return validatePythonExample(dir)
    }

    return fmt.Errorf("unsupported language: %v", lang)
}
```

### Integration Testing

Optionally run examples against live Diode server:

```go
// IntegrationTest runs example against live server
func IntegrationTest(dir string, diodeTarget, apiKey string) error {
    cmd := exec.Command("go", "run", "main.go")
    cmd.Dir = dir
    cmd.Env = append(os.Environ(),
        fmt.Sprintf("DIODE_TARGET=%s", diodeTarget),
        fmt.Sprintf("DIODE_API_KEY=%s", apiKey),
    )

    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("run failed: %s", output)
    }

    // Check for success message
    if !strings.Contains(string(output), "Successfully ingested") {
        return fmt.Errorf("unexpected output: %s", output)
    }

    return nil
}
```

## Makefile Integration

```makefile
# diode-sdk-go/Makefile

.PHONY: gen-examples
gen-examples:
	@echo "Generating entity examples..."
	go run ./tools/gen-entity-docs \
		-proto=../diode/diode-proto/diode/v1/ingester.proto \
		-examples-dir=./docs/examples \
		-lang=go \
		-variation=typical

.PHONY: gen-examples-all
gen-examples-all:
	@echo "Generating all example variations..."
	for variation in minimal typical comprehensive; do \
		go run ./tools/gen-entity-docs \
			-proto=../diode/diode-proto/diode/v1/ingester.proto \
			-examples-dir=./docs/examples-$$variation \
			-lang=go \
			-variation=$$variation; \
	done

.PHONY: test-examples
test-examples:
	@echo "Validating all examples compile..."
	@for dir in docs/examples/*/; do \
		echo "Building $$dir..."; \
		(cd "$$dir" && go build -o /dev/null .) || exit 1; \
	done
	@echo "✓ All examples compile successfully"

.PHONY: test-examples-integration
test-examples-integration:
	@echo "Running examples against Diode server..."
	@if [ -z "$$DIODE_API_KEY" ]; then \
		echo "Error: DIODE_API_KEY not set"; \
		exit 1; \
	fi
	@for dir in docs/examples/*/; do \
		echo "Running $$dir..."; \
		(cd "$$dir" && go run main.go) || echo "⚠ $$dir failed"; \
	done

.PHONY: verify-coverage
verify-coverage:
	@echo "Checking example coverage..."
	go run ./tools/gen-entity-docs \
		-proto=../diode/diode-proto/diode/v1/ingester.proto \
		-examples-dir=./docs/examples \
		-verify-only

.PHONY: coverage-report
coverage-report:
	@echo "Generating coverage report..."
	go run ./tools/gen-entity-docs \
		-proto=../diode/diode-proto/diode/v1/ingester.proto \
		-examples-dir=./docs/examples \
		-report=./docs/COVERAGE.md
```

## Success Criteria

### Phase 1: Go Examples
- ✅ All 96+ entity types have generated examples
- ✅ All examples compile without errors
- ✅ Examples demonstrate typical usage (not just minimal)
- ✅ Examples include proper error handling
- ✅ Consistent code style across all examples
- ✅ CI validates examples on every proto change

### Phase 2: Python Examples
- ✅ All entity types have Python examples
- ✅ Python examples mirror Go example patterns
- ✅ Update Python SDK docs with generated examples
- ✅ Automated testing for Python examples

### Phase 3: Advanced Features
- ✅ Multi-entity examples (relationships)
- ✅ Batch ingestion examples
- ✅ Error handling examples
- ✅ Retry and resilience patterns
- ✅ Custom field examples

## References

- Proto parsing: `diode-server/cmd/protograph/parser.go`
- Go templates: https://pkg.go.dev/text/template
- Protocol Buffers: https://protobuf.dev/
- Existing manual examples: `diode-sdk-python/docs/entities.md`
