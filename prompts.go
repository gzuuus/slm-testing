package main

import (
	"fmt"
	"strings"
)

type PromptKey string

const (
	// Cultural Knowledge
	PromptIdiom    PromptKey = "idiom"
	PromptProverb  PromptKey = "proverb"
	PromptMetaphor PromptKey = "metaphor"

	// Mathematics
	PromptMathSimple  PromptKey = "math_simple"
	PromptMathLogic   PromptKey = "math_logic"
	PromptProbability PromptKey = "probability"

	// Reasoning
	PromptLogic      PromptKey = "logic"
	PromptCausation  PromptKey = "causation"
	PromptComparison PromptKey = "comparison"

	// Technical
	PromptCodeConcept PromptKey = "code_concept"
	PromptAlgorithm   PromptKey = "algorithm"
	PromptTechExplain PromptKey = "tech_explain"

	// Scientific
	PromptPhysics   PromptKey = "physics"
	PromptChemistry PromptKey = "chemistry"
	PromptBiology   PromptKey = "biology"

	// Language
	PromptGrammar PromptKey = "grammar"
	PromptSynonym PromptKey = "synonym"
	PromptContext PromptKey = "context"

	// Practical
	PromptFinance    PromptKey = "finance"
	PromptHealth     PromptKey = "health"
	PromptTechnology PromptKey = "technology"

	// Multilingual Capabilities
	PromptTranslation  PromptKey = "translation"
	PromptMultiLingual PromptKey = "multilingual"
	PromptLocalContext PromptKey = "local_context"

	// Technical Deep Dives
	PromptSystemDesign  PromptKey = "system_design"
	PromptDebugScenario PromptKey = "debug_scenario"
	PromptCodeReview    PromptKey = "code_review"

	// Knowledge Integration
	PromptCrossDomain   PromptKey = "cross_domain"
	PromptTrendAnalysis PromptKey = "trend_analysis"
	PromptInnovation    PromptKey = "innovation"

	// Advanced Mathematics
	PromptMathComplex PromptKey = "math_complex"
	PromptMathProof   PromptKey = "math_proof"
	PromptMathOptimal PromptKey = "math_optimal"

	// Historical Analysis
	PromptHistoryCause   PromptKey = "history_cause"
	PromptHistoryCompare PromptKey = "history_compare"

	// Ethical Reasoning
	PromptEthicalDilemma  PromptKey = "ethical_dilemma"
	PromptMoralPhilosophy PromptKey = "moral_philosophy"

	// Music and Art Theory
	PromptMusicTheory PromptKey = "music_theory"
	PromptArtAnalysis PromptKey = "art_analysis"

	// Game Theory
	PromptGameStrategy PromptKey = "game_strategy"
	PromptGameTheory   PromptKey = "game_theory"
)

