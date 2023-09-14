# Issue Specification

Before reporting bugs, we highly recommend utilizing the [GitHub Issues](https://github.com/openimsdk/openkf/issues) page. Please ensure to search for similar questions on the [issue board](https://github.com/openimsdk/openkf/issues) as your query might have already been addressed or is currently being resolved. For new questions, kindly submit them through the [question assistant](https://github.com/openimsdk/openkf/issues/new/choose).

When reporting bugs, include the necessary code to reproduce the problem and provide step-by-step instructions. For new feature suggestions, clearly describe the desired changes and expected behavior.

## For Committers

### Working on Pull Requests

- Committers should re-run unit tests on their code and document the test results. If necessary, conduct smoke tests.
- It is essential for committers to adhere to the [code covenant](https://github.com/openimsdk/openkf/blob/main/CONTRIBUTING.md) regarding code formatting.

### Working on Issue Management

- Assign appropriate labels to each issue and regularly review them.
- Format issue descriptions when required.
- Maintain effective communication with contributors.
- Stay updated on the status of each issue.
- Assign one or more individuals to work on resolving the issues.

## For Contributors

### Raising an Issue

developers should add full labels for issue you raised.(including **kind**(问题类型 such as,feat,bug),**Degree of importance**(重要程度 emergency，important，normal),**strategy**(处置策略 such as self ,paticipant,coperation), **progress**(进度 such as running,later),**size**(问题规模 such as L,M,XXL))

### Claiming an Issue and Making Pull Requests

- Developers can claim an issue by commenting "want to claim + issue ID."
- When creating a pull request, developers should conduct incremental testing before proceeding with the process outlined in `/contributor.md`.

**Additional Details:**

The labels used within the project are synced from the [openim-yaml.yaml](https://github.com/kubecub/github-label-syncer/blob/main/labels-templates/openim-yaml.yaml) file. The [github-label-syncer](https://github.com/kubecub/github-label-syncer) project automatically validates and syncs these labels with the OpenIM-Server repository. A script is employed to periodically retrieve or sync the labels from OpenIM-Server using the actions mechanism. This synchronization process is incremental and also takes into consideration modifications. When adding labels to OpenKF, existing labels are not removed.

The labeling process is primarily automated through robots, with minimal manual intervention for label selection. In the future, the powerful Prow system will be leveraged to further enhance community management tasks and enable a high degree of automation.
