# Design — In the Wild

Cinematic, cool-futuristic-with-nature. Dark burgundy base with warm sunset
glows and a cool periwinkle counterpoint — built from the two Dribbble refs
("Waterfall from Heaven" cool-warm diptych and "Wildlife Landing" burgundy +
sunset). Dark is default, with a light theme that flips to cream + plum.

Theme is remembered in `localStorage` and set on `<html data-theme>` before
paint, so there's no flash.

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
| `--header-bg` | `#F5EEDE` | `#241217` | Header strip (always inverts the page) |
| `--header-text` | `#241217` | `#F0E6D6` | Header text and nav |

Header is intentionally always the opposite of the page body — light bar
on dark theme, dark bar on light. It frames the content like a wildlife
poster.

## Typography

- **Headings:** Fraunces — editorial, slightly italic for scientific names.
- **Body:** Josefin Sans — soft geometric, airy; 300–700 loaded.
- Hero titles use `clamp(2.8rem, 7vw, 5.2rem)` with a tight 1.02 line-height.

## Page map

| Route | Page | Purpose |
|-------|------|---------|
| `/` | `pages.Home` | Hero carousel, slim stats strip, animals preview, globe/map toggle, gallery preview, charts |
| `/animals` | `pages.Animals` | Every photographed species as a tile, ranked by count |
| `/gallery` | `pages.Gallery` | Full masonry grid; click opens lightbox with "Open sighting" button |
| `/sightings/{id}` | `pages.Sighting` | One sighting: hero photo, prose + facts panel, related sightings from the same country |
| anything else | `pages.NotFound` | Themed 404 with links back |

The home nav highlights the current page via `aria-current="page"` on a sunset
underline. `BaseProps.Current` is `"home" / "animals" / "gallery" / "sighting" / "404"`.

## Key components

- **Hero carousel** (`components/Hero`) — Swiper, 1 slide at a time, auto-advance (5s, pauses on hover, disabled under reduced motion), prev/next circular buttons, slim bullet pagination. Photos take a gentle 0.78 brightness + a two-stop ink scrim — enough mood, not a wash.
- **Stats strip** (`.stats-strip`) — single horizontal line of totals under the hero, bordered top & bottom, ochre dots as separators.
- **Animal tile** (`.animal-tile`) — portrait 4:5 photo, caption fades out of a bottom scrim, hover lifts and brightens the border to sunset.
- **Globe / Map** (`components/ThreeScene` + `components/Map` in a `.geo` wrapper) — a floating pill-shaped tab toggle sits top-left over the view. Globe uses `web/assets/js/scenes/globe.js`: dark matte sphere, faint wireframe shell for the "futuristic grid" feel, periwinkle halo, sunset points per sighting. Rotates slowly, parallax on pointer, respects reduced motion.
- **Gallery masonry** (`.gallery-mason`) — CSS columns (1→4 depending on width), varied heights, `data-lightbox` opens the dialog; `data-sighting` is read by `lightbox.js` to show "Open sighting →".
- **Charts** (`components/Bars`, `components/Sparkline`) — unchanged plain markup, tinted with `--chart-bar` and `--chart-track`.
- **Sighting detail** (`.detail-hero` + `.detail-body` + `.related-strip`) — big hero photo with scrim and title, 2fr/1fr prose + facts grid, related photos from the same country below.

## Blocks installed

- `favicon` (SVG mark)
- `swiper` (via hero only — the standard swiper row isn't on any page)
- `charts`, `map`, `lightbox`, `reveal`, `transitions`, `threejs`

## Avoid

- Inline styles built from GBIF text (URLs, captions). URLs go through `templ.URL`; captions through `textContent`.
- A canvas per card — the globe is one `THREE.Points` for every point.
- Overriding theme colours with hard-coded hexes — use the tokens, so the toggle keeps working.
- Decorative emoji; use SVGs (the sun icon on the theme toggle is inline).

## Future

When adding a page, give it a `Current` string, hang it off `web.New`, and
add a nav link in `layout.templ`. Keep the header, nav, footer, lightbox and
globe import map in the layout — not in individual pages.