var TestPrompts = map[PromptKey]string{
	PromptIdiom:           "What does the phrase 'beating around the bush' mean?",
	PromptProverb:         "Explain the meaning of 'a stitch in time saves nine'",
	PromptMetaphor:        "What does it mean when someone says 'life is a roller coaster'?",
	PromptMathSimple:      "What is the fastest way to calculate 15% of a number?",
	PromptMathLogic:       "If a clock shows 3:15, what is the angle between the hour and minute hands?",
	PromptProbability:     "What are the odds of rolling two sixes with two dice?",
	PromptLogic:           "If all A are B, and all B are C, what can we conclude about A and C?",
	PromptCausation:       "Why does ice float in water?",
	PromptComparison:      "What's the key difference between correlation and causation?",
	PromptCodeConcept:     "What is a recursive function?",
	PromptAlgorithm:       "Explain how binary search works",
	PromptTechExplain:     "What is the difference between HTTP and HTTPS?",
	PromptPhysics:         "Why do we see lightning before we hear thunder?",
	PromptChemistry:       "Why does salt dissolve in water?",
	PromptBiology:         "What is the main function of red blood cells?",
	PromptGrammar:         "When should we use 'who' versus 'whom'?",
	PromptSynonym:         "What's the difference between 'eager' and 'anxious'?",
	PromptContext:         "In the phrase 'bank account' and 'river bank', what causes the different meanings of 'bank'?",
	PromptFinance:         "What's the difference between a debit card and a credit card?",
	PromptHealth:          "Why is breakfast considered important?",
	PromptTechnology:      "What's the purpose of a CPU in a computer?",
	PromptTranslation:     "Translate 'Welcome to the future of AI' into French, Japanese, and Arabic.",
	PromptMultiLingual:    "Respond to this question in Spanish: What are the key benefits of renewable energy?",
	PromptLocalContext:    "Explain how the concept of 'time' is viewed differently across various cultures.",
	PromptSystemDesign:    "Design a high-level architecture for a real-time chat application that needs to support millions of users.",
	PromptDebugScenario:   "Given a Node.js application with high CPU usage and memory leaks, what steps would you take to diagnose and fix the issues?",
	PromptCodeReview:      "Review this code snippet for potential issues: `function fetchData(callback) { const data = getData(); callback(data); }`",
	PromptCrossDomain:     "Explain how principles of biology could be applied to improve computer network design.",
	PromptTrendAnalysis:   "Analyze the intersection of AI advancement and its impact on human creativity in various fields.",
	PromptInnovation:      "Propose a novel solution for reducing urban traffic congestion using emerging technologies.",
	PromptMathComplex:     `Solve this calculus problem: Find the volume of the solid obtained by rotating the region bounded by y = x², y = 2x, and the y-axis about the x-axis. Show your work and explain each step.`,
	PromptMathProof:       `Prove that the square root of 2 is irrational using a proof by contradiction. Explain your reasoning in detail.`,
	PromptMathOptimal:     `A company produces two types of products, A and B. Product A requires 2 hours of labor and 3 units of raw material, while Product B requires 3 hours of labor and 2 units of raw material. The company has 100 hours of labor and 120 units of raw material available. Product A sells for $40 and Product B for $50. How many of each product should be produced to maximize profit?`,
	PromptHistoryCause:    `Analyze the primary causes of the Industrial Revolution and their interconnections. How did these factors influence each other?`,
	PromptHistoryCompare:  `Compare and contrast the Renaissance in Italy and the Golden Age in China. What were the key similarities and differences in their cultural and scientific achievements?`,
	PromptEthicalDilemma:  `A self-driving car must make a split-second decision: swerve to avoid a group of pedestrians but put its passenger at risk, or maintain course to protect its passenger but harm the pedestrians. What ethical frameworks could guide this decision?`,
	PromptMoralPhilosophy: `Compare utilitarian and deontological approaches to privacy rights in the digital age. How would each framework address data collection practices?`,
	PromptMusicTheory:     `Explain the concept of modal interchange in music theory. How does it differ from standard diatonic harmony, and what emotional effects can it create?`,
	PromptArtAnalysis:     `Analyze the use of perspective, light, and symbolism in Vermeer's "Girl with a Pearl Earring". How do these elements contribute to the painting's impact?`,
	PromptGameStrategy:    `In a game of prisoner's dilemma repeated 100 times, what would be the optimal strategy? Consider both theoretical and practical aspects.`,
	PromptGameTheory:      `Explain how Nash Equilibrium applies to market competition between two companies setting prices for similar products.`,
}

// Get all prompts as a slice
func AllPrompts() []PromptKey {
	prompts := make([]PromptKey, 0, len(TestPrompts))
	for p := range TestPrompts {
		prompts = append(prompts, p)
	}
	return prompts
}

func ParsePrompts(input string) ([]PromptKey, error) {
	if input == "" {
		return nil, nil
	}

	items := strings.Split(input, ",")
	result := make([]PromptKey, 0, len(items))

	for _, item := range items {
		key := PromptKey(strings.TrimSpace(item))
		if _, exists := TestPrompts[key]; !exists {
			return nil, fmt.Errorf("invalid prompt key: %s", item)
		}
		result = append(result, key)
	}

	return result, nil
}
