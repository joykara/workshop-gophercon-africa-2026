package pages

import (
	"fmt"
	"strings"

	"workshop/internal/domain/sighting"
	"workshop/internal/domain/species"
	"workshop/web/views/components"
)

// AnimalTile is the compact card used in the home preview row and on
// the /animals page. Built from a group of sightings for one species.
type AnimalTile struct {
	Name       string
	Latin      string
	PhotoURL   string
	Count      int
	SightingID string // first sighting of this species that has a photo
}

// PhotoTile is one entry in the gallery: a photo and the sighting it
// belongs to, so the lightbox can link through.
type PhotoTile struct {
	URL      string
	Alt      string
	Creator  string
	SightingID string
	Species  string
}

// animalTiles groups sightings by species, keeps only ones with a
// photo, and ranks by count descending.
func animalTiles(sightings []sighting.Sighting, limit int) []AnimalTile {
	groups := make(map[string][]sighting.Sighting)
	order := make([]string, 0)
	for _, s := range sightings {
		key := s.Species.ScientificName
		if key == "" {
			continue
		}
		if _, ok := groups[key]; !ok {
			order = append(order, key)
		}
		groups[key] = append(groups[key], s)
	}

	tiles := make([]AnimalTile, 0, len(order))
	for _, k := range order {
		g := groups[k]
		photo, idx := firstPhotoWithIndex(g)
		if photo == "" {
			continue
		}
		sp := g[0].Species
		tiles = append(tiles, AnimalTile{
			Name:       animalName(sp),
			Latin:      sp.ScientificName,
			PhotoURL:   photo,
			Count:      len(g),
			SightingID: g[idx].ID.String(),
		})
	}

	for i := 0; i < len(tiles); i++ {
		best := i
		for j := i + 1; j < len(tiles); j++ {
			if tiles[j].Count > tiles[best].Count {
				best = j
			}
		}
		tiles[i], tiles[best] = tiles[best], tiles[i]
	}

	if limit > 0 && len(tiles) > limit {
		tiles = tiles[:limit]
	}
	return tiles
}

// photoTiles picks the first photo of each sighting that has one, up
// to limit (0 = all). Used for the gallery and the home preview strip.
func photoTiles(sightings []sighting.Sighting, limit int) []PhotoTile {
	tiles := make([]PhotoTile, 0, len(sightings))
	for _, s := range sightings {
		for _, m := range s.Media {
			if m.Type != "StillImage" || m.URL == "" {
				continue
			}
			tiles = append(tiles, PhotoTile{
				URL:        m.URL,
				Alt:        animalName(s.Species),
				Creator:    m.Creator,
				SightingID: s.ID.String(),
				Species:    animalName(s.Species),
			})
			break
		}
		if limit > 0 && len(tiles) >= limit {
			break
		}
	}
	return tiles
}

// globePoints renders every sighting's position as a [lat, lng] pair
// for the three.js globe. The JSON that templ.JSONScript writes stays
// small because each entry is just two numbers.
func globePoints(sightings []sighting.Sighting) [][2]float64 {
	out := make([][2]float64, 0, len(sightings))
	for _, s := range sightings {
		if s.Location.Latitude == 0 && s.Location.Longitude == 0 {
			continue
		}
		out = append(out, [2]float64{s.Location.Latitude, s.Location.Longitude})
	}
	return out
}


// topAnimals ranks the species seen most, by readable name.
func topAnimals(sightings []sighting.Sighting, n int) []sighting.Tally {
	return sighting.Top(sighting.CountBy(sightings, func(s sighting.Sighting) string {
		return animalName(s.Species)
	}), n)
}

// topCountries ranks the countries with the most sightings.
func topCountries(sightings []sighting.Sighting, n int) []sighting.Tally {
	return sighting.Top(sighting.CountBy(sightings, func(s sighting.Sighting) string {
		return s.Location.Country
	}), n)
}

// heroItems picks the top-k most-sighted animals that have a photo, and
// builds a slide for each: name, latin, a one-line summary, cover photo,
// how many times seen, and up to three countries seen in.
func heroItems(sightings []sighting.Sighting, n int) []components.HeroItem {
	groups := make(map[string][]sighting.Sighting)
	order := make([]string, 0)
	for _, s := range sightings {
		key := s.Species.ScientificName
		if key == "" {
			continue
		}
		if _, ok := groups[key]; !ok {
			order = append(order, key)
		}
		groups[key] = append(groups[key], s)
	}

	type ranked struct {
		key   string
		count int
	}
	ranks := make([]ranked, 0, len(groups))
	for _, k := range order {
		ranks = append(ranks, ranked{k, len(groups[k])})
	}
	for i := 0; i < len(ranks); i++ {
		best := i
		for j := i + 1; j < len(ranks); j++ {
			if ranks[j].count > ranks[best].count {
				best = j
			}
		}
		ranks[i], ranks[best] = ranks[best], ranks[i]
	}

	items := make([]components.HeroItem, 0, n)
	for _, r := range ranks {
		g := groups[r.key]
		photo := firstPhoto(g)
		if photo == "" {
			continue
		}
		sp := g[0].Species
		name := animalName(sp)
		items = append(items, components.HeroItem{
			Title:    name,
			Latin:    sp.ScientificName,
			Summary:  heroSummary(sp, r.count, g),
			PhotoURL: photo,
			Alt:      name + " photographed in the wild",
			Count:    r.count,
			Places:   countriesSeen(g, 3),
		})
		if n > 0 && len(items) >= n {
			break
		}
	}
	return items
}

func firstPhoto(group []sighting.Sighting) string {
	url, _ := firstPhotoWithIndex(group)
	return url
}

func firstPhotoWithIndex(group []sighting.Sighting) (string, int) {
	for i, s := range group {
		for _, m := range s.Media {
			if m.Type == "StillImage" && m.URL != "" {
				return m.URL, i
			}
		}
	}
	return "", 0
}

func countriesSeen(group []sighting.Sighting, limit int) string {
	seen := make(map[string]struct{})
	out := make([]string, 0)
	for _, s := range group {
		c := s.Location.Country
		if c == "" {
			continue
		}
		if _, ok := seen[c]; ok {
			continue
		}
		seen[c] = struct{}{}
		out = append(out, c)
		if len(out) >= limit {
			break
		}
	}
	return strings.Join(out, ", ")
}

func heroSummary(sp species.Species, count int, g []sighting.Sighting) string {
	group := sp.Family
	if group == "" {
		group = sp.Order
	}
	if group == "" {
		group = sp.Class
	}
	first := g[len(g)-1]
	latest := g[0]
	span := fmt.Sprintf("between %d and %d", first.HappenedAt.Year(), latest.HappenedAt.Year())
	if first.HappenedAt.Year() == latest.HappenedAt.Year() {
		span = fmt.Sprintf("in %d", first.HappenedAt.Year())
	}
	if group != "" {
		return fmt.Sprintf("A member of the %s family, recorded %d times %s across our dataset.", group, count, span)
	}
	return fmt.Sprintf("Recorded %d times %s across our dataset.", count, span)
}

func animalName(sp species.Species) string {
	if sp.VernacularName != "" {
		return sp.VernacularName
	}
	if sp.CanonicalName != "" {
		return sp.CanonicalName
	}
	return sp.ScientificName
}
