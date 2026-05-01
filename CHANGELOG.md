# Changelog

## 0.5.0 (2026-05-01)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/zavudev/cli/compare/v0.4.0...v0.5.0)

### Features

* **api:** api update ([d9200b0](https://github.com/zavudev/cli/commit/d9200b09fa86ca1c86f2ae140a2b2f22554840a9))
* **api:** api update ([0a0f73f](https://github.com/zavudev/cli/commit/0a0f73f9f8a329e1370c2787c27904752e6ef801))
* **api:** api update ([771f45e](https://github.com/zavudev/cli/commit/771f45e52871e2e8499b5477b9f19c04e4167a1d))
* **api:** api update ([ace0081](https://github.com/zavudev/cli/commit/ace00816df72353a38c8d9427988c61b21631528))
* **api:** api update ([00f5deb](https://github.com/zavudev/cli/commit/00f5deb659ce9dca56f617f547d31ef461bb0f7f))
* **api:** api update ([6126b7e](https://github.com/zavudev/cli/commit/6126b7ed48909667c2f3b2979332602b7b5b154c))
* **api:** api update ([bcb63ef](https://github.com/zavudev/cli/commit/bcb63ef6b1b4e3f5fad06d07650c34f8bd3450d3))
* **api:** manual updates ([c40c4cb](https://github.com/zavudev/cli/commit/c40c4cb409948540bb7feef81b0ae11de056e7f5))
* **cli:** add `--raw-output`/`-r` option to print raw (non-JSON) strings ([d764186](https://github.com/zavudev/cli/commit/d76418669cafbae10aedd4c7e7d42f3fa356cf20))
* **cli:** alias parameters in data with `x-stainless-cli-data-alias` ([3711e06](https://github.com/zavudev/cli/commit/3711e0654e91a0345effe9b901f8b8635fbe0c62))
* **cli:** send filename and content type when reading input from files ([82faa3c](https://github.com/zavudev/cli/commit/82faa3cc225b7e931602af05d49a1482d42d0fcd))
* support passing path and query params over stdin ([b8f994d](https://github.com/zavudev/cli/commit/b8f994dc81af29fb1526cbe06cfc7704e886b573))


### Bug Fixes

* **cli:** correctly load zsh autocompletion ([a29fdc5](https://github.com/zavudev/cli/commit/a29fdc57fec69b50adf10f5ca9d76da5267c60cc))
* flags for nullable body scalar fields are strictly typed ([f71454c](https://github.com/zavudev/cli/commit/f71454c4db5b528adceda8850a41857a80da27d7))


### Chores

* add documentation for ./scripts/link ([df7e71b](https://github.com/zavudev/cli/commit/df7e71b731758877a22074668862d865d71603eb))
* **ci:** support manually triggering release workflow ([0c7441d](https://github.com/zavudev/cli/commit/0c7441d9ef8c2da629a88ba8fc750319ca4e6225))
* **cli:** fall back to JSON when using default "explore" with non-TTY ([b9b7c60](https://github.com/zavudev/cli/commit/b9b7c6049f69268590f059883babc559076e5945))
* **cli:** switch long lists of positional args over to param structs ([c97dd87](https://github.com/zavudev/cli/commit/c97dd8703f6ca089d3c29307eab81e23bec7c1d4))
* **cli:** use `ShowJSONOpts` as argument to `formatJSON` instead of many positionals ([c0c6221](https://github.com/zavudev/cli/commit/c0c62210622f72686086063f315649eb8bc35ec4))
* **internal:** more robust bootstrap script ([873650f](https://github.com/zavudev/cli/commit/873650f94169cc8ee552d5dab9e426b6b1007fa8))

## 0.4.0 (2026-04-12)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/zavudev/cli/compare/v0.3.0...v0.4.0)

### Features

* **api:** cli tap build ([1d5807e](https://github.com/zavudev/cli/commit/1d5807e5454e32d74fb3c90df99382d50b6a4fa4))

## 0.3.0 (2026-04-12)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/zavudev/cli/compare/v0.2.0...v0.3.0)

### Features

* **api:** api update ([dc7d30f](https://github.com/zavudev/cli/commit/dc7d30f02c58d9d69f6c12e8d963e33f55b48bf9))


### Bug Fixes

* fix for failing to drop invalid module replace in link script ([fc6113e](https://github.com/zavudev/cli/commit/fc6113e2f868775e64288e4716e812faff08b924))


### Chores

* **cli:** additional test cases for `ShowJSONIterator` ([8c8aadb](https://github.com/zavudev/cli/commit/8c8aadb90efed1c30f597d638f219aef1d1b7eb8))
* **cli:** let `--format raw` be used in conjunction with `--transform` ([95966e4](https://github.com/zavudev/cli/commit/95966e4a90d431fd3d1a60adcd93d289f7a674df))

## 0.2.0 (2026-04-07)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/zavudev/cli/compare/v0.1.0...v0.2.0)

### Features

* allow `-` as value representing stdin to binary-only file parameters in CLIs ([0a374cd](https://github.com/zavudev/cli/commit/0a374cd5a6f48d846f88b6ec7a959a8746217387))
* **api:** api update ([639a685](https://github.com/zavudev/cli/commit/639a685199ce4abc12d80dcfc70e7283384ed6c7))
* better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url` ([d063382](https://github.com/zavudev/cli/commit/d063382073fd01452d325f6fcf66e30662049b28))
* binary-only parameters become CLI flags that take filenames only ([6331844](https://github.com/zavudev/cli/commit/63318441188ef6e625d7e3ec93572625d26acec4))
* set CLI flag constant values automatically where `x-stainless-const` is set ([b684b6d](https://github.com/zavudev/cli/commit/b684b6d9a959a0c26d6fd2460a5de6ef8a8a5de6))


### Bug Fixes

* fall back to main branch if linking fails in CI ([beddccf](https://github.com/zavudev/cli/commit/beddccf98406903f6658ffe1be8cc5f6bbd4a333))
* fix for off-by-one error in pagination logic ([fa6a1e1](https://github.com/zavudev/cli/commit/fa6a1e13c61969a77cd4efd20bc393beead3dbf7))
* fix quoting typo ([c49dac8](https://github.com/zavudev/cli/commit/c49dac8f9bea1c7634ff188ceb3b842635a64bb2))
* handle empty data set using `--format explore` ([b05d221](https://github.com/zavudev/cli/commit/b05d221fc141ddb84cdbafb6415c27b21a8ce855))
* use `RawJSON` when iterating items with `--format explore` in the CLI ([3063e5b](https://github.com/zavudev/cli/commit/3063e5bfe65a2e4c6a35ef9fa9e1f91f7855fb82))


### Chores

* mark all CLI-related tests in Go with `t.Parallel()` ([bb98431](https://github.com/zavudev/cli/commit/bb98431d96a944ef8c0f521ecd3d4f27a699663f))
* modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary ([8e3dd42](https://github.com/zavudev/cli/commit/8e3dd420527897288802a3d70b964abdb1c75a02))
* omit full usage information when missing required CLI parameters ([5b162ac](https://github.com/zavudev/cli/commit/5b162ac682d194e419e67ffe569080986c71642b))
* switch some CLI Go tests from `os.Chdir` to `t.Chdir` ([17e4d2c](https://github.com/zavudev/cli/commit/17e4d2c67716786540c6ec23a718e280b55f84c2))

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
