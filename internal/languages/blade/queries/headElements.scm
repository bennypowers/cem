; Head element for finding insertion points
(element
  (start_tag
    (tag_name) @tag.name
    (#eq? @tag.name "head")
  ) @start.tag
  (end_tag)? @end.tag
) @head.element