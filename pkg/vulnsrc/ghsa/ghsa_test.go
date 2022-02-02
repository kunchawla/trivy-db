package ghsa

import (
	"github.com/aquasecurity/trivy-db/pkg/dbtest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/aquasecurity/trivy-db/pkg/db"
	"github.com/aquasecurity/trivy-db/pkg/types"
	"github.com/aquasecurity/trivy-db/pkg/vulnsrc/vulnerability"
)

func TestVulnSrc_Update(t *testing.T) {
	type want struct {
		key   []string
		value interface{}
	}
	tests := []struct {
		name       string
		dir        string
		wantValues []want
		wantErr    string
	}{
		{
			name: "happy path",
			dir:  filepath.Join("testdata", "happy"),
			wantValues: []want{
				{
					key: []string{"data-source", "composer::GitHub Security Advisory Composer"},
					value: types.DataSource{
						ID:   vulnerability.GHSA,
						Name: "GitHub Security Advisory Composer",
						URL:  "https://github.com/advisories?query=type%3Areviewed+ecosystem%3Acomposer",
					},
				},
				{
					key: []string{"advisory-detail", "CVE-2019-19745", "composer::GitHub Security Advisory Composer", "contao/core-bundle"},
					value: types.Advisory{
						PatchedVersions:    []string{"4.8.6", "4.4.46"},
						VulnerableVersions: []string{"\u003e= 4.5.0, \u003c 4.8.6", "\u003e= 4.0.0, \u003c 4.4.46"},
					},
				},
				{
					key: []string{"vulnerability-detail", "CVE-2019-19745", ghsaDir},
					value: types.VulnerabilityDetail{
						ID:          "CVE-2019-19745",
						Title:       "Unrestricted file uploads in Contao",
						Description: "### Impact\n\nA back end user with access to the form generator can upload arbitrary files and execute them on the server.\n\n### Patches\n\nUpdate to Contao 4.4.46 or 4.8.6.\n\n### Workarounds\n\nConfigure your web server so it does not execute PHP files and other scripts in the Contao file upload directory.\n\n### References\n\nhttps://contao.org/en/security-advisories/unrestricted-file-uploads.html\n\n### For more information\n\nIf you have any questions or comments about this advisory, open an issue in [contao/contao](https://github.com/contao/contao/issues/new/choose).",
						References: []string{
							"https://github.com/contao/contao/security/advisories/GHSA-wjx8-cgrm-hh8p",
							"https://nvd.nist.gov/vuln/detail/CVE-2019-19745",
							"https://contao.org/en/news.html",
							"https://contao.org/en/security-advisories/unrestricted-file-uploads.html",
							"https://github.com/FriendsOfPHP/security-advisories/blob/master/contao/core-bundle/CVE-2019-19745.yaml",
							"https://github.com/advisories/GHSA-wjx8-cgrm-hh8p",
						},
						Severity: types.SeverityHigh,
					},
				},
				{
					key:   []string{"vulnerability-id", "CVE-2019-19745"},
					value: map[string]interface{}{},
				},
				{
					key: []string{"data-source", "maven::GitHub Security Advisory Maven"},
					value: types.DataSource{
						ID:   vulnerability.GHSA,
						Name: "GitHub Security Advisory Maven",
						URL:  "https://github.com/advisories?query=type%3Areviewed+ecosystem%3Amaven",
					},
				},
				{
					key: []string{"advisory-detail", "CVE-2018-1196", "maven::GitHub Security Advisory Maven", "org.springframework.boot:spring-boot"},
					value: types.Advisory{
						PatchedVersions:    []string{"1.5.10"},
						VulnerableVersions: []string{"\u003e= 1.5.0, \u003c 1.5.10"},
					},
				},
				{
					key: []string{"vulnerability-detail", "CVE-2018-1196", ghsaDir},
					value: types.VulnerabilityDetail{
						ID:          "CVE-2018-1196",
						Title:       "Moderate severity vulnerability that affects org.springframework.boot:spring-boot",
						Description: "Spring Boot supports an embedded launch script that can be used to easily run the application as a systemd or init.d linux service. The script included with Spring Boot 1.5.9 and earlier and 2.0.0.M1 through 2.0.0.M7 is susceptible to a symlink attack which allows the \"run_user\" to overwrite and take ownership of any file on the same system. In order to instigate the attack, the application must be installed as a service and the \"run_user\" requires shell access to the server. Spring Boot application that are not installed as a service, or are not using the embedded launch script are not susceptible.",
						References: []string{
							"https://nvd.nist.gov/vuln/detail/CVE-2018-1196",
							"https://github.com/advisories/GHSA-xx65-cc7g-9pfp",
							"https://pivotal.io/security/cve-2018-1196",
						},
						Severity: types.SeverityMedium,
					},
				},
				{
					key:   []string{"vulnerability-id", "CVE-2018-1196"},
					value: map[string]interface{}{},
				},
				{
					key: []string{"data-source", "npm::GitHub Security Advisory Npm"},
					value: types.DataSource{
						ID:   vulnerability.GHSA,
						Name: "GitHub Security Advisory Npm",
						URL:  "https://github.com/advisories?query=type%3Areviewed+ecosystem%3Anpm",
					},
				},
				{
					key: []string{"advisory-detail", "CVE-2018-3745", "npm::GitHub Security Advisory Npm", "atob"},
					value: types.Advisory{
						PatchedVersions:    []string{"2.1.0"},
						VulnerableVersions: []string{"\u003c 2.1.0"},
					},
				},
				{
					key: []string{"vulnerability-detail", "CVE-2018-3745", ghsaDir},
					value: types.VulnerabilityDetail{
						ID:          "CVE-2018-3745",
						Title:       "Out-of-bounds Read in atob",
						Description: "Versions of `atob` before 2.1.0  uninitialized Buffers when number is passed in input on Node.js 4.x and below.\n\n\n## Recommendation\n\nUpdate to version 2.1.0 or later.",
						References: []string{
							"https://nvd.nist.gov/vuln/detail/CVE-2018-3745",
							"https://hackerone.com/reports/321686",
							"https://github.com/advisories/GHSA-8w4h-3cm3-2pm2",
							"https://www.npmjs.com/advisories/646",
						},
						Severity: types.SeverityMedium,
					},
				},
				{
					key:   []string{"vulnerability-id", "CVE-2018-3745"},
					value: map[string]interface{}{},
				},
				{
					key: []string{"data-source", "nuget::GitHub Security Advisory Nuget"},
					value: types.DataSource{
						ID:   vulnerability.GHSA,
						Name: "GitHub Security Advisory Nuget",
						URL:  "https://github.com/advisories?query=type%3Areviewed+ecosystem%3Anuget",
					},
				},
				{
					key: []string{"advisory-detail", "CVE-2019-1010113", "nuget::GitHub Security Advisory Nuget", "CLEditor"},
					value: types.Advisory{
						VulnerableVersions: []string{"\u003c= 1.4.5"},
					},
				},
				{
					key: []string{"vulnerability-detail", "CVE-2019-1010113", ghsaDir},
					value: types.VulnerabilityDetail{
						ID:          "CVE-2019-1010113",
						Title:       "Moderate severity vulnerability that affects CLEditor",
						Description: "Premium Software CLEditor 1.4.5 and earlier is affected by: Cross Site Scripting (XSS). The impact is: An attacker might be able to inject arbitrary html and script code into the web site. The component is: jQuery plug-in. The attack vector is: the victim must open a crafted href attribute of a link (A) element.",
						References: []string{
							"https://nvd.nist.gov/vuln/detail/CVE-2019-1010113",
						},
						Severity: types.SeverityMedium,
					},
				},
				{
					key:   []string{"vulnerability-id", "CVE-2019-1010113"},
					value: map[string]interface{}{},
				},
				{
					key: []string{"data-source", "pip::GitHub Security Advisory Pip"},
					value: types.DataSource{
						ID:   vulnerability.GHSA,
						Name: "GitHub Security Advisory Pip",
						URL:  "https://github.com/advisories?query=type%3Areviewed+ecosystem%3Apip",
					},
				},
				{
					key: []string{"advisory-detail", "CVE-2018-14574", "pip::GitHub Security Advisory Pip", "django"},
					value: types.Advisory{
						PatchedVersions:    []string{"2.0.8", "1.11.15"},
						VulnerableVersions: []string{"\u003e= 2.0, \u003c 2.0.8", "\u003e= 1.11.0, \u003c 1.11.15"},
					},
				},
				{
					key: []string{"vulnerability-detail", "CVE-2018-14574", ghsaDir},
					value: types.VulnerabilityDetail{
						ID:          "CVE-2018-14574",
						Title:       "Moderate severity vulnerability that affects django",
						Description: "django.middleware.common.CommonMiddleware in Django 1.11.x before 1.11.15 and 2.0.x before 2.0.8 has an Open Redirect.",
						References: []string{
							"https://nvd.nist.gov/vuln/detail/CVE-2018-14574",
							"https://access.redhat.com/errata/RHSA-2019:0265",
							"https://github.com/advisories/GHSA-5hg3-6c2f-f3wr",
							"https://usn.ubuntu.com/3726-1/",
							"https://www.debian.org/security/2018/dsa-4264",
							"https://www.djangoproject.com/weblog/2018/aug/01/security-releases/",
							"http://www.securityfocus.com/bid/104970",
							"http://www.securitytracker.com/id/1041403",
						},
						Severity: types.SeverityMedium,
					},
				},
				{
					key:   []string{"vulnerability-id", "CVE-2018-14574"},
					value: map[string]interface{}{},
				},
				{
					key: []string{"data-source", "rubygems::GitHub Security Advisory Rubygems"},
					value: types.DataSource{
						ID:   vulnerability.GHSA,
						Name: "GitHub Security Advisory Rubygems",
						URL:  "https://github.com/advisories?query=type%3Areviewed+ecosystem%3Arubygems",
					},
				},
				{
					key: []string{"advisory-detail", "CVE-2018-16477", "rubygems::GitHub Security Advisory Rubygems", "activestorage"},
					value: types.Advisory{
						PatchedVersions:    []string{"5.2.1.1"},
						VulnerableVersions: []string{"\u003e= 5.2.0, \u003c= 5.2.1.0"},
					},
				},
				{
					key: []string{"vulnerability-detail", "CVE-2018-16477", ghsaDir},
					value: types.VulnerabilityDetail{
						ID:          "CVE-2018-16477",
						Title:       "High severity vulnerability that affects activestorage",
						Description: "A bypass vulnerability in Active Storage >= 5.2.0 for Google Cloud Storage and Disk services allow an attacker to modify the `content-disposition` and `content-type` parameters which can be used in with HTML files and have them executed inline. Additionally, if combined with other techniques such as cookie bombing and specially crafted AppCache manifests, an attacker can gain access to private signed URLs within a specific storage path.",
						References: []string{
							"https://nvd.nist.gov/vuln/detail/CVE-2018-16477",
							"https://github.com/advisories/GHSA-7rr7-rcjw-56vj",
							"https://groups.google.com/d/msg/rubyonrails-security/3KQRnXDIuLg/mByx5KkqBAAJ",
							"https://weblog.rubyonrails.org/2018/11/27/Rails-4-2-5-0-5-1-5-2-have-been-released/",
						},
						Severity: types.SeverityMedium,
					},
				},
				{
					key:   []string{"vulnerability-id", "CVE-2018-16477"},
					value: map[string]interface{}{},
				},
			},
		},
		{
			name:    "sad path (dir doesn't exist)",
			dir:     filepath.Join("testdata", "badpathdoesnotexist"),
			wantErr: "no such file or directory",
		},
		{
			name:    "sad path (failed to decode)",
			dir:     filepath.Join("testdata", "sad"),
			wantErr: "failed to decode GHSA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()

			err := db.Init(tempDir)
			require.NoError(t, err)
			defer db.Close()

			vs := NewVulnSrc()
			err = vs.Update(tt.dir)
			if tt.wantErr != "" {
				require.NotNil(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}

			require.NoError(t, err)
			require.NoError(t, db.Close()) // Need to close before dbtest.JSONEq is called
			for _, w := range tt.wantValues {
				dbtest.JSONEq(t, db.Path(tempDir), w.key, w.value, w.key)
			}
		})
	}
}
