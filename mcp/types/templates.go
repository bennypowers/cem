/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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
package types

// TemplateDataProvider defines the interface for all template data types
type TemplateDataProvider interface {
	// Core element information
	GetElement() ElementInfo
	GetContext() string
	GetOptions() map[string]string
}

// BaseTemplateData provides common template data fields for both tools and resources
type BaseTemplateData struct {
	Element           ElementInfo
	Context           string
	Options           map[string]string
	SchemaDefinitions interface{} // Schema data for template functions
}

// GetElement implements TemplateDataProvider
func (b BaseTemplateData) GetElement() ElementInfo { return b.Element }

// GetContext implements TemplateDataProvider
func (b BaseTemplateData) GetContext() string { return b.Context }

// GetOptions implements TemplateDataProvider
func (b BaseTemplateData) GetOptions() map[string]string { return b.Options }

// NewBaseTemplateData creates base template data
func NewBaseTemplateData(element ElementInfo, context string, options map[string]string) BaseTemplateData {
	return BaseTemplateData{
		Element: element,
		Context: context,
		Options: options,
	}
}

// NewBaseTemplateDataWithSchema creates base template data with schema context
func NewBaseTemplateDataWithSchema(element ElementInfo, context string, options map[string]string, schemaDefinitions interface{}) BaseTemplateData {
	return BaseTemplateData{
		Element:           element,
		Context:           context,
		Options:           options,
		SchemaDefinitions: schemaDefinitions,
	}
}

// FetchedData represents data fetched from various sources for template processing
type FetchedData map[string]any
