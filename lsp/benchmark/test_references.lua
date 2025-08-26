#!/usr/bin/env nvim --headless --clean -l
-- Quick test of go-to-references functionality
-- Usage: nvim --headless --clean -u configs/cem-minimal.lua -l test_references.lua

-- Ensure the benchmark directory is in the path
local script_dir = vim.fn.getcwd()
package.path = script_dir .. '/?.lua;' .. package.path

local function test_references()
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  local config = _G.BENCHMARK_LSP_CONFIG
  
  if not config then
    print("ERROR: No LSP configuration found.")
    return
  end
  
  print(string.format("Testing go-to-references for %s LSP server", server_name))
  
  local fixture_dir = script_dir .. '/fixtures/large_project'
  
  -- Test files with custom elements
  local test_files = {
    fixture_dir .. '/components/button-examples.html',
    fixture_dir .. '/pages/dashboard.html',
    fixture_dir .. '/components/card-examples.html'
  }
  
  local results = {
    success = false,
    files_tested = 0,
    references_found = 0,
    avg_response_time = 0,
    response_times = {},
    errors = {}
  }
  
  local success, error_msg = pcall(function()
    for _, test_file in ipairs(test_files) do
      if vim.fn.filereadable(test_file) == 0 then
        table.insert(results.errors, "File not readable: " .. test_file)
        goto continue
      end
      
      print(string.format("Testing references in: %s", vim.fn.fnamemodify(test_file, ':t')))
      
      -- Open the file
      vim.cmd('edit ' .. test_file)
      vim.wait(2000, function() return false end) -- Wait for LSP initialization
      
      results.files_tested = results.files_tested + 1
      
      -- Find some custom elements in the file to test
      local lines = vim.api.nvim_buf_get_lines(0, 0, -1, false)
      local custom_elements = {}
      
      -- Look for custom elements (tags with hyphens)
      for line_num, line in ipairs(lines) do
        local start_pos = 1
        while true do
          local tag_start, tag_end = string.find(line, '<([a-z]+-[a-z-]*)', start_pos)
          if not tag_start then break end
          
          local tag_name = string.match(line, '<([a-z]+-[a-z-]*)', start_pos)
          if tag_name then
            table.insert(custom_elements, {
              tag = tag_name,
              line = line_num,
              col = tag_start + 1 -- +1 to skip '<'
            })
          end
          start_pos = tag_end + 1
        end
      end
      
      print(string.format("Found %d custom elements to test", #custom_elements))
      
      -- Test references for up to 3 elements per file
      local elements_to_test = math.min(3, #custom_elements)
      for i = 1, elements_to_test do
        local element = custom_elements[i]
        
        -- Position cursor on the element
        vim.api.nvim_win_set_cursor(0, {element.line, element.col})
        
        print(string.format("Testing references for <%s> at line %d", element.tag, element.line))
        
        -- Try to get references
        local ref_start = vim.fn.reltime()
        local ref_success = pcall(function()
          vim.lsp.buf.references()
        end)
        local ref_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(ref_start)))
        
        table.insert(results.response_times, ref_time)
        
        if ref_success then
          results.references_found = results.references_found + 1
          print(string.format("  ✅ References call succeeded (%.3fs)", ref_time))
        else
          table.insert(results.errors, string.format("References failed for <%s>", element.tag))
          print(string.format("  ❌ References call failed"))
        end
        
        -- Small delay between requests
        vim.wait(500, function() return false end)
      end
      
      ::continue::
    end
    
    -- Calculate average response time
    if #results.response_times > 0 then
      local total_time = 0
      for _, time in ipairs(results.response_times) do
        total_time = total_time + time
      end
      results.avg_response_time = total_time / #results.response_times
    end
    
    results.success = true
  end)
  
  if success then
    print("\n" .. string.rep("=", 50))
    print("REFERENCES TEST SUMMARY")
    print(string.rep("=", 50))
    print(string.format("Files tested: %d", results.files_tested))
    print(string.format("References found: %d", results.references_found))
    print(string.format("Average response time: %.3fs", results.avg_response_time))
    
    if #results.errors > 0 then
      print("\nErrors:")
      for _, error in ipairs(results.errors) do
        print("  " .. error)
      end
    end
    
    local success_rate = results.references_found / math.max(1, #results.response_times)
    print(string.format("References success rate: %.1f%%", success_rate * 100))
    
    if success_rate >= 0.8 and results.avg_response_time < 1.0 then
      print("🎉 References performance is EXCELLENT")
    elseif success_rate >= 0.5 and results.avg_response_time < 2.0 then
      print("⚠️  References performance is GOOD")
    else
      print("🚨 References performance needs improvement")
    end
  else
    print(string.format("❌ Test failed: %s", error_msg))
  end
end

test_references()