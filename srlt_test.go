package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dependency", func() {

	Describe("Restore", func() {
		It("Exists before", func() {
			for _, dep := range srlt.Deps {
				Ω(dep.Exists()).Should(BeFalse())
			}
		})

		It("Clone", func() {
			for _, dep := range srlt.Deps {
				Ω(dep.Clone()).Should(BeNil())
			}
		})

		It("Pull", func() {
			for _, dep := range srlt.Deps {
				Ω(dep.Pull()).Should(BeNil())
			}
		})

		It("Checkout", func() {
			for _, dep := range srlt.Deps {
				Ω(dep.Checkout()).Should(BeNil())
			}
		})
	})

	Describe("Snapshot", func() {

		var newDeps map[string]*Dependency

		It("Scan and find deps", func() {
			_new, err := snapshot(file, base, false, false)
			Ω(err).Should(BeNil())
			Ω(_new).Should(HaveLen(len(srlt.Deps)))
			newDeps = _new
		})

		It("Compare scanned and from file", func() {
			for name, dep := range newDeps {
				d, ok := srlt.Deps[name]
				Ω(ok).Should(BeTrue())
				Ω(d.Name).Should(Equal(dep.Name))
				Ω(d.Type).Should(Equal(dep.Type))
				Ω(d.Remote).Should(Equal(dep.Remote))
				Ω(d.Commit).Should(Equal(dep.Commit))
			}
		})
	})

	Describe("Validate", func() {
		It("fixteres", func() {
			var depsFixtures = []struct {
				Dep     Dependency
				IsValid bool
			}{
				{Dependency{Name: "", Type: "no valid", Commit: "", Remote: ""}, false},
				{Dependency{Name: "srlt", Type: "git", Commit: "527b9bb611009a567eaf0f47c6b59c301a71e20b", Remote: "git@github.com:olebedev/srlt.git"}, true},
				{Dependency{Name: "rest", Type: "hg", Commit: "", Remote: "git@github.com:olebedev/srlt.git"}, false},
				{Dependency{Name: "snake_case", Type: "bzr", Commit: "5", Remote: ""}, false},
				{Dependency{Name: "launchpad.net/~niemeyer/goyaml/beta", Type: "bzr", Commit: "4", Remote: "launchpad.net/~niemeyer/goyaml/beta"}, true},
			}

			for _, item := range depsFixtures {
				Ω(item.Dep.Validate() == nil).Should(Equal(item.IsValid))
			}
		})
	})
})
