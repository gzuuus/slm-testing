# SML Testing Tool

A comprehensive testing tool for evaluating Small Language Models (SLMs) across various configurations and prompt patterns. This tool enables the optimization of model responses through the exploration of different prompting strategies and configuration settings, leveraging [Ollama](https://ollama.com/) for inference.

## Features

- **Flexible Testing**: Test models with predefined patterns or specific prompts
- **Configuration Presets**: Multiple preset configurations for different use cases
- **Response Metrics**: Track response time, character count, and word count
- **Export Capability**: Save results to JSON for further analysis
- **Interactive Output**: Real-time console feedback during testing
- **Category-based Testing**: Pre-organized test patterns for different domains

## Installation

Make sure you have [Ollama](https://ollama.com/) installed in your system, then:

```bash
# Clone the repository
git https://github.com/gzuuus/slm-testing

# Navigate to the project directory
cd sml-testing

# Install dependencies
go mod download
```

## Usage

The tool provides a flexible command-line interface with various flags:

```bash
go run . [flags]
```

### Command-line Flags

#### Output Control
- `-print`: Print results to console (default: true)
- `-export`: Export results to JSON file

#### Connection Settings
- `-url`: LLM server URL (default: "http://localhost:11434")
- `-model`: Model name (default: "qwen2.5:0.5b-instruct")

#### Test Configuration
- `-patterns`: Comma-separated list of test patterns to run
- `-configs`: Comma-separated list of model configurations
- `-prompts`: Comma-separated list of specific prompts to test

#### Utility
- `-help`: Display help message

### Configuration Presets

Each preset is optimized for specific use cases:

| Preset | Description |
|--------|-------------|
| Ultra-Precise | High precision settings for factual responses |
| Creative-High | Settings optimized for creative outputs |
| Mirostat-Balanced | Balanced temperature with Mirostat control |
| Mirostat2-Dynamic | Dynamic adjustment using Mirostat 2 |
| Analytical | Settings tuned for analytical responses |

### Test Patterns

Available test categories:

| Pattern | Description |
|---------|-------------|
| language | Language understanding and usage |
| math | Mathematical and logical reasoning |
| technical | Technical concepts and explanations |
| science | Scientific principles and phenomena |
| practical | Practical knowledge |
| precision | High-precision responses |
| creative | Creative explanations |
| multilang | Multilingual capabilities |
| advanced-math | Advanced mathematical problems |
| humanities | Humanities and arts |
| game-theory | Game theory and strategy |

### Example Commands

```bash
# Run language and technical patterns
go run . -patterns=language,technical

# Use specific configurations
go run . -configs=Ultra-Precise,Creative-High

# Test specific prompts
go run . -prompts=idiom,proverb,metaphor

# Combine patterns and configs
go run . -patterns=language -configs=Ultra-Precise

# Export results
go run . -patterns=technical -export
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.