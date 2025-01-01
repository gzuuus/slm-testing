package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type PatternKey string

type TestPattern struct {
	prompts []PromptKey
	configs []ConfigKey
}

const (
	// Test Patterns
	PatternLanguage   PatternKey = "language"
	PatternMathLogic  PatternKey = "math"
	PatternTechnical  PatternKey = "technical"
	PatternScience    PatternKey = "science"
	PatternPractical  PatternKey = "practical"
	PatternPrecision  PatternKey = "precision"
	PatternCreative   PatternKey = "creative"
	PatternMultilang  PatternKey = "multilang"
	PatternAdvMath    PatternKey = "advanced-math"
	PatternHumanities PatternKey = "humanities"
	PatternGameTheory PatternKey = "game-theory"
)

func CustomTest(prompts []PromptKey, configs []ConfigKey) TestPattern {
	return TestPattern{
		prompts: map[bool][]PromptKey{true: prompts, false: AllPrompts()}[len(prompts) > 0],
		configs: map[bool][]ConfigKey{true: configs, false: AllConfigs()}[len(configs) > 0],
	}
}

func CustomRangeTest(promptRange, configRange [2]int) TestPattern {
	prompts := AllPrompts()
	configs := AllConfigs()

	pStart, pEnd := promptRange[0], promptRange[1]
	cStart, cEnd := configRange[0], configRange[1]

	return TestPattern{
		prompts: prompts[pStart:pEnd],
		configs: configs[cStart:cEnd],
	}
}

func RandomTest(promptCount, configCount int) TestPattern {
	prompts := AllPrompts()
	configs := AllConfigs()

	rand.Shuffle(len(prompts), func(i, j int) {
		prompts[i], prompts[j] = prompts[j], prompts[i]
	})

	rand.Shuffle(len(configs), func(i, j int) {
		configs[i], configs[j] = configs[j], configs[i]
	})

	return TestPattern{
		prompts: prompts[:promptCount],
		configs: configs[:configCount],
	}
}

// Category-based test patterns
func LanguageTest(configs []ConfigKey) TestPattern {
	return TestPattern{
		prompts: []PromptKey{
			PromptIdiom, PromptProverb, PromptMetaphor,
			PromptGrammar, PromptSynonym, PromptContext,
		},
		configs: map[bool][]ConfigKey{true: configs, false: AllConfigs()}[len(configs) > 0],
	}
}

func MathAndLogicTest(configs []ConfigKey) TestPattern {
	return TestPattern{
		prompts: []PromptKey{
			PromptMathSimple, PromptMathLogic, PromptProbability,
			PromptLogic, PromptCausation, PromptComparison,
		},
		configs: map[bool][]ConfigKey{true: configs, false: AllConfigs()}[len(configs) > 0],
	}
}

func TechnicalTest(configs []ConfigKey) TestPattern {
	return TestPattern{
		prompts: []PromptKey{
			PromptCodeConcept, PromptAlgorithm, PromptTechExplain,
			PromptTechnology,
		},
		configs: map[bool][]ConfigKey{true: configs, false: AllConfigs()}[len(configs) > 0],
	}
}

func ScienceTest(configs []ConfigKey) TestPattern {
	return TestPattern{
		prompts: []PromptKey{
			PromptPhysics, PromptChemistry, PromptBiology,
			PromptCausation,
		},
		configs: map[bool][]ConfigKey{true: configs, false: AllConfigs()}[len(configs) > 0],
	}
}

func PracticalKnowledgeTest(configs []ConfigKey) TestPattern {
	return TestPattern{
		prompts: []PromptKey{
			PromptFinance, PromptHealth, PromptTechnology,
		},
		configs: map[bool][]ConfigKey{true: configs, false: AllConfigs()}[len(configs) > 0],
	}
}

func PrecisionTest(configs []ConfigKey) TestPattern {
	return TestPattern{
		prompts: []PromptKey{
			PromptMathSimple, PromptLogic, PromptCodeConcept,
			PromptTechExplain, PromptGrammar,
		},
		configs: map[bool][]ConfigKey{true: configs, false: AllConfigs()}[len(configs) > 0],
	}
}

func CreativeExplanationTest(configs []ConfigKey) TestPattern {
	return TestPattern{
		prompts: []PromptKey{
			PromptMetaphor, PromptContext, PromptCausation,
			PromptPhysics, PromptBiology,
		},
		configs: map[bool][]ConfigKey{true: configs, false: {ConfigMirostat2Dynamic, ConfigMirostatBalanced}}[len(configs) > 0],
	}
}

func MultilangTechnicalTest(configs []ConfigKey) TestPattern {
	return TestPattern{
		prompts: []PromptKey{
			PromptTranslation, PromptMultiLingual,
		},
		configs: map[bool][]ConfigKey{true: configs, false: AllConfigs()}[len(configs) > 0],
	}
}

func AdvancedMathTest(configs []ConfigKey) TestPattern {
	return TestPattern{
		prompts: []PromptKey{
			PromptMathComplex, PromptMathProof, PromptMathOptimal,
			PromptMathLogic, PromptProbability,
		},
		configs: map[bool][]ConfigKey{true: configs, false: AllConfigs()}[len(configs) > 0],
	}
}

func HumanitiesTest(configs []ConfigKey) TestPattern {
	return TestPattern{
		prompts: []PromptKey{
			PromptHistoryCause, PromptHistoryCompare,
			PromptEthicalDilemma, PromptMoralPhilosophy,
			PromptMusicTheory, PromptArtAnalysis,
		},
		configs: map[bool][]ConfigKey{true: configs, false: AllConfigs()}[len(configs) > 0],
	}
}

func GameTheoryTest(configs []ConfigKey) TestPattern {
	return TestPattern{
		prompts: []PromptKey{
			PromptGameStrategy, PromptGameTheory,
			PromptLogic, PromptProbability,
		},
		configs: map[bool][]ConfigKey{true: configs, false: AllConfigs()}[len(configs) > 0],
	}
}

var PatternMap = map[PatternKey]func([]ConfigKey) TestPattern{
	PatternLanguage:   LanguageTest,
	PatternMathLogic:  MathAndLogicTest,
	PatternTechnical:  TechnicalTest,
	PatternScience:    ScienceTest,
	PatternPractical:  PracticalKnowledgeTest,
	PatternPrecision:  PrecisionTest,
	PatternCreative:   CreativeExplanationTest,
	PatternMultilang:  MultilangTechnicalTest,
	PatternAdvMath:    AdvancedMathTest,
	PatternHumanities: HumanitiesTest,
	PatternGameTheory: GameTheoryTest,
}

// GetAllPatterns returns all available pattern keys
func GetAllPatterns() []PatternKey {
	patterns := make([]PatternKey, 0, len(PatternMap))
	for p := range PatternMap {
		patterns = append(patterns, p)
	}
	return patterns
}

func ParsePatterns(input string) ([]PatternKey, error) {
	if input == "" {
		return nil, nil
	}

	items := strings.Split(input, ",")
	result := make([]PatternKey, 0, len(items))

	for _, item := range items {
		key := PatternKey(strings.TrimSpace(item))
		if _, exists := PatternMap[key]; !exists {
			return nil, fmt.Errorf("invalid pattern key: %s", item)
		}
		result = append(result, key)
	}

	return result, nil
}
