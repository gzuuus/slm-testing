package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/parakeet-nest/parakeet/completion"
	"github.com/parakeet-nest/parakeet/llm"
)

// ResponseMetrics holds the basic metrics for each test response
type ResponseMetrics struct {
	ResponseTime time.Duration `json:"responseTime"`
	CharCount    int           `json:"charCount"`
	WordCount    int           `json:"wordCount"`
}

// TestResult combines the test configuration, prompt, response, and metrics
type TestResult struct {
	Config    ConfigKey       `json:"config"`
	Prompt    PromptKey       `json:"prompt"`
	Response  string          `json:"response"`
	Metrics   ResponseMetrics `json:"metrics"`
	Timestamp time.Time       `json:"timestamp"`
}

// Flags holds the program's command line flags
type Flags struct {
	ExportJSON   bool
	PrintResults bool
	URL          string
	Model        string
	Patterns     string
	Configs      string
	Prompts      string
	Help         bool
}

const helpText = `SML Testing Tool
=============
A comprehensive testing tool for evaluating Small Language Models (SLMs) across 
different configurations and prompt patterns.

Usage:
  go run . [flags]

Flags:
  Output Control:
    -print        Print results to console (default: true)
    -export       Export results to JSON file

  Connection Settings:
    -url string   LLM server URL (default: "http://localhost:11434")
    -model string Model name (default: "qwen2.5:0.5b-instruct")

  Test Configuration:
    -patterns     Comma-separated list of test patterns to run
                 Available: %s
    
    -configs      Comma-separated list of model configurations
                 Available: %s
    
    -prompts      Comma-separated list of specific prompts to test
                 Available: %s

  Utility:
    -help        Display this help message

Examples:
  # Run specific test patterns
  go run . -patterns=language,technical

  # Use specific model configurations
  go run . -configs=Ultra-Precise,Creative-High

  # Test specific prompts
  go run . -prompts=idiom,proverb,metaphor

  # Combine patterns and configs
  go run . -patterns=language -configs=Ultra-Precise

  # Export results to JSON
  go run . -patterns=technical -export
`

func toStrings[T ~string](items []T) []string {
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = string(item)
	}
	return result
}

func parseFlags() *Flags {
	flags := &Flags{}

	flag.BoolVar(&flags.Help, "help", false, "Show help message")
	flag.BoolVar(&flags.ExportJSON, "export", false, "Export results to JSON file")
	flag.BoolVar(&flags.PrintResults, "print", true, "Print results to console")
	flag.StringVar(&flags.URL, "url", "http://localhost:11434", "LLM server URL")
	flag.StringVar(&flags.Model, "model", "deepseek-r1:1.5b", "Model name")
	flag.StringVar(&flags.Patterns, "patterns", "", "Comma-separated list of test patterns")
	flag.StringVar(&flags.Configs, "configs", "", "Comma-separated list of configurations")
	flag.StringVar(&flags.Prompts, "prompts", "", "Comma-separated list of specific prompts")

	// Custom usage message
	flag.Usage = func() {
		patterns := strings.Join(toStrings(GetAllPatterns()), ", ")
		configs := strings.Join(toStrings(AllConfigs()), ", ")
		prompts := strings.Join(toStrings(AllPrompts()), ", ")
		fmt.Printf(helpText, patterns, configs, prompts)
	}

	flag.Parse()

	if flags.Help {
		flag.Usage()
		os.Exit(0)
	}

	return flags
}

// TestLLM runs a single test with the specified configuration
func TestLLM(url, model string, promptKey PromptKey, configKey ConfigKey, config map[string]interface{}) (TestResult, error) {
	// Start timing from the moment we begin processing
	startTime := time.Now()

	options := llm.SetOptions(config)
	question := llm.GenQuery{
		Model:   model,
		Prompt:  TestPrompts[promptKey],
		Options: options,
		System:  "You're a friendly and helpful assistant providing concise and accurate answers.",
	}

	// Generate response and measure total time including network and processing
	answer, err := completion.Generate(url, question)
	if err != nil {
		return TestResult{}, fmt.Errorf("generation error: %v", err)
	}

	// Calculate metrics
	endTime := time.Now()
	responseTime := endTime.Sub(startTime)

	metrics := ResponseMetrics{
		ResponseTime: responseTime,
		CharCount:    len(answer.Response),
		WordCount:    len(strings.Fields(answer.Response)),
	}

	result := TestResult{
		Config:    configKey,
		Prompt:    promptKey,
		Response:  answer.Response,
		Metrics:   metrics,
		Timestamp: endTime,
	}

	return result, nil
}

