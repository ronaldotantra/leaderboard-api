package healthcheck

type Status string

const (
	StatusHealthy   = "Healthy"
	StatusUnhealthy = "Unhealthy"
)

type HealthReport struct {
	Status     Status            `json:"status"`
	Components []component       `json:"components"`
	Failures   map[string]string `json:"failures"`
}

func (r HealthReport) MapReport() map[string]any {
	components := make([]string, 0, len(r.Components))
	for _, cmp := range r.Components {
		components = append(components, cmp.Name)
	}

	return map[string]any{
		"status":     r.Status,
		"components": components,
		"failures":   r.Failures,
	}
}
