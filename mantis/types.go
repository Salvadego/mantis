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
	OrgID             int64   `json:"orgId"`
	Quantity          float64 `json:"quantity"`
	SalesOrderLine    int64   `json:"salesOrderLine"`
	SalesOrder        int64   `json:"salesOrder"`
	ReferenceID       int64   `json:"referenceId"`
	DocumentNo        string  `json:"documentNo"`
	NonBusinessDayID  int64   `json:"nonBusinessDayId"`
	IsSurplusActivity bool    `json:"isSurplusActivity"`
	IsDuty            bool    `json:"isDuty"`
	Fase              int64   `json:"fase"`
	ExternalCode      string  `json:"externalCode"`
	DocNumWarranty    int64   `json:"docNumWarranty"`
	IsApproved        bool    `json:"isApproved"`
	IsProcessed       bool    `json:"isProcessed"`
	IsActive          bool    `json:"isActive"`
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

type ErrorDetail struct {
	Message string `json:"message"`
}

type APIErrorItem struct {
	Status      string        `json:"status"`
	Date        string        `json:"date"`
	Type        string        `json:"type"`
	Code        string        `json:"code"`
	Message     string        `json:"message"`
	MoreInfo    string        `json:"moreInfo"`
	MessageType string        `json:"messageType"`
	Details     []ErrorDetail `json:"details"`
}

type ErrorsResponse struct {
	Errors []APIErrorItem `json:"errors"`
}

func (e ErrorsResponse) IsError() bool {
	return e.Errors != nil
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

type TicketsResponse struct {
	TicketNo      string `json:"ticketNo"`
	Description   string `json:"description"`
	ContractID    string `json:"contractId"`
	PCApproved    string `json:"pcApproved"`
	KurzText      string `json:"kurzText"`
	ContractTitle string `json:"contractTitle"`
	ContractType  string `json:"contractType"`
}

type ContractResponse struct {
	LtContracts []LtContract `json:"Lt_Contracts"`
}

type LtContract struct {
	ContractID   string       `json:"Contract_ID"`
	Title        string       `json:"Title"`
	ContractType ContractType `json:"Contract_Type"`
}

type ContractType string

const (
	ServiceDesk ContractType = "SERVICE DESK"
)

type S_Employee struct {
	AdUserID     int64      `json:"Ad_User_ID"`
	Name         string     `json:"Name"`
	SUserID      string     `json:"S_User_ID"`
	TipoLider    string     `json:"Tipo_Lider"`
	Active       bool       `json:"Active"`
	Module       string     `json:"Module"`
	EmployeeCode string     `json:"Employee_Code"`
	Bp           string     `json:"BP"`
	Unit         string     `json:"Unit"`
	Terceiro     bool       `json:"Terceiro"`
	FullTime     bool       `json:"Full_Time"`
	DataInicio   *time.Time `json:"Data_Inicio"`
	DataFim      *time.Time `json:"Data_Fim"`
	CUser        string     `json:"C_User"`
}

type SupportInfoResponse struct {
	Description            string         `json:"Description"`
	Priority               string         `json:"Priority"`
	Category               any            `json:"Category"`
	CategoryDescription    any            `json:"Category_Description"`
	Text                   any            `json:"Text"`
	ProcessType            string         `json:"Process_Type"`
	Reference              any            `json:"Reference"`
	ObjectID               string         `json:"Object_ID"`
	GUID                   string         `json:"Guid"`
	UserStatus             string         `json:"User_Status"`
	Attachments            []Attachment   `json:"Attachments"`
	File                   []any          `json:"File"`
	CreatedAt              time.Time      `json:"Created_At"`
	CreatedBy              CreatedBy      `json:"Created_By"`
	ChangedBy              any            `json:"Changed_By"`
	ChangedAt              time.Time      `json:"Changed_At"`
	CategoryID             string         `json:"Category_ID"`
	ContractID             any            `json:"Contract_ID"`
	AttachmentNo           any            `json:"Attachment_No"`
	Sla                    any            `json:"Sla"`
	DateReport             any            `json:"Date_Report"`
	Processor              any            `json:"Processor"`
	PriorityDescription    string         `json:"Priority_Description"`
	UserStatusDescription  string         `json:"User_Status_Description"`
	ProcessTypeDescription string         `json:"Process_Type_Description"`
	Texts                  []Text         `json:"Texts"`
	Processadores          []Processadore `json:"Processadores"`
	SubstatusDescription   string         `json:"Substatus_Description"`
	DateEndPlan            time.Time      `json:"Date_End_Plan"`
	DateStartPlan          time.Time      `json:"Date_Start_Plan"`
	CreatedAtFrom          any            `json:"Created_At_From"`
	CreatedAtTo            any            `json:"Created_At_To"`
	ChangedAtFrom          any            `json:"Changed_At_From"`
	ChangedAtTo            any            `json:"Changed_At_To"`
	VisibleUpdate          bool           `json:"Visible_Update"`
	VisibleRef             bool           `json:"Visible_Ref"`
	VisibleApprove         bool           `json:"Visible_Approve"`
	VisibleDisapprove      bool           `json:"Visible_Disapprove"`
	VisibleClose           bool           `json:"Visible_Close"`
	EnableUpdate           bool           `json:"Enable_Update"`
	EnableRef              bool           `json:"Enable_Ref"`
	EnableApprove          bool           `json:"Enable_Approve"`
	EnableDisapprove       bool           `json:"Enable_Disapprove"`
	EnableClose            bool           `json:"Enable_Close"`
	DocumentFlows          []any          `json:"Document_Flows"`
	ProcessorDetail        CreatedBy      `json:"Processor_Detail"`
	EvIsPC                 string         `json:"Ev_Is_Pc"`
	EvPCStatDesc           string         `json:"Ev_Pc_Stat_Desc"`
	FilterProcessType      any            `json:"Filter_Process_Type"`
	TotHrChamado           string         `json:"Tot_Hr_Chamado"`
	TotHrPC                string         `json:"Tot_Hr_PC"`
	Baseline               string         `json:"Baseline"`
	TotHrAprovadaPC        string         `json:"Tot_Hr_Aprovada_PC"`
	TicketNumber           any            `json:"Ticket_Number"`
}

type Attachment struct {
	GUID        string    `json:"Guid"`
	FileContent string    `json:"File_Content"`
	FileName    string    `json:"File_Name"`
	CreatedBy   CreatedBy `json:"Created_By"`
	CreatedAt   time.Time `json:"Created_At"`
}

type CreatedBy struct {
	ADUserID         int64  `json:"AD_User_ID"`
	Name             string `json:"Name"`
	Description      string `json:"Description"`
	CellPhone        any    `json:"CellPhone"`
	Phone            string `json:"Phone"`
	Email            string `json:"Email"`
	UserOtherService string `json:"User_Other_Service"`
	PictureID        int64  `json:"Picture_ID"`
	BusinessUnitID   *int64 `json:"BusinessUnit_ID"`
}

type Processadore struct {
	UserInformation CreatedBy `json:"User_Information"`
}

type Text struct {
	UserInformation *CreatedBy `json:"User_Information"`
	Text            string     `json:"Text"`
	TDFCreatedAt    time.Time  `json:"TDFCreated_At"`
	TDFUser         string     `json:"TDFUser"`
	GUID            string     `json:"Guid"`
	TdID            string     `json:"TD_ID"`
	Description     string     `json:"Description"`
}
