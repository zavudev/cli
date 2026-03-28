# Changelog

## 0.2.0 (2026-03-28)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/zavudev/cli/compare/v0.1.0...v0.2.0)

### Features

* set CLI flag constant values automatically where `x-stainless-const` is set ([b684b6d](https://github.com/zavudev/cli/commit/b684b6d9a959a0c26d6fd2460a5de6ef8a8a5de6))


### Bug Fixes

* fix for off-by-one error in pagination logic ([fa6a1e1](https://github.com/zavudev/cli/commit/fa6a1e13c61969a77cd4efd20bc393beead3dbf7))


### Chores

* omit full usage information when missing required CLI parameters ([5b162ac](https://github.com/zavudev/cli/commit/5b162ac682d194e419e67ffe569080986c71642b))

## 0.1.0 (2026-03-27)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/zavudev/cli/compare/v0.0.1...v0.1.0)

### Features

* add `--max-items` flag for paginated/streaming endpoints ([d099d9a](https://github.com/zavudev/cli/commit/d099d9ace57362e7db34dc382ed80090a87b8857))
* add default description for enum CLI flags without an explicit description ([a21c0ef](https://github.com/zavudev/cli/commit/a21c0ef64234d264c67e022032cb025152aecb13))
* **api:** api update ([f0cbc5c](https://github.com/zavudev/cli/commit/f0cbc5c1f699d371df7b3aee5dd97057caed77f2))
* **api:** api update ([5c66648](https://github.com/zavudev/cli/commit/5c666482a7f36db7d72990ceb8e123f438e23b7c))
* **api:** api update ([f5e3ad8](https://github.com/zavudev/cli/commit/f5e3ad8d0836bf56327e41950912d4d55f97c289))
* **api:** api update ([71fffa2](https://github.com/zavudev/cli/commit/71fffa24cae4b95e0e3f178e8634a3e8b33a5544))
* **api:** api update ([057edc2](https://github.com/zavudev/cli/commit/057edc26a692aa047fb9badde8e0dac6baac548e))
* support passing required body params through pipes ([b6b120d](https://github.com/zavudev/cli/commit/b6b120d8c4cf46317514da94c9421c18cf154d4e))


### Bug Fixes

* avoid printing usage errors twice ([a19a61b](https://github.com/zavudev/cli/commit/a19a61b57771aeb9ee9bd0650b17f516d49101fa))
* avoid reading from stdin unless request body is form encoded or json ([79e85e8](https://github.com/zavudev/cli/commit/79e85e8191c783cfc88f35062f6f9faee2b19cc1))
* better support passing client args in any position ([4eead6e](https://github.com/zavudev/cli/commit/4eead6ea5da5cc2052eaf1ebc019b29895026ddb))
* cli no longer hangs when stdin is attached to a pipe with empty input ([3249921](https://github.com/zavudev/cli/commit/32499219db79204cedd8f19d695edce3ef23a91a))
* fix for encoding arrays with `any` type items ([b23125c](https://github.com/zavudev/cli/commit/b23125cbb911d008cfbe9130383bced7d4d10529))
* fix for test cases with newlines in YAML and better error reporting ([30cd58a](https://github.com/zavudev/cli/commit/30cd58aedf593c4e88a79adefe9f1bb1f999a8f8))
* improve linking behavior when developing on a branch not in the Go SDK ([cfd29a6](https://github.com/zavudev/cli/commit/cfd29a6b99e787ef4530501addbb00bb1bb38606))
* improved workflow for developing on branches ([be05274](https://github.com/zavudev/cli/commit/be05274a1e236c32ece9befad881329df509a0fb))
* no longer require an API key when building on production repos ([31aa3e5](https://github.com/zavudev/cli/commit/31aa3e5d71a543f2447e9d8cb466eb5f761230e5))
* only set client options when the corresponding CLI flag or env var is explicitly set ([8a56d44](https://github.com/zavudev/cli/commit/8a56d44b9d1b9104cfd423f6a8aacb27f2405018))


### Chores

* **ci:** skip lint on metadata-only changes ([e460924](https://github.com/zavudev/cli/commit/e460924a21eafe475f58a0de22500420f0b0c328))
* **ci:** skip uploading artifacts on stainless-internal branches ([06a368d](https://github.com/zavudev/cli/commit/06a368d377f422aa3f6ef24f5911fa186d2e41eb))
* configure new SDK language ([83838d7](https://github.com/zavudev/cli/commit/83838d75526387dab0bb51096b9865f7ed64f871))
* **internal:** codegen related update ([f16a48b](https://github.com/zavudev/cli/commit/f16a48b6649401c3ba9eaecd715c86b179d7459e))
* **internal:** tweak CI branches ([0fa2136](https://github.com/zavudev/cli/commit/0fa213620bf2cfda858e1a20b28cfd07667dac2e))
* **internal:** update gitignore ([e875b04](https://github.com/zavudev/cli/commit/e875b04166543c8705c19b952917f4c87e856bd5))
