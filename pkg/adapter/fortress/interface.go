package fortress

import (
	"time"

	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type FortressAdapter interface {
	GetIcyWeeklyDistribution() (icys *model.AdapterIcy, err error)

	GetCommunityEarn() (earns *model.AdapterEarn, err error)
	GetCommunityMemoEarn() (earns *model.AdapterMemoEarn, err error)

	GetTechRadar(ringFilter string, query *string) (techradars *model.AdapterTechRadar, err error)

	GetNewSubscribers() (subscribers *model.AdapterSubscriber, err error)

	GetOpenPositions() (positions *model.AdapterHiringPosition, err error)

	GetUpcomingEvents() (events *model.AdapterEvent, err error)
	CreateGuildScheduledEvent(ev *model.DiscordEvent) error
	GetGuildScheduledEvents() ([]*model.DiscordEvent, error)
	SetSpeakers(eventID string, mapSpeakers map[string][]string) (*model.Event, error)

	GetStaffingDemands() (staffs *model.AdapterStaffingDemands, err error)

	GetProjects(q string) (res *model.ProjectListResponse, err error)
	GetProjectMilestones(q string) (milestones *model.AdapterProjectMilestone, err error)
	GetProjectCommissionModels(projectID string) (cm []model.ProjectCommissionModel, err error)

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
	GetDeliveryMetricsMonthlyReportDiscordMsg(now bool) (msg *model.AdapterDeliveryMetricsReportMsg, err error)
	SyncDeliveryMetricsData() (err error)

	GetEmployees(in EmployeeSearch) (rs []model.Employee, err error)
	GetEmployeesWithMMAScore() (employees []model.EmployeeMMAScore, err error)

	GetTrendingRepos(spokenLang string, programLang string, dateRange string) (repos *model.AdapterTrend, err error)
	SalaryAdvance(discordID, amount string) (salaryAdvance *model.AdapterSalaryAdvance, err error)
	CheckAdvanceSalary(discordID string) (salaryAdvance *model.AdapterCheckSalaryAdvance, err error)

	SalaryAdvanceReport() (unpaidSalaryAdvances *model.AdapterSalaryAdvanceReport, err error)
	GetIcyAccounting() (icyAccounting *model.AdapterIcyAccounting, err error)
	ListICYEarnedTransactions(discordID string, page, size int) (*model.AdapterICYEarnedTransactions, error)
	GetICYTotalEarned(discordID string) (*model.AdapterICYTotalEarned, error)
	Get30daysTotalReward() (*model.AdapterICYTotalEarned, error)

	CheckWithdrawCondition(discordID string) (rs *model.AdapterCheckWithdrawCondition, err error)
	GetBanks(id, bin, swiftCode string) (banks *model.AdapterBank, err error)

	SyncMemoLogs() (memos *model.MemoLogsResponse, err error)
	GetMemoLogs(from, to *time.Time) (memos *model.MemoLogsResponse, err error)
	GetMemoOpenPullRequest() (memos *model.MemoPullRequestResponse, err error)
	GetMemoLogsByDiscordID(discordID string) (memos *model.MemoLogsByDiscordIDResponse, err error)
	GetTopAuthors(limit, days int) (topAuthors *model.MemoTopAuthorsResponse, err error)

	GetDiscordResearchTopics(timeRange string) (data *model.DiscordResearchTopicResponse, err error)

	FetchNews(platform, topic string) ([]model.News, error)
	GetOgifStats(discordID string, after time.Time) (*model.OgifStatsResponse, error)
	GetOgifLeaderboard(after time.Time, limit int) ([]model.OgifLeaderboardRecord, error)
}
