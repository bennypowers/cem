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

// DesignTokensLoader provides an interface for loading design tokens
// This interface allows dependency injection without circular imports between
// designtokens and workspace packages
type DesignTokensLoader interface {
	// Load loads design tokens from the workspace context
	// Returns the loaded design tokens and any error encountered
	Load(ctx any) (any, error)
}

// DesignTokensCache provides caching for design tokens to avoid redundant loading
type DesignTokensCache interface {
	// LoadOrReuse loads design tokens if not cached, or returns cached result
	LoadOrReuse(ctx WorkspaceContext) (any, error)
	// Clear resets the cache
	Clear()
}
