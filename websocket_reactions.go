package slack

// ReactionItem is a lighter-weight item than is returned by the reactions list.
type ReactionItem struct {
	Type        string `json:"type"`
	Channel     string `json:"channel,omitempty"`
	File        string `json:"file,omitempty"`
	FileComment string `json:"file_comment,omitempty"`
	Timestamp   string `json:"ts,omitempty"`
}

// ReactionEvent is a event that reaction added or removed
type ReactionEvent struct {
	Type           string       `json:"type"`
	User           string       `json:"user"`
	ItemUser       string       `json:"item_user"`
	Item           ReactionItem `json:"item"`
	Reaction       string       `json:"reaction"`
	EventTimestamp string       `json:"event_ts"`
}

// Type of ReactionEvent
const (
	ReactionAdded   = "reaction_added"
	ReactionRemoved = "reaction_removed"
)

// // ReactionAddedEvent represents the Reaction added event
// type ReactionAddedEvent ReactionEvent

// // ReactionRemovedEvent represents the Reaction removed event
// type ReactionRemovedEvent ReactionEvent
