package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"unicode"
)

const formHTML = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Combo Generator</title>
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<style>
		:root {
			--bg: #f5f5f4;
			--page-bg: #fbfbfa;
			--border: #e4e4e0;
			--text-main: #111827;
			--text-muted: #6b7280;
			--accent: #2563eb;
			--accent-soft: #dbeafe;
		}

		* {
			box-sizing: border-box;
		}

		body {
			margin: 0;
			padding: 32px 16px;
			font-family: system-ui, -apple-system, BlinkMacSystemFont, "SF Pro Text",
				Segoe UI, sans-serif;
			background: radial-gradient(circle at top left, #f9fafb, #e5e7eb);
			color: var(--text-main);
			display: flex;
			justify-content: center;
		}

		.page {
			width: 100%;
			max-width: 860px;
			background: var(--page-bg);
			border-radius: 18px;
			border: 1px solid var(--border);
			box-shadow: 0 18px 45px rgba(15, 23, 42, 0.12);
			padding: 28px 30px 22px;
		}

		.page-header {
			display: flex;
			align-items: center;
			gap: 12px;
			margin-bottom: 24px;
		}

		.page-icon {
			width: 34px;
			height: 34px;
			border-radius: 10px;
			background: #f97316;
			display: flex;
			align-items: center;
			justify-content: center;
			font-size: 20px;
		}

		.page-title {
			font-size: 24px;
			font-weight: 600;
		}

		.page-subtitle {
			font-size: 13px;
			color: var(--text-muted);
			margin-top: 4px;
		}

		form {
			margin-top: 10px;
		}

		.field-group {
			margin-bottom: 18px;
		}

		label {
			display: block;
			font-size: 13px;
			font-weight: 500;
			margin-bottom: 6px;
		}

		.input-row {
			display: flex;
			gap: 10px;
			align-items: center;
		}

		input[type="text"] {
			width: 100%;
			padding: 9px 10px;
			border-radius: 9px;
			border: 1px solid var(--border);
			background: #ffffff;
			font-size: 13px;
			outline: none;
			transition: border 0.15s ease, box-shadow 0.15s ease, background 0.15s ease;
		}

		input[type="text"]::placeholder {
			color: #9ca3af;
		}

		input[type="text"]:focus {
			border-color: var(--accent);
			box-shadow: 0 0 0 1px var(--accent-soft);
			background: #ffffff;
		}

		.helper {
			font-size: 12px;
			color: var(--text-muted);
			margin-top: 4px;
		}

		.tag {
			display: inline-flex;
			align-items: center;
			gap: 6px;
			border-radius: 999px;
			border: 1px solid #e5e7eb;
			padding: 3px 9px;
			font-size: 11px;
			color: #4b5563;
			background: #f9fafb;
		}

		.tag-dot {
			width: 7px;
			height: 7px;
			border-radius: 999px;
			background: var(--accent);
		}

		button[type="submit"] {
			margin-top: 8px;
			padding: 8px 18px;
			border-radius: 999px;
			border: 1px solid #d1d5db;
			background: #111827;
			color: white;
			font-size: 13px;
			font-weight: 500;
			cursor: pointer;
			display: inline-flex;
			align-items: center;
			gap: 8px;
			box-shadow: 0 8px 18px rgba(15, 23, 42, 0.35);
			transition: transform 0.08s ease, box-shadow 0.08s ease, background 0.08s ease;
		}

		button[type="submit"]:hover {
			background: #030712;
			transform: translateY(-1px);
			box-shadow: 0 14px 26px rgba(15, 23, 42, 0.55);
		}

		button[type="submit"]:active {
			transform: translateY(0);
			box-shadow: 0 6px 12px rgba(15, 23, 42, 0.3);
		}

		.btn-icon {
			font-size: 14px;
		}

		.footer-line {
			margin-top: 22px;
			padding-top: 10px;
			border-top: 1px solid #e5e7eb;
			display: flex;
			justify-content: space-between;
			align-items: center;
			font-size: 11px;
			color: #9ca3af;
		}

		.footer-right {
			color: #6b7280;
			font-style: italic;
		}

		.footer-left {
			display: flex;
			align-items: center;
			gap: 8px;
		}

		.footer-dot {
			width: 6px;
			height: 6px;
			border-radius: 999px;
			background: #22c55e;
		}

		@media (max-width: 640px) {
			.page {
				padding: 22px 18px 16px;
				border-radius: 14px;
			}
			.page-header {
				flex-direction: row;
				align-items: flex-start;
			}
		}
	</style>
</head>
<body>
	<div class="page">
		<div class="page-header">
			<div class="page-icon">🔐</div>
			<div>
				<div class="page-title">Combo Generator</div>
				<div class="page-subtitle">
					Create all combinations: two text parts + number tail from birth date / custom list.
				</div>
			</div>
		</div>

		<form method="POST">
			<div class="field-group">
				<label for="firstNames">First names</label>
				<input type="text" id="firstNames" name="firstNames" placeholder="Kylie, Ky">
				<div class="helper">Comma-separated variations. Example: <b>Kylie, Ky, KJ</b>.</div>
			</div>

			<div class="field-group">
				<label for="lastName">Last name</label>
				<input type="text" id="lastName" name="lastName" placeholder="Johnson">
			</div>

			<div class="field-group">
				<label for="birthDate">Birth date (DD.MM.YYYY)</label>
				<input type="text" id="birthDate" name="birthDate" placeholder="10.10.2001">
				<div class="helper">Numbers will be used only at the end, like <b>KylieJohn1010</b>.</div>
			</div>

			<div class="field-group">
				<label for="customNumbers">Numbers</label>
				<input type="text" id="customNumbers" name="customNumbers" placeholder="12, 007, 99">
				<div class="helper">Comma-separated numbers that will also be appended at the end (same as birth date tails).</div>
			</div>

			<div class="field-group">
				<div class="input-row">
					<div style="flex:1">
						<label for="relatives">Relatives</label>
						<input type="text" id="relatives" name="relatives" placeholder="John, Maria, Masha">
					</div>
					<div>
						<div class="tag">
							<span class="tag-dot"></span>
							<span>2 text parts + number</span>
						</div>
					</div>
				</div>
				<div class="helper">
					Any pair: (first / last / relative) + (first / last / relative) + numeric tail.
				</div>
			</div>

			<button type="submit">
				<span class="btn-icon">⬇</span>
				<span>Generate TXT file</span>
			</button>
		</form>

		<div class="footer-line">
			<div class="footer-left">
				<div class="footer-dot"></div>
				<div>Server: localhost:8080</div>
			</div>
			<div class="footer-right">
				i do not know why u need this for
			</div>
		</div>
	</div>
</body>
</html>`

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, formHTML)
		return
	}

	// POST: generate file
	if err := r.ParseForm(); err != nil {
		http.Error(w, "cannot parse form", http.StatusBadRequest)
		return
	}

	firstNames := parseList(r.FormValue("firstNames"))
	lastName := strings.TrimSpace(r.FormValue("lastName"))
	relatives := parseList(r.FormValue("relatives"))
	birthDate := strings.TrimSpace(r.FormValue("birthDate"))
	customNumbers := parseList(r.FormValue("customNumbers"))

	// collect all text components
	var components []string
	components = append(components, firstNames...)
	if lastName != "" {
		components = append(components, lastName)
	}
	components = append(components, relatives...)

	// numeric tails from birth date
	dateNumbers := buildDateNumbers(birthDate)

	// append custom numbers as additional tails
	for _, cn := range customNumbers {
		cnDigits := digitsOnly(cn)
		if cnDigits != "" {
			dateNumbers = append(dateNumbers, cnDigits)
		}
	}

	// require at least one numeric tail source
	if len(dateNumbers) == 0 {
		http.Error(w, "provide a valid birth date and/or a non-empty numbers list", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment; filename=\"combinations.txt\"")

	// generate: any component1 + component2 + number
	for i := 0; i < len(components); i++ {
		for j := 0; j < len(components); j++ {
			if i == j {
				continue
			}
			for _, num := range dateNumbers {
				writeVariants(w, components[i], components[j], num)
			}
		}
	}
}

// parseList splits by comma and trims spaces
func parseList(s string) []string {
	var res []string
	parts := strings.Split(s, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			res = append(res, p)
		}
	}
	return res
}

// buildDateNumbers builds numeric endings from birth date.
func buildDateNumbers(dateStr string) []string {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" {
		return nil
	}

	parts := strings.Split(dateStr, ".")
	if len(parts) != 3 {
		// fallback: just keep digits from whatever was provided
		digits := digitsOnly(dateStr)
		if digits == "" {
			return nil
		}
		return []string{digits}
	}

	dd := digitsOnly(parts[0])
	mm := digitsOnly(parts[1])
	yyyy := digitsOnly(parts[2])

	if dd == "" || mm == "" || yyyy == "" {
		return nil
	}

	var nums []string
	add := func(s string) {
		if s != "" {
			nums = append(nums, s)
		}
	}

	// basic variants
	add(dd + mm) // 1010
	if len(yyyy) == 4 {
		add(dd + mm + yyyy[2:]) // 101001
	}
	add(dd + mm + yyyy) // 10102001
	add(yyyy)           // 2001
	add(yyyy + mm + dd) // 20011010

	// dedupe
	seen := make(map[string]struct{})
	var out []string
	for _, n := range nums {
		if _, ok := seen[n]; !ok {
			seen[n] = struct{}{}
			out = append(out, n)
		}
	}
	return out
}

// generate all case variants for one combo and write to response
func writeVariants(w http.ResponseWriter, part1, part2, num string) {
	p1 := strings.ToLower(strings.TrimSpace(part1))
	p2 := strings.ToLower(strings.TrimSpace(part2))

	base := p1 + p2 + num

	local := make(map[string]struct{})
	add := func(s string) {
		if s == "" {
			return
		}
		if _, ok := local[s]; ok {
			return
		}
		local[s] = struct{}{}
		fmt.Fprintln(w, s)
	}

	// all lower
	add(base)
	// Capitalized parts
	cp1 := capitalize(p1)
	cp2 := capitalize(p2)
	add(cp1 + cp2 + num)
	// all upper
	add(strings.ToUpper(base))
	// first letter of whole combo capitalized
	add(capitalize(base))
}

// capitalize first rune, rest lower
func capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

// digitsOnly keeps only 0–9
func digitsOnly(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
