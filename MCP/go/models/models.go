package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ServerSideEncryptionConfiguration represents the ServerSideEncryptionConfiguration schema from the OpenAPI specification
type ServerSideEncryptionConfiguration struct {
	Kmskeyid interface{} `json:"KmsKeyId"`
}

// CreateDomainRequest represents the CreateDomainRequest schema from the OpenAPI specification
type CreateDomainRequest struct {
	Name interface{} `json:"Name"`
	Serversideencryptionconfiguration interface{} `json:"ServerSideEncryptionConfiguration"`
	Tags interface{} `json:"Tags,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// ListFraudsterRegistrationJobsRequest represents the ListFraudsterRegistrationJobsRequest schema from the OpenAPI specification
type ListFraudsterRegistrationJobsRequest struct {
	Domainid interface{} `json:"DomainId"`
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DomainSummary represents the DomainSummary schema from the OpenAPI specification
type DomainSummary struct {
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Serversideencryptionconfiguration interface{} `json:"ServerSideEncryptionConfiguration,omitempty"`
	Serversideencryptionupdatedetails interface{} `json:"ServerSideEncryptionUpdateDetails,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Domainstatus interface{} `json:"DomainStatus,omitempty"`
	Watchlistdetails interface{} `json:"WatchlistDetails,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
}

// DescribeFraudsterResponse represents the DescribeFraudsterResponse schema from the OpenAPI specification
type DescribeFraudsterResponse struct {
	Fraudster interface{} `json:"Fraudster,omitempty"`
}

// Domain represents the Domain schema from the OpenAPI specification
type Domain struct {
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
	Serversideencryptionupdatedetails interface{} `json:"ServerSideEncryptionUpdateDetails,omitempty"`
	Watchlistdetails interface{} `json:"WatchlistDetails,omitempty"`
	Serversideencryptionconfiguration interface{} `json:"ServerSideEncryptionConfiguration,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Domainstatus interface{} `json:"DomainStatus,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
}

// DisassociateFraudsterResponse represents the DisassociateFraudsterResponse schema from the OpenAPI specification
type DisassociateFraudsterResponse struct {
	Fraudster Fraudster `json:"Fraudster,omitempty"` // Contains all the information about a fraudster.
}

// DescribeFraudsterRequest represents the DescribeFraudsterRequest schema from the OpenAPI specification
type DescribeFraudsterRequest struct {
	Domainid interface{} `json:"DomainId"`
	Fraudsterid interface{} `json:"FraudsterId"`
}

