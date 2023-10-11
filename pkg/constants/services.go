package constants

const (
	Analytic     = "analytic"
	Auth         = "auth"
	Contact      = "contact"
	Content      = "content"
	Coupon       = "coupon"
	Gateway      = "gateway"
	Message      = "message"
	Notification = "notification"
	Organization = "organization"
	Payment      = "payment"
	Person       = "person"
	Plans        = "plans"
	Tutorial     = "tutorial"
	Utility      = "utility"
	UtilityType  = "utility_type"
	Hermes       = "hermes"

	// ─── Publisher ───────────────────────────────────────────────────────
	publisher             = "publisher"
	AnalyticPublisher     = Analytic + "." + publisher
	PersonPublisher       = Person + "." + publisher
	NotificationPublisher = Notification + "." + publisher
	UtilityPublisher      = "utility" + "." + publisher
	SchedulerPublisher    = "scheduler" + "." + publisher

	JobQueue = "job-queue"
)
