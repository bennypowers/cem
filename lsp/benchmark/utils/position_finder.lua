--[[
Position Finder Utility for LSP Benchmarks
Uses Neovim's tree-sitter API to dynamically find element and attribute positions
Replaces hardcoded line/character positions with semantic position finding
]]

local M = {}

--- Find position of HTML/TS element by tag name
--- @param file_path string Path to HTML/TypeScript file
--- @param element_name string Tag name to find (e.g., "my-button")
--- @param occurrence number|nil Which occurrence to find (1-indexed, default: 1)
--- @return table|nil position {line=N, character=M} (0-indexed) or nil if not found
--- @return string|nil error Error message if failed
function M.find_element_position(file_path, element_name, occurrence)
	occurrence = occurrence or 1

	-- Read file content
	local ok_read, lines = pcall(vim.fn.readfile, file_path)
	if not ok_read then
		return nil, string.format("Failed to read file: %s", file_path)
	end

	local content = table.concat(lines, "\n")

	-- Detect language from file extension
	local lang = file_path:match("%.html$") and "html" or "typescript"

	-- Parse with tree-sitter
	local ok_parser, parser = pcall(vim.treesitter.get_string_parser, content, lang)
	if not ok_parser then
		return nil, string.format("Tree-sitter parser unavailable for %s", lang)
	end

	local trees = parser:parse()
	if not trees or #trees == 0 then
		return nil, "Failed to parse file with tree-sitter"
	end

	local tree = trees[1]
	local root = tree:root()

	-- Query for matching elements
	local query_str = [[
    (element
      (start_tag
        (tag_name) @tag.name))
  ]]

	local ok_query, query = pcall(vim.treesitter.query.parse, lang, query_str)
	if not ok_query then
		return nil, string.format("Failed to parse tree-sitter query: %s", query)
	end

	local count = 0

	for id, node in query:iter_captures(root, content, 0, -1) do
		local name = query.captures[id]
		if name == "tag.name" then
			local text = vim.treesitter.get_node_text(node, content)
			if text == element_name then
				count = count + 1
				if count == occurrence then
					local row, col, _, _ = node:range()
					return { line = row, character = col }, nil
				end
			end
		end
	end

	return nil, string.format("Element '%s' occurrence %d not found in %s", element_name, occurrence, file_path)
end

--- Find position of attribute within HTML element
--- @param file_path string Path to HTML file
--- @param element_name string Tag name containing the attribute
--- @param attr_name string Attribute name to find
--- @param occurrence number|nil Which occurrence of element to search (default: 1)
--- @return table|nil position {line=N, character=M} (0-indexed) or nil if not found
--- @return string|nil error Error message if failed
function M.find_attribute_position(file_path, element_name, attr_name, occurrence)
	occurrence = occurrence or 1

	-- Read file content
	local ok_read, lines = pcall(vim.fn.readfile, file_path)
	if not ok_read then
		return nil, string.format("Failed to read file: %s", file_path)
	end

	local content = table.concat(lines, "\n")

	-- Detect language (currently only HTML supported for attributes)
	local lang = file_path:match("%.html$") and "html" or "typescript"

	-- Parse with tree-sitter
	local ok_parser, parser = pcall(vim.treesitter.get_string_parser, content, lang)
	if not ok_parser then
		return nil, string.format("Tree-sitter parser unavailable for %s", lang)
	end

	local trees = parser:parse()
	if not trees or #trees == 0 then
		return nil, "Failed to parse file with tree-sitter"
	end

	local tree = trees[1]
	local root = tree:root()

	-- Query for attributes in matching elements
	local query_str = [[
    (element
      (start_tag
        (tag_name) @tag.name
        (attribute
          (attribute_name) @attr.name)))
  ]]

	local ok_query, query = pcall(vim.treesitter.query.parse, lang, query_str)
	if not ok_query then
		return nil, string.format("Failed to parse tree-sitter query: %s", query)
	end

	-- Track elements we've seen (by node ID to avoid duplicates)
	local seen_elements = {}
	local element_count = 0
	local target_element_node = nil

	-- First pass: find the target element node
	local tag_query_str = [[(element (start_tag (tag_name) @tag.name))]]
	local tag_query = vim.treesitter.query.parse(lang, tag_query_str)

	for id, node in tag_query:iter_captures(root, content, 0, -1) do
		if tag_query.captures[id] == "tag.name" then
			local text = vim.treesitter.get_node_text(node, content)
			local node_id = node:id()

			-- Only count each unique tag node once
			if text == element_name and not seen_elements[node_id] then
				seen_elements[node_id] = true
				element_count = element_count + 1

				if element_count == occurrence then
					-- Get the parent element node
					target_element_node = node:parent():parent()
					break
				end
			end
		end
	end

	if not target_element_node then
		return nil,
			string.format(
				"Element '%s' occurrence %d not found in %s",
				element_name,
				occurrence,
				file_path
			)
	end

	-- Second pass: find the attribute within the target element
	local attr_query_str = [[(attribute (attribute_name) @attr.name)]]
	local attr_query = vim.treesitter.query.parse(lang, attr_query_str)

	for id, node in attr_query:iter_captures(target_element_node, content, 0, -1) do
		if attr_query.captures[id] == "attr.name" then
			local text = vim.treesitter.get_node_text(node, content)
			if text == attr_name then
				local row, col, _, _ = node:range()
				return { line = row, character = col }, nil
			end
		end
	end

	return nil,
		string.format(
			"Attribute '%s' in element '%s' (occurrence %d) not found in %s",
			attr_name,
			element_name,
			occurrence,
			file_path
		)
end

--- Batch find multiple element positions
--- @param file_path string Path to file
--- @param elements table Array of {element=string, occurrence=number|nil}
--- @return table results Array of {element=string, line=number, character=number, error=string|nil}
function M.find_elements_batch(file_path, elements)
	local results = {}
	for _, spec in ipairs(elements) do
		local pos, err = M.find_element_position(file_path, spec.element, spec.occurrence)
		if pos then
			table.insert(results, {
				element = spec.element,
				line = pos.line,
				character = pos.character,
			})
		else
			table.insert(results, {
				element = spec.element,
				error = err,
			})
		end
	end
	return results
end

--- Batch find multiple attribute positions
--- @param file_path string Path to file
--- @param attributes table Array of {element=string, attr=string, occurrence=number|nil, expected_content=string|nil}
--- @return table results Array of {element=string, attr=string, line=number, character=number, expected_content=string|nil, error=string|nil}
function M.find_attributes_batch(file_path, attributes)
	local results = {}
	for _, spec in ipairs(attributes) do
		local pos, err = M.find_attribute_position(file_path, spec.element, spec.attr, spec.occurrence)
		if pos then
			table.insert(results, {
				element = spec.element,
				attr = spec.attr,
				line = pos.line,
				character = pos.character,
				expected_content = spec.expected_content,
			})
		else
			table.insert(results, {
				element = spec.element,
				attr = spec.attr,
				error = err,
			})
		end
	end
	return results
end

return M
