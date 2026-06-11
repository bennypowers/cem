package jsoncmerge

import (
	"encoding/json"
	"fmt"

	"github.com/tailscale/hujson"
)

// Merge inserts or replaces topKey.subKey in raw JSONC bytes,
// preserving all comments and formatting via byte-level splicing.
// Uses hujson AST for key location to avoid matching inside comments/strings.
func Merge(data []byte, topKey, subKey string, value any) ([]byte, error) {
	v, err := hujson.Parse(data)
	if err != nil {
		return nil, fmt.Errorf("invalid JSONC: %w", err)
	}
	v.UpdateOffsets()

	topNode := v.Find("/" + topKey)

	if topNode == nil {
		rootObj, ok := v.Value.(*hujson.Object)
		if !ok {
			return nil, fmt.Errorf("root is not an object")
		}
		closeBrace := closeOffset(&v, rootObj)
		topKeyStr := fmt.Sprintf("%q", topKey)
		innerJSON, _ := json.MarshalIndent(map[string]any{subKey: value}, "  ", "  ")
		newBlock := fmt.Sprintf("  %s: %s", topKeyStr, string(innerJSON))
		return insertBeforeClose(data, closeBrace, newBlock, "")
	}

	subNode := v.Find("/" + topKey + "/" + subKey)

	if subNode != nil {
		valueJSON, _ := json.MarshalIndent(value, "    ", "  ")
		return splice(data, subNode.StartOffset, subNode.EndOffset, valueJSON), nil
	}

	topObj, ok := topNode.Value.(*hujson.Object)
	if !ok {
		return nil, fmt.Errorf("%q is not an object", topKey)
	}
	closeBrace := closeOffset(topNode, topObj)
	valueJSON, _ := json.MarshalIndent(value, "    ", "  ")
	newMember := fmt.Sprintf("    %q: %s", subKey, string(valueJSON))
	return insertBeforeClose(data, closeBrace, newMember, "  ")
}

// closeOffset returns the byte position of the closing } in the original data.
// EndOffset points past the }, and AfterExtra (whitespace inside before })
// is accounted for within the object's Pack representation.
func closeOffset(v *hujson.Value, obj *hujson.Object) int {
	return v.EndOffset - len("}") - len(obj.AfterExtra)
}

func insertBeforeClose(data []byte, closeBrace int, content, closeIndent string) ([]byte, error) {
	lastSig := closeBrace - 1
	for lastSig >= 0 {
		ch := data[lastSig]
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			lastSig--
			continue
		}
		if ch == '/' && lastSig > 0 && data[lastSig-1] == '/' {
			for lastSig > 0 && data[lastSig-1] != '\n' {
				lastSig--
			}
			lastSig--
			continue
		}
		if ch == '/' && lastSig >= 2 && data[lastSig-1] == '*' {
			j := lastSig - 2
			for j >= 1 {
				if data[j] == '/' && data[j+1] == '*' {
					lastSig = j - 1
					break
				}
				j--
			}
			if j < 1 {
				break
			}
			continue
		}
		break
	}

	if lastSig < 0 || data[lastSig] == '{' || data[lastSig] == '[' {
		inserted := "\n" + content + "\n" + closeIndent
		return splice(data, closeBrace, closeBrace, []byte(inserted)), nil
	}

	if data[lastSig] == ',' {
		inserted := "\n" + content + "\n" + closeIndent
		return splice(data, lastSig+1, closeBrace, []byte(inserted)), nil
	}

	inserted := ",\n" + content + "\n" + closeIndent
	return splice(data, lastSig+1, closeBrace, []byte(inserted)), nil
}

func splice(data []byte, start, end int, insert []byte) []byte {
	var buf []byte
	buf = append(buf, data[:start]...)
	buf = append(buf, insert...)
	buf = append(buf, data[end:]...)
	return buf
}
