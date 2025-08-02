package parser

import (
	"regexp"
	"strings"
	"time"

	"github.com/amirmtaati/task/internal/models"
)

const emptyStr = ""

type PatternHandler func(matches []string, task *models.Task, line string) (string, error)

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
		tagsRgx:           regexp.MustCompile(`(\w+):(\S+)`),
		contextRgx:        regexp.MustCompile(`(^|\s+)@(\S+)`),
		projectsRgx:       regexp.MustCompile(`(^|\s+)\+(\S+)`),
	}
}

func (p *Parser) PopulateTask(task *models.Task, line string) error {
	line = strings.TrimSpace(line)

	patterns := p.getPatterns()

	// Process each pattern
	for regex, handler := range patterns {
		var err error
		line, err = p.processPattern(regex, line, task, handler)
		if err != nil {
			return err
		}
	}

	task.Todo = strings.TrimSpace(line)
	return nil
}

func (p *Parser) Parse(line string) (*models.Task, error) {
	task := models.NewTask(line)
	err := p.PopulateTask(task, line)
	return task, err
}

func (p *Parser) processPattern(regex *regexp.Regexp, line string, task *models.Task, handler PatternHandler) (string, error) {
	if !regex.MatchString(line) {
		return line, nil
	}

	matches := regex.FindStringSubmatch(line)
	if len(matches) == 0 {
		return line, nil
	}

	return handler(matches, task, line)
}

func (p *Parser) handleCompleted(matches []string, task *models.Task, line string) (string, error) {
	task.Done = true
	return p.completedRgx.ReplaceAllString(line, emptyStr), nil
}

func (p *Parser) handlePriority(matches []string, task *models.Task, line string) (string, error) {
	if len(matches) >= 3 {
		task.Priority = matches[2]
	}
	return p.priorityRgx.ReplaceAllString(line, emptyStr), nil
}

func (p *Parser) handleCreationDate(matches []string, task *models.Task, line string) (string, error) {
	if len(matches) >= 3 {
		date, err := time.ParseInLocation("2006-01-02", matches[2], time.Local)
		if err != nil {
			return line, err
		}
		task.CreationDate = date
	}
	return p.creationDateRgx.ReplaceAllString(line, emptyStr), nil
}

func (p *Parser) handleCompletionDate(matches []string, task *models.Task, line string) (string, error) {
	if len(matches) >= 2 {
		date, err := time.ParseInLocation("2006-01-02", matches[1], time.Local)
		if err != nil {
			return line, err
		}
		task.CompletionDate = date
	}
	return p.completionDateRgx.ReplaceAllString(line, emptyStr), nil
}

func (p *Parser) handleContext(matches []string, task *models.Task, line string) (string, error) {
	allMatches := p.contextRgx.FindAllStringSubmatch(line, -1)
	for _, match := range allMatches {
		if len(match) >= 3 {
			context := match[2]
			found := false
			for _, existing := range task.Contexts {
				if existing == context {
					found = true
					break
				}
			}
			if !found {
				task.Contexts = append(task.Contexts, context)
			}
		}
	}
	return p.contextRgx.ReplaceAllString(line, "$1"), nil
}

func (p *Parser) handleProjects(matches []string, task *models.Task, line string) (string, error) {
	allMatches := p.projectsRgx.FindAllStringSubmatch(line, -1)
	for _, match := range allMatches {
		if len(match) >= 3 {
			project := match[2]
			found := false
			for _, existing := range task.Projects {
				if existing == project {
					found = true
					break
				}
			}
			if !found {
				task.Projects = append(task.Projects, project)
			}
		}
	}
	return p.projectsRgx.ReplaceAllString(line, "$1"), nil
}

func (p *Parser) handleTags(matches []string, task *models.Task, line string) (string, error) {
	if task.Tags == nil {
		task.Tags = make(map[string]string)
	}

	allMatches := p.tagsRgx.FindAllStringSubmatch(line, -1)
	for _, match := range allMatches {
		if len(match) >= 3 {
			key := match[1]
			value := match[2]
			task.Tags[key] = value
		}
	}
	return p.tagsRgx.ReplaceAllString(line, emptyStr), nil
}

func (p *Parser) getPatterns() map[*regexp.Regexp]PatternHandler {
	return map[*regexp.Regexp]PatternHandler{
		p.completedRgx:      p.handleCompleted,
		p.priorityRgx:       p.handlePriority,
		p.creationDateRgx:   p.handleCreationDate,
		p.completionDateRgx: p.handleCompletionDate,
		p.contextRgx:        p.handleContext,
		p.projectsRgx:       p.handleProjects,
		p.tagsRgx:           p.handleTags,
	}
}
