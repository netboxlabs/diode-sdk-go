# Entity Documentation Generation Plan for diode-sdk-go

## Understanding of the Task

The goal is to **automate the generation of Diode entity documentation** for the Go SDK, similar to the existing [Python SDK entities.md](https://github.com/netboxlabs/diode-sdk-python/blob/5dae6c3222fb99124c3e22bd312ebc4a537ccf9f/docs/entities.md) which is currently manually maintained.

### Current State

- **Python SDK**: Has manually created `docs/entities.md` with 96+ entity types documented with attributes and usage examples
- **Go SDK**: No equivalent documentation exists
- **Problem**: Manual documentation is time-consuming and becomes outdated as entities evolve
- **Solution**: Automate documentation generation from the source of truth (protobuf definitions)

### Architecture Overview

```
NetBox Django Models
    ↓
Django REST Framework Serializers
    ↓
diode-drf-extract Tool (SerializerTraverser)
    ↓
Protobuf Schema (ingester.proto) ← **SOURCE OF TRUTH**
    ↓
┌─────────────┬──────────────┬─────────────────┐
│             │              │                 │
buf codegen   buf codegen    protograph        New: doc generator
    ↓             ↓              ↓                    ↓
Python SDK    Go SDK      entity_mappings.go     entities.md
```

## Existing Code That Analyzes/Translates Django Marshalers

### ⚠️ Critical Finding: Metadata Loss in Current Implementation

**The current extraction process DOES NOT preserve field documentation from Django serializers.**

The `SerializerTraverser` in the diode-drf-extract tool extracts **structural metadata** (types, constraints, relationships) but **ignores semantic metadata** (field descriptions, help_text, labels, docstrings). This means:

- ✅ **Extracted**: Field names, types, required/optional status, enums, min/max values, relationships
- ❌ **Lost**: Field descriptions, help_text, labels, docstrings, semantic meaning

**Evidence**: The generated `ingester.proto` file (4,223 lines) contains:
- Message definitions with correct field types and numbers
- Validation rules (enums, constraints)
- NO field-level documentation comments
- NO explanation of what fields represent

Example from `ingester.proto`:
```protobuf
message Device {
  optional string name = 1;
  DeviceType device_type = 2;
  DeviceRole role = 3;
  // ... no comments explaining what these fields are for
}
```

### 1. Django Serializer Extraction Tool
**Location**: `/Users/pstuart/CODE/github.com/netboxlabs/eng-observability-skunkworks/diode-drf-extract/`

**Key File**: `diode_drf_extract/management/commands/extract_protobuf.py`
- Analyzes Django REST Framework serializers from NetBox
- Uses `SerializerTraverser` to extract field definitions, types, and relationships
- Generates protobuf definitions, Python wrappers, Go wrappers, and utility files
- Configuration: `output/ingester.yaml`

**What It Extracts** (from `visitor.py`):
```python
@dataclass
class FieldInfo:
    name: str
    field_type: str
    is_required: bool
    is_read_only: bool
    enum_values: list[str]
    min_value: Optional[int]
    max_value: Optional[int]
    # ... but NO help_text, label, or description
```

**What It Doesn't Extract**:
- `field.help_text` - Human-readable field descriptions
- `field.label` - Display labels
- Docstrings from serializer classes
- Any semantic documentation

**Command**:
```bash
cd /Users/pstuart/CODE/github.com/netboxlabs/eng-observability-skunkworks/diode-drf-extract
make NETBOX_VERSION=v4.5.0 docker-compose-extract
make copy-artifacts
```

### 2. Protobuf Parser for Go Code Generation
**Location**: `/Users/pstuart/CODE/github.com/netboxlabs/diode/diode-server/cmd/protograph/`

**Key Files**:
- `parser.go` (370 lines) - Parses `.proto` files and extracts entity types and fields
- `generator.go` (141 lines) - Generates Go code from parsed entities
- `templates.go` (122 lines) - Go text templates for code generation
- `main.go` - Entry point

**Capabilities**:
- Parses protobuf syntax to extract message definitions
- Identifies entity types, fields, field types, and relationships
- Generates Go switch statements and helper functions
- **Output**: `diode-server/gen/protograph/entity_mappings.go`

**Command**:
```bash
cd /Users/pstuart/CODE/github.com/netboxlabs/diode/diode-server
make gen-entity-mappings
```

### 3. Django Transformation Layer
**Location**: `/Users/pstuart/CODE/github.com/netboxlabs/diode-netbox-plugin/netbox_diode_plugin/api/`

**Key Files**:
- `transformer.py` (720 lines) - Converts protobuf JSON to Django ORM-compatible dictionaries
- `matcher.py` (885 lines) - Object matching and fingerprinting logic
- `supported_models.py` (85 lines) - Extracts supported models and field metadata

These files show how entities are **used** but don't generate documentation.

### 4. Existing Documentation Generator
**Location**: `/Users/pstuart/CODE/github.com/netboxlabs/diode-netbox-plugin/netbox_diode_plugin/management/commands/generate_matching_docs.py`

**Purpose**: Generates markdown documentation for entity matching criteria
**Output**: `docs/matching-criteria-documentation.md`

This demonstrates the **pattern** for automated documentation generation from code introspection.

## Source of Truth Analysis

### Django Serializers vs Proto Files: Who's Really in Charge?

**Initial Source of Truth**: Django REST Framework Serializers in NetBox
- Define the canonical data model structure
- Contain field definitions, types, validation rules
- Have access to `help_text`, `label`, docstrings (but NOT currently extracted)
- Reflect NetBox's internal data model

**Maintained Source of Truth**: Proto Files (`ingester.proto`)
- Generated from serializers, then **merged with previous versions**
- Field numbers are **preserved** across generations (critical for backward compatibility)
- Comments and custom options are maintained through merge process
- Acts as the stable, version-controlled wire format specification

**The Merge Process** preserves:
- Field numbers from previous proto versions
- Deprecated field markers
- Message structure and hierarchy
- Any manually added documentation (though none currently exists)

**Conclusion**: Django serializers are the **initial** source of truth, but proto files become the **maintained** source due to the merge process. However, **neither currently contains field documentation**.

## Plan: Generate diode-sdk-go Documentation

### Three Possible Approaches

Given the metadata loss issue, we have three options:

### Approach A: Direct Django Serializer Documentation (Richest Data)

**Extract documentation directly from Django serializers** before the proto generation step.

**Advantages**:
- Access to ALL metadata: `help_text`, `label`, docstrings
- Most accurate, human-readable field descriptions
- Can include NetBox-specific context and examples
- No dependency on proto file preservation

**Disadvantages**:
- Requires running in Django/NetBox environment
- Need to sync with NetBox version updates
- Separate from Go SDK code generation pipeline
- More complex setup (Docker, NetBox dependencies)

**Implementation**:
1. Create Django management command in diode-drf-extract
2. Use SerializerTraverser to discover all serializers
3. Extract `field.help_text`, `field.label` from each serializer field
4. Generate markdown directly from serializer metadata
5. Include Go SDK usage examples (from templates)

### Approach B: Enhanced Proto with Documentation (Hybrid)

**Enhance the extraction tool to preserve documentation in proto files**, then generate docs from enriched proto.

**Advantages**:
- Proto file becomes comprehensive source of truth
- Documentation preserved through merge process
- Works with existing protograph parser
- Standard protobuf documentation comments

**Disadvantages**:
- Requires modifying diode-drf-extract tool
- Need to coordinate with upstream team
- One-time enhancement needed before documentation generation
- Requires re-running extraction for all NetBox versions

**Implementation**:
1. Modify `visitor.py` to extract `field.help_text` and add to `FieldInfo`
2. Modify `protobuf.py` to write field comments into generated proto
3. Ensure `proto_merger.py` preserves field comments
4. Re-run extraction to generate documented proto file
5. Generate markdown from enriched proto using protograph

### Approach C: Protobuf-Only Documentation (Current State)

**Generate documentation from existing proto files** without field descriptions.

**Advantages**:
- Simplest implementation (extend protograph)
- No dependencies on Django or extraction tool
- Works with current proto files immediately
- Part of Go SDK build pipeline

**Disadvantages**:
- ❌ **No field descriptions** - Users won't know what fields mean
- Only structural information (types, required/optional)
- Less helpful than Python SDK's manual documentation
- May require manual augmentation

**Implementation**:
1. Extend protograph to parse proto and generate markdown
2. Include field types, required status, validation rules
3. Generate Go usage examples from templates
4. **Manually add field descriptions later** (defeats automation goal)

## Recommended Approach: Approach B (Enhanced Proto with Documentation)

**Why**: This provides the best long-term solution:

1. **Single Source of Truth**: Proto file becomes comprehensive, documented source
2. **Automated Updates**: Documentation regenerates when NetBox updates
3. **Consistent Format**: Same proto used for Go SDK, Python SDK, server, and docs
4. **Standard Practice**: Protobuf comments are industry standard for API documentation
5. **Preservation**: Merge process maintains documentation across versions

**Two-Phase Implementation**:

**Phase 1**: Enhance Extraction Tool (Upstream Work)
- Modify diode-drf-extract to capture help_text
- Write field comments to proto file
- Coordinate with NetBoxLabs team

**Phase 2**: Generate Documentation (This Project)
- Parse enriched proto file
- Generate markdown with field descriptions
- Add Go usage examples

### Implementation: Extend protograph Tool

**Advantages**:
- Existing protobuf parser already in `protograph/parser.go`
- Already parses entity types, fields, and relationships
- Mature Go codebase
- Can share parsing logic with entity mappings generation

**Implementation Steps**:

1. **Create Standalone Documentation Generator**

   **Why separate from protograph**:
   - Different output artifacts (markdown + Go examples vs Go code)
   - Different templates and logic
   - Simpler to maintain and test independently
   - Can reuse protograph's parser as a library if needed

   **New tool structure**:
   ```
   tools/gen-entity-docs/
   ├── main.go              # Entry point
   ├── parser.go            # Proto parser (or import from protograph)
   ├── doc_generator.go     # Markdown generation
   ├── example_generator.go # Go example generation
   ├── templates/
   │   ├── doc.md.tmpl     # Entity documentation template
   │   ├── example.go.tmpl # Go example template
   │   └── go.mod.tmpl     # go.mod template for examples
   └── go.mod
   ```

2. **Command-Line Interface**
   ```go
   var (
       protoFile   = flag.String("proto", "", "path to .proto file")
       outputFile  = flag.String("output", "", "output markdown file path")
       examplesDir = flag.String("examples-dir", "", "directory for example code")
       sdkImport   = flag.String("sdk-import", "github.com/netboxlabs/diode-sdk-go", "SDK import path")
       sdkVersion  = flag.String("sdk-version", "", "SDK version for examples (default: use replace directive)")
   )
   ```

   **Usage**:
   ```bash
   go run ./tools/gen-entity-docs \
       -proto=../diode/diode-proto/diode/v1/ingester.proto \
       -output=./docs/entities.md \
       -examples-dir=./docs/examples
   ```

3. **Generate Documentation Format and Testable Examples**

   **Proposed Directory Structure**:
   ```
   diode-sdk-go/
   ├── docs/
   │   ├── entities.md                    # Main documentation (generated)
   │   └── examples/
   │       ├── README.md                  # How to run examples
   │       ├── device/
   │       │   ├── main.go               # Complete, runnable Device example
   │       │   └── go.mod                # Module file
   │       ├── interface/
   │       │   ├── main.go               # Complete, runnable Interface example
   │       │   └── go.mod                # Module file
   │       ├── ip_address/
   │       │   ├── main.go               # Complete, runnable IPAddress example
   │       │   └── go.mod                # Module file
   │       └── ... (96+ entity directories)
   ├── tools/
   │   └── gen-entity-docs/
   │       ├── main.go                    # Documentation generator
   │       ├── templates/
   │       │   ├── doc.md.tmpl           # Markdown template
   │       │   └── example.go.tmpl       # Go example template
   │       └── testdata/
   │           └── example_entities.txt   # Test data
   ```

   **Generated Example Structure** (`docs/examples/device/main.go`):
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
       // Get configuration from environment
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

       // Create Device entity
       device := &diodepb.Device{
           Name: "router01",
           DeviceType: &diodepb.DeviceType{
               Name: "Cisco Catalyst 9300",
           },
           Role: &diodepb.Role{
               Name: "Router",
           },
           Site: &diodepb.Site{
               Name: "NYC-DC1",
           },
       }

       // Wrap in Entity message
       entity := &diodepb.Entity{
           Entity: &diodepb.Entity_Device{Device: device},
       }

       // Ingest to Diode
       ctx := context.Background()
       response, err := client.Ingest(ctx, []*diodepb.Entity{entity})
       if err != nil {
           log.Fatalf("Failed to ingest: %v", err)
       }

       log.Printf("Successfully ingested Device: %s", response.String())
   }
   ```

   **Generated Documentation Format** (`docs/entities.md`):
   ```markdown
   ## Device

   **Description**: Represents a physical device in the network infrastructure.

   **Attributes**:
   - `name` (string, required) - The unique name of the device
   - `device_type` (DeviceTypeRef) - Reference to the device type
   - `role` (RoleRef) - Device role assignment
   - `site` (SiteRef, required) - Site where device is located
   - `status` (string) - Operational status
   - `tags` ([]TagRef) - Associated tags
   - `custom_fields` (map[string]CustomFieldValue) - Custom field values

   **Usage Example**:

   See the complete, runnable example at [docs/examples/device/](../examples/device/)

   ```go
   // Create Device entity
   device := &diodepb.Device{
       Name: "router01",
       DeviceType: &diodepb.DeviceType{Name: "Cisco Catalyst 9300"},
       Role: &diodepb.Role{Name: "Router"},
       Site: &diodepb.Site{Name: "NYC-DC1"},
   }

   // Wrap in Entity and ingest
   entity := &diodepb.Entity{Entity: &diodepb.Entity_Device{Device: device}}
   response, err := client.Ingest(ctx, []*diodepb.Entity{entity})
   ```

   Run the example:
   ```bash
   cd docs/examples/device
   export DIODE_API_KEY="your-api-key"
   go run main.go
   ```
   ```

   **Benefits of This Approach**:
   - ✅ Each example is a complete, runnable Go program
   - ✅ Examples can be built and tested in CI to verify correctness
   - ✅ Each example has its own directory (one `main()` per package)
   - ✅ Users can copy entire example directories to get started quickly
   - ✅ Examples include proper error handling and context
   - ✅ Documentation references working code, not inline snippets

4. **Makefile Integration**
   Add to `diode-sdk-go/Makefile`:
   ```makefile
   .PHONY: gen-entity-docs
   gen-entity-docs:
       @echo "Generating entity documentation and examples..."
       go run ./tools/gen-entity-docs \
           -proto=../diode/diode-proto/diode/v1/ingester.proto \
           -output=./docs/entities.md \
           -examples-dir=./docs/examples

   .PHONY: test-examples
   test-examples:
       @echo "Building all entity examples..."
       @for dir in docs/examples/*/; do \
           echo "Building $$dir..."; \
           (cd "$$dir" && go build -o /dev/null .) || exit 1; \
       done
       @echo "All examples build successfully!"

   .PHONY: docs
   docs: gen-entity-docs test-examples
       @echo "Documentation and examples generated and verified!"
   ```

5. **CI/CD Integration**

   **GitHub Actions Workflow** (`.github/workflows/docs.yml`):
   ```yaml
   name: Documentation

   on:
     pull_request:
       paths:
         - 'docs/**'
         - 'diode/**'
         - 'tools/gen-entity-docs/**'
     push:
       branches: [main, develop]

   jobs:
     generate-and-test:
       runs-on: ubuntu-latest
       steps:
         - uses: actions/checkout@v3

         - uses: actions/setup-go@v4
           with:
             go-version: '1.21'

         - name: Generate documentation
           run: make gen-entity-docs

         - name: Build all examples
           run: make test-examples

         - name: Check for uncommitted changes
           run: |
             git diff --exit-code docs/entities.md || \
               (echo "Documentation is out of date. Run 'make gen-entity-docs'" && exit 1)
             git diff --exit-code docs/examples/ || \
               (echo "Examples are out of date. Run 'make gen-entity-docs'" && exit 1)
   ```

   **CI Checks**:
   - Generate documentation and examples
   - Build all example programs to verify they compile
   - Validate documentation is up-to-date (fail if uncommitted changes)
   - Run example tests (if test files are added later)

### Alternative: Direct Django Serializer Documentation

If enhancing the proto extraction is not feasible, we could generate documentation directly from Django serializers:

**Advantages**:
- Access to ALL field metadata (help_text, labels, docstrings)
- Most accurate and complete documentation
- Independent from proto generation

**Disadvantages**:
- Requires Django/NetBox environment
- More complex setup
- Separate from Go SDK pipeline

**Implementation**:
1. Create Django management command in diode-drf-extract
2. Use SerializerTraverser to discover serializers
3. Extract help_text from each field
4. Generate markdown with Go usage examples
5. Copy to diode-sdk-go repository

## Recommended Implementation Plan

### Decision Point: Choose Implementation Path

**Before starting implementation, decide which approach to take:**

1. **Option 1 (Recommended)**: Enhance extraction tool first, then generate docs from enriched proto
   - Timeline: 2-3 weeks (1 week upstream enhancement + 1-2 weeks doc generation)
   - Quality: Excellent (automated, complete, maintainable)
   - Coordination: Requires NetBoxLabs team collaboration

2. **Option 2 (Quick Win)**: Generate docs from current proto, manually add descriptions
   - Timeline: 1 week
   - Quality: Good structure, but missing field descriptions initially
   - Coordination: None needed

3. **Option 3 (Alternative)**: Generate directly from Django serializers
   - Timeline: 2 weeks
   - Quality: Excellent (complete metadata)
   - Coordination: Independent, but requires Django environment setup

### Phase 1: Prototype with Standalone Generator (Week 1)

**Assumes current proto files without field documentation**

1. **Create proto parser** to extract:
   - Entity message names and descriptions (from comments, if any)
   - Field names, types, and required/optional status
   - Field documentation from comments (currently none, but parser should support)
   - Reference types (identify *Ref message types)
   - Validation rules (enums, min/max values)
   - Optional: reuse/adapt protograph's parser.go

2. **Create templates**:
   - **Markdown template** following Python entities.md format:
     - Header with entity name
     - Description from protobuf comments (or placeholder)
     - Attributes table with name, type, required status, validation rules
     - **Note**: Field descriptions will be "TBD" until proto is enhanced
     - Reference to example code directory
     - Minimal inline code snippet

   - **Go example template** for complete, runnable programs:
     - Package main with imports
     - main() function with proper error handling
     - Environment variable configuration (DIODE_TARGET, DIODE_API_KEY)
     - Entity creation with required fields
     - Client setup and ingestion
     - Logging and error handling

3. **Generate initial documentation and examples** for 2-3 entities as proof of concept:
   - Device (common, well-understood entity)
   - IPAddress (shows different field types)
   - Interface (shows relationships)

   **Deliverables**:
   - `docs/entities.md` with 3 entities documented
   - `docs/examples/device/main.go` - buildable example
   - `docs/examples/ip_address/main.go` - buildable example
   - `docs/examples/interface/main.go` - buildable example
   - Makefile target to build all examples
   - CI workflow to validate examples build

4. **Validate proof of concept**:
   - Run `make test-examples` to verify all examples compile
   - Review generated documentation format
   - Present to team for feedback on approach
   - Identify what's missing (field descriptions) and prioritize enhancement

### Phase 2: Full Implementation (Week 2)

**Path A: If proto enhancement is approved**
1. **Coordinate with diode-drf-extract team**:
   - Share requirements for field documentation extraction
   - Review `visitor.py` modifications
   - Test enhanced proto generation
   - Verify field comments in generated proto

2. **Update protograph parser**:
   - Parse field-level comments from enhanced proto
   - Extract validation rules and constraints
   - Build entity relationship graph

3. **Implement full documentation generator**:
   - Generate docs for all 96+ entity types with descriptions
   - Add proper Go code examples with imports
   - Include entity relationships documentation
   - Add table of contents

**Path B: If using current proto**
1. **Implement full documentation generator**:
   - Generate `docs/entities.md` for all 96+ entity types
   - Generate `docs/examples/<entity>/main.go` for all entities
   - Include field types, validation rules, relationships
   - Add "Description: TBD" placeholders for field descriptions
   - Create proper table of contents with links
   - Add examples index in `docs/examples/README.md`

2. **Generate all example code**:
   - 96+ directories under `docs/examples/`
   - Each with complete, runnable `main.go`
   - Each with `go.mod` using replace directive for local SDK
   - Consistent structure and error handling across all examples
   - Environment variable configuration for flexibility

3. **Add Makefile targets**:
   - `make gen-entity-docs` - Generate documentation and examples
   - `make test-examples` - Build all examples to verify compilation
   - `make docs` - Full documentation generation + verification
   - Integrate with existing `make generate` target

4. **Update CI/CD**:
   - Add documentation generation to GitHub Actions
   - Build all examples in CI to catch compilation errors
   - Validate documentation is up-to-date in PRs
   - Fail CI if generated files have uncommitted changes

5. **Create companion enhancement plan**:
   - Document what's needed in proto files (field descriptions)
   - Create proposal for diode-drf-extract enhancement
   - Estimate effort for manual documentation augmentation
   - Open issue tracking proto enhancement work

### Phase 3: Enhancement (Week 3)

5. **Add advanced features**:
   - Cross-references between entity types
   - Field constraints and validation rules displayed prominently
   - Common usage patterns and best practices
   - Troubleshooting section

6. **Documentation website integration**:
   - Integrate with existing docs site
   - Add search functionality
   - Add interactive examples

7. **Manual augmentation** (if needed):
   - Review Python SDK entities.md for field descriptions
   - Manually add descriptions to generated markdown
   - Create issue tracking remaining documentation needs

## Key Design Decisions

### 1. Source of Truth
**Decision**: Use `ingester.proto` as the source of truth for structure, with awareness of its limitations
**Rationale**:
- Already the canonical definition for wire format
- Contains all entity structures with correct types and validation
- Stable field numbers preserved across versions
- Used by both Python and Go SDKs
- **Limitation**: Currently lacks field-level documentation (can be enhanced)

### 2. Parser Implementation
**Decision**: Extend existing `protograph` tool
**Rationale**:
- Avoid duplicating protobuf parsing logic
- Leverage existing, tested code
- Maintains consistency with entity mappings generation
- Simpler than protoc plugins for this use case

### 3. Documentation Format
**Decision**: Follow Python SDK `entities.md` format
**Rationale**:
- Users already familiar with the format
- Proven effective for entity documentation
- Easy to maintain consistency across SDKs

### 4. Example Code Generation
**Decision**: Generate complete, buildable Go programs in separate directories, not inline code
**Rationale**:
- Ensures examples are always syntactically correct (can be built in CI)
- Each example is a complete, runnable program with proper error handling
- One `main()` per directory avoids Go package conflicts
- Users can copy entire example directories as starting points
- Examples can be tested independently
- Easy to update when SDK API changes via regeneration
- Documentation references example directories, keeping docs clean

**Structure**:
- Each entity gets its own `docs/examples/<entity>/` directory
- Each directory contains `main.go` and `go.mod`
- Examples use environment variables for configuration (DIODE_API_KEY, DIODE_TARGET)
- All examples follow consistent patterns (imports, error handling, logging)

## Success Criteria

### Minimum Viable Product (MVP)
1. ✅ Documentation generated for all 96+ entity types
2. ✅ Each entity has: entity name, attributes table with types, reference to example code
3. ✅ Attributes table includes: field name, type, required/optional status, validation rules
4. ✅ Complete, runnable Go example for each entity in `docs/examples/<entity>/main.go`
5. ✅ All examples compile successfully (verified in CI)
6. ✅ Documentation automatically updates when protobuf changes
7. ✅ CI validates documentation and examples are up-to-date
8. ✅ Format matches Python SDK structure for consistency
9. ✅ Examples include proper error handling, environment config, and logging

### Stretch Goals (if proto enhancement happens)
8. ✅ Each field has human-readable description explaining its purpose
9. ✅ Entity-level descriptions from NetBox documentation
10. ✅ Cross-references between related entities
11. ✅ Documentation published to docs site with search

## Timeline

- **Week 1**: Prototype and proof of concept (3 entities)
- **Week 2**: Full implementation and CI integration
- **Week 3**: Enhancements and documentation site integration

## Next Steps

### Immediate Actions

1. **Review and approve this plan** with stakeholders
2. **Make decision** on implementation path:
   - Option A: Enhance proto first (best long-term)
   - Option B: Generate from current proto (quick win)
   - Option C: Generate from Django serializers (alternative)

### If choosing Option A (Enhanced Proto)
3. **Coordinate with diode-drf-extract team**:
   - Share this analysis
   - Discuss adding help_text extraction to SerializerTraverser
   - Estimate effort for upstream enhancement
4. **Wait for enhanced proto** or proceed with Option B in parallel

### If choosing Option B (Current Proto)
3. **Create GitHub issue** for tracking
4. **Set up development branch** in diode-sdk-go
5. **Begin Phase 1 implementation** (protograph extension)
6. **Generate prototype** with 2-3 entities
7. **Review prototype** with team
8. **Complete full implementation**
9. **Create follow-up issue** for proto enhancement

### If choosing Option C (Django Serializers)
3. **Set up Django environment** for diode-drf-extract
4. **Create Django management command** for doc generation
5. **Implement serializer documentation extractor**
6. **Generate documentation** with full metadata
7. **Automate copying** to diode-sdk-go

## References

### Documentation
- **Python SDK entities.md**: https://github.com/netboxlabs/diode-sdk-python/blob/5dae6c3222fb99124c3e22bd312ebc4a537ccf9f/docs/entities.md
- **diode-drf-extract README**: `/Users/pstuart/CODE/github.com/netboxlabs/eng-observability-skunkworks/diode-drf-extract/README.md`

### Source Code
- **Protobuf schema**: `/Users/pstuart/CODE/github.com/netboxlabs/diode/diode-proto/diode/v1/ingester.proto` (4,223 lines, no field docs)
- **Protograph tool**: `/Users/pstuart/CODE/github.com/netboxlabs/diode/diode-server/cmd/protograph/`
  - `parser.go` (370 lines) - Proto file parser
  - `generator.go` (141 lines) - Code generator
  - `templates.go` (122 lines) - Go templates
- **Django extraction tool**: `/Users/pstuart/CODE/github.com/netboxlabs/eng-observability-skunkworks/diode-drf-extract/`
  - `visitor.py` - SerializerTraverser implementation
  - `protobuf.py` - Proto generation
  - `proto_merger.py` - Proto merging logic
  - `common.py` - DataModel dataclasses

### Key Configuration
- **Entity definitions**: `/Users/pstuart/CODE/github.com/netboxlabs/eng-observability-skunkworks/diode-drf-extract/output/ingester.yaml`

## Appendix: Metadata Extraction Enhancement

### Concern: Proto File "Pollution"

**Valid concern**: Adding help_text comments to proto files might make them verbose for users who don't need documentation.

**Solution Options**:

### Option 1: Make Help Text Optional (Recommended)

Add configuration to control whether help_text is included in proto generation:

**Modify `ingester.yaml`**:
```yaml
output:
  include_field_documentation: true  # Set to false for clean proto
  documentation_style: "comments"    # Options: "comments", "separate", "none"
```

**Modify `protobuf.py`**:
```python
def generate_field_proto(field_info: FieldInfo, config: Config) -> str:
    lines = []

    # Only add comments if configured
    if config.include_field_documentation and field_info.description:
        # Keep comments concise (max 80 chars per line)
        lines.append(f"  // {field_info.description}")

    lines.append(f"  {field_type} {field_name} = {field_num};")
    return "\n".join(lines)
```

**Benefits**:
- ✅ Default behavior unchanged (clean proto)
- ✅ Opt-in for documentation generation
- ✅ Single proto file to maintain
- ✅ Standard protobuf practice

### Option 2: Separate Documentation File

Generate a companion documentation file alongside the proto:

**File Structure**:
```
diode-proto/diode/v1/
├── ingester.proto           # Clean, no comments
└── ingester.proto.docs.yaml # Documentation metadata
```

**Documentation file format** (`ingester.proto.docs.yaml`):
```yaml
messages:
  Device:
    description: "Represents a physical device in the network infrastructure"
    fields:
      name:
        description: "The unique name of the device"
        help_text: "Must be unique within the site"
      serial:
        description: "Manufacturer-assigned serial number"
        help_text: "Used for hardware inventory tracking"
```

**Benefits**:
- ✅ Proto file stays completely clean
- ✅ Structured documentation format
- ✅ Easy to parse programmatically
- ✅ Can include more metadata than comments allow

**Disadvantages**:
- ❌ Two files to maintain (though generation is automated)
- ❌ Need to coordinate updates between files
- ❌ Non-standard approach

### Option 3: Protobuf Field Options (Advanced)

Use custom protobuf field options instead of comments:

**Define custom options**:
```protobuf
extend google.protobuf.FieldOptions {
  string help_text = 50002;
  string label = 50003;
}

message Device {
  optional string name = 1 [(help_text) = "The unique name of the device"];
  optional string serial = 2 [(help_text) = "Manufacturer-assigned serial number"];
}
```

**Benefits**:
- ✅ Structured, machine-readable metadata
- ✅ Preserved through protoc compilation
- ✅ Can be accessed via reflection in generated code
- ✅ More "proper" protobuf approach

**Disadvantages**:
- ❌ More complex implementation
- ❌ Harder to read for humans scanning the file
- ❌ Not as widely supported by documentation tools

### Recommended Approach: Option 1 (Optional Comments)

**Why**:
1. **Standard Practice**: Protobuf comments are the industry standard for API documentation
2. **No Pollution**: Can be disabled for those who prefer clean protos
3. **Simple Implementation**: Small modification to existing generation code
4. **Tool Support**: Most proto documentation tools expect comments

**Implementation**:

**Modify `visitor.py`**:
```python
@dataclass
class FieldInfo:
    name: str
    field_type: str
    is_required: bool
    is_read_only: bool
    enum_values: list[str]
    min_value: Optional[int]
    max_value: Optional[int]
    # ADD THESE:
    help_text: Optional[str] = None  # From field.help_text
    label: Optional[str] = None      # From field.label
    description: Optional[str] = None # Computed/combined
```

**In SerializerTraverser.traverse_serializer()**:
```python
# Extract help text and label
help_text = getattr(field, 'help_text', None)
label = getattr(field, 'label', None)

field_info = FieldInfo(
    name=field_name,
    # ... other fields ...
    help_text=help_text,
    label=label,
    description=help_text or label or None
)
```

**Modify `protobuf.py` to conditionally write field comments**:
```python
def generate_field_proto(field_info: FieldInfo, include_docs: bool = True) -> str:
    lines = []

    # Only add comments if enabled and description exists
    if include_docs and field_info.description:
        # Format description nicely (wrap long lines)
        desc = format_comment(field_info.description, max_width=80)
        lines.append(f"  // {desc}")

    lines.append(f"  {field_type} {field_name} = {field_num};")
    return "\n".join(lines)

def format_comment(text: str, max_width: int = 80) -> str:
    """Format comment text, wrapping at max_width."""
    # Implementation to wrap long descriptions
    pass
```

**Add configuration to `ingester.yaml`**:
```yaml
# Documentation generation settings
documentation:
  include_in_proto: true  # Set to false for clean proto files
  max_comment_length: 80  # Wrap comments at this width
```

**Update `proto_merger.py` to preserve comments**:
```python
# proto_merger likely already preserves comments, but verify:
def merge_field_comments(old_field, new_field):
    """Preserve comments from old proto if new proto has none."""
    if old_field.comments and not new_field.comments:
        new_field.comments = old_field.comments
    return new_field
```

**Result**:
- Default: Generate proto WITH documentation comments
- Optional: Set `documentation.include_in_proto: false` for clean proto
- Documentation generation tool works with either configuration

This enhancement would make the proto file the comprehensive source of truth for both structure AND documentation, with the flexibility to disable documentation if desired.
