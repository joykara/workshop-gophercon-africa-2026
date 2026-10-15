# Design — Night Sightings

Dark, field-notebook feel: photos glow against a near-black background with
sand-coloured type and warm ochre accents. "Savannah at night."

## Tokens

Set as CSS custom properties on `:root` in `web/assets/css/app.css`.

| Token | Value | Where it's used |
|-------|-------|-----------------|
| `--ink` | `#1A1612` | Page background, scrim bottom |
| `--ink-raised` | `#26211B` | Cards, stat tiles, chart cards, muted sections |
| `--sand` | `#F2E8D5` | Body text, headings |
| `--sand-dim` | `#BFB29A` | Secondary text, labels, nav links |
| `--ochre` | `#D4932C` | Primary accent, chart bars, map markers, section eyebrows |
| `--ochre-bright` | `#E8A33D` | Links, hover, active pagination bullet |
| `--acacia` | `#6B8E4E` | Reserved secondary accent (unused for now) |
| `--acacia-deep` | `#2D4A2B` | Reserved deep accent |
| `--border` | `#3A332A` | Hairlines around raised surfaces |

## Typography

- **Headings:** Fraunces, 500 for section titles, 700 for the hero. Italic 500 for the latin names. Loaded from Google Fonts.
- **Body:** Inter 400 / 500 / 600.
- Fluid sizes with `clamp()`: hero up to ~4rem, h2 up to ~2.1rem, body 16px/1.6.

## Spacing and shape

- Page gutter `--page-gutter: 28px`, max width `--page-max: 72rem`.
- Section vertical rhythm: `72px` top/bottom, `56px` on narrow screens.
- Card radius `--radius-lg: 18px`, stat/chart cards the same.
- Border on every raised surface: `1px solid var(--border)`.

## Components

- **Hero carousel** (`components/Hero`): one slide fills the viewport (clamped 460–720px), background image dimmed to 55% and tinted with a two-stop ink scrim, title + latin + one-line summary + "N sightings, seen in …". Prev/next arrows are circular, 48px, ochre on hover. Pagination is slim bars at the bottom. Autoplay runs on a 5s cycle and pauses on hover; disabled when `prefers-reduced-motion: reduce`.
- **Stats row** (`.stats`): auto-fit grid, 180px minimum. Each tile shows a label (uppercase, dim) and a big Fraunces number.
- **Charts** (`components/Bars`, `components/Sparkline`): plain markup from the charts block, tinted via `--chart-bar` and `--chart-track`.
- **Map** (`components/Map`): Leaflet, clustered markers, `map--dark` class to invert the tiles. `--map-marker` and `--map-marker-stroke` take the palette.
- **Sightings grid** (`.sightings`): auto-fill, 320px minimum. Card has a 4/3 cover photo, title in Fraunces, dim uppercase meta labels, warm sand values. Hover lifts 3px and brightens the border to ochre.

## Avoid

- Pure black (`#000`) backgrounds — use `--ink` so the shadow of real black can still read.
- Any inline style built from GBIF data (URLs, text). Images always go through `templ.URL`.
- Over-animating. Autoplay on the hero only; everything else is static or hover-only.

## Future sessions

If you add a page, extend `layout.BaseProps` (e.g. a `Description`) rather
than making a new layout. Keep the header, nav and footer in one place.
