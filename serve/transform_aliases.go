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

package serve

import "bennypowers.dev/cem/serve/middleware/transform"

// Type aliases for backward compatibility with existing code and tests
// These types now live in the transform middleware package

type (
	// Transform engine types
	Loader        = transform.Loader
	Target        = transform.Target
	SourceMapMode = transform.SourceMapMode

	// Transform cache types
	TransformCache = transform.Cache
	CacheKey       = transform.CacheKey
	CacheEntry     = transform.CacheEntry
	CacheStats     = transform.CacheStats

	// Transform options and results
	TransformOptions = transform.TransformOptions
	TransformResult  = transform.TransformResult
)

// Constants
const (
	// Loaders
	LoaderTS  = transform.LoaderTS
	LoaderTSX = transform.LoaderTSX
	LoaderJS  = transform.LoaderJS
	LoaderJSX = transform.LoaderJSX

	// Targets
	ES2015 = transform.ES2015
	ES2016 = transform.ES2016
	ES2017 = transform.ES2017
	ES2018 = transform.ES2018
	ES2019 = transform.ES2019
	ES2020 = transform.ES2020
	ES2021 = transform.ES2021
	ES2022 = transform.ES2022
	ES2023 = transform.ES2023
	ESNext = transform.ESNext

	// Source map modes
	SourceMapInline   = transform.SourceMapInline
	SourceMapExternal = transform.SourceMapExternal
	SourceMapNone     = transform.SourceMapNone
)

// Functions
var (
	NewTransformCache   = transform.NewCache
	TransformTypeScript = transform.TransformTypeScript
	TransformCSS        = transform.TransformCSS
	ValidTargets        = transform.ValidTargets
	IsValidTarget       = transform.IsValidTarget
)
