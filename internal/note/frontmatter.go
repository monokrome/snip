package note

import (
	"encoding/json"
	"strings"

	"gopkg.in/yaml.v3"
)

func FrontmatterErr(yamlErr error) string {
	if yamlErr != nil {
		return ": " + yamlErr.Error()
	}
	return ""
}

func FormatFrontmatter(metadata string) string {
	if metadata == "" || metadata == "{}" {
		return ""
	}

	var data map[string]any
	if err := json.Unmarshal([]byte(metadata), &data); err != nil {
		return ""
	}

	if len(data) == 0 {
		return ""
	}

	yamlBytes, err := yaml.Marshal(data)
	if err != nil {
		return ""
	}

	return "---\n" + string(yamlBytes) + "---\n"
}

func (n *NoteWithTags) FullContent() string {
	return FormatFrontmatter(n.Metadata) + n.Content
}

func ParseFrontmatter(content string) (metadata string, body string, ok bool, yamlErr error) {
	content = strings.TrimSpace(content)

	if !strings.HasPrefix(content, "---") {
		return "", content, true, nil
	}

	rest := content[3:]
	endIdx := strings.Index(rest, "\n---")
	if endIdx < 0 {
		return "", content, false, nil
	}

	frontmatter := rest[:endIdx]
	body = strings.TrimSpace(rest[endIdx+4:])

	var data map[string]any
	if err := yaml.Unmarshal([]byte(frontmatter), &data); err != nil {
		return "", content, false, err
	}

	metaBytes, err := json.Marshal(data)
	if err != nil {
		return "", content, false, nil
	}

	return string(metaBytes), body, true, nil
}
