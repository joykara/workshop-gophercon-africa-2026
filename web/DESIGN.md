# Design — In the Wild

Cinematic, cool-futuristic-with-nature. Dark burgundy base with warm sunset
glows, cool periwinkle counterpoints, and sections framed as frosted glass
panels over full-bleed animal photography. The chrome floats on top — nav
is absolute and transparent, scrolling away with the page.

Dark theme is default. Theme is remembered in `localStorage` and applied
on `<html data-theme>` before paint so there's no flash.

## Tokens

Both themes live in `web/assets/css/app.css`. Pick a token, not a hex, when
writing styles — the toggle swaps both palettes under the same names.

| Token | Dark | Light | Role |
|-------|------|-------|------|
| `--bg` | `#241217` | `#F2EBDC` | Page background |
| `--bg-raised` | `#3A2024` | `#EBE1D0` | Cards, tiles |
| `--bg-sunk` | `#1C0D10` | `#E6DBC5` | Hero backdrop, sunk sections |
| `--text` | `#F0E6D6` | `#241217` | Body text, headings |
| `--text-dim` | `#B9A99A` | `#6E4B4F` | Labels, secondary text |
| `--sunset` | `#F0A858` | `#E89B2E` | Primary accent, chart bars, map markers |
| `--sunset-bright` | `#FFC077` | `#CA7D13` | Hover, bright highlights |
| `--periwinkle` | `#8599D6` | `#5A6DB5` | Cool secondary accent (globe, lines) |

The header is **absolutely positioned and transparent** on the home page
and on sighting detail (both have a photo at the top). On `/animals`,
`/gallery` and the 404 page, the `.site-header--inner` modifier adds a
very subtle gradient scrim so the nav still reads over the light body.
Pages without a hero use `.has-header-space` on the first section to
push content past the floating nav.

Nav layout: `grid-template-columns: 1fr auto 1fr` — wordmark on the
left (justify-self: start), nav links **centered**, theme toggle on
the right (justify-self: end).

## Typography

- **Headings — Almendra Display.** A tall, slightly medieval serif. Used
  for the hero title, section titles, pattern card names and the
  wordmark.
- **Body — Playpen Sans.** A rounded, hand-drawn-feeling sans. Used for
  body text, labels, nav links and UI chrome. Weight 400 at body size,
  500–600 for emphasis.
- **Fraunces** stays loaded as `--font-ui` for places where numbers or
  smaller type need a steadier serif (chart card headings, the stats
  strip hint — mostly a fallback token).

## Page map

| Route | Page | Purpose |
|-------|------|---------|
| `/` | `pages.Home` | Hero carousel, stats strip, five frosted-glass sections over animal photo backdrops: Animals · Geo (globe + map) · Gallery · Patterns · (sections live inside `.section--photo`) |
| `/animals` | `pages.Animals` | Every photographed species as a masonry tile |
| `/gallery` | `pages.Gallery` | Full masonry grid; lightbox has "Open sighting →" |
| `/sightings/{id}` | `pages.Sighting` | Hero photo, prose + facts, masonry related strip |
| anything else | `pages.NotFound` | Themed 404 with links back |

`BaseProps.Current` is `"home" | "animals" | "gallery" | "sighting" | "404"`;
nav highlights the current page with a sunset underline.

## Key components and effects

### Hero (`components/Hero`)
- Full-width Swiper carousel; one slide per animal.
- Each slide: a 3D-perspective `.hero__photo-wrap` wrapping the photo so
  it tilts on pointer move (hero.js updates `--tilt-x` / `--tilt-y`).
- **Ken Burns** CSS animation on `.hero__photo` scales + translates the
  photo over 14s on loop.
- **Kinetic title**: `kineticTitle()` splits the name into one span per
  character; a CSS `letter-in` keyframe staggers each in by 40ms (capped
  at 1200ms) when `.swiper-slide-active` arrives.
- **Frosted glass panel** on the content block (`.hero__content-inner`):
  blurred rgba, warm plum tint, soft shadow.
- Autoplay 5.5s, pauses on hover, respects reduced motion.

### Sections with photo backgrounds (`.section--photo`)
- Each major home section has a full-bleed `<img>` in `.section__bg`,
  dimmed to 50%, overlaid with a plum gradient.
- Content lives inside a `.glass` panel: `background: rgba(28,13,16,0.38)`,
  `backdrop-filter: blur(16px)`, border, soft shadow.
- Photos are picked by `sectionBG(sightings, hints...)`: it searches the
  real sightings for a vernacular name containing any hint, falling back
  to the first photo. This never leaks GBIF text into inline styles —
  the URL goes through `templ.URL` on an `<img>`.

### Grids (irregular / staggered)
- The animal row and the related strip now use **CSS columns** (1→4
  responsively) with per-nth-child aspect ratios (3/4, 4/5, 1/1, 5/7)
  so the stack looks intentionally irregular rather than tiled.
- The gallery was already CSS columns; unchanged.

### Globe (`three_scene` + `scenes/globe.js`)
- Matte sphere + periwinkle wire shell + sunset halo + one THREE.Points
  per sighting from `[lat, lng]` JSON.
- **Legend** panel floats bottom-right over the view (`.globe-legend`):
  three rows explaining the sunset point, the periwinkle rim, and the
  drag/hover controls.

### Flat map (Leaflet + markercluster)
- Starts at `[20, 0], zoom 2` (world view) when `World: true` is passed
  on `MapProps`. The home panel sets it.
- `.geo__view[data-view="map"]::after` adds an inset vignette so it
  reads lit-from-below. `.leaflet-marker-icon` gets a sunset drop-shadow;
  `.marker-cluster` gets a sunset gradient + glow.

### Patterns (`.pattern-featured`, `.pattern-cards`)
- Featured year sparkline at the top on its own glass card.
- Below: two groups of up to 6 cards. Each card has a very faint animal
  photo in the background (`opacity: 0.22`, brightens on hover), a rank
  number, the name in Almendra, a mini sparkline for that species or
  country's own counts per year, and the total in large sunset numerals.

## Avoid

- Inline styles built from GBIF text (URLs, captions). URLs go through
  `templ.URL` on `<img>` elements; captions through `textContent`.
- Hard-coded hexes in new rules — use the tokens so the toggle keeps
  working.
- A canvas per card. The globe is one `THREE.Points` for every point.
- Fraunces or Josefin as the primary pair now — those are superseded.
  Almendra Display + Playpen Sans is the new direction.

## Future

New pages: give them a `Current` string, add the handler + route in
`web.New`, and add a link in `layout.templ`. If the page has no hero,
add `has-header-space` to the first section so content clears the
floating nav. If the page has a hero, the nav sits over it naturally.
