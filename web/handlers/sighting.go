package handlers

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"uuid"

	"workshop/internal/domain/sighting"
	"workshop/web/views/pages"
)

type SightingFinder interface {
	Find(ctx context.Context, id uuid.UUID) (sighting.Sighting, error)
	List(ctx context.Context, filter sighting.ListFilter) ([]sighting.Sighting, error)
}

func Sighting(logger *slog.Logger, svc SightingFinder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			renderNotFound(ctx, w, logger)
			return
		}

		s, err := svc.Find(ctx, id)
		if err != nil {
			if errors.Is(err, sighting.ErrNotFound) {
				renderNotFound(ctx, w, logger)
				return
			}
			logger.ErrorContext(ctx, "handlers: finding sighting: "+err.Error())
			renderError(ctx, w, logger, http.StatusInternalServerError, "Unable to load sighting.")
			return
		}

		related := relatedSightings(ctx, svc, s, 6)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err = pages.Sighting(s, related).Render(ctx, w); err != nil {
			logger.ErrorContext(ctx, "handlers: rendering sighting: "+err.Error())
			renderError(ctx, w, logger, http.StatusInternalServerError, "Unable to render page.")
		}
	}
}

// relatedSightings finds up to n other sightings from the same country,
// prefers ones with photos, and never includes the current sighting.
func relatedSightings(ctx context.Context, svc SightingFinder, s sighting.Sighting, n int) []sighting.Sighting {
	if s.Location.CountryCode == "" {
		return nil
	}
	all, err := svc.List(ctx, sighting.ListFilter{
		Country: s.Location.CountryCode,
		Limit:   100,
	})
	if err != nil {
		return nil
	}
	out := make([]sighting.Sighting, 0, n)
	for _, r := range all {
		if r.ID == s.ID {
			continue
		}
		hasPhoto := false
		for _, m := range r.Media {
			if m.Type == "StillImage" && m.URL != "" {
				hasPhoto = true
				break
			}
		}
		if !hasPhoto {
			continue
		}
		out = append(out, r)
		if len(out) >= n {
			break
		}
	}
	return out
}

func renderNotFound(ctx context.Context, w http.ResponseWriter, logger *slog.Logger) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	if err := pages.NotFound().Render(ctx, w); err != nil {
		logger.ErrorContext(ctx, "handlers: rendering 404: "+err.Error())
	}
}

func NotFound(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderNotFound(r.Context(), w, logger)
	}
}
