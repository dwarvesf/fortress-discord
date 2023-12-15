package fortress

import "github.com/dwarvesf/fortress-discord/pkg/model"

type FortressAdapter interface {
	GetIcyWeeklyDistribution() (icys *model.AdapterIcy, err error)

	GetCommunityEarn() (earns *model.AdapterEarn, err error)

	GetTechRadar(ringFilter string, query *string) (techradars *model.AdapterTechRadar, err error)

	GetNewSubscribers() (subscribers *model.AdapterSubscriber, err error)

	GetOpenPositions() (positions *model.AdapterHiringPosition, err error)

	GetUpcomingEvents() (events *model.AdapterEvent, err error)

	GetStaffingDemands() (staffs *model.AdapterStaffingDemands, err error)

	GetProjectMilestones(q string) (milestones *model.AdapterProjectMilestone, err error)

	GetInternalDigest() (digests *model.AdapterDigest, err error)
	GetExternalDigest() (digests *model.AdapterDigest, err error)

	GetMemos() (memos *model.AdapterMemo, err error)

	GetActiveIssues() (issues *model.AdapterIssue, err error)

	LogTechRadarTopic(topicName string, discordId string) error

	GetChangelogs() (digests *model.ChangelogDigest, err error)
	SendChangelog(changelog *model.Changelog) error

	UpsertRollupRecord(record *model.EngagementsRollupRecord) error

	CreateBraineryPost(post *model.CreateBraineryLogRequest) error
	GetBraineryReport(view string, date string) (*model.BraineryMetric, error)

	GetDeliveryMetricsWeeklyReportDiscordMsg() (msg *model.AdapterDeliveryMetricsReportMsg, err error)
	GetDeliveryMetricsMonthlyReportDiscordMsg() (msg *model.AdapterDeliveryMetricsReportMsg, err error)
	SyncDeliveryMetricsData() (err error)

	GetEmployees(in EmployeeSearch) (rs []model.Employee, err error)
	GetEmployeesWithMMAScore() (employees []model.EmployeeMMAScore, err error)

	GetTrendingRepos(spokenLang string, programLang string, dateRange string) (repos *model.AdapterTrend, err error)
	SalaryAdvance(discordID, amount string) (salaryAdvance *model.AdapterSalaryAdvance, err error)
	CheckAdvanceSalary(discordID string) (salaryAdvance *model.AdapterCheckSalaryAdvance, err error)

	SalaryAdvanceReport() (unpaidSalaryAdvances *model.AdapterSalaryAdvanceReport, err error)
	GetIcyAccounting() (icyAccounting *model.AdapterIcyAccounting, err error)
}
