/**
 * Fetches health data and adds score badges to listing page cards.
 * Cards must have a data-tag-name attribute to be matched.
 */
const STATUS_COLORS = { pass: 'green', warn: 'orange', fail: 'red' };

const badgeTemplate = document.createElement('template');
badgeTemplate.innerHTML = `
  <pf-v6-popover slot="title"
                 position="bottom"
                 trigger-action="hover">
    <pf-v6-label slot="trigger"
                 style="cursor: default"></pf-v6-label>
    <span slot="header"></span>
    <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact"></dl>
  </pf-v6-popover>
`;

const categoryTemplate = document.createElement('template');
categoryTemplate.innerHTML = `
  <div class="pf-v6-c-description-list__group">
    <dt class="pf-v6-c-description-list__term"></dt>
    <dd class="pf-v6-c-description-list__description"></dd>
  </div>
`;

async function addHealthBadges() {
  const cards = document.querySelectorAll('.listing-grid > pf-v6-card[data-tag-name]');
  if (cards.length === 0) return;

  let result;
  try {
    const response = await fetch('/__cem/api/health');
    if (!response.ok) return;
    result = await response.json();
  } catch {
    return;
  }

  // Build a map of tag name -> component report
  const scoreMap = new Map();
  for (const mod of result.modules) {
    for (const decl of mod.declarations) {
      const key = decl.tagName || decl.name;
      if (key) scoreMap.set(key, decl);
    }
  }

  for (const card of cards) {
    const tagName = card.dataset.tagName;
    const decl = scoreMap.get(tagName);
    if (!decl) continue;

    const pct = decl.maxScore > 0
      ? Math.round((decl.score / decl.maxScore) * 100)
      : 0;
    const status = pct >= 80 ? 'pass' : pct >= 40 ? 'warn' : 'fail';

    const fragment = badgeTemplate.content.cloneNode(true);
    const label = fragment.querySelector('pf-v6-label');
    label.setAttribute('color', STATUS_COLORS[status]);
    label.textContent = `${pct}%`;

    const header = fragment.querySelector('[slot="header"]');
    header.textContent = `Health: ${pct}%`;

    const dl = fragment.querySelector('dl');
    for (const cat of decl.categories) {
      const catPct = cat.maxPoints > 0
        ? Math.round((cat.points / cat.maxPoints) * 100)
        : 0;
      const icon = cat.status === 'pass' ? '✓'
        : cat.status === 'warn' ? '⚠' : '✗';

      const row = categoryTemplate.content.cloneNode(true);
      row.querySelector('dt').textContent = cat.category;
      row.querySelector('dd').textContent = `${icon} ${catPct}% (${cat.points}/${cat.maxPoints})`;
      dl.append(row);
    }

    card.append(fragment);
  }
}

addHealthBadges();
