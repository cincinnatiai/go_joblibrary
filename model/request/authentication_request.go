package request

type AuthenticationRequest struct {
	Method              string `json:"method"`
	AttemptingServiceId string `json:"attempting_service_id"`
	GoalServiceKey      string `json:"goal_service_key"`
}
