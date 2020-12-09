# FASTEN Reporter
This reporter intends to format all data collected by the analysis phase and formats it to a [kafka topic](#kafka-topic-format-message)

## Development proposal
- Query the Dgraph to collect metadata from `datanodes` connected to a `infonode` of a `filenode`/`packadenode`.
- Format it as `json` following the proposed kafka topic example below.

## Kafka Topic format message

### `fasten.qmstr.*`

**Input:** `fasten.RepoCloner.out`

<!-- In the future the description should be something like: The Compliance Analyzer plug-in (QMSTR) consumes the repository path from the `Repo Cloner`, copy the source files to its cluster,  and starts QMSTR process by: ... -->
**Description:** The Compliance Analyser plug-in (QMSTR) consumes the repository URL from the `Repo Cloner`, downloads the source files, and starts the QMSTR process by: building the project and creating a build graph, analyzing the build graph to find license and compliance metadata, and reporting all findings back to the Kafka topic (`fasten.qmstr.out`) to finally insert its information into the [Metadata database](https://github.com/fasten-project/fasten/wiki/Metadata-Database-Schema).


Example output message:

<!-- Check with FASTEN partners about our pattern here: Should we call the plugin `qmstr` or should we call it `ComplianceAnalyzer` (just like in our java codebase)? -->
```json
{
    "plugin_name": "QMSTR", 
    "plugin_version": "0.0.1",
      "input":{
          "product":"org.project.compilers",
          "forge":"mvn",
          "version":"4.12"
      },
    "created_at": 1594820036478, 
    "payload":{
        "package_metadata":{
            "licenses":["BSD-3-Clause", "Apache 2.0"],
            "authors":["Frances Allen", "Grace Hopper"],
        },
        "file_metadata":[
            {
                "filename":"path/to/file.java",
                "SHA-1":"70e133a8ab344227e2f7b218de1b9a0d2e1cc2a1 ",
                "licenses":["Apache 2.0"]
            },
            {
                "filename":"path/to/file.class",
                "SHA-1":"8adaac85a044dfc4c7e3ce457020003427daac81 ",
                "licenses":["Apache 2.0"]
            },
            {
                ...
            }
        ]
    },
}
```
<!-- This error message example should be update with some real error message from QMSTR plugin -->
Example error message
```json
{
    "plugin_name": "QMSTR", 
    "plugin_version": "0.0.1",
      "input":{
          "product":"org.project.compilers",
          "forge":"mvn",
          "version":"4.12"
      },
    "created_at": 1594820036478, 
    "error":{
        "message":"git@github.com:knowm/XChart.git: remote hung up unexpectedly",
        "stacktrace":[
         "org.eclipse.jgit.api.FetchCommand.call(FetchCommand.java:222)",
         "org.eclipse.jgit.api.CloneCommand.fetch(CloneCommand.java:292)",
         "org.eclipse.jgit.api.CloneCommand.call(CloneCommand.java:176)",
         "eu.fasten.analyzer.repoclonerplugin.utils.GitCloner.cloneRepo(GitCloner.java:58)",
         "eu.fasten.analyzer.repoclonerplugin.RepoClonerPlugin$RepoCloner.cloneRepo(RepoClonerPlugin.java:146)",
         "eu.fasten.analyzer.repoclonerplugin.RepoClonerPlugin$RepoCloner.consume(RepoClonerPlugin.java:126)",
         "eu.fasten.server.plugins.kafka.FastenKafkaPlugin.handleConsuming(FastenKafkaPlugin.java:153)",
         "eu.fasten.server.plugins.kafka.FastenKafkaPlugin.run(FastenKafkaPlugin.java:105)",
         "java.base/java.lang.Thread.run(Thread.java:834)"
      ],
      "name":"TransportException"
    }
}
```