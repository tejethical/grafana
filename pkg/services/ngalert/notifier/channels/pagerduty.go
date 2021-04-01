package channels

import (
	"context"
	"encoding/json"
	"os"

	"github.com/pkg/errors"
	"github.com/prometheus/alertmanager/notify"
	"github.com/prometheus/alertmanager/types"
	"github.com/prometheus/common/model"

	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/alerting"
	old_notifiers "github.com/grafana/grafana/pkg/services/alerting/notifiers"
)

const (
	pagerdutyEventAPIURL  = "https://events.pagerduty.com/v2/enqueue"
	pagerDutyEventTrigger = "trigger"
	pagerDutyEventResolve = "resolve"
)

// NewPagerdutyNotifier is the constructor for the PagerDuty notifier
func NewPagerdutyNotifier(model *models.AlertNotification) (*PagerdutyNotifier, error) {
	// TODO: validate these.
	severity := model.Settings.Get("severity").MustString("critical")
	autoResolve := model.Settings.Get("autoResolve").MustBool(true)
	key := model.DecryptedValue("integrationKey", model.Settings.Get("integrationKey").MustString())
	messageInDetails := model.Settings.Get("messageInDetails").MustBool(false)
	if key == "" {
		return nil, alerting.ValidationError{Reason: "Could not find integration key property in settings"}
	}

	return &PagerdutyNotifier{
		NotifierBase:     old_notifiers.NewNotifierBase(model),
		Key:              key,
		Severity:         severity,
		AutoResolve:      autoResolve,
		MessageInDetails: messageInDetails,
		log:              log.New("alerting.notifier." + model.Name),
	}, nil
}

// PagerdutyNotifier is responsible for sending
// alert notifications to pagerduty
type PagerdutyNotifier struct {
	old_notifiers.NotifierBase
	Key              string
	Severity         string
	AutoResolve      bool
	MessageInDetails bool
	log              log.Logger
}

// Notify sends an alert notification to PagerDuty
func (pn *PagerdutyNotifier) Notify(ctx context.Context, as ...*types.Alert) (bool, error) {
	key, err := notify.ExtractGroupKey(ctx)
	if err != nil {
		return false, err
	}

	alerts := types.Alerts(as...)
	eventType := pagerDutyEventTrigger
	if alerts.Status() == model.AlertResolved {
		eventType = pagerDutyEventResolve
	}

	if alerts.Status() == model.AlertResolved && !pn.AutoResolve {
		pn.log.Info("Not sending a trigger to Pagerduty", "status", alerts.Status(), "auto resolve", pn.AutoResolve)
		return true, nil
	}

	pn.log.Info("Notifying Pagerduty", "event_type", eventType)
	msg := &pagerDutyMessage{
		Client:      "Grafana",
		ClientURL:   "TODO URL", // TODO: external URL, but which one?
		RoutingKey:  pn.Key,
		EventAction: eventType,
		DedupKey:    key.Hash(),
		Links:       make([]pagerDutyLink, 0), // TODO: preconfigured links? Or external URLs of all alerts?
		Payload: &pagerDutyPayload{
			Component:     "Grafana",
			Summary:       "TODO", // TODO: preconfigured description with template?
			Severity:      pn.Severity,
			CustomDetails: nil,    // TODO: preconfigured template with data from alerts?
			Class:         "TODO", // TODO: preconfigured?
			Group:         "TODO", // TODO: preconfigured?
		},
	}

	if hostname, err := os.Hostname(); err == nil {
		// TODO: should this be configured like in Prometheus AM?
		msg.Payload.Source = hostname
	}

	body, err := json.Marshal(&msg)
	if err != nil {
		return false, errors.Wrap(err, "marshal json")
	}

	cmd := &models.SendWebhookSync{
		Url:        pagerdutyEventAPIURL,
		Body:       string(body),
		HttpMethod: "POST",
		HttpHeader: map[string]string{
			"Content-Type": "application/json",
		},
	}

	if err := bus.DispatchCtx(ctx, cmd); err != nil {
		pn.log.Error("Failed to send notification to Pagerduty", "error", err, "body", string(body))
		return false, err
	}

	return true, nil
}

type pagerDutyMessage struct {
	RoutingKey  string            `json:"routing_key,omitempty"`
	ServiceKey  string            `json:"service_key,omitempty"`
	DedupKey    string            `json:"dedup_key,omitempty"`
	IncidentKey string            `json:"incident_key,omitempty"`
	EventType   string            `json:"event_type,omitempty"`
	Description string            `json:"description,omitempty"`
	EventAction string            `json:"event_action"`
	Payload     *pagerDutyPayload `json:"payload"`
	Client      string            `json:"client,omitempty"`
	ClientURL   string            `json:"client_url,omitempty"`
	Details     map[string]string `json:"details,omitempty"`
	Images      []pagerDutyImage  `json:"images,omitempty"`
	Links       []pagerDutyLink   `json:"links,omitempty"`
}

type pagerDutyLink struct {
	HRef string `json:"href"`
	Text string `json:"text"`
}

type pagerDutyImage struct {
	Src  string `json:"src"`
	Alt  string `json:"alt"`
	Href string `json:"href"`
}

type pagerDutyPayload struct {
	Summary       string            `json:"summary"`
	Source        string            `json:"source"`
	Severity      string            `json:"severity"`
	Timestamp     string            `json:"timestamp,omitempty"`
	Class         string            `json:"class,omitempty"`
	Component     string            `json:"component,omitempty"`
	Group         string            `json:"group,omitempty"`
	CustomDetails map[string]string `json:"custom_details,omitempty"`
}

func (pn *PagerdutyNotifier) SendResolved() bool {
	return pn.AutoResolve
}
