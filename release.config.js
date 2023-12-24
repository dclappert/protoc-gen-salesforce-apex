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
  ],
  branches: ["main"],
  repositoryUrl: ".",
};