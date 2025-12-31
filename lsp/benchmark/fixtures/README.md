# LSP Benchmark Test Fixtures

This directory contains test fixtures for benchmarking LSP performance at different scales.

## Project Sizes

### Small Project (3 elements)
- **Location**: `small_project/`
- **Elements**: 3 basic components (button, card, input)
- **Purpose**: Baseline performance testing
- **Use Case**: Small applications or proof-of-concepts

### Medium Project (realistic scale) 
- **Location**: `medium_project/`
- **Elements**: 5 components with realistic complexity
- **Purpose**: Real-world performance testing
- **Use Case**: Typical web application component library

### Large Project (stress testing)
- **Location**: `large_project/`
- **Elements**: 20+ components with full feature sets
- **Purpose**: Stress testing and scalability analysis
- **Use Case**: Enterprise-scale component libraries

## Fixture Structure

Each fixture contains:
- `custom-elements.json` - Component manifest for LSP analysis
- `package.json` - Project metadata (where applicable)
- Sample HTML/TypeScript files for testing (generated during benchmarks)

## Element Complexity

The fixtures progressively increase in complexity:

1. **Basic**: Simple attributes, minimal slots
2. **Intermediate**: Multiple variants, events, complex attributes
3. **Advanced**: Full feature sets, extensive slot systems, type unions

## Usage in Benchmarks

These fixtures are automatically used by the benchmark scenarios:
- Startup tests use all fixture sizes
- Performance tests focus on medium project for realistic metrics
- Stress tests use large project for scalability analysis

## Customization

To add new fixtures:
1. Create a new directory under `fixtures/`
2. Add `custom-elements.json` with your component definitions
3. Update benchmark scenarios to include the new fixture
4. Document the fixture purpose and scale in this README