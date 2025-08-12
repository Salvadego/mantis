package mantis

import "time"

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	UserId       string `json:"user_id"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type NonBusinessDay struct {
	NonBusinessDayID int64     `json:"nonBusinessDayId"`
	OrganizationID   int64     `json:"organizationId"`
	CountryID        int64     `json:"countryId"`
	Date             time.Time `json:"date"`
	TotalHour        int64     `json:"totalHour"`
	Name             string    `json:"name"`
}

type UserRole struct {
	ADRoleID int64  `json:"AD_Role_ID"`
	Name     string `json:"Name"`
}

type Timesheet struct {
	UserID              int     `json:"userId"`
	DateDoc             string  `json:"dateDoc"`
	SalesOrderLine      int     `json:"salesOrderLine"`
	TicketNo            string  `json:"ticketNo"`
	TicketContractTitle string  `json:"ticketContractTitle"`
	Fase                *string `json:"fase,omitempty"`
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
		Details []struct {
			Message string `json:"message"`
		}
	} `json:"errors"`
}

type Employee struct {
	FullName     string  `json:"fullName"`
	Email        string  `json:"email"`
	UserID       int     `json:"userId"`
	EmployeeCode int     `json:"employeeCode"`
	DailyJourney float64 `json:"dailyJourney"`
}

type TicketResponse struct {
	TicketNumber           string  `json:"objectID"`
	ProcessType            string  `json:"processType"`
	Status                 string  `json:"ZDesc_Stat"`
	Description            string  `json:"description"`
	Correlation            string  `json:"correlation"`
	Notificador            string  `json:"notificador"`
	Processador            string  `json:"processador"`
	TicketCreated          string  `json:"ticketCreated"`
	TicketConfirmationDate string  `json:"ticketConfirmationDate"`
	ZhrStatE0001           float64 `json:"zhrStatE0001"`
	ZhrStatE0002           float64 `json:"zhrStatE0002"`
	ZhrStatE0003           float64 `json:"zhrStatE0003"`
	ZhrStatE0005           float64 `json:"zhrStatE0005"`
	Priority               string  `json:"priority"`
	PercSLA                string  `json:"percSLA"`
	Catergory              string  `json:"catDesc"`
	Cat01                  string  `json:"cat01"`
	Cat02                  string  `json:"cat02"`
	Cat03                  string  `json:"cat03"`
	Cat04                  string  `json:"cat04"`
	Cat05                  string  `json:"cat05"`
	Cat06                  string  `json:"cat06"`
	Cat07                  string  `json:"cat07"`
	SoldToPartLis          string  `json:"soldToPartLis"`
	Comp                   string  `json:"comp"`
	TicketCustomer         string  `json:"ticketCustomer"`
	HasDoc                 string  `json:"hasDoc"`
	ProcessorName1         string  `json:"processorName1"`
	ProcessorName2         string  `json:"processorName2"`
	ProcessorName3         string  `json:"processorName3"`
	ProcessorName4         string  `json:"processorName4"`
	ProcessorName5         string  `json:"processorName5"`
	ProcessorID1           string  `json:"processorID1"`
	ProcessorID2           string  `json:"processorID2"`
	ProcessorID3           string  `json:"processorID3"`
	ProcessorID4           string  `json:"processorID4"`
	ProcessorID5           string  `json:"processorID5"`
	ConfirmAt              any     `json:"confirmAt"`
	ZhrLiq                 string  `json:"zhrLiq"`
	ZhrLiqNet              string  `json:"zhrLiqNet"`
	ZhrTot                 string  `json:"zhrTot"`
	ZTotalDias             string  `json:"zTotalDias"`
	HrsConvertidas         string  `json:"hrsConvertidas"`
	Mes                    string  `json:"mes"`
	ExpSLA                 string  `json:"expSLA"`
	ExpMotivo              string  `json:"expMotivo"`
	Destaque               string  `json:"destaque"`
	DestaqueMotivo         string  `json:"destaqueMotivo"`
	IDContrato             string  `json:"idContrato"`
	TituloContrato         string  `json:"tituloContrato"`
	TipoContrato           string  `json:"tipoContrato"`
	NameDM                 string  `json:"nameDM"`
	TicketCreatedHr        string  `json:"ticketCreatedHr"`
	TicketConfirmationHr   string  `json:"ticketConfirmationHr"`
	ConfirmHr              any     `json:"confirmHr"`
	Category               string  `json:"category"`
	SLAGoal                any     `json:"slaGoal"`
}

type ReferenceType struct {
	ColumnName string `json:"ColumnName"`
	TableName  string `json:"TableName"`
	Value      string `json:"value"`
	Name       string `json:"name"`
	IsDistinct bool   `json:"Is_Distinct"`
}
