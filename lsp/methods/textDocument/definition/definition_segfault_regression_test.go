package definition_test

import (
	"sync"
	"testing"
	"time"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/definition"
	"bennypowers.dev/cem/lsp/testhelpers"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestDefinitionNilPointerRegression tests for the segmentation fault
// that was occurring when Document.Content() was called on a corrupted document
func TestDefinitionNilPointerRegression(t *testing.T) {
	// Create a mock server context
	ctx := testhelpers.NewMockServerContext()
	
	// Create a real DocumentManager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Create a document with button-element
	content := `<button-element>Test</button-element>`
	doc := dm.OpenDocument("file:///test.html", content, 1)
	ctx.AddDocument("file:///test.html", doc)

	// Test basic definition call (should not crash)
	params := &protocol.DefinitionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: "file:///test.html",
			},
			Position: protocol.Position{Line: 0, Character: 5}, // Inside button-element
		},
	}

	// This should not crash even if element is not found in registry
	result, err := definition.Definition(ctx, nil, params)
	if err != nil {
		t.Errorf("Definition call failed: %v", err)
	}
	
	// Result should be nil since button-element is not in registry, but no crash
	if result != nil {
		t.Logf("Definition result: %v", result)
	}
}

// TestDefinitionConcurrentAccess tests concurrent access to documents
// to ensure no race conditions cause segfaults
func TestDefinitionConcurrentAccess(t *testing.T) {
	// Create a mock server context
	ctx := testhelpers.NewMockServerContext()
	
	// Create a real DocumentManager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Create a document with button-element
	content := `<button-element>Test</button-element>`
	doc := dm.OpenDocument("file:///test.html", content, 1)
	ctx.AddDocument("file:///test.html", doc)

	var wg sync.WaitGroup
	var errCount int
	var mu sync.Mutex
	
	// Start multiple goroutines that try to access the document concurrently
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Simulate concurrent access to documents
			for j := 0; j < 10; j++ {
				// Try to get document content
				doc := ctx.Document("file:///test.html")
				if doc != nil {
					_, err := doc.Content()
					if err != nil {
						mu.Lock()
						errCount++
						mu.Unlock()
					}
				}
				
				// Small delay to increase chance of race conditions
				time.Sleep(time.Microsecond * 10)
			}
		}(i)
	}

	// Concurrently open/close documents to create race conditions
	wg.Add(1)
	go func() {
		defer wg.Done()
		
		for i := 0; i < 5; i++ {
			// Open a new document
			newDoc := dm.OpenDocument("file:///test.html", content, int32(i+2))
			ctx.AddDocument("file:///test.html", newDoc)
			
			time.Sleep(time.Microsecond * 50)
			
			// Close the document
			dm.CloseDocument("file:///test.html")
			
			time.Sleep(time.Microsecond * 50)
		}
	}()

	// Try definition requests while documents are being accessed concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		
		for i := 0; i < 3; i++ {
			params := &protocol.DefinitionParams{
				TextDocumentPositionParams: protocol.TextDocumentPositionParams{
					TextDocument: protocol.TextDocumentIdentifier{
						URI: "file:///test.html",
					},
					Position: protocol.Position{Line: 0, Character: 5}, // Inside button-element
				},
			}

			_, err := definition.Definition(ctx, nil, params)
			if err != nil {
				mu.Lock()
				errCount++
				mu.Unlock()
			}
			
			time.Sleep(time.Millisecond * 50)
		}
	}()

	// Wait for all goroutines to complete
	wg.Wait()
	
	// The test passes if we reach here without a segfault
	// Some errors are expected due to document lifecycle, but no crashes
	t.Logf("Test completed with %d errors (expected in concurrent access)", errCount)
}