// FraudsterRegistrationJob represents the FraudsterRegistrationJob schema from the OpenAPI specification
type FraudsterRegistrationJob struct {
	Domainid interface{} `json:"DomainId,omitempty"`
	Endedat interface{} `json:"EndedAt,omitempty"`
	Jobname interface{} `json:"JobName,omitempty"`
	Jobprogress interface{} `json:"JobProgress,omitempty"`
	Outputdataconfig interface{} `json:"OutputDataConfig,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Registrationconfig interface{} `json:"RegistrationConfig,omitempty"`
	Failuredetails interface{} `json:"FailureDetails,omitempty"`
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Inputdataconfig interface{} `json:"InputDataConfig,omitempty"`
	Dataaccessrolearn interface{} `json:"DataAccessRoleArn,omitempty"`
}

// ListFraudstersResponse represents the ListFraudstersResponse schema from the OpenAPI specification
type ListFraudstersResponse struct {
	Fraudstersummaries interface{} `json:"FraudsterSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// SpeakerEnrollmentJobSummary represents the SpeakerEnrollmentJobSummary schema from the OpenAPI specification
type SpeakerEnrollmentJobSummary struct {
	Endedat interface{} `json:"EndedAt,omitempty"`
	Failuredetails interface{} `json:"FailureDetails,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Jobname interface{} `json:"JobName,omitempty"`
	Jobprogress interface{} `json:"JobProgress,omitempty"`
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
}

// SpeakerSummary represents the SpeakerSummary schema from the OpenAPI specification
type SpeakerSummary struct {
	Lastaccessedat interface{} `json:"LastAccessedAt,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Customerspeakerid interface{} `json:"CustomerSpeakerId,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
	Generatedspeakerid interface{} `json:"GeneratedSpeakerId,omitempty"`
}

// DeleteFraudsterRequest represents the DeleteFraudsterRequest schema from the OpenAPI specification
type DeleteFraudsterRequest struct {
	Fraudsterid interface{} `json:"FraudsterId"`
	Domainid interface{} `json:"DomainId"`
}

// Watchlist represents the Watchlist schema from the OpenAPI specification
type Watchlist struct {
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Watchlistid interface{} `json:"WatchlistId,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Defaultwatchlist interface{} `json:"DefaultWatchlist,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// DescribeSpeakerRequest represents the DescribeSpeakerRequest schema from the OpenAPI specification
type DescribeSpeakerRequest struct {
	Domainid interface{} `json:"DomainId"`
	Speakerid interface{} `json:"SpeakerId"`
}

// InputDataConfig represents the InputDataConfig schema from the OpenAPI specification
type InputDataConfig struct {
	S3uri interface{} `json:"S3Uri"`
}

// SpeakerEnrollmentJob represents the SpeakerEnrollmentJob schema from the OpenAPI specification
type SpeakerEnrollmentJob struct {
	Inputdataconfig interface{} `json:"InputDataConfig,omitempty"`
	Enrollmentconfig interface{} `json:"EnrollmentConfig,omitempty"`
	Failuredetails interface{} `json:"FailureDetails,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Endedat interface{} `json:"EndedAt,omitempty"`
	Jobprogress interface{} `json:"JobProgress,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
	Jobname interface{} `json:"JobName,omitempty"`
	Outputdataconfig interface{} `json:"OutputDataConfig,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Dataaccessrolearn interface{} `json:"DataAccessRoleArn,omitempty"`
}

// DescribeFraudsterRegistrationJobResponse represents the DescribeFraudsterRegistrationJobResponse schema from the OpenAPI specification
type DescribeFraudsterRegistrationJobResponse struct {
	Job interface{} `json:"Job,omitempty"`
}

// EnrollmentConfig represents the EnrollmentConfig schema from the OpenAPI specification
type EnrollmentConfig struct {
	Existingenrollmentaction interface{} `json:"ExistingEnrollmentAction,omitempty"`
	Frauddetectionconfig interface{} `json:"FraudDetectionConfig,omitempty"`
}

// ListWatchlistsRequest represents the ListWatchlistsRequest schema from the OpenAPI specification
type ListWatchlistsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Domainid interface{} `json:"DomainId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// FraudRiskDetails represents the FraudRiskDetails schema from the OpenAPI specification
type FraudRiskDetails struct {
	Knownfraudsterrisk interface{} `json:"KnownFraudsterRisk"`
	Voicespoofingrisk interface{} `json:"VoiceSpoofingRisk"`
}

// ListDomainsResponse represents the ListDomainsResponse schema from the OpenAPI specification
type ListDomainsResponse struct {
	Domainsummaries interface{} `json:"DomainSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListSpeakersRequest represents the ListSpeakersRequest schema from the OpenAPI specification
type ListSpeakersRequest struct {
	Domainid interface{} `json:"DomainId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeDomainResponse represents the DescribeDomainResponse schema from the OpenAPI specification
type DescribeDomainResponse struct {
	Domain interface{} `json:"Domain,omitempty"`
}

// DescribeSpeakerEnrollmentJobRequest represents the DescribeSpeakerEnrollmentJobRequest schema from the OpenAPI specification
type DescribeSpeakerEnrollmentJobRequest struct {
	Jobid interface{} `json:"JobId"`
	Domainid interface{} `json:"DomainId"`
}

// EnrollmentJobFraudDetectionConfig represents the EnrollmentJobFraudDetectionConfig schema from the OpenAPI specification
type EnrollmentJobFraudDetectionConfig struct {
	Frauddetectionaction interface{} `json:"FraudDetectionAction,omitempty"`
	Riskthreshold interface{} `json:"RiskThreshold,omitempty"`
	Watchlistids interface{} `json:"WatchlistIds,omitempty"`
}

// DeleteWatchlistRequest represents the DeleteWatchlistRequest schema from the OpenAPI specification
type DeleteWatchlistRequest struct {
	Domainid interface{} `json:"DomainId"`
	Watchlistid interface{} `json:"WatchlistId"`
}

// OptOutSpeakerResponse represents the OptOutSpeakerResponse schema from the OpenAPI specification
type OptOutSpeakerResponse struct {
	Speaker interface{} `json:"Speaker,omitempty"`
}

// CreateDomainResponse represents the CreateDomainResponse schema from the OpenAPI specification
type CreateDomainResponse struct {
	Domain interface{} `json:"Domain,omitempty"`
}

// ListFraudstersRequest represents the ListFraudstersRequest schema from the OpenAPI specification
type ListFraudstersRequest struct {
	Watchlistid interface{} `json:"WatchlistId,omitempty"`
	Domainid interface{} `json:"DomainId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Fraudster represents the Fraudster schema from the OpenAPI specification
type Fraudster struct {
	Domainid interface{} `json:"DomainId,omitempty"`
	Generatedfraudsterid interface{} `json:"GeneratedFraudsterId,omitempty"`
	Watchlistids interface{} `json:"WatchlistIds,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// FraudsterRegistrationJobSummary represents the FraudsterRegistrationJobSummary schema from the OpenAPI specification
type FraudsterRegistrationJobSummary struct {
	Domainid interface{} `json:"DomainId,omitempty"`
	Endedat interface{} `json:"EndedAt,omitempty"`
	Failuredetails interface{} `json:"FailureDetails,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Jobname interface{} `json:"JobName,omitempty"`
	Jobprogress interface{} `json:"JobProgress,omitempty"`
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Tagkeys interface{} `json:"TagKeys"`
	Resourcearn interface{} `json:"ResourceArn"`
}

// UpdateWatchlistRequest represents the UpdateWatchlistRequest schema from the OpenAPI specification
type UpdateWatchlistRequest struct {
	Name interface{} `json:"Name,omitempty"`
	Watchlistid interface{} `json:"WatchlistId"`
	Description interface{} `json:"Description,omitempty"`
	Domainid interface{} `json:"DomainId"`
}

// StartSpeakerEnrollmentJobResponse represents the StartSpeakerEnrollmentJobResponse schema from the OpenAPI specification
type StartSpeakerEnrollmentJobResponse struct {
	Job interface{} `json:"Job,omitempty"`
}

// RegistrationConfig represents the RegistrationConfig schema from the OpenAPI specification
type RegistrationConfig struct {
	Duplicateregistrationaction interface{} `json:"DuplicateRegistrationAction,omitempty"`
	Fraudstersimilaritythreshold interface{} `json:"FraudsterSimilarityThreshold,omitempty"`
	Watchlistids interface{} `json:"WatchlistIds,omitempty"`
}

// AssociateFraudsterRequest represents the AssociateFraudsterRequest schema from the OpenAPI specification
type AssociateFraudsterRequest struct {
	Domainid interface{} `json:"DomainId"`
	Fraudsterid interface{} `json:"FraudsterId"`
	Watchlistid interface{} `json:"WatchlistId"`
}

// DeleteSpeakerRequest represents the DeleteSpeakerRequest schema from the OpenAPI specification
type DeleteSpeakerRequest struct {
	Domainid interface{} `json:"DomainId"`
	Speakerid interface{} `json:"SpeakerId"`
}

// ListSpeakerEnrollmentJobsResponse represents the ListSpeakerEnrollmentJobsResponse schema from the OpenAPI specification
type ListSpeakerEnrollmentJobsResponse struct {
	Jobsummaries interface{} `json:"JobSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateDomainRequest represents the UpdateDomainRequest schema from the OpenAPI specification
type UpdateDomainRequest struct {
	Domainid interface{} `json:"DomainId"`
	Name interface{} `json:"Name"`
	Serversideencryptionconfiguration interface{} `json:"ServerSideEncryptionConfiguration"`
	Description interface{} `json:"Description,omitempty"`
}

// AuthenticationResult represents the AuthenticationResult schema from the OpenAPI specification
type AuthenticationResult struct {
	Authenticationresultid interface{} `json:"AuthenticationResultId,omitempty"`
	Configuration interface{} `json:"Configuration,omitempty"`
	Customerspeakerid interface{} `json:"CustomerSpeakerId,omitempty"`
	Decision interface{} `json:"Decision,omitempty"`
	Generatedspeakerid interface{} `json:"GeneratedSpeakerId,omitempty"`
	Score interface{} `json:"Score,omitempty"`
	Audioaggregationendedat interface{} `json:"AudioAggregationEndedAt,omitempty"`
	Audioaggregationstartedat interface{} `json:"AudioAggregationStartedAt,omitempty"`
}

// Speaker represents the Speaker schema from the OpenAPI specification
type Speaker struct {
	Status interface{} `json:"Status,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Customerspeakerid interface{} `json:"CustomerSpeakerId,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
	Generatedspeakerid interface{} `json:"GeneratedSpeakerId,omitempty"`
	Lastaccessedat interface{} `json:"LastAccessedAt,omitempty"`
}

// ListSpeakerEnrollmentJobsRequest represents the ListSpeakerEnrollmentJobsRequest schema from the OpenAPI specification
type ListSpeakerEnrollmentJobsRequest struct {
	Domainid interface{} `json:"DomainId"`
	Jobstatus interface{} `json:"JobStatus,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// EvaluateSessionResponse represents the EvaluateSessionResponse schema from the OpenAPI specification
type EvaluateSessionResponse struct {
	Frauddetectionresult interface{} `json:"FraudDetectionResult,omitempty"`
	Sessionid interface{} `json:"SessionId,omitempty"`
	Sessionname interface{} `json:"SessionName,omitempty"`
	Streamingstatus interface{} `json:"StreamingStatus,omitempty"`
	Authenticationresult interface{} `json:"AuthenticationResult,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
}

// JobProgress represents the JobProgress schema from the OpenAPI specification
type JobProgress struct {
	Percentcomplete interface{} `json:"PercentComplete,omitempty"`
}

// WatchlistDetails represents the WatchlistDetails schema from the OpenAPI specification
type WatchlistDetails struct {
	Defaultwatchlistid interface{} `json:"DefaultWatchlistId"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Tags interface{} `json:"Tags"`
}

// DescribeSpeakerEnrollmentJobResponse represents the DescribeSpeakerEnrollmentJobResponse schema from the OpenAPI specification
type DescribeSpeakerEnrollmentJobResponse struct {
	Job interface{} `json:"Job,omitempty"`
}

// DescribeWatchlistResponse represents the DescribeWatchlistResponse schema from the OpenAPI specification
type DescribeWatchlistResponse struct {
	Watchlist interface{} `json:"Watchlist,omitempty"`
}

// StartSpeakerEnrollmentJobRequest represents the StartSpeakerEnrollmentJobRequest schema from the OpenAPI specification
type StartSpeakerEnrollmentJobRequest struct {
	Dataaccessrolearn interface{} `json:"DataAccessRoleArn"`
	Domainid interface{} `json:"DomainId"`
	Enrollmentconfig interface{} `json:"EnrollmentConfig,omitempty"`
	Inputdataconfig interface{} `json:"InputDataConfig"`
	Jobname interface{} `json:"JobName,omitempty"`
	Outputdataconfig interface{} `json:"OutputDataConfig"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
}

// ListWatchlistsResponse represents the ListWatchlistsResponse schema from the OpenAPI specification
type ListWatchlistsResponse struct {
	Watchlistsummaries interface{} `json:"WatchlistSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// WatchlistSummary represents the WatchlistSummary schema from the OpenAPI specification
type WatchlistSummary struct {
	Domainid interface{} `json:"DomainId,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Watchlistid interface{} `json:"WatchlistId,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Defaultwatchlist interface{} `json:"DefaultWatchlist,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// UpdateWatchlistResponse represents the UpdateWatchlistResponse schema from the OpenAPI specification
type UpdateWatchlistResponse struct {
	Watchlist interface{} `json:"Watchlist,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// DeleteDomainRequest represents the DeleteDomainRequest schema from the OpenAPI specification
type DeleteDomainRequest struct {
	Domainid interface{} `json:"DomainId"`
}

// FraudsterSummary represents the FraudsterSummary schema from the OpenAPI specification
type FraudsterSummary struct {
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
	Generatedfraudsterid interface{} `json:"GeneratedFraudsterId,omitempty"`
	Watchlistids interface{} `json:"WatchlistIds,omitempty"`
}

// FailureDetails represents the FailureDetails schema from the OpenAPI specification
type FailureDetails struct {
	Message interface{} `json:"Message,omitempty"`
	Statuscode interface{} `json:"StatusCode,omitempty"`
}

// DisassociateFraudsterRequest represents the DisassociateFraudsterRequest schema from the OpenAPI specification
type DisassociateFraudsterRequest struct {
	Domainid interface{} `json:"DomainId"`
	Fraudsterid interface{} `json:"FraudsterId"`
	Watchlistid interface{} `json:"WatchlistId"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// VoiceSpoofingRisk represents the VoiceSpoofingRisk schema from the OpenAPI specification
type VoiceSpoofingRisk struct {
	Riskscore interface{} `json:"RiskScore"`
}

// KnownFraudsterRisk represents the KnownFraudsterRisk schema from the OpenAPI specification
type KnownFraudsterRisk struct {
	Riskscore interface{} `json:"RiskScore"`
	Generatedfraudsterid interface{} `json:"GeneratedFraudsterId,omitempty"`
}

// DescribeFraudsterRegistrationJobRequest represents the DescribeFraudsterRegistrationJobRequest schema from the OpenAPI specification
type DescribeFraudsterRegistrationJobRequest struct {
	Jobid interface{} `json:"JobId"`
	Domainid interface{} `json:"DomainId"`
}

// ListSpeakersResponse represents the ListSpeakersResponse schema from the OpenAPI specification
type ListSpeakersResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Speakersummaries interface{} `json:"SpeakerSummaries,omitempty"`
}

// EvaluateSessionRequest represents the EvaluateSessionRequest schema from the OpenAPI specification
type EvaluateSessionRequest struct {
	Domainid interface{} `json:"DomainId"`
	Sessionnameorid interface{} `json:"SessionNameOrId"`
}

// FraudDetectionConfiguration represents the FraudDetectionConfiguration schema from the OpenAPI specification
type FraudDetectionConfiguration struct {
	Riskthreshold interface{} `json:"RiskThreshold,omitempty"`
	Watchlistid interface{} `json:"WatchlistId,omitempty"`
}

// OptOutSpeakerRequest represents the OptOutSpeakerRequest schema from the OpenAPI specification
type OptOutSpeakerRequest struct {
	Speakerid interface{} `json:"SpeakerId"`
	Domainid interface{} `json:"DomainId"`
}

// ListFraudsterRegistrationJobsResponse represents the ListFraudsterRegistrationJobsResponse schema from the OpenAPI specification
type ListFraudsterRegistrationJobsResponse struct {
	Jobsummaries interface{} `json:"JobSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListDomainsRequest represents the ListDomainsRequest schema from the OpenAPI specification
type ListDomainsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// OutputDataConfig represents the OutputDataConfig schema from the OpenAPI specification
type OutputDataConfig struct {
	S3uri interface{} `json:"S3Uri"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// FraudDetectionResult represents the FraudDetectionResult schema from the OpenAPI specification
type FraudDetectionResult struct {
	Audioaggregationendedat interface{} `json:"AudioAggregationEndedAt,omitempty"`
	Audioaggregationstartedat interface{} `json:"AudioAggregationStartedAt,omitempty"`
	Configuration interface{} `json:"Configuration,omitempty"`
	Decision interface{} `json:"Decision,omitempty"`
	Frauddetectionresultid interface{} `json:"FraudDetectionResultId,omitempty"`
	Reasons interface{} `json:"Reasons,omitempty"`
	Riskdetails interface{} `json:"RiskDetails,omitempty"`
}

// StartFraudsterRegistrationJobResponse represents the StartFraudsterRegistrationJobResponse schema from the OpenAPI specification
type StartFraudsterRegistrationJobResponse struct {
	Job interface{} `json:"Job,omitempty"`
}

// CreateWatchlistRequest represents the CreateWatchlistRequest schema from the OpenAPI specification
type CreateWatchlistRequest struct {
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Domainid interface{} `json:"DomainId"`
	Name interface{} `json:"Name"`
}

// DescribeWatchlistRequest represents the DescribeWatchlistRequest schema from the OpenAPI specification
type DescribeWatchlistRequest struct {
	Domainid interface{} `json:"DomainId"`
	Watchlistid interface{} `json:"WatchlistId"`
}

// AuthenticationConfiguration represents the AuthenticationConfiguration schema from the OpenAPI specification
type AuthenticationConfiguration struct {
	Acceptancethreshold interface{} `json:"AcceptanceThreshold"`
}

// AssociateFraudsterResponse represents the AssociateFraudsterResponse schema from the OpenAPI specification
type AssociateFraudsterResponse struct {
	Fraudster Fraudster `json:"Fraudster,omitempty"` // Contains all the information about a fraudster.
}

// StartFraudsterRegistrationJobRequest represents the StartFraudsterRegistrationJobRequest schema from the OpenAPI specification
type StartFraudsterRegistrationJobRequest struct {
	Domainid interface{} `json:"DomainId"`
	Inputdataconfig interface{} `json:"InputDataConfig"`
	Jobname interface{} `json:"JobName,omitempty"`
	Outputdataconfig interface{} `json:"OutputDataConfig"`
	Registrationconfig interface{} `json:"RegistrationConfig,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Dataaccessrolearn interface{} `json:"DataAccessRoleArn"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// ServerSideEncryptionUpdateDetails represents the ServerSideEncryptionUpdateDetails schema from the OpenAPI specification
type ServerSideEncryptionUpdateDetails struct {
	Message interface{} `json:"Message,omitempty"`
	Oldkmskeyid interface{} `json:"OldKmsKeyId,omitempty"`
	Updatestatus interface{} `json:"UpdateStatus,omitempty"`
}

// UpdateDomainResponse represents the UpdateDomainResponse schema from the OpenAPI specification
type UpdateDomainResponse struct {
	Domain interface{} `json:"Domain,omitempty"`
}

// DescribeSpeakerResponse represents the DescribeSpeakerResponse schema from the OpenAPI specification
type DescribeSpeakerResponse struct {
	Speaker interface{} `json:"Speaker,omitempty"`
}

// CreateWatchlistResponse represents the CreateWatchlistResponse schema from the OpenAPI specification
type CreateWatchlistResponse struct {
	Watchlist interface{} `json:"Watchlist,omitempty"`
}

// DescribeDomainRequest represents the DescribeDomainRequest schema from the OpenAPI specification
type DescribeDomainRequest struct {
	Domainid interface{} `json:"DomainId"`
}
