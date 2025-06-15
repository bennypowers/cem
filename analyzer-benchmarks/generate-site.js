import fs from 'fs';
import md from 'markdown-it';
const input = fs.readFileSync('README.md', 'utf8');
const html = md().render(input);

fs.mkdirSync('site', { recursive: true });
fs.writeFileSync('site/index.html', /*html*/`
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Analyzer Benchmarks</title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/github-markdown-css/github-markdown.min.css">
</head>
<body class="markdown-body">
${html}
</body>
</html>
`);
