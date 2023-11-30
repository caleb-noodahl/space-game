package events

type EnvironmentalStorageNotificationPayload struct {
	StorageType string  `json:"storage_type"`
	Supply      float64 `json:"supply"`
	Capacity    float64 `json:"capacity"`
}
