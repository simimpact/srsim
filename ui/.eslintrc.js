/* eslint-disable */
module.exports = {
  root: true,
  env: {
    browser: true,
    node: true,
    es6: true,
  },

  parser: "@typescript-eslint/parser",
  parserOptions: {
    ecmaVersion: "latest",
    sourceType: "module",
    project: ["packages/web/tsconfig.json"],
    tsconfigRootDir: __dirname,
    createDefaultProgram: true,
  },
  ignorePatterns: ["**/dist/*", "**/protos.d.ts"],

  plugins: ["@typescript-eslint", "import", "react", "react-hooks", "workspaces"],
  extends: [
    "eslint:recommended",
    "plugin:import/recommended",
    "plugin:import/typescript",
    "plugin:@typescript-eslint/recommended",
    "plugin:@typescript-eslint/recommended-requiring-type-checking",
    "plugin:@typescript-eslint/strict",
    "plugin:react/recommended",
    "plugin:react-hooks/recommended",
    "plugin:workspaces/recommended",
    "plugin:prettier/recommended",
  ],
  settings: {
    react: { version: "detect" },
    "import/parsers": {
      "@typescript-eslint/parser": [".ts", ".tsx"],
    },
    "import/resolver": {
      typescript: {
        project: ["packages/web/tsconfig.json"],
      },
    },
  },

  rules: {
    semi: "error",
    "linebreak-style": ["error", "unix"],

    "import/order": [
      "error",
      {
        groups: ["builtin", "external", "internal", "parent", "sibling", "index", "object"],
        "newlines-between": "never",
        alphabetize: {
          order: "asc",
          caseInsensitive: true,
        },
      },
    ],
    "import/default": "off",
    "import/no-named-as-default": "off",
    "import/no-named-as-default-member": "off",

    "react/react-in-jsx-scope": "off",

    "no-unused-vars": "off",
    "@typescript-eslint/no-unused-vars": ["error", { ignoreRestSiblings: true }],
    "@typescript-eslint/explicit-function-return-type": ["off"],
    "@typescript-eslint/no-empty-function": ["off"],
    "@typescript-eslint/no-explicit-any": ["off"],
    "@typescript-eslint/no-empty-interface": ["off"],

    "prettier/prettier": ["error", {}, { usePrettierrc: true }],
  },
};
