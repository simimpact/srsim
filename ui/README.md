# Getting Started

## Prerequisites

- Node 18+
- Go v1.19+
- Yarn 2
- (Recommended) If running on windows, recommend using WSL 2

To setup the UI workspace, execute the following commands:

```
git clone https://github.com/simimpact/srsim.git
cd srsim/ui
yarn install
```

### **All commands should be ran from the `ui` folder**

#### `yarn start`

Runs the srsim website in development mode. \
Open http://localhost:5173 to view it in the browser.

#### `yarn storybook`

Launches the storybook instance for testing components in isolation.\
Open http://localhost:6006 to view the storybook.

#### `yarn preview`

Build and run the srsim website in release mode. This will be an identical experience to how it will deploy. \
Open http://localhost:4173 to view it in the browser.

#### `yarn workspace @srsim/web add <pkg>`

Add a new npm package to the `@srsim/web`/`website` build.
