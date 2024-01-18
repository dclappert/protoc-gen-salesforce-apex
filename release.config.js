module.exports = {
  plugins: [
    "@semantic-release/commit-analyzer",
    '@semantic-release/github'
  ],
  branches: ["main"],
  repositoryUrl: "https://github.com/dclappert/protoc-gen-salesforce-apex.git",
};