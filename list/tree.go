/*
Copyright © 2025 Benny Powers

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package list

import (
	"fmt"
	"strings"

	"github.com/pterm/pterm"

	M "bennypowers.dev/cem/manifest"
)

// FilterByDeprecated filters a slice of members based on deprecated status.
// If showDeprecated is true, returns only deprecated members.
// If showDeprecated is false, returns only non-deprecated members.
func FilterByDeprecated[T M.RenderableMemberWithContext](items []T, showDeprecated bool) []T {
	var filtered []T
	for _, item := range items {
		if item.IsDeprecated() == showDeprecated {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

// RenderTree renders members in a tree format, grouped by member type.
func RenderTree[T M.RenderableMemberWithContext](title string, items []T, showDeprecated *bool) error {
	if len(items) == 0 {
		return nil
	}

	// Filter by deprecated status if specified
	if showDeprecated != nil {
		items = FilterByDeprecated(items, *showDeprecated)
		if len(items) == 0 {
			return nil
		}
	}

	// Group items by type
	groups := make(map[string][]T)
	for _, item := range items {
		memberType := item.GetMemberType()
		groups[memberType] = append(groups[memberType], item)
	}

	if len(groups) == 0 {
		return nil
	}

	// Print title
	pterm.DefaultSection.Println(title)

	// Define display order for member types
	typeOrder := []string{"attribute", "field", "method", "event", "slot", "css-property", "css-state", "css-part"}
	typeTitles := map[string]string{
		"attribute":    "Attributes",
		"field":        "Fields", 
		"method":       "Methods",
		"event":        "Events",
		"slot":         "Slots",
		"css-property": "CSS Properties",
		"css-state":    "CSS States",
		"css-part":     "CSS Parts",
	}

	// Render each group
	for _, memberType := range typeOrder {
		if group, exists := groups[memberType]; exists {
			title := typeTitles[memberType]
			if showDeprecated != nil && *showDeprecated {
				title = "Deprecated " + title
			}
			
			pterm.DefaultBasicText.Println("")
			pterm.DefaultHeader.WithFullWidth(false).WithBackgroundStyle(pterm.NewStyle(pterm.BgBlue)).Println(title)
			
			for _, item := range group {
				renderTreeItem(item)
			}
		}
	}

	return nil
}

// RenderTreeGrouped renders all member types grouped together in tree format.
func RenderTreeGrouped(title string, allMembers map[string][]M.RenderableMemberWithContext, showDeprecated *bool) error {
	if len(allMembers) == 0 {
		return nil
	}

	// Print title
	pterm.DefaultSection.Println(title)

	// Define display order for member types
	typeOrder := []string{"attributes", "fields", "methods", "events", "slots"}
	typeTitles := map[string]string{
		"attributes": "Attributes",
		"fields":     "Fields",
		"methods":    "Methods",
		"events":     "Events", 
		"slots":      "Slots",
	}

	hasAnyContent := false

	// Render each group
	for _, memberType := range typeOrder {
		if members, exists := allMembers[memberType]; exists {
			// Filter by deprecated status if specified
			if showDeprecated != nil {
				var filtered []M.RenderableMemberWithContext
				for _, member := range members {
					if member.IsDeprecated() == *showDeprecated {
						filtered = append(filtered, member)
					}
				}
				members = filtered
			}

			if len(members) == 0 {
				continue
			}

			hasAnyContent = true
			title := typeTitles[memberType]
			if showDeprecated != nil && *showDeprecated {
				title = "Deprecated " + title
			}

			pterm.DefaultBasicText.Println("")
			pterm.DefaultHeader.WithFullWidth(false).WithBackgroundStyle(pterm.NewStyle(pterm.BgBlue)).Println(title)

			for _, member := range members {
				renderTreeItem(member)
			}
		}
	}

	if !hasAnyContent {
		if showDeprecated != nil && *showDeprecated {
			pterm.Info.Println("No deprecated members found.")
		} else if showDeprecated != nil && !*showDeprecated {
			pterm.Info.Println("No non-deprecated members found.")
		} else {
			pterm.Info.Println("No members found.")
		}
	}

	return nil
}

// renderTreeItem renders a single member in tree format.
func renderTreeItem(item M.RenderableMemberWithContext) {
	row := item.ToTableRow()
	if len(row) == 0 {
		return
	}

	// Create the tree item display
	name := row[0]
	details := []string{}
	
	// Add details based on member type
	switch item.GetMemberType() {
	case "attribute":
		if len(row) >= 4 {
			if row[1] != "" {
				details = append(details, fmt.Sprintf("Property: %s", row[1]))
			}
			if row[2] == "✅" {
				details = append(details, "Reflects")
			}
			if row[3] != "" {
				details = append(details, row[3])
			}
		}
	case "field":
		if len(row) >= 5 {
			if row[1] != "" {
				details = append(details, fmt.Sprintf("Type: %s", row[1]))
			}
			if row[2] != "public" {
				details = append(details, fmt.Sprintf("Privacy: %s", row[2]))
			}
			if row[3] == "true" {
				details = append(details, "Static")
			}
			if row[4] != "" {
				details = append(details, row[4])
			}
		}
	case "method":
		if len(row) >= 5 {
			if row[1] != "void" {
				details = append(details, fmt.Sprintf("Returns: %s", row[1]))
			}
			if row[2] != "public" {
				details = append(details, fmt.Sprintf("Privacy: %s", row[2]))
			}
			if row[3] == "true" {
				details = append(details, "Static")
			}
			if row[4] != "" {
				details = append(details, row[4])
			}
		}
	case "event":
		if len(row) >= 3 {
			if row[1] != "" {
				details = append(details, fmt.Sprintf("Type: %s", row[1]))
			}
			if row[2] != "" {
				details = append(details, row[2])
			}
		}
	case "slot":
		if len(row) >= 2 && row[1] != "" {
			details = append(details, row[1])
		}
	default:
		// For other types, just add remaining columns as details
		for i := 1; i < len(row); i++ {
			if row[i] != "" {
				details = append(details, row[i])
			}
		}
	}

	// Format the output
	prefix := "├─"
	if item.IsDeprecated() {
		pterm.Printf("  %s ⚠ %s (deprecated)\n", prefix, name)
	} else {
		pterm.Printf("  %s %s\n", prefix, name)
	}

	if len(details) > 0 {
		detailText := strings.Join(details, " • ")
		pterm.Printf("    %s\n", detailText)
	}
}