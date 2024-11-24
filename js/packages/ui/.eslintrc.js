/** @type {import("eslint").Linter.Config} */
module.exports = {
  root: true,
  extends: ["@srsim/eslint-config/react-internal.js"],
  parser: "@typescript-eslint/parser",
  parserOptions: {
    project: "./tsconfig.lint.json",
  },
};
