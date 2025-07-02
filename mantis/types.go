package mantis

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	UserId       string `json:"user_id"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type Timesheet struct {
	UserID              int     `json:"userId"`
	DateDoc             string  `json:"dateDoc"`
	SalesOrderLine      int     `json:"salesOrderLine"`
	TicketNo            string  `json:"ticketNo"`
	TicketContractTitle string  `json:"ticketContractTitle"`
	Fase                *string `json:"fase"`
	Quantity            float64 `json:"quantity"`
	TicketDescription   string  `json:"ticketDescription"`
	Description         string  `json:"description"`
	TimesheetType       string  `json:"timesheetType"`
	SalesOrder          int     `json:"salesOrder"`
}

type TimesheetResponse struct {
	Value []Timesheet `json:"value"`
}

type TimesheetsResponse struct {
	TimesheetID       int     `json:"timesheetId"`
	TimesheetType     string  `json:"timesheetType"`
	UserID            int     `json:"userId"`
	Quantity          float64 `json:"quantity"`
	TicketNo          string  `json:"ticketNo"`
	Description       string  `json:"description"`
	DateDoc           string  `json:"dateDoc"`
	ProjectName       string  `json:"projectName"`
	ProjectManager    string  `json:"projectManager"`
	TicketDescription string  `json:"ticketDescription"`
	TicketContract    string  `json:"ticketContractTitle"`
	Cost              float64 `json:"cost"`
}

type ProjectTimesheet struct {
	ProjectNumber      int    `json:"projectNumber"`
	ProjectTitle       string `json:"projectTitle"`
	EmployeeLineNumber int    `json:"employeeLineNumber"`
	ProjectNeedTicket  bool   `json:"projectNeedTicket"`
}

type ErrorsResponse struct {
	Errors []struct {
		Status  string `json:"status"`
		Date    string `json:"date"`
		Message string `json:"message"`
	} `json:"errors"`
}

type Employee struct {
	FullName     string  `json:"fullName"`
	Email        string  `json:"email"`
	UserID       int     `json:"userId"`
	EmployeeCode int     `json:"employeeCode"`
	DailyJourney float64 `json:"dailyJourney"`
}
