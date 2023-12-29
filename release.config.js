const LIBRARY_VERSION_FILENAME = "VERSION";

module.exports = {
  plugins: [
    "@semantic-release/commit-analyzer",
    [
      "@semantic-release/exec",
      {
        publishCmd: "echo ${nextRelease.version} > " + LIBRARY_VERSION_FILENAME,
      },
    ],
    '@semantic-release/github'
  ],
  branches: ["main"],
  repositoryUrl: "https://github.com/dclappert/protoc-gen-salesforce-apex.git",
};