// RunTestPattern executes all tests in a pattern and collects results
func RunTestPattern(url, model string, pattern TestPattern, shouldPrint bool) ([]TestResult, error) {
	results := make([]TestResult, 0)

	for _, prompt := range pattern.prompts {
		for _, config := range pattern.configs {
			fmt.Printf("🤖 Running '%s' with %s configuration:\n", prompt, config)

			result, err := TestLLM(url, model, prompt, config, Configs[config])
			if err != nil {
				fmt.Printf("Error testing prompt %s with config %s: %v\n", prompt, config, err)
				continue
			}

			if shouldPrint {
				printResponse(result)
			}

			results = append(results, result)
		}
	}

	return results, nil
}

// ExportResults saves test results to a JSON file
func ExportResults(results []TestResult, baseFilename string) error {
	timestamp := time.Now().Format("2006-01-02_150405")
	filename := fmt.Sprintf("%s_%s.json", baseFilename, timestamp)

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		return fmt.Errorf("failed to encode results: %v", err)
	}

	fmt.Printf("\nResults exported to: %s\n", filename)
	return nil
}

// printResponse outputs the test results to console in a readable format
func printResponse(result TestResult) {
	fmt.Printf("\n%s\n", strings.Repeat("-", 40))
	fmt.Printf(">> Test Configuration:\n")
	fmt.Printf("> Config: %s\n", result.Config)
	fmt.Printf("> Prompt: %s\n", result.Prompt)
	fmt.Printf("> Prompt text: %s\n", TestPrompts[result.Prompt])

	fmt.Printf("\nMetrics:\n")
	fmt.Printf("- Response time: %v\n", result.Metrics.ResponseTime)
	fmt.Printf("- Character count: %d\n", result.Metrics.CharCount)
	fmt.Printf("- Word count: %d\n", result.Metrics.WordCount)

	fmt.Printf("\nResponse:\n%s\n", result.Response)
	fmt.Printf("\n%s\n", strings.Repeat("-", 40))
}

func main() {
	flags := parseFlags()

	// Parse and validate configurations
	selectedConfigs, err := ParseConfigs(flags.Configs)
	if err != nil {
		fmt.Printf("Error parsing configurations: %v\n", err)
		fmt.Println("Available configs:", AllConfigs())
		return
	}

	// Parse and validate specific prompts if provided
	selectedPrompts, err := ParsePrompts(flags.Prompts)
	if err != nil {
		fmt.Printf("Error parsing prompts: %v\n", err)
		fmt.Println("Available prompts:", AllPrompts())
		return
	}

	// Parse and validate patterns
	selectedPatterns, err := ParsePatterns(flags.Patterns)
	if err != nil {
		fmt.Printf("Error parsing patterns: %v\n", err)
		fmt.Println("Available patterns:", GetAllPatterns())
		return
	}

	// Determine which test pattern to use
	var pattern TestPattern

	switch {
	case len(selectedPrompts) > 0:
		// If specific prompts are provided, use them directly
		pattern = CustomTest(selectedPrompts, selectedConfigs)
	case len(selectedPatterns) > 0:
		// If patterns are provided, combine their prompts
		var combinedPrompts []PromptKey
		for _, p := range selectedPatterns {
			patternResult := PatternMap[p](nil)
			combinedPrompts = append(combinedPrompts, patternResult.prompts...)
		}
		pattern = CustomTest(combinedPrompts, selectedConfigs)
	default:
		// Default to RandomTest if no pattern or prompts are specified
		pattern = RandomTest(5, 5)
	}

	// Run tests
	startTime := time.Now()
	results, err := RunTestPattern(flags.URL, flags.Model, pattern, flags.PrintResults)
	if err != nil {
		fmt.Printf("Error running tests: %v\n", err)
		return
	}

	// Export results if flag is set
	if flags.ExportJSON {
		if err := ExportResults(results, "test_results"); err != nil {
			fmt.Printf("Error exporting results: %v\n", err)
			return
		}
	}

	// Print summary
	fmt.Printf("\nCompleted %d tests in %v\n", len(results), time.Since(startTime))
}
