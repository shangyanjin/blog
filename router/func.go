package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"math"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// GetTemplateFuncMap returns a template.FuncMap with common template functions
func GetTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		// Add basic template functions
		"add": func(a, b int) int {
			return a + b
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"mul": func(a, b int) int {
			return a * b
		},
		"div": func(a, b int) int {
			if b == 0 {
				return 0
			}
			return a / b
		},
		"mod": func(a, b int) int {
			if b == 0 {
				return 0
			}
			return a % b
		},
		"eq": func(a, b interface{}) bool {
			return a == b
		},
		"neq": func(a, b interface{}) bool {
			return a != b
		},
		"lt": func(a, b int) bool {
			return a < b
		},
		"lte": func(a, b int) bool {
			return a <= b
		},
		"gt": func(a, b int) bool {
			return a > b
		},
		"gte": func(a, b int) bool {
			return a >= b
		},
		"hasPrefix":    strings.HasPrefix,
		"hasSuffix":    strings.HasSuffix,
		"contains":     strings.Contains,
		"lower":        strings.ToLower,
		"upper":        strings.ToUpper,
		"join":         strings.Join,
		"split":        strings.Split,
		"trim":         strings.Trim,
		"trimSpace":    strings.TrimSpace,
		"trimPrefix":   strings.TrimPrefix,
		"trimSuffix":   strings.TrimSuffix,
		"trimLeft":     strings.TrimLeft,
		"trimRight":    strings.TrimRight,
		"replace":      strings.Replace,
		"replaceAll":   strings.ReplaceAll,
		"repeat":       strings.Repeat,
		"containsAny":  strings.ContainsAny,
		"containsRune": strings.ContainsRune,
		"count":        strings.Count,
		"index":        strings.Index,
		"indexAny":     strings.IndexAny,
		"indexRune":    strings.IndexRune,
		"lastIndex":    strings.LastIndex,
		"lastIndexAny": strings.LastIndexAny,
		"unescapeHTML": html.UnescapeString,
		"prettyPrint": func(v interface{}) string {
			b, err := json.Marshal(v)
			if err != nil {
				return fmt.Sprintf("Error marshaling: %v", err)
			}

			var out bytes.Buffer
			err = json.Indent(&out, b, "", "    ")
			if err != nil {
				return fmt.Sprintf("Error indenting: %v", err)
			}

			return out.String()
		},
		"formatDate": func(t interface{}) string {
			if t == nil {
				return "-"
			}

			var timeStr string
			switch v := t.(type) {
			case string:
				timeStr = v
			case time.Time:
				timeStr = v.Format("2006-01-02 15:04:05")
			default:
				// Handle any other type by converting to string
				timeStr = fmt.Sprint(v)
			}

			if timeStr == "" {
				return "-"
			}

			// Try parsing with multiple formats
			formats := []string{
				"2006-01-02 15:04:05",
				"2006-01-02T15:04:05Z",
				time.RFC3339,
				time.RFC3339Nano,
				time.RFC822,
				time.RFC822Z,
				time.RFC850,
				time.RFC1123,
				time.RFC1123Z,
				"2006-01-02",
				"2006/01/02 15:04:05",
				"02/01/2006 15:04:05",
				"2006-01-02 15:04",
				"2006/01/02",
				"20060102150405",
				"2006.01.02 15:04:05",
				"2006.01.02",
				"02-01-2006",
				"02/01/2006",
				"02.01.2006",
				"January 2, 2006",
				"Jan 2, 2006",
				"2006年01月02日",
				"2006年1月2日",
				"2006-1-2",
				"2006/1/2",
				"2006.1.2",
				"20060102",
			}

			var timeValue time.Time
			var err error
			for _, format := range formats {
				timeValue, err = time.Parse(format, timeStr)
				if err == nil {
					break
				}
			}
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"value": timeStr,
					"error": err,
				}).Warn("Failed to parse date string")
				return timeStr
			}

			now := time.Now()
			diff := now.Sub(timeValue)

			switch {
			case diff < time.Second*30:
				return "just now"
			case diff < time.Minute:
				seconds := int(diff.Seconds())
				return fmt.Sprintf("%d seconds ago", seconds)
			case diff < time.Hour:
				minutes := int(diff.Minutes())
				if minutes == 1 {
					return "1 minute ago"
				}
				return fmt.Sprintf("%d minutes ago", minutes)
			case diff < time.Hour*24:
				hours := int(diff.Hours())
				if hours == 1 {
					return "1 hour ago"
				}
				return fmt.Sprintf("%d hours ago", hours)
			case diff < time.Hour*24*30:
				days := int(diff.Hours() / 24)
				if days == 1 {
					return "1 day ago"
				}
				return fmt.Sprintf("%d days ago", days)
			case diff < time.Hour*24*365:
				months := int(diff.Hours() / 24 / 30)
				if months == 1 {
					return "1 month ago"
				}
				return fmt.Sprintf("%d months ago", months)
			default:
				years := int(diff.Hours() / 24 / 365)
				if years == 1 {
					return "1 year ago"
				}
				return fmt.Sprintf("%d years ago", years)
			}
		},

		"ceil": func(x interface{}) int {
			var f float64
			switch v := x.(type) {
			case float64:
				f = v
			case int:
				f = float64(v)
			case int64:
				f = float64(v)
			default:
				return 0
			}
			return int(math.Ceil(f))
		},
		"unescape": func(s string) string {
			return html.UnescapeString(s)
		},
		"raw": func(s interface{}) template.HTML {
			switch v := s.(type) {
			case template.HTML:
				return v
			case string:
				return template.HTML(v)
			default:
				return template.HTML(fmt.Sprint(v))
			}
		},
		"seq": func(start, end int) []int {
			if end < start {
				return []int{}
			}
			n := end - start + 1
			result := make([]int, n)
			for i := 0; i < n; i++ {
				result[i] = start + i
			}
			return result
		},
		"Loop": func(n int) []int {
			result := make([]int, n)
			for i := 0; i < n; i++ {
				result[i] = i
			}
			return result
		},
	}
}
