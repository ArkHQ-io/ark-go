# Changelog

## 0.19.0 (2026-04-18)

Full Changelog: [v0.18.1...v0.19.0](https://github.com/ArkHQ-io/ark-go/compare/v0.18.1...v0.19.0)

### Features

* **api:** add tenantId to send ([87f7e60](https://github.com/ArkHQ-io/ark-go/commit/87f7e60b3bea70683ee3379f77db939c71969fc5))
* **internal:** support comma format in multipart form encoding ([dfb8f82](https://github.com/ArkHQ-io/ark-go/commit/dfb8f8280956cc50555308f8dd59ebb69dc9efc0))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([f8b7f54](https://github.com/ArkHQ-io/ark-go/commit/f8b7f5401325ff2bf1781754293c1e581a0eb94d))
* **client:** use correct format specifier for header serialization ([83f0257](https://github.com/ArkHQ-io/ark-go/commit/83f02579e712ffdb40afbc63f562ed4a3937790c))
* fixes for pagination and iteration, plus iter.Seq support ([1cfca47](https://github.com/ArkHQ-io/ark-go/commit/1cfca478f9ca1ef7d41e31734b24927e4c3f4889))
* prevent duplicate ? in query params ([efaa427](https://github.com/ArkHQ-io/ark-go/commit/efaa427cd24addafc85ab0a7c24e036045af3c61))


### Chores

* **ci:** skip lint on metadata-only changes ([3b11c2f](https://github.com/ArkHQ-io/ark-go/commit/3b11c2f9141798aabd576356c176efef8b5b7757))
* **ci:** skip uploading artifacts on stainless-internal branches ([84bbafd](https://github.com/ArkHQ-io/ark-go/commit/84bbafd241085c45a0b2f93fe68f51b077779bbc))
* **ci:** support opting out of skipping builds on metadata-only commits ([8c7210a](https://github.com/ArkHQ-io/ark-go/commit/8c7210a1960872833154c00b740696bce35f5fae))
* **client:** fix multipart serialisation of Default() fields ([d0a12c9](https://github.com/ArkHQ-io/ark-go/commit/d0a12c9da0f3e47dc878048d1125f5e9ea8914b6))
* **internal:** codegen related update ([44699f5](https://github.com/ArkHQ-io/ark-go/commit/44699f5720fff87010bde0747d8f41c0b68c6769))
* **internal:** codegen related update ([58093ea](https://github.com/ArkHQ-io/ark-go/commit/58093ea52bbc645a5bdaaaa8b705234b164ae695))
* **internal:** minor cleanup ([7081fcc](https://github.com/ArkHQ-io/ark-go/commit/7081fccda980946e9f6b40f1b4635e1a9973ea4a))
* **internal:** move custom custom `json` tags to `api` ([23837f9](https://github.com/ArkHQ-io/ark-go/commit/23837f92b51813054acf0f3357ab61f584d32dab))
* **internal:** support default value struct tag ([5817254](https://github.com/ArkHQ-io/ark-go/commit/58172549cc4e8b536981e361cdf91ffc3a38cf14))
* **internal:** tweak CI branches ([a8a04b8](https://github.com/ArkHQ-io/ark-go/commit/a8a04b80e88dc4ad5b157f8518edbcd1e8167f80))
* **internal:** update gitignore ([bb1a340](https://github.com/ArkHQ-io/ark-go/commit/bb1a3405ff41573c133403fedcb88ecac4f20ec0))
* **internal:** use explicit returns ([00e57a4](https://github.com/ArkHQ-io/ark-go/commit/00e57a4fcc1c71cffa750931a0387ba6633ddc2f))
* **internal:** use explicit returns in more places ([5cbbb26](https://github.com/ArkHQ-io/ark-go/commit/5cbbb2656003d3ce3929468a47be1e67da3b7c63))
* remove unnecessary error check for url parsing ([5c4a43d](https://github.com/ArkHQ-io/ark-go/commit/5c4a43dd838ff12a9e92f6b74deabeba024d41fa))
* **test:** do not count install time for mock server timeout ([e8cb037](https://github.com/ArkHQ-io/ark-go/commit/e8cb0377cd7554d58f0877767a461953f7e319a0))
* **tests:** bump steady to v0.19.4 ([9671262](https://github.com/ArkHQ-io/ark-go/commit/9671262fca7aa6c29eb9bd3b8c7ea04b0d6703b3))
* **tests:** bump steady to v0.19.5 ([2345e53](https://github.com/ArkHQ-io/ark-go/commit/2345e53e7e568c4d3a7287360d8ffbad8fb91abe))
* **tests:** bump steady to v0.19.6 ([0a20cf3](https://github.com/ArkHQ-io/ark-go/commit/0a20cf38ac67735013b4bbf564b1ce3e37bcf108))
* **tests:** bump steady to v0.19.7 ([48d4d9d](https://github.com/ArkHQ-io/ark-go/commit/48d4d9de684ac2af111295e9d05b798d08dc0add))
* **tests:** bump steady to v0.20.1 ([51762ee](https://github.com/ArkHQ-io/ark-go/commit/51762ee98c217373b2e91939680abfd4ac71458b))
* **tests:** bump steady to v0.20.2 ([e09ff01](https://github.com/ArkHQ-io/ark-go/commit/e09ff01f287247c221fcef784b863672bba1cdce))
* **tests:** bump steady to v0.22.1 ([8363ad4](https://github.com/ArkHQ-io/ark-go/commit/8363ad4901667fe2689324ff83f1ce882d49b357))
* update docs for api:"required" ([69e480e](https://github.com/ArkHQ-io/ark-go/commit/69e480eb1864d430b4c6a9ca139ac62e9fb2aa10))
* update mock server docs ([412d3e7](https://github.com/ArkHQ-io/ark-go/commit/412d3e7b706e679820b83eb68a4a04f57be39559))


### Refactors

* **tests:** switch from prism to steady ([134b963](https://github.com/ArkHQ-io/ark-go/commit/134b963c0cc77aa09a0d23f679101e9757349f48))

## 0.18.1 (2026-02-18)

Full Changelog: [v0.18.0...v0.18.1](https://github.com/ArkHQ-io/ark-go/compare/v0.18.0...v0.18.1)

### Bug Fixes

* **encoder:** correctly serialize NullStruct ([4280a68](https://github.com/ArkHQ-io/ark-go/commit/4280a68e9c08b8b46486f2c9c2719e84619f5f5b))

## 0.18.0 (2026-02-05)

Full Changelog: [v0.17.0...v0.18.0](https://github.com/ArkHQ-io/ark-go/compare/v0.17.0...v0.18.0)

### Features

* **api:** add Credentials endpoint ([9b7c497](https://github.com/ArkHQ-io/ark-go/commit/9b7c4970eb682a91a72656e168871d71120d95ba))
* **api:** add Platform webhooks ([f387b78](https://github.com/ArkHQ-io/ark-go/commit/f387b78f3b3dceedcf7bdf4bdf1107bd7204dc6a))
* **api:** endpoint updates ([2dee4ee](https://github.com/ArkHQ-io/ark-go/commit/2dee4ee2e1b6e7cd3c136119a47a720990237a66))
* **api:** standardization improvements ([ce7ca41](https://github.com/ArkHQ-io/ark-go/commit/ce7ca41fd695efd9b090ee74523d13826d014239))
* **api:** tenant usage ([3bfc6fc](https://github.com/ArkHQ-io/ark-go/commit/3bfc6fccbebb3661765d3aed8b7cc353b50506f5))

## 0.17.0 (2026-02-03)

Full Changelog: [v0.16.0...v0.17.0](https://github.com/ArkHQ-io/ark-go/compare/v0.16.0...v0.17.0)

### Features

* **api:** Add Tenants ([94427a2](https://github.com/ArkHQ-io/ark-go/commit/94427a2c99d5f052c91ad31859c97a531e5a3486))
* **api:** api update ([b084289](https://github.com/ArkHQ-io/ark-go/commit/b08428924f86a3999c6642492dc194ce3178d7ea))
* **api:** manual updates ([be47347](https://github.com/ArkHQ-io/ark-go/commit/be4734765c174f25bcdddb476a51051e3f9eb3df))
* **api:** manual updates ([6ab5f3a](https://github.com/ArkHQ-io/ark-go/commit/6ab5f3a0a0455339c890b0ff93642576617ef02c))
* **api:** manual updates ([0e66ae0](https://github.com/ArkHQ-io/ark-go/commit/0e66ae04cdb1ef6d9e81e56619da076d37546e5a))

## 0.16.0 (2026-01-30)

Full Changelog: [v0.15.0...v0.16.0](https://github.com/ArkHQ-io/ark-go/compare/v0.15.0...v0.16.0)

### Features

* **api:** api update ([444086c](https://github.com/ArkHQ-io/ark-go/commit/444086c7e3193ab8e8392b13115caf8a0841ade6))
* **api:** manual updates ([09cb13b](https://github.com/ArkHQ-io/ark-go/commit/09cb13bb456ad9bb55e5e228b513727471b5cc59))

## 0.15.0 (2026-01-30)

Full Changelog: [v0.14.0...v0.15.0](https://github.com/ArkHQ-io/ark-go/compare/v0.14.0...v0.15.0)

### Features

* **api:** api update ([edba0d9](https://github.com/ArkHQ-io/ark-go/commit/edba0d9e36c3ead8e41e025057484920eb1f0c42))
* **api:** api update ([7430518](https://github.com/ArkHQ-io/ark-go/commit/743051899b4087337117ec15805aa84ca306466e))
* **api:** manual updates ([a322131](https://github.com/ArkHQ-io/ark-go/commit/a322131bf8d8fd94ccea9d7b4ff1a6bdd753f89a))
* **api:** manual updates ([2c866d2](https://github.com/ArkHQ-io/ark-go/commit/2c866d244b91200ef1e5fa4f957aef59695423f9))
* **api:** manual updates ([cf646e1](https://github.com/ArkHQ-io/ark-go/commit/cf646e19338cb1afabd7fe2f5a20dfd5efcedc6f))

## 0.14.0 (2026-01-29)

Full Changelog: [v0.13.0...v0.14.0](https://github.com/ArkHQ-io/ark-go/compare/v0.13.0...v0.14.0)

### Features

* **api:** add usage and SendLimit Headers ([d85c334](https://github.com/ArkHQ-io/ark-go/commit/d85c33491391046d08b00c228fc4c298de6d5b73))
* **api:** api update ([1e56852](https://github.com/ArkHQ-io/ark-go/commit/1e56852f528fdb1351cac80c4b2c34a33552b9a4))
* **api:** domain list improvement ([90bda05](https://github.com/ArkHQ-io/ark-go/commit/90bda0580bf1fc5f573248d5b9a8402603855795))


### Bug Fixes

* **docs:** fix mcp installation instructions for remote servers ([0806037](https://github.com/ArkHQ-io/ark-go/commit/08060379e4b962ad794f317e1d8cd83a31775a13))

## 0.13.0 (2026-01-25)

Full Changelog: [v0.12.0...v0.13.0](https://github.com/ArkHQ-io/ark-go/compare/v0.12.0...v0.13.0)

### Features

* **api:** manual updates ([4b20d24](https://github.com/ArkHQ-io/ark-go/commit/4b20d24cd50227ee8227be8a4ce678d573888b53))
* **api:** update email details to include attachments ([ee24249](https://github.com/ArkHQ-io/ark-go/commit/ee24249404df578645aca958d833f0eeb2a0bc61))
* **client:** add a convenient param.SetJSON helper ([bb784b4](https://github.com/ArkHQ-io/ark-go/commit/bb784b40fa959977242993f72f71313733145ba2))

## 0.12.0 (2026-01-23)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/ArkHQ-io/ark-go/compare/v0.11.0...v0.12.0)

### Features

* **api:** fix from in Send raw MIME email ([b28cc4e](https://github.com/ArkHQ-io/ark-go/commit/b28cc4e42d9668a9d9e2a4171af47f82fa9b7200))
* **api:** improve raw MIME error handling ([e081b2d](https://github.com/ArkHQ-io/ark-go/commit/e081b2dbf57e25e592e4ea378601e1fbc6a8e6c6))

## 0.11.0 (2026-01-23)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/ArkHQ-io/ark-go/compare/v0.10.0...v0.11.0)

### Features

* **api:** improve raw endpoint ([567dffe](https://github.com/ArkHQ-io/ark-go/commit/567dffe3d889f730071b951d73e8787107dae2fe))

## 0.10.0 (2026-01-22)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/ArkHQ-io/ark-go/compare/v0.9.0...v0.10.0)

### Features

* **api:** api update ([a85e9b9](https://github.com/ArkHQ-io/ark-go/commit/a85e9b96cf489714533f83a0a6f3f9f64eff0c49))
* **api:** api update ([570a407](https://github.com/ArkHQ-io/ark-go/commit/570a4079dc4f67a0efce331a6b339329dfde0fce))
* **api:** fix incorrect webhook payload examples ([fd2e5f0](https://github.com/ArkHQ-io/ark-go/commit/fd2e5f03ee7154540882f29b90f7b0392a31ff60))
* **api:** manual updates ([7eb1580](https://github.com/ArkHQ-io/ark-go/commit/7eb158021562c76dbd6eb1c10166f52a3b25e7cd))

## 0.9.0 (2026-01-22)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/ArkHQ-io/ark-go/compare/v0.8.0...v0.9.0)

### Features

* **api:** improve GET delivery attempts ([74e92c4](https://github.com/ArkHQ-io/ark-go/commit/74e92c4d8d914bb632c133b5a6bd225abedc1266))

## 0.8.0 (2026-01-21)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/ArkHQ-io/ark-go/compare/v0.7.0...v0.8.0)

### Features

* **api:** add sandbox domain ([aad38db](https://github.com/ArkHQ-io/ark-go/commit/aad38db966779c461688e616fcd4af30012decb8))

## 0.7.0 (2026-01-20)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/ArkHQ-io/ark-go/compare/v0.6.0...v0.7.0)

### Features

* **api:** add webhook deliveries ([06c916b](https://github.com/ArkHQ-io/ark-go/commit/06c916b874acaccca186b1d21582a88c576a6e23))
* **api:** api update ([c974547](https://github.com/ArkHQ-io/ark-go/commit/c974547d98521ef9a78350d00ba95c1010eba265))
* **api:** api update ([dad5ae6](https://github.com/ArkHQ-io/ark-go/commit/dad5ae6c61f5cb24dd3f12769904e06b1595b874))
* **api:** manual updates ([cf57deb](https://github.com/ArkHQ-io/ark-go/commit/cf57deba939ba163ada7d5271eb3d44e80957b0e))


### Bug Fixes

* **docs:** add missing pointer prefix to api.md return types ([b0ba78c](https://github.com/ArkHQ-io/ark-go/commit/b0ba78c0ac8ebbb236206f1748918a2f44f68a86))


### Chores

* **internal:** update `actions/checkout` version ([d303dd4](https://github.com/ArkHQ-io/ark-go/commit/d303dd4fa3ade4f79c35a28971c4b7fc6e574dce))

## 0.6.0 (2026-01-14)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/ArkHQ-io/ark-go/compare/v0.5.0...v0.6.0)

### Features

* **api:** add metadata ([a5c91dc](https://github.com/ArkHQ-io/ark-go/commit/a5c91dcab47761c881ff344df29b14441e3a3168))

## 0.5.0 (2026-01-13)

Full Changelog: [v0.4.1...v0.5.0](https://github.com/ArkHQ-io/ark-go/compare/v0.4.1...v0.5.0)

### Features

* **api:** manual updates ([ac339f1](https://github.com/ArkHQ-io/ark-go/commit/ac339f192c84a8980d116fe40c15fef395cde975))
* **api:** manual updates ([9f9e676](https://github.com/ArkHQ-io/ark-go/commit/9f9e676cdbf3cabe0e65cfbd92f53822a04dc08c))

## 0.4.1 (2026-01-13)

Full Changelog: [v0.4.0...v0.4.1](https://github.com/ArkHQ-io/ark-go/compare/v0.4.0...v0.4.1)

### Chores

* configure new SDK language ([6d99ded](https://github.com/ArkHQ-io/ark-go/commit/6d99dedde8a84abb646b218178a197a09a9f4bc2))

## 0.4.0 (2026-01-13)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/ArkHQ-io/ark-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** manual updates ([fc9184c](https://github.com/ArkHQ-io/ark-go/commit/fc9184c26a0bd99e15d286059d45114b6ac34b64))

## 0.3.0 (2026-01-13)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/ArkHQ-io/ark-go/compare/v0.2.0...v0.3.0)

### Features

* **api:** api update ([837dd58](https://github.com/ArkHQ-io/ark-go/commit/837dd58ac46d162da07186e57e94307a1b9f93fd))
* **api:** api update ([dcb6ba7](https://github.com/ArkHQ-io/ark-go/commit/dcb6ba7bac9e458510858f01b694016a069327dc))
* **api:** api update ([a832ace](https://github.com/ArkHQ-io/ark-go/commit/a832acec44b591ca59a88593bb2e06cba0e998e8))
* **api:** api update ([ec78459](https://github.com/ArkHQ-io/ark-go/commit/ec7845961af43ebdacb5ec81f9f76a3b19212343))
* **api:** manual updates ([875cccb](https://github.com/ArkHQ-io/ark-go/commit/875cccb3ecb1bc3f012886ea5ba68a2b62148715))
* **api:** manual updates ([9d442ed](https://github.com/ArkHQ-io/ark-go/commit/9d442edafbc15f80506eb937390f726cd6be347d))
* **api:** manual updates ([83ed8db](https://github.com/ArkHQ-io/ark-go/commit/83ed8db84b58389416cde7d646e33739f7266809))
* **api:** manual updates ([c929a68](https://github.com/ArkHQ-io/ark-go/commit/c929a684b3d6da8827ba4fe8fca6a53c3d200784))
* **api:** manual updates ([cd77672](https://github.com/ArkHQ-io/ark-go/commit/cd77672504c9217d8195b62e6eb4df8b671f3071))
* **api:** manual updates ([14ee31e](https://github.com/ArkHQ-io/ark-go/commit/14ee31eb772e139bd174d5ec850658f2e843e08e))
* **api:** manual updates ([05a73d4](https://github.com/ArkHQ-io/ark-go/commit/05a73d4ecd56c7a5b4d6742cc6027a0cff4e4d05))
* **api:** manual updates ([e1ff71e](https://github.com/ArkHQ-io/ark-go/commit/e1ff71ecdd2c0b0436a0b566f2bc2afb27ce4cf2))
* **api:** manual updates ([fb14f7f](https://github.com/ArkHQ-io/ark-go/commit/fb14f7f876a340d1143a7cba51efaf3e1a0f7be5))
* **api:** manual updates ([b5426d2](https://github.com/ArkHQ-io/ark-go/commit/b5426d2e0c9e1777ccbc5071d3a2660cfd87dd34))
* **api:** manual updates ([1d98da0](https://github.com/ArkHQ-io/ark-go/commit/1d98da037a83ce09719bbc0597cb9710ad683765))

## 0.2.0 (2026-01-12)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/ArkHQ-io/ark-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** api update ([0e2b8ff](https://github.com/ArkHQ-io/ark-go/commit/0e2b8ffec787e316f77b1dda5820ad9341f71d9d))

## 0.1.0 (2026-01-12)

Full Changelog: [v0.0.2...v0.1.0](https://github.com/ArkHQ-io/ark-go/compare/v0.0.2...v0.1.0)

### Features

* **api:** api update ([d1fb411](https://github.com/ArkHQ-io/ark-go/commit/d1fb411ee4d3a1343c23782d4b7453486ee93f1f))
* **api:** api update ([7d424b1](https://github.com/ArkHQ-io/ark-go/commit/7d424b1b0bbb89a76617fec18bb48ebe7cd9ffb0))

## 0.0.2 (2026-01-12)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/ArkHQ-io/ark-go/compare/v0.0.1...v0.0.2)

### Chores

* configure new SDK language ([00b3850](https://github.com/ArkHQ-io/ark-go/commit/00b3850b42ba38bffd493bb7822beb3ce6c91034))
* update SDK settings ([5f2c63e](https://github.com/ArkHQ-io/ark-go/commit/5f2c63ec7b09a41d2f492a5f3d6eb00fdf8fe9bc))
