package trie

import (
	"math/rand"
	"testing"
	"time"
)

var (
	testWords = []string{
		"apple", "app", "application", "apply", "appreciate", "approach",
		"banana", "band", "bandana", "bank", "banking", "bankruptcy",
		"cat", "car", "card", "care", "careful", "careless",
		"dog", "door", "down", "download", "downward", "downtown",
		"elephant", "electric", "electricity", "electronic", "element",
		"fire", "first", "fish", "fishing", "fisherman", "fit",
		"go", "good", "google", "government", "governor", "grace",
		"house", "home", "hospital", "host", "hosting", "hotel",
	}

	longWords  = generateLongWords(1000, 20)  // 1000 words, avg 20 chars each
	shortWords = generateShortWords(10000, 5) // 10000 words, avg 5 chars each
)

func generateLongWords(count, avgLength int) []string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	words := make([]string, count)

	for i := range count {
		length := max(avgLength+rand.Intn(10)-5, 1)

		word := make([]rune, length)
		for j := 0; j < length; j++ {
			word[j] = rune('a' + rand.Intn(26))
		}
		words[i] = string(word)
	}
	return words
}

func generateShortWords(count, avgLength int) []string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano() + 1))
	words := make([]string, count)

	for i := range count {
		length := max(avgLength+rng.Intn(4)-2, 1)

		word := make([]rune, length)
		for j := range length {
			word[j] = rune('a' + rng.Intn(26))
		}
		words[i] = string(word)
	}
	return words
}

// Map Trie Benchmarks
func BenchmarkMapTrie_Insert(b *testing.B) {
	b.Run("SmallWords", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			trie := NewMapTrie()
			for _, word := range testWords {
				trie.Insert(word)
			}
		}
	})

	b.Run("LongWords", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			trie := NewMapTrie()
			for _, word := range longWords[:100] { // Use subset for reasonable benchmark time
				trie.Insert(word)
			}
		}
	})

	b.Run("ManyShortWords", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			trie := NewMapTrie()
			for _, word := range shortWords[:1000] { // Use subset
				trie.Insert(word)
			}
		}
	})
}

func BenchmarkMapTrie_Search(b *testing.B) {
	trie := NewMapTrie()
	for _, word := range testWords {
		trie.Insert(word)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, word := range testWords {
			trie.Search(word)
		}
	}
}

func BenchmarkMapTrie_StartsWith(b *testing.B) {
	trie := NewMapTrie()
	for _, word := range testWords {
		trie.Insert(word)
	}

	prefixes := []string{"app", "ban", "car", "do", "ele"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, prefix := range prefixes {
			trie.StartsWith(prefix)
		}
	}
}

// List Trie Benchmarks
func BenchmarkListTrie_Insert(b *testing.B) {
	b.Run("SmallWords", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			trie := NewListTrie()
			for _, word := range testWords {
				trie.Insert(word)
			}
		}
	})

	b.Run("LongWords", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			trie := NewListTrie()
			for _, word := range longWords[:100] {
				trie.Insert(word)
			}
		}
	})

	b.Run("ManyShortWords", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			trie := NewListTrie()
			for _, word := range shortWords[:1000] {
				trie.Insert(word)
			}
		}
	})
}

func BenchmarkListTrie_Search(b *testing.B) {
	trie := NewListTrie()
	for _, word := range testWords {
		trie.Insert(word)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, word := range testWords {
			trie.Search(word)
		}
	}
}

func BenchmarkListTrie_StartsWith(b *testing.B) {
	trie := NewListTrie()
	for _, word := range testWords {
		trie.Insert(word)
	}

	prefixes := []string{"app", "ban", "car", "do", "ele"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, prefix := range prefixes {
			trie.StartsWith(prefix)
		}
	}
}

// Memory usage comparison
func BenchmarkMemoryUsage(b *testing.B) {
	b.Run("MapTrie_Memory", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			trie := NewMapTrie()
			for _, word := range testWords {
				trie.Insert(word)
			}
		}
	})

	b.Run("ListTrie_Memory", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			trie := NewListTrie()
			for _, word := range testWords {
				trie.Insert(word)
			}
		}
	})

	b.Run("Trie_Memory", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			trie := NewTrie()
			for _, word := range testWords {
				trie.Insert(word)
			}
		}
	})

}

func BenchmarkTrie_Search(b *testing.B) {
	trie := NewTrie()
	for _, word := range testWords {
		trie.Insert(word)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, word := range testWords {
			trie.Search(word)
		}
	}
}

func BenchmarkTrie_Insert(b *testing.B) {
	b.Run("SmallWords", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			trie := NewTrie()
			for _, word := range testWords {
				trie.Insert(word)
			}
		}
	})

	b.Run("LongWords", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			trie := NewTrie()
			for _, word := range longWords[:100] {
				trie.Insert(word)
			}
		}
	})

	b.Run("ManyShortWords", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			trie := NewTrie()
			for _, word := range shortWords[:1000] {
				trie.Insert(word)
			}
		}
	})
}

func BenchmarkTrie_Delete(b *testing.B) {
	b.Run("SmallWords", func(b *testing.B) {
		trie := NewTrie()
		for _, word := range testWords {
			trie.Insert(word)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, word := range testWords {
				trie.Delete(word)
			}
		}
	})

	b.Run("LongWords", func(b *testing.B) {
		trie := NewTrie()
		for _, word := range longWords[:100] {
			trie.Insert(word)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, word := range longWords[:100] {
				trie.Delete(word)
			}
		}
	})

	b.Run("ManyShortWords", func(b *testing.B) {
		trie := NewTrie()
		for _, word := range shortWords[:1000] {
			trie.Insert(word)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, word := range shortWords[:1000] {
				trie.Delete(word)
			}
		}
	})
}

func BenchmarkTrie_AutoComplete(b *testing.B) {
	trie := NewTrie()
	for _, word := range testWords {
		trie.Insert(word)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, word := range testWords {
			trie.AutoComplete(word)
		}
	}
}
