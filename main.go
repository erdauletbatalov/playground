package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	rateLimiter = sync.Map{}
	cooldown    = 10 * time.Second
)

func main() {
	for {
		time.Sleep(time.Second)

		InstantStats()

		continue
	}
}

func InstantStats() error {
	users := []string{"erdauletbatalov", "erictronic", "rshezarr"}

	// Собираем данные для топ-3
	type UserStats struct {
		FirstName      string
		Username       string
		LeetCodeUser   string
		WeeklyProblems int
		DailyStats     string
	}

	var userStatsSlice []UserStats

	for _, user := range users {
		stats, err := GetLeetCodeStats(user)
		if err != nil {
			continue
		}

		weekStats, totalWeek, _ := GetCurrentWeekStats(stats)
		userStatsSlice = append(userStatsSlice, UserStats{
			LeetCodeUser:   user,
			WeeklyProblems: totalWeek,
			DailyStats:     weekStats,
		})
	}

	// Сортируем по количеству решенных задач за неделю
	sort.Slice(userStatsSlice, func(i, j int) bool {
		return userStatsSlice[i].WeeklyProblems > userStatsSlice[j].WeeklyProblems
	})

	statsMsg := "🏆 Топ 3 на за эту неделю на:\n\n"

	// Выводим топ-3
	for i := 0; i < len(userStatsSlice) && i < 3; i++ {
		statsMsg += fmt.Sprintf("%d. %s (@%s) - %d задач за неделю ля наху\n",
			i+1,
			userStatsSlice[i].FirstName,
			userStatsSlice[i].LeetCodeUser,
			userStatsSlice[i].WeeklyProblems,
		)
	}

	statsMsg += "\n📊 Опщи статистика:\n\n"

	// Основная статистика
	for _, user := range users {
		stats, err := GetLeetCodeStats(user)
		if err != nil {
			statsMsg += fmt.Sprintf("⚠️ Не удалось получить статистику для %s\n", user)
			continue
		}

		statsMsg += fmt.Sprintf(
			"👤 %s:\n- Решено задач: %d\n- Изи: %d\n- Нормис: %d\n- Хардовые: %d\n\n",
			user,
			stats.TotalSolved,
			stats.EasySolved,
			stats.MediumSolved,
			stats.HardSolved,
		)
	}

	fmt.Println(statsMsg)
	return nil
}

func GetCurrentWeekStats(stats *LeetCodeStats) (string, int, error) {
	now := time.Now()
	weekStart := now.AddDate(0, 0, -int(now.Weekday()))

	weekDays := make(map[string]int)
	totalSolved := 0

	fmt.Println("stats.SubmissionCalendar: ", stats.SubmissionCalendar, "/n")

	// Get all days of current week
	for i := 0; i < 7; i++ {
		currentDay := weekStart.AddDate(0, 0, i)
		timestamp := strconv.FormatInt(currentDay.Unix(), 10)
		fmt.Println("current day: " + timestamp + "/n")

		if count, exists := stats.SubmissionCalendar[timestamp]; exists {
			weekDays[currentDay.Format("Monday")] = count
			totalSolved += count
		} else {
			weekDays[currentDay.Format("Monday")] = 0
		}
	}

	var result strings.Builder
	result.WriteString("Статистика текущей недели:\n\n")

	for day, count := range weekDays {
		result.WriteString(fmt.Sprintf("%s: %d задач\n", day, count))
	}

	result.WriteString(fmt.Sprintf("\nВсего за неделю: %d задач", totalSolved))

	return result.String(), totalSolved, nil
}

type LeetCodeStats struct {
	TotalSolved        int            `json:"totalSolved"`
	TotalQuestions     int            `json:"totalQuestions"`
	EasySolved         int            `json:"easySolved"`
	MediumSolved       int            `json:"mediumSolved"`
	HardSolved         int            `json:"hardSolved"`
	SubmissionCalendar map[string]int `json:"submissionCalendar"`
}

var (
	statsCache = sync.Map{}
	cacheTTL   = 10 * time.Second
)

type CachedStats struct {
	Stats    *LeetCodeStats
	ExpireAt time.Time
}

func GetLeetCodeStats(username string) (*LeetCodeStats, error) {
	// Check cache first
	if cached, ok := statsCache.Load(username); ok {
		cachedData := cached.(CachedStats)
		if time.Now().Before(cachedData.ExpireAt) {
			return cachedData.Stats, nil
		}
		statsCache.Delete(username)
	}

	// Make API request if not in cache or expired
	url := fmt.Sprintf("https://leetcode-stats-api.herokuapp.com/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch stats for %s: %s", username, resp.Status)
	}

	var stats LeetCodeStats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, err
	}

	// Store in cache
	statsCache.Store(username, CachedStats{
		Stats:    &stats,
		ExpireAt: time.Now().Add(cacheTTL),
	})

	return &stats, nil
}
