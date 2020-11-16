package action

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionToolchains() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("rustup", "toolchain", "list").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)

			re := regexp.MustCompile(`^(?P<toolchain>[^ ]+)( \((?P<info>.*)\))?$`)
			for _, line := range strings.Split(string(output), "\n") {
				if re.MatchString(line) {
					matches := re.FindStringSubmatch(line)
					vals = append(vals, matches[1], matches[3])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}

func ActionTargets(installedOnly bool) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		opts := []string{"target", "list"}
		if installedOnly {
			opts = append(opts, "--installed")
		}

		if output, err := exec.Command("rustup", opts...).Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)

			re := regexp.MustCompile(`^(?P<target>[^ ]+)( \((?P<info>.*)\))?$`)
			for _, line := range strings.Split(string(output), "\n") {
				if re.MatchString(line) {
					matches := re.FindStringSubmatch(line)
					vals = append(vals, matches[1], matches[3])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}

func ActionComponents(installedOnly bool) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		opts := []string{"component", "list"}
		if installedOnly {
			opts = append(opts, "--installed")
		}

		if output, err := exec.Command("rustup", opts...).Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)

			re := regexp.MustCompile(`^(?P<target>[^ ]+)( \((?P<info>.*)\))?$`)
			for _, line := range strings.Split(string(output), "\n") {
				if re.MatchString(line) {
					matches := re.FindStringSubmatch(line)
					vals = append(vals, matches[1], matches[3])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}

func ActionAvailableComponents() carapace.Action {
	return carapace.ActionValuesDescribed(
		"rustc", "The Rust compiler.",
		"cargo", "Cargo is a package manager and build tool.",
		"rustfmt", "Rustfmt is a tool for automatically formatting code.",
		"rust-std", "This is the Rust standard library.",
		"rust-docs", "This is a local copy of the Rust documentation.",
		"rls", "RLS is a language server",
		"rust-analyzer-preview", "Rust anaylzer is a new language server.",
		"clippy", "Clippy is a lint tool.",
		"miri", "Miri is an experimental Rust interpreter.",
		"rust-src", "This is a local copy of the source code of the Rust standard library.",
		"rust-analysis", "Metadata about the standard library, used by tools like RLS.",
		"rust-mingw", "This contains a linker and platform libraries.",
		"llvm-tools-preview", "This contains a collection of LLVM tools.",
		"rustc-dev", "This component contains the compiler as a library.",
	)
}

func ActionOverrides() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("rustup", "override", "list").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)

			re := regexp.MustCompile(`^(?P<path>.*)\t(?P<toolchain>[^ ]+)$`)
			for _, line := range strings.Split(string(output), "\n") {
				if re.MatchString(line) {
					matches := re.FindStringSubmatch(line)
					vals = append(vals, matches[1], matches[2])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}
