/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
package cmd

type FieldValue = fieldValue

const FieldValueCustom = fieldValueCustom

var (
	NormalizeGitURL     = normalizeGitURL
	DetectSourceFiles   = detectSourceFiles
	DetectDemoFiles     = detectDemoFiles
	DetectCSSPatterns   = detectCSSPatterns
	DetectIndent        = detectIndent
	MarshalConfig       = marshalConfig
	SelectASCIILogo     = selectASCIILogo
	SplitCommaList      = splitCommaList
	ReorderYAMLMapping  = reorderYAMLMapping
	ValidatePackageSpec = validatePackageSpecifiers
)
