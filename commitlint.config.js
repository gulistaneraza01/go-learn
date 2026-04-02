// commitlint.config.js
export default {
  extends: ["@commitlint/config-conventional"],

  rules: {
    // --- Type ---
    "type-enum": [
      2,
      "always",
      [
        "feat", // New feature
        "fix", // Bug fix
        "docs", // Documentation only
        "style", // Formatting, missing semicolons, etc. (no logic change)
        "refactor", // Code change that's neither a fix nor a feature
        "perf", // Performance improvement
        "test", // Adding or correcting tests
        "build", // Build system or external dependencies (e.g. webpack, npm)
        "ci", // CI/CD configuration (e.g. GitHub Actions, CircleCI)
        "chore", // Maintenance tasks (e.g. updating .gitignore)
        "revert", // Reverts a previous commit
      ],
    ],
    "type-case": [2, "always", "lower-case"],
    "type-empty": [2, "never"],

    // --- Scope ---
    "scope-case": [2, "always", "lower-case"],
    // Optional: enforce a fixed list of scopes for monorepos
    // "scope-enum": [2, "always", ["api", "web", "mobile", "shared"]],

    // --- Subject ---
    "subject-empty": [2, "never"],
    "subject-full-stop": [2, "never", "."], // No trailing period
    "subject-case": [
      2,
      "never",
      ["sentence-case", "start-case", "pascal-case", "upper-case"],
    ],
    "subject-min-length": [2, "always", 10],
    "subject-max-length": [2, "always", 72],

    // --- Body ---
    "body-leading-blank": [2, "always"], // Blank line before body
    "body-max-line-length": [2, "always", 100],

    // --- Footer ---
    "footer-leading-blank": [2, "always"], // Blank line before footer
    "footer-max-line-length": [2, "always", 100],

    // --- Header ---
    "header-max-length": [2, "always", 72],
  },

  // Optional: custom prompt config for @commitlint/prompt-cli or czg
  prompt: {
    questions: {
      type: {
        description: "Select the type of change you're committing",
      },
      scope: {
        description:
          "What is the scope of this change (e.g. component or file name)",
      },
      subject: {
        description:
          "Write a short, imperative tense description of the change",
      },
      body: {
        description: "Provide a longer description of the change (optional)",
      },
      breaking: {
        description: "Describe any breaking changes (optional)",
      },
      issues: {
        description:
          "Add issue references, e.g. 'fix #123', 'closes #456' (optional)",
      },
    },
  },
};
