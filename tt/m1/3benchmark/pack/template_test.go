package pack

import (
	"bytes"
	"html/template"
	"testing"
)

func BenchmarkExample(b *testing.B) {
	temp, _ := template.New("").Parse("Hello, Go")
	b.ResetTimer()
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		temp.Execute(&buf, nil)
		buf.Reset()
	}
}
