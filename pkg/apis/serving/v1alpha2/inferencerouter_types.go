// ABTestSpec defines parameters required for an A/B test.
type ABTestSpec {
	// Name of numeric metric we are A/B testing for improvement.
	// +required
	MetricName string `json:"metricName"`
	// Absolute minimum improvement in the metric that should yield a positive result from the A/B test.
	// +required
	MinimumDetectableEffect float `json:"minimumDetectableEffect"`
	// Percent chance that if the minimum detectable effect threshold is not met, the test yields a negative result.
	// +optional
	Confidence int `json:"confidence,omitempty"`
	// Percent chance that if the minimum detectable effect threshold is met, the test yields a poistive result.
	// +optional
	Power int `json:"power,omitempty"`
	// Base group's metric value, if known a priori. Ignores MaximumPercentError and EstimationConfidence if set.
	// +optional
	BaseRate float `json:"baseRate,omitempty"`
	// Largest allowed percent deviation from the true metric value in base rate metric estimation.
	// +optional
	MaximumPercentError int `json:"maximumPercentError,omitempty"`
	// Confidence that the metric estimate deviates from its true value by no more than MaximumPercentError.
	// +optional
	EstimationConfidence `json:"estimationConfidence,omitempty"`
	// Random seed for assigning users to a group.
	// +optional
	Seed int `json:"seed,omitempty"`
	// TrafficPercent defines the percentage of users routed to the B group, if not 50%.
	// +optional
	TrafficPercent int `json:"trafficPercent,omitempty"`
}
