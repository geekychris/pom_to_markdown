package main


import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// POM represents the Maven POM structure
type POM struct {
	XMLName      xml.Name     `xml:"project"`
	GroupId      string       `xml:"groupId"`
	ArtifactId   string       `xml:"artifactId"`
	Version      string       `xml:"version"`
	Packaging    string       `xml:"packaging"`
	Name         string       `xml:"name"`
	Description  string       `xml:"description"`
	URL          string       `xml:"url"`
	Parent       *Parent      `xml:"parent"`
	Properties   *Properties  `xml:"properties"`
	Dependencies []Dependency `xml:"dependencies>dependency"`
	Plugins      []Plugin     `xml:"build>plugins>plugin"`
	Developers   []Developer  `xml:"developers>developer"`
	Licenses     []License    `xml:"licenses>license"`
	SCM          *SCM         `xml:"scm"`
}

type Parent struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
}


type Properties struct {
	JavaVersion    string `xml:"java.version"`
	MavenCompiler  string `xml:"maven.compiler.source"`
	ProjectBuild   string `xml:"project.build.sourceEncoding"`
	SpringVersion  string `xml:"spring.version"`
	JunitVersion   string `xml:"junit.version"`
}

type Dependency struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
	Scope      string `xml:"scope"`
	Type       string `xml:"type"`
	Optional   string `xml:"optional"`
}

type Plugin struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
}

type Developer struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
	ID    string `xml:"id"`
}

type License struct {
	Name string `xml:"name"`
	URL  string `xml:"url"`
}

type SCM struct {
	URL        string `xml:"url"`
	Connection string `xml:"connection"`
	Tag        string `xml:"tag"`
}

func main() {
	var inputFile = flag.String("input", "pom.xml", "Input POM XML file path")
	var outputFile = flag.String("output", "", "Output Markdown file path (default: stdout)")
	flag.Parse()

	// Read the POM XML file
	xmlData, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", *inputFile, err)
		os.Exit(1)
	}

	// Parse the XML
	var pom POM
	err = xml.Unmarshal(xmlData, &pom)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing XML: %v\n", err)
		os.Exit(1)
	}

	// Convert to Markdown
	markdown := convertToMarkdown(&pom)

	// Output the result
	if *outputFile != "" {
		err = ioutil.WriteFile(*outputFile, []byte(markdown), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file %s: %v\n", *outputFile, err)
			os.Exit(1)
		}
		fmt.Printf("Markdown documentation generated: %s\n", *outputFile)
	} else {
		fmt.Print(markdown)
	}
}

