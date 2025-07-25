package parser

import (
	"github.com/amirmtaati/tempus/internal/core/task"
	"regexp"
	"strings"
	"time"
)

const emptyStr = ""

type Parser struct {
	priorityRgx       *regexp.Regexp
	completedRgx      *regexp.Regexp
	creationDateRgx   *regexp.Regexp
	completionDateRgx *regexp.Regexp
	tagsRgx           *regexp.Regexp
	contextRgx        *regexp.Regexp
	projectsRgx       *regexp.Regexp
}

func NewParser() *Parser {
	return &Parser{
		priorityRgx:       regexp.MustCompile(`^(x|x \d{4}-\d{2}-\d{2}|)\s*\(([A-Z])\)\s+`),
		completedRgx:      regexp.MustCompile(`^x\s+`),
		creationDateRgx:   regexp.MustCompile(`^(\([A-Z]\)|x \d{4}-\d{2}-\d{2} \([A-Z]\)|x \([A-Z]\)|x \d{4}-\d{2}-\d{2}|)\s*(\d{4}-\d{2}-\d{2})\s+`),
		completionDateRgx: regexp.MustCompile(`^x\s*(\d{4}-\d{2}-\d{2})\s+`),
		tagsRgx:           regexp.MustCompile(`^x\s*(\d{4}-\d{2}-\d{2})\s+`),
		contextRgx:        regexp.MustCompile(`(^|\s+)@(\S+)`),
		projectsRgx:       regexp.MustCompile(`(^|\s+)\+(\S+)`),
	}
}

func (p *Parser) Parse(line string) (*task.Task, error) {
	// init
	task := task.NewTask()
	line = strings.TrimSpace(line)

	// 1. Completed?
	if p.completedRgx.MatchString(line) {
		task.Done = true
		task.Raw = p.completedRgx.ReplaceAllString(task.Raw, emptyStr)
	}

	// 2. Priority
	if p.priorityRgx.MatchString(line) {
		task.Priority = p.priorityRgx.FindStringSubmatch(line)[2]
		task.Raw = p.priorityRgx.ReplaceAllString(task.Raw, emptyStr)
	}

	// 3. Creation, Completion Date
	if p.creationDateRgx.MatchString(line) {
		date, err := time.ParseInLocation("2006-01-02", p.creationDateRgx.FindStringSubmatch(line)[2], time.Local)

		if err != nil {

		}

		task.CreationDate = date
		task.Raw = p.creationDateRgx.ReplaceAllString(task.Raw, emptyStr)
	}

	if p.completionDateRgx.MatchString(line) {
		date, err := time.ParseInLocation("2006-01-02", p.completionDateRgx.FindStringSubmatch(line)[1], time.Local)

		if err != nil {

		}

		task.CompletionDate = date
		task.Raw = p.completionDateRgx.ReplaceAllString(task.Raw, emptyStr)
	}

	// 4. Context, Projects, Tags
	if p.contextRgx.MatchString(line) {

	}

	if p.projectsRgx.MatchString(line) {

	}

	if p.tagsRgx.MatchString(line) {

	}

	return &task, nil
}
