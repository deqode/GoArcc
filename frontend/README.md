# Alfred UI

## Getting Started

- Install **nodejs** & **npm**
- First, run the development server:

```bash
cd frontend/
npm install
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

- To build for **Production**

```bash
 npm run build 
```

- To Start server for **Production**

```bash
 npm start 
```

- To Perform linting and type check

```bash
 npm run lint
 npm run type-check
```

- Make sure you have **.env.locals** & `SERVER, CENTRIFUGO, GRAPHQL` configured

## Major libraries & framework  used

- [next](https://nextjs.org/)
- [react](https://reactjs.org/)
- [@material-ui](http://material-ui.com/)
- [@apollo/client](http://apollographql.com/)
- [next-iron-session](https://www.npmjs.com/package/next-iron-session)

## Development style guide

- **Naming Conventions:-**
  - Every  `Next Page Component` should use `PascalCase`  
  - File Name of `Next Page Component` should follow [next naming conventions for routing](https://nextjs.org/docs/routing/introduction)
  - (check) Every Hooks must start with `use`
  - Every other custom component file name should follow `PascalCase` and name must be equal to `default export`
  - All environment constants must follow `UPPER_CASE_SNAKE_CASE`
  - All other helper files should be `camelCase` (non-component files)
  - All the folder names should be `camelCase` (without overriding the next framework)
  - All interfaces name should be `PascalCase`
- **Tsx Style Conventions:-**
  - Do not use inline styles
  - Use only `material UI styles (makeStyles)` object
  - all styles must be in `styles` folder
  - `commonStyles.ts` should contains common styles
  - (check)`pageNameStyles.ts` should contains page related styles
  - Only Component specific styles must add or override in components itself.
  - Use only Material UI components (Do not tamper by adding extra divs)
  - Use constant Color Codes for Color swatch & theming
  - Use material UI class and css handling mechanism
- **imports in an order of :-**
  - React import
  - Library imports (Alphabetical order)
  - Absolute imports from the project (Alphabetical order)
  - Relative imports (Alphabetical order)
- Don't use `any` as a type
- Avoid setting of state outside of hooks
- Use arrow function wherever possible
- Use Custom Hooks to make code more precise (for subscription , lazyQuery, multiple fetching  etc )
- Use interface & Define in same file (where it is used ) or in `interfaces.ts` if commonly used
- Use `next/image` instead of `<img/>`
- Use `next/link` instead of `<a/>`
- Create individual files for `gql queries`
- Use hashmaps instead of switch case
- Use enums instead of hard-coded values
- Use axios for fetching instead of fetch
- use `Bearer header` in `gql, rest, centrifuge` middleware for **Authentication**
- Use the DRY principle (Don't repeat yourself).
  - Re-use components
- Create multiple files instead of writing a big file
- Review your code before creating a pull request or committing (use husky)
- api folder (call rest and gql apis)
  - rest & gql sub-folders
- Use useReducer when useState becomes complex

## Testing

## Folder structure