func convertToMarkdown(pom *POM) string {
	var md strings.Builder

	// Project title
	if pom.Name != "" {
		md.WriteString(fmt.Sprintf("# %s\n\n", pom.Name))
	} else {
		md.WriteString(fmt.Sprintf("# %s\n\n", pom.ArtifactId))
	}

	// Basic project information
	md.WriteString("## Project Information\n\n")
	md.WriteString(fmt.Sprintf("- **Group ID**: %s\n", pom.GroupId))
	md.WriteString(fmt.Sprintf("- **Artifact ID**: %s\n", pom.ArtifactId))
	md.WriteString(fmt.Sprintf("- **Version**: %s\n", pom.Version))
	
	if pom.Packaging != "" {
		md.WriteString(fmt.Sprintf("- **Packaging**: %s\n", pom.Packaging))
	}
	
	if pom.URL != "" {
		md.WriteString(fmt.Sprintf("- **URL**: [%s](%s)\n", pom.URL, pom.URL))
	}
	
	md.WriteString("\n")

	// Description
	if pom.Description != "" {
		md.WriteString("## Description\n\n")
		md.WriteString(fmt.Sprintf("%s\n\n", pom.Description))
	}

	// Parent project
	if pom.Parent != nil {
		md.WriteString("## Parent Project\n\n")
		md.WriteString(fmt.Sprintf("- **Group ID**: %s\n", pom.Parent.GroupId))
		md.WriteString(fmt.Sprintf("- **Artifact ID**: %s\n", pom.Parent.ArtifactId))
		md.WriteString(fmt.Sprintf("- **Version**: %s\n\n", pom.Parent.Version))
	}

	// Properties
	if pom.Properties != nil {
		md.WriteString("## Properties\n\n")
		if pom.Properties.JavaVersion != "" {
			md.WriteString(fmt.Sprintf("- **Java Version**: %s\n", pom.Properties.JavaVersion))
		}
		if pom.Properties.MavenCompiler != "" {
			md.WriteString(fmt.Sprintf("- **Maven Compiler Source**: %s\n", pom.Properties.MavenCompiler))
		}
		if pom.Properties.ProjectBuild != "" {
			md.WriteString(fmt.Sprintf("- **Source Encoding**: %s\n", pom.Properties.ProjectBuild))
		}
		if pom.Properties.SpringVersion != "" {
			md.WriteString(fmt.Sprintf("- **Spring Version**: %s\n", pom.Properties.SpringVersion))
		}
		if pom.Properties.JunitVersion != "" {
			md.WriteString(fmt.Sprintf("- **JUnit Version**: %s\n", pom.Properties.JunitVersion))
		}
		md.WriteString("\n")
	}

	// Dependencies
	if len(pom.Dependencies) > 0 {
		md.WriteString("## Dependencies\n\n")
		md.WriteString("| Group ID | Artifact ID | Version | Scope | Type |\n")
		md.WriteString("|----------|-------------|---------|-------|------|\n")
		
		for _, dep := range pom.Dependencies {
			scope := dep.Scope
			if scope == "" {
				scope = "compile"
			}
			depType := dep.Type
			if depType == "" {
				depType = "jar"
			}
			
			md.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
				dep.GroupId, dep.ArtifactId, dep.Version, scope, depType))
		}
		md.WriteString("\n")
	}

	// Build Plugins
	if len(pom.Plugins) > 0 {
		md.WriteString("## Build Plugins\n\n")
		md.WriteString("| Group ID | Artifact ID | Version |\n")
		md.WriteString("|----------|-------------|---------|\n")
		
		for _, plugin := range pom.Plugins {
			md.WriteString(fmt.Sprintf("| %s | %s | %s |\n",
				plugin.GroupId, plugin.ArtifactId, plugin.Version))
		}
		md.WriteString("\n")
	}

	// Developers
	if len(pom.Developers) > 0 {
		md.WriteString("## Developers\n\n")
		for _, dev := range pom.Developers {
			if dev.Name != "" {
				md.WriteString(fmt.Sprintf("- **%s**", dev.Name))
				if dev.Email != "" {
					md.WriteString(fmt.Sprintf(" <%s>", dev.Email))
				}
				if dev.ID != "" {
					md.WriteString(fmt.Sprintf(" (ID: %s)", dev.ID))
				}
				md.WriteString("\n")
			}
		}
		md.WriteString("\n")
	}

	// Licenses
	if len(pom.Licenses) > 0 {
		md.WriteString("## Licenses\n\n")
		for _, license := range pom.Licenses {
			if license.URL != "" {
				md.WriteString(fmt.Sprintf("- [%s](%s)\n", license.Name, license.URL))
			} else {
				md.WriteString(fmt.Sprintf("- %s\n", license.Name))
			}
		}
		md.WriteString("\n")
	}

	// Source Control Management
	if pom.SCM != nil {
		md.WriteString("## Source Control\n\n")
		if pom.SCM.URL != "" {
			md.WriteString(fmt.Sprintf("- **URL**: [%s](%s)\n", pom.SCM.URL, pom.SCM.URL))
		}
		if pom.SCM.Connection != "" {
			md.WriteString(fmt.Sprintf("- **Connection**: %s\n", pom.SCM.Connection))
		}
		if pom.SCM.Tag != "" {
			md.WriteString(fmt.Sprintf("- **Tag**: %s\n", pom.SCM.Tag))
		}
		md.WriteString("\n")
	}

	return md.String()
